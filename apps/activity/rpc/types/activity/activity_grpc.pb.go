// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.11.2
// source: rpc/activity.proto

package activity

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
	Activity_GrabPoints_FullMethodName            = "/activity.activity/GrabPoints"
	Activity_LookAllActivities_FullMethodName     = "/activity.activity/LookAllActivities"
	Activity_UserSignUpActiity_FullMethodName     = "/activity.activity/UserSignUpActiity"
	Activity_UserViewAllActivities_FullMethodName = "/activity.activity/UserViewAllActivities"
)

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	// 抢积分
	GrabPoints(ctx context.Context, in *GrabPointsRequest, opts ...grpc.CallOption) (*GrabPointsResponse, error)
	// 查看所有的活动
	LookAllActivities(ctx context.Context, in *LookAllActivitiesRequest, opts ...grpc.CallOption) (*LookAllActivitiesResponse, error)
	// 用户进行报名
	UserSignUpActiity(ctx context.Context, in *UserSignUpActivityRequest, opts ...grpc.CallOption) (*UserSignUpActivityResponse, error)
	// 用户查看自己的报名活动
	UserViewAllActivities(ctx context.Context, in *UserViewAllActivitiesRequest, opts ...grpc.CallOption) (*UserViewAllActivitiesResponse, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) GrabPoints(ctx context.Context, in *GrabPointsRequest, opts ...grpc.CallOption) (*GrabPointsResponse, error) {
	out := new(GrabPointsResponse)
	err := c.cc.Invoke(ctx, Activity_GrabPoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) LookAllActivities(ctx context.Context, in *LookAllActivitiesRequest, opts ...grpc.CallOption) (*LookAllActivitiesResponse, error) {
	out := new(LookAllActivitiesResponse)
	err := c.cc.Invoke(ctx, Activity_LookAllActivities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) UserSignUpActiity(ctx context.Context, in *UserSignUpActivityRequest, opts ...grpc.CallOption) (*UserSignUpActivityResponse, error) {
	out := new(UserSignUpActivityResponse)
	err := c.cc.Invoke(ctx, Activity_UserSignUpActiity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) UserViewAllActivities(ctx context.Context, in *UserViewAllActivitiesRequest, opts ...grpc.CallOption) (*UserViewAllActivitiesResponse, error) {
	out := new(UserViewAllActivitiesResponse)
	err := c.cc.Invoke(ctx, Activity_UserViewAllActivities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations must embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	// 抢积分
	GrabPoints(context.Context, *GrabPointsRequest) (*GrabPointsResponse, error)
	// 查看所有的活动
	LookAllActivities(context.Context, *LookAllActivitiesRequest) (*LookAllActivitiesResponse, error)
	// 用户进行报名
	UserSignUpActiity(context.Context, *UserSignUpActivityRequest) (*UserSignUpActivityResponse, error)
	// 用户查看自己的报名活动
	UserViewAllActivities(context.Context, *UserViewAllActivitiesRequest) (*UserViewAllActivitiesResponse, error)
	mustEmbedUnimplementedActivityServer()
}

// UnimplementedActivityServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) GrabPoints(context.Context, *GrabPointsRequest) (*GrabPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrabPoints not implemented")
}
func (UnimplementedActivityServer) LookAllActivities(context.Context, *LookAllActivitiesRequest) (*LookAllActivitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookAllActivities not implemented")
}
func (UnimplementedActivityServer) UserSignUpActiity(context.Context, *UserSignUpActivityRequest) (*UserSignUpActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignUpActiity not implemented")
}
func (UnimplementedActivityServer) UserViewAllActivities(context.Context, *UserViewAllActivitiesRequest) (*UserViewAllActivitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserViewAllActivities not implemented")
}
func (UnimplementedActivityServer) mustEmbedUnimplementedActivityServer() {}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_GrabPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrabPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GrabPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_GrabPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GrabPoints(ctx, req.(*GrabPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_LookAllActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookAllActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).LookAllActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_LookAllActivities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).LookAllActivities(ctx, req.(*LookAllActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_UserSignUpActiity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignUpActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).UserSignUpActiity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_UserSignUpActiity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).UserSignUpActiity(ctx, req.(*UserSignUpActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_UserViewAllActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserViewAllActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).UserViewAllActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_UserViewAllActivities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).UserViewAllActivities(ctx, req.(*UserViewAllActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activity.activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GrabPoints",
			Handler:    _Activity_GrabPoints_Handler,
		},
		{
			MethodName: "LookAllActivities",
			Handler:    _Activity_LookAllActivities_Handler,
		},
		{
			MethodName: "UserSignUpActiity",
			Handler:    _Activity_UserSignUpActiity_Handler,
		},
		{
			MethodName: "UserViewAllActivities",
			Handler:    _Activity_UserViewAllActivities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/activity.proto",
}
