// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.11.2
// source: rpc/group.proto

package group

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

// 开启打卡记录
type StartCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *StartCheckRequest) Reset() {
	*x = StartCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCheckRequest) ProtoMessage() {}

func (x *StartCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCheckRequest.ProtoReflect.Descriptor instead.
func (*StartCheckRequest) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{0}
}

func (x *StartCheckRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type StartCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckId         uint32 `protobuf:"varint,1,opt,name=check_id,json=checkId,proto3" json:"check_id,omitempty"`
	UserId          uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ContinuousDays  int32  `protobuf:"varint,3,opt,name=continuous_days,json=continuousDays,proto3" json:"continuous_days,omitempty"`
	CreateTime      int32  `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	LastCheckInTime int32  `protobuf:"varint,5,opt,name=last_check_in_time,json=lastCheckInTime,proto3" json:"last_check_in_time,omitempty"`
}

func (x *StartCheckResponse) Reset() {
	*x = StartCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCheckResponse) ProtoMessage() {}

func (x *StartCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCheckResponse.ProtoReflect.Descriptor instead.
func (*StartCheckResponse) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{1}
}

func (x *StartCheckResponse) GetCheckId() uint32 {
	if x != nil {
		return x.CheckId
	}
	return 0
}

func (x *StartCheckResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *StartCheckResponse) GetContinuousDays() int32 {
	if x != nil {
		return x.ContinuousDays
	}
	return 0
}

func (x *StartCheckResponse) GetCreateTime() int32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *StartCheckResponse) GetLastCheckInTime() int32 {
	if x != nil {
		return x.LastCheckInTime
	}
	return 0
}

// 书法记录的简单信息
type RecordSimpleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId   uint32  `protobuf:"varint,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	UserId     uint32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content    string  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Image      string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Score      float32 `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
	CreateTime int32   `protobuf:"varint,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *RecordSimpleInfo) Reset() {
	*x = RecordSimpleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordSimpleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSimpleInfo) ProtoMessage() {}

func (x *RecordSimpleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSimpleInfo.ProtoReflect.Descriptor instead.
func (*RecordSimpleInfo) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{2}
}

func (x *RecordSimpleInfo) GetRecordId() uint32 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

func (x *RecordSimpleInfo) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RecordSimpleInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *RecordSimpleInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *RecordSimpleInfo) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *RecordSimpleInfo) GetCreateTime() int32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

// 上传书法信息
type CreateRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content string  `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Image   string  `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Score   float32 `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *CreateRecordRequest) Reset() {
	*x = CreateRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecordRequest) ProtoMessage() {}

func (x *CreateRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateRecordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRecordRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRecordRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateRecordRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateRecordRequest) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type CreateRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordInfo *RecordSimpleInfo `protobuf:"bytes,1,opt,name=record_info,json=recordInfo,proto3" json:"record_info,omitempty"`
}

func (x *CreateRecordResponse) Reset() {
	*x = CreateRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecordResponse) ProtoMessage() {}

func (x *CreateRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecordResponse.ProtoReflect.Descriptor instead.
func (*CreateRecordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRecordResponse) GetRecordInfo() *RecordSimpleInfo {
	if x != nil {
		return x.RecordInfo
	}
	return nil
}

// 查看某人的书法记录
type LookRecordByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LookRecordByUserIdRequest) Reset() {
	*x = LookRecordByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookRecordByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookRecordByUserIdRequest) ProtoMessage() {}

func (x *LookRecordByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookRecordByUserIdRequest.ProtoReflect.Descriptor instead.
func (*LookRecordByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{5}
}

func (x *LookRecordByUserIdRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LookRecordByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordInfo []*RecordSimpleInfo `protobuf:"bytes,1,rep,name=record_info,json=recordInfo,proto3" json:"record_info,omitempty"`
}

func (x *LookRecordByUserIdResponse) Reset() {
	*x = LookRecordByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookRecordByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookRecordByUserIdResponse) ProtoMessage() {}

func (x *LookRecordByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookRecordByUserIdResponse.ProtoReflect.Descriptor instead.
func (*LookRecordByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{6}
}

func (x *LookRecordByUserIdResponse) GetRecordInfo() []*RecordSimpleInfo {
	if x != nil {
		return x.RecordInfo
	}
	return nil
}

// 检查打卡模式是否开启
type CheckPunchCardModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CheckPunchCardModelRequest) Reset() {
	*x = CheckPunchCardModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPunchCardModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPunchCardModelRequest) ProtoMessage() {}

func (x *CheckPunchCardModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPunchCardModelRequest.ProtoReflect.Descriptor instead.
func (*CheckPunchCardModelRequest) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{7}
}

func (x *CheckPunchCardModelRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckPunchCardModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CheckPunchCardModelResponse) Reset() {
	*x = CheckPunchCardModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPunchCardModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPunchCardModelResponse) ProtoMessage() {}

func (x *CheckPunchCardModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPunchCardModelResponse.ProtoReflect.Descriptor instead.
func (*CheckPunchCardModelResponse) Descriptor() ([]byte, []int) {
	return file_rpc_group_proto_rawDescGZIP(), []int{8}
}

func (x *CheckPunchCardModelResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

var File_rpc_group_proto protoreflect.FileDescriptor

var file_rpc_group_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2c, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xbf, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x5f,
	0x64, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x74, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x50, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x34, 0x0a, 0x19, 0x4c, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x1a, 0x4c, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x35, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x43, 0x61,
	0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x1b, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x75, 0x6e, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xcc, 0x02, 0x0a, 0x05, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x41, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x59, 0x0a, 0x12, 0x4c, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4c,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x13, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x75, 0x6e, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x50, 0x75, 0x6e, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_group_proto_rawDescOnce sync.Once
	file_rpc_group_proto_rawDescData = file_rpc_group_proto_rawDesc
)

func file_rpc_group_proto_rawDescGZIP() []byte {
	file_rpc_group_proto_rawDescOnce.Do(func() {
		file_rpc_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_group_proto_rawDescData)
	})
	return file_rpc_group_proto_rawDescData
}

var file_rpc_group_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rpc_group_proto_goTypes = []interface{}{
	(*StartCheckRequest)(nil),           // 0: group.StartCheckRequest
	(*StartCheckResponse)(nil),          // 1: group.StartCheckResponse
	(*RecordSimpleInfo)(nil),            // 2: group.RecordSimpleInfo
	(*CreateRecordRequest)(nil),         // 3: group.CreateRecordRequest
	(*CreateRecordResponse)(nil),        // 4: group.CreateRecordResponse
	(*LookRecordByUserIdRequest)(nil),   // 5: group.LookRecordByUserIdRequest
	(*LookRecordByUserIdResponse)(nil),  // 6: group.LookRecordByUserIdResponse
	(*CheckPunchCardModelRequest)(nil),  // 7: group.CheckPunchCardModelRequest
	(*CheckPunchCardModelResponse)(nil), // 8: group.CheckPunchCardModelResponse
}
var file_rpc_group_proto_depIdxs = []int32{
	2, // 0: group.CreateRecordResponse.record_info:type_name -> group.RecordSimpleInfo
	2, // 1: group.LookRecordByUserIdResponse.record_info:type_name -> group.RecordSimpleInfo
	0, // 2: group.Group.StartCheck:input_type -> group.StartCheckRequest
	3, // 3: group.Group.CreateRecord:input_type -> group.CreateRecordRequest
	5, // 4: group.Group.LookRecordByUserId:input_type -> group.LookRecordByUserIdRequest
	7, // 5: group.Group.CheckPunchCardModel:input_type -> group.CheckPunchCardModelRequest
	1, // 6: group.Group.StartCheck:output_type -> group.StartCheckResponse
	4, // 7: group.Group.CreateRecord:output_type -> group.CreateRecordResponse
	6, // 8: group.Group.LookRecordByUserId:output_type -> group.LookRecordByUserIdResponse
	8, // 9: group.Group.CheckPunchCardModel:output_type -> group.CheckPunchCardModelResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_group_proto_init() }
func file_rpc_group_proto_init() {
	if File_rpc_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartCheckRequest); i {
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
		file_rpc_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartCheckResponse); i {
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
		file_rpc_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordSimpleInfo); i {
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
		file_rpc_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRecordRequest); i {
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
		file_rpc_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRecordResponse); i {
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
		file_rpc_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookRecordByUserIdRequest); i {
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
		file_rpc_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookRecordByUserIdResponse); i {
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
		file_rpc_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPunchCardModelRequest); i {
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
		file_rpc_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPunchCardModelResponse); i {
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
			RawDescriptor: file_rpc_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_group_proto_goTypes,
		DependencyIndexes: file_rpc_group_proto_depIdxs,
		MessageInfos:      file_rpc_group_proto_msgTypes,
	}.Build()
	File_rpc_group_proto = out.File
	file_rpc_group_proto_rawDesc = nil
	file_rpc_group_proto_goTypes = nil
	file_rpc_group_proto_depIdxs = nil
}
