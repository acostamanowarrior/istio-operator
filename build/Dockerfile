FROM registry.redhat.io/ubi8/ubi-minimal:latest

ARG build_type=maistra

ADD tmp/_output/bin/istio-operator /usr/local/bin/istio-operator
ADD resources/smcp-templates/ /usr/local/share/istio-operator/default-templates
ADD manifests-servicemesh /manifests
COPY resources/helm/ /usr/local/share/istio-operator/helm/

ENTRYPOINT [ "/usr/local/bin/istio-operator" ]
