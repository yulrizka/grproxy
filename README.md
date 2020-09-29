# grproxy
GRPC proxy to inspect, record, and reply requests

```
$ ./grproxy -addr "127.0.0.1:9999" -target "127.0.0.1:10000"
```

Arguments
* **addr** local address of the proxy
* **target** upstream (backend) of the GRPC server to forward to

## Example

Say that you are running a GRPC service **foo** on kubernetes port **10000**.

To debug it, you want to port forward the remote service locally

```
$ kubectl port-forward svc/strawberry-core 10000:10000
```

At this point you will have the original GRPC server running on port **10000** locally.

Start the proxy on port 9999

```
$ ./grproxy -addr ":9999" -target "127.0.0.1:10000"
```

This will open the proxy on port **9999**

In your app, you can start using the proxy

```
// example listing all methods with grpcurl
$ grpcurl -plaintext localhost:9999 list

Fooservice
grpc.reflection.v1alpha.ServerReflection
```
