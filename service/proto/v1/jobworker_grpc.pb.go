// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: service/proto/v1/jobworker.proto

package proto

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
	JobWorker_StartJob_FullMethodName   = "/jobworker.v1.JobWorker/StartJob"
	JobWorker_StopJob_FullMethodName    = "/jobworker.v1.JobWorker/StopJob"
	JobWorker_QueryJob_FullMethodName   = "/jobworker.v1.JobWorker/QueryJob"
	JobWorker_StreamLogs_FullMethodName = "/jobworker.v1.JobWorker/StreamLogs"
)

// JobWorkerClient is the client API for JobWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobWorkerClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	StopJob(ctx context.Context, in *StopJobRequest, opts ...grpc.CallOption) (*StopJobResponse, error)
	QueryJob(ctx context.Context, in *QueryJobRequest, opts ...grpc.CallOption) (*QueryJobResponse, error)
	StreamLogs(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamLogsResponse], error)
}

type jobWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewJobWorkerClient(cc grpc.ClientConnInterface) JobWorkerClient {
	return &jobWorkerClient{cc}
}

func (c *jobWorkerClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartJobResponse)
	err := c.cc.Invoke(ctx, JobWorker_StartJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobWorkerClient) StopJob(ctx context.Context, in *StopJobRequest, opts ...grpc.CallOption) (*StopJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopJobResponse)
	err := c.cc.Invoke(ctx, JobWorker_StopJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobWorkerClient) QueryJob(ctx context.Context, in *QueryJobRequest, opts ...grpc.CallOption) (*QueryJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryJobResponse)
	err := c.cc.Invoke(ctx, JobWorker_QueryJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobWorkerClient) StreamLogs(ctx context.Context, in *StreamLogsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamLogsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &JobWorker_ServiceDesc.Streams[0], JobWorker_StreamLogs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamLogsRequest, StreamLogsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JobWorker_StreamLogsClient = grpc.ServerStreamingClient[StreamLogsResponse]

// JobWorkerServer is the server API for JobWorker service.
// All implementations must embed UnimplementedJobWorkerServer
// for forward compatibility.
type JobWorkerServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	StopJob(context.Context, *StopJobRequest) (*StopJobResponse, error)
	QueryJob(context.Context, *QueryJobRequest) (*QueryJobResponse, error)
	StreamLogs(*StreamLogsRequest, grpc.ServerStreamingServer[StreamLogsResponse]) error
	mustEmbedUnimplementedJobWorkerServer()
}

// UnimplementedJobWorkerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJobWorkerServer struct{}

func (UnimplementedJobWorkerServer) StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (UnimplementedJobWorkerServer) StopJob(context.Context, *StopJobRequest) (*StopJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopJob not implemented")
}
func (UnimplementedJobWorkerServer) QueryJob(context.Context, *QueryJobRequest) (*QueryJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryJob not implemented")
}
func (UnimplementedJobWorkerServer) StreamLogs(*StreamLogsRequest, grpc.ServerStreamingServer[StreamLogsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamLogs not implemented")
}
func (UnimplementedJobWorkerServer) mustEmbedUnimplementedJobWorkerServer() {}
func (UnimplementedJobWorkerServer) testEmbeddedByValue()                   {}

// UnsafeJobWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobWorkerServer will
// result in compilation errors.
type UnsafeJobWorkerServer interface {
	mustEmbedUnimplementedJobWorkerServer()
}

func RegisterJobWorkerServer(s grpc.ServiceRegistrar, srv JobWorkerServer) {
	// If the following call pancis, it indicates UnimplementedJobWorkerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JobWorker_ServiceDesc, srv)
}

func _JobWorker_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobWorkerServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobWorker_StartJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobWorkerServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobWorker_StopJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobWorkerServer).StopJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobWorker_StopJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobWorkerServer).StopJob(ctx, req.(*StopJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobWorker_QueryJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobWorkerServer).QueryJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobWorker_QueryJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobWorkerServer).QueryJob(ctx, req.(*QueryJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobWorker_StreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobWorkerServer).StreamLogs(m, &grpc.GenericServerStream[StreamLogsRequest, StreamLogsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type JobWorker_StreamLogsServer = grpc.ServerStreamingServer[StreamLogsResponse]

// JobWorker_ServiceDesc is the grpc.ServiceDesc for JobWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jobworker.v1.JobWorker",
	HandlerType: (*JobWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _JobWorker_StartJob_Handler,
		},
		{
			MethodName: "StopJob",
			Handler:    _JobWorker_StopJob_Handler,
		},
		{
			MethodName: "QueryJob",
			Handler:    _JobWorker_QueryJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLogs",
			Handler:       _JobWorker_StreamLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/proto/v1/jobworker.proto",
}
