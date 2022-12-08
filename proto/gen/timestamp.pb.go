// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: timestamp.proto

package example

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

type ExampleTimestamp_ENUM int32

const (
	ExampleTimestamp_UNKONW ExampleTimestamp_ENUM = 0
	ExampleTimestamp_HTTP   ExampleTimestamp_ENUM = 1
	ExampleTimestamp_GRPC   ExampleTimestamp_ENUM = 2
)

// Enum value maps for ExampleTimestamp_ENUM.
var (
	ExampleTimestamp_ENUM_name = map[int32]string{
		0: "UNKONW",
		1: "HTTP",
		2: "GRPC",
	}
	ExampleTimestamp_ENUM_value = map[string]int32{
		"UNKONW": 0,
		"HTTP":   1,
		"GRPC":   2,
	}
)

func (x ExampleTimestamp_ENUM) Enum() *ExampleTimestamp_ENUM {
	p := new(ExampleTimestamp_ENUM)
	*p = x
	return p
}

func (x ExampleTimestamp_ENUM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExampleTimestamp_ENUM) Descriptor() protoreflect.EnumDescriptor {
	return file_timestamp_proto_enumTypes[0].Descriptor()
}

func (ExampleTimestamp_ENUM) Type() protoreflect.EnumType {
	return &file_timestamp_proto_enumTypes[0]
}

func (x ExampleTimestamp_ENUM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExampleTimestamp_ENUM.Descriptor instead.
func (ExampleTimestamp_ENUM) EnumDescriptor() ([]byte, []int) {
	return file_timestamp_proto_rawDescGZIP(), []int{0, 0}
}

type ExampleTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I64Example  int64                 `protobuf:"varint,1,opt,name=i64_example,json=i64Example,proto3" json:"i64_example,omitempty"`
	StrExample  string                `protobuf:"bytes,2,opt,name=str_example,json=strExample,proto3" json:"str_example,omitempty"`
	EnumExample ExampleTimestamp_ENUM `protobuf:"varint,3,opt,name=enum_example,json=enumExample,proto3,enum=example.ExampleTimestamp_ENUM" json:"enum_example,omitempty"`
}

func (x *ExampleTimestamp) Reset() {
	*x = ExampleTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_timestamp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleTimestamp) ProtoMessage() {}

func (x *ExampleTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_timestamp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleTimestamp.ProtoReflect.Descriptor instead.
func (*ExampleTimestamp) Descriptor() ([]byte, []int) {
	return file_timestamp_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleTimestamp) GetI64Example() int64 {
	if x != nil {
		return x.I64Example
	}
	return 0
}

func (x *ExampleTimestamp) GetStrExample() string {
	if x != nil {
		return x.StrExample
	}
	return ""
}

func (x *ExampleTimestamp) GetEnumExample() ExampleTimestamp_ENUM {
	if x != nil {
		return x.EnumExample
	}
	return ExampleTimestamp_UNKONW
}

var File_timestamp_proto protoreflect.FileDescriptor

var file_timestamp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x10, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x36, 0x34, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x36, 0x34, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x41, 0x0a, 0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x45, 0x4e, 0x55, 0x4d, 0x52, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x22, 0x26, 0x0a, 0x04, 0x45, 0x4e, 0x55, 0x4d, 0x12, 0x0a, 0x0a, 0x06,
	0x55, 0x4e, 0x4b, 0x4f, 0x4e, 0x57, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x10, 0x02, 0x42, 0x0f, 0x5a, 0x0d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_timestamp_proto_rawDescOnce sync.Once
	file_timestamp_proto_rawDescData = file_timestamp_proto_rawDesc
)

func file_timestamp_proto_rawDescGZIP() []byte {
	file_timestamp_proto_rawDescOnce.Do(func() {
		file_timestamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_timestamp_proto_rawDescData)
	})
	return file_timestamp_proto_rawDescData
}

var file_timestamp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_timestamp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_timestamp_proto_goTypes = []interface{}{
	(ExampleTimestamp_ENUM)(0), // 0: example.ExampleTimestamp.ENUM
	(*ExampleTimestamp)(nil),   // 1: example.ExampleTimestamp
}
var file_timestamp_proto_depIdxs = []int32{
	0, // 0: example.ExampleTimestamp.enum_example:type_name -> example.ExampleTimestamp.ENUM
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_timestamp_proto_init() }
func file_timestamp_proto_init() {
	if File_timestamp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_timestamp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleTimestamp); i {
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
			RawDescriptor: file_timestamp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_timestamp_proto_goTypes,
		DependencyIndexes: file_timestamp_proto_depIdxs,
		EnumInfos:         file_timestamp_proto_enumTypes,
		MessageInfos:      file_timestamp_proto_msgTypes,
	}.Build()
	File_timestamp_proto = out.File
	file_timestamp_proto_rawDesc = nil
	file_timestamp_proto_goTypes = nil
	file_timestamp_proto_depIdxs = nil
}