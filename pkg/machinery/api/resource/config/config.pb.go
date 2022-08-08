// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: resource/config/config.proto

package config

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MachineType matches machine.Type constants.
type MachineType int32

const (
	MachineType_UNKNOWN       MachineType = 0
	MachineType_INIT          MachineType = 1
	MachineType_CONTROL_PLANE MachineType = 2
	MachineType_WORKER        MachineType = 3
)

// Enum value maps for MachineType.
var (
	MachineType_name = map[int32]string{
		0: "UNKNOWN",
		1: "INIT",
		2: "CONTROL_PLANE",
		3: "WORKER",
	}
	MachineType_value = map[string]int32{
		"UNKNOWN":       0,
		"INIT":          1,
		"CONTROL_PLANE": 2,
		"WORKER":        3,
	}
)

func (x MachineType) Enum() *MachineType {
	p := new(MachineType)
	*p = x
	return p
}

func (x MachineType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineType) Descriptor() protoreflect.EnumDescriptor {
	return file_resource_config_config_proto_enumTypes[0].Descriptor()
}

func (MachineType) Type() protoreflect.EnumType {
	return &file_resource_config_config_proto_enumTypes[0]
}

func (x MachineType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineType.Descriptor instead.
func (MachineType) EnumDescriptor() ([]byte, []int) {
	return file_resource_config_config_proto_rawDescGZIP(), []int{0}
}

// MessageConfigSpec is the spec for the config.MachineConfig resource.
type MachineConfigSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains YAML marshalled machine configuration.
	//
	// Byte representation is preserved as the machine configuration was submitted to the node.
	YamlMarshalled []byte `protobuf:"bytes,1,opt,name=yaml_marshalled,json=yamlMarshalled,proto3" json:"yaml_marshalled,omitempty"`
}

func (x *MachineConfigSpec) Reset() {
	*x = MachineConfigSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineConfigSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineConfigSpec) ProtoMessage() {}

func (x *MachineConfigSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineConfigSpec.ProtoReflect.Descriptor instead.
func (*MachineConfigSpec) Descriptor() ([]byte, []int) {
	return file_resource_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *MachineConfigSpec) GetYamlMarshalled() []byte {
	if x != nil {
		return x.YamlMarshalled
	}
	return nil
}

// MachineTypeSpec is the spec for the config.MachineType resource.
type MachineTypeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineType MachineType `protobuf:"varint,1,opt,name=machine_type,json=machineType,proto3,enum=resource.config.MachineType" json:"machine_type,omitempty"`
}

func (x *MachineTypeSpec) Reset() {
	*x = MachineTypeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineTypeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineTypeSpec) ProtoMessage() {}

func (x *MachineTypeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineTypeSpec.ProtoReflect.Descriptor instead.
func (*MachineTypeSpec) Descriptor() ([]byte, []int) {
	return file_resource_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *MachineTypeSpec) GetMachineType() MachineType {
	if x != nil {
		return x.MachineType
	}
	return MachineType_UNKNOWN
}

var File_resource_config_config_proto protoreflect.FileDescriptor

var file_resource_config_config_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x3c, 0x0a, 0x11, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x27, 0x0a, 0x0f, 0x79, 0x61, 0x6d, 0x6c, 0x5f, 0x6d, 0x61, 0x72,
	0x73, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x79,
	0x61, 0x6d, 0x6c, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x22, 0x52, 0x0a,
	0x0f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x3f, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x2a, 0x43, 0x0a, 0x0b, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x52,
	0x4f, 0x4c, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x4f,
	0x52, 0x4b, 0x45, 0x52, 0x10, 0x03, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_resource_config_config_proto_rawDescOnce sync.Once
	file_resource_config_config_proto_rawDescData = file_resource_config_config_proto_rawDesc
)

func file_resource_config_config_proto_rawDescGZIP() []byte {
	file_resource_config_config_proto_rawDescOnce.Do(func() {
		file_resource_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_config_config_proto_rawDescData)
	})
	return file_resource_config_config_proto_rawDescData
}

var file_resource_config_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resource_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resource_config_config_proto_goTypes = []interface{}{
	(MachineType)(0),          // 0: resource.config.MachineType
	(*MachineConfigSpec)(nil), // 1: resource.config.MachineConfigSpec
	(*MachineTypeSpec)(nil),   // 2: resource.config.MachineTypeSpec
}
var file_resource_config_config_proto_depIdxs = []int32{
	0, // 0: resource.config.MachineTypeSpec.machine_type:type_name -> resource.config.MachineType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resource_config_config_proto_init() }
func file_resource_config_config_proto_init() {
	if File_resource_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineConfigSpec); i {
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
		file_resource_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineTypeSpec); i {
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
			RawDescriptor: file_resource_config_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_config_config_proto_goTypes,
		DependencyIndexes: file_resource_config_config_proto_depIdxs,
		EnumInfos:         file_resource_config_config_proto_enumTypes,
		MessageInfos:      file_resource_config_config_proto_msgTypes,
	}.Build()
	File_resource_config_config_proto = out.File
	file_resource_config_config_proto_rawDesc = nil
	file_resource_config_config_proto_goTypes = nil
	file_resource_config_config_proto_depIdxs = nil
}
