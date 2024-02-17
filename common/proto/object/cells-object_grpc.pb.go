// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package object

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

// ObjectsEndpointClient is the client API for ObjectsEndpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectsEndpointClient interface {
	GetMinioConfig(ctx context.Context, in *GetMinioConfigRequest, opts ...grpc.CallOption) (*GetMinioConfigResponse, error)
	StorageStats(ctx context.Context, in *StorageStatsRequest, opts ...grpc.CallOption) (*StorageStatsResponse, error)
}

type objectsEndpointClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectsEndpointClient(cc grpc.ClientConnInterface) ObjectsEndpointClient {
	return &objectsEndpointClient{cc}
}

func (c *objectsEndpointClient) GetMinioConfig(ctx context.Context, in *GetMinioConfigRequest, opts ...grpc.CallOption) (*GetMinioConfigResponse, error) {
	out := new(GetMinioConfigResponse)
	err := c.cc.Invoke(ctx, "/object.ObjectsEndpoint/GetMinioConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectsEndpointClient) StorageStats(ctx context.Context, in *StorageStatsRequest, opts ...grpc.CallOption) (*StorageStatsResponse, error) {
	out := new(StorageStatsResponse)
	err := c.cc.Invoke(ctx, "/object.ObjectsEndpoint/StorageStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectsEndpointServer is the server API for ObjectsEndpoint service.
// All implementations must embed UnimplementedObjectsEndpointServer
// for forward compatibility
type ObjectsEndpointServer interface {
	GetMinioConfig(context.Context, *GetMinioConfigRequest) (*GetMinioConfigResponse, error)
	StorageStats(context.Context, *StorageStatsRequest) (*StorageStatsResponse, error)
	mustEmbedUnimplementedObjectsEndpointServer()
}

// UnimplementedObjectsEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedObjectsEndpointServer struct {
}

func (UnimplementedObjectsEndpointServer) GetMinioConfig(context.Context, *GetMinioConfigRequest) (*GetMinioConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinioConfig not implemented")
}
func (UnimplementedObjectsEndpointServer) StorageStats(context.Context, *StorageStatsRequest) (*StorageStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageStats not implemented")
}
func (UnimplementedObjectsEndpointServer) mustEmbedUnimplementedObjectsEndpointServer() {}

// UnsafeObjectsEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectsEndpointServer will
// result in compilation errors.
type UnsafeObjectsEndpointServer interface {
	mustEmbedUnimplementedObjectsEndpointServer()
}

func RegisterObjectsEndpointServer(s grpc.ServiceRegistrar, srv ObjectsEndpointServer) {
	s.RegisterService(&ObjectsEndpoint_ServiceDesc, srv)
}

func _ObjectsEndpoint_GetMinioConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinioConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectsEndpointServer).GetMinioConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ObjectsEndpoint/GetMinioConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectsEndpointServer).GetMinioConfig(ctx, req.(*GetMinioConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectsEndpoint_StorageStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectsEndpointServer).StorageStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ObjectsEndpoint/StorageStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectsEndpointServer).StorageStats(ctx, req.(*StorageStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectsEndpoint_ServiceDesc is the grpc.ServiceDesc for ObjectsEndpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectsEndpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object.ObjectsEndpoint",
	HandlerType: (*ObjectsEndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMinioConfig",
			Handler:    _ObjectsEndpoint_GetMinioConfig_Handler,
		},
		{
			MethodName: "StorageStats",
			Handler:    _ObjectsEndpoint_StorageStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cells-object.proto",
}

// DataSourceEndpointClient is the client API for DataSourceEndpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataSourceEndpointClient interface {
	GetDataSourceConfig(ctx context.Context, in *GetDataSourceConfigRequest, opts ...grpc.CallOption) (*GetDataSourceConfigResponse, error)
}

type dataSourceEndpointClient struct {
	cc grpc.ClientConnInterface
}

func NewDataSourceEndpointClient(cc grpc.ClientConnInterface) DataSourceEndpointClient {
	return &dataSourceEndpointClient{cc}
}

func (c *dataSourceEndpointClient) GetDataSourceConfig(ctx context.Context, in *GetDataSourceConfigRequest, opts ...grpc.CallOption) (*GetDataSourceConfigResponse, error) {
	out := new(GetDataSourceConfigResponse)
	err := c.cc.Invoke(ctx, "/object.DataSourceEndpoint/GetDataSourceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataSourceEndpointServer is the server API for DataSourceEndpoint service.
// All implementations must embed UnimplementedDataSourceEndpointServer
// for forward compatibility
type DataSourceEndpointServer interface {
	GetDataSourceConfig(context.Context, *GetDataSourceConfigRequest) (*GetDataSourceConfigResponse, error)
	mustEmbedUnimplementedDataSourceEndpointServer()
}

// UnimplementedDataSourceEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedDataSourceEndpointServer struct {
}

func (UnimplementedDataSourceEndpointServer) GetDataSourceConfig(context.Context, *GetDataSourceConfigRequest) (*GetDataSourceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataSourceConfig not implemented")
}
func (UnimplementedDataSourceEndpointServer) mustEmbedUnimplementedDataSourceEndpointServer() {}

// UnsafeDataSourceEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataSourceEndpointServer will
// result in compilation errors.
type UnsafeDataSourceEndpointServer interface {
	mustEmbedUnimplementedDataSourceEndpointServer()
}

func RegisterDataSourceEndpointServer(s grpc.ServiceRegistrar, srv DataSourceEndpointServer) {
	s.RegisterService(&DataSourceEndpoint_ServiceDesc, srv)
}

func _DataSourceEndpoint_GetDataSourceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataSourceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceEndpointServer).GetDataSourceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.DataSourceEndpoint/GetDataSourceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceEndpointServer).GetDataSourceConfig(ctx, req.(*GetDataSourceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataSourceEndpoint_ServiceDesc is the grpc.ServiceDesc for DataSourceEndpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataSourceEndpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object.DataSourceEndpoint",
	HandlerType: (*DataSourceEndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDataSourceConfig",
			Handler:    _DataSourceEndpoint_GetDataSourceConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cells-object.proto",
}

// ResourceCleanerEndpointClient is the client API for ResourceCleanerEndpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceCleanerEndpointClient interface {
	CleanResourcesBeforeDelete(ctx context.Context, in *CleanResourcesRequest, opts ...grpc.CallOption) (*CleanResourcesResponse, error)
}

type resourceCleanerEndpointClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceCleanerEndpointClient(cc grpc.ClientConnInterface) ResourceCleanerEndpointClient {
	return &resourceCleanerEndpointClient{cc}
}

func (c *resourceCleanerEndpointClient) CleanResourcesBeforeDelete(ctx context.Context, in *CleanResourcesRequest, opts ...grpc.CallOption) (*CleanResourcesResponse, error) {
	out := new(CleanResourcesResponse)
	err := c.cc.Invoke(ctx, "/object.ResourceCleanerEndpoint/CleanResourcesBeforeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceCleanerEndpointServer is the server API for ResourceCleanerEndpoint service.
// All implementations must embed UnimplementedResourceCleanerEndpointServer
// for forward compatibility
type ResourceCleanerEndpointServer interface {
	CleanResourcesBeforeDelete(context.Context, *CleanResourcesRequest) (*CleanResourcesResponse, error)
	mustEmbedUnimplementedResourceCleanerEndpointServer()
}

// UnimplementedResourceCleanerEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedResourceCleanerEndpointServer struct {
}

func (UnimplementedResourceCleanerEndpointServer) CleanResourcesBeforeDelete(context.Context, *CleanResourcesRequest) (*CleanResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanResourcesBeforeDelete not implemented")
}
func (UnimplementedResourceCleanerEndpointServer) mustEmbedUnimplementedResourceCleanerEndpointServer() {
}

// UnsafeResourceCleanerEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceCleanerEndpointServer will
// result in compilation errors.
type UnsafeResourceCleanerEndpointServer interface {
	mustEmbedUnimplementedResourceCleanerEndpointServer()
}

func RegisterResourceCleanerEndpointServer(s grpc.ServiceRegistrar, srv ResourceCleanerEndpointServer) {
	s.RegisterService(&ResourceCleanerEndpoint_ServiceDesc, srv)
}

func _ResourceCleanerEndpoint_CleanResourcesBeforeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceCleanerEndpointServer).CleanResourcesBeforeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ResourceCleanerEndpoint/CleanResourcesBeforeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceCleanerEndpointServer).CleanResourcesBeforeDelete(ctx, req.(*CleanResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceCleanerEndpoint_ServiceDesc is the grpc.ServiceDesc for ResourceCleanerEndpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceCleanerEndpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object.ResourceCleanerEndpoint",
	HandlerType: (*ResourceCleanerEndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CleanResourcesBeforeDelete",
			Handler:    _ResourceCleanerEndpoint_CleanResourcesBeforeDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cells-object.proto",
}
