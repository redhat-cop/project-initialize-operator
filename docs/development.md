Project Initialize Operator - Development
========================================

### Run Locally (OpenShift)
`This should only be for development purposes`
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