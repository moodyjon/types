// Code generated by protoc-gen-go. DO NOT EDIT.
// source: legacy_claim.proto

package legacy_pb

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

type Claim_Version int32

const (
	Claim_UNKNOWN_VERSION Claim_Version = 0
	Claim__0_0_1          Claim_Version = 1
)

var Claim_Version_name = map[int32]string{
	0: "UNKNOWN_VERSION",
	1: "_0_0_1",
}

var Claim_Version_value = map[string]int32{
	"UNKNOWN_VERSION": 0,
	"_0_0_1":          1,
}

func (x Claim_Version) Enum() *Claim_Version {
	p := new(Claim_Version)
	*p = x
	return p
}

func (x Claim_Version) String() string {
	return proto.EnumName(Claim_Version_name, int32(x))
}

func (x *Claim_Version) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Claim_Version_value, data, "Claim_Version")
	if err != nil {
		return err
	}
	*x = Claim_Version(value)
	return nil
}

func (Claim_Version) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b52e2f635402822, []int{0, 0}
}

type Claim_ClaimType int32

const (
	Claim_UNKNOWN_CLAIM_TYPE Claim_ClaimType = 0
	Claim_streamType         Claim_ClaimType = 1
	Claim_certificateType    Claim_ClaimType = 2
)

var Claim_ClaimType_name = map[int32]string{
	0: "UNKNOWN_CLAIM_TYPE",
	1: "streamType",
	2: "certificateType",
}

var Claim_ClaimType_value = map[string]int32{
	"UNKNOWN_CLAIM_TYPE": 0,
	"streamType":         1,
	"certificateType":    2,
}

func (x Claim_ClaimType) Enum() *Claim_ClaimType {
	p := new(Claim_ClaimType)
	*p = x
	return p
}

func (x Claim_ClaimType) String() string {
	return proto.EnumName(Claim_ClaimType_name, int32(x))
}

func (x *Claim_ClaimType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Claim_ClaimType_value, data, "Claim_ClaimType")
	if err != nil {
		return err
	}
	*x = Claim_ClaimType(value)
	return nil
}

func (Claim_ClaimType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b52e2f635402822, []int{0, 1}
}

type Claim struct {
	Version              *Claim_Version   `protobuf:"varint,1,req,name=version,enum=legacy_pb.Claim_Version" json:"version,omitempty"`
	ClaimType            *Claim_ClaimType `protobuf:"varint,2,req,name=claimType,enum=legacy_pb.Claim_ClaimType" json:"claimType,omitempty"`
	Stream               *Stream          `protobuf:"bytes,3,opt,name=stream" json:"stream,omitempty"`
	Certificate          *Certificate     `protobuf:"bytes,4,opt,name=certificate" json:"certificate,omitempty"`
	PublisherSignature   *Signature       `protobuf:"bytes,5,opt,name=publisherSignature" json:"publisherSignature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b52e2f635402822, []int{0}
}

func (m *Claim) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Claim.Unmarshal(m, b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return xxx_messageInfo_Claim.Size(m)
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetVersion() Claim_Version {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Claim_UNKNOWN_VERSION
}

func (m *Claim) GetClaimType() Claim_ClaimType {
	if m != nil && m.ClaimType != nil {
		return *m.ClaimType
	}
	return Claim_UNKNOWN_CLAIM_TYPE
}

func (m *Claim) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

func (m *Claim) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Claim) GetPublisherSignature() *Signature {
	if m != nil {
		return m.PublisherSignature
	}
	return nil
}

func init() {
	proto.RegisterEnum("legacy_pb.Claim_Version", Claim_Version_name, Claim_Version_value)
	proto.RegisterEnum("legacy_pb.Claim_ClaimType", Claim_ClaimType_name, Claim_ClaimType_value)
	proto.RegisterType((*Claim)(nil), "legacy_pb.Claim")
}

func init() { proto.RegisterFile("legacy_claim.proto", fileDescriptor_3b52e2f635402822) }

var fileDescriptor_3b52e2f635402822 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4f, 0xfa, 0x40,
	0x14, 0xc4, 0x29, 0xfc, 0x81, 0xf0, 0xf8, 0x07, 0xca, 0xd3, 0x90, 0x0d, 0x27, 0xc2, 0x09, 0x3d,
	0x34, 0xd8, 0x13, 0x57, 0x53, 0x49, 0x24, 0x6a, 0x31, 0x5b, 0xc4, 0x78, 0xda, 0x94, 0x66, 0xc5,
	0x4d, 0x2a, 0x6d, 0xb6, 0xc5, 0x84, 0x2f, 0xe2, 0xe7, 0x35, 0xdd, 0x6e, 0x61, 0x13, 0xbd, 0xce,
	0xfc, 0xde, 0xbc, 0x19, 0xc0, 0x98, 0xef, 0xc2, 0xe8, 0xc8, 0xa2, 0x38, 0x14, 0x9f, 0x4e, 0x2a,
	0x93, 0x3c, 0xc1, 0x8e, 0xd6, 0xd2, 0xed, 0xe8, 0x7f, 0x96, 0x4b, 0x1e, 0x6a, 0x63, 0x34, 0x88,
	0xb8, 0xcc, 0xc5, 0xbb, 0x88, 0xc2, 0x9c, 0x6b, 0xa9, 0x9f, 0x89, 0xdd, 0x3e, 0xcc, 0x0f, 0x52,
	0x0b, 0x93, 0xef, 0x06, 0x34, 0xbd, 0x22, 0x0c, 0x5d, 0x68, 0x7f, 0x71, 0x99, 0x89, 0x64, 0x4f,
	0xac, 0x71, 0x7d, 0xda, 0x73, 0x89, 0x73, 0x0a, 0x76, 0x14, 0xe2, 0x6c, 0x4a, 0x9f, 0x56, 0x20,
	0xce, 0xa1, 0xa3, 0x9a, 0xac, 0x8f, 0x29, 0x27, 0x75, 0x75, 0x35, 0xfa, 0x75, 0xe5, 0x55, 0x04,
	0x3d, 0xc3, 0x78, 0x05, 0xad, 0xb2, 0x2b, 0x69, 0x8c, 0xad, 0x69, 0xd7, 0x1d, 0x18, 0x67, 0x81,
	0x32, 0xa8, 0x06, 0x70, 0x0e, 0x5d, 0x63, 0x08, 0xf9, 0xa7, 0xf8, 0xa1, 0xf9, 0xe6, 0xec, 0x52,
	0x13, 0xc5, 0x3b, 0xc0, 0xf4, 0xb0, 0x8d, 0x45, 0xf6, 0xc1, 0x65, 0x50, 0x0d, 0x27, 0x4d, 0x15,
	0x70, 0x69, 0x3e, 0xac, 0x3c, 0xfa, 0x07, 0x3f, 0xb9, 0x86, 0xb6, 0x1e, 0x8e, 0x17, 0xd0, 0x7f,
	0xf1, 0x1f, 0xfc, 0xd5, 0xab, 0xcf, 0x36, 0x0b, 0x1a, 0x2c, 0x57, 0xbe, 0x5d, 0x43, 0x80, 0x16,
	0x9b, 0xb1, 0x19, 0xbb, 0xb1, 0xad, 0xc9, 0x3d, 0x74, 0x4e, 0x73, 0x71, 0x08, 0x58, 0xd1, 0xde,
	0xe3, 0xed, 0xf2, 0x89, 0xad, 0xdf, 0x9e, 0x17, 0x76, 0x0d, 0x7b, 0x00, 0xe5, 0xb4, 0x82, 0xb2,
	0xad, 0x22, 0xd5, 0x68, 0xad, 0xc4, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xb6, 0x8d,
	0x7b, 0xea, 0x01, 0x00, 0x00,
}
