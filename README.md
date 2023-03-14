# helm-charts

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/jaconi)](https://artifacthub.io/packages/search?repo=jaconi)

```
helm repo add jaconi https://charts.jaconi.io
```

## Testing

Create a [kind](https://kind.sigs.k8s.io) cluster:

```
kind create cluster --config kind.yaml
```

Install [MetalLB](https://metallb.universe.tf) in the created cluster:

```
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.9/config/manifests/metallb-native.yaml
```

Determine the kind IP range:

```
docker network inspect -f '{{ .IPAM.Config }}' kind
```

Configure an IP address pool for MetalLB:

```
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

```
docker compose up --detach
```

Install the Helm charts for testing:

```
for f in */Chart.yaml; do
  chart=$(dirname $f)
  helm install --create-namespace --namespace $chart $chart $chart
done
```

After changing things, update the Helm charts:

```
for f in */Chart.yaml; do
  chart=$(dirname $f)
  helm upgrade --namespace $chart $chart $chart
done
```

## NetBird

Forward the NetBird management server to port `8081`:

```
kubectl port-forward -n netbird service/netbird-management 8081:80
```

Forward the NetBird dashboard to port `8080`:

```
kubectl port-forward -n netbird-dashboard service/netbird-dashboard 8080:80
```
