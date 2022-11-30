// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: source.proto

package _go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Source_Version int32

const (
	Source_UNKNOWN_VERSION Source_Version = 0
	Source__0_0_1          Source_Version = 1
)

// Enum value maps for Source_Version.
var (
	Source_Version_name = map[int32]string{
		0: "UNKNOWN_VERSION",
		1: "_0_0_1",
	}
	Source_Version_value = map[string]int32{
		"UNKNOWN_VERSION": 0,
		"_0_0_1":          1,
	}
)

func (x Source_Version) Enum() *Source_Version {
	p := new(Source_Version)
	*p = x
	return p
}

func (x Source_Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source_Version) Descriptor() protoreflect.EnumDescriptor {
	return file_source_proto_enumTypes[0].Descriptor()
}

func (Source_Version) Type() protoreflect.EnumType {
	return &file_source_proto_enumTypes[0]
}

func (x Source_Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Source_Version) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Source_Version(num)
	return nil
}

// Deprecated: Use Source_Version.Descriptor instead.
func (Source_Version) EnumDescriptor() ([]byte, []int) {
	return file_source_proto_rawDescGZIP(), []int{0, 0}
}

type Source_SourceTypes int32

const (
	Source_UNKNOWN_SOURCE_TYPE Source_SourceTypes = 0
	Source_lbry_sd_hash        Source_SourceTypes = 1
)

// Enum value maps for Source_SourceTypes.
var (
	Source_SourceTypes_name = map[int32]string{
		0: "UNKNOWN_SOURCE_TYPE",
		1: "lbry_sd_hash",
	}
	Source_SourceTypes_value = map[string]int32{
		"UNKNOWN_SOURCE_TYPE": 0,
		"lbry_sd_hash":        1,
	}
)

func (x Source_SourceTypes) Enum() *Source_SourceTypes {
	p := new(Source_SourceTypes)
	*p = x
	return p
}

func (x Source_SourceTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source_SourceTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_source_proto_enumTypes[1].Descriptor()
}

func (Source_SourceTypes) Type() protoreflect.EnumType {
	return &file_source_proto_enumTypes[1]
}

func (x Source_SourceTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Source_SourceTypes) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Source_SourceTypes(num)
	return nil
}

// Deprecated: Use Source_SourceTypes.Descriptor instead.
func (Source_SourceTypes) EnumDescriptor() ([]byte, []int) {
	return file_source_proto_rawDescGZIP(), []int{0, 1}
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     *Source_Version     `protobuf:"varint,1,req,name=version,enum=legacy_pb.Source_Version" json:"version"`
	SourceType  *Source_SourceTypes `protobuf:"varint,2,req,name=sourceType,enum=legacy_pb.Source_SourceTypes" json:"sourceType"`
	Source      []byte              `protobuf:"bytes,3,req,name=source" json:"source"`
	ContentType *string             `protobuf:"bytes,4,req,name=contentType" json:"contentType"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_source_proto_rawDescGZIP(), []int{0}
}

func (x *Source) GetVersion() Source_Version {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return Source_UNKNOWN_VERSION
}

func (x *Source) GetSourceType() Source_SourceTypes {
	if x != nil && x.SourceType != nil {
		return *x.SourceType
	}
	return Source_UNKNOWN_SOURCE_TYPE
}

func (x *Source) GetSource() []byte {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Source) GetContentType() string {
	if x != nil && x.ContentType != nil {
		return *x.ContentType
	}
	return ""
}

var File_source_proto protoreflect.FileDescriptor

var file_source_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x62, 0x22, 0x9c, 0x02, 0x0a, 0x06, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x70,
	0x62, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x2a, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x5f, 0x30, 0x5f, 0x30, 0x5f, 0x31, 0x10, 0x01, 0x22, 0x38,
	0x0a, 0x0b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x17, 0x0a,
	0x13, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x6c, 0x62, 0x72, 0x79, 0x5f, 0x73,
	0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x10, 0x01,
}

var (
	file_source_proto_rawDescOnce sync.Once
	file_source_proto_rawDescData = file_source_proto_rawDesc
)

func file_source_proto_rawDescGZIP() []byte {
	file_source_proto_rawDescOnce.Do(func() {
		file_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_proto_rawDescData)
	})
	return file_source_proto_rawDescData
}

var file_source_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_source_proto_goTypes = []interface{}{
	(Source_Version)(0),     // 0: legacy_pb.Source.Version
	(Source_SourceTypes)(0), // 1: legacy_pb.Source.SourceTypes
	(*Source)(nil),          // 2: legacy_pb.Source
}
var file_source_proto_depIdxs = []int32{
	0, // 0: legacy_pb.Source.version:type_name -> legacy_pb.Source.Version
	1, // 1: legacy_pb.Source.sourceType:type_name -> legacy_pb.Source.SourceTypes
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_source_proto_init() }
func file_source_proto_init() {
	if File_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_source_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_source_proto_goTypes,
		DependencyIndexes: file_source_proto_depIdxs,
		EnumInfos:         file_source_proto_enumTypes,
		MessageInfos:      file_source_proto_msgTypes,
	}.Build()
	File_source_proto = out.File
	file_source_proto_rawDesc = nil
	file_source_proto_goTypes = nil
	file_source_proto_depIdxs = nil
}
