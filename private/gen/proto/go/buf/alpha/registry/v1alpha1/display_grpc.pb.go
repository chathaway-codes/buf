// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             (unknown)
// source: buf/alpha/registry/v1alpha1/display.proto

package registryv1alpha1

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

// DisplayServiceClient is the client API for DisplayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DisplayServiceClient interface {
	// DisplayRepositoryRolesToManage returns which roles should be displayed
	// to the user when they are managing contributors on the repository.
	DisplayRepositoryRolesToManage(ctx context.Context, in *DisplayRepositoryRolesToManageRequest, opts ...grpc.CallOption) (*DisplayRepositoryRolesToManageResponse, error)
	// DisplayRepositoryRolesToManageByUser returns which roles should be displayed
	// to the user when they are managing a specific contributor on the repository.
	DisplayRepositoryRolesToManageByUser(ctx context.Context, in *DisplayRepositoryRolesToManageByUserRequest, opts ...grpc.CallOption) (*DisplayRepositoryRolesToManageByUserResponse, error)
	// DisplayPluginRolesToManage returns which roles should be displayed
	// to the user when they are managing contributors on the plugin.
	DisplayPluginRolesToManage(ctx context.Context, in *DisplayPluginRolesToManageRequest, opts ...grpc.CallOption) (*DisplayPluginRolesToManageResponse, error)
	// DisplayPluginRolesToManageByUser returns which roles should be displayed
	// to the user when they are managing a specific contributor on the plugin.
	DisplayPluginRolesToManageByUser(ctx context.Context, in *DisplayPluginRolesToManageByUserRequest, opts ...grpc.CallOption) (*DisplayPluginRolesToManageByUserResponse, error)
	// DisplayTemplateRolesToManage returns which roles should be displayed
	// to the user when they are managing contributors on the template.
	DisplayTemplateRolesToManage(ctx context.Context, in *DisplayTemplateRolesToManageRequest, opts ...grpc.CallOption) (*DisplayTemplateRolesToManageResponse, error)
	// DisplayTemplateRolesToManageByUser returns which roles should be displayed
	// to the user when they are managing a specific contributor on the template.
	DisplayTemplateRolesToManageByUser(ctx context.Context, in *DisplayTemplateRolesToManageByUserRequest, opts ...grpc.CallOption) (*DisplayTemplateRolesToManageByUserResponse, error)
}

type displayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDisplayServiceClient(cc grpc.ClientConnInterface) DisplayServiceClient {
	return &displayServiceClient{cc}
}

func (c *displayServiceClient) DisplayRepositoryRolesToManage(ctx context.Context, in *DisplayRepositoryRolesToManageRequest, opts ...grpc.CallOption) (*DisplayRepositoryRolesToManageResponse, error) {
	out := new(DisplayRepositoryRolesToManageResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.DisplayService/DisplayRepositoryRolesToManage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayRepositoryRolesToManageByUser(ctx context.Context, in *DisplayRepositoryRolesToManageByUserRequest, opts ...grpc.CallOption) (*DisplayRepositoryRolesToManageByUserResponse, error) {
	out := new(DisplayRepositoryRolesToManageByUserResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.DisplayService/DisplayRepositoryRolesToManageByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayPluginRolesToManage(ctx context.Context, in *DisplayPluginRolesToManageRequest, opts ...grpc.CallOption) (*DisplayPluginRolesToManageResponse, error) {
	out := new(DisplayPluginRolesToManageResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.DisplayService/DisplayPluginRolesToManage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayPluginRolesToManageByUser(ctx context.Context, in *DisplayPluginRolesToManageByUserRequest, opts ...grpc.CallOption) (*DisplayPluginRolesToManageByUserResponse, error) {
	out := new(DisplayPluginRolesToManageByUserResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.DisplayService/DisplayPluginRolesToManageByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayTemplateRolesToManage(ctx context.Context, in *DisplayTemplateRolesToManageRequest, opts ...grpc.CallOption) (*DisplayTemplateRolesToManageResponse, error) {
	out := new(DisplayTemplateRolesToManageResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.DisplayService/DisplayTemplateRolesToManage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayTemplateRolesToManageByUser(ctx context.Context, in *DisplayTemplateRolesToManageByUserRequest, opts ...grpc.CallOption) (*DisplayTemplateRolesToManageByUserResponse, error) {
	out := new(DisplayTemplateRolesToManageByUserResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.DisplayService/DisplayTemplateRolesToManageByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisplayServiceServer is the server API for DisplayService service.
// All implementations should embed UnimplementedDisplayServiceServer
// for forward compatibility
type DisplayServiceServer interface {
	// DisplayRepositoryRolesToManage returns which roles should be displayed
	// to the user when they are managing contributors on the repository.
	DisplayRepositoryRolesToManage(context.Context, *DisplayRepositoryRolesToManageRequest) (*DisplayRepositoryRolesToManageResponse, error)
	// DisplayRepositoryRolesToManageByUser returns which roles should be displayed
	// to the user when they are managing a specific contributor on the repository.
	DisplayRepositoryRolesToManageByUser(context.Context, *DisplayRepositoryRolesToManageByUserRequest) (*DisplayRepositoryRolesToManageByUserResponse, error)
	// DisplayPluginRolesToManage returns which roles should be displayed
	// to the user when they are managing contributors on the plugin.
	DisplayPluginRolesToManage(context.Context, *DisplayPluginRolesToManageRequest) (*DisplayPluginRolesToManageResponse, error)
	// DisplayPluginRolesToManageByUser returns which roles should be displayed
	// to the user when they are managing a specific contributor on the plugin.
	DisplayPluginRolesToManageByUser(context.Context, *DisplayPluginRolesToManageByUserRequest) (*DisplayPluginRolesToManageByUserResponse, error)
	// DisplayTemplateRolesToManage returns which roles should be displayed
	// to the user when they are managing contributors on the template.
	DisplayTemplateRolesToManage(context.Context, *DisplayTemplateRolesToManageRequest) (*DisplayTemplateRolesToManageResponse, error)
	// DisplayTemplateRolesToManageByUser returns which roles should be displayed
	// to the user when they are managing a specific contributor on the template.
	DisplayTemplateRolesToManageByUser(context.Context, *DisplayTemplateRolesToManageByUserRequest) (*DisplayTemplateRolesToManageByUserResponse, error)
}

// UnimplementedDisplayServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDisplayServiceServer struct {
}

func (UnimplementedDisplayServiceServer) DisplayRepositoryRolesToManage(context.Context, *DisplayRepositoryRolesToManageRequest) (*DisplayRepositoryRolesToManageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayRepositoryRolesToManage not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayRepositoryRolesToManageByUser(context.Context, *DisplayRepositoryRolesToManageByUserRequest) (*DisplayRepositoryRolesToManageByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayRepositoryRolesToManageByUser not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayPluginRolesToManage(context.Context, *DisplayPluginRolesToManageRequest) (*DisplayPluginRolesToManageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayPluginRolesToManage not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayPluginRolesToManageByUser(context.Context, *DisplayPluginRolesToManageByUserRequest) (*DisplayPluginRolesToManageByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayPluginRolesToManageByUser not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayTemplateRolesToManage(context.Context, *DisplayTemplateRolesToManageRequest) (*DisplayTemplateRolesToManageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayTemplateRolesToManage not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayTemplateRolesToManageByUser(context.Context, *DisplayTemplateRolesToManageByUserRequest) (*DisplayTemplateRolesToManageByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayTemplateRolesToManageByUser not implemented")
}

// UnsafeDisplayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DisplayServiceServer will
// result in compilation errors.
type UnsafeDisplayServiceServer interface {
	mustEmbedUnimplementedDisplayServiceServer()
}

func RegisterDisplayServiceServer(s grpc.ServiceRegistrar, srv DisplayServiceServer) {
	s.RegisterService(&DisplayService_ServiceDesc, srv)
}

func _DisplayService_DisplayRepositoryRolesToManage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayRepositoryRolesToManageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayRepositoryRolesToManage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.DisplayService/DisplayRepositoryRolesToManage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayRepositoryRolesToManage(ctx, req.(*DisplayRepositoryRolesToManageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayRepositoryRolesToManageByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayRepositoryRolesToManageByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayRepositoryRolesToManageByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.DisplayService/DisplayRepositoryRolesToManageByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayRepositoryRolesToManageByUser(ctx, req.(*DisplayRepositoryRolesToManageByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayPluginRolesToManage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayPluginRolesToManageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayPluginRolesToManage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.DisplayService/DisplayPluginRolesToManage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayPluginRolesToManage(ctx, req.(*DisplayPluginRolesToManageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayPluginRolesToManageByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayPluginRolesToManageByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayPluginRolesToManageByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.DisplayService/DisplayPluginRolesToManageByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayPluginRolesToManageByUser(ctx, req.(*DisplayPluginRolesToManageByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayTemplateRolesToManage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayTemplateRolesToManageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayTemplateRolesToManage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.DisplayService/DisplayTemplateRolesToManage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayTemplateRolesToManage(ctx, req.(*DisplayTemplateRolesToManageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayTemplateRolesToManageByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayTemplateRolesToManageByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayTemplateRolesToManageByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.DisplayService/DisplayTemplateRolesToManageByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayTemplateRolesToManageByUser(ctx, req.(*DisplayTemplateRolesToManageByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DisplayService_ServiceDesc is the grpc.ServiceDesc for DisplayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DisplayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.DisplayService",
	HandlerType: (*DisplayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DisplayRepositoryRolesToManage",
			Handler:    _DisplayService_DisplayRepositoryRolesToManage_Handler,
		},
		{
			MethodName: "DisplayRepositoryRolesToManageByUser",
			Handler:    _DisplayService_DisplayRepositoryRolesToManageByUser_Handler,
		},
		{
			MethodName: "DisplayPluginRolesToManage",
			Handler:    _DisplayService_DisplayPluginRolesToManage_Handler,
		},
		{
			MethodName: "DisplayPluginRolesToManageByUser",
			Handler:    _DisplayService_DisplayPluginRolesToManageByUser_Handler,
		},
		{
			MethodName: "DisplayTemplateRolesToManage",
			Handler:    _DisplayService_DisplayTemplateRolesToManage_Handler,
		},
		{
			MethodName: "DisplayTemplateRolesToManageByUser",
			Handler:    _DisplayService_DisplayTemplateRolesToManageByUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/display.proto",
}
