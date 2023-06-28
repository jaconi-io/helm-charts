# koptimize

Helm chart for [koptimize](https://github.com/jaconi-io/koptimize). Note, that
while the Helm chart is open source and the Docker image is public, the code for
koptimize itself is private.

koptimize aims at reducing cost of a Kubernetes cluster, going further than, for
example, the cluster autoscaler.

## Features

- A mutating webhook, that forces workloads (that tolerate spot instances) onto
  spot instances. This is helpful, to ramp up spot instance usage. Example:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
  namespace: default
spec:
  ...
  tolerations:
    - key: cloud.google.com/gke-spot
      operator: Equal
      value: "true"
      effect: NoSchedule
```

becomes

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
  namespace: default
spec:
  ...
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
          - matchExpressions:
              - key: cloud.google.com/gke-spot
                operator: In
                values:
                  - "true"
  tolerations:
    - key: cloud.google.com/gke-spot
      operator: Equal
      value: "true"
      effect: NoSchedule
```

## Parameters

**This section is generated using metadata from [values.yaml](values.yaml)!**
