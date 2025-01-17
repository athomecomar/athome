// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared/entity.proto

package pbshared

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Entity struct {
	EntityId             uint64   `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	EntityTable          string   `protobuf:"bytes,2,opt,name=entity_table,json=entityTable,proto3" json:"entity_table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f8e2dc6d1e0a88f, []int{0}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetEntityId() uint64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *Entity) GetEntityTable() string {
	if m != nil {
		return m.EntityTable
	}
	return ""
}

func init() {
	proto.RegisterType((*Entity)(nil), "shared.Entity")
}

func init() {
	proto.RegisterFile("shared/entity.proto", fileDescriptor_6f8e2dc6d1e0a88f)
}

var fileDescriptor_6f8e2dc6d1e0a88f = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xce, 0x48, 0x2c,
	0x4a, 0x4d, 0xd1, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x83, 0x08, 0x2a, 0x79, 0x70, 0xb1, 0xb9, 0x82, 0xc5, 0x85, 0xa4, 0xb9, 0x38, 0x21, 0x2a,
	0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x38, 0x20, 0x02, 0x9e, 0x29, 0x42,
	0x8a, 0x5c, 0x3c, 0x50, 0xc9, 0x92, 0xc4, 0xa4, 0x9c, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x6e, 0x88, 0x58, 0x08, 0x48, 0xc8, 0x49, 0x3b, 0x4a, 0x33, 0x3d, 0xb3, 0x24, 0xa3, 0x34,
	0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xb1, 0x24, 0x23, 0x3f, 0x37, 0x35, 0x39, 0x3f, 0x37, 0xb1,
	0x08, 0xca, 0xd6, 0x2f, 0x48, 0xd2, 0x2f, 0x48, 0x82, 0x58, 0x9b, 0xc4, 0x06, 0x76, 0x85, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x54, 0x16, 0x9d, 0xf1, 0x9c, 0x00, 0x00, 0x00,
}
