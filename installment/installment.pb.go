// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: installment/installment.proto

package installment

import (
	pagination "github.com/djoonta/kpmasterproto/pagination"
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

type InstallmentWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Messsage string `protobuf:"bytes,2,opt,name=messsage,proto3" json:"messsage,omitempty"`
}

func (x *InstallmentWriteResponse) Reset() {
	*x = InstallmentWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentWriteResponse) ProtoMessage() {}

func (x *InstallmentWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentWriteResponse.ProtoReflect.Descriptor instead.
func (*InstallmentWriteResponse) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{0}
}

func (x *InstallmentWriteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *InstallmentWriteResponse) GetMesssage() string {
	if x != nil {
		return x.Messsage
	}
	return ""
}

type InstallmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *InstallmentInfo) Reset() {
	*x = InstallmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentInfo) ProtoMessage() {}

func (x *InstallmentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentInfo.ProtoReflect.Descriptor instead.
func (*InstallmentInfo) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{1}
}

func (x *InstallmentInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InstallmentInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstallmentInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InstallmentInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type InstallmentFindIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InstallmentFindIDRequest) Reset() {
	*x = InstallmentFindIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentFindIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentFindIDRequest) ProtoMessage() {}

func (x *InstallmentFindIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentFindIDRequest.ProtoReflect.Descriptor instead.
func (*InstallmentFindIDRequest) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{2}
}

func (x *InstallmentFindIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InstallmentFindIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installment *InstallmentInfo `protobuf:"bytes,1,opt,name=installment,proto3" json:"installment,omitempty"`
}

func (x *InstallmentFindIDResponse) Reset() {
	*x = InstallmentFindIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentFindIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentFindIDResponse) ProtoMessage() {}

func (x *InstallmentFindIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentFindIDResponse.ProtoReflect.Descriptor instead.
func (*InstallmentFindIDResponse) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{3}
}

func (x *InstallmentFindIDResponse) GetInstallment() *InstallmentInfo {
	if x != nil {
		return x.Installment
	}
	return nil
}

type InstallmentFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  string `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit string `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *InstallmentFindAllRequest) Reset() {
	*x = InstallmentFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentFindAllRequest) ProtoMessage() {}

func (x *InstallmentFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentFindAllRequest.ProtoReflect.Descriptor instead.
func (*InstallmentFindAllRequest) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{4}
}

func (x *InstallmentFindAllRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *InstallmentFindAllRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type InstallmentFindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installments []*InstallmentInfo           `protobuf:"bytes,1,rep,name=installments,proto3" json:"installments,omitempty"`
	Paginations  *pagination.PaginationMaster `protobuf:"bytes,2,opt,name=paginations,proto3" json:"paginations,omitempty"`
}

func (x *InstallmentFindAllResponse) Reset() {
	*x = InstallmentFindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentFindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentFindAllResponse) ProtoMessage() {}

func (x *InstallmentFindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentFindAllResponse.ProtoReflect.Descriptor instead.
func (*InstallmentFindAllResponse) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{5}
}

func (x *InstallmentFindAllResponse) GetInstallments() []*InstallmentInfo {
	if x != nil {
		return x.Installments
	}
	return nil
}

func (x *InstallmentFindAllResponse) GetPaginations() *pagination.PaginationMaster {
	if x != nil {
		return x.Paginations
	}
	return nil
}

type InstallmentCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InstallmentCreateRequest) Reset() {
	*x = InstallmentCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentCreateRequest) ProtoMessage() {}

func (x *InstallmentCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentCreateRequest.ProtoReflect.Descriptor instead.
func (*InstallmentCreateRequest) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{6}
}

func (x *InstallmentCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstallmentCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type InstallmentCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateResponse *InstallmentWriteResponse `protobuf:"bytes,1,opt,name=create_response,json=createResponse,proto3" json:"create_response,omitempty"`
}

func (x *InstallmentCreateResponse) Reset() {
	*x = InstallmentCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentCreateResponse) ProtoMessage() {}

func (x *InstallmentCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentCreateResponse.ProtoReflect.Descriptor instead.
func (*InstallmentCreateResponse) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{7}
}

func (x *InstallmentCreateResponse) GetCreateResponse() *InstallmentWriteResponse {
	if x != nil {
		return x.CreateResponse
	}
	return nil
}

type InstallmentDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InstallmentDeleteRequest) Reset() {
	*x = InstallmentDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentDeleteRequest) ProtoMessage() {}

func (x *InstallmentDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentDeleteRequest.ProtoReflect.Descriptor instead.
func (*InstallmentDeleteRequest) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{8}
}

func (x *InstallmentDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InstallmentDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeleteResponse *InstallmentWriteResponse `protobuf:"bytes,1,opt,name=delete_response,json=deleteResponse,proto3" json:"delete_response,omitempty"`
}

func (x *InstallmentDeleteResponse) Reset() {
	*x = InstallmentDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentDeleteResponse) ProtoMessage() {}

func (x *InstallmentDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentDeleteResponse.ProtoReflect.Descriptor instead.
func (*InstallmentDeleteResponse) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{9}
}

func (x *InstallmentDeleteResponse) GetDeleteResponse() *InstallmentWriteResponse {
	if x != nil {
		return x.DeleteResponse
	}
	return nil
}

type InstallmentUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InstallmentUpdateRequest) Reset() {
	*x = InstallmentUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentUpdateRequest) ProtoMessage() {}

func (x *InstallmentUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentUpdateRequest.ProtoReflect.Descriptor instead.
func (*InstallmentUpdateRequest) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{10}
}

func (x *InstallmentUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InstallmentUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstallmentUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type InstallmentUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateResponse *InstallmentWriteResponse `protobuf:"bytes,1,opt,name=update_response,json=updateResponse,proto3" json:"update_response,omitempty"`
}

func (x *InstallmentUpdateResponse) Reset() {
	*x = InstallmentUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_installment_installment_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallmentUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallmentUpdateResponse) ProtoMessage() {}

func (x *InstallmentUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_installment_installment_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallmentUpdateResponse.ProtoReflect.Descriptor instead.
func (*InstallmentUpdateResponse) Descriptor() ([]byte, []int) {
	return file_installment_installment_proto_rawDescGZIP(), []int{11}
}

func (x *InstallmentUpdateResponse) GetUpdateResponse() *InstallmentWriteResponse {
	if x != nil {
		return x.UpdateResponse
	}
	return nil
}

var File_installment_installment_proto protoreflect.FileDescriptor

var file_installment_installment_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x18,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x92,
	0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x2a, 0x0a, 0x18, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5d, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69,
	0x6e, 0x64, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x45,
	0x0a, 0x19, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x1a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x70, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x50, 0x0a, 0x18, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a,
	0x19, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x18,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x18, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x19, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x6b, 0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x6f, 0x6f, 0x6e, 0x74, 0x61, 0x2f, 0x6b,
	0x70, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_installment_installment_proto_rawDescOnce sync.Once
	file_installment_installment_proto_rawDescData = file_installment_installment_proto_rawDesc
)

func file_installment_installment_proto_rawDescGZIP() []byte {
	file_installment_installment_proto_rawDescOnce.Do(func() {
		file_installment_installment_proto_rawDescData = protoimpl.X.CompressGZIP(file_installment_installment_proto_rawDescData)
	})
	return file_installment_installment_proto_rawDescData
}

var file_installment_installment_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_installment_installment_proto_goTypes = []interface{}{
	(*InstallmentWriteResponse)(nil),    // 0: kpmasterproto.InstallmentWriteResponse
	(*InstallmentInfo)(nil),             // 1: kpmasterproto.InstallmentInfo
	(*InstallmentFindIDRequest)(nil),    // 2: kpmasterproto.InstallmentFindIDRequest
	(*InstallmentFindIDResponse)(nil),   // 3: kpmasterproto.InstallmentFindIDResponse
	(*InstallmentFindAllRequest)(nil),   // 4: kpmasterproto.InstallmentFindAllRequest
	(*InstallmentFindAllResponse)(nil),  // 5: kpmasterproto.InstallmentFindAllResponse
	(*InstallmentCreateRequest)(nil),    // 6: kpmasterproto.InstallmentCreateRequest
	(*InstallmentCreateResponse)(nil),   // 7: kpmasterproto.InstallmentCreateResponse
	(*InstallmentDeleteRequest)(nil),    // 8: kpmasterproto.InstallmentDeleteRequest
	(*InstallmentDeleteResponse)(nil),   // 9: kpmasterproto.InstallmentDeleteResponse
	(*InstallmentUpdateRequest)(nil),    // 10: kpmasterproto.InstallmentUpdateRequest
	(*InstallmentUpdateResponse)(nil),   // 11: kpmasterproto.InstallmentUpdateResponse
	(*timestamppb.Timestamp)(nil),       // 12: google.protobuf.Timestamp
	(*pagination.PaginationMaster)(nil), // 13: kpmasterproto.PaginationMaster
}
var file_installment_installment_proto_depIdxs = []int32{
	12, // 0: kpmasterproto.InstallmentInfo.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 1: kpmasterproto.InstallmentFindIDResponse.installment:type_name -> kpmasterproto.InstallmentInfo
	1,  // 2: kpmasterproto.InstallmentFindAllResponse.installments:type_name -> kpmasterproto.InstallmentInfo
	13, // 3: kpmasterproto.InstallmentFindAllResponse.paginations:type_name -> kpmasterproto.PaginationMaster
	0,  // 4: kpmasterproto.InstallmentCreateResponse.create_response:type_name -> kpmasterproto.InstallmentWriteResponse
	0,  // 5: kpmasterproto.InstallmentDeleteResponse.delete_response:type_name -> kpmasterproto.InstallmentWriteResponse
	0,  // 6: kpmasterproto.InstallmentUpdateResponse.update_response:type_name -> kpmasterproto.InstallmentWriteResponse
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_installment_installment_proto_init() }
func file_installment_installment_proto_init() {
	if File_installment_installment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_installment_installment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentWriteResponse); i {
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
		file_installment_installment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentInfo); i {
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
		file_installment_installment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentFindIDRequest); i {
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
		file_installment_installment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentFindIDResponse); i {
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
		file_installment_installment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentFindAllRequest); i {
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
		file_installment_installment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentFindAllResponse); i {
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
		file_installment_installment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentCreateRequest); i {
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
		file_installment_installment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentCreateResponse); i {
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
		file_installment_installment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentDeleteRequest); i {
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
		file_installment_installment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentDeleteResponse); i {
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
		file_installment_installment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentUpdateRequest); i {
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
		file_installment_installment_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallmentUpdateResponse); i {
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
			RawDescriptor: file_installment_installment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_installment_installment_proto_goTypes,
		DependencyIndexes: file_installment_installment_proto_depIdxs,
		MessageInfos:      file_installment_installment_proto_msgTypes,
	}.Build()
	File_installment_installment_proto = out.File
	file_installment_installment_proto_rawDesc = nil
	file_installment_installment_proto_goTypes = nil
	file_installment_installment_proto_depIdxs = nil
}
