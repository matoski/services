// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/services/m3o/secret/proto/secret.proto

package go_micro_service_m3o_secret

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for SecretService service

func NewSecretServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SecretService service

type SecretService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
}

type secretService struct {
	c    client.Client
	name string
}

func NewSecretService(name string, c client.Client) SecretService {
	return &secretService{
		c:    c,
		name: name,
	}
}

func (c *secretService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "SecretService.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SecretService service

type SecretServiceHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
}

func RegisterSecretServiceHandler(s server.Server, hdlr SecretServiceHandler, opts ...server.HandlerOption) error {
	type secretService interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
	}
	type SecretService struct {
		secretService
	}
	h := &secretServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SecretService{h}, opts...))
}

type secretServiceHandler struct {
	SecretServiceHandler
}

func (h *secretServiceHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.SecretServiceHandler.Create(ctx, in, out)
}