// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint.proto

package tendermint

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	github_com_hyperledger_burrow_binary "github.com/hyperledger/burrow/binary"
	github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type NodeInfo struct {
	ID                   github_com_hyperledger_burrow_crypto.Address  `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"ID"`
	ListenAddress        string                                        `protobuf:"bytes,2,opt,name=ListenAddress,proto3" json:"ListenAddress,omitempty"`
	Network              string                                        `protobuf:"bytes,3,opt,name=Network,proto3" json:"Network,omitempty"`
	Version              string                                        `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`
	Channels             github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,5,opt,name=Channels,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"Channels"`
	Moniker              string                                        `protobuf:"bytes,6,opt,name=Moniker,proto3" json:"Moniker,omitempty"`
	RPCAddress           string                                        `protobuf:"bytes,7,opt,name=RPCAddress,proto3" json:"RPCAddress,omitempty"`
	TxIndex              string                                        `protobuf:"bytes,8,opt,name=TxIndex,proto3" json:"TxIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_04f926c8da23c367, []int{0}
}
func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetListenAddress() string {
	if m != nil {
		return m.ListenAddress
	}
	return ""
}

func (m *NodeInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *NodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *NodeInfo) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *NodeInfo) GetRPCAddress() string {
	if m != nil {
		return m.RPCAddress
	}
	return ""
}

func (m *NodeInfo) GetTxIndex() string {
	if m != nil {
		return m.TxIndex
	}
	return ""
}

func (*NodeInfo) XXX_MessageName() string {
	return "tendermint.NodeInfo"
}
func init() {
	proto.RegisterType((*NodeInfo)(nil), "tendermint.NodeInfo")
	golang_proto.RegisterType((*NodeInfo)(nil), "tendermint.NodeInfo")
}

func init() { proto.RegisterFile("tendermint.proto", fileDescriptor_04f926c8da23c367) }
func init() { golang_proto.RegisterFile("tendermint.proto", fileDescriptor_04f926c8da23c367) }

var fileDescriptor_04f926c8da23c367 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0xed, 0xab, 0x80, 0x1b, 0x4d, 0x4c, 0x4f, 0x1b, 0x0f, 0x85, 0x18, 0x0f, 0x1c,
	0xa4, 0x4d, 0xfc, 0xf3, 0x01, 0x84, 0x1e, 0x6c, 0xa2, 0x44, 0x1b, 0xe3, 0xc1, 0x1b, 0xa5, 0x43,
	0xd9, 0x00, 0x3b, 0x64, 0x77, 0x1b, 0xe8, 0xb7, 0xf3, 0xc8, 0xd5, 0x9b, 0xf1, 0x40, 0x0c, 0x7c,
	0x11, 0xd3, 0x6d, 0x11, 0xbc, 0xe8, 0x6d, 0x7f, 0xcf, 0xb3, 0x33, 0xf3, 0xec, 0x0e, 0x3d, 0xd6,
	0x20, 0x62, 0x90, 0x13, 0x2e, 0xb4, 0x3b, 0x95, 0xa8, 0xd1, 0xa6, 0x5b, 0xe5, 0xa4, 0x95, 0x70,
	0x3d, 0x4c, 0x23, 0xb7, 0x8f, 0x13, 0x2f, 0xc1, 0x04, 0x3d, 0x73, 0x25, 0x4a, 0x07, 0x86, 0x0c,
	0x98, 0x53, 0x51, 0x7a, 0xfa, 0x66, 0xd1, 0x5a, 0x17, 0x63, 0x08, 0xc4, 0x00, 0x6d, 0x9f, 0x5a,
	0x81, 0xcf, 0x48, 0x83, 0x34, 0x0f, 0xdb, 0x57, 0x8b, 0x65, 0xfd, 0xdf, 0xc7, 0xb2, 0x7e, 0xbe,
	0xd3, 0x6f, 0x98, 0x4d, 0x41, 0x8e, 0x21, 0x4e, 0x40, 0x7a, 0x51, 0x2a, 0x25, 0xce, 0xbc, 0xbe,
	0xcc, 0xa6, 0x1a, 0xdd, 0x9b, 0x38, 0x96, 0xa0, 0x54, 0x68, 0x05, 0xbe, 0x7d, 0x46, 0x8f, 0xee,
	0xb8, 0xd2, 0x20, 0x4a, 0x91, 0x59, 0x0d, 0xd2, 0x3c, 0x08, 0x7f, 0x8a, 0x36, 0xa3, 0xd5, 0x2e,
	0xe8, 0x19, 0xca, 0x11, 0xfb, 0x6f, 0xfc, 0x0d, 0xe6, 0xce, 0x33, 0x48, 0xc5, 0x51, 0xb0, 0xbd,
	0xc2, 0x29, 0xd1, 0x7e, 0xa4, 0xb5, 0xce, 0xb0, 0x27, 0x04, 0x8c, 0x15, 0xdb, 0x37, 0x29, 0xaf,
	0xcb, 0x94, 0xad, 0xdf, 0x53, 0x46, 0x5c, 0xf4, 0x64, 0xe6, 0xde, 0xc2, 0xbc, 0x9d, 0x69, 0x50,
	0xe1, 0x77, 0x9b, 0x7c, 0xd8, 0x3d, 0x0a, 0x3e, 0x02, 0xc9, 0x2a, 0xc5, 0xb0, 0x12, 0x6d, 0x87,
	0xd2, 0xf0, 0xa1, 0xb3, 0x79, 0x43, 0xd5, 0x98, 0x3b, 0x4a, 0x5e, 0xf9, 0x34, 0x0f, 0x44, 0x0c,
	0x73, 0x56, 0x2b, 0x2a, 0x4b, 0x6c, 0xfb, 0xef, 0x2b, 0x87, 0x7c, 0xae, 0x1c, 0xf2, 0xba, 0x76,
	0xc8, 0x62, 0xed, 0x90, 0x97, 0x8b, 0x3f, 0x3e, 0x11, 0x85, 0x02, 0xa1, 0x52, 0xe5, 0x6d, 0x17,
	0x19, 0x55, 0xcc, 0x82, 0x2e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba, 0x9e, 0x0e, 0xac, 0xef,
	0x01, 0x00, 0x00,
}

func (m *NodeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovTendermint(uint64(l))
	l = len(m.ListenAddress)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	l = m.Channels.Size()
	n += 1 + l + sovTendermint(uint64(l))
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	l = len(m.RPCAddress)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	l = len(m.TxIndex)
	if l > 0 {
		n += 1 + l + sovTendermint(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTendermint(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTendermint(x uint64) (n int) {
	return sovTendermint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
