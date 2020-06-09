// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control.proto

package ctl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MgmtCtlClient is the client API for MgmtCtl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MgmtCtlClient interface {
	// Prepare nonvolatile storage devices for use with DAOS
	StoragePrepare(ctx context.Context, in *StoragePrepareReq, opts ...grpc.CallOption) (*StoragePrepareResp, error)
	// Retrieve details of nonvolatile storage on server, including health info
	StorageScan(ctx context.Context, in *StorageScanReq, opts ...grpc.CallOption) (*StorageScanResp, error)
	// Format nonvolatile storage devices for use with DAOS
	StorageFormat(ctx context.Context, in *StorageFormatReq, opts ...grpc.CallOption) (*StorageFormatResp, error)
	// Query DAOS system status
	SystemQuery(ctx context.Context, in *SystemQueryReq, opts ...grpc.CallOption) (*SystemQueryResp, error)
	// Stop DAOS system (shutdown data-plane instances)
	SystemStop(ctx context.Context, in *SystemStopReq, opts ...grpc.CallOption) (*SystemStopResp, error)
	// ResetFormat DAOS system (restart data-plane instances)
	SystemResetFormat(ctx context.Context, in *SystemResetFormatReq, opts ...grpc.CallOption) (*SystemResetFormatResp, error)
	// Start DAOS system (restart data-plane instances)
	SystemStart(ctx context.Context, in *SystemStartReq, opts ...grpc.CallOption) (*SystemStartResp, error)
	// Perform a fabric scan to determine the available provider, device, NUMA node combinations
	NetworkScan(ctx context.Context, in *NetworkScanReq, opts ...grpc.CallOption) (*NetworkScanResp, error)
}

type mgmtCtlClient struct {
	cc *grpc.ClientConn
}

func NewMgmtCtlClient(cc *grpc.ClientConn) MgmtCtlClient {
	return &mgmtCtlClient{cc}
}

func (c *mgmtCtlClient) StoragePrepare(ctx context.Context, in *StoragePrepareReq, opts ...grpc.CallOption) (*StoragePrepareResp, error) {
	out := new(StoragePrepareResp)
	err := c.cc.Invoke(ctx, "/ctl.MgmtCtl/StoragePrepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtCtlClient) StorageScan(ctx context.Context, in *StorageScanReq, opts ...grpc.CallOption) (*StorageScanResp, error) {
	out := new(StorageScanResp)
	err := c.cc.Invoke(ctx, "/ctl.MgmtCtl/StorageScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtCtlClient) StorageFormat(ctx context.Context, in *StorageFormatReq, opts ...grpc.CallOption) (*StorageFormatResp, error) {
	out := new(StorageFormatResp)
	err := c.cc.Invoke(ctx, "/ctl.MgmtCtl/StorageFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtCtlClient) SystemQuery(ctx context.Context, in *SystemQueryReq, opts ...grpc.CallOption) (*SystemQueryResp, error) {
	out := new(SystemQueryResp)
	err := c.cc.Invoke(ctx, "/ctl.MgmtCtl/SystemQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtCtlClient) SystemStop(ctx context.Context, in *SystemStopReq, opts ...grpc.CallOption) (*SystemStopResp, error) {
	out := new(SystemStopResp)
	err := c.cc.Invoke(ctx, "/ctl.MgmtCtl/SystemStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtCtlClient) SystemResetFormat(ctx context.Context, in *SystemResetFormatReq, opts ...grpc.CallOption) (*SystemResetFormatResp, error) {
	out := new(SystemResetFormatResp)
	err := c.cc.Invoke(ctx, "/ctl.MgmtCtl/SystemResetFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtCtlClient) SystemStart(ctx context.Context, in *SystemStartReq, opts ...grpc.CallOption) (*SystemStartResp, error) {
	out := new(SystemStartResp)
	err := c.cc.Invoke(ctx, "/ctl.MgmtCtl/SystemStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtCtlClient) NetworkScan(ctx context.Context, in *NetworkScanReq, opts ...grpc.CallOption) (*NetworkScanResp, error) {
	out := new(NetworkScanResp)
	err := c.cc.Invoke(ctx, "/ctl.MgmtCtl/NetworkScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MgmtCtlServer is the server API for MgmtCtl service.
type MgmtCtlServer interface {
	// Prepare nonvolatile storage devices for use with DAOS
	StoragePrepare(context.Context, *StoragePrepareReq) (*StoragePrepareResp, error)
	// Retrieve details of nonvolatile storage on server, including health info
	StorageScan(context.Context, *StorageScanReq) (*StorageScanResp, error)
	// Format nonvolatile storage devices for use with DAOS
	StorageFormat(context.Context, *StorageFormatReq) (*StorageFormatResp, error)
	// Query DAOS system status
	SystemQuery(context.Context, *SystemQueryReq) (*SystemQueryResp, error)
	// Stop DAOS system (shutdown data-plane instances)
	SystemStop(context.Context, *SystemStopReq) (*SystemStopResp, error)
	// ResetFormat DAOS system (restart data-plane instances)
	SystemResetFormat(context.Context, *SystemResetFormatReq) (*SystemResetFormatResp, error)
	// Start DAOS system (restart data-plane instances)
	SystemStart(context.Context, *SystemStartReq) (*SystemStartResp, error)
	// Perform a fabric scan to determine the available provider, device, NUMA node combinations
	NetworkScan(context.Context, *NetworkScanReq) (*NetworkScanResp, error)
}

func RegisterMgmtCtlServer(s *grpc.Server, srv MgmtCtlServer) {
	s.RegisterService(&_MgmtCtl_serviceDesc, srv)
}

func _MgmtCtl_StoragePrepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoragePrepareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtCtlServer).StoragePrepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.MgmtCtl/StoragePrepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtCtlServer).StoragePrepare(ctx, req.(*StoragePrepareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtCtl_StorageScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageScanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtCtlServer).StorageScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.MgmtCtl/StorageScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtCtlServer).StorageScan(ctx, req.(*StorageScanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtCtl_StorageFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageFormatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtCtlServer).StorageFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.MgmtCtl/StorageFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtCtlServer).StorageFormat(ctx, req.(*StorageFormatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtCtl_SystemQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtCtlServer).SystemQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.MgmtCtl/SystemQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtCtlServer).SystemQuery(ctx, req.(*SystemQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtCtl_SystemStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemStopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtCtlServer).SystemStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.MgmtCtl/SystemStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtCtlServer).SystemStop(ctx, req.(*SystemStopReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtCtl_SystemResetFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemResetFormatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtCtlServer).SystemResetFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.MgmtCtl/SystemResetFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtCtlServer).SystemResetFormat(ctx, req.(*SystemResetFormatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtCtl_SystemStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemStartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtCtlServer).SystemStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.MgmtCtl/SystemStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtCtlServer).SystemStart(ctx, req.(*SystemStartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtCtl_NetworkScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkScanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtCtlServer).NetworkScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.MgmtCtl/NetworkScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtCtlServer).NetworkScan(ctx, req.(*NetworkScanReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MgmtCtl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ctl.MgmtCtl",
	HandlerType: (*MgmtCtlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoragePrepare",
			Handler:    _MgmtCtl_StoragePrepare_Handler,
		},
		{
			MethodName: "StorageScan",
			Handler:    _MgmtCtl_StorageScan_Handler,
		},
		{
			MethodName: "StorageFormat",
			Handler:    _MgmtCtl_StorageFormat_Handler,
		},
		{
			MethodName: "SystemQuery",
			Handler:    _MgmtCtl_SystemQuery_Handler,
		},
		{
			MethodName: "SystemStop",
			Handler:    _MgmtCtl_SystemStop_Handler,
		},
		{
			MethodName: "SystemResetFormat",
			Handler:    _MgmtCtl_SystemResetFormat_Handler,
		},
		{
			MethodName: "SystemStart",
			Handler:    _MgmtCtl_SystemStart_Handler,
		},
		{
			MethodName: "NetworkScan",
			Handler:    _MgmtCtl_NetworkScan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}

func init() { proto.RegisterFile("control.proto", fileDescriptor_control_5589eb17ef6ac216) }

var fileDescriptor_control_5589eb17ef6ac216 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd2, 0x4f, 0x4a, 0xc3, 0x50,
	0x10, 0x06, 0x70, 0x41, 0x51, 0x18, 0x8d, 0xe0, 0x54, 0x2b, 0x66, 0xe9, 0x01, 0xb2, 0xd0, 0x85,
	0xe0, 0x4a, 0x28, 0xb8, 0x52, 0xd1, 0xe6, 0x04, 0xcf, 0x30, 0x74, 0x61, 0x92, 0x89, 0xf3, 0x46,
	0xa4, 0x27, 0xf0, 0xda, 0xf2, 0xfe, 0xd4, 0x4e, 0x6c, 0xbb, 0x7c, 0xbf, 0xcc, 0x97, 0x2f, 0x93,
	0x04, 0x8a, 0x86, 0x7b, 0x15, 0x6e, 0xab, 0x41, 0x58, 0x19, 0xf7, 0x1b, 0x6d, 0xcb, 0xc2, 0x2b,
	0x8b, 0x5b, 0x50, 0xb2, 0xb2, 0xe8, 0x49, 0xbf, 0x59, 0x3e, 0xf2, 0xf1, 0xc4, 0x2f, 0xbd, 0x52,
	0x97, 0x4e, 0x37, 0x3f, 0x07, 0x70, 0xf4, 0xbc, 0xe8, 0x74, 0xa6, 0x2d, 0xce, 0xe0, 0xb4, 0x4e,
	0xc9, 0x57, 0xa1, 0xc1, 0x09, 0xe1, 0xb4, 0x6a, 0xb4, 0xad, 0xc6, 0x38, 0xa7, 0xcf, 0xf2, 0x72,
	0xab, 0xfb, 0xe1, 0x7a, 0x0f, 0xef, 0xe1, 0x38, 0x7b, 0xdd, 0xb8, 0x1e, 0x27, 0x76, 0x32, 0x48,
	0x88, 0x9f, 0x6f, 0x62, 0xcc, 0x3e, 0x40, 0x91, 0xf1, 0x91, 0xa5, 0x73, 0x8a, 0x17, 0x76, 0x30,
	0x59, 0xc8, 0x4f, 0xb7, 0xf1, 0x5f, 0x7b, 0x5c, 0xef, 0xed, 0x8b, 0x64, 0xb9, 0x6a, 0x5f, 0x8b,
	0x69, 0xb7, 0x18, 0xb3, 0x77, 0x00, 0x09, 0x6b, 0xe5, 0x01, 0xd1, 0x4c, 0x05, 0x08, 0xc9, 0xc9,
	0x86, 0xc5, 0xe0, 0x13, 0x9c, 0x25, 0x9b, 0x93, 0x27, 0xcd, 0x8f, 0x7e, 0x65, 0x66, 0x8d, 0x87,
	0xdb, 0x94, 0xbb, 0x2e, 0x8d, 0x57, 0xa8, 0xd5, 0x89, 0xe2, 0xb8, 0xd3, 0x89, 0xfe, 0x5f, 0x21,
	0xe3, 0x2a, 0xfb, 0x92, 0x3e, 0xb6, 0x79, 0xf9, 0x46, 0xd6, 0xd9, 0x11, 0x86, 0xec, 0xfb, 0x61,
	0xfc, 0x21, 0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xfc, 0x24, 0xea, 0x52, 0x02, 0x00,
	0x00,
}
