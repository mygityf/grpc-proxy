package model

import "errors"

var (
	ErrOfNoServiceEndpoint = errors.New("endpoint nil for service")
)

const (
	GrpcProxyEndpoint = "127.0.0.1:9991"
)

// service endpoint
type ServiceEndpoint struct {
	// rpc service name
	ServiceName string
	// endpoint of the rpc service name
	Endpoint string
}
