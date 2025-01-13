# helm-charts

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/jaconi)](https://artifacthub.io/packages/search?repo=jaconi)

```shell
helm repo add jaconi https://charts.jaconi.io
```

## Testing

Create a [kind](https://kind.sigs.k8s.io) cluster:

```shell
kind create cluster --config kind.yaml
```

Install [MetalLB](https://metallb.io) in the created cluster:

```shell
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.14.9/config/manifests/metallb-native.yaml
```

Determine the kind IP range:

```shell
docker network inspect -f '{{ .IPAM.Config }}' kind
```

Configure an IP address pool for MetalLB:

```shell
kubectl apply -f - << EOF
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: kind
  namespace: metallb-system
spec:
  addresses:
    - 172.18.255.200-172.18.255.250
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: kind
  namespace: metallb-system
spec:
  ipAddressPools:
    - kind
EOF
```

Start [Keycloak](https://www.keycloak.org):

```shell
docker compose up --detach
```

Update the Netbird configuration like this:

```yaml
auth:
  device:
    # provider: none
    provider: hosted
    audience: account
    authority: http://localhost
    clientID: netbird-management
    deviceAuthorizationEndpoint: http://localhost/auth/device
    tokenEndpoint: http://localhost/token
    scope: openid
    useIDToken: false
```

Install the Helm charts for testing:

```shell
for f in */Chart.yaml; do
  chart=$(dirname $f)
  helm install --create-namespace --namespace $chart $chart $chart
done
```

After changing things, update the Helm charts:

```shell
for f in */Chart.yaml; do
  chart=$(dirname $f)
  helm upgrade --namespace $chart $chart $chart
done
```

## NetBird

Forward the NetBird management server to port `8081`:

```shell
kubectl port-forward -n netbird service/netbird-management 8081:80
```

Forward the NetBird dashboard to port `8080`:

```shell
kubectl port-forward -n netbird-dashboard service/netbird-dashboard 8080:80
```

Forward the NetBird signal server to port `10000`:

```shell
kubectl port-forward -n netbird service/netbird-signal 10000:80
```

Join the network by running

```shell
netbird up --management-url http://localhost:8081 --admin-url http://localhost:8080
```
