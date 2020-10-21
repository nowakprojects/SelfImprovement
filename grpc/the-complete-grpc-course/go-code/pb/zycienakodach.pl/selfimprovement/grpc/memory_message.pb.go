// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: memory_message.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Memory_Unit int32

const (
	Memory_UNKNOWN  Memory_Unit = 0
	Memory_BIT      Memory_Unit = 1
	Memory_BYTE     Memory_Unit = 2
	Memory_KILOBYTE Memory_Unit = 3
	Memory_MEGABYTE Memory_Unit = 4
	Memory_GIGABYTE Memory_Unit = 5
	Memory_TERABYTE Memory_Unit = 6
)

// Enum value maps for Memory_Unit.
var (
	Memory_Unit_name = map[int32]string{
		0: "UNKNOWN",
		1: "BIT",
		2: "BYTE",
		3: "KILOBYTE",
		4: "MEGABYTE",
		5: "GIGABYTE",
		6: "TERABYTE",
	}
	Memory_Unit_value = map[string]int32{
		"UNKNOWN":  0,
		"BIT":      1,
		"BYTE":     2,
		"KILOBYTE": 3,
		"MEGABYTE": 4,
		"GIGABYTE": 5,
		"TERABYTE": 6,
	}
)

func (x Memory_Unit) Enum() *Memory_Unit {
	p := new(Memory_Unit)
	*p = x
	return p
}

func (x Memory_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Memory_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_memory_message_proto_enumTypes[0].Descriptor()
}

func (Memory_Unit) Type() protoreflect.EnumType {
	return &file_memory_message_proto_enumTypes[0]
}

func (x Memory_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Memory_Unit.Descriptor instead.
func (Memory_Unit) EnumDescriptor() ([]byte, []int) {
	return file_memory_message_proto_rawDescGZIP(), []int{0, 0}
}

type Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64      `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit  Memory_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=zycienakodach.pl.Memory_Unit" json:"unit,omitempty"`
}

func (x *Memory) Reset() {
	*x = Memory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memory) ProtoMessage() {}

func (x *Memory) ProtoReflect() protoreflect.Message {
	mi := &file_memory_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memory.ProtoReflect.Descriptor instead.
func (*Memory) Descriptor() ([]byte, []int) {
	return file_memory_message_proto_rawDescGZIP(), []int{0}
}

func (x *Memory) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Memory) GetUnit() Memory_Unit {
	if x != nil {
		return x.Unit
	}
	return Memory_UNKNOWN
}

var File_memory_message_proto protoreflect.FileDescriptor

var file_memory_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x7a, 0x79, 0x63, 0x69, 0x65, 0x6e, 0x61, 0x6b,
	0x6f, 0x64, 0x61, 0x63, 0x68, 0x2e, 0x70, 0x6c, 0x22, 0xb1, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x7a, 0x79, 0x63, 0x69, 0x65, 0x6e,
	0x61, 0x6b, 0x6f, 0x64, 0x61, 0x63, 0x68, 0x2e, 0x70, 0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x5e, 0x0a, 0x04,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x59,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x49, 0x4c, 0x4f, 0x42, 0x59, 0x54, 0x45,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x47, 0x41, 0x42, 0x59, 0x54, 0x45, 0x10, 0x04,
	0x12, 0x0c, 0x0a, 0x08, 0x47, 0x49, 0x47, 0x41, 0x42, 0x59, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0c,
	0x0a, 0x08, 0x54, 0x45, 0x52, 0x41, 0x42, 0x59, 0x54, 0x45, 0x10, 0x06, 0x42, 0x27, 0x5a, 0x25,
	0x7a, 0x79, 0x63, 0x69, 0x65, 0x6e, 0x61, 0x6b, 0x6f, 0x64, 0x61, 0x63, 0x68, 0x2e, 0x70, 0x6c,
	0x2f, 0x73, 0x65, 0x6c, 0x66, 0x69, 0x6d, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memory_message_proto_rawDescOnce sync.Once
	file_memory_message_proto_rawDescData = file_memory_message_proto_rawDesc
)

func file_memory_message_proto_rawDescGZIP() []byte {
	file_memory_message_proto_rawDescOnce.Do(func() {
		file_memory_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_memory_message_proto_rawDescData)
	})
	return file_memory_message_proto_rawDescData
}

var file_memory_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_memory_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_memory_message_proto_goTypes = []interface{}{
	(Memory_Unit)(0), // 0: zycienakodach.pl.Memory.Unit
	(*Memory)(nil),   // 1: zycienakodach.pl.Memory
}
var file_memory_message_proto_depIdxs = []int32{
	0, // 0: zycienakodach.pl.Memory.unit:type_name -> zycienakodach.pl.Memory.Unit
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_memory_message_proto_init() }
func file_memory_message_proto_init() {
	if File_memory_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memory_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memory); i {
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
			RawDescriptor: file_memory_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_memory_message_proto_goTypes,
		DependencyIndexes: file_memory_message_proto_depIdxs,
		EnumInfos:         file_memory_message_proto_enumTypes,
		MessageInfos:      file_memory_message_proto_msgTypes,
	}.Build()
	File_memory_message_proto = out.File
	file_memory_message_proto_rawDesc = nil
	file_memory_message_proto_goTypes = nil
	file_memory_message_proto_depIdxs = nil
}
