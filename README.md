Project Initialize Operator
========================================

[![Build Status](https://github.com/redhat-cop/project-initialize-operator/workflows/project-initialize-operator/badge.svg?branch=master)](https://github.com/redhat-cop/project-initialize-operator/actions?workflow=project-initialize-operator)
 [![Docker Repository on Quay](https://quay.io/repository/redhat-cop/project-initialize-operator/status "Docker Repository on Quay")](https://quay.io/repository/redhat-cop/project-initialize-operator)

_This repository is currently undergoing active development. Functionality may be in flux._

## Overview
This repository contains the Project Initialize Operator which provides functionality for creating new projects within OpenShift and triggering custom on-boarding processes, specifically around the GitOps solution [ArgoCD](https://argoproj.github.io/argo-cd/).


### Install (OpenShift)
The operator will require `cluster-admin` permissions that can be applied using the resources provided in the deploy/ folder.

Create the expected namespace
```
$ oc new-project project-operator
```

Add the `ProjectInitialize` CRD and resources to the cluster
```
$ oc apply -f deploy/service_account.yaml
$ oc apply -f deploy/role.yaml
$ oc apply -f deploy/role_binding.yaml
```

### Add ProjectInitialize CRD
#### 4.X OCP
```
$ oc apply -f deploy/crds/redhatcop.redhat.io_projectinitializes_crd.yaml
```
#### 3.X OCP
```
$ oc apply -f deploy/crds/redhatcop.redhat.io_projectinitializes_crd_3x.yaml
```

### Add ProjectInitializeQuota CRD
#### 4.X OCP
```
$ oc apply -f deploy/crds/redhatcop.redhat.io_projectinitializequota_crd.yaml
```
#### 3.X OCP
```
$ oc apply -f deploy/crds/redhatcop.redhat.io_projectinitializequota_crd_3x.yaml
```
### Deploy Operator (OpenShift)
Run the following command when ready to deploy the operator into cluster it will monitor

```
$ oc apply -f deploy/operator.yaml
```

### Namespace Labels/Annotations
Labels and annotations can be added to the namespace that is generated through the operator by specifying the values within the `ProjectInitialize` CR.

```yaml
apiVersion: redhatcop.redhat.io/v1alpha1
kind: ProjectInitialize
metadata:
  name: example-projectinitialize
spec:
  team: test
  env: dev
  cluster: clusterA
  displayName: "Test Project"
  desc: "A test project for showing the functionality of the Project Initialize Operator"
  quotaSize: small
  namespaceDetails:
    annotations:
      testKey: testValue
    labels:
      testKey: testValue
```

### Adding Defined Quota Sizes to Cluster
When the `quotaSize` attribute is defined in the `ProjectInitializeQuota` Custom Resource (CR) the operator will search for a cluster level `ProjectInitializeQuota` CR that defines a particular quota size. This can be used to define predetermined t-shirt sizes when creating new projects (small, medium, large, etc)

```yaml
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


## Example Workflow
The Project Initialize Operator will need to be running in the `project-operator` namespace before running the following example workflow.


### Apply T-Shirt Size
First start by applying the `ProjectInitializeQuota` CR that will be a global t-shirt size placeholder that the  initializer can reference when applying quotas to new projects.
```
$ oc apply -f deploy/examples/small_projectqouta_cr.yaml
```

### Apply Project Initializer
Apply the `ProjectInitialize` CR which contains details about the dev team name, cluster name, and a reference to the `ProjectInitializeQuota` which will specify the quota to assign the namespace. 

Creating a `ProjectInitialize` object will result in a new project (namespace) being created.

```
$ oc apply -f deploy/examples/basic_projectinit_cr.yaml
```

The project name will be a derivation of the `team` and `env` specified in the `ProjectInitalize` object.  The result will be `${team}-${env}`.  For example

```yaml
apiVersion: redhatcop.redhat.io/v1alpha1
kind: ProjectInitialize
metadata:
  name: phoenix-dev-projectinitialize
spec:
  team: phoenix
  env: dev
  cluster: us-west-2
  displayName: "Phoenix project for Dev environment"
  desc: "a test project for showing the functionality of the project initialize operator"
  quotaSize: small
  namespaceDetails:
    annotations:
      testkey: testValue
    labels:
      testkey: testValue
```

Will result in a namespace like this:

```
$ oc apply -f phoenix-dev.yaml
projectinitialize.redhatcop.redhat.io/phoenix-dev-projectinitialize created
$ oc get projects phoenix-dev
NAME          DISPLAY NAME                          STATUS
phoenix-dev   Phoenix project for Dev environment   Active
```

Examining the YAML definition is instructive: 

```
$ oc get projects phoenix-dev -o yaml
```

```yaml
apiVersion: project.openshift.io/v1
kind: Project
metadata:
  annotations:
    openshift.io/description: a test project for showing the functionality of the
      project initialize operator
    openshift.io/display-name: Phoenix project for Dev environment
    openshift.io/requester: system:serviceaccount:project-operator:project-initialize
    openshift.io/sa.scc.mcs: s0:c24,c9
    openshift.io/sa.scc.supplemental-groups: 1000570000/10000
    openshift.io/sa.scc.uid-range: 1000570000/10000
    testkey: testvalue
  creationTimestamp: "2020-10-01T19:07:20Z"
  labels:
    app: phoenix
    env: dev
  name: phoenix-dev
  resourceVersion: "233538"
  selfLink: /apis/project.openshift.io/v1/projects/phoenix-dev
  uid: c2ce8b0a-8354-4777-b7fb-fae08354ccb5
spec:
  finalizers:
  - kubernetes
status:
  phase: Active
```


## Development

For help with development, see [docs/development.md](docs/development.md)
