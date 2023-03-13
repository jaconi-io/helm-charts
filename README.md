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

Start [Keycloak](https://www.keycloak.org):

```
docker compose up
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
kubectl port-forward -n netbird-management service/netbird-management 8081:80
```

Forward the NetBird dashboard to port `8080`:

```
kubectl port-forward -n netbird-dashboard service/netbird-dashboard 8080:80
```
