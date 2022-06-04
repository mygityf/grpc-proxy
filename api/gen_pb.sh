mkdir -p ../api-gen/hello/v1 ../api-gen/proxy/v1
protoc --go_out=plugins=grpc:../api-gen/hello/v1  ./hello.proto
protoc --go_out=plugins=grpc:../api-gen/proxy/v1  ./proxy_register.proto
