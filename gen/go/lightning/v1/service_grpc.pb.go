// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: lightning/v1/service.proto

package lightning

import (
	context "context"
	lnrpc "github.com/lightningnetwork/lnd/lnrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	LightningService_CreateInvoice_FullMethodName          = "/lightning.v1.LightningService/CreateInvoice"
	LightningService_SendToRoute_FullMethodName            = "/lightning.v1.LightningService/SendToRoute"
	LightningService_InitializeUser_FullMethodName         = "/lightning.v1.LightningService/InitializeUser"
	LightningService_GetBalance_FullMethodName             = "/lightning.v1.LightningService/GetBalance"
	LightningService_ListTransactions_FullMethodName       = "/lightning.v1.LightningService/ListTransactions"
	LightningService_ProcessIncomingInvoice_FullMethodName = "/lightning.v1.LightningService/ProcessIncomingInvoice"
)

// LightningServiceClient is the client API for LightningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LightningServiceClient interface {
	CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*CreateInvoiceResponse, error)
	SendToRoute(ctx context.Context, in *SendToRouteRequest, opts ...grpc.CallOption) (*SendToRouteResponse, error)
	InitializeUser(ctx context.Context, in *InitializeBalanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetBalance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	ProcessIncomingInvoice(ctx context.Context, in *lnrpc.Invoice, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type lightningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLightningServiceClient(cc grpc.ClientConnInterface) LightningServiceClient {
	return &lightningServiceClient{cc}
}

func (c *lightningServiceClient) CreateInvoice(ctx context.Context, in *CreateInvoiceRequest, opts ...grpc.CallOption) (*CreateInvoiceResponse, error) {
	out := new(CreateInvoiceResponse)
	err := c.cc.Invoke(ctx, LightningService_CreateInvoice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningServiceClient) SendToRoute(ctx context.Context, in *SendToRouteRequest, opts ...grpc.CallOption) (*SendToRouteResponse, error) {
	out := new(SendToRouteResponse)
	err := c.cc.Invoke(ctx, LightningService_SendToRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningServiceClient) InitializeUser(ctx context.Context, in *InitializeBalanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LightningService_InitializeUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningServiceClient) GetBalance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, LightningService_GetBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningServiceClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, LightningService_ListTransactions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningServiceClient) ProcessIncomingInvoice(ctx context.Context, in *lnrpc.Invoice, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LightningService_ProcessIncomingInvoice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LightningServiceServer is the server API for LightningService service.
// All implementations should embed UnimplementedLightningServiceServer
// for forward compatibility
type LightningServiceServer interface {
	CreateInvoice(context.Context, *CreateInvoiceRequest) (*CreateInvoiceResponse, error)
	SendToRoute(context.Context, *SendToRouteRequest) (*SendToRouteResponse, error)
	InitializeUser(context.Context, *InitializeBalanceRequest) (*emptypb.Empty, error)
	GetBalance(context.Context, *emptypb.Empty) (*GetBalanceResponse, error)
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	ProcessIncomingInvoice(context.Context, *lnrpc.Invoice) (*emptypb.Empty, error)
}

// UnimplementedLightningServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLightningServiceServer struct {
}

func (UnimplementedLightningServiceServer) CreateInvoice(context.Context, *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvoice not implemented")
}
func (UnimplementedLightningServiceServer) SendToRoute(context.Context, *SendToRouteRequest) (*SendToRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToRoute not implemented")
}
func (UnimplementedLightningServiceServer) InitializeUser(context.Context, *InitializeBalanceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeUser not implemented")
}
func (UnimplementedLightningServiceServer) GetBalance(context.Context, *emptypb.Empty) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedLightningServiceServer) ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (UnimplementedLightningServiceServer) ProcessIncomingInvoice(context.Context, *lnrpc.Invoice) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessIncomingInvoice not implemented")
}

// UnsafeLightningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LightningServiceServer will
// result in compilation errors.
type UnsafeLightningServiceServer interface {
	mustEmbedUnimplementedLightningServiceServer()
}

func RegisterLightningServiceServer(s grpc.ServiceRegistrar, srv LightningServiceServer) {
	s.RegisterService(&LightningService_ServiceDesc, srv)
}

func _LightningService_CreateInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServiceServer).CreateInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LightningService_CreateInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServiceServer).CreateInvoice(ctx, req.(*CreateInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightningService_SendToRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendToRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServiceServer).SendToRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LightningService_SendToRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServiceServer).SendToRoute(ctx, req.(*SendToRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightningService_InitializeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServiceServer).InitializeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LightningService_InitializeUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServiceServer).InitializeUser(ctx, req.(*InitializeBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightningService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LightningService_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServiceServer).GetBalance(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightningService_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServiceServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LightningService_ListTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServiceServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LightningService_ProcessIncomingInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(lnrpc.Invoice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServiceServer).ProcessIncomingInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LightningService_ProcessIncomingInvoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServiceServer).ProcessIncomingInvoice(ctx, req.(*lnrpc.Invoice))
	}
	return interceptor(ctx, in, info, handler)
}

// LightningService_ServiceDesc is the grpc.ServiceDesc for LightningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LightningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lightning.v1.LightningService",
	HandlerType: (*LightningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvoice",
			Handler:    _LightningService_CreateInvoice_Handler,
		},
		{
			MethodName: "SendToRoute",
			Handler:    _LightningService_SendToRoute_Handler,
		},
		{
			MethodName: "InitializeUser",
			Handler:    _LightningService_InitializeUser_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _LightningService_GetBalance_Handler,
		},
		{
			MethodName: "ListTransactions",
			Handler:    _LightningService_ListTransactions_Handler,
		},
		{
			MethodName: "ProcessIncomingInvoice",
			Handler:    _LightningService_ProcessIncomingInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lightning/v1/service.proto",
}
