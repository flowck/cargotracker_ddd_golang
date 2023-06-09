// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: cargotracker.proto

package static

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

// CargoTrackerClient is the client API for CargoTracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CargoTrackerClient interface {
	BookNewCargo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type cargoTrackerClient struct {
	cc grpc.ClientConnInterface
}

func NewCargoTrackerClient(cc grpc.ClientConnInterface) CargoTrackerClient {
	return &cargoTrackerClient{cc}
}

func (c *cargoTrackerClient) BookNewCargo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cargotracker.CargoTracker/BookNewCargo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CargoTrackerServer is the server API for CargoTracker service.
// All implementations must embed UnimplementedCargoTrackerServer
// for forward compatibility
type CargoTrackerServer interface {
	BookNewCargo(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedCargoTrackerServer()
}

// UnimplementedCargoTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedCargoTrackerServer struct {
}

func (UnimplementedCargoTrackerServer) BookNewCargo(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookNewCargo not implemented")
}
func (UnimplementedCargoTrackerServer) mustEmbedUnimplementedCargoTrackerServer() {}

// UnsafeCargoTrackerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CargoTrackerServer will
// result in compilation errors.
type UnsafeCargoTrackerServer interface {
	mustEmbedUnimplementedCargoTrackerServer()
}

func RegisterCargoTrackerServer(s grpc.ServiceRegistrar, srv CargoTrackerServer) {
	s.RegisterService(&CargoTracker_ServiceDesc, srv)
}

func _CargoTracker_BookNewCargo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CargoTrackerServer).BookNewCargo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cargotracker.CargoTracker/BookNewCargo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CargoTrackerServer).BookNewCargo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CargoTracker_ServiceDesc is the grpc.ServiceDesc for CargoTracker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CargoTracker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cargotracker.CargoTracker",
	HandlerType: (*CargoTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookNewCargo",
			Handler:    _CargoTracker_BookNewCargo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cargotracker.proto",
}
