// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commonfate/authz/v1alpha1/entity.proto

package authzv1alpha1

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

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_entity_proto_rawDescGZIP(), []int{0}
}

func (x *Attribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Attribute) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Value_Str
	//	*Value_Bool
	//	*Value_Long
	//	*Value_Entity
	//	*Value_Record
	//	*Value_Set
	Value isValue_Value `protobuf_oneof:"value"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[1]
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
	return file_commonfate_authz_v1alpha1_entity_proto_rawDescGZIP(), []int{1}
}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Value) GetStr() string {
	if x, ok := x.GetValue().(*Value_Str); ok {
		return x.Str
	}
	return ""
}

func (x *Value) GetBool() bool {
	if x, ok := x.GetValue().(*Value_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *Value) GetLong() int64 {
	if x, ok := x.GetValue().(*Value_Long); ok {
		return x.Long
	}
	return 0
}

func (x *Value) GetEntity() *UID {
	if x, ok := x.GetValue().(*Value_Entity); ok {
		return x.Entity
	}
	return nil
}

func (x *Value) GetRecord() *Record {
	if x, ok := x.GetValue().(*Value_Record); ok {
		return x.Record
	}
	return nil
}

func (x *Value) GetSet() *Set {
	if x, ok := x.GetValue().(*Value_Set); ok {
		return x.Set
	}
	return nil
}

type isValue_Value interface {
	isValue_Value()
}

type Value_Str struct {
	Str string `protobuf:"bytes,1,opt,name=str,proto3,oneof"`
}

type Value_Bool struct {
	Bool bool `protobuf:"varint,2,opt,name=bool,proto3,oneof"`
}

type Value_Long struct {
	Long int64 `protobuf:"varint,3,opt,name=long,proto3,oneof"`
}

type Value_Entity struct {
	Entity *UID `protobuf:"bytes,4,opt,name=entity,proto3,oneof"`
}

type Value_Record struct {
	Record *Record `protobuf:"bytes,5,opt,name=record,proto3,oneof"`
}

type Value_Set struct {
	Set *Set `protobuf:"bytes,6,opt,name=set,proto3,oneof"`
}

func (*Value_Str) isValue_Value() {}

func (*Value_Bool) isValue_Value() {}

func (*Value_Long) isValue_Value() {}

func (*Value_Entity) isValue_Value() {}

func (*Value_Record) isValue_Value() {}

func (*Value_Set) isValue_Value() {}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*Attribute `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_entity_proto_rawDescGZIP(), []int{2}
}

func (x *Record) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type Set struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Set) Reset() {
	*x = Set{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Set) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Set) ProtoMessage() {}

func (x *Set) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Set.ProtoReflect.Descriptor instead.
func (*Set) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_entity_proto_rawDescGZIP(), []int{3}
}

func (x *Set) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type UID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UID) Reset() {
	*x = UID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UID) ProtoMessage() {}

func (x *UID) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UID.ProtoReflect.Descriptor instead.
func (*UID) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_entity_proto_rawDescGZIP(), []int{4}
}

func (x *UID) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        *UID         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Attributes []*Attribute `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Parents    []*UID       `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_entity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_entity_proto_rawDescGZIP(), []int{5}
}

func (x *Entity) GetUid() *UID {
	if x != nil {
		return x.Uid
	}
	return nil
}

func (x *Entity) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Entity) GetParents() []*UID {
	if x != nil {
		return x.Parents
	}
	return nil
}

var File_commonfate_authz_v1alpha1_entity_proto protoreflect.FileDescriptor

var file_commonfate_authz_v1alpha1_entity_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x22, 0x55, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xfb, 0x01, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x14,
	0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x04,
	0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x55, 0x49, 0x44, 0x48, 0x00, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3b,
	0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x03, 0x73,
	0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x03, 0x73, 0x65, 0x74, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4e, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x44, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12,
	0x38, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x03, 0x55, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x30, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x49, 0x44, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x44, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x49, 0x44, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0xfb, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x69, 0x65, 0x6d, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41,
	0x58, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a,
	0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_authz_v1alpha1_entity_proto_rawDescOnce sync.Once
	file_commonfate_authz_v1alpha1_entity_proto_rawDescData = file_commonfate_authz_v1alpha1_entity_proto_rawDesc
)

func file_commonfate_authz_v1alpha1_entity_proto_rawDescGZIP() []byte {
	file_commonfate_authz_v1alpha1_entity_proto_rawDescOnce.Do(func() {
		file_commonfate_authz_v1alpha1_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_authz_v1alpha1_entity_proto_rawDescData)
	})
	return file_commonfate_authz_v1alpha1_entity_proto_rawDescData
}

var file_commonfate_authz_v1alpha1_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_commonfate_authz_v1alpha1_entity_proto_goTypes = []interface{}{
	(*Attribute)(nil), // 0: commonfate.authz.v1alpha1.Attribute
	(*Value)(nil),     // 1: commonfate.authz.v1alpha1.Value
	(*Record)(nil),    // 2: commonfate.authz.v1alpha1.Record
	(*Set)(nil),       // 3: commonfate.authz.v1alpha1.Set
	(*UID)(nil),       // 4: commonfate.authz.v1alpha1.UID
	(*Entity)(nil),    // 5: commonfate.authz.v1alpha1.Entity
}
var file_commonfate_authz_v1alpha1_entity_proto_depIdxs = []int32{
	1, // 0: commonfate.authz.v1alpha1.Attribute.value:type_name -> commonfate.authz.v1alpha1.Value
	4, // 1: commonfate.authz.v1alpha1.Value.entity:type_name -> commonfate.authz.v1alpha1.UID
	2, // 2: commonfate.authz.v1alpha1.Value.record:type_name -> commonfate.authz.v1alpha1.Record
	3, // 3: commonfate.authz.v1alpha1.Value.set:type_name -> commonfate.authz.v1alpha1.Set
	0, // 4: commonfate.authz.v1alpha1.Record.attributes:type_name -> commonfate.authz.v1alpha1.Attribute
	1, // 5: commonfate.authz.v1alpha1.Set.values:type_name -> commonfate.authz.v1alpha1.Value
	4, // 6: commonfate.authz.v1alpha1.Entity.uid:type_name -> commonfate.authz.v1alpha1.UID
	0, // 7: commonfate.authz.v1alpha1.Entity.attributes:type_name -> commonfate.authz.v1alpha1.Attribute
	4, // 8: commonfate.authz.v1alpha1.Entity.parents:type_name -> commonfate.authz.v1alpha1.UID
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_commonfate_authz_v1alpha1_entity_proto_init() }
func file_commonfate_authz_v1alpha1_entity_proto_init() {
	if File_commonfate_authz_v1alpha1_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_authz_v1alpha1_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_commonfate_authz_v1alpha1_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_commonfate_authz_v1alpha1_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_commonfate_authz_v1alpha1_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Set); i {
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
		file_commonfate_authz_v1alpha1_entity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UID); i {
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
		file_commonfate_authz_v1alpha1_entity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
	file_commonfate_authz_v1alpha1_entity_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Value_Str)(nil),
		(*Value_Bool)(nil),
		(*Value_Long)(nil),
		(*Value_Entity)(nil),
		(*Value_Record)(nil),
		(*Value_Set)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_authz_v1alpha1_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_authz_v1alpha1_entity_proto_goTypes,
		DependencyIndexes: file_commonfate_authz_v1alpha1_entity_proto_depIdxs,
		MessageInfos:      file_commonfate_authz_v1alpha1_entity_proto_msgTypes,
	}.Build()
	File_commonfate_authz_v1alpha1_entity_proto = out.File
	file_commonfate_authz_v1alpha1_entity_proto_rawDesc = nil
	file_commonfate_authz_v1alpha1_entity_proto_goTypes = nil
	file_commonfate_authz_v1alpha1_entity_proto_depIdxs = nil
}
