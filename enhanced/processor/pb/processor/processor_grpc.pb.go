// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.4
// source: processor.proto

package processor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Processor_GenerateDemoMetadata_FullMethodName   = "/processor.Processor/GenerateDemoMetadata"
	Processor_UpdateMetadata_FullMethodName         = "/processor.Processor/UpdateMetadata"
	Processor_GetUpdateResult_FullMethodName        = "/processor.Processor/GetUpdateResult"
	Processor_GenerateDemoDeviceData_FullMethodName = "/processor.Processor/GenerateDemoDeviceData"
)

// ProcessorClient is the client API for Processor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessorClient interface {
	// DEMO use
	GenerateDemoMetadata(ctx context.Context, in *GenerateDemoMetadataReq, opts ...grpc.CallOption) (*GenerateDemoMetadataResp, error)
	// Benchmark use
	UpdateMetadata(ctx context.Context, in *UpdateMetadataReq, opts ...grpc.CallOption) (*UpdateMetadataResp, error)
	// Benchmark use
	GetUpdateResult(ctx context.Context, in *GetUpdateResultReq, opts ...grpc.CallOption) (*GetUpdateResultResp, error)
	// DEMO use
	GenerateDemoDeviceData(ctx context.Context, in *GenerateDemoDeviceDataReq, opts ...grpc.CallOption) (*GenerateDemoDeviceDataResp, error)
}

type processorClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessorClient(cc grpc.ClientConnInterface) ProcessorClient {
	return &processorClient{cc}
}

func (c *processorClient) GenerateDemoMetadata(ctx context.Context, in *GenerateDemoMetadataReq, opts ...grpc.CallOption) (*GenerateDemoMetadataResp, error) {
	out := new(GenerateDemoMetadataResp)
	err := c.cc.Invoke(ctx, Processor_GenerateDemoMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorClient) UpdateMetadata(ctx context.Context, in *UpdateMetadataReq, opts ...grpc.CallOption) (*UpdateMetadataResp, error) {
	out := new(UpdateMetadataResp)
	err := c.cc.Invoke(ctx, Processor_UpdateMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorClient) GetUpdateResult(ctx context.Context, in *GetUpdateResultReq, opts ...grpc.CallOption) (*GetUpdateResultResp, error) {
	out := new(GetUpdateResultResp)
	err := c.cc.Invoke(ctx, Processor_GetUpdateResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processorClient) GenerateDemoDeviceData(ctx context.Context, in *GenerateDemoDeviceDataReq, opts ...grpc.CallOption) (*GenerateDemoDeviceDataResp, error) {
	out := new(GenerateDemoDeviceDataResp)
	err := c.cc.Invoke(ctx, Processor_GenerateDemoDeviceData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorServer is the server API for Processor service.
// All implementations must embed UnimplementedProcessorServer
// for forward compatibility
type ProcessorServer interface {
	// DEMO use
	GenerateDemoMetadata(context.Context, *GenerateDemoMetadataReq) (*GenerateDemoMetadataResp, error)
	// Benchmark use
	UpdateMetadata(context.Context, *UpdateMetadataReq) (*UpdateMetadataResp, error)
	// Benchmark use
	GetUpdateResult(context.Context, *GetUpdateResultReq) (*GetUpdateResultResp, error)
	// DEMO use
	GenerateDemoDeviceData(context.Context, *GenerateDemoDeviceDataReq) (*GenerateDemoDeviceDataResp, error)
	mustEmbedUnimplementedProcessorServer()
}

// UnimplementedProcessorServer must be embedded to have forward compatible implementations.
type UnimplementedProcessorServer struct {
}

func (UnimplementedProcessorServer) GenerateDemoMetadata(context.Context, *GenerateDemoMetadataReq) (*GenerateDemoMetadataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDemoMetadata not implemented")
}
func (UnimplementedProcessorServer) UpdateMetadata(context.Context, *UpdateMetadataReq) (*UpdateMetadataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetadata not implemented")
}
func (UnimplementedProcessorServer) GetUpdateResult(context.Context, *GetUpdateResultReq) (*GetUpdateResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdateResult not implemented")
}
func (UnimplementedProcessorServer) GenerateDemoDeviceData(context.Context, *GenerateDemoDeviceDataReq) (*GenerateDemoDeviceDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDemoDeviceData not implemented")
}
func (UnimplementedProcessorServer) mustEmbedUnimplementedProcessorServer() {}

// UnsafeProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessorServer will
// result in compilation errors.
type UnsafeProcessorServer interface {
	mustEmbedUnimplementedProcessorServer()
}

func RegisterProcessorServer(s grpc.ServiceRegistrar, srv ProcessorServer) {
	s.RegisterService(&Processor_ServiceDesc, srv)
}

func _Processor_GenerateDemoMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDemoMetadataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServer).GenerateDemoMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Processor_GenerateDemoMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServer).GenerateDemoMetadata(ctx, req.(*GenerateDemoMetadataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Processor_UpdateMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetadataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServer).UpdateMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Processor_UpdateMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServer).UpdateMetadata(ctx, req.(*UpdateMetadataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Processor_GetUpdateResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpdateResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServer).GetUpdateResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Processor_GetUpdateResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServer).GetUpdateResult(ctx, req.(*GetUpdateResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Processor_GenerateDemoDeviceData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDemoDeviceDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServer).GenerateDemoDeviceData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Processor_GenerateDemoDeviceData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServer).GenerateDemoDeviceData(ctx, req.(*GenerateDemoDeviceDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Processor_ServiceDesc is the grpc.ServiceDesc for Processor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Processor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "processor.Processor",
	HandlerType: (*ProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateDemoMetadata",
			Handler:    _Processor_GenerateDemoMetadata_Handler,
		},
		{
			MethodName: "UpdateMetadata",
			Handler:    _Processor_UpdateMetadata_Handler,
		},
		{
			MethodName: "GetUpdateResult",
			Handler:    _Processor_GetUpdateResult_Handler,
		},
		{
			MethodName: "GenerateDemoDeviceData",
			Handler:    _Processor_GenerateDemoDeviceData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "processor.proto",
}
