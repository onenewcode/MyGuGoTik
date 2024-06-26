// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: relation.proto

package relation

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
	RelationService_Follow_FullMethodName            = "/rpc.Relation.RelationService/Follow"
	RelationService_Unfollow_FullMethodName          = "/rpc.Relation.RelationService/Unfollow"
	RelationService_GetFollowList_FullMethodName     = "/rpc.Relation.RelationService/GetFollowList"
	RelationService_CountFollowList_FullMethodName   = "/rpc.Relation.RelationService/CountFollowList"
	RelationService_GetFollowerList_FullMethodName   = "/rpc.Relation.RelationService/GetFollowerList"
	RelationService_CountFollowerList_FullMethodName = "/rpc.Relation.RelationService/CountFollowerList"
	RelationService_GetFriendList_FullMethodName     = "/rpc.Relation.RelationService/GetFriendList"
	RelationService_IsFollow_FullMethodName          = "/rpc.Relation.RelationService/IsFollow"
)

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	Follow(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error)
	Unfollow(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error)
	GetFollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error)
	CountFollowList(ctx context.Context, in *CountFollowListRequest, opts ...grpc.CallOption) (*CountFollowListResponse, error)
	GetFollowerList(ctx context.Context, in *FollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
	CountFollowerList(ctx context.Context, in *CountFollowerListRequest, opts ...grpc.CallOption) (*CountFollowerListResponse, error)
	GetFriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) Follow(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error) {
	out := new(RelationActionResponse)
	err := c.cc.Invoke(ctx, RelationService_Follow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) Unfollow(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error) {
	out := new(RelationActionResponse)
	err := c.cc.Invoke(ctx, RelationService_Unfollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error) {
	out := new(FollowListResponse)
	err := c.cc.Invoke(ctx, RelationService_GetFollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) CountFollowList(ctx context.Context, in *CountFollowListRequest, opts ...grpc.CallOption) (*CountFollowListResponse, error) {
	out := new(CountFollowListResponse)
	err := c.cc.Invoke(ctx, RelationService_CountFollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFollowerList(ctx context.Context, in *FollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	out := new(FollowerListResponse)
	err := c.cc.Invoke(ctx, RelationService_GetFollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) CountFollowerList(ctx context.Context, in *CountFollowerListRequest, opts ...grpc.CallOption) (*CountFollowerListResponse, error) {
	out := new(CountFollowerListResponse)
	err := c.cc.Invoke(ctx, RelationService_CountFollowerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	out := new(FriendListResponse)
	err := c.cc.Invoke(ctx, RelationService_GetFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error) {
	out := new(IsFollowResponse)
	err := c.cc.Invoke(ctx, RelationService_IsFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility
type RelationServiceServer interface {
	Follow(context.Context, *RelationActionRequest) (*RelationActionResponse, error)
	Unfollow(context.Context, *RelationActionRequest) (*RelationActionResponse, error)
	GetFollowList(context.Context, *FollowListRequest) (*FollowListResponse, error)
	CountFollowList(context.Context, *CountFollowListRequest) (*CountFollowListResponse, error)
	GetFollowerList(context.Context, *FollowerListRequest) (*FollowerListResponse, error)
	CountFollowerList(context.Context, *CountFollowerListRequest) (*CountFollowerListResponse, error)
	GetFriendList(context.Context, *FriendListRequest) (*FriendListResponse, error)
	IsFollow(context.Context, *IsFollowRequest) (*IsFollowResponse, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (UnimplementedRelationServiceServer) Follow(context.Context, *RelationActionRequest) (*RelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedRelationServiceServer) Unfollow(context.Context, *RelationActionRequest) (*RelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedRelationServiceServer) GetFollowList(context.Context, *FollowListRequest) (*FollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowList not implemented")
}
func (UnimplementedRelationServiceServer) CountFollowList(context.Context, *CountFollowListRequest) (*CountFollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFollowList not implemented")
}
func (UnimplementedRelationServiceServer) GetFollowerList(context.Context, *FollowerListRequest) (*FollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerList not implemented")
}
func (UnimplementedRelationServiceServer) CountFollowerList(context.Context, *CountFollowerListRequest) (*CountFollowerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountFollowerList not implemented")
}
func (UnimplementedRelationServiceServer) GetFriendList(context.Context, *FriendListRequest) (*FriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedRelationServiceServer) IsFollow(context.Context, *IsFollowRequest) (*IsFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFollow not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_Follow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).Follow(ctx, req.(*RelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_Unfollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).Unfollow(ctx, req.(*RelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetFollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFollowList(ctx, req.(*FollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_CountFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).CountFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_CountFollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).CountFollowList(ctx, req.(*CountFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetFollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFollowerList(ctx, req.(*FollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_CountFollowerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).CountFollowerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_CountFollowerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).CountFollowerList(ctx, req.(*CountFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFriendList(ctx, req.(*FriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_IsFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).IsFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_IsFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).IsFollow(ctx, req.(*IsFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Relation.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Follow",
			Handler:    _RelationService_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _RelationService_Unfollow_Handler,
		},
		{
			MethodName: "GetFollowList",
			Handler:    _RelationService_GetFollowList_Handler,
		},
		{
			MethodName: "CountFollowList",
			Handler:    _RelationService_CountFollowList_Handler,
		},
		{
			MethodName: "GetFollowerList",
			Handler:    _RelationService_GetFollowerList_Handler,
		},
		{
			MethodName: "CountFollowerList",
			Handler:    _RelationService_CountFollowerList_Handler,
		},
		{
			MethodName: "GetFriendList",
			Handler:    _RelationService_GetFriendList_Handler,
		},
		{
			MethodName: "IsFollow",
			Handler:    _RelationService_IsFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation.proto",
}
