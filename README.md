Project Initialize Operator
========================================

_This repository is currently undergoing active development. Functionality may be in flux_

Operator to support initializing a new project under a GitOps management pattern [OpenShift Container Platform](https://www.openshift.com/container-platform/index.html)

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

Add crd and resources to cluster
```
$ oc apply -f deploy/crds/redhatcop.redhat.io_projectinitializes_crd.yaml
$ oc apply -f deploy/service_account.yaml
$ oc apply -f deploy/role.yaml
$ oc apply -f deploy/role_binding.yaml
```

Login to the cluster via the Service Account above
```
$ oc sa get-token project-initialize
$ oc login --token="{above_token}"
```

Run Operator-SDK
```
$ operator-sdk up local --namespace="project-operator" 
```