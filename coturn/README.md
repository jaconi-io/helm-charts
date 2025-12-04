# coturn

Helm chart for [coturn](https://github.com/coturn/coturn) TURN and STUN server.

## Testing

Install the coturn Helm chart with `--allow-loopback-peers` and `--no-auth`.

Emulate a peer by running `turnutils_peer`:

```shell
turnutils_peer -L 127.0.0.1
```

Port forward coturn:

```shell
kubectl port-forward --namespace coturn services/coturn 3478:3478
```

Emulate a client by running `turnutils_uclient`:

```shell
turnutils_uclient -t -e 127.0.0.1 -X 127.0.0.1
```

## Parameters

**This section is generated using metadata from [values.yaml](values.yaml)!**
