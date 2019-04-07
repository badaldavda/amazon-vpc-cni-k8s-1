// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	AddNetworkRequest
	AddNetworkReply
	DelNetworkRequest
	DelNetworkReply
*/
package rpc

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

type AddNetworkRequest struct {
	K8S_POD_NAME               string `protobuf:"bytes,1,opt,name=K8S_POD_NAME,json=K8SPODNAME" json:"K8S_POD_NAME,omitempty"`
	K8S_POD_NAMESPACE          string `protobuf:"bytes,2,opt,name=K8S_POD_NAMESPACE,json=K8SPODNAMESPACE" json:"K8S_POD_NAMESPACE,omitempty"`
	K8S_POD_INFRA_CONTAINER_ID string `protobuf:"bytes,3,opt,name=K8S_POD_INFRA_CONTAINER_ID,json=K8SPODINFRACONTAINERID" json:"K8S_POD_INFRA_CONTAINER_ID,omitempty"`
	Netns                      string `protobuf:"bytes,4,opt,name=Netns" json:"Netns,omitempty"`
	IfName                     string `protobuf:"bytes,5,opt,name=IfName" json:"IfName,omitempty"`
}

func (m *AddNetworkRequest) Reset()                    { *m = AddNetworkRequest{} }
func (m *AddNetworkRequest) String() string            { return proto.CompactTextString(m) }
func (*AddNetworkRequest) ProtoMessage()               {}
func (*AddNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddNetworkRequest) GetK8S_POD_NAME() string {
	if m != nil {
		return m.K8S_POD_NAME
	}
	return ""
}

func (m *AddNetworkRequest) GetK8S_POD_NAMESPACE() string {
	if m != nil {
		return m.K8S_POD_NAMESPACE
	}
	return ""
}

func (m *AddNetworkRequest) GetK8S_POD_INFRA_CONTAINER_ID() string {
	if m != nil {
		return m.K8S_POD_INFRA_CONTAINER_ID
	}
	return ""
}

func (m *AddNetworkRequest) GetNetns() string {
	if m != nil {
		return m.Netns
	}
	return ""
}

func (m *AddNetworkRequest) GetIfName() string {
	if m != nil {
		return m.IfName
	}
	return ""
}

type AddNetworkReply struct {
	Success         bool     `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	IPv4Addr        string   `protobuf:"bytes,2,opt,name=IPv4Addr" json:"IPv4Addr,omitempty"`
	IPv4Subnet      string   `protobuf:"bytes,3,opt,name=IPv4Subnet" json:"IPv4Subnet,omitempty"`
	DeviceNumber    int32    `protobuf:"varint,4,opt,name=DeviceNumber" json:"DeviceNumber,omitempty"`
	UseExternalSNAT bool     `protobuf:"varint,5,opt,name=UseExternalSNAT" json:"UseExternalSNAT,omitempty"`
	VPCcidrs        []string `protobuf:"bytes,6,rep,name=VPCcidrs" json:"VPCcidrs,omitempty"`
	CGWIP           string   `protobuf:"bytes,7,opt,name=CGWIP" json:"CGWIP,omitempty"`
	VNI             int32    `protobuf:"varint,8,opt,name=VNI" json:"VNI,omitempty"`
	DstPort         int32    `protobuf:"varint,9,opt,name=DstPort" json:"DstPort,omitempty"`
}

func (m *AddNetworkReply) Reset()                    { *m = AddNetworkReply{} }
func (m *AddNetworkReply) String() string            { return proto.CompactTextString(m) }
func (*AddNetworkReply) ProtoMessage()               {}
func (*AddNetworkReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddNetworkReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AddNetworkReply) GetIPv4Addr() string {
	if m != nil {
		return m.IPv4Addr
	}
	return ""
}

func (m *AddNetworkReply) GetIPv4Subnet() string {
	if m != nil {
		return m.IPv4Subnet
	}
	return ""
}

func (m *AddNetworkReply) GetDeviceNumber() int32 {
	if m != nil {
		return m.DeviceNumber
	}
	return 0
}

func (m *AddNetworkReply) GetUseExternalSNAT() bool {
	if m != nil {
		return m.UseExternalSNAT
	}
	return false
}

func (m *AddNetworkReply) GetVPCcidrs() []string {
	if m != nil {
		return m.VPCcidrs
	}
	return nil
}

func (m *AddNetworkReply) GetCGWIP() string {
	if m != nil {
		return m.CGWIP
	}
	return ""
}

func (m *AddNetworkReply) GetVNI() int32 {
	if m != nil {
		return m.VNI
	}
	return 0
}

func (m *AddNetworkReply) GetDstPort() int32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

type DelNetworkRequest struct {
	K8S_POD_NAME               string `protobuf:"bytes,1,opt,name=K8S_POD_NAME,json=K8SPODNAME" json:"K8S_POD_NAME,omitempty"`
	K8S_POD_NAMESPACE          string `protobuf:"bytes,2,opt,name=K8S_POD_NAMESPACE,json=K8SPODNAMESPACE" json:"K8S_POD_NAMESPACE,omitempty"`
	K8S_POD_INFRA_CONTAINER_ID string `protobuf:"bytes,3,opt,name=K8S_POD_INFRA_CONTAINER_ID,json=K8SPODINFRACONTAINERID" json:"K8S_POD_INFRA_CONTAINER_ID,omitempty"`
	IPv4Addr                   string `protobuf:"bytes,4,opt,name=IPv4Addr" json:"IPv4Addr,omitempty"`
	Reason                     string `protobuf:"bytes,5,opt,name=Reason" json:"Reason,omitempty"`
}

func (m *DelNetworkRequest) Reset()                    { *m = DelNetworkRequest{} }
func (m *DelNetworkRequest) String() string            { return proto.CompactTextString(m) }
func (*DelNetworkRequest) ProtoMessage()               {}
func (*DelNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DelNetworkRequest) GetK8S_POD_NAME() string {
	if m != nil {
		return m.K8S_POD_NAME
	}
	return ""
}

func (m *DelNetworkRequest) GetK8S_POD_NAMESPACE() string {
	if m != nil {
		return m.K8S_POD_NAMESPACE
	}
	return ""
}

func (m *DelNetworkRequest) GetK8S_POD_INFRA_CONTAINER_ID() string {
	if m != nil {
		return m.K8S_POD_INFRA_CONTAINER_ID
	}
	return ""
}

func (m *DelNetworkRequest) GetIPv4Addr() string {
	if m != nil {
		return m.IPv4Addr
	}
	return ""
}

func (m *DelNetworkRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type DelNetworkReply struct {
	Success      bool   `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	IPv4Addr     string `protobuf:"bytes,2,opt,name=IPv4Addr" json:"IPv4Addr,omitempty"`
	DeviceNumber int32  `protobuf:"varint,3,opt,name=DeviceNumber" json:"DeviceNumber,omitempty"`
}

func (m *DelNetworkReply) Reset()                    { *m = DelNetworkReply{} }
func (m *DelNetworkReply) String() string            { return proto.CompactTextString(m) }
func (*DelNetworkReply) ProtoMessage()               {}
func (*DelNetworkReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DelNetworkReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DelNetworkReply) GetIPv4Addr() string {
	if m != nil {
		return m.IPv4Addr
	}
	return ""
}

func (m *DelNetworkReply) GetDeviceNumber() int32 {
	if m != nil {
		return m.DeviceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*AddNetworkRequest)(nil), "rpc.AddNetworkRequest")
	proto.RegisterType((*AddNetworkReply)(nil), "rpc.AddNetworkReply")
	proto.RegisterType((*DelNetworkRequest)(nil), "rpc.DelNetworkRequest")
	proto.RegisterType((*DelNetworkReply)(nil), "rpc.DelNetworkReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CNIBackend service

type CNIBackendClient interface {
	AddNetwork(ctx context.Context, in *AddNetworkRequest, opts ...grpc.CallOption) (*AddNetworkReply, error)
	DelNetwork(ctx context.Context, in *DelNetworkRequest, opts ...grpc.CallOption) (*DelNetworkReply, error)
}

type cNIBackendClient struct {
	cc *grpc.ClientConn
}

func NewCNIBackendClient(cc *grpc.ClientConn) CNIBackendClient {
	return &cNIBackendClient{cc}
}

func (c *cNIBackendClient) AddNetwork(ctx context.Context, in *AddNetworkRequest, opts ...grpc.CallOption) (*AddNetworkReply, error) {
	out := new(AddNetworkReply)
	err := grpc.Invoke(ctx, "/rpc.CNIBackend/AddNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cNIBackendClient) DelNetwork(ctx context.Context, in *DelNetworkRequest, opts ...grpc.CallOption) (*DelNetworkReply, error) {
	out := new(DelNetworkReply)
	err := grpc.Invoke(ctx, "/rpc.CNIBackend/DelNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CNIBackend service

type CNIBackendServer interface {
	AddNetwork(context.Context, *AddNetworkRequest) (*AddNetworkReply, error)
	DelNetwork(context.Context, *DelNetworkRequest) (*DelNetworkReply, error)
}

func RegisterCNIBackendServer(s *grpc.Server, srv CNIBackendServer) {
	s.RegisterService(&_CNIBackend_serviceDesc, srv)
}

func _CNIBackend_AddNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNIBackendServer).AddNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CNIBackend/AddNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNIBackendServer).AddNetwork(ctx, req.(*AddNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CNIBackend_DelNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CNIBackendServer).DelNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CNIBackend/DelNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CNIBackendServer).DelNetwork(ctx, req.(*DelNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CNIBackend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.CNIBackend",
	HandlerType: (*CNIBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNetwork",
			Handler:    _CNIBackend_AddNetwork_Handler,
		},
		{
			MethodName: "DelNetwork",
			Handler:    _CNIBackend_DelNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x5d, 0x96, 0xb5, 0x6b, 0x3f, 0x4d, 0x2a, 0xb5, 0xa6, 0xca, 0xea, 0x01, 0x55, 0x39, 0x4d,
	0x1c, 0x76, 0x00, 0x0e, 0x13, 0xe2, 0x12, 0x92, 0x80, 0xac, 0x09, 0x37, 0x72, 0xc6, 0x38, 0x46,
	0x69, 0x62, 0xa4, 0xa9, 0x59, 0x12, 0x6c, 0x67, 0xb0, 0x7f, 0xc0, 0x9d, 0x9f, 0xc5, 0x89, 0x7f,
	0x84, 0xec, 0x24, 0x6b, 0x48, 0x6e, 0x9c, 0xb8, 0xf9, 0x3d, 0x7f, 0x4f, 0x79, 0xcf, 0x2f, 0x1f,
	0xcc, 0x45, 0x95, 0x5e, 0x56, 0xa2, 0x54, 0x25, 0xb2, 0x45, 0x95, 0x3a, 0xbf, 0x2c, 0x58, 0xba,
	0x59, 0x46, 0xb9, 0xfa, 0x56, 0x8a, 0x3d, 0xe3, 0x5f, 0x6b, 0x2e, 0x15, 0xda, 0xc0, 0xd9, 0xf5,
	0x55, 0x14, 0x87, 0x5b, 0x3f, 0xa6, 0xee, 0xc7, 0x00, 0x5b, 0x1b, 0xeb, 0x62, 0xce, 0xe0, 0xfa,
	0x2a, 0x0a, 0xb7, 0xbe, 0x66, 0xd0, 0x0b, 0x58, 0xf6, 0x27, 0xa2, 0xd0, 0xf5, 0x02, 0x7c, 0x6c,
	0xc6, 0x16, 0x87, 0x31, 0x43, 0xa3, 0x37, 0xb0, 0xee, 0x66, 0x09, 0x7d, 0xcf, 0xdc, 0xd8, 0xdb,
	0xd2, 0x1b, 0x97, 0xd0, 0x80, 0xc5, 0xc4, 0xc7, 0xb6, 0x11, 0xad, 0x1a, 0x91, 0xb9, 0x7f, 0xba,
	0x26, 0x3e, 0x3a, 0x87, 0x09, 0xe5, 0xaa, 0x90, 0xf8, 0xc4, 0x8c, 0x35, 0x00, 0xad, 0x60, 0x4a,
	0xbe, 0xd0, 0xe4, 0x9e, 0xe3, 0x89, 0xa1, 0x5b, 0xe4, 0xfc, 0x3c, 0x86, 0x45, 0x3f, 0x4d, 0x95,
	0x3f, 0x22, 0x0c, 0xa7, 0x51, 0x9d, 0xa6, 0x5c, 0x4a, 0x13, 0x63, 0xc6, 0x3a, 0x88, 0xd6, 0x30,
	0x23, 0xe1, 0xc3, 0x6b, 0x37, 0xcb, 0x44, 0x6b, 0xfd, 0x09, 0xa3, 0xe7, 0x00, 0xfa, 0x1c, 0xd5,
	0xbb, 0x82, 0xab, 0xd6, 0x63, 0x8f, 0x41, 0x0e, 0x9c, 0xf9, 0xfc, 0xe1, 0x2e, 0xe5, 0xb4, 0xbe,
	0xdf, 0x71, 0x61, 0xec, 0x4d, 0xd8, 0x5f, 0x1c, 0xba, 0x80, 0xc5, 0x27, 0xc9, 0x83, 0xef, 0x8a,
	0x8b, 0x22, 0xc9, 0x23, 0xea, 0xde, 0x18, 0xbb, 0x33, 0x36, 0xa4, 0xb5, 0x93, 0xdb, 0xd0, 0x4b,
	0xef, 0x32, 0x21, 0xf1, 0x74, 0x63, 0x6b, 0x27, 0x1d, 0xd6, 0x2f, 0xe0, 0x7d, 0xf8, 0x4c, 0x42,
	0x7c, 0xda, 0xbc, 0x80, 0x01, 0xe8, 0x19, 0xd8, 0xb7, 0x94, 0xe0, 0x99, 0xf9, 0xac, 0x3e, 0xea,
	0x9c, 0xbe, 0x54, 0x61, 0x29, 0x14, 0x9e, 0x1b, 0xb6, 0x83, 0xce, 0x6f, 0x0b, 0x96, 0x3e, 0xcf,
	0xff, 0xdb, 0x8e, 0xfb, 0x3d, 0x9c, 0x0c, 0x7a, 0x58, 0xc1, 0x94, 0xf1, 0x44, 0x96, 0x45, 0xd7,
	0x74, 0x83, 0x9c, 0x3d, 0x2c, 0xfa, 0x91, 0xfe, 0xbd, 0xe8, 0x61, 0x91, 0xf6, 0xb8, 0xc8, 0x97,
	0x3f, 0x2c, 0x00, 0x8f, 0x92, 0x77, 0x49, 0xba, 0xe7, 0x45, 0x86, 0xde, 0x02, 0x1c, 0x7e, 0x32,
	0xb4, 0xba, 0xd4, 0x2b, 0x35, 0xda, 0xa1, 0xf5, 0xf9, 0x88, 0xaf, 0xf2, 0x47, 0xe7, 0x48, 0xab,
	0x0f, 0xce, 0x5b, 0xf5, 0xa8, 0x9d, 0x56, 0x3d, 0x88, 0xe8, 0x1c, 0xed, 0xa6, 0x66, 0x77, 0x5f,
	0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xbb, 0xfd, 0x9c, 0xc8, 0x03, 0x00, 0x00,
}
