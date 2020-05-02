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
  hard:
    cpu: "5"
    memory: "10Gi"
    pods: "10"
```


### Install (OpenShift)
The operator will require `cluster-admin` permissions that can be applied using the resources provided in the deploy/ folder.

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

### Run Locally (OpenShift)
Prerequisites:

In order to run the operator locally, you will need to meet these [prerequisites](https://github.com/operator-framework/operator-sdk#prerequisites) and then follow these [instructions](https://github.com/operator-framework/operator-sdk/blob/master/doc/user/install-operator-sdk.md#install-the-operator-sdk-cli) to install the operator-sdk.

Run the following steps to run the operator locally.

Pull in dependences
```
$ export GO111MODULE=on
$ go mod vendor
```

Login to the cluster via the Service Account shown in the above install step
```
$ TOKEN=$(oc sa get-token project-initialize)
$ oc login --token="${TOKEN}"
```

Run Operator-SDK locally
```
$ operator-sdk run --local --namespace="project-operator" 
```

### Deploy Operator (OpenShift)
Run the following command when ready to deploy the operator into cluster it will monitor

```
$ oc apply -f deploy/operator.yaml
```