// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: pb/fee/fee.proto

package fee

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Fee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeeId      string                 `protobuf:"bytes,1,opt,name=fee_id,json=feeId,proto3" json:"fee_id,omitempty"`
	FeeName    string                 `protobuf:"bytes,2,opt,name=fee_name,json=feeName,proto3" json:"fee_name,omitempty"`
	Nominal    float64                `protobuf:"fixed64,3,opt,name=nominal,proto3" json:"nominal,omitempty"`
	Percentage float64                `protobuf:"fixed64,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	ChargeTo   string                 `protobuf:"bytes,5,opt,name=charge_to,json=chargeTo,proto3" json:"charge_to,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedBy  *AdminInfo             `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
}

func (x *Fee) Reset() {
	*x = Fee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fee_fee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fee) ProtoMessage() {}

func (x *Fee) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fee_fee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fee.ProtoReflect.Descriptor instead.
func (*Fee) Descriptor() ([]byte, []int) {
	return file_pb_fee_fee_proto_rawDescGZIP(), []int{0}
}

func (x *Fee) GetFeeId() string {
	if x != nil {
		return x.FeeId
	}
	return ""
}

func (x *Fee) GetFeeName() string {
	if x != nil {
		return x.FeeName
	}
	return ""
}

func (x *Fee) GetNominal() float64 {
	if x != nil {
		return x.Nominal
	}
	return 0
}

func (x *Fee) GetPercentage() float64 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *Fee) GetChargeTo() string {
	if x != nil {
		return x.ChargeTo
	}
	return ""
}

func (x *Fee) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Fee) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Fee) GetCreatedBy() *AdminInfo {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

type AdminInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUid string `protobuf:"bytes,1,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AdminInfo) Reset() {
	*x = AdminInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fee_fee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminInfo) ProtoMessage() {}

func (x *AdminInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fee_fee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminInfo.ProtoReflect.Descriptor instead.
func (*AdminInfo) Descriptor() ([]byte, []int) {
	return file_pb_fee_fee_proto_rawDescGZIP(), []int{1}
}

func (x *AdminInfo) GetUserUid() string {
	if x != nil {
		return x.UserUid
	}
	return ""
}

func (x *AdminInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code  int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Data  *Fee   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FeeResponse) Reset() {
	*x = FeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fee_fee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeResponse) ProtoMessage() {}

func (x *FeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fee_fee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeResponse.ProtoReflect.Descriptor instead.
func (*FeeResponse) Descriptor() ([]byte, []int) {
	return file_pb_fee_fee_proto_rawDescGZIP(), []int{2}
}

func (x *FeeResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FeeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *FeeResponse) GetData() *Fee {
	if x != nil {
		return x.Data
	}
	return nil
}

type RequestFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param  string `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	Key    string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RequestFee) Reset() {
	*x = RequestFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fee_fee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFee) ProtoMessage() {}

func (x *RequestFee) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fee_fee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFee.ProtoReflect.Descriptor instead.
func (*RequestFee) Descriptor() ([]byte, []int) {
	return file_pb_fee_fee_proto_rawDescGZIP(), []int{3}
}

func (x *RequestFee) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *RequestFee) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *RequestFee) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ResponseFeeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code  int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Data  []*Fee `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseFeeList) Reset() {
	*x = ResponseFeeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fee_fee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFeeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFeeList) ProtoMessage() {}

func (x *ResponseFeeList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fee_fee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFeeList.ProtoReflect.Descriptor instead.
func (*ResponseFeeList) Descriptor() ([]byte, []int) {
	return file_pb_fee_fee_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseFeeList) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResponseFeeList) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseFeeList) GetData() []*Fee {
	if x != nil {
		return x.Data
	}
	return nil
}

type FeeDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code  int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *FeeDeleteResponse) Reset() {
	*x = FeeDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fee_fee_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeeDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeeDeleteResponse) ProtoMessage() {}

func (x *FeeDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fee_fee_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeeDeleteResponse.ProtoReflect.Descriptor instead.
func (*FeeDeleteResponse) Descriptor() ([]byte, []int) {
	return file_pb_fee_fee_proto_rawDescGZIP(), []int{5}
}

func (x *FeeDeleteResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FeeDeleteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_pb_fee_fee_proto protoreflect.FileDescriptor

var file_pb_fee_fee_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x66, 0x65, 0x65, 0x2f, 0x66, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x66, 0x65, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x66, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x65, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x54, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x3a, 0x0a, 0x09,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x4c, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x59, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x65, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x46,
	0x65, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x11, 0x46, 0x65, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xe8, 0x03, 0x0a, 0x0a, 0x46, 0x65, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x46, 0x65, 0x65,
	0x12, 0x08, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x46, 0x65, 0x65, 0x1a, 0x10, 0x2e, 0x66, 0x65, 0x65,
	0x2e, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65,
	0x65, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x46, 0x65, 0x65, 0x12,
	0x08, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x46, 0x65, 0x65, 0x1a, 0x10, 0x2e, 0x66, 0x65, 0x65, 0x2e,
	0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x32, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65,
	0x73, 0x2f, 0x7b, 0x66, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x09,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x46, 0x65, 0x65, 0x12, 0x08, 0x2e, 0x66, 0x65, 0x65, 0x2e,
	0x46, 0x65, 0x65, 0x1a, 0x10, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x5b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x65,
	0x12, 0x08, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x46, 0x65, 0x65, 0x1a, 0x16, 0x2e, 0x66, 0x65, 0x65,
	0x2e, 0x46, 0x65, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x2a, 0x24, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x65, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x5a, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x12, 0x0f, 0x2e, 0x66, 0x65,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x65, 0x1a, 0x14, 0x2e, 0x66,
	0x65, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x65, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x12, 0x08, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x46, 0x65, 0x65,
	0x1a, 0x10, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_fee_fee_proto_rawDescOnce sync.Once
	file_pb_fee_fee_proto_rawDescData = file_pb_fee_fee_proto_rawDesc
)

func file_pb_fee_fee_proto_rawDescGZIP() []byte {
	file_pb_fee_fee_proto_rawDescOnce.Do(func() {
		file_pb_fee_fee_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_fee_fee_proto_rawDescData)
	})
	return file_pb_fee_fee_proto_rawDescData
}

var file_pb_fee_fee_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_fee_fee_proto_goTypes = []interface{}{
	(*Fee)(nil),                   // 0: fee.Fee
	(*AdminInfo)(nil),             // 1: fee.AdminInfo
	(*FeeResponse)(nil),           // 2: fee.FeeResponse
	(*RequestFee)(nil),            // 3: fee.RequestFee
	(*ResponseFeeList)(nil),       // 4: fee.ResponseFeeList
	(*FeeDeleteResponse)(nil),     // 5: fee.FeeDeleteResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_pb_fee_fee_proto_depIdxs = []int32{
	6,  // 0: fee.Fee.created_at:type_name -> google.protobuf.Timestamp
	6,  // 1: fee.Fee.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: fee.Fee.created_by:type_name -> fee.AdminInfo
	0,  // 3: fee.FeeResponse.data:type_name -> fee.Fee
	0,  // 4: fee.ResponseFeeList.data:type_name -> fee.Fee
	0,  // 5: fee.FeeService.AddFee:input_type -> fee.Fee
	0,  // 6: fee.FeeService.EditFee:input_type -> fee.Fee
	0,  // 7: fee.FeeService.DetailFee:input_type -> fee.Fee
	0,  // 8: fee.FeeService.DeleteFee:input_type -> fee.Fee
	3,  // 9: fee.FeeService.ListFee:input_type -> fee.RequestFee
	0,  // 10: fee.FeeService.GetFee:input_type -> fee.Fee
	2,  // 11: fee.FeeService.AddFee:output_type -> fee.FeeResponse
	2,  // 12: fee.FeeService.EditFee:output_type -> fee.FeeResponse
	2,  // 13: fee.FeeService.DetailFee:output_type -> fee.FeeResponse
	5,  // 14: fee.FeeService.DeleteFee:output_type -> fee.FeeDeleteResponse
	4,  // 15: fee.FeeService.ListFee:output_type -> fee.ResponseFeeList
	2,  // 16: fee.FeeService.GetFee:output_type -> fee.FeeResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pb_fee_fee_proto_init() }
func file_pb_fee_fee_proto_init() {
	if File_pb_fee_fee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_fee_fee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fee); i {
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
		file_pb_fee_fee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminInfo); i {
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
		file_pb_fee_fee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeResponse); i {
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
		file_pb_fee_fee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFee); i {
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
		file_pb_fee_fee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFeeList); i {
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
		file_pb_fee_fee_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeeDeleteResponse); i {
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
			RawDescriptor: file_pb_fee_fee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_fee_fee_proto_goTypes,
		DependencyIndexes: file_pb_fee_fee_proto_depIdxs,
		MessageInfos:      file_pb_fee_fee_proto_msgTypes,
	}.Build()
	File_pb_fee_fee_proto = out.File
	file_pb_fee_fee_proto_rawDesc = nil
	file_pb_fee_fee_proto_goTypes = nil
	file_pb_fee_fee_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeeServiceClient is the client API for FeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeeServiceClient interface {
	AddFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeResponse, error)
	EditFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeResponse, error)
	DetailFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeResponse, error)
	DeleteFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeDeleteResponse, error)
	ListFee(ctx context.Context, in *RequestFee, opts ...grpc.CallOption) (*ResponseFeeList, error)
	//this endpoint only use fot get fees by fee names
	GetFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeResponse, error)
}

type feeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeeServiceClient(cc grpc.ClientConnInterface) FeeServiceClient {
	return &feeServiceClient{cc}
}

func (c *feeServiceClient) AddFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeResponse, error) {
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, "/fee.FeeService/AddFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feeServiceClient) EditFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeResponse, error) {
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, "/fee.FeeService/EditFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feeServiceClient) DetailFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeResponse, error) {
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, "/fee.FeeService/DetailFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feeServiceClient) DeleteFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeDeleteResponse, error) {
	out := new(FeeDeleteResponse)
	err := c.cc.Invoke(ctx, "/fee.FeeService/DeleteFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feeServiceClient) ListFee(ctx context.Context, in *RequestFee, opts ...grpc.CallOption) (*ResponseFeeList, error) {
	out := new(ResponseFeeList)
	err := c.cc.Invoke(ctx, "/fee.FeeService/ListFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feeServiceClient) GetFee(ctx context.Context, in *Fee, opts ...grpc.CallOption) (*FeeResponse, error) {
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, "/fee.FeeService/GetFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeeServiceServer is the server API for FeeService service.
type FeeServiceServer interface {
	AddFee(context.Context, *Fee) (*FeeResponse, error)
	EditFee(context.Context, *Fee) (*FeeResponse, error)
	DetailFee(context.Context, *Fee) (*FeeResponse, error)
	DeleteFee(context.Context, *Fee) (*FeeDeleteResponse, error)
	ListFee(context.Context, *RequestFee) (*ResponseFeeList, error)
	//this endpoint only use fot get fees by fee names
	GetFee(context.Context, *Fee) (*FeeResponse, error)
}

// UnimplementedFeeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeeServiceServer struct {
}

func (*UnimplementedFeeServiceServer) AddFee(context.Context, *Fee) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFee not implemented")
}
func (*UnimplementedFeeServiceServer) EditFee(context.Context, *Fee) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditFee not implemented")
}
func (*UnimplementedFeeServiceServer) DetailFee(context.Context, *Fee) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailFee not implemented")
}
func (*UnimplementedFeeServiceServer) DeleteFee(context.Context, *Fee) (*FeeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFee not implemented")
}
func (*UnimplementedFeeServiceServer) ListFee(context.Context, *RequestFee) (*ResponseFeeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFee not implemented")
}
func (*UnimplementedFeeServiceServer) GetFee(context.Context, *Fee) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFee not implemented")
}

func RegisterFeeServiceServer(s *grpc.Server, srv FeeServiceServer) {
	s.RegisterService(&_FeeService_serviceDesc, srv)
}

func _FeeService_AddFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).AddFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fee.FeeService/AddFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).AddFee(ctx, req.(*Fee))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeeService_EditFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).EditFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fee.FeeService/EditFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).EditFee(ctx, req.(*Fee))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeeService_DetailFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).DetailFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fee.FeeService/DetailFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).DetailFee(ctx, req.(*Fee))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeeService_DeleteFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).DeleteFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fee.FeeService/DeleteFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).DeleteFee(ctx, req.(*Fee))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeeService_ListFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).ListFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fee.FeeService/ListFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).ListFee(ctx, req.(*RequestFee))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeeService_GetFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeeServiceServer).GetFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fee.FeeService/GetFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeeServiceServer).GetFee(ctx, req.(*Fee))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fee.FeeService",
	HandlerType: (*FeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFee",
			Handler:    _FeeService_AddFee_Handler,
		},
		{
			MethodName: "EditFee",
			Handler:    _FeeService_EditFee_Handler,
		},
		{
			MethodName: "DetailFee",
			Handler:    _FeeService_DetailFee_Handler,
		},
		{
			MethodName: "DeleteFee",
			Handler:    _FeeService_DeleteFee_Handler,
		},
		{
			MethodName: "ListFee",
			Handler:    _FeeService_ListFee_Handler,
		},
		{
			MethodName: "GetFee",
			Handler:    _FeeService_GetFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/fee/fee.proto",
}
