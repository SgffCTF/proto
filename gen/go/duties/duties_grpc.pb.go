// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: duties/duties.proto

package ssov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Duties_ReadDuty_FullMethodName           = "/duties.Duties/ReadDuty"
	Duties_ReadMyOwnerDuties_FullMethodName  = "/duties.Duties/ReadMyOwnerDuties"
	Duties_ReadMyTargetDuties_FullMethodName = "/duties.Duties/ReadMyTargetDuties"
	Duties_CreateDuty_FullMethodName         = "/duties.Duties/CreateDuty"
	Duties_UpdateDuty_FullMethodName         = "/duties.Duties/UpdateDuty"
	Duties_DeleteDuty_FullMethodName         = "/duties.Duties/DeleteDuty"
)

// DutiesClient is the client API for Duties service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DutiesClient interface {
	ReadDuty(ctx context.Context, in *ReadDutyRequest, opts ...grpc.CallOption) (*ReadDutyResponse, error)
	ReadMyOwnerDuties(ctx context.Context, in *ReadMyOwnerDutiesRequest, opts ...grpc.CallOption) (*ReadMyDutiesResponse, error)
	ReadMyTargetDuties(ctx context.Context, in *ReadMyTargetDutiesRequest, opts ...grpc.CallOption) (*ReadMyDutiesResponse, error)
	CreateDuty(ctx context.Context, in *CreateDutyRequest, opts ...grpc.CallOption) (*CreateDutyResponse, error)
	UpdateDuty(ctx context.Context, in *UpdateDutyRequest, opts ...grpc.CallOption) (*UpdateDutyResponse, error)
	DeleteDuty(ctx context.Context, in *DeleteDutyRequest, opts ...grpc.CallOption) (*DeleteDutyResponse, error)
}

type dutiesClient struct {
	cc grpc.ClientConnInterface
}

func NewDutiesClient(cc grpc.ClientConnInterface) DutiesClient {
	return &dutiesClient{cc}
}

func (c *dutiesClient) ReadDuty(ctx context.Context, in *ReadDutyRequest, opts ...grpc.CallOption) (*ReadDutyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadDutyResponse)
	err := c.cc.Invoke(ctx, Duties_ReadDuty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dutiesClient) ReadMyOwnerDuties(ctx context.Context, in *ReadMyOwnerDutiesRequest, opts ...grpc.CallOption) (*ReadMyDutiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadMyDutiesResponse)
	err := c.cc.Invoke(ctx, Duties_ReadMyOwnerDuties_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dutiesClient) ReadMyTargetDuties(ctx context.Context, in *ReadMyTargetDutiesRequest, opts ...grpc.CallOption) (*ReadMyDutiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadMyDutiesResponse)
	err := c.cc.Invoke(ctx, Duties_ReadMyTargetDuties_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dutiesClient) CreateDuty(ctx context.Context, in *CreateDutyRequest, opts ...grpc.CallOption) (*CreateDutyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDutyResponse)
	err := c.cc.Invoke(ctx, Duties_CreateDuty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dutiesClient) UpdateDuty(ctx context.Context, in *UpdateDutyRequest, opts ...grpc.CallOption) (*UpdateDutyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDutyResponse)
	err := c.cc.Invoke(ctx, Duties_UpdateDuty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dutiesClient) DeleteDuty(ctx context.Context, in *DeleteDutyRequest, opts ...grpc.CallOption) (*DeleteDutyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDutyResponse)
	err := c.cc.Invoke(ctx, Duties_DeleteDuty_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DutiesServer is the server API for Duties service.
// All implementations must embed UnimplementedDutiesServer
// for forward compatibility.
type DutiesServer interface {
	ReadDuty(context.Context, *ReadDutyRequest) (*ReadDutyResponse, error)
	ReadMyOwnerDuties(context.Context, *ReadMyOwnerDutiesRequest) (*ReadMyDutiesResponse, error)
	ReadMyTargetDuties(context.Context, *ReadMyTargetDutiesRequest) (*ReadMyDutiesResponse, error)
	CreateDuty(context.Context, *CreateDutyRequest) (*CreateDutyResponse, error)
	UpdateDuty(context.Context, *UpdateDutyRequest) (*UpdateDutyResponse, error)
	DeleteDuty(context.Context, *DeleteDutyRequest) (*DeleteDutyResponse, error)
	mustEmbedUnimplementedDutiesServer()
}

// UnimplementedDutiesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDutiesServer struct{}

func (UnimplementedDutiesServer) ReadDuty(context.Context, *ReadDutyRequest) (*ReadDutyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadDuty not implemented")
}
func (UnimplementedDutiesServer) ReadMyOwnerDuties(context.Context, *ReadMyOwnerDutiesRequest) (*ReadMyDutiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMyOwnerDuties not implemented")
}
func (UnimplementedDutiesServer) ReadMyTargetDuties(context.Context, *ReadMyTargetDutiesRequest) (*ReadMyDutiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMyTargetDuties not implemented")
}
func (UnimplementedDutiesServer) CreateDuty(context.Context, *CreateDutyRequest) (*CreateDutyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDuty not implemented")
}
func (UnimplementedDutiesServer) UpdateDuty(context.Context, *UpdateDutyRequest) (*UpdateDutyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDuty not implemented")
}
func (UnimplementedDutiesServer) DeleteDuty(context.Context, *DeleteDutyRequest) (*DeleteDutyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDuty not implemented")
}
func (UnimplementedDutiesServer) mustEmbedUnimplementedDutiesServer() {}
func (UnimplementedDutiesServer) testEmbeddedByValue()                {}

// UnsafeDutiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DutiesServer will
// result in compilation errors.
type UnsafeDutiesServer interface {
	mustEmbedUnimplementedDutiesServer()
}

func RegisterDutiesServer(s grpc.ServiceRegistrar, srv DutiesServer) {
	// If the following call pancis, it indicates UnimplementedDutiesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Duties_ServiceDesc, srv)
}

func _Duties_ReadDuty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadDutyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).ReadDuty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Duties_ReadDuty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).ReadDuty(ctx, req.(*ReadDutyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Duties_ReadMyOwnerDuties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMyOwnerDutiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).ReadMyOwnerDuties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Duties_ReadMyOwnerDuties_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).ReadMyOwnerDuties(ctx, req.(*ReadMyOwnerDutiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Duties_ReadMyTargetDuties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMyTargetDutiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).ReadMyTargetDuties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Duties_ReadMyTargetDuties_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).ReadMyTargetDuties(ctx, req.(*ReadMyTargetDutiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Duties_CreateDuty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDutyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).CreateDuty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Duties_CreateDuty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).CreateDuty(ctx, req.(*CreateDutyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Duties_UpdateDuty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDutyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).UpdateDuty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Duties_UpdateDuty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).UpdateDuty(ctx, req.(*UpdateDutyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Duties_DeleteDuty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDutyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).DeleteDuty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Duties_DeleteDuty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).DeleteDuty(ctx, req.(*DeleteDutyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Duties_ServiceDesc is the grpc.ServiceDesc for Duties service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Duties_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "duties.Duties",
	HandlerType: (*DutiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadDuty",
			Handler:    _Duties_ReadDuty_Handler,
		},
		{
			MethodName: "ReadMyOwnerDuties",
			Handler:    _Duties_ReadMyOwnerDuties_Handler,
		},
		{
			MethodName: "ReadMyTargetDuties",
			Handler:    _Duties_ReadMyTargetDuties_Handler,
		},
		{
			MethodName: "CreateDuty",
			Handler:    _Duties_CreateDuty_Handler,
		},
		{
			MethodName: "UpdateDuty",
			Handler:    _Duties_UpdateDuty_Handler,
		},
		{
			MethodName: "DeleteDuty",
			Handler:    _Duties_DeleteDuty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "duties/duties.proto",
}
