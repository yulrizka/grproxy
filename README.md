# grproxy
GRPC proxy to inspect, record, and reply requests

**Current limitations**
* Only support support inspecting request and response
* Server needs to implement [reflection API](https://github.com/grpc/grpc/blob/master/doc/server-reflection.md)

## Running
```
$ ./grproxy -addr "127.0.0.1:9999" -target "127.0.0.1:10000"
```

Arguments
* **addr** local address of the proxy
* **target** upstream (backend) of the GRPC server to forward to

## Running with docker

```
$ docker run --rm -it -p 9999:9999 yulrizka/grproxy:latest -target "host.docker.internal:10000"
```

`host.docker.internal` will be translated by docker to the host IP address as a workaround
since OSX or Windows does not support `--net=host`. This is if you
bind remote GRPC server to your local machine (localhost). See example below.

## Example

Say that you are running a GRPC service **FooService** on remote kubernetes port **10000**.

To debug it, first you need to forward the grpc service port to localhost

```
$ kubectl port-forward svc/foo-service 10000:10000
```

At this point you will have the original GRPC server running on port **10000** locally.

Start the proxy on port 9999

```
$ docker run --rm -it -p 9999:9999 yulrizka/grproxy:latest -target "host.docker.internal:10000"
```

This will open the proxy on port **9999** on your local (host) machine that forward the request to `localhost:10000` which eventually
forward the request to the original (upstream) GRPC server.

In your app, you can start using the proxy as GRPC server.

```
# example listing all methods with grpcurl
$ grpcurl -plaintext localhost:9999 list

Fooservice
grpc.reflection.v1alpha.ServerReflection
```
