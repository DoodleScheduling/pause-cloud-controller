FROM gcr.io/distroless/static:latest@sha256:3f2b64ef97bd285e36132c684e6b2ae8f2723293d09aae046196cca64251acac
WORKDIR /
COPY manager manager
USER 65532:65532

ENTRYPOINT ["/manager"]
