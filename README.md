Project Initialize Operator
========================================

_This repository is currently undergoing active development. Functionality may be in flux_

Operator to support initializing a new project under a GitOps management pattern [OpenShift Container Platform](https://www.openshift.com/container-platform/index.html)

### Adding Defined Quota Sizes to Cluster

When the `quotaSize` attribute is defined in the `ProjectInitializeQuota` Custom Resource (CR) the operator will search for a cluster level `ProjectInitializeQuota` CR that defines a praticular quota size. This can be used to define predetermined t-shirt sizes when creating new projects (small, medium, large, etc)

```
apiVersion: redhatcop.redhat.io/v1alpha1
kind: ProjectInitializeQuota
metadata:
  name: small
spec:
  cpu: "5"
  memory: "10Gi"
  pods: "10"
```


### Run Locally (OpenShift)

Run the following steps to run the operator locally. The operator will require `cluster-admin` permissions that can be applied using the resources provided in the deploy/ folder.

Pull in dependences
```
$ export GO111MODULE=on
$ go mod vendor
```

Create the expected namespace
```
$ oc new-project project-operator
```

Add projectinitialize crd and resources to cluster
```
$ oc apply -f deploy/crds/redhatcop.redhat.io_projectinitializes_crd.yaml
$ oc apply -f deploy/service_account.yaml
$ oc apply -f deploy/role.yaml
$ oc apply -f deploy/role_binding.yaml
```

Add cluster level quota crd if needing to add defined quotas (Optional) 
```
$ oc apply -f deploy/crds/redhatcop.redhat.io_projectinitializequota_crd.yaml
```

Login to the cluster via the Service Account above
```
$ TOKEN=$(oc sa get-token project-initialize)
$ oc login --token="${TOKEN}"
```

Run Operator-SDK
```
$ operator-sdk run --local --namespace="project-operator" 
```