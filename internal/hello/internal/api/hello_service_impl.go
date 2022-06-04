package api

import (
	"context"
	hellopb "grpc-proxy/api-gen/hello/v1"
	"log"
)

// hello service impl
type helloServiceImpl struct {
}

// new
func NewHelloServiceImpl() *helloServiceImpl {
	return &helloServiceImpl{}
}

// hello rpc
func (h *helloServiceImpl) Hello(
	ctx context.Context, request *hellopb.HelloRequest,
) (*hellopb.HelloResponse, error) {
	log.Println("hello-server begin", request.Message)
	return &hellopb.HelloResponse{
		Message: request.Message,
		Code:    0,
	}, nil
}
