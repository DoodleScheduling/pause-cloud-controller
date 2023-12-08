/*
Copyright 2022 Doodle.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"errors"

	infrav1beta1 "github.com/doodlescheduling/cloud-autoscale-controller/api/v1beta1"
	"github.com/go-logr/logr"
	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

//+kubebuilder:rbac:groups=cloudautoscale.infra.doodle.com,resources=mongodbatlasclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cloudautoscale.infra.doodle.com,resources=mongodbatlasclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cloudautoscale.infra.doodle.com,resources=mongodbatlasclusters/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch

// MongoDBAtlasClusterReconciler reconciles a Namespace object
type MongoDBAtlasClusterReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

type MongoDBAtlasClusterReconcilerOptions struct {
	MaxConcurrentReconciles int
}

// SetupWithManager sets up the controller with the Manager.
func (r *MongoDBAtlasClusterReconciler) SetupWithManager(mgr ctrl.Manager, opts MongoDBAtlasClusterReconcilerOptions) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1beta1.MongoDBAtlasCluster{}).
		Watches(
			&corev1.Pod{},
			handler.EnqueueRequestsFromMapFunc(r.requestsForChangeBySelector),
		).
		WithOptions(controller.Options{MaxConcurrentReconciles: opts.MaxConcurrentReconciles}).
		Complete(r)
}

func (r *MongoDBAtlasClusterReconciler) requestsForChangeBySelector(ctx context.Context, o client.Object) []reconcile.Request {
	var list infrav1beta1.MongoDBAtlasClusterList
	if err := r.List(ctx, &list, client.InNamespace(o.GetNamespace())); err != nil {
		return nil
	}

	var reqs []reconcile.Request
	for _, cluster := range list.Items {
		for _, selector := range cluster.Spec.ScaleToZero {
			labelSel, err := metav1.LabelSelectorAsSelector(&selector)
			if err != nil {
				r.Log.Error(err, "can not select scaleToZero selectors")
				continue
			}

			if labelSel.Matches(labels.Set(o.GetLabels())) {
				r.Log.V(1).Info("change of referenced resource detected", "namespace", o.GetNamespace(), "name", o.GetName(), "kind", o.GetObjectKind().GroupVersionKind().Kind, "resource", cluster.GetName())
				reqs = append(reqs, reconcile.Request{NamespacedName: objectKey(&cluster)})
			}
		}
	}

	return reqs
}

func (r *MongoDBAtlasClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("namespace", req.Namespace, "name", req.Name)

	cluster := infrav1beta1.MongoDBAtlasCluster{}
	err := r.Client.Get(ctx, req.NamespacedName, &cluster)
	if err != nil {
		if kerrors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if cluster.Spec.Suspend {
		return ctrl.Result{}, nil
	}

	result, err := r.reconcile(ctx, cluster, logger)
	cluster.Status.ObservedGeneration = cluster.GetGeneration()

	if err != nil {
		logger.Error(err, "reconcile error occured")
		cluster = infrav1beta1.MongoDBAtlasClusterReady(cluster, metav1.ConditionFalse, "ReconciliationFailed", err.Error())
		r.Recorder.Event(&cluster, "Normal", "error", err.Error())
		result.Requeue = true
	}

	// Update status after reconciliation.
	if err := r.patchStatus(ctx, &cluster); err != nil {
		logger.Error(err, "unable to update status after reconciliation")
		return ctrl.Result{Requeue: true}, err
	}

	if err == nil && cluster.Spec.Interval.Duration != 0 {
		result.RequeueAfter = cluster.Spec.Interval.Duration
	}

	return result, err

}

func (r *MongoDBAtlasClusterReconciler) reconcile(ctx context.Context, cluster infrav1beta1.MongoDBAtlasCluster, logger logr.Logger) (ctrl.Result, error) {
	var atlas corev1.Secret
	if err := r.Get(ctx, types.NamespacedName{
		Name:      cluster.Spec.Secret.Name,
		Namespace: cluster.Name,
	}, &atlas); err != nil {
		return reconcile.Result{}, err
	}

	opts := atlasOptions{}
	if val, ok := atlas.Data["publicKey"]; !ok {
		return ctrl.Result{}, errors.New("publicKey not found in secret")
	} else {
		opts.PublicKey = string(val)
	}

	if val, ok := atlas.Data["privateKey"]; !ok {
		return ctrl.Result{}, errors.New("privateKey not found in secret")
	} else {
		opts.PrivateKey = string(val)
	}

	if cluster.Spec.ClusterName == "" {
		opts.ClusterName = cluster.Name
	} else {
		opts.ClusterName = cluster.Spec.ClusterName
	}

	logger = logger.WithValues("cluster", opts.ClusterName)
	suspend, err := r.hasRunningPods(ctx, cluster, logger)
	if err != nil {
		return ctrl.Result{}, err
	}

	var (
		res ctrl.Result
	)

	if suspend {
		logger.Info("make sure atlas clusters are suspended", "cluster", opts.ClusterName)
		res, err = r.suspend(ctx, logger, opts)
	} else {
		logger.Info("make sure atlas clusters are resumed", "cluster", opts.ClusterName)
		res, err = r.resume(ctx, logger, opts)
	}

	if err != nil {
		return res, err
	}

	return res, err
}

func (r *MongoDBAtlasClusterReconciler) hasRunningPods(ctx context.Context, cluster infrav1beta1.MongoDBAtlasCluster, logger logr.Logger) (bool, error) {
	if len(cluster.Spec.ScaleToZero) == 0 {
		return true, nil
	}

	for _, selector := range cluster.Spec.ScaleToZero {
		selector, err := metav1.LabelSelectorAsSelector(&selector)
		if err != nil {
			return false, err
		}

		var list corev1.PodList
		if err := r.Client.List(ctx, &list, client.InNamespace(cluster.Name), &client.MatchingLabelsSelector{Selector: selector}); err != nil {
			return false, err
		}

		//Compatibility to DoodleScheduling/k8s-pause, otherwise no pods would be sufficient
		for _, pod := range list.Items {
			if pod.Status.Phase != "Suspended" {
				return true, nil
			}
		}
	}

	return false, nil
}

type atlasOptions struct {
	infrav1beta1.MongoDBAtlasClusterSpec
	ClusterName string
	PublicKey   string
	PrivateKey  string
}

func (r *MongoDBAtlasClusterReconciler) resume(ctx context.Context, logger logr.Logger, opts atlasOptions) (ctrl.Result, error) {
	t := digest.NewTransport(opts.PublicKey, opts.PrivateKey)
	tc, err := t.Client()
	if err != nil {
		return ctrl.Result{}, err
	}

	atlas := mongodbatlas.NewClient(tc)
	cluster, _, err := atlas.Clusters.Get(ctx, opts.GroupID, opts.ClusterName)
	if err != nil {
		return ctrl.Result{}, err
	}

	if *cluster.Paused {
		logger.Info("resume atlas cluster")
		p := false
		cluster := &mongodbatlas.Cluster{
			Paused: &p,
		}

		_, res, err := atlas.Clusters.Update(ctx, opts.GroupID, opts.ClusterName, cluster)
		if err != nil {
			logger.Info("error response", "response", res)
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *MongoDBAtlasClusterReconciler) suspend(ctx context.Context, logger logr.Logger, opts atlasOptions) (ctrl.Result, error) {
	t := digest.NewTransport(opts.PublicKey, opts.PrivateKey)
	tc, err := t.Client()
	if err != nil {
		return ctrl.Result{}, err
	}

	atlas := mongodbatlas.NewClient(tc)
	cluster, _, err := atlas.Clusters.Get(ctx, opts.GroupID, opts.ClusterName)

	if err != nil {
		return ctrl.Result{}, err
	}

	if !*cluster.Paused {
		logger.Info("suspend atlas cluster")
		p := true
		cluster := &mongodbatlas.Cluster{
			Paused: &p,
		}
		_, res, err := atlas.Clusters.Update(ctx, opts.GroupID, opts.ClusterName, cluster)
		if err != nil {
			logger.Info("error response", "response", res)
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *MongoDBAtlasClusterReconciler) patchStatus(ctx context.Context, cluster *infrav1beta1.MongoDBAtlasCluster) error {
	key := client.ObjectKeyFromObject(cluster)
	latest := &infrav1beta1.MongoDBAtlasCluster{}
	if err := r.Client.Get(ctx, key, latest); err != nil {
		return err
	}

	return r.Client.Status().Patch(ctx, cluster, client.MergeFrom(latest))
}

// objectKey returns client.ObjectKey for the object.
func objectKey(object metav1.Object) client.ObjectKey {
	return client.ObjectKey{
		Namespace: object.GetNamespace(),
		Name:      object.GetName(),
	}
}
