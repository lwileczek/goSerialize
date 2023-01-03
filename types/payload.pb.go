// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: payload.proto

package types

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

type SubStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cat     string `protobuf:"bytes,1,opt,name=Cat,proto3" json:"Cat,omitempty"`
	Feeling string `protobuf:"bytes,2,opt,name=Feeling,proto3" json:"Feeling,omitempty"`
}

func (x *SubStruct) Reset() {
	*x = SubStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubStruct) ProtoMessage() {}

func (x *SubStruct) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubStruct.ProtoReflect.Descriptor instead.
func (*SubStruct) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{0}
}

func (x *SubStruct) GetCat() string {
	if x != nil {
		return x.Cat
	}
	return ""
}

func (x *SubStruct) GetFeeling() string {
	if x != nil {
		return x.Feeling
	}
	return ""
}

type PbPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringEntry         string           `protobuf:"bytes,1,opt,name=StringEntry,proto3" json:"StringEntry,omitempty"`
	SmallInteger        uint32           `protobuf:"varint,2,opt,name=SmallInteger,proto3" json:"SmallInteger,omitempty"`
	NormalInteger       int64            `protobuf:"varint,3,opt,name=NormalInteger,proto3" json:"NormalInteger,omitempty"`
	Boolean             bool             `protobuf:"varint,4,opt,name=Boolean,proto3" json:"Boolean,omitempty"`
	SomeFloat           float32          `protobuf:"fixed32,5,opt,name=SomeFloat,proto3" json:"SomeFloat,omitempty"`
	IntArray            []int32          `protobuf:"varint,6,rep,packed,name=IntArray,proto3" json:"IntArray,omitempty"`
	Chart               map[string]int32 `protobuf:"bytes,7,rep,name=Chart,proto3" json:"Chart,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	SubShop             *SubStruct       `protobuf:"bytes,8,opt,name=SubShop,proto3" json:"SubShop,omitempty"`
	SerializationMethod string           `protobuf:"bytes,9,opt,name=SerializationMethod,proto3" json:"SerializationMethod,omitempty"`
}

func (x *PbPayload) Reset() {
	*x = PbPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbPayload) ProtoMessage() {}

func (x *PbPayload) ProtoReflect() protoreflect.Message {
	mi := &file_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbPayload.ProtoReflect.Descriptor instead.
func (*PbPayload) Descriptor() ([]byte, []int) {
	return file_payload_proto_rawDescGZIP(), []int{1}
}

func (x *PbPayload) GetStringEntry() string {
	if x != nil {
		return x.StringEntry
	}
	return ""
}

func (x *PbPayload) GetSmallInteger() uint32 {
	if x != nil {
		return x.SmallInteger
	}
	return 0
}

func (x *PbPayload) GetNormalInteger() int64 {
	if x != nil {
		return x.NormalInteger
	}
	return 0
}

func (x *PbPayload) GetBoolean() bool {
	if x != nil {
		return x.Boolean
	}
	return false
}

func (x *PbPayload) GetSomeFloat() float32 {
	if x != nil {
		return x.SomeFloat
	}
	return 0
}

func (x *PbPayload) GetIntArray() []int32 {
	if x != nil {
		return x.IntArray
	}
	return nil
}

func (x *PbPayload) GetChart() map[string]int32 {
	if x != nil {
		return x.Chart
	}
	return nil
}

func (x *PbPayload) GetSubShop() *SubStruct {
	if x != nil {
		return x.SubShop
	}
	return nil
}

func (x *PbPayload) GetSerializationMethod() string {
	if x != nil {
		return x.SerializationMethod
	}
	return ""
}

var File_payload_proto protoreflect.FileDescriptor

var file_payload_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x65, 0x78, 0x22, 0x37, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x43, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43,
	0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x65, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x90, 0x03, 0x0a,
	0x09, 0x50, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x0d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x53, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x49, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x08, 0x49, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x2e, 0x50,
	0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x53, 0x75,
	0x62, 0x53, 0x68, 0x6f, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x78,
	0x2e, 0x53, 0x75, 0x62, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x53, 0x75, 0x62, 0x53,
	0x68, 0x6f, 0x70, 0x12, 0x30, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x38, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_payload_proto_rawDescOnce sync.Once
	file_payload_proto_rawDescData = file_payload_proto_rawDesc
)

func file_payload_proto_rawDescGZIP() []byte {
	file_payload_proto_rawDescOnce.Do(func() {
		file_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_payload_proto_rawDescData)
	})
	return file_payload_proto_rawDescData
}

var file_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payload_proto_goTypes = []interface{}{
	(*SubStruct)(nil), // 0: ex.SubStruct
	(*PbPayload)(nil), // 1: ex.PbPayload
	nil,               // 2: ex.PbPayload.ChartEntry
}
var file_payload_proto_depIdxs = []int32{
	2, // 0: ex.PbPayload.Chart:type_name -> ex.PbPayload.ChartEntry
	0, // 1: ex.PbPayload.SubShop:type_name -> ex.SubStruct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payload_proto_init() }
func file_payload_proto_init() {
	if File_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubStruct); i {
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
		file_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbPayload); i {
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
			RawDescriptor: file_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payload_proto_goTypes,
		DependencyIndexes: file_payload_proto_depIdxs,
		MessageInfos:      file_payload_proto_msgTypes,
	}.Build()
	File_payload_proto = out.File
	file_payload_proto_rawDesc = nil
	file_payload_proto_goTypes = nil
	file_payload_proto_depIdxs = nil
}
