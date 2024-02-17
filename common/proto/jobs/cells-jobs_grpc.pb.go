// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package jobs

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

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobServiceClient interface {
	PutJob(ctx context.Context, in *PutJobRequest, opts ...grpc.CallOption) (*PutJobResponse, error)
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error)
	DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobResponse, error)
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (JobService_ListJobsClient, error)
	PutTask(ctx context.Context, in *PutTaskRequest, opts ...grpc.CallOption) (*PutTaskResponse, error)
	PutTaskStream(ctx context.Context, opts ...grpc.CallOption) (JobService_PutTaskStreamClient, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (JobService_ListTasksClient, error)
	DeleteTasks(ctx context.Context, in *DeleteTasksRequest, opts ...grpc.CallOption) (*DeleteTasksResponse, error)
	DetectStuckTasks(ctx context.Context, in *DetectStuckTasksRequest, opts ...grpc.CallOption) (*DetectStuckTasksResponse, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) PutJob(ctx context.Context, in *PutJobRequest, opts ...grpc.CallOption) (*PutJobResponse, error) {
	out := new(PutJobResponse)
	err := c.cc.Invoke(ctx, "/jobs.JobService/PutJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error) {
	out := new(GetJobResponse)
	err := c.cc.Invoke(ctx, "/jobs.JobService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobResponse, error) {
	out := new(DeleteJobResponse)
	err := c.cc.Invoke(ctx, "/jobs.JobService/DeleteJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (JobService_ListJobsClient, error) {
	stream, err := c.cc.NewStream(ctx, &JobService_ServiceDesc.Streams[0], "/jobs.JobService/ListJobs", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobServiceListJobsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JobService_ListJobsClient interface {
	Recv() (*ListJobsResponse, error)
	grpc.ClientStream
}

type jobServiceListJobsClient struct {
	grpc.ClientStream
}

func (x *jobServiceListJobsClient) Recv() (*ListJobsResponse, error) {
	m := new(ListJobsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobServiceClient) PutTask(ctx context.Context, in *PutTaskRequest, opts ...grpc.CallOption) (*PutTaskResponse, error) {
	out := new(PutTaskResponse)
	err := c.cc.Invoke(ctx, "/jobs.JobService/PutTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) PutTaskStream(ctx context.Context, opts ...grpc.CallOption) (JobService_PutTaskStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &JobService_ServiceDesc.Streams[1], "/jobs.JobService/PutTaskStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobServicePutTaskStreamClient{stream}
	return x, nil
}

type JobService_PutTaskStreamClient interface {
	Send(*PutTaskRequest) error
	Recv() (*PutTaskResponse, error)
	grpc.ClientStream
}

type jobServicePutTaskStreamClient struct {
	grpc.ClientStream
}

func (x *jobServicePutTaskStreamClient) Send(m *PutTaskRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *jobServicePutTaskStreamClient) Recv() (*PutTaskResponse, error) {
	m := new(PutTaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (JobService_ListTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &JobService_ServiceDesc.Streams[2], "/jobs.JobService/ListTasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobServiceListTasksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JobService_ListTasksClient interface {
	Recv() (*ListTasksResponse, error)
	grpc.ClientStream
}

type jobServiceListTasksClient struct {
	grpc.ClientStream
}

func (x *jobServiceListTasksClient) Recv() (*ListTasksResponse, error) {
	m := new(ListTasksResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobServiceClient) DeleteTasks(ctx context.Context, in *DeleteTasksRequest, opts ...grpc.CallOption) (*DeleteTasksResponse, error) {
	out := new(DeleteTasksResponse)
	err := c.cc.Invoke(ctx, "/jobs.JobService/DeleteTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) DetectStuckTasks(ctx context.Context, in *DetectStuckTasksRequest, opts ...grpc.CallOption) (*DetectStuckTasksResponse, error) {
	out := new(DetectStuckTasksResponse)
	err := c.cc.Invoke(ctx, "/jobs.JobService/DetectStuckTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
// All implementations must embed UnimplementedJobServiceServer
// for forward compatibility
type JobServiceServer interface {
	PutJob(context.Context, *PutJobRequest) (*PutJobResponse, error)
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)
	DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobResponse, error)
	ListJobs(*ListJobsRequest, JobService_ListJobsServer) error
	PutTask(context.Context, *PutTaskRequest) (*PutTaskResponse, error)
	PutTaskStream(JobService_PutTaskStreamServer) error
	ListTasks(*ListTasksRequest, JobService_ListTasksServer) error
	DeleteTasks(context.Context, *DeleteTasksRequest) (*DeleteTasksResponse, error)
	DetectStuckTasks(context.Context, *DetectStuckTasksRequest) (*DetectStuckTasksResponse, error)
	mustEmbedUnimplementedJobServiceServer()
}

// UnimplementedJobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (UnimplementedJobServiceServer) PutJob(context.Context, *PutJobRequest) (*PutJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutJob not implemented")
}
func (UnimplementedJobServiceServer) GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedJobServiceServer) DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
func (UnimplementedJobServiceServer) ListJobs(*ListJobsRequest, JobService_ListJobsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}
func (UnimplementedJobServiceServer) PutTask(context.Context, *PutTaskRequest) (*PutTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutTask not implemented")
}
func (UnimplementedJobServiceServer) PutTaskStream(JobService_PutTaskStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PutTaskStream not implemented")
}
func (UnimplementedJobServiceServer) ListTasks(*ListTasksRequest, JobService_ListTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedJobServiceServer) DeleteTasks(context.Context, *DeleteTasksRequest) (*DeleteTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTasks not implemented")
}
func (UnimplementedJobServiceServer) DetectStuckTasks(context.Context, *DetectStuckTasksRequest) (*DetectStuckTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectStuckTasks not implemented")
}
func (UnimplementedJobServiceServer) mustEmbedUnimplementedJobServiceServer() {}

// UnsafeJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServiceServer will
// result in compilation errors.
type UnsafeJobServiceServer interface {
	mustEmbedUnimplementedJobServiceServer()
}

func RegisterJobServiceServer(s grpc.ServiceRegistrar, srv JobServiceServer) {
	s.RegisterService(&JobService_ServiceDesc, srv)
}

func _JobService_PutJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).PutJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.JobService/PutJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).PutJob(ctx, req.(*PutJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.JobService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJob(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.JobService/DeleteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).DeleteJob(ctx, req.(*DeleteJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_ListJobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListJobsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobServiceServer).ListJobs(m, &jobServiceListJobsServer{stream})
}

type JobService_ListJobsServer interface {
	Send(*ListJobsResponse) error
	grpc.ServerStream
}

type jobServiceListJobsServer struct {
	grpc.ServerStream
}

func (x *jobServiceListJobsServer) Send(m *ListJobsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _JobService_PutTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).PutTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.JobService/PutTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).PutTask(ctx, req.(*PutTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_PutTaskStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JobServiceServer).PutTaskStream(&jobServicePutTaskStreamServer{stream})
}

type JobService_PutTaskStreamServer interface {
	Send(*PutTaskResponse) error
	Recv() (*PutTaskRequest, error)
	grpc.ServerStream
}

type jobServicePutTaskStreamServer struct {
	grpc.ServerStream
}

func (x *jobServicePutTaskStreamServer) Send(m *PutTaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *jobServicePutTaskStreamServer) Recv() (*PutTaskRequest, error) {
	m := new(PutTaskRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _JobService_ListTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTasksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobServiceServer).ListTasks(m, &jobServiceListTasksServer{stream})
}

type JobService_ListTasksServer interface {
	Send(*ListTasksResponse) error
	grpc.ServerStream
}

type jobServiceListTasksServer struct {
	grpc.ServerStream
}

func (x *jobServiceListTasksServer) Send(m *ListTasksResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _JobService_DeleteTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).DeleteTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.JobService/DeleteTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).DeleteTasks(ctx, req.(*DeleteTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_DetectStuckTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectStuckTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).DetectStuckTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.JobService/DetectStuckTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).DetectStuckTasks(ctx, req.(*DetectStuckTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobService_ServiceDesc is the grpc.ServiceDesc for JobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jobs.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutJob",
			Handler:    _JobService_PutJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _JobService_GetJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _JobService_DeleteJob_Handler,
		},
		{
			MethodName: "PutTask",
			Handler:    _JobService_PutTask_Handler,
		},
		{
			MethodName: "DeleteTasks",
			Handler:    _JobService_DeleteTasks_Handler,
		},
		{
			MethodName: "DetectStuckTasks",
			Handler:    _JobService_DetectStuckTasks_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListJobs",
			Handler:       _JobService_ListJobs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PutTaskStream",
			Handler:       _JobService_PutTaskStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListTasks",
			Handler:       _JobService_ListTasks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cells-jobs.proto",
}

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	Control(ctx context.Context, in *CtrlCommand, opts ...grpc.CallOption) (*CtrlCommandResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Control(ctx context.Context, in *CtrlCommand, opts ...grpc.CallOption) (*CtrlCommandResponse, error) {
	out := new(CtrlCommandResponse)
	err := c.cc.Invoke(ctx, "/jobs.TaskService/Control", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	Control(context.Context, *CtrlCommand) (*CtrlCommandResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) Control(context.Context, *CtrlCommand) (*CtrlCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Control not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_Control_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CtrlCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Control(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jobs.TaskService/Control",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Control(ctx, req.(*CtrlCommand))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jobs.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Control",
			Handler:    _TaskService_Control_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cells-jobs.proto",
}
