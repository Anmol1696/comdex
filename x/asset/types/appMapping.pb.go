// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/appMapping.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AppMapping struct {
	Id               uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string                                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	ShortName        string                                 `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty" yaml:"short_name"`
	MinGovDeposit    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=min_gov_deposit,json=minGovDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_gov_deposit" yaml:"min_gov_deposit"`
	GovTimeInSeconds float64                                `protobuf:"fixed64,5,opt,name=gov_time_in_seconds,json=govTimeInSeconds,proto3" json:"gov_time_in_seconds,omitempty" yaml:"gov_time_in_seconds"`
	GenesisToken     []MintGenesisToken                     `protobuf:"bytes,6,rep,name=genesis_token,json=genesisToken,proto3" json:"genesis_token" yaml:"genesis_token"`
}

func (m *AppMapping) Reset()         { *m = AppMapping{} }
func (m *AppMapping) String() string { return proto.CompactTextString(m) }
func (*AppMapping) ProtoMessage()    {}
func (*AppMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfd2a3ad9e2bc5f, []int{0}
}
func (m *AppMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppMapping.Merge(m, src)
}
func (m *AppMapping) XXX_Size() int {
	return m.Size()
}
func (m *AppMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_AppMapping.DiscardUnknown(m)
}

var xxx_messageInfo_AppMapping proto.InternalMessageInfo

type MintGenesisToken struct {
	AssetId       uint64                                  `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	GenesisSupply *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=genesis_supply,json=genesisSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"genesis_supply,omitempty"`
	IsGovToken    bool                                    `protobuf:"varint,3,opt,name=is_gov_token,json=isGovToken,proto3" json:"is_gov_token,omitempty" yaml:"is_gov_token"`
	Recipient     string                                  `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
}

func (m *MintGenesisToken) Reset()         { *m = MintGenesisToken{} }
func (m *MintGenesisToken) String() string { return proto.CompactTextString(m) }
func (*MintGenesisToken) ProtoMessage()    {}
func (*MintGenesisToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfd2a3ad9e2bc5f, []int{1}
}
func (m *MintGenesisToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintGenesisToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintGenesisToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintGenesisToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintGenesisToken.Merge(m, src)
}
func (m *MintGenesisToken) XXX_Size() int {
	return m.Size()
}
func (m *MintGenesisToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MintGenesisToken.DiscardUnknown(m)
}

var xxx_messageInfo_MintGenesisToken proto.InternalMessageInfo

type AppAndGovTime struct {
	AppId            uint64  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	GovTimeInSeconds float64 `protobuf:"fixed64,2,opt,name=gov_time_in_seconds,json=govTimeInSeconds,proto3" json:"gov_time_in_seconds,omitempty" yaml:"gov_time_in_seconds"`
}

func (m *AppAndGovTime) Reset()         { *m = AppAndGovTime{} }
func (m *AppAndGovTime) String() string { return proto.CompactTextString(m) }
func (*AppAndGovTime) ProtoMessage()    {}
func (*AppAndGovTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfd2a3ad9e2bc5f, []int{2}
}
func (m *AppAndGovTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppAndGovTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppAndGovTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppAndGovTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppAndGovTime.Merge(m, src)
}
func (m *AppAndGovTime) XXX_Size() int {
	return m.Size()
}
func (m *AppAndGovTime) XXX_DiscardUnknown() {
	xxx_messageInfo_AppAndGovTime.DiscardUnknown(m)
}

var xxx_messageInfo_AppAndGovTime proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AppMapping)(nil), "comdex.asset.v1beta1.AppMapping")
	proto.RegisterType((*MintGenesisToken)(nil), "comdex.asset.v1beta1.MintGenesisToken")
	proto.RegisterType((*AppAndGovTime)(nil), "comdex.asset.v1beta1.AppAndGovTime")
}

func init() {
	proto.RegisterFile("comdex/asset/v1beta1/appMapping.proto", fileDescriptor_2cfd2a3ad9e2bc5f)
}

var fileDescriptor_2cfd2a3ad9e2bc5f = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x93, 0xb6, 0x2b, 0xab, 0xb7, 0xae, 0x5d, 0x5a, 0xa0, 0x9a, 0x50, 0x52, 0x19, 0x31,
	0x55, 0x48, 0x4b, 0xb4, 0xc1, 0x05, 0x6e, 0x8d, 0x90, 0x4a, 0x0f, 0x45, 0x5a, 0xb6, 0x13, 0x97,
	0x28, 0x4d, 0xbc, 0xcc, 0x5a, 0x63, 0x5b, 0x75, 0x56, 0xd1, 0x3b, 0x07, 0x8e, 0xbc, 0x03, 0x17,
	0x1e, 0xa5, 0xc7, 0x1d, 0x11, 0x87, 0x08, 0xda, 0x37, 0xc8, 0x13, 0xa0, 0xd8, 0xc9, 0x5a, 0xa6,
	0x21, 0xa1, 0x9d, 0x62, 0xfb, 0xff, 0xf3, 0xf7, 0xb7, 0xbf, 0x2f, 0x06, 0x2f, 0x7c, 0x1a, 0x05,
	0xe8, 0x93, 0xe5, 0x71, 0x8e, 0x62, 0x6b, 0x76, 0x3c, 0x46, 0xb1, 0x77, 0x6c, 0x79, 0x8c, 0x8d,
	0x3c, 0xc6, 0x30, 0x09, 0x4d, 0x36, 0xa5, 0x31, 0xd5, 0xda, 0x12, 0x33, 0x05, 0x66, 0xe6, 0xd8,
	0x41, 0x3b, 0xa4, 0x21, 0x15, 0x80, 0x95, 0x8d, 0x24, 0x0b, 0xbf, 0x95, 0x01, 0xe8, 0xdf, 0x0a,
	0x68, 0x7b, 0xa0, 0x84, 0x83, 0x8e, 0xda, 0x55, 0x7b, 0x15, 0xa7, 0x84, 0x03, 0xed, 0x39, 0xa8,
	0x10, 0x2f, 0x42, 0x9d, 0x52, 0x57, 0xed, 0xd5, 0xec, 0x46, 0x9a, 0x18, 0x3b, 0x73, 0x2f, 0x9a,
	0xbc, 0x85, 0xd9, 0x2a, 0x74, 0x44, 0x51, 0x7b, 0x0d, 0x00, 0xbf, 0xa4, 0xd3, 0xd8, 0x15, 0x68,
	0x59, 0xa0, 0x8f, 0xd3, 0xc4, 0xd8, 0x97, 0xe8, 0xba, 0x06, 0x9d, 0x9a, 0x98, 0x7c, 0xc8, 0x76,
	0x31, 0xd0, 0x88, 0x30, 0x71, 0x43, 0x3a, 0x73, 0x03, 0xc4, 0x28, 0xc7, 0x71, 0xa7, 0x22, 0xb6,
	0xbe, 0x5f, 0x24, 0x86, 0xf2, 0x33, 0x31, 0x0e, 0x43, 0x1c, 0x5f, 0x5e, 0x8f, 0x4d, 0x9f, 0x46,
	0x96, 0x4f, 0x79, 0x44, 0x79, 0xfe, 0x39, 0xe2, 0xc1, 0x95, 0x15, 0xcf, 0x19, 0xe2, 0xe6, 0x90,
	0xc4, 0x69, 0x62, 0x3c, 0x91, 0x8d, 0xee, 0xc8, 0x41, 0xa7, 0x1e, 0x61, 0x32, 0xa0, 0xb3, 0x77,
	0x72, 0xae, 0x8d, 0x40, 0x2b, 0x2b, 0xc7, 0x38, 0x42, 0x2e, 0x26, 0x2e, 0x47, 0x3e, 0x25, 0x01,
	0xef, 0x6c, 0x75, 0xd5, 0x9e, 0x6a, 0xeb, 0x69, 0x62, 0x1c, 0x48, 0x9d, 0x7b, 0x20, 0xe8, 0x34,
	0x43, 0x3a, 0x3b, 0xc7, 0x11, 0x1a, 0x92, 0x33, 0xb9, 0xa4, 0x61, 0x50, 0x0f, 0x11, 0x41, 0x1c,
	0x73, 0x37, 0xa6, 0x57, 0x88, 0x74, 0xaa, 0xdd, 0x72, 0x6f, 0xe7, 0xe4, 0xd0, 0xbc, 0xcf, 0x7e,
	0x73, 0x84, 0x49, 0x3c, 0x90, 0xf8, 0x79, 0x46, 0xdb, 0xcf, 0xb2, 0x6b, 0xa6, 0x89, 0xd1, 0xce,
	0x9b, 0x6e, 0x4a, 0x41, 0x67, 0x37, 0xdc, 0x60, 0xe1, 0xe7, 0x12, 0x68, 0xde, 0x15, 0xd0, 0x4c,
	0xb0, 0x2d, 0x5a, 0xb8, 0x45, 0x62, 0x76, 0x2b, 0x4d, 0x8c, 0x86, 0x94, 0x2b, 0x2a, 0xd0, 0x79,
	0x24, 0x86, 0xc3, 0x40, 0x3b, 0x05, 0x7b, 0x45, 0x13, 0x7e, 0xcd, 0xd8, 0x64, 0x9e, 0xa7, 0xfa,
	0xf2, 0xff, 0xbd, 0x76, 0x8a, 0x1b, 0x9f, 0x09, 0x01, 0xed, 0x0d, 0xd8, 0xc5, 0x5c, 0x78, 0x2e,
	0x1d, 0xc8, 0xb2, 0xdf, 0xb6, 0x9f, 0xa6, 0x89, 0xd1, 0x92, 0xc7, 0xd8, 0xac, 0x42, 0x07, 0x60,
	0x3e, 0xa0, 0x33, 0x79, 0xfa, 0x13, 0x50, 0x9b, 0x22, 0x1f, 0x33, 0x8c, 0x48, 0x11, 0x7c, 0x3b,
	0x4d, 0x8c, 0xa6, 0xdc, 0x77, 0x5b, 0x82, 0xce, 0x1a, 0x83, 0x5f, 0x54, 0x50, 0xef, 0x33, 0xd6,
	0x27, 0xc1, 0x40, 0x86, 0xa1, 0xf5, 0x40, 0xd5, 0x63, 0x6c, 0xed, 0xc0, 0x7e, 0x9a, 0x18, 0xf5,
	0xdc, 0x01, 0xb1, 0x0e, 0x9d, 0x2d, 0x8f, 0xb1, 0x61, 0xf0, 0xaf, 0xf0, 0x4b, 0x0f, 0x0b, 0xdf,
	0x3e, 0x5d, 0xfc, 0xd6, 0x95, 0xef, 0x4b, 0x5d, 0x59, 0x2c, 0x75, 0xf5, 0x66, 0xa9, 0xab, 0xbf,
	0x96, 0xba, 0xfa, 0x75, 0xa5, 0x2b, 0x37, 0x2b, 0x5d, 0xf9, 0xb1, 0xd2, 0x95, 0x8f, 0xd6, 0x5f,
	0x96, 0x66, 0x7f, 0xc4, 0x11, 0xbd, 0xb8, 0xc0, 0x3e, 0xf6, 0x26, 0xf9, 0xdc, 0x2a, 0x5e, 0xb2,
	0xf0, 0x77, 0x5c, 0x15, 0x2f, 0xf2, 0xd5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x0b, 0x84,
	0x58, 0xe6, 0x03, 0x00, 0x00,
}

func (m *AppMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenesisToken) > 0 {
		for iNdEx := len(m.GenesisToken) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisToken[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAppMapping(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.GovTimeInSeconds != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.GovTimeInSeconds))))
		i--
		dAtA[i] = 0x29
	}
	{
		size := m.MinGovDeposit.Size()
		i -= size
		if _, err := m.MinGovDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAppMapping(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ShortName) > 0 {
		i -= len(m.ShortName)
		copy(dAtA[i:], m.ShortName)
		i = encodeVarintAppMapping(dAtA, i, uint64(len(m.ShortName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAppMapping(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAppMapping(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MintGenesisToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintGenesisToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintGenesisToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintAppMapping(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x22
	}
	if m.IsGovToken {
		i--
		if m.IsGovToken {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.GenesisSupply != nil {
		{
			size := m.GenesisSupply.Size()
			i -= size
			if _, err := m.GenesisSupply.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintAppMapping(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.AssetId != 0 {
		i = encodeVarintAppMapping(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AppAndGovTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppAndGovTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppAndGovTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GovTimeInSeconds != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.GovTimeInSeconds))))
		i--
		dAtA[i] = 0x11
	}
	if m.AppId != 0 {
		i = encodeVarintAppMapping(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAppMapping(dAtA []byte, offset int, v uint64) int {
	offset -= sovAppMapping(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AppMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAppMapping(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAppMapping(uint64(l))
	}
	l = len(m.ShortName)
	if l > 0 {
		n += 1 + l + sovAppMapping(uint64(l))
	}
	l = m.MinGovDeposit.Size()
	n += 1 + l + sovAppMapping(uint64(l))
	if m.GovTimeInSeconds != 0 {
		n += 9
	}
	if len(m.GenesisToken) > 0 {
		for _, e := range m.GenesisToken {
			l = e.Size()
			n += 1 + l + sovAppMapping(uint64(l))
		}
	}
	return n
}

func (m *MintGenesisToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetId != 0 {
		n += 1 + sovAppMapping(uint64(m.AssetId))
	}
	if m.GenesisSupply != nil {
		l = m.GenesisSupply.Size()
		n += 1 + l + sovAppMapping(uint64(l))
	}
	if m.IsGovToken {
		n += 2
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovAppMapping(uint64(l))
	}
	return n
}

func (m *AppAndGovTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppId != 0 {
		n += 1 + sovAppMapping(uint64(m.AppId))
	}
	if m.GovTimeInSeconds != 0 {
		n += 9
	}
	return n
}

func sovAppMapping(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAppMapping(x uint64) (n int) {
	return sovAppMapping(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppMapping
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
			return fmt.Errorf("proto: AppMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGovDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinGovDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovTimeInSeconds", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.GovTimeInSeconds = float64(math.Float64frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
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
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisToken = append(m.GenesisToken, MintGenesisToken{})
			if err := m.GenesisToken[len(m.GenesisToken)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAppMapping(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAppMapping
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MintGenesisToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppMapping
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
			return fmt.Errorf("proto: MintGenesisToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintGenesisToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.GenesisSupply = &v
			if err := m.GenesisSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsGovToken", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
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
			m.IsGovToken = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAppMapping(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAppMapping
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AppAndGovTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppMapping
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
			return fmt.Errorf("proto: AppAndGovTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppAndGovTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovTimeInSeconds", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.GovTimeInSeconds = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipAppMapping(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAppMapping
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAppMapping(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAppMapping
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
					return 0, ErrIntOverflowAppMapping
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAppMapping
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
				return 0, ErrInvalidLengthAppMapping
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAppMapping
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAppMapping
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAppMapping        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAppMapping          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAppMapping = fmt.Errorf("proto: unexpected end of group")
)
