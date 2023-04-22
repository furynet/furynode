// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furynode/tokenregistry/v1/types.proto

package types

import (
	fmt "fmt"
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

type Permission int32

const (
	Permission_UNSPECIFIED  Permission = 0
	Permission_CLP          Permission = 1
	Permission_IBCEXPORT    Permission = 2
	Permission_IBCIMPORT    Permission = 3
	Permission_DISABLE_BUY  Permission = 4
	Permission_DISABLE_SELL Permission = 5
)

var Permission_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "CLP",
	2: "IBCEXPORT",
	3: "IBCIMPORT",
	4: "DISABLE_BUY",
	5: "DISABLE_SELL",
}

var Permission_value = map[string]int32{
	"UNSPECIFIED":  0,
	"CLP":          1,
	"IBCEXPORT":    2,
	"IBCIMPORT":    3,
	"DISABLE_BUY":  4,
	"DISABLE_SELL": 5,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d08afdaf425e66ea, []int{0}
}

type GenesisState struct {
	Registry *Registry `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08afdaf425e66ea, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetRegistry() *Registry {
	if m != nil {
		return m.Registry
	}
	return nil
}

type Registry struct {
	Entries []*RegistryEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08afdaf425e66ea, []int{1}
}
func (m *Registry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(m, src)
}
func (m *Registry) XXX_Size() int {
	return m.Size()
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetEntries() []*RegistryEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RegistryEntry struct {
	Decimals                 int64        `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Denom                    string       `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	BaseDenom                string       `protobuf:"bytes,4,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	Path                     string       `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	IbcChannelId             string       `protobuf:"bytes,6,opt,name=ibc_channel_id,json=ibcChannelId,proto3" json:"ibc_channel_id,omitempty"`
	IbcCounterpartyChannelId string       `protobuf:"bytes,7,opt,name=ibc_counterparty_channel_id,json=ibcCounterpartyChannelId,proto3" json:"ibc_counterparty_channel_id,omitempty"`
	DisplayName              string       `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	DisplaySymbol            string       `protobuf:"bytes,9,opt,name=display_symbol,json=displaySymbol,proto3" json:"display_symbol,omitempty"`
	Network                  string       `protobuf:"bytes,10,opt,name=network,proto3" json:"network,omitempty"`
	Address                  string       `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
	ExternalSymbol           string       `protobuf:"bytes,12,opt,name=external_symbol,json=externalSymbol,proto3" json:"external_symbol,omitempty"`
	TransferLimit            string       `protobuf:"bytes,13,opt,name=transfer_limit,json=transferLimit,proto3" json:"transfer_limit,omitempty"`
	Permissions              []Permission `protobuf:"varint,15,rep,packed,name=permissions,proto3,enum=furynode.tokenregistry.v1.Permission" json:"permissions,omitempty"`
	// The name of denomination unit of this token that is the smallest unit
	// stored. IBC imports of this RegistryEntry convert and store funds as
	// unit_denom. Several different denom units of a token may be imported into
	// this same unit denom, they should all be stored under the same unit_denom
	// if they are the same token. When exporting a RegistryEntry where unit_denom
	// != denom, then unit_denom can, in future, be used to indicate the source of
	// funds for a denom unit that does not actually exist on chain, enabling
	// other chains to overcome the uint64 limit on the packet level and import
	// large amounts of high precision tokens easily. ie. microfury -> fury i.e
	// fury -> fury
	UnitDenom string `protobuf:"bytes,16,opt,name=unit_denom,json=unitDenom,proto3" json:"unit_denom,omitempty"`
	// The name of denomination unit of this token that should appear on
	// counterparty chain when this unit is exported. If empty, the denom is
	// exported as is. Generally this will only be used to map a high precision
	// (unit_denom) to a lower precision, to overcome the current uint64 limit on
	// the packet level. i.e fury -> microfury i.e microfury -> microfury
	IbcCounterpartyDenom   string `protobuf:"bytes,17,opt,name=ibc_counterparty_denom,json=ibcCounterpartyDenom,proto3" json:"ibc_counterparty_denom,omitempty"`
	IbcCounterpartyChainId string `protobuf:"bytes,18,opt,name=ibc_counterparty_chain_id,json=ibcCounterpartyChainId,proto3" json:"ibc_counterparty_chain_id,omitempty"`
}

func (m *RegistryEntry) Reset()         { *m = RegistryEntry{} }
func (m *RegistryEntry) String() string { return proto.CompactTextString(m) }
func (*RegistryEntry) ProtoMessage()    {}
func (*RegistryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08afdaf425e66ea, []int{2}
}
func (m *RegistryEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegistryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegistryEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegistryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryEntry.Merge(m, src)
}
func (m *RegistryEntry) XXX_Size() int {
	return m.Size()
}
func (m *RegistryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryEntry proto.InternalMessageInfo

func (m *RegistryEntry) GetDecimals() int64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *RegistryEntry) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *RegistryEntry) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *RegistryEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RegistryEntry) GetIbcChannelId() string {
	if m != nil {
		return m.IbcChannelId
	}
	return ""
}

func (m *RegistryEntry) GetIbcCounterpartyChannelId() string {
	if m != nil {
		return m.IbcCounterpartyChannelId
	}
	return ""
}

func (m *RegistryEntry) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *RegistryEntry) GetDisplaySymbol() string {
	if m != nil {
		return m.DisplaySymbol
	}
	return ""
}

func (m *RegistryEntry) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *RegistryEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegistryEntry) GetExternalSymbol() string {
	if m != nil {
		return m.ExternalSymbol
	}
	return ""
}

func (m *RegistryEntry) GetTransferLimit() string {
	if m != nil {
		return m.TransferLimit
	}
	return ""
}

func (m *RegistryEntry) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *RegistryEntry) GetUnitDenom() string {
	if m != nil {
		return m.UnitDenom
	}
	return ""
}

func (m *RegistryEntry) GetIbcCounterpartyDenom() string {
	if m != nil {
		return m.IbcCounterpartyDenom
	}
	return ""
}

func (m *RegistryEntry) GetIbcCounterpartyChainId() string {
	if m != nil {
		return m.IbcCounterpartyChainId
	}
	return ""
}

func init() {
	proto.RegisterEnum("furynode.tokenregistry.v1.Permission", Permission_name, Permission_value)
	proto.RegisterType((*GenesisState)(nil), "furynode.tokenregistry.v1.GenesisState")
	proto.RegisterType((*Registry)(nil), "furynode.tokenregistry.v1.Registry")
	proto.RegisterType((*RegistryEntry)(nil), "furynode.tokenregistry.v1.RegistryEntry")
}

func init() {
	proto.RegisterFile("furynode/tokenregistry/v1/types.proto", fileDescriptor_d08afdaf425e66ea)
}

var fileDescriptor_d08afdaf425e66ea = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4f, 0x4f, 0x13, 0x41,
	0x18, 0xc6, 0xbb, 0xb4, 0xa5, 0xed, 0xdb, 0x3f, 0xac, 0x13, 0x42, 0x46, 0x8c, 0x4d, 0x6d, 0x20,
	0x34, 0x1e, 0xda, 0x80, 0x5e, 0x3c, 0x68, 0x42, 0x4b, 0x31, 0xab, 0x05, 0x9b, 0x56, 0x12, 0xf5,
	0xd2, 0xcc, 0x76, 0x87, 0x32, 0x61, 0x77, 0x66, 0x33, 0x33, 0x20, 0xfd, 0x00, 0xde, 0xfd, 0x58,
	0x1e, 0x39, 0x7a, 0x34, 0xf0, 0x45, 0xcc, 0xce, 0xee, 0x42, 0x01, 0x89, 0xb7, 0x79, 0x9f, 0xe7,
	0xf7, 0xbc, 0x99, 0xec, 0x93, 0x1d, 0xd8, 0x50, 0xec, 0x98, 0x0b, 0x8f, 0x76, 0xb4, 0x38, 0xa5,
	0x5c, 0xd2, 0x19, 0x53, 0x5a, 0xce, 0x3b, 0xe7, 0xdb, 0x1d, 0x3d, 0x0f, 0xa9, 0x6a, 0x87, 0x52,
	0x68, 0x81, 0x70, 0x42, 0xb5, 0xef, 0x50, 0xed, 0xf3, 0xed, 0xf5, 0xd5, 0x99, 0x98, 0x09, 0x03,
	0x75, 0xa2, 0x53, 0xcc, 0x37, 0x0f, 0xa1, 0xf2, 0x9e, 0x72, 0xaa, 0x98, 0x1a, 0x6b, 0xa2, 0x29,
	0x7a, 0x07, 0xc5, 0x34, 0x84, 0x97, 0x1a, 0x56, 0xab, 0xbc, 0xd3, 0x6c, 0x3f, 0xb6, 0xb2, 0x3d,
	0x4a, 0xce, 0xa3, 0x9b, 0x4c, 0xf3, 0x00, 0x8a, 0xa9, 0x8a, 0x76, 0xa1, 0x40, 0xb9, 0x96, 0x8c,
	0x2a, 0x6c, 0x35, 0xb2, 0xad, 0xf2, 0xce, 0xd6, 0xff, 0x57, 0xf5, 0x79, 0xb4, 0x2f, 0xcd, 0x35,
	0x7f, 0xe4, 0xa1, 0x7a, 0xc7, 0x42, 0xeb, 0x50, 0xf4, 0xe8, 0x94, 0x05, 0xc4, 0x57, 0xe6, 0x82,
	0xd9, 0xd1, 0xcd, 0x8c, 0x56, 0x21, 0xef, 0x51, 0x2e, 0x02, 0x9c, 0x6d, 0x58, 0xad, 0xd2, 0x28,
	0x1e, 0xd0, 0x73, 0x00, 0x97, 0x28, 0x3a, 0x89, 0xad, 0x9c, 0xb1, 0x4a, 0x91, 0xb2, 0x67, 0x6c,
	0x04, 0xb9, 0x90, 0xe8, 0x13, 0x9c, 0x37, 0x86, 0x39, 0xa3, 0x0d, 0xa8, 0x31, 0x77, 0x3a, 0x99,
	0x9e, 0x10, 0xce, 0xa9, 0x3f, 0x61, 0x1e, 0x5e, 0x36, 0x6e, 0x85, 0xb9, 0xd3, 0x5e, 0x2c, 0x3a,
	0x1e, 0x7a, 0x0b, 0xcf, 0x0c, 0x25, 0xce, 0xb8, 0xa6, 0x32, 0x24, 0x52, 0xcf, 0x17, 0x23, 0x05,
	0x13, 0xc1, 0x51, 0x64, 0x81, 0xb8, 0x8d, 0xbf, 0x80, 0x8a, 0xc7, 0x54, 0xe8, 0x93, 0xf9, 0x84,
	0x93, 0x80, 0xe2, 0xa2, 0xe1, 0xcb, 0x89, 0x76, 0x48, 0x02, 0x8a, 0x36, 0xa1, 0x96, 0x22, 0x6a,
	0x1e, 0xb8, 0xc2, 0xc7, 0x25, 0x03, 0x55, 0x13, 0x75, 0x6c, 0x44, 0x84, 0xa1, 0xc0, 0xa9, 0xfe,
	0x2e, 0xe4, 0x29, 0x06, 0xe3, 0xa7, 0x63, 0xe4, 0x10, 0xcf, 0x93, 0x54, 0x29, 0x5c, 0x8e, 0x9d,
	0x64, 0x44, 0x5b, 0xb0, 0x42, 0x2f, 0x34, 0x95, 0x9c, 0xf8, 0xe9, 0xee, 0x8a, 0x21, 0x6a, 0xa9,
	0x9c, 0x2c, 0xdf, 0x84, 0x9a, 0x96, 0x84, 0xab, 0x63, 0x2a, 0x27, 0x3e, 0x0b, 0x98, 0xc6, 0xd5,
	0xf8, 0x0e, 0xa9, 0x3a, 0x88, 0x44, 0xb4, 0x0f, 0xe5, 0x90, 0xca, 0x80, 0x29, 0xc5, 0x04, 0x57,
	0x78, 0xa5, 0x91, 0x6d, 0xd5, 0x76, 0x36, 0x1e, 0x2f, 0x7c, 0x78, 0x03, 0x8f, 0x16, 0x83, 0x51,
	0x5b, 0x67, 0x9c, 0xe9, 0xa4, 0x2d, 0x3b, 0x6e, 0x2b, 0x52, 0xe2, 0xb6, 0x5e, 0xc3, 0xda, 0x83,
	0x6f, 0x1e, 0xa3, 0x4f, 0x0c, 0xba, 0x7a, 0xef, 0x73, 0xc7, 0xa9, 0x37, 0xf0, 0xf4, 0x5f, 0x4d,
	0x31, 0x1e, 0xf5, 0x84, 0x4c, 0x70, 0xed, 0x61, 0x4f, 0x8c, 0x3b, 0xde, 0x87, 0x5c, 0xd1, 0xb2,
	0x97, 0x5e, 0xce, 0x00, 0x6e, 0x2f, 0x8c, 0x56, 0xa0, 0x7c, 0x74, 0x38, 0x1e, 0xf6, 0x7b, 0xce,
	0xbe, 0xd3, 0xdf, 0xb3, 0x33, 0xa8, 0x00, 0xd9, 0xde, 0x60, 0x68, 0x5b, 0xa8, 0x0a, 0x25, 0xa7,
	0xdb, 0xeb, 0x7f, 0x19, 0x7e, 0x1a, 0x7d, 0xb6, 0x97, 0x92, 0xd1, 0x39, 0x30, 0x63, 0x36, 0xca,
	0xed, 0x39, 0xe3, 0xdd, 0xee, 0xa0, 0x3f, 0xe9, 0x1e, 0x7d, 0xb5, 0x73, 0xc8, 0x86, 0x4a, 0x2a,
	0x8c, 0xfb, 0x83, 0x81, 0x9d, 0xef, 0x7e, 0xfc, 0x75, 0x55, 0xb7, 0x2e, 0xaf, 0xea, 0xd6, 0x9f,
	0xab, 0xba, 0xf5, 0xf3, 0xba, 0x9e, 0xb9, 0xbc, 0xae, 0x67, 0x7e, 0x5f, 0xd7, 0x33, 0xdf, 0xb6,
	0x67, 0x4c, 0x9f, 0x9c, 0xb9, 0xed, 0xa9, 0x08, 0x3a, 0x63, 0x76, 0x6c, 0xae, 0xdf, 0x49, 0xdf,
	0x84, 0x8b, 0x7b, 0xaf, 0x82, 0x79, 0x12, 0xdc, 0x65, 0xf3, 0x8f, 0xbf, 0xfa, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x80, 0xb9, 0xb6, 0x3b, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Registry != nil {
		{
			size, err := m.Registry.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *Registry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Registry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Registry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RegistryEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegistryEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegistryEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcCounterpartyChainId) > 0 {
		i -= len(m.IbcCounterpartyChainId)
		copy(dAtA[i:], m.IbcCounterpartyChainId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.IbcCounterpartyChainId)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if len(m.IbcCounterpartyDenom) > 0 {
		i -= len(m.IbcCounterpartyDenom)
		copy(dAtA[i:], m.IbcCounterpartyDenom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.IbcCounterpartyDenom)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.UnitDenom) > 0 {
		i -= len(m.UnitDenom)
		copy(dAtA[i:], m.UnitDenom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.UnitDenom)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.Permissions) > 0 {
		dAtA3 := make([]byte, len(m.Permissions)*10)
		var j2 int
		for _, num := range m.Permissions {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintTypes(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.TransferLimit) > 0 {
		i -= len(m.TransferLimit)
		copy(dAtA[i:], m.TransferLimit)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TransferLimit)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.ExternalSymbol) > 0 {
		i -= len(m.ExternalSymbol)
		copy(dAtA[i:], m.ExternalSymbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ExternalSymbol)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.DisplaySymbol) > 0 {
		i -= len(m.DisplaySymbol)
		copy(dAtA[i:], m.DisplaySymbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DisplaySymbol)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.DisplayName) > 0 {
		i -= len(m.DisplayName)
		copy(dAtA[i:], m.DisplayName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DisplayName)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.IbcCounterpartyChannelId) > 0 {
		i -= len(m.IbcCounterpartyChannelId)
		copy(dAtA[i:], m.IbcCounterpartyChannelId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.IbcCounterpartyChannelId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.IbcChannelId) > 0 {
		i -= len(m.IbcChannelId)
		copy(dAtA[i:], m.IbcChannelId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.IbcChannelId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Decimals != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Registry != nil {
		l = m.Registry.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Registry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *RegistryEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Decimals != 0 {
		n += 1 + sovTypes(uint64(m.Decimals))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.IbcChannelId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.IbcCounterpartyChannelId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DisplaySymbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ExternalSymbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.TransferLimit)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Permissions) > 0 {
		l = 0
		for _, e := range m.Permissions {
			l += sovTypes(uint64(e))
		}
		n += 1 + sovTypes(uint64(l)) + l
	}
	l = len(m.UnitDenom)
	if l > 0 {
		n += 2 + l + sovTypes(uint64(l))
	}
	l = len(m.IbcCounterpartyDenom)
	if l > 0 {
		n += 2 + l + sovTypes(uint64(l))
	}
	l = len(m.IbcCounterpartyChainId)
	if l > 0 {
		n += 2 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Registry == nil {
				m.Registry = &Registry{}
			}
			if err := m.Registry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Registry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Registry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Registry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &RegistryEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RegistryEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RegistryEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegistryEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcCounterpartyChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcCounterpartyChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplaySymbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplaySymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalSymbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferLimit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType == 0 {
				var v Permission
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Permission(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Permissions = append(m.Permissions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypes
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Permissions) == 0 {
					m.Permissions = make([]Permission, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Permission
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Permission(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Permissions = append(m.Permissions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnitDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnitDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcCounterpartyDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcCounterpartyDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcCounterpartyChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcCounterpartyChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
