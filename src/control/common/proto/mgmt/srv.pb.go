// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv.proto

package mgmt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Identifier for server rank within DAOS pool
type DaosRank struct {
	PoolUuid             string   `protobuf:"bytes,1,opt,name=pool_uuid,json=poolUuid,proto3" json:"pool_uuid,omitempty"`
	Rank                 uint32   `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DaosRank) Reset()         { *m = DaosRank{} }
func (m *DaosRank) String() string { return proto.CompactTextString(m) }
func (*DaosRank) ProtoMessage()    {}
func (*DaosRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_8ffad8837fce5c06, []int{0}
}
func (m *DaosRank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DaosRank.Unmarshal(m, b)
}
func (m *DaosRank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DaosRank.Marshal(b, m, deterministic)
}
func (dst *DaosRank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DaosRank.Merge(dst, src)
}
func (m *DaosRank) XXX_Size() int {
	return xxx_messageInfo_DaosRank.Size(m)
}
func (m *DaosRank) XXX_DiscardUnknown() {
	xxx_messageInfo_DaosRank.DiscardUnknown(m)
}

var xxx_messageInfo_DaosRank proto.InternalMessageInfo

func (m *DaosRank) GetPoolUuid() string {
	if m != nil {
		return m.PoolUuid
	}
	return ""
}

func (m *DaosRank) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

type DaosResp struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DaosResp) Reset()         { *m = DaosResp{} }
func (m *DaosResp) String() string { return proto.CompactTextString(m) }
func (*DaosResp) ProtoMessage()    {}
func (*DaosResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_8ffad8837fce5c06, []int{1}
}
func (m *DaosResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DaosResp.Unmarshal(m, b)
}
func (m *DaosResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DaosResp.Marshal(b, m, deterministic)
}
func (dst *DaosResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DaosResp.Merge(dst, src)
}
func (m *DaosResp) XXX_Size() int {
	return xxx_messageInfo_DaosResp.Size(m)
}
func (m *DaosResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DaosResp.DiscardUnknown(m)
}

var xxx_messageInfo_DaosResp proto.InternalMessageInfo

func (m *DaosResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type SetRankReq struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRankReq) Reset()         { *m = SetRankReq{} }
func (m *SetRankReq) String() string { return proto.CompactTextString(m) }
func (*SetRankReq) ProtoMessage()    {}
func (*SetRankReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_8ffad8837fce5c06, []int{2}
}
func (m *SetRankReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRankReq.Unmarshal(m, b)
}
func (m *SetRankReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRankReq.Marshal(b, m, deterministic)
}
func (dst *SetRankReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRankReq.Merge(dst, src)
}
func (m *SetRankReq) XXX_Size() int {
	return xxx_messageInfo_SetRankReq.Size(m)
}
func (m *SetRankReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRankReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetRankReq proto.InternalMessageInfo

func (m *SetRankReq) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

type CreateMsReq struct {
	Bootstrap bool `protobuf:"varint,1,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	// Server UUID of this MS replica.
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Server management address of this MS replica.
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMsReq) Reset()         { *m = CreateMsReq{} }
func (m *CreateMsReq) String() string { return proto.CompactTextString(m) }
func (*CreateMsReq) ProtoMessage()    {}
func (*CreateMsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_srv_8ffad8837fce5c06, []int{3}
}
func (m *CreateMsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMsReq.Unmarshal(m, b)
}
func (m *CreateMsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMsReq.Marshal(b, m, deterministic)
}
func (dst *CreateMsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMsReq.Merge(dst, src)
}
func (m *CreateMsReq) XXX_Size() int {
	return xxx_messageInfo_CreateMsReq.Size(m)
}
func (m *CreateMsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMsReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMsReq proto.InternalMessageInfo

func (m *CreateMsReq) GetBootstrap() bool {
	if m != nil {
		return m.Bootstrap
	}
	return false
}

func (m *CreateMsReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CreateMsReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*DaosRank)(nil), "mgmt.DaosRank")
	proto.RegisterType((*DaosResp)(nil), "mgmt.DaosResp")
	proto.RegisterType((*SetRankReq)(nil), "mgmt.SetRankReq")
	proto.RegisterType((*CreateMsReq)(nil), "mgmt.CreateMsReq")
}

func init() { proto.RegisterFile("srv.proto", fileDescriptor_srv_8ffad8837fce5c06) }

var fileDescriptor_srv_8ffad8837fce5c06 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x0f, 0xc2, 0x20,
	0x10, 0x85, 0x43, 0xad, 0x4d, 0x39, 0xe3, 0xc2, 0x60, 0x9a, 0xe8, 0xd0, 0x30, 0x75, 0x72, 0x71,
	0x74, 0xd4, 0xd5, 0x85, 0xc6, 0xd9, 0xd0, 0x40, 0x4c, 0x53, 0x5b, 0x10, 0x0e, 0x7f, 0xbf, 0x01,
	0x1b, 0x75, 0x7b, 0x7c, 0xb9, 0xf7, 0xe5, 0x01, 0xd4, 0xbb, 0xd7, 0xde, 0x3a, 0x83, 0x86, 0xe5,
	0xe3, 0x7d, 0x44, 0x7e, 0x84, 0xf2, 0x2c, 0x8d, 0x17, 0x72, 0x1a, 0xd8, 0x16, 0xa8, 0x35, 0xe6,
	0x71, 0x0b, 0xa1, 0x57, 0x15, 0xa9, 0x49, 0x43, 0x45, 0x19, 0xc1, 0x35, 0xf4, 0x8a, 0x31, 0xc8,
	0x9d, 0x9c, 0x86, 0x2a, 0xab, 0x49, 0xb3, 0x16, 0x29, 0x73, 0x3e, 0x97, 0xb5, 0xb7, 0x6c, 0x03,
	0x85, 0x47, 0x89, 0xc1, 0xa7, 0xe6, 0x52, 0xcc, 0x2f, 0x5e, 0x03, 0xb4, 0x1a, 0xa3, 0x5f, 0xe8,
	0xe7, 0xd7, 0x42, 0xfe, 0x2c, 0x2d, 0xac, 0x4e, 0x4e, 0x4b, 0xd4, 0x17, 0x1f, 0x4f, 0x76, 0x40,
	0x3b, 0x63, 0xd0, 0xa3, 0x93, 0x36, 0xdd, 0x95, 0xe2, 0x07, 0xa2, 0x20, 0xcd, 0xcb, 0xd2, 0xbc,
	0x94, 0x23, 0x93, 0x4a, 0xb9, 0x6a, 0xf1, 0x61, 0x31, 0x77, 0x45, 0xfa, 0xe4, 0xe1, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x26, 0x65, 0x14, 0x6d, 0xf1, 0x00, 0x00, 0x00,
}