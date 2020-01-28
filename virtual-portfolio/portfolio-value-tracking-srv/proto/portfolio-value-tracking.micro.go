// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/portfolio-value-tracking.proto

package portfolio_value_tracking

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PortfolioValueTracking service

type PortfolioValueTrackingService interface {
	GetDailyHistory(ctx context.Context, in *Portfolio, opts ...client.CallOption) (*Portfolio, error)
	GetIntradayHistory(ctx context.Context, in *Portfolio, opts ...client.CallOption) (*Portfolio, error)
	GetPriceMovement(ctx context.Context, in *GetPriceMovementsRequest, opts ...client.CallOption) (*GetPriceMovementsResponse, error)
	ListPriceMovements(ctx context.Context, in *ListPriceMovementsRequest, opts ...client.CallOption) (*ListPriceMovementsResponse, error)
	ListValuations(ctx context.Context, in *ListValuationsRequest, opts ...client.CallOption) (*ListValuationsResponse, error)
}

type portfolioValueTrackingService struct {
	c    client.Client
	name string
}

func NewPortfolioValueTrackingService(name string, c client.Client) PortfolioValueTrackingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "portfoliovaluetracking"
	}
	return &portfolioValueTrackingService{
		c:    c,
		name: name,
	}
}

func (c *portfolioValueTrackingService) GetDailyHistory(ctx context.Context, in *Portfolio, opts ...client.CallOption) (*Portfolio, error) {
	req := c.c.NewRequest(c.name, "PortfolioValueTracking.GetDailyHistory", in)
	out := new(Portfolio)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioValueTrackingService) GetIntradayHistory(ctx context.Context, in *Portfolio, opts ...client.CallOption) (*Portfolio, error) {
	req := c.c.NewRequest(c.name, "PortfolioValueTracking.GetIntradayHistory", in)
	out := new(Portfolio)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioValueTrackingService) GetPriceMovement(ctx context.Context, in *GetPriceMovementsRequest, opts ...client.CallOption) (*GetPriceMovementsResponse, error) {
	req := c.c.NewRequest(c.name, "PortfolioValueTracking.GetPriceMovement", in)
	out := new(GetPriceMovementsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioValueTrackingService) ListPriceMovements(ctx context.Context, in *ListPriceMovementsRequest, opts ...client.CallOption) (*ListPriceMovementsResponse, error) {
	req := c.c.NewRequest(c.name, "PortfolioValueTracking.ListPriceMovements", in)
	out := new(ListPriceMovementsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfolioValueTrackingService) ListValuations(ctx context.Context, in *ListValuationsRequest, opts ...client.CallOption) (*ListValuationsResponse, error) {
	req := c.c.NewRequest(c.name, "PortfolioValueTracking.ListValuations", in)
	out := new(ListValuationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PortfolioValueTracking service

type PortfolioValueTrackingHandler interface {
	GetDailyHistory(context.Context, *Portfolio, *Portfolio) error
	GetIntradayHistory(context.Context, *Portfolio, *Portfolio) error
	GetPriceMovement(context.Context, *GetPriceMovementsRequest, *GetPriceMovementsResponse) error
	ListPriceMovements(context.Context, *ListPriceMovementsRequest, *ListPriceMovementsResponse) error
	ListValuations(context.Context, *ListValuationsRequest, *ListValuationsResponse) error
}

func RegisterPortfolioValueTrackingHandler(s server.Server, hdlr PortfolioValueTrackingHandler, opts ...server.HandlerOption) error {
	type portfolioValueTracking interface {
		GetDailyHistory(ctx context.Context, in *Portfolio, out *Portfolio) error
		GetIntradayHistory(ctx context.Context, in *Portfolio, out *Portfolio) error
		GetPriceMovement(ctx context.Context, in *GetPriceMovementsRequest, out *GetPriceMovementsResponse) error
		ListPriceMovements(ctx context.Context, in *ListPriceMovementsRequest, out *ListPriceMovementsResponse) error
		ListValuations(ctx context.Context, in *ListValuationsRequest, out *ListValuationsResponse) error
	}
	type PortfolioValueTracking struct {
		portfolioValueTracking
	}
	h := &portfolioValueTrackingHandler{hdlr}
	return s.Handle(s.NewHandler(&PortfolioValueTracking{h}, opts...))
}

type portfolioValueTrackingHandler struct {
	PortfolioValueTrackingHandler
}

func (h *portfolioValueTrackingHandler) GetDailyHistory(ctx context.Context, in *Portfolio, out *Portfolio) error {
	return h.PortfolioValueTrackingHandler.GetDailyHistory(ctx, in, out)
}

func (h *portfolioValueTrackingHandler) GetIntradayHistory(ctx context.Context, in *Portfolio, out *Portfolio) error {
	return h.PortfolioValueTrackingHandler.GetIntradayHistory(ctx, in, out)
}

func (h *portfolioValueTrackingHandler) GetPriceMovement(ctx context.Context, in *GetPriceMovementsRequest, out *GetPriceMovementsResponse) error {
	return h.PortfolioValueTrackingHandler.GetPriceMovement(ctx, in, out)
}

func (h *portfolioValueTrackingHandler) ListPriceMovements(ctx context.Context, in *ListPriceMovementsRequest, out *ListPriceMovementsResponse) error {
	return h.PortfolioValueTrackingHandler.ListPriceMovements(ctx, in, out)
}

func (h *portfolioValueTrackingHandler) ListValuations(ctx context.Context, in *ListValuationsRequest, out *ListValuationsResponse) error {
	return h.PortfolioValueTrackingHandler.ListValuations(ctx, in, out)
}