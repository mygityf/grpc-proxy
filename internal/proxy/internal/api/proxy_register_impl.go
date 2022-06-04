package api

import (
	"context"
	proxypb "grpc-proxy/api-gen/proxy/v1"
	"grpc-proxy/internal/proxy/internal/logic"
	"grpc-proxy/internal/proxy/pkg/model"
	"log"
)

// impl
type proxyRegisterImpl struct {
}

// New
func NewProxyRegisterImpl() *proxyRegisterImpl {
	return &proxyRegisterImpl{}
}

// register
func (p *proxyRegisterImpl) Register(
	ctx context.Context, request *proxypb.RegisterRequest,
) (*proxypb.RegisterResponse, error) {
	log.Println("proxy register", request.ServiceName, request.Endpoint)
	logicImpl := logic.GetRegisterImpl()
	endpoint, has := logicImpl.Get(request.ServiceName)
	if has {
		endpoint.Endpoint = request.Endpoint
	} else {
		endpoint = &model.ServiceEndpoint{
			ServiceName: request.ServiceName,
			Endpoint:    request.Endpoint,
		}
	}
	logic.GetRegisterImpl().Add(request.ServiceName, endpoint)

	return &proxypb.RegisterResponse{}, nil
}
