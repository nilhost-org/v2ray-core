package p2p

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Seed struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	PubId                string   `protobuf:"bytes,4,opt,name=pub_id,json=pubId,proto3" json:"pub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Seed) Reset()         { *m = Seed{} }
func (m *Seed) String() string { return proto.CompactTextString(m) }
func (*Seed) ProtoMessage()    {}
func (*Seed) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d9df4fd11cc3b, []int{0}
}

func (m *Seed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Seed.Unmarshal(m, b)
}
func (m *Seed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Seed.Marshal(b, m, deterministic)
}
func (m *Seed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Seed.Merge(m, src)
}
func (m *Seed) XXX_Size() int {
	return xxx_messageInfo_Seed.Size(m)
}
func (m *Seed) XXX_DiscardUnknown() {
	xxx_messageInfo_Seed.DiscardUnknown(m)
}

var xxx_messageInfo_Seed proto.InternalMessageInfo

func (m *Seed) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Seed) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Seed) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Seed) GetPubId() string {
	if m != nil {
		return m.PubId
	}
	return ""
}

type Config struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	PubId                string   `protobuf:"bytes,4,opt,name=pub_id,json=pubId,proto3" json:"pub_id,omitempty"`
	Seedlist             []*Seed  `protobuf:"bytes,5,rep,name=seedlist,proto3" json:"seedlist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d9df4fd11cc3b, []int{1}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Config) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Config) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Config) GetPubId() string {
	if m != nil {
		return m.PubId
	}
	return ""
}

func (m *Config) GetSeedlist() []*Seed {
	if m != nil {
		return m.Seedlist
	}
	return nil
}

func init() {
	proto.RegisterType((*Seed)(nil), "v2ray.core.app.p2p.seed")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.p2p.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/p2p/config.proto", fileDescriptor_5d7d9df4fd11cc3b)
}

var fileDescriptor_5d7d9df4fd11cc3b = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0xcf, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0x06, 0x70, 0x36, 0xff, 0x38, 0x47, 0xb0, 0x18, 0x50, 0x16, 0xab, 0x70, 0x36, 0xa9, 0x36,
	0xb0, 0xde, 0x0b, 0xe8, 0x55, 0x76, 0x21, 0x85, 0x85, 0x20, 0x92, 0x6c, 0xa2, 0x2c, 0xdc, 0x39,
	0x1f, 0x7b, 0x39, 0xe1, 0x1e, 0xc4, 0x97, 0xf0, 0x29, 0x25, 0x0b, 0xa6, 0xd1, 0xd6, 0x6e, 0x77,
	0xe6, 0x07, 0xf3, 0x7d, 0x74, 0xf3, 0x61, 0x43, 0x77, 0x32, 0x4e, 0xf6, 0xb5, 0x93, 0x30, 0xd6,
	0x1d, 0x50, 0xc3, 0xa2, 0x76, 0xf2, 0xfe, 0xea, 0xdf, 0x0c, 0x82, 0x4c, 0xc2, 0xfc, 0x83, 0xc2,
	0x68, 0x3a, 0xc0, 0xc0, 0x62, 0xfd, 0x4c, 0xd9, 0x61, 0x1c, 0x07, 0xbe, 0xa6, 0x55, 0x44, 0x4e,
	0x76, 0x5a, 0x95, 0xaa, 0x3a, 0x6b, 0x97, 0x3f, 0x5f, 0x50, 0xe2, 0xa1, 0x93, 0x38, 0x4d, 0x3c,
	0x98, 0x29, 0x83, 0x84, 0x49, 0xa7, 0xa5, 0xaa, 0xf2, 0x36, 0xbe, 0xf9, 0x92, 0x0a, 0x1c, 0xfb,
	0x17, 0x3f, 0xe8, 0x2c, 0xba, 0x1c, 0xc7, 0xfe, 0x61, 0x58, 0x7f, 0x2a, 0x2a, 0xb6, 0x31, 0xc3,
	0x3f, 0x5d, 0xe0, 0x0d, 0xad, 0xe6, 0x02, 0x3b, 0x7f, 0x98, 0x74, 0x5e, 0xa6, 0xd5, 0xb9, 0xd5,
	0xe6, 0x77, 0x4f, 0x33, 0x9b, 0x76, 0x91, 0xf7, 0x1b, 0xba, 0x72, 0xb2, 0xff, 0x03, 0x36, 0xea,
	0x29, 0x85, 0xc5, 0x57, 0xc2, 0x8f, 0xb6, 0xed, 0x4e, 0x66, 0x3b, 0xef, 0xee, 0x00, 0xd3, 0xd8,
	0xa6, 0x2f, 0x62, 0xe0, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x77, 0x1b, 0xb8, 0x6e,
	0x01, 0x00, 0x00,
}
