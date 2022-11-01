// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: internal/message/message.proto

package message

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

type PublicMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*PublicMessage_Join_
	//	*PublicMessage_Leave_
	//	*PublicMessage_Move_
	//	*PublicMessage_LieDown_
	//	*PublicMessage_Underlay_
	Action isPublicMessage_Action `protobuf_oneof:"action"`
}

func (x *PublicMessage) Reset() {
	*x = PublicMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicMessage) ProtoMessage() {}

func (x *PublicMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicMessage.ProtoReflect.Descriptor instead.
func (*PublicMessage) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{0}
}

func (m *PublicMessage) GetAction() isPublicMessage_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *PublicMessage) GetJoin() *PublicMessage_Join {
	if x, ok := x.GetAction().(*PublicMessage_Join_); ok {
		return x.Join
	}
	return nil
}

func (x *PublicMessage) GetLeave() *PublicMessage_Leave {
	if x, ok := x.GetAction().(*PublicMessage_Leave_); ok {
		return x.Leave
	}
	return nil
}

func (x *PublicMessage) GetMove() *PublicMessage_Move {
	if x, ok := x.GetAction().(*PublicMessage_Move_); ok {
		return x.Move
	}
	return nil
}

func (x *PublicMessage) GetLieDown() *PublicMessage_LieDown {
	if x, ok := x.GetAction().(*PublicMessage_LieDown_); ok {
		return x.LieDown
	}
	return nil
}

func (x *PublicMessage) GetUnderlay() *PublicMessage_Underlay {
	if x, ok := x.GetAction().(*PublicMessage_Underlay_); ok {
		return x.Underlay
	}
	return nil
}

type isPublicMessage_Action interface {
	isPublicMessage_Action()
}

type PublicMessage_Join_ struct {
	Join *PublicMessage_Join `protobuf:"bytes,4,opt,name=join,proto3,oneof"`
}

type PublicMessage_Leave_ struct {
	Leave *PublicMessage_Leave `protobuf:"bytes,5,opt,name=leave,proto3,oneof"`
}

type PublicMessage_Move_ struct {
	Move *PublicMessage_Move `protobuf:"bytes,6,opt,name=move,proto3,oneof"`
}

type PublicMessage_LieDown_ struct {
	LieDown *PublicMessage_LieDown `protobuf:"bytes,7,opt,name=lie_down,json=lieDown,proto3,oneof"`
}

type PublicMessage_Underlay_ struct {
	Underlay *PublicMessage_Underlay `protobuf:"bytes,8,opt,name=underlay,proto3,oneof"`
}

func (*PublicMessage_Join_) isPublicMessage_Action() {}

func (*PublicMessage_Leave_) isPublicMessage_Action() {}

func (*PublicMessage_Move_) isPublicMessage_Action() {}

func (*PublicMessage_LieDown_) isPublicMessage_Action() {}

func (*PublicMessage_Underlay_) isPublicMessage_Action() {}

type UpcomingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*UpcomingMessage_Move_
	//	*UpcomingMessage_LieDown_
	Action isUpcomingMessage_Action `protobuf_oneof:"action"`
}

func (x *UpcomingMessage) Reset() {
	*x = UpcomingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpcomingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpcomingMessage) ProtoMessage() {}

func (x *UpcomingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpcomingMessage.ProtoReflect.Descriptor instead.
func (*UpcomingMessage) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{1}
}

func (m *UpcomingMessage) GetAction() isUpcomingMessage_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *UpcomingMessage) GetMove() *UpcomingMessage_Move {
	if x, ok := x.GetAction().(*UpcomingMessage_Move_); ok {
		return x.Move
	}
	return nil
}

func (x *UpcomingMessage) GetLieDown() *UpcomingMessage_LieDown {
	if x, ok := x.GetAction().(*UpcomingMessage_LieDown_); ok {
		return x.LieDown
	}
	return nil
}

type isUpcomingMessage_Action interface {
	isUpcomingMessage_Action()
}

type UpcomingMessage_Move_ struct {
	Move *UpcomingMessage_Move `protobuf:"bytes,3,opt,name=move,proto3,oneof"`
}

type UpcomingMessage_LieDown_ struct {
	LieDown *UpcomingMessage_LieDown `protobuf:"bytes,4,opt,name=lie_down,json=lieDown,proto3,oneof"`
}

func (*UpcomingMessage_Move_) isUpcomingMessage_Action() {}

func (*UpcomingMessage_LieDown_) isUpcomingMessage_Action() {}

type PublicMessage_Join struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid int32 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
}

func (x *PublicMessage_Join) Reset() {
	*x = PublicMessage_Join{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicMessage_Join) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicMessage_Join) ProtoMessage() {}

func (x *PublicMessage_Join) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicMessage_Join.ProtoReflect.Descriptor instead.
func (*PublicMessage_Join) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PublicMessage_Join) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

type PublicMessage_Leave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid int32 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
}

func (x *PublicMessage_Leave) Reset() {
	*x = PublicMessage_Leave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicMessage_Leave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicMessage_Leave) ProtoMessage() {}

func (x *PublicMessage_Leave) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicMessage_Leave.ProtoReflect.Descriptor instead.
func (*PublicMessage_Leave) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PublicMessage_Leave) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

type PublicMessage_Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid int32   `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	X   float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y   float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *PublicMessage_Move) Reset() {
	*x = PublicMessage_Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicMessage_Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicMessage_Move) ProtoMessage() {}

func (x *PublicMessage_Move) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicMessage_Move.ProtoReflect.Descriptor instead.
func (*PublicMessage_Move) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{0, 2}
}

func (x *PublicMessage_Move) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *PublicMessage_Move) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PublicMessage_Move) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type PublicMessage_LieDown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid int32 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Bed int32 `protobuf:"varint,2,opt,name=bed,proto3" json:"bed,omitempty"`
}

func (x *PublicMessage_LieDown) Reset() {
	*x = PublicMessage_LieDown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicMessage_LieDown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicMessage_LieDown) ProtoMessage() {}

func (x *PublicMessage_LieDown) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicMessage_LieDown.ProtoReflect.Descriptor instead.
func (*PublicMessage_LieDown) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{0, 3}
}

func (x *PublicMessage_LieDown) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *PublicMessage_LieDown) GetBed() int32 {
	if x != nil {
		return x.Bed
	}
	return 0
}

type PublicMessage_Underlay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid int32 `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
}

func (x *PublicMessage_Underlay) Reset() {
	*x = PublicMessage_Underlay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicMessage_Underlay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicMessage_Underlay) ProtoMessage() {}

func (x *PublicMessage_Underlay) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicMessage_Underlay.ProtoReflect.Descriptor instead.
func (*PublicMessage_Underlay) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{0, 4}
}

func (x *PublicMessage_Underlay) GetTid() int32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

type UpcomingMessage_Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *UpcomingMessage_Move) Reset() {
	*x = UpcomingMessage_Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpcomingMessage_Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpcomingMessage_Move) ProtoMessage() {}

func (x *UpcomingMessage_Move) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpcomingMessage_Move.ProtoReflect.Descriptor instead.
func (*UpcomingMessage_Move) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{1, 0}
}

func (x *UpcomingMessage_Move) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *UpcomingMessage_Move) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type UpcomingMessage_LieDown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bed int32 `protobuf:"varint,1,opt,name=bed,proto3" json:"bed,omitempty"`
}

func (x *UpcomingMessage_LieDown) Reset() {
	*x = UpcomingMessage_LieDown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_message_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpcomingMessage_LieDown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpcomingMessage_LieDown) ProtoMessage() {}

func (x *UpcomingMessage_LieDown) ProtoReflect() protoreflect.Message {
	mi := &file_internal_message_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpcomingMessage_LieDown.ProtoReflect.Descriptor instead.
func (*UpcomingMessage_LieDown) Descriptor() ([]byte, []int) {
	return file_internal_message_message_proto_rawDescGZIP(), []int{1, 1}
}

func (x *UpcomingMessage_LieDown) GetBed() int32 {
	if x != nil {
		return x.Bed
	}
	return 0
}

var File_internal_message_message_proto protoreflect.FileDescriptor

var file_internal_message_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe9, 0x03, 0x0a, 0x0d, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6a,
	0x6f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x34,
	0x0a, 0x05, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x48, 0x00, 0x52, 0x05, 0x6c,
	0x65, 0x61, 0x76, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x48,
	0x00, 0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x6c, 0x69, 0x65, 0x5f, 0x64,
	0x6f, 0x77, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4c, 0x69, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x61, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55,
	0x6e, 0x64, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x48, 0x00, 0x52, 0x08, 0x75, 0x6e, 0x64, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x1a, 0x18, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x1a, 0x19, 0x0a,
	0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x1a, 0x34, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74,
	0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x1a, 0x2d,
	0x0a, 0x07, 0x4c, 0x69, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x65, 0x64, 0x1a, 0x1c, 0x0a,
	0x08, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd0, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x6f, 0x76,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x3d,
	0x0a, 0x08, 0x6c, 0x69, 0x65, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x63, 0x6f, 0x6d,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x65, 0x44, 0x6f,
	0x77, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x1a, 0x22, 0x0a,
	0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x79, 0x1a, 0x1b, 0x0a, 0x07, 0x4c, 0x69, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x62, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x65, 0x64, 0x42, 0x08,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_message_message_proto_rawDescOnce sync.Once
	file_internal_message_message_proto_rawDescData = file_internal_message_message_proto_rawDesc
)

func file_internal_message_message_proto_rawDescGZIP() []byte {
	file_internal_message_message_proto_rawDescOnce.Do(func() {
		file_internal_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_message_message_proto_rawDescData)
	})
	return file_internal_message_message_proto_rawDescData
}

var file_internal_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_message_message_proto_goTypes = []interface{}{
	(*PublicMessage)(nil),           // 0: message.PublicMessage
	(*UpcomingMessage)(nil),         // 1: message.UpcomingMessage
	(*PublicMessage_Join)(nil),      // 2: message.PublicMessage.Join
	(*PublicMessage_Leave)(nil),     // 3: message.PublicMessage.Leave
	(*PublicMessage_Move)(nil),      // 4: message.PublicMessage.Move
	(*PublicMessage_LieDown)(nil),   // 5: message.PublicMessage.LieDown
	(*PublicMessage_Underlay)(nil),  // 6: message.PublicMessage.Underlay
	(*UpcomingMessage_Move)(nil),    // 7: message.UpcomingMessage.Move
	(*UpcomingMessage_LieDown)(nil), // 8: message.UpcomingMessage.LieDown
}
var file_internal_message_message_proto_depIdxs = []int32{
	2, // 0: message.PublicMessage.join:type_name -> message.PublicMessage.Join
	3, // 1: message.PublicMessage.leave:type_name -> message.PublicMessage.Leave
	4, // 2: message.PublicMessage.move:type_name -> message.PublicMessage.Move
	5, // 3: message.PublicMessage.lie_down:type_name -> message.PublicMessage.LieDown
	6, // 4: message.PublicMessage.underlay:type_name -> message.PublicMessage.Underlay
	7, // 5: message.UpcomingMessage.move:type_name -> message.UpcomingMessage.Move
	8, // 6: message.UpcomingMessage.lie_down:type_name -> message.UpcomingMessage.LieDown
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_internal_message_message_proto_init() }
func file_internal_message_message_proto_init() {
	if File_internal_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicMessage); i {
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
		file_internal_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpcomingMessage); i {
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
		file_internal_message_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicMessage_Join); i {
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
		file_internal_message_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicMessage_Leave); i {
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
		file_internal_message_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicMessage_Move); i {
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
		file_internal_message_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicMessage_LieDown); i {
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
		file_internal_message_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicMessage_Underlay); i {
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
		file_internal_message_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpcomingMessage_Move); i {
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
		file_internal_message_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpcomingMessage_LieDown); i {
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
	file_internal_message_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PublicMessage_Join_)(nil),
		(*PublicMessage_Leave_)(nil),
		(*PublicMessage_Move_)(nil),
		(*PublicMessage_LieDown_)(nil),
		(*PublicMessage_Underlay_)(nil),
	}
	file_internal_message_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UpcomingMessage_Move_)(nil),
		(*UpcomingMessage_LieDown_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_message_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_message_message_proto_goTypes,
		DependencyIndexes: file_internal_message_message_proto_depIdxs,
		MessageInfos:      file_internal_message_message_proto_msgTypes,
	}.Build()
	File_internal_message_message_proto = out.File
	file_internal_message_message_proto_rawDesc = nil
	file_internal_message_message_proto_goTypes = nil
	file_internal_message_message_proto_depIdxs = nil
}
