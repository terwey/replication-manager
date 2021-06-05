// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: config.proto

package config_store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Environment int32

const (
	Environment_NONE        Environment = 0
	Environment_PRODUCTION  Environment = 1
	Environment_STAGING     Environment = 2
	Environment_TESTING     Environment = 3
	Environment_DEVELOPMENT Environment = 4
)

// Enum value maps for Environment.
var (
	Environment_name = map[int32]string{
		0: "NONE",
		1: "PRODUCTION",
		2: "STAGING",
		3: "TESTING",
		4: "DEVELOPMENT",
	}
	Environment_value = map[string]int32{
		"NONE":        0,
		"PRODUCTION":  1,
		"STAGING":     2,
		"TESTING":     3,
		"DEVELOPMENT": 4,
	}
)

func (x Environment) Enum() *Environment {
	p := new(Environment)
	*p = x
	return p
}

func (x Environment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Environment) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[0].Descriptor()
}

func (Environment) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[0]
}

func (x Environment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Environment.Descriptor instead.
func (Environment) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_STRING Type = 0
	Type_INT    Type = 1
	Type_FLOAT  Type = 2
	Type_BOOL   Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "STRING",
		1: "INT",
		2: "FLOAT",
		3: "BOOL",
	}
	Type_value = map[string]int32{
		"STRING": 0,
		"INT":    1,
		"FLOAT":  2,
		"BOOL":   3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string      `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values      []*Value    `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	Namespace   string      `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Environment Environment `protobuf:"varint,4,opt,name=environment,proto3,enum=signal18.config_store.v1.Environment" json:"environment,omitempty"`
	Section     []string    `protobuf:"bytes,5,rep,name=section,proto3" json:"section,omitempty"`
	Secret      bool        `protobuf:"varint,6,opt,name=secret,proto3" json:"secret,omitempty"`
	Revision    int64       `protobuf:"varint,10,opt,name=revision,proto3" json:"revision,omitempty"`
	Version     string      `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`
	// either created or modified can be dropped as they are the same value
	Created *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created,proto3" json:"created,omitempty"`
	Deleted *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Property) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Property) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *Property) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Property) GetEnvironment() Environment {
	if x != nil {
		return x.Environment
	}
	return Environment_NONE
}

func (x *Property) GetSection() []string {
	if x != nil {
		return x.Section
	}
	return nil
}

func (x *Property) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

func (x *Property) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *Property) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Property) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Property) GetDeleted() *timestamppb.Timestamp {
	if x != nil {
		return x.Deleted
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type     Type   `protobuf:"varint,2,opt,name=type,proto3,enum=signal18.config_store.v1.Type" json:"type,omitempty"`
	Checksum string `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"` // used to prevent restoring encrypted values
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Value) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_STRING
}

func (x *Value) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Property    *Property `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Limit       int64     `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	IgnoreValue bool      `protobuf:"varint,3,opt,name=ignore_value,json=ignoreValue,proto3" json:"ignore_value,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *Query) GetProperty() *Property {
	if x != nil {
		return x.Property
	}
	return nil
}

func (x *Query) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Query) GetIgnoreValue() bool {
	if x != nil {
		return x.IgnoreValue
	}
	return false
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x03, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x47, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x6b, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x22, 0x80, 0x01, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x3e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x52, 0x0a, 0x0b,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x47, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04,
	0x2a, 0x30, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c,
	0x10, 0x03, 0x32, 0xb7, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x55, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a,
	0x22, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x22, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x31, 0x38, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_proto_goTypes = []interface{}{
	(Environment)(0),              // 0: signal18.config_store.v1.Environment
	(Type)(0),                     // 1: signal18.config_store.v1.Type
	(*Property)(nil),              // 2: signal18.config_store.v1.Property
	(*Value)(nil),                 // 3: signal18.config_store.v1.Value
	(*Query)(nil),                 // 4: signal18.config_store.v1.Query
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_config_proto_depIdxs = []int32{
	3, // 0: signal18.config_store.v1.Property.values:type_name -> signal18.config_store.v1.Value
	0, // 1: signal18.config_store.v1.Property.environment:type_name -> signal18.config_store.v1.Environment
	5, // 2: signal18.config_store.v1.Property.created:type_name -> google.protobuf.Timestamp
	5, // 3: signal18.config_store.v1.Property.deleted:type_name -> google.protobuf.Timestamp
	1, // 4: signal18.config_store.v1.Value.type:type_name -> signal18.config_store.v1.Type
	2, // 5: signal18.config_store.v1.Query.property:type_name -> signal18.config_store.v1.Property
	2, // 6: signal18.config_store.v1.ConfigStore.Store:input_type -> signal18.config_store.v1.Property
	4, // 7: signal18.config_store.v1.ConfigStore.Search:input_type -> signal18.config_store.v1.Query
	2, // 8: signal18.config_store.v1.ConfigStore.Store:output_type -> signal18.config_store.v1.Property
	2, // 9: signal18.config_store.v1.ConfigStore.Search:output_type -> signal18.config_store.v1.Property
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		EnumInfos:         file_config_proto_enumTypes,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}