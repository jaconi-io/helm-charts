# coturn

Helm chart for [coturn](https://github.com/coturn/coturn) TURN and STUN server.

## cert-manager

When using [cert-manager](https://cert-manager.io), a [Certificate](https://cert-manager.io/docs/usage/certificate/)
resource can be created by this chart:

```yaml
certificate:
  enabled: true
  dnsName: coturn.example.com
  issuerName: letsencrypt
```

The resulting TLS certificate and private key will be passed to coturn via the
`--cert` and `--pkey` flags. The settings `cert` and `pkey` will be ignored.

## Traefik

When using [Traefik](https://traefik.io), a [IngressRouteTCP](https://doc.traefik.io/traefik/routing/providers/kubernetes-crd/#kind-ingressroutetcp)
and a [IngressRouteUDP](https://doc.traefik.io/traefik/routing/providers/kubernetes-crd/#kind-ingressrouteudp)
resource can be created by this chart:

```yaml
traefik:
  enabled: true
  entryPoints:
    tcp: tcp
    tcp-tls: tcp-tls
    udp: udp
    udp-tls: udp-tls
```

When using the official Traefik Helm chart, the following (or similar)
configuration is needed for Traefik:

```yaml
ports:
  tcp:
    port: 3478
    nodePort: 30478
    expose: true
    exposedPort: 3478
    protocol: TCP
  tcp-tls:
    port: 5349
    nodePort: 30349
    expose: true
    exposedPort: 5349
    protocol: TCP
  udp:
    port: 3478
    nodePort: 30478
    expose: true
    exposedPort: 3478
    protocol: UDP
  udp-tls:
    port: 5349
    nodePort: 30349
    expose: true
    exposedPort: 5349
    protocol: UDP
```
