package wire

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ActionType int32

const (
	ActionType_SEED_ONLINE  ActionType = 0
	ActionType_SEED_OFFLINE ActionType = 1
	ActionType_SEED_SUSPEND ActionType = 2
	ActionType_SEED_UPDATE  ActionType = 3
)

var ActionType_name = map[int32]string{
	0: "SEED_ONLINE",
	1: "SEED_OFFLINE",
	2: "SEED_SUSPEND",
	3: "SEED_UPDATE",
}
var ActionType_value = map[string]int32{
	"SEED_ONLINE":  0,
	"SEED_OFFLINE": 1,
	"SEED_SUSPEND": 2,
	"SEED_UPDATE":  3,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_seedlist_53892d8491644e45, []int{0}
}

type MessageData struct {
	// shared between all requests
	ClientVersion        string   `protobuf:"bytes,1,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Gossip               bool     `protobuf:"varint,4,opt,name=gossip,proto3" json:"gossip,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	NodePubKey           []byte   `protobuf:"bytes,6,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Sign                 []byte   `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageData) Reset()      { *m = MessageData{} }
func (*MessageData) ProtoMessage() {}
func (*MessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_seedlist_53892d8491644e45, []int{0}
}
func (m *MessageData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *MessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageData.Merge(dst, src)
}
func (m *MessageData) XXX_Size() int {
	return m.Size()
}
func (m *MessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageData.DiscardUnknown(m)
}

var xxx_messageInfo_MessageData proto.InternalMessageInfo

func (m *MessageData) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *MessageData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessageData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MessageData) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *MessageData) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MessageData) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *MessageData) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

// a protocol define a set of reuqest and responses
type SeedListRequest struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// method specific data
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeedListRequest) Reset()      { *m = SeedListRequest{} }
func (*SeedListRequest) ProtoMessage() {}
func (*SeedListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_seedlist_53892d8491644e45, []int{1}
}
func (m *SeedListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SeedListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SeedListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SeedListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeedListRequest.Merge(dst, src)
}
func (m *SeedListRequest) XXX_Size() int {
	return m.Size()
}
func (m *SeedListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SeedListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SeedListRequest proto.InternalMessageInfo

func (m *SeedListRequest) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *SeedListRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SeedListResponse struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// response specific data
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeedListResponse) Reset()      { *m = SeedListResponse{} }
func (*SeedListResponse) ProtoMessage() {}
func (*SeedListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_seedlist_53892d8491644e45, []int{2}
}
func (m *SeedListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SeedListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SeedListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SeedListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeedListResponse.Merge(dst, src)
}
func (m *SeedListResponse) XXX_Size() int {
	return m.Size()
}
func (m *SeedListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SeedListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SeedListResponse proto.InternalMessageInfo

func (m *SeedListResponse) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *SeedListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SeedInfo struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	HostID               string   `protobuf:"bytes,4,opt,name=hostID,proto3" json:"hostID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeedInfo) Reset()      { *m = SeedInfo{} }
func (*SeedInfo) ProtoMessage() {}
func (*SeedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_seedlist_53892d8491644e45, []int{3}
}
func (m *SeedInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SeedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SeedInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SeedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeedInfo.Merge(dst, src)
}
func (m *SeedInfo) XXX_Size() int {
	return m.Size()
}
func (m *SeedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SeedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SeedInfo proto.InternalMessageInfo

func (m *SeedInfo) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *SeedInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *SeedInfo) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *SeedInfo) GetHostID() string {
	if m != nil {
		return m.HostID
	}
	return ""
}

type HelloSeedList struct {
	Action               ActionType `protobuf:"varint,1,opt,name=action,proto3,enum=wire.ActionType" json:"action,omitempty"`
	Seed                 *SeedInfo  `protobuf:"bytes,2,opt,name=seed" json:"seed,omitempty"`
	Nonce                string     `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HelloSeedList) Reset()      { *m = HelloSeedList{} }
func (*HelloSeedList) ProtoMessage() {}
func (*HelloSeedList) Descriptor() ([]byte, []int) {
	return fileDescriptor_seedlist_53892d8491644e45, []int{4}
}
func (m *HelloSeedList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloSeedList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloSeedList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HelloSeedList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloSeedList.Merge(dst, src)
}
func (m *HelloSeedList) XXX_Size() int {
	return m.Size()
}
func (m *HelloSeedList) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloSeedList.DiscardUnknown(m)
}

var xxx_messageInfo_HelloSeedList proto.InternalMessageInfo

func (m *HelloSeedList) GetAction() ActionType {
	if m != nil {
		return m.Action
	}
	return ActionType_SEED_ONLINE
}

func (m *HelloSeedList) GetSeed() *SeedInfo {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *HelloSeedList) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageData)(nil), "wire.MessageData")
	proto.RegisterType((*SeedListRequest)(nil), "wire.SeedListRequest")
	proto.RegisterType((*SeedListResponse)(nil), "wire.SeedListResponse")
	proto.RegisterType((*SeedInfo)(nil), "wire.SeedInfo")
	proto.RegisterType((*HelloSeedList)(nil), "wire.HelloSeedList")
	proto.RegisterEnum("wire.ActionType", ActionType_name, ActionType_value)
}
func (this *MessageData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MessageData)
	if !ok {
		that2, ok := that.(MessageData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ClientVersion != that1.ClientVersion {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Gossip != that1.Gossip {
		return false
	}
	if this.NodeId != that1.NodeId {
		return false
	}
	if !bytes.Equal(this.NodePubKey, that1.NodePubKey) {
		return false
	}
	if !bytes.Equal(this.Sign, that1.Sign) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SeedListRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SeedListRequest)
	if !ok {
		that2, ok := that.(SeedListRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MessageData.Equal(that1.MessageData) {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SeedListResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SeedListResponse)
	if !ok {
		that2, ok := that.(SeedListResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MessageData.Equal(that1.MessageData) {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SeedInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SeedInfo)
	if !ok {
		that2, ok := that.(SeedInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if this.Ip != that1.Ip {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if this.HostID != that1.HostID {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HelloSeedList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelloSeedList)
	if !ok {
		that2, ok := that.(HelloSeedList)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Action != that1.Action {
		return false
	}
	if !this.Seed.Equal(that1.Seed) {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MessageData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&wire.MessageData{")
	s = append(s, "ClientVersion: "+fmt.Sprintf("%#v", this.ClientVersion)+",\n")
	s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Gossip: "+fmt.Sprintf("%#v", this.Gossip)+",\n")
	s = append(s, "NodeId: "+fmt.Sprintf("%#v", this.NodeId)+",\n")
	s = append(s, "NodePubKey: "+fmt.Sprintf("%#v", this.NodePubKey)+",\n")
	s = append(s, "Sign: "+fmt.Sprintf("%#v", this.Sign)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SeedListRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&wire.SeedListRequest{")
	if this.MessageData != nil {
		s = append(s, "MessageData: "+fmt.Sprintf("%#v", this.MessageData)+",\n")
	}
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SeedListResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&wire.SeedListResponse{")
	if this.MessageData != nil {
		s = append(s, "MessageData: "+fmt.Sprintf("%#v", this.MessageData)+",\n")
	}
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SeedInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&wire.SeedInfo{")
	s = append(s, "Protocol: "+fmt.Sprintf("%#v", this.Protocol)+",\n")
	s = append(s, "Ip: "+fmt.Sprintf("%#v", this.Ip)+",\n")
	s = append(s, "Port: "+fmt.Sprintf("%#v", this.Port)+",\n")
	s = append(s, "HostID: "+fmt.Sprintf("%#v", this.HostID)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HelloSeedList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&wire.HelloSeedList{")
	s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	if this.Seed != nil {
		s = append(s, "Seed: "+fmt.Sprintf("%#v", this.Seed)+",\n")
	}
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSeedlist(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *MessageData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClientVersion) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.ClientVersion)))
		i += copy(dAtA[i:], m.ClientVersion)
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(m.Timestamp))
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Gossip {
		dAtA[i] = 0x20
		i++
		if m.Gossip {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.NodeId) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.NodeId)))
		i += copy(dAtA[i:], m.NodeId)
	}
	if len(m.NodePubKey) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.NodePubKey)))
		i += copy(dAtA[i:], m.NodePubKey)
	}
	if len(m.Sign) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.Sign)))
		i += copy(dAtA[i:], m.Sign)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SeedListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SeedListRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MessageData != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(m.MessageData.Size()))
		n1, err := m.MessageData.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SeedListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SeedListResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MessageData != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(m.MessageData.Size()))
		n2, err := m.MessageData.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SeedInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SeedInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Protocol) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.Protocol)))
		i += copy(dAtA[i:], m.Protocol)
	}
	if len(m.Ip) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	if m.Port != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(m.Port))
	}
	if len(m.HostID) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.HostID)))
		i += copy(dAtA[i:], m.HostID)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *HelloSeedList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloSeedList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Action != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(m.Action))
	}
	if m.Seed != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(m.Seed.Size()))
		n3, err := m.Seed.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSeedlist(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSeedlist(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedMessageData(r randySeedlist, easy bool) *MessageData {
	this := &MessageData{}
	this.ClientVersion = string(randStringSeedlist(r))
	this.Timestamp = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Timestamp *= -1
	}
	this.Id = string(randStringSeedlist(r))
	this.Gossip = bool(bool(r.Intn(2) == 0))
	this.NodeId = string(randStringSeedlist(r))
	v1 := r.Intn(100)
	this.NodePubKey = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.NodePubKey[i] = byte(r.Intn(256))
	}
	v2 := r.Intn(100)
	this.Sign = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Sign[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSeedlist(r, 8)
	}
	return this
}

func NewPopulatedSeedListRequest(r randySeedlist, easy bool) *SeedListRequest {
	this := &SeedListRequest{}
	if r.Intn(10) != 0 {
		this.MessageData = NewPopulatedMessageData(r, easy)
	}
	this.Message = string(randStringSeedlist(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSeedlist(r, 3)
	}
	return this
}

func NewPopulatedSeedListResponse(r randySeedlist, easy bool) *SeedListResponse {
	this := &SeedListResponse{}
	if r.Intn(10) != 0 {
		this.MessageData = NewPopulatedMessageData(r, easy)
	}
	this.Message = string(randStringSeedlist(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSeedlist(r, 3)
	}
	return this
}

func NewPopulatedSeedInfo(r randySeedlist, easy bool) *SeedInfo {
	this := &SeedInfo{}
	this.Protocol = string(randStringSeedlist(r))
	this.Ip = string(randStringSeedlist(r))
	this.Port = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Port *= -1
	}
	this.HostID = string(randStringSeedlist(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSeedlist(r, 5)
	}
	return this
}

func NewPopulatedHelloSeedList(r randySeedlist, easy bool) *HelloSeedList {
	this := &HelloSeedList{}
	this.Action = ActionType([]int32{0, 1, 2, 3}[r.Intn(4)])
	if r.Intn(10) != 0 {
		this.Seed = NewPopulatedSeedInfo(r, easy)
	}
	this.Nonce = string(randStringSeedlist(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSeedlist(r, 4)
	}
	return this
}

type randySeedlist interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSeedlist(r randySeedlist) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSeedlist(r randySeedlist) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneSeedlist(r)
	}
	return string(tmps)
}
func randUnrecognizedSeedlist(r randySeedlist, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSeedlist(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSeedlist(dAtA []byte, r randySeedlist, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSeedlist(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateSeedlist(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateSeedlist(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSeedlist(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSeedlist(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSeedlist(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSeedlist(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *MessageData) Size() (n int) {
	var l int
	_ = l
	l = len(m.ClientVersion)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovSeedlist(uint64(m.Timestamp))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	if m.Gossip {
		n += 2
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	l = len(m.NodePubKey)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SeedListRequest) Size() (n int) {
	var l int
	_ = l
	if m.MessageData != nil {
		l = m.MessageData.Size()
		n += 1 + l + sovSeedlist(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SeedListResponse) Size() (n int) {
	var l int
	_ = l
	if m.MessageData != nil {
		l = m.MessageData.Size()
		n += 1 + l + sovSeedlist(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SeedInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovSeedlist(uint64(m.Port))
	}
	l = len(m.HostID)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HelloSeedList) Size() (n int) {
	var l int
	_ = l
	if m.Action != 0 {
		n += 1 + sovSeedlist(uint64(m.Action))
	}
	if m.Seed != nil {
		l = m.Seed.Size()
		n += 1 + l + sovSeedlist(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovSeedlist(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSeedlist(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSeedlist(x uint64) (n int) {
	return sovSeedlist(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MessageData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MessageData{`,
		`ClientVersion:` + fmt.Sprintf("%v", this.ClientVersion) + `,`,
		`Timestamp:` + fmt.Sprintf("%v", this.Timestamp) + `,`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Gossip:` + fmt.Sprintf("%v", this.Gossip) + `,`,
		`NodeId:` + fmt.Sprintf("%v", this.NodeId) + `,`,
		`NodePubKey:` + fmt.Sprintf("%v", this.NodePubKey) + `,`,
		`Sign:` + fmt.Sprintf("%v", this.Sign) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SeedListRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SeedListRequest{`,
		`MessageData:` + strings.Replace(fmt.Sprintf("%v", this.MessageData), "MessageData", "MessageData", 1) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SeedListResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SeedListResponse{`,
		`MessageData:` + strings.Replace(fmt.Sprintf("%v", this.MessageData), "MessageData", "MessageData", 1) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SeedInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SeedInfo{`,
		`Protocol:` + fmt.Sprintf("%v", this.Protocol) + `,`,
		`Ip:` + fmt.Sprintf("%v", this.Ip) + `,`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`HostID:` + fmt.Sprintf("%v", this.HostID) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HelloSeedList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HelloSeedList{`,
		`Action:` + fmt.Sprintf("%v", this.Action) + `,`,
		`Seed:` + strings.Replace(fmt.Sprintf("%v", this.Seed), "SeedInfo", "SeedInfo", 1) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSeedlist(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MessageData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSeedlist
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gossip", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Gossip = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodePubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodePubKey = append(m.NodePubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.NodePubKey == nil {
				m.NodePubKey = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = append(m.Sign[:0], dAtA[iNdEx:postIndex]...)
			if m.Sign == nil {
				m.Sign = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSeedlist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSeedlist
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SeedListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSeedlist
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SeedListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeedListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MessageData == nil {
				m.MessageData = &MessageData{}
			}
			if err := m.MessageData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSeedlist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSeedlist
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SeedListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSeedlist
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SeedListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeedListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MessageData == nil {
				m.MessageData = &MessageData{}
			}
			if err := m.MessageData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSeedlist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSeedlist
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SeedInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSeedlist
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SeedInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeedInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSeedlist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSeedlist
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HelloSeedList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSeedlist
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloSeedList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloSeedList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= (ActionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Seed == nil {
				m.Seed = &SeedInfo{}
			}
			if err := m.Seed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSeedlist
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSeedlist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSeedlist
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSeedlist(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSeedlist
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSeedlist
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSeedlist
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSeedlist
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSeedlist(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSeedlist = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSeedlist   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("seedlist.proto", fileDescriptor_seedlist_53892d8491644e45) }

var fileDescriptor_seedlist_53892d8491644e45 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xbb, 0xf9, 0x6a, 0x32, 0x69, 0x53, 0xb3, 0x42, 0xc8, 0xaa, 0xd0, 0x2a, 0xb2, 0x7a,
	0xb0, 0x90, 0x48, 0xa5, 0xf4, 0x09, 0x8a, 0xec, 0x8a, 0x88, 0x12, 0xa2, 0x4d, 0xcb, 0x15, 0xfc,
	0xb1, 0x75, 0x57, 0x72, 0xbc, 0x26, 0xbb, 0x11, 0xea, 0x8d, 0xc7, 0xe0, 0x11, 0x78, 0x11, 0x24,
	0x8e, 0x1c, 0x39, 0x12, 0xf3, 0x02, 0x1c, 0x39, 0x22, 0x8f, 0x9d, 0x0f, 0xee, 0xdc, 0xe6, 0xf7,
	0x9f, 0x19, 0xcf, 0xce, 0x7f, 0x0c, 0x03, 0x2d, 0x44, 0x9c, 0x4a, 0x6d, 0x46, 0xf9, 0x52, 0x19,
	0x45, 0x5b, 0x1f, 0xe5, 0x52, 0x9c, 0x3e, 0x4f, 0xa4, 0xb9, 0x5f, 0x85, 0xa3, 0x48, 0x2d, 0xce,
	0x13, 0x95, 0xa8, 0x73, 0x4c, 0x86, 0xab, 0x3b, 0x24, 0x04, 0x8c, 0xaa, 0x26, 0xe7, 0x2b, 0x81,
	0xfe, 0x6b, 0xa1, 0x75, 0x90, 0x08, 0x2f, 0x30, 0x01, 0x3d, 0x83, 0xe3, 0x28, 0x95, 0x22, 0x33,
	0x6f, 0xc5, 0x52, 0x4b, 0x95, 0xd9, 0x64, 0x48, 0xdc, 0x1e, 0xff, 0x57, 0xa4, 0x4f, 0xa1, 0x67,
	0xe4, 0x42, 0x68, 0x13, 0x2c, 0x72, 0xbb, 0x31, 0x24, 0x6e, 0x93, 0xef, 0x04, 0x3a, 0x80, 0x86,
	0x8c, 0xed, 0x26, 0x36, 0x36, 0x64, 0x4c, 0x9f, 0x40, 0x27, 0x51, 0x5a, 0xcb, 0xdc, 0x6e, 0x0d,
	0x89, 0xdb, 0xe5, 0x35, 0x95, 0x7a, 0xa6, 0x62, 0x31, 0x89, 0xed, 0x36, 0xd6, 0xd6, 0x44, 0x19,
	0x40, 0x19, 0xcd, 0x56, 0xe1, 0x2b, 0xf1, 0x60, 0x77, 0x86, 0xc4, 0x3d, 0xe2, 0x7b, 0x0a, 0xa5,
	0xd0, 0xd2, 0x32, 0xc9, 0xec, 0x43, 0xcc, 0x60, 0xec, 0xbc, 0x87, 0x93, 0xb9, 0x10, 0xf1, 0xb5,
	0xd4, 0x86, 0x8b, 0x0f, 0x2b, 0xa1, 0x0d, 0xbd, 0x80, 0xfe, 0x62, 0xb7, 0x19, 0x2e, 0xd2, 0x1f,
	0x3f, 0x1a, 0x95, 0x2e, 0x8d, 0xf6, 0x56, 0xe6, 0xfb, 0x55, 0xd4, 0x86, 0xc3, 0x1a, 0x71, 0xaf,
	0x1e, 0xdf, 0xa0, 0x13, 0x80, 0xb5, 0x9b, 0xa0, 0x73, 0x95, 0x69, 0xf1, 0xbf, 0x47, 0x84, 0xd0,
	0x2d, 0x47, 0x4c, 0xb2, 0x3b, 0x45, 0x4f, 0xa1, 0x8b, 0x17, 0x8a, 0x54, 0x5a, 0xdf, 0x60, 0xcb,
	0x68, 0x70, 0x5e, 0x37, 0x37, 0x64, 0x5e, 0x1a, 0x92, 0xab, 0xa5, 0x41, 0xcb, 0xdb, 0x1c, 0xe3,
	0xd2, 0xdc, 0x7b, 0xa5, 0xcd, 0xc4, 0x43, 0xd3, 0x7b, 0xbc, 0x26, 0x47, 0xc3, 0xf1, 0x4b, 0x91,
	0xa6, 0x6a, 0xb3, 0x0b, 0x75, 0xa1, 0x13, 0x44, 0x66, 0x73, 0xea, 0xc1, 0xd8, 0xaa, 0x9e, 0x7f,
	0x89, 0xda, 0xcd, 0x43, 0x2e, 0x78, 0x9d, 0xa7, 0x0e, 0xb4, 0xca, 0x5f, 0x0e, 0x07, 0xf7, 0xc7,
	0x83, 0xaa, 0x6e, 0xf3, 0x60, 0x8e, 0x39, 0xfa, 0x18, 0xda, 0x99, 0xca, 0x22, 0x51, 0x9f, 0xbf,
	0x82, 0x67, 0x1c, 0x60, 0xf7, 0x3d, 0x7a, 0x02, 0xfd, 0xb9, 0xef, 0x7b, 0xef, 0xde, 0x4c, 0xaf,
	0x27, 0x53, 0xdf, 0x3a, 0xa0, 0x16, 0x1c, 0x55, 0xc2, 0xd5, 0x15, 0x2a, 0x64, 0xab, 0xcc, 0x6f,
	0xe7, 0x33, 0x7f, 0xea, 0x59, 0x8d, 0x6d, 0xd3, 0xed, 0xcc, 0xbb, 0xbc, 0xf1, 0xad, 0xe6, 0x8b,
	0xb3, 0x1f, 0x6b, 0x76, 0xf0, 0x7b, 0xcd, 0xc8, 0x9f, 0x35, 0x23, 0x9f, 0x0a, 0x46, 0xbe, 0x14,
	0x8c, 0x7c, 0x2b, 0x18, 0xf9, 0x5e, 0x30, 0xf2, 0xb3, 0x60, 0xe4, 0xf3, 0x2f, 0x76, 0x10, 0x76,
	0xd0, 0xb4, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0xdb, 0x48, 0x94, 0x2d, 0x03, 0x00,
	0x00,
}
