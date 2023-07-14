// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: admin/admin.proto

package admin

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
	// CreateUserRole creates role for user
	CreateUserRole(ctx context.Context, in *CreateUserRoleRequest, opts ...grpc.CallOption) (*UserRole, error)
	// DeleteUserRole deletes user role
	DeleteUserRole(ctx context.Context, in *DeleteUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetUserRoles returns user roles by user ID
	GetUserRolesByUserID(ctx context.Context, in *GetUserRolesByUserIDRequest, opts ...grpc.CallOption) (*GetUserRolesByUserIDResponse, error)
	// ListUserRoles deletes user role
	ListUserRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUserRolesResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateUserRole(ctx context.Context, in *CreateUserRoleRequest, opts ...grpc.CallOption) (*UserRole, error) {
	out := new(UserRole)
	err := c.cc.Invoke(ctx, "/admin.Service/CreateUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteUserRole(ctx context.Context, in *DeleteUserRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/admin.Service/DeleteUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetUserRolesByUserID(ctx context.Context, in *GetUserRolesByUserIDRequest, opts ...grpc.CallOption) (*GetUserRolesByUserIDResponse, error) {
	out := new(GetUserRolesByUserIDResponse)
	err := c.cc.Invoke(ctx, "/admin.Service/GetUserRolesByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListUserRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUserRolesResponse, error) {
	out := new(ListUserRolesResponse)
	err := c.cc.Invoke(ctx, "/admin.Service/ListUserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// CreateUserRole creates role for user
	CreateUserRole(context.Context, *CreateUserRoleRequest) (*UserRole, error)
	// DeleteUserRole deletes user role
	DeleteUserRole(context.Context, *DeleteUserRoleRequest) (*emptypb.Empty, error)
	// GetUserRoles returns user roles by user ID
	GetUserRolesByUserID(context.Context, *GetUserRolesByUserIDRequest) (*GetUserRolesByUserIDResponse, error)
	// ListUserRoles deletes user role
	ListUserRoles(context.Context, *emptypb.Empty) (*ListUserRolesResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateUserRole(context.Context, *CreateUserRoleRequest) (*UserRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserRole not implemented")
}
func (UnimplementedServiceServer) DeleteUserRole(context.Context, *DeleteUserRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserRole not implemented")
}
func (UnimplementedServiceServer) GetUserRolesByUserID(context.Context, *GetUserRolesByUserIDRequest) (*GetUserRolesByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRolesByUserID not implemented")
}
func (UnimplementedServiceServer) ListUserRoles(context.Context, *emptypb.Empty) (*ListUserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRoles not implemented")
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

func _Service_CreateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Service/CreateUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateUserRole(ctx, req.(*CreateUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Service/DeleteUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteUserRole(ctx, req.(*DeleteUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetUserRolesByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRolesByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetUserRolesByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Service/GetUserRolesByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetUserRolesByUserID(ctx, req.(*GetUserRolesByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Service/ListUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListUserRoles(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserRole",
			Handler:    _Service_CreateUserRole_Handler,
		},
		{
			MethodName: "DeleteUserRole",
			Handler:    _Service_DeleteUserRole_Handler,
		},
		{
			MethodName: "GetUserRolesByUserID",
			Handler:    _Service_GetUserRolesByUserID_Handler,
		},
		{
			MethodName: "ListUserRoles",
			Handler:    _Service_ListUserRoles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/admin.proto",
}
