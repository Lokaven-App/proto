// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: delivery/grpc/wallet/wallet.proto

package wallet

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	UserUid   string               `protobuf:"bytes,2,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
	Amount    int32                `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_delivery_grpc_wallet_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *Balance) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Balance) GetUserUid() string {
	if x != nil {
		return x.UserUid
	}
	return ""
}

func (x *Balance) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Balance) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Balances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balances []*Balance `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (x *Balances) Reset() {
	*x = Balances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balances) ProtoMessage() {}

func (x *Balances) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balances.ProtoReflect.Descriptor instead.
func (*Balances) Descriptor() ([]byte, []int) {
	return file_delivery_grpc_wallet_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *Balances) GetBalances() []*Balance {
	if x != nil {
		return x.Balances
	}
	return nil
}

type BalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Balance `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BalanceResponse) Reset() {
	*x = BalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResponse) ProtoMessage() {}

func (x *BalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResponse.ProtoReflect.Descriptor instead.
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return file_delivery_grpc_wallet_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *BalanceResponse) GetData() *Balance {
	if x != nil {
		return x.Data
	}
	return nil
}

type BalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *BalanceRequest) Reset() {
	*x = BalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceRequest) ProtoMessage() {}

func (x *BalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceRequest.ProtoReflect.Descriptor instead.
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return file_delivery_grpc_wallet_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *BalanceRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type LogBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid           string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	BalanceUid    string               `protobuf:"bytes,2,opt,name=balance_uid,json=balanceUid,proto3" json:"balance_uid,omitempty"`
	UserUid       string               `protobuf:"bytes,3,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
	TourUid       string               `protobuf:"bytes,4,opt,name=tour_uid,json=tourUid,proto3" json:"tour_uid,omitempty"`
	ScheduleUid   string               `protobuf:"bytes,5,opt,name=schedule_uid,json=scheduleUid,proto3" json:"schedule_uid,omitempty"`
	Amount        int32                `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Type          string               `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Message       string               `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
	TransactionId string               `protobuf:"bytes,9,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	UpdatedAt     *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *LogBalance) Reset() {
	*x = LogBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogBalance) ProtoMessage() {}

func (x *LogBalance) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogBalance.ProtoReflect.Descriptor instead.
func (*LogBalance) Descriptor() ([]byte, []int) {
	return file_delivery_grpc_wallet_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *LogBalance) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *LogBalance) GetBalanceUid() string {
	if x != nil {
		return x.BalanceUid
	}
	return ""
}

func (x *LogBalance) GetUserUid() string {
	if x != nil {
		return x.UserUid
	}
	return ""
}

func (x *LogBalance) GetTourUid() string {
	if x != nil {
		return x.TourUid
	}
	return ""
}

func (x *LogBalance) GetScheduleUid() string {
	if x != nil {
		return x.ScheduleUid
	}
	return ""
}

func (x *LogBalance) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *LogBalance) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LogBalance) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogBalance) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *LogBalance) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type LogsBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*LogBalance `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *LogsBalanceResponse) Reset() {
	*x = LogsBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsBalanceResponse) ProtoMessage() {}

func (x *LogsBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsBalanceResponse.ProtoReflect.Descriptor instead.
func (*LogsBalanceResponse) Descriptor() ([]byte, []int) {
	return file_delivery_grpc_wallet_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *LogsBalanceResponse) GetData() []*LogBalance {
	if x != nil {
		return x.Data
	}
	return nil
}

type LogBalancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TourId     string `protobuf:"bytes,2,opt,name=tour_id,json=tourId,proto3" json:"tour_id,omitempty"`
	ScheduleId string `protobuf:"bytes,3,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *LogBalancesRequest) Reset() {
	*x = LogBalancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogBalancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogBalancesRequest) ProtoMessage() {}

func (x *LogBalancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_delivery_grpc_wallet_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogBalancesRequest.ProtoReflect.Descriptor instead.
func (*LogBalancesRequest) Descriptor() ([]byte, []int) {
	return file_delivery_grpc_wallet_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *LogBalancesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LogBalancesRequest) GetTourId() string {
	if x != nil {
		return x.TourId
	}
	return ""
}

func (x *LogBalancesRequest) GetScheduleId() string {
	if x != nil {
		return x.ScheduleId
	}
	return ""
}

var File_delivery_grpc_wallet_wallet_proto protoreflect.FileDescriptor

var file_delivery_grpc_wallet_wallet_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x07, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x37, 0x0a, 0x08, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0x36, 0x0a, 0x0f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0xc0, 0x02, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x75, 0x72,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x75, 0x72,
	0x55, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3d,
	0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x67,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x67, 0x0a,
	0x12, 0x4c, 0x6f, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x6f, 0x75, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x6f, 0x75, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x32, 0xc6, 0x02, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x1a, 0x17, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d,
	0x12, 0x6d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_delivery_grpc_wallet_wallet_proto_rawDescOnce sync.Once
	file_delivery_grpc_wallet_wallet_proto_rawDescData = file_delivery_grpc_wallet_wallet_proto_rawDesc
)

func file_delivery_grpc_wallet_wallet_proto_rawDescGZIP() []byte {
	file_delivery_grpc_wallet_wallet_proto_rawDescOnce.Do(func() {
		file_delivery_grpc_wallet_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_delivery_grpc_wallet_wallet_proto_rawDescData)
	})
	return file_delivery_grpc_wallet_wallet_proto_rawDescData
}

var file_delivery_grpc_wallet_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_delivery_grpc_wallet_wallet_proto_goTypes = []interface{}{
	(*Balance)(nil),             // 0: wallet.Balance
	(*Balances)(nil),            // 1: wallet.Balances
	(*BalanceResponse)(nil),     // 2: wallet.BalanceResponse
	(*BalanceRequest)(nil),      // 3: wallet.BalanceRequest
	(*LogBalance)(nil),          // 4: wallet.LogBalance
	(*LogsBalanceResponse)(nil), // 5: wallet.LogsBalanceResponse
	(*LogBalancesRequest)(nil),  // 6: wallet.LogBalancesRequest
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_delivery_grpc_wallet_wallet_proto_depIdxs = []int32{
	7, // 0: wallet.Balance.updated_at:type_name -> google.protobuf.Timestamp
	0, // 1: wallet.Balances.balances:type_name -> wallet.Balance
	0, // 2: wallet.BalanceResponse.data:type_name -> wallet.Balance
	7, // 3: wallet.LogBalance.updated_at:type_name -> google.protobuf.Timestamp
	4, // 4: wallet.LogsBalanceResponse.data:type_name -> wallet.LogBalance
	4, // 5: wallet.WalletService.CreateOrAddBalance:input_type -> wallet.LogBalance
	3, // 6: wallet.WalletService.GetBalance:input_type -> wallet.BalanceRequest
	6, // 7: wallet.WalletService.GetLogBalances:input_type -> wallet.LogBalancesRequest
	2, // 8: wallet.WalletService.CreateOrAddBalance:output_type -> wallet.BalanceResponse
	2, // 9: wallet.WalletService.GetBalance:output_type -> wallet.BalanceResponse
	5, // 10: wallet.WalletService.GetLogBalances:output_type -> wallet.LogsBalanceResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_delivery_grpc_wallet_wallet_proto_init() }
func file_delivery_grpc_wallet_wallet_proto_init() {
	if File_delivery_grpc_wallet_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_delivery_grpc_wallet_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_delivery_grpc_wallet_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balances); i {
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
		file_delivery_grpc_wallet_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceResponse); i {
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
		file_delivery_grpc_wallet_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceRequest); i {
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
		file_delivery_grpc_wallet_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogBalance); i {
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
		file_delivery_grpc_wallet_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsBalanceResponse); i {
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
		file_delivery_grpc_wallet_wallet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogBalancesRequest); i {
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
			RawDescriptor: file_delivery_grpc_wallet_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_delivery_grpc_wallet_wallet_proto_goTypes,
		DependencyIndexes: file_delivery_grpc_wallet_wallet_proto_depIdxs,
		MessageInfos:      file_delivery_grpc_wallet_wallet_proto_msgTypes,
	}.Build()
	File_delivery_grpc_wallet_wallet_proto = out.File
	file_delivery_grpc_wallet_wallet_proto_rawDesc = nil
	file_delivery_grpc_wallet_wallet_proto_goTypes = nil
	file_delivery_grpc_wallet_wallet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletServiceClient interface {
	CreateOrAddBalance(ctx context.Context, in *LogBalance, opts ...grpc.CallOption) (*BalanceResponse, error)
	GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	GetLogBalances(ctx context.Context, in *LogBalancesRequest, opts ...grpc.CallOption) (*LogsBalanceResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) CreateOrAddBalance(ctx context.Context, in *LogBalance, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/CreateOrAddBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetLogBalances(ctx context.Context, in *LogBalancesRequest, opts ...grpc.CallOption) (*LogsBalanceResponse, error) {
	out := new(LogsBalanceResponse)
	err := c.cc.Invoke(ctx, "/wallet.WalletService/GetLogBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
type WalletServiceServer interface {
	CreateOrAddBalance(context.Context, *LogBalance) (*BalanceResponse, error)
	GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	GetLogBalances(context.Context, *LogBalancesRequest) (*LogsBalanceResponse, error)
}

// UnimplementedWalletServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (*UnimplementedWalletServiceServer) CreateOrAddBalance(context.Context, *LogBalance) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrAddBalance not implemented")
}
func (*UnimplementedWalletServiceServer) GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (*UnimplementedWalletServiceServer) GetLogBalances(context.Context, *LogBalancesRequest) (*LogsBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogBalances not implemented")
}

func RegisterWalletServiceServer(s *grpc.Server, srv WalletServiceServer) {
	s.RegisterService(&_WalletService_serviceDesc, srv)
}

func _WalletService_CreateOrAddBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogBalance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreateOrAddBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/CreateOrAddBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreateOrAddBalance(ctx, req.(*LogBalance))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetLogBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetLogBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.WalletService/GetLogBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetLogBalances(ctx, req.(*LogBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WalletService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrAddBalance",
			Handler:    _WalletService_CreateOrAddBalance_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _WalletService_GetBalance_Handler,
		},
		{
			MethodName: "GetLogBalances",
			Handler:    _WalletService_GetLogBalances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery/grpc/wallet/wallet.proto",
}