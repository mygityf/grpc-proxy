all:
	go build -o gRpcHelloServer internal/hello/pkg/server/server.go
	go build -o gRpcHelloClient internal/hello/pkg/client/client.go
	go build -o gRpcProxy internal/main.go
clean:
	rm -rf gRpcHelloServer gRpcHelloClient gRpcProxy