// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: txs.proto

package txs

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	github_com_hyperledger_burrow_binary "github.com/hyperledger/burrow/binary"
	crypto "github.com/hyperledger/burrow/crypto"
	github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
	github_com_hyperledger_burrow_txs_payload "github.com/hyperledger/burrow/txs/payload"
	io "io"
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

// An envelope contains both the signable Tx and the signatures for each input (in signatories)
type Envelope struct {
	Signatories []Signatory `protobuf:"bytes,1,rep,name=Signatories,proto3" json:"Signatories"`
	// Canonical bytes of the Tx ready to be signed
	Tx                   *Tx      `protobuf:"bytes,2,opt,name=Tx,proto3,customtype=Tx" json:"Tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()      { *m = Envelope{} }
func (*Envelope) ProtoMessage() {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_372ebcf753025bdc, []int{0}
}
func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return m.Size()
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetSignatories() []Signatory {
	if m != nil {
		return m.Signatories
	}
	return nil
}

func (*Envelope) XXX_MessageName() string {
	return "txs.Envelope"
}

// Signatory contains signature and one or both of Address and PublicKey to identify the signer
type Signatory struct {
	Address              *github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,1,opt,name=Address,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"Address,omitempty"`
	PublicKey            *crypto.PublicKey                             `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Signature            *crypto.Signature                             `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *Signatory) Reset()         { *m = Signatory{} }
func (m *Signatory) String() string { return proto.CompactTextString(m) }
func (*Signatory) ProtoMessage()    {}
func (*Signatory) Descriptor() ([]byte, []int) {
	return fileDescriptor_372ebcf753025bdc, []int{1}
}
func (m *Signatory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Signatory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Signatory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signatory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signatory.Merge(m, src)
}
func (m *Signatory) XXX_Size() int {
	return m.Size()
}
func (m *Signatory) XXX_DiscardUnknown() {
	xxx_messageInfo_Signatory.DiscardUnknown(m)
}

var xxx_messageInfo_Signatory proto.InternalMessageInfo

func (m *Signatory) GetPublicKey() *crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Signatory) GetSignature() *crypto.Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (*Signatory) XXX_MessageName() string {
	return "txs.Signatory"
}

// BroadcastTx or Transaction receipt
type Receipt struct {
	// Transaction type
	TxType github_com_hyperledger_burrow_txs_payload.Type `protobuf:"varint,1,opt,name=TxType,proto3,casttype=github.com/hyperledger/burrow/txs/payload.Type" json:"TxType,omitempty"`
	// The hash of the transaction that caused this event to be generated
	TxHash github_com_hyperledger_burrow_binary.HexBytes `protobuf:"bytes,2,opt,name=TxHash,proto3,customtype=github.com/hyperledger/burrow/binary.HexBytes" json:"TxHash"`
	// Whether the transaction creates a contract
	CreatesContract bool `protobuf:"varint,3,opt,name=CreatesContract,proto3" json:"CreatesContract,omitempty"`
	// The address of the contract being called
	ContractAddress      github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,4,opt,name=ContractAddress,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"ContractAddress"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_372ebcf753025bdc, []int{2}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return m.Size()
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetTxType() github_com_hyperledger_burrow_txs_payload.Type {
	if m != nil {
		return m.TxType
	}
	return 0
}

func (m *Receipt) GetCreatesContract() bool {
	if m != nil {
		return m.CreatesContract
	}
	return false
}

func (*Receipt) XXX_MessageName() string {
	return "txs.Receipt"
}
func init() {
	proto.RegisterType((*Envelope)(nil), "txs.Envelope")
	golang_proto.RegisterType((*Envelope)(nil), "txs.Envelope")
	proto.RegisterType((*Signatory)(nil), "txs.Signatory")
	golang_proto.RegisterType((*Signatory)(nil), "txs.Signatory")
	proto.RegisterType((*Receipt)(nil), "txs.Receipt")
	golang_proto.RegisterType((*Receipt)(nil), "txs.Receipt")
}

func init() { proto.RegisterFile("txs.proto", fileDescriptor_372ebcf753025bdc) }
func init() { golang_proto.RegisterFile("txs.proto", fileDescriptor_372ebcf753025bdc) }

var fileDescriptor_372ebcf753025bdc = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xee, 0x39, 0x51, 0xda, 0x5c, 0x0a, 0x15, 0x37, 0xa0, 0xa8, 0x83, 0x1d, 0x32, 0x65, 0xa0,
	0x36, 0x0a, 0xbf, 0x24, 0x36, 0x5c, 0x21, 0x55, 0x45, 0x48, 0xe8, 0xf0, 0xc4, 0x80, 0xb0, 0x9d,
	0x87, 0x63, 0xc9, 0xf8, 0xac, 0xbb, 0x33, 0xdc, 0xfd, 0x27, 0x8c, 0xfc, 0x05, 0xec, 0x6c, 0x8c,
	0x19, 0x99, 0x33, 0x58, 0x28, 0xfd, 0x2f, 0x98, 0x90, 0xaf, 0xe7, 0xb4, 0x74, 0x08, 0x62, 0x7b,
	0x3f, 0xbe, 0xef, 0x7b, 0xdf, 0xbd, 0x77, 0x78, 0x28, 0x95, 0xf0, 0x2b, 0xce, 0x24, 0x23, 0x3d,
	0xa9, 0xc4, 0xf1, 0x49, 0x96, 0xcb, 0x65, 0x9d, 0xf8, 0x29, 0xfb, 0x18, 0x64, 0x2c, 0x63, 0x81,
	0xe9, 0x25, 0xf5, 0x07, 0x93, 0x99, 0xc4, 0x44, 0x97, 0x9c, 0xe3, 0xc3, 0x94, 0xeb, 0x4a, 0xda,
	0x6c, 0xfa, 0x1e, 0x1f, 0xbc, 0x28, 0x3f, 0x41, 0xc1, 0x2a, 0x20, 0x4f, 0xf0, 0xe8, 0x4d, 0x9e,
	0x95, 0xb1, 0x64, 0x3c, 0x07, 0x31, 0x46, 0x93, 0xde, 0x6c, 0x34, 0xbf, 0xed, 0xb7, 0xe3, 0xba,
	0xba, 0x0e, 0xfb, 0xab, 0xc6, 0xdb, 0xa3, 0xd7, 0x81, 0xe4, 0x2e, 0x76, 0x22, 0x35, 0x76, 0x26,
	0x68, 0x76, 0x18, 0x0e, 0xd6, 0x8d, 0xe7, 0x44, 0x8a, 0x3a, 0x91, 0x7a, 0xd6, 0xff, 0xf2, 0xd5,
	0xdb, 0x9b, 0x7e, 0x47, 0x78, 0xb8, 0xa5, 0x93, 0x73, 0xbc, 0xff, 0x7c, 0xb1, 0xe0, 0x20, 0x5a,
	0xfd, 0x96, 0xf0, 0x60, 0xdd, 0x78, 0xf7, 0xaf, 0xbd, 0x60, 0xa9, 0x2b, 0xe0, 0x05, 0x2c, 0x32,
	0xe0, 0x41, 0x52, 0x73, 0xce, 0x3e, 0x07, 0xd6, 0xb0, 0xe5, 0xd1, 0x4e, 0x80, 0x04, 0x78, 0xf8,
	0xba, 0x4e, 0x8a, 0x3c, 0x7d, 0x09, 0xda, 0x8c, 0x1f, 0xcd, 0xef, 0xf8, 0x16, 0xbc, 0x6d, 0xd0,
	0x2b, 0x4c, 0x4b, 0xb8, 0x74, 0x52, 0x73, 0x18, 0xf7, 0xff, 0x26, 0x6c, 0x1b, 0xf4, 0x0a, 0x33,
	0xfd, 0xe6, 0xe0, 0x7d, 0x0a, 0x29, 0xe4, 0x95, 0x24, 0xe7, 0x78, 0x10, 0xa9, 0x48, 0x57, 0x60,
	0x8c, 0xdf, 0x0a, 0xe7, 0xbf, 0x1b, 0xcf, 0xdf, 0x6d, 0x5c, 0x2a, 0x11, 0x54, 0xb1, 0x2e, 0x58,
	0xbc, 0xf0, 0x5b, 0x26, 0xb5, 0x0a, 0xe4, 0x55, 0xab, 0x75, 0x16, 0x8b, 0xa5, 0xdd, 0xda, 0xe3,
	0x76, 0xa9, 0xeb, 0xc6, 0x3b, 0xd9, 0xad, 0x97, 0xe4, 0x65, 0xcc, 0xb5, 0x7f, 0x06, 0x2a, 0xd4,
	0x12, 0x04, 0xb5, 0x22, 0x64, 0x86, 0x8f, 0x4e, 0x39, 0xc4, 0x12, 0xc4, 0x29, 0x2b, 0x25, 0x8f,
	0x53, 0x39, 0xee, 0x4d, 0xd0, 0xec, 0x80, 0xde, 0x2c, 0x93, 0x77, 0xf8, 0xa8, 0x8b, 0xbb, 0x33,
	0xf4, 0x8d, 0x83, 0x47, 0xd6, 0xc1, 0xff, 0x9d, 0xe2, 0xa6, 0x58, 0xf8, 0x74, 0xb5, 0x71, 0xd1,
	0xcf, 0x8d, 0x8b, 0x7e, 0x6d, 0x5c, 0xf4, 0xe3, 0xc2, 0x45, 0xab, 0x0b, 0x17, 0xbd, 0xbd, 0xf7,
	0xcf, 0x35, 0x25, 0x03, 0xf3, 0x1d, 0x1f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xe2, 0x91,
	0x22, 0xdd, 0x02, 0x00, 0x00,
}

func (m *Envelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Envelope) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Signatories) > 0 {
		for _, msg := range m.Signatories {
			dAtA[i] = 0xa
			i++
			i = encodeVarintTxs(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Tx != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.Tx.Size()))
		n1, err := m.Tx.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Signatory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signatory) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Address != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.Address.Size()))
		n2, err := m.Address.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.PublicKey != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.PublicKey.Size()))
		n3, err := m.PublicKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Signature != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.Signature.Size()))
		n4, err := m.Signature.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Receipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TxType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTxs(dAtA, i, uint64(m.TxType))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintTxs(dAtA, i, uint64(m.TxHash.Size()))
	n5, err := m.TxHash.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.CreatesContract {
		dAtA[i] = 0x18
		i++
		if m.CreatesContract {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintTxs(dAtA, i, uint64(m.ContractAddress.Size()))
	n6, err := m.ContractAddress.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTxs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Envelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signatories) > 0 {
		for _, e := range m.Signatories {
			l = e.Size()
			n += 1 + l + sovTxs(uint64(l))
		}
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Signatory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.PublicKey != nil {
		l = m.PublicKey.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovTxs(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Receipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxType != 0 {
		n += 1 + sovTxs(uint64(m.TxType))
	}
	l = m.TxHash.Size()
	n += 1 + l + sovTxs(uint64(l))
	if m.CreatesContract {
		n += 2
	}
	l = m.ContractAddress.Size()
	n += 1 + l + sovTxs(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTxs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTxs(x uint64) (n int) {
	return sovTxs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Envelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatories = append(m.Signatories, Signatory{})
			if err := m.Signatories[len(m.Signatories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v Tx
			m.Tx = &v
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTxs
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
func (m *Signatory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Signatory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signatory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_hyperledger_burrow_crypto.Address
			m.Address = &v
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PublicKey == nil {
				m.PublicKey = &crypto.PublicKey{}
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signature == nil {
				m.Signature = &crypto.Signature{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTxs
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
func (m *Receipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Receipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			m.TxType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxType |= github_com_hyperledger_burrow_txs_payload.Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatesContract", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CreatesContract = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ContractAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTxs
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTxs
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
func skipTxs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxs
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
					return 0, ErrIntOverflowTxs
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
					return 0, ErrIntOverflowTxs
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
			if length < 0 {
				return 0, ErrInvalidLengthTxs
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTxs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTxs
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
				next, err := skipTxs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTxs
				}
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
	ErrInvalidLengthTxs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxs   = fmt.Errorf("proto: integer overflow")
)
