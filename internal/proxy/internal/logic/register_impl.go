package logic

import (
	"context"
	"google.golang.org/grpc"
	"grpc-proxy/internal/proxy/pkg/model"
	"grpc-proxy/internal/proxy/pkg/proxy"
	"sync"
)

var (
	registerLogicImplPtr  *registerLogicImpl
	registerLogicImplOnce sync.Once
)

// impl
type registerLogicImpl struct {
	serviceCache          sync.Map
	endpointGrpcConnCache sync.Map
}

// once
func GetRegisterImpl() *registerLogicImpl {
	registerLogicImplOnce.Do(func() {
		registerLogicImplPtr = &registerLogicImpl{}
	})
	return registerLogicImplPtr
}

// add service
func (r *registerLogicImpl) Add(serviceName string, endpoint *model.ServiceEndpoint) {
	r.serviceCache.Store(serviceName, endpoint)
}

// delete service
func (r *registerLogicImpl) Delete(serviceName string) {
	r.serviceCache.Delete(serviceName)
}

// get service
func (r *registerLogicImpl) Get(serviceName string) (*model.ServiceEndpoint, bool) {
	valueObj, ok := r.serviceCache.Load(serviceName)
	if ok {
		return valueObj.(*model.ServiceEndpoint), true
	}
	return nil, false
}

// get conn
func (r *registerLogicImpl) GetConn(ctx context.Context, serviceName string) (*grpc.ClientConn, error) {
	endpoint, ok := r.Get(serviceName)
	if !ok {
		return nil, model.ErrOfNoServiceEndpoint
	}
	if valueObj, has := r.endpointGrpcConnCache.Load(endpoint.Endpoint); has {
		return valueObj.(*grpc.ClientConn), nil
	}
	// new conn
	conn, err := grpc.DialContext(ctx, endpoint.Endpoint,
		grpc.WithCodec(proxy.Codec()),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	r.endpointGrpcConnCache.Store(endpoint.Endpoint, conn)
	return conn, nil
}
