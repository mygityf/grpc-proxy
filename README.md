# grpc-proxy
    gRpchelloClient ----> gRpcProxy ----> gRpcHelloServer
                          ^                    |
                          |______register______|
# building
```bash
go build -o gRpcHelloServer internal/hello/pkg/server/server.go
go build -o gRpcHelloClient internal/hello/pkg/client/client.go
go build -o gRpcProxy internal/main.go
```

## grpc-proxy
* listen address 0.0.0.0:9991
  - internal/main.go
  - internal/proxy/pkg/cmd/cmd.go
  - internal/proxy/pkg/boot/boot.go
  - internal/proxy/internal/api/proxy_register_impl.go
* forward service
  - receive register request from 'gRpcHelloServer'
  - forward request of '/grpc.proxy.hello.v1.HelloService/Hello' from 'gRpchelloClient' to 'gRpcHelloServer'
  * param fullMethodName for Director like '/grpc.proxy.hello.v1.HelloService/Hello'
  * parse service name to 'grpc.proxy.hello.v1.HelloService' and forward to 'gRpcHelloServer'
* running log
```bash
grpc-proxy ./gRpcProxy      
2022/06/04 14:16:55 Listen at: 0.0.0.0:9991
2022/06/04 14:16:58 proxy register grpc.proxy.hello.v1.HelloService 127.0.0.1:9999
2022/06/04 14:17:01 proxy.forward.Director /grpc.proxy.hello.v1.HelloService/Hello grpc.proxy.hello.v1.HelloService
```
## hello server
* listen address 0.0.0.0:9999
  - internal/hello/pkg/server/server.go
  - internal/hello/pkg/cmd/cmd.go
  - internal/hello/pkg/boot/boot.go
  - internal/hello/internal/api/hello_service_impl.go
* register hello service to gRpcProxy
  - serviceName: grpc.proxy.hello.v1.HelloService
  - serviceEndpoint: 127.0.0.1:9999
* running log
```bash
grpc-proxy ./gRpcHelloServer
2022/06/04 14:16:58 Listen outer at: 0.0.0.0:9999
2022/06/04 14:17:01 hello-server begin hello

```
## hello client
* connect to gRpcProxy by 127.0.0.1:9991
* gRpcProxy forward hello request to gRpcHelloServer by 127.0.0.1:9999
  - internal/hello/pkg/client/client.go
* running log
```bash
grpc-proxy ./gRpcHelloClient
2022/06/04 14:17:01 Hello response hello
```