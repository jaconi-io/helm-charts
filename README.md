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

Create the `netbird-relay-secret`:

```shell
kubectl create secret generic -n netbird netbird-relay-secret --from-literal=netbird-relay-secret-key=t0pS3cr37!
```

Forward the NetBird dashboard to port `8081`:

```shell
kubectl port-forward -n netbird-dashboard service/netbird-dashboard 8081:80
```

Forward the NetBird management server to port `8082`:

```shell
kubectl port-forward -n netbird service/netbird-management 8082:80
```

Forward the NetBird signal server to port `10000`:

```shell
kubectl port-forward -n netbird service/netbird-signal 10000:80
```

Join the network by running

```shell
netbird up --admin-url http://localhost:8081 --management-url http://localhost:8082
```

Access the NetBird dashboard at [http://localhost:8081](http://localhost:8081). When redirected to Keycloak, use
`admin:admin` as your credentials.
