//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRDSInstance) DeepCopyInto(out *AWSRDSInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRDSInstance.
func (in *AWSRDSInstance) DeepCopy() *AWSRDSInstance {
	if in == nil {
		return nil
	}
	out := new(AWSRDSInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSRDSInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRDSInstanceList) DeepCopyInto(out *AWSRDSInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSRDSInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRDSInstanceList.
func (in *AWSRDSInstanceList) DeepCopy() *AWSRDSInstanceList {
	if in == nil {
		return nil
	}
	out := new(AWSRDSInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSRDSInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRDSInstanceSpec) DeepCopyInto(out *AWSRDSInstanceSpec) {
	*out = *in
	if in.ScaleToZero != nil {
		in, out := &in.ScaleToZero, &out.ScaleToZero
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.GracePeriod = in.GracePeriod
	out.Interval = in.Interval
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRDSInstanceSpec.
func (in *AWSRDSInstanceSpec) DeepCopy() *AWSRDSInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(AWSRDSInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRDSInstanceStatus) DeepCopyInto(out *AWSRDSInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRDSInstanceStatus.
func (in *AWSRDSInstanceStatus) DeepCopy() *AWSRDSInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AWSRDSInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasCluster) DeepCopyInto(out *MongoDBAtlasCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasCluster.
func (in *MongoDBAtlasCluster) DeepCopy() *MongoDBAtlasCluster {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasClusterList) DeepCopyInto(out *MongoDBAtlasClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBAtlasCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasClusterList.
func (in *MongoDBAtlasClusterList) DeepCopy() *MongoDBAtlasClusterList {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBAtlasClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasClusterSpec) DeepCopyInto(out *MongoDBAtlasClusterSpec) {
	*out = *in
	if in.ScaleToZero != nil {
		in, out := &in.ScaleToZero, &out.ScaleToZero
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.GracePeriod = in.GracePeriod
	out.Interval = in.Interval
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasClusterSpec.
func (in *MongoDBAtlasClusterSpec) DeepCopy() *MongoDBAtlasClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBAtlasClusterStatus) DeepCopyInto(out *MongoDBAtlasClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBAtlasClusterStatus.
func (in *MongoDBAtlasClusterStatus) DeepCopy() *MongoDBAtlasClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBAtlasClusterStatus)
	in.DeepCopyInto(out)
	return out
}
