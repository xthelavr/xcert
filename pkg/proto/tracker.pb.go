// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracker.proto

package tracker

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_3047cc6308c72048, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type GVRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GVRequest) Reset()         { *m = GVRequest{} }
func (m *GVRequest) String() string { return proto.CompactTextString(m) }
func (*GVRequest) ProtoMessage()    {}
func (*GVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_3047cc6308c72048, []int{1}
}
func (m *GVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GVRequest.Unmarshal(m, b)
}
func (m *GVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GVRequest.Marshal(b, m, deterministic)
}
func (dst *GVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GVRequest.Merge(dst, src)
}
func (m *GVRequest) XXX_Size() int {
	return xxx_messageInfo_GVRequest.Size(m)
}
func (m *GVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GVRequest proto.InternalMessageInfo

type GVResponse struct {
	Ips                  map[string]*Empty `protobuf:"bytes,1,rep,name=ips" json:"ips,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GVResponse) Reset()         { *m = GVResponse{} }
func (m *GVResponse) String() string { return proto.CompactTextString(m) }
func (*GVResponse) ProtoMessage()    {}
func (*GVResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_3047cc6308c72048, []int{2}
}
func (m *GVResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GVResponse.Unmarshal(m, b)
}
func (m *GVResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GVResponse.Marshal(b, m, deterministic)
}
func (dst *GVResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GVResponse.Merge(dst, src)
}
func (m *GVResponse) XXX_Size() int {
	return xxx_messageInfo_GVResponse.Size(m)
}
func (m *GVResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GVResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GVResponse proto.InternalMessageInfo

func (m *GVResponse) GetIps() map[string]*Empty {
	if m != nil {
		return m.Ips
	}
	return nil
}

type NVRequest struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NVRequest) Reset()         { *m = NVRequest{} }
func (m *NVRequest) String() string { return proto.CompactTextString(m) }
func (*NVRequest) ProtoMessage()    {}
func (*NVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_3047cc6308c72048, []int{3}
}
func (m *NVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVRequest.Unmarshal(m, b)
}
func (m *NVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVRequest.Marshal(b, m, deterministic)
}
func (dst *NVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVRequest.Merge(dst, src)
}
func (m *NVRequest) XXX_Size() int {
	return xxx_messageInfo_NVRequest.Size(m)
}
func (m *NVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NVRequest proto.InternalMessageInfo

func (m *NVRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type NVResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NVResponse) Reset()         { *m = NVResponse{} }
func (m *NVResponse) String() string { return proto.CompactTextString(m) }
func (*NVResponse) ProtoMessage()    {}
func (*NVResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_3047cc6308c72048, []int{4}
}
func (m *NVResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NVResponse.Unmarshal(m, b)
}
func (m *NVResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NVResponse.Marshal(b, m, deterministic)
}
func (dst *NVResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NVResponse.Merge(dst, src)
}
func (m *NVResponse) XXX_Size() int {
	return xxx_messageInfo_NVResponse.Size(m)
}
func (m *NVResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NVResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NVResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "tracker.Empty")
	proto.RegisterType((*GVRequest)(nil), "tracker.GVRequest")
	proto.RegisterType((*GVResponse)(nil), "tracker.GVResponse")
	proto.RegisterMapType((map[string]*Empty)(nil), "tracker.GVResponse.IpsEntry")
	proto.RegisterType((*NVRequest)(nil), "tracker.NVRequest")
	proto.RegisterType((*NVResponse)(nil), "tracker.NVResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tracker service

type TrackerClient interface {
	GetValidators(ctx context.Context, in *GVRequest, opts ...grpc.CallOption) (*GVResponse, error)
	NewValidator(ctx context.Context, in *NVRequest, opts ...grpc.CallOption) (*NVResponse, error)
}

type trackerClient struct {
	cc *grpc.ClientConn
}

func NewTrackerClient(cc *grpc.ClientConn) TrackerClient {
	return &trackerClient{cc}
}

func (c *trackerClient) GetValidators(ctx context.Context, in *GVRequest, opts ...grpc.CallOption) (*GVResponse, error) {
	out := new(GVResponse)
	err := grpc.Invoke(ctx, "/tracker.Tracker/GetValidators", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) NewValidator(ctx context.Context, in *NVRequest, opts ...grpc.CallOption) (*NVResponse, error) {
	out := new(NVResponse)
	err := grpc.Invoke(ctx, "/tracker.Tracker/NewValidator", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tracker service

type TrackerServer interface {
	GetValidators(context.Context, *GVRequest) (*GVResponse, error)
	NewValidator(context.Context, *NVRequest) (*NVResponse, error)
}

func RegisterTrackerServer(s *grpc.Server, srv TrackerServer) {
	s.RegisterService(&_Tracker_serviceDesc, srv)
}

func _Tracker_GetValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).GetValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracker.Tracker/GetValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).GetValidators(ctx, req.(*GVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_NewValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).NewValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracker.Tracker/NewValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).NewValidator(ctx, req.(*NVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tracker.Tracker",
	HandlerType: (*TrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValidators",
			Handler:    _Tracker_GetValidators_Handler,
		},
		{
			MethodName: "NewValidator",
			Handler:    _Tracker_NewValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracker.proto",
}

func init() { proto.RegisterFile("tracker.proto", fileDescriptor_tracker_3047cc6308c72048) }

var fileDescriptor_tracker_3047cc6308c72048 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0xa2, 0x12, 0xf2, 0xd2, 0x56, 0xe8, 0x58, 0xa2, 0xc2, 0x10, 0x59, 0x0c, 0x99,
	0x32, 0x84, 0x05, 0xba, 0x97, 0x8a, 0xc5, 0x43, 0x84, 0xba, 0x07, 0xb8, 0x21, 0x6a, 0x69, 0x8c,
	0xed, 0x82, 0xb2, 0x31, 0xf0, 0xc3, 0x51, 0xd2, 0x26, 0x14, 0xa9, 0x9b, 0xcf, 0x7e, 0xef, 0xd3,
	0xe7, 0xc3, 0xc4, 0x99, 0xf2, 0x75, 0xcd, 0x26, 0xd3, 0xa6, 0x76, 0x35, 0x05, 0x87, 0x51, 0x06,
	0x18, 0x2d, 0xde, 0xb5, 0x6b, 0x64, 0x84, 0x70, 0xb9, 0x2a, 0xf8, 0x63, 0xc7, 0xd6, 0xc9, 0x1f,
	0x01, 0xb4, 0x93, 0xd5, 0xf5, 0xd6, 0x32, 0x65, 0xf0, 0x2b, 0x6d, 0x63, 0x91, 0xf8, 0x69, 0x94,
	0xdf, 0x64, 0x3d, 0xea, 0x2f, 0x91, 0x3d, 0x69, 0xbb, 0xd8, 0x3a, 0xd3, 0x14, 0x6d, 0x70, 0xf6,
	0x88, 0x8b, 0xfe, 0x82, 0x2e, 0xe1, 0xaf, 0xb9, 0x89, 0x45, 0x22, 0xd2, 0xb0, 0x68, 0x8f, 0x74,
	0x8b, 0xd1, 0x67, 0xb9, 0xd9, 0x71, 0xec, 0x25, 0x22, 0x8d, 0xf2, 0xe9, 0xc0, 0xeb, 0x44, 0x8a,
	0xfd, 0xe3, 0xdc, 0xbb, 0x17, 0xf2, 0x1a, 0xa1, 0xea, 0x9d, 0x68, 0x0a, 0xaf, 0xd2, 0x07, 0x8e,
	0x57, 0x69, 0x39, 0x06, 0xd4, 0x20, 0x90, 0x7f, 0x0b, 0x04, 0xcf, 0x7b, 0x0e, 0xcd, 0x31, 0x59,
	0xb2, 0x5b, 0x95, 0x9b, 0xea, 0xad, 0x74, 0xb5, 0xb1, 0x44, 0xff, 0x94, 0x3b, 0xdc, 0xec, 0xea,
	0xc4, 0x37, 0xe4, 0x19, 0x3d, 0x60, 0xac, 0xf8, 0x6b, 0xe8, 0x1e, 0x55, 0xd5, 0x89, 0xaa, 0x3a,
	0xaa, 0xbe, 0x9c, 0x77, 0xab, 0xbd, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xdd, 0x3d, 0x50,
	0x6b, 0x01, 0x00, 0x00,
}