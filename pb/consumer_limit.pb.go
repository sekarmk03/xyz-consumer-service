// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: consumer_limit.proto

package pb

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

type ConsumerLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ConsumerId     uint64 `protobuf:"varint,2,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	Tenor          uint32 `protobuf:"varint,3,opt,name=tenor,proto3" json:"tenor,omitempty"`
	LimitAmount    uint64 `protobuf:"varint,4,opt,name=limit_amount,json=limitAmount,proto3" json:"limit_amount,omitempty"`
	LimitAvailable uint64 `protobuf:"varint,5,opt,name=limit_available,json=limitAvailable,proto3" json:"limit_available,omitempty"`
	CreatedAt      string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ConsumerLimit) Reset() {
	*x = ConsumerLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerLimit) ProtoMessage() {}

func (x *ConsumerLimit) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerLimit.ProtoReflect.Descriptor instead.
func (*ConsumerLimit) Descriptor() ([]byte, []int) {
	return file_consumer_limit_proto_rawDescGZIP(), []int{0}
}

func (x *ConsumerLimit) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConsumerLimit) GetConsumerId() uint64 {
	if x != nil {
		return x.ConsumerId
	}
	return 0
}

func (x *ConsumerLimit) GetTenor() uint32 {
	if x != nil {
		return x.Tenor
	}
	return 0
}

func (x *ConsumerLimit) GetLimitAmount() uint64 {
	if x != nil {
		return x.LimitAmount
	}
	return 0
}

func (x *ConsumerLimit) GetLimitAvailable() uint64 {
	if x != nil {
		return x.LimitAvailable
	}
	return 0
}

func (x *ConsumerLimit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ConsumerLimit) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ConsumerLimitListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*ConsumerLimit `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ConsumerLimitListResponse) Reset() {
	*x = ConsumerLimitListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_limit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerLimitListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerLimitListResponse) ProtoMessage() {}

func (x *ConsumerLimitListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_limit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerLimitListResponse.ProtoReflect.Descriptor instead.
func (*ConsumerLimitListResponse) Descriptor() ([]byte, []int) {
	return file_consumer_limit_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumerLimitListResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConsumerLimitListResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConsumerLimitListResponse) GetData() []*ConsumerLimit {
	if x != nil {
		return x.Data
	}
	return nil
}

type ConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId uint64 `protobuf:"varint,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
}

func (x *ConsumerRequest) Reset() {
	*x = ConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_limit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerRequest) ProtoMessage() {}

func (x *ConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_limit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerRequest.ProtoReflect.Descriptor instead.
func (*ConsumerRequest) Descriptor() ([]byte, []int) {
	return file_consumer_limit_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumerRequest) GetConsumerId() uint64 {
	if x != nil {
		return x.ConsumerId
	}
	return 0
}

type UpdateAvailableLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerId        uint64 `protobuf:"varint,1,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	Tenor             uint32 `protobuf:"varint,2,opt,name=tenor,proto3" json:"tenor,omitempty"`
	AmountTransaction uint64 `protobuf:"varint,3,opt,name=amount_transaction,json=amountTransaction,proto3" json:"amount_transaction,omitempty"`
}

func (x *UpdateAvailableLimitRequest) Reset() {
	*x = UpdateAvailableLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_limit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAvailableLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvailableLimitRequest) ProtoMessage() {}

func (x *UpdateAvailableLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_limit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvailableLimitRequest.ProtoReflect.Descriptor instead.
func (*UpdateAvailableLimitRequest) Descriptor() ([]byte, []int) {
	return file_consumer_limit_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAvailableLimitRequest) GetConsumerId() uint64 {
	if x != nil {
		return x.ConsumerId
	}
	return 0
}

func (x *UpdateAvailableLimitRequest) GetTenor() uint32 {
	if x != nil {
		return x.Tenor
	}
	return 0
}

func (x *UpdateAvailableLimitRequest) GetAmountTransaction() uint64 {
	if x != nil {
		return x.AmountTransaction
	}
	return 0
}

type ConsumerLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *ConsumerLimit `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConsumerLimitResponse) Reset() {
	*x = ConsumerLimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_limit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerLimitResponse) ProtoMessage() {}

func (x *ConsumerLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_limit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerLimitResponse.ProtoReflect.Descriptor instead.
func (*ConsumerLimitResponse) Descriptor() ([]byte, []int) {
	return file_consumer_limit_proto_rawDescGZIP(), []int{4}
}

func (x *ConsumerLimitResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConsumerLimitResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConsumerLimitResponse) GetData() *ConsumerLimit {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_consumer_limit_proto protoreflect.FileDescriptor

var file_consumer_limit_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x78, 0x79, 0x7a, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x22, 0xe0, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x76, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x78,
	0x79, 0x7a, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x32, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x83, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x74, 0x65, 0x6e, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x11, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x78, 0x79, 0x7a,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xa8, 0x02, 0x0a, 0x14, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x78, 0x79, 0x7a, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x78, 0x79, 0x7a, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x17, 0x2e, 0x78, 0x79, 0x7a,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x1a, 0x1f, 0x2e, 0x78, 0x79, 0x7a, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x25, 0x2e, 0x78,
	0x79, 0x7a, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x78, 0x79, 0x7a, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consumer_limit_proto_rawDescOnce sync.Once
	file_consumer_limit_proto_rawDescData = file_consumer_limit_proto_rawDesc
)

func file_consumer_limit_proto_rawDescGZIP() []byte {
	file_consumer_limit_proto_rawDescOnce.Do(func() {
		file_consumer_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_consumer_limit_proto_rawDescData)
	})
	return file_consumer_limit_proto_rawDescData
}

var file_consumer_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_consumer_limit_proto_goTypes = []interface{}{
	(*ConsumerLimit)(nil),               // 0: xyz_grpc.ConsumerLimit
	(*ConsumerLimitListResponse)(nil),   // 1: xyz_grpc.ConsumerLimitListResponse
	(*ConsumerRequest)(nil),             // 2: xyz_grpc.ConsumerRequest
	(*UpdateAvailableLimitRequest)(nil), // 3: xyz_grpc.UpdateAvailableLimitRequest
	(*ConsumerLimitResponse)(nil),       // 4: xyz_grpc.ConsumerLimitResponse
}
var file_consumer_limit_proto_depIdxs = []int32{
	0, // 0: xyz_grpc.ConsumerLimitListResponse.data:type_name -> xyz_grpc.ConsumerLimit
	0, // 1: xyz_grpc.ConsumerLimitResponse.data:type_name -> xyz_grpc.ConsumerLimit
	2, // 2: xyz_grpc.ConsumerLimitService.GetConsumerLimitsByConsumerId:input_type -> xyz_grpc.ConsumerRequest
	0, // 3: xyz_grpc.ConsumerLimitService.CreateConsumerLimit:input_type -> xyz_grpc.ConsumerLimit
	3, // 4: xyz_grpc.ConsumerLimitService.UpdateAvailableLimit:input_type -> xyz_grpc.UpdateAvailableLimitRequest
	1, // 5: xyz_grpc.ConsumerLimitService.GetConsumerLimitsByConsumerId:output_type -> xyz_grpc.ConsumerLimitListResponse
	4, // 6: xyz_grpc.ConsumerLimitService.CreateConsumerLimit:output_type -> xyz_grpc.ConsumerLimitResponse
	4, // 7: xyz_grpc.ConsumerLimitService.UpdateAvailableLimit:output_type -> xyz_grpc.ConsumerLimitResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_consumer_limit_proto_init() }
func file_consumer_limit_proto_init() {
	if File_consumer_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consumer_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerLimit); i {
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
		file_consumer_limit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerLimitListResponse); i {
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
		file_consumer_limit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerRequest); i {
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
		file_consumer_limit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAvailableLimitRequest); i {
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
		file_consumer_limit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerLimitResponse); i {
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
			RawDescriptor: file_consumer_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consumer_limit_proto_goTypes,
		DependencyIndexes: file_consumer_limit_proto_depIdxs,
		MessageInfos:      file_consumer_limit_proto_msgTypes,
	}.Build()
	File_consumer_limit_proto = out.File
	file_consumer_limit_proto_rawDesc = nil
	file_consumer_limit_proto_goTypes = nil
	file_consumer_limit_proto_depIdxs = nil
}
