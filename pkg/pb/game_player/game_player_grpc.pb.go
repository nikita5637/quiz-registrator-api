// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: game_player/game_player.proto

package gameplayer

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// CreateGamePlayer creates new game player
	CreateGamePlayer(ctx context.Context, in *CreateGamePlayerRequest, opts ...grpc.CallOption) (*GamePlayer, error)
	// DeleteGamePlayer deletes game player
	DeleteGamePlayer(ctx context.Context, in *DeleteGamePlayerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetGamePlayer returns game player by ID
	GetGamePlayer(ctx context.Context, in *GetGamePlayerRequest, opts ...grpc.CallOption) (*GamePlayer, error)
	// GetGamePlayersByGameID returns list of game players by game ID
	GetGamePlayersByGameID(ctx context.Context, in *GetGamePlayersByGameIDRequest, opts ...grpc.CallOption) (*GetGamePlayersByGameIDResponse, error)
	// GetUserGameIDs returns complete list of user game IDs
	GetUserGameIDs(ctx context.Context, in *GetUserGameIDsRequest, opts ...grpc.CallOption) (*GetUserGameIDsResponse, error)
	// PatchGamePlayer patches game player
	PatchGamePlayer(ctx context.Context, in *PatchGamePlayerRequest, opts ...grpc.CallOption) (*GamePlayer, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateGamePlayer(ctx context.Context, in *CreateGamePlayerRequest, opts ...grpc.CallOption) (*GamePlayer, error) {
	out := new(GamePlayer)
	err := c.cc.Invoke(ctx, "/game_player.Service/CreateGamePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteGamePlayer(ctx context.Context, in *DeleteGamePlayerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/game_player.Service/DeleteGamePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetGamePlayer(ctx context.Context, in *GetGamePlayerRequest, opts ...grpc.CallOption) (*GamePlayer, error) {
	out := new(GamePlayer)
	err := c.cc.Invoke(ctx, "/game_player.Service/GetGamePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetGamePlayersByGameID(ctx context.Context, in *GetGamePlayersByGameIDRequest, opts ...grpc.CallOption) (*GetGamePlayersByGameIDResponse, error) {
	out := new(GetGamePlayersByGameIDResponse)
	err := c.cc.Invoke(ctx, "/game_player.Service/GetGamePlayersByGameID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetUserGameIDs(ctx context.Context, in *GetUserGameIDsRequest, opts ...grpc.CallOption) (*GetUserGameIDsResponse, error) {
	out := new(GetUserGameIDsResponse)
	err := c.cc.Invoke(ctx, "/game_player.Service/GetUserGameIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PatchGamePlayer(ctx context.Context, in *PatchGamePlayerRequest, opts ...grpc.CallOption) (*GamePlayer, error) {
	out := new(GamePlayer)
	err := c.cc.Invoke(ctx, "/game_player.Service/PatchGamePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// CreateGamePlayer creates new game player
	CreateGamePlayer(context.Context, *CreateGamePlayerRequest) (*GamePlayer, error)
	// DeleteGamePlayer deletes game player
	DeleteGamePlayer(context.Context, *DeleteGamePlayerRequest) (*emptypb.Empty, error)
	// GetGamePlayer returns game player by ID
	GetGamePlayer(context.Context, *GetGamePlayerRequest) (*GamePlayer, error)
	// GetGamePlayersByGameID returns list of game players by game ID
	GetGamePlayersByGameID(context.Context, *GetGamePlayersByGameIDRequest) (*GetGamePlayersByGameIDResponse, error)
	// GetUserGameIDs returns complete list of user game IDs
	GetUserGameIDs(context.Context, *GetUserGameIDsRequest) (*GetUserGameIDsResponse, error)
	// PatchGamePlayer patches game player
	PatchGamePlayer(context.Context, *PatchGamePlayerRequest) (*GamePlayer, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateGamePlayer(context.Context, *CreateGamePlayerRequest) (*GamePlayer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGamePlayer not implemented")
}
func (UnimplementedServiceServer) DeleteGamePlayer(context.Context, *DeleteGamePlayerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGamePlayer not implemented")
}
func (UnimplementedServiceServer) GetGamePlayer(context.Context, *GetGamePlayerRequest) (*GamePlayer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGamePlayer not implemented")
}
func (UnimplementedServiceServer) GetGamePlayersByGameID(context.Context, *GetGamePlayersByGameIDRequest) (*GetGamePlayersByGameIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGamePlayersByGameID not implemented")
}
func (UnimplementedServiceServer) GetUserGameIDs(context.Context, *GetUserGameIDsRequest) (*GetUserGameIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGameIDs not implemented")
}
func (UnimplementedServiceServer) PatchGamePlayer(context.Context, *PatchGamePlayerRequest) (*GamePlayer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchGamePlayer not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateGamePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGamePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateGamePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.Service/CreateGamePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateGamePlayer(ctx, req.(*CreateGamePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteGamePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGamePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteGamePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.Service/DeleteGamePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteGamePlayer(ctx, req.(*DeleteGamePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetGamePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGamePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetGamePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.Service/GetGamePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetGamePlayer(ctx, req.(*GetGamePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetGamePlayersByGameID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGamePlayersByGameIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetGamePlayersByGameID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.Service/GetGamePlayersByGameID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetGamePlayersByGameID(ctx, req.(*GetGamePlayersByGameIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetUserGameIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGameIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetUserGameIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.Service/GetUserGameIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetUserGameIDs(ctx, req.(*GetUserGameIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PatchGamePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchGamePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PatchGamePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.Service/PatchGamePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PatchGamePlayer(ctx, req.(*PatchGamePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game_player.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGamePlayer",
			Handler:    _Service_CreateGamePlayer_Handler,
		},
		{
			MethodName: "DeleteGamePlayer",
			Handler:    _Service_DeleteGamePlayer_Handler,
		},
		{
			MethodName: "GetGamePlayer",
			Handler:    _Service_GetGamePlayer_Handler,
		},
		{
			MethodName: "GetGamePlayersByGameID",
			Handler:    _Service_GetGamePlayersByGameID_Handler,
		},
		{
			MethodName: "GetUserGameIDs",
			Handler:    _Service_GetUserGameIDs_Handler,
		},
		{
			MethodName: "PatchGamePlayer",
			Handler:    _Service_PatchGamePlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_player/game_player.proto",
}

// RegistratorServiceClient is the client API for RegistratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistratorServiceClient interface {
	// RegisterPlayer registers player for a game
	RegisterPlayer(ctx context.Context, in *RegisterPlayerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UnregisterPlayer unregisters player from a game
	UnregisterPlayer(ctx context.Context, in *UnregisterPlayerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UpdatePlayerDegree updates player degree
	UpdatePlayerDegree(ctx context.Context, in *UpdatePlayerDegreeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type registratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistratorServiceClient(cc grpc.ClientConnInterface) RegistratorServiceClient {
	return &registratorServiceClient{cc}
}

func (c *registratorServiceClient) RegisterPlayer(ctx context.Context, in *RegisterPlayerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/game_player.RegistratorService/RegisterPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registratorServiceClient) UnregisterPlayer(ctx context.Context, in *UnregisterPlayerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/game_player.RegistratorService/UnregisterPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registratorServiceClient) UpdatePlayerDegree(ctx context.Context, in *UpdatePlayerDegreeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/game_player.RegistratorService/UpdatePlayerDegree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistratorServiceServer is the server API for RegistratorService service.
// All implementations must embed UnimplementedRegistratorServiceServer
// for forward compatibility
type RegistratorServiceServer interface {
	// RegisterPlayer registers player for a game
	RegisterPlayer(context.Context, *RegisterPlayerRequest) (*emptypb.Empty, error)
	// UnregisterPlayer unregisters player from a game
	UnregisterPlayer(context.Context, *UnregisterPlayerRequest) (*emptypb.Empty, error)
	// UpdatePlayerDegree updates player degree
	UpdatePlayerDegree(context.Context, *UpdatePlayerDegreeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRegistratorServiceServer()
}

// UnimplementedRegistratorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegistratorServiceServer struct {
}

func (UnimplementedRegistratorServiceServer) RegisterPlayer(context.Context, *RegisterPlayerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPlayer not implemented")
}
func (UnimplementedRegistratorServiceServer) UnregisterPlayer(context.Context, *UnregisterPlayerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterPlayer not implemented")
}
func (UnimplementedRegistratorServiceServer) UpdatePlayerDegree(context.Context, *UpdatePlayerDegreeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlayerDegree not implemented")
}
func (UnimplementedRegistratorServiceServer) mustEmbedUnimplementedRegistratorServiceServer() {}

// UnsafeRegistratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistratorServiceServer will
// result in compilation errors.
type UnsafeRegistratorServiceServer interface {
	mustEmbedUnimplementedRegistratorServiceServer()
}

func RegisterRegistratorServiceServer(s grpc.ServiceRegistrar, srv RegistratorServiceServer) {
	s.RegisterService(&RegistratorService_ServiceDesc, srv)
}

func _RegistratorService_RegisterPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistratorServiceServer).RegisterPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.RegistratorService/RegisterPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistratorServiceServer).RegisterPlayer(ctx, req.(*RegisterPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistratorService_UnregisterPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistratorServiceServer).UnregisterPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.RegistratorService/UnregisterPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistratorServiceServer).UnregisterPlayer(ctx, req.(*UnregisterPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistratorService_UpdatePlayerDegree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlayerDegreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistratorServiceServer).UpdatePlayerDegree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game_player.RegistratorService/UpdatePlayerDegree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistratorServiceServer).UpdatePlayerDegree(ctx, req.(*UpdatePlayerDegreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegistratorService_ServiceDesc is the grpc.ServiceDesc for RegistratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegistratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game_player.RegistratorService",
	HandlerType: (*RegistratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPlayer",
			Handler:    _RegistratorService_RegisterPlayer_Handler,
		},
		{
			MethodName: "UnregisterPlayer",
			Handler:    _RegistratorService_UnregisterPlayer_Handler,
		},
		{
			MethodName: "UpdatePlayerDegree",
			Handler:    _RegistratorService_UpdatePlayerDegree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_player/game_player.proto",
}
