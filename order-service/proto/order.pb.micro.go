// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order.proto

package proto

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

// Api Endpoints for Order service

func NewOrderEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Order service

type OrderService interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...client.CallOption) (*CommonResult, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...client.CallOption) (*CommonResult, error) {
	req := c.c.NewRequest(c.name, "Order.CreateOrder", in)
	out := new(CommonResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	CreateOrder(context.Context, *CreateOrderRequest, *CommonResult) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		CreateOrder(ctx context.Context, in *CreateOrderRequest, out *CommonResult) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) CreateOrder(ctx context.Context, in *CreateOrderRequest, out *CommonResult) error {
	return h.OrderHandler.CreateOrder(ctx, in, out)
}
