# Router

<br />

Simulating a router with Go programming language. A router is a device that 
connects two or more packet-switched networks or subnetworks. 
It serves two primary functions: managing traffic between these 
networks by forwarding data packets to their intended IP addresses, 
and allowing multiple devices to use the same Internet connection.

<br />

## Type

I implemented each router as follow:

```go
type Router struct {
	defaultGateWay string

	routingTable map[string]string

	metrics metrics.Metrics
}
```

And each router has two mehtod, ```RegisterIp``` and ```Next```, where register ip is used
to add a new gateway for router and next is the routing function.

## Start

```shell
go run main.go
```

### Metrics

```shell
localhost:4040/metrics
```

### Grafana

Dashboards are in _Grafana_ directory.
