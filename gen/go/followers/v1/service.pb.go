// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: followers/v1/service.proto

package followers

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Follower struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Follower) Reset() {
	*x = Follower{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Follower) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Follower) ProtoMessage() {}

func (x *Follower) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Follower.ProtoReflect.Descriptor instead.
func (*Follower) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *Follower) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Follower) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserToFollowId string `protobuf:"bytes,2,opt,name=user_to_follow_id,json=userToFollowId,proto3" json:"user_to_follow_id,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *FollowRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FollowRequest) GetUserToFollowId() string {
	if x != nil {
		return x.UserToFollowId
	}
	return ""
}

type UnFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UnFollowRequest) Reset() {
	*x = UnFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnFollowRequest) ProtoMessage() {}

func (x *UnFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnFollowRequest.ProtoReflect.Descriptor instead.
func (*UnFollowRequest) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *UnFollowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListFollowersRequest) Reset() {
	*x = ListFollowersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowersRequest) ProtoMessage() {}

func (x *ListFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowersRequest.ProtoReflect.Descriptor instead.
func (*ListFollowersRequest) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListFollowersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followers []*Follower `protobuf:"bytes,1,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *ListFollowersResponse) Reset() {
	*x = ListFollowersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowersResponse) ProtoMessage() {}

func (x *ListFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowersResponse.ProtoReflect.Descriptor instead.
func (*ListFollowersResponse) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListFollowersResponse) GetFollowers() []*Follower {
	if x != nil {
		return x.Followers
	}
	return nil
}

type ListFollowingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListFollowingsRequest) Reset() {
	*x = ListFollowingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowingsRequest) ProtoMessage() {}

func (x *ListFollowingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowingsRequest.ProtoReflect.Descriptor instead.
func (*ListFollowingsRequest) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListFollowingsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListFollowingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followings []*Follower `protobuf:"bytes,1,rep,name=followings,proto3" json:"followings,omitempty"`
}

func (x *ListFollowingsResponse) Reset() {
	*x = ListFollowingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowingsResponse) ProtoMessage() {}

func (x *ListFollowingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowingsResponse.ProtoReflect.Descriptor instead.
func (*ListFollowingsResponse) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListFollowingsResponse) GetFollowings() []*Follower {
	if x != nil {
		return x.Followings
	}
	return nil
}

type CountFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CountFollowersRequest) Reset() {
	*x = CountFollowersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFollowersRequest) ProtoMessage() {}

func (x *CountFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFollowersRequest.ProtoReflect.Descriptor instead.
func (*CountFollowersRequest) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *CountFollowersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CountFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountFollowersResponse) Reset() {
	*x = CountFollowersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFollowersResponse) ProtoMessage() {}

func (x *CountFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFollowersResponse.ProtoReflect.Descriptor instead.
func (*CountFollowersResponse) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *CountFollowersResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CountFollowingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CountFollowingsRequest) Reset() {
	*x = CountFollowingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFollowingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFollowingsRequest) ProtoMessage() {}

func (x *CountFollowingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFollowingsRequest.ProtoReflect.Descriptor instead.
func (*CountFollowingsRequest) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *CountFollowingsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CountFollowingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountFollowingsResponse) Reset() {
	*x = CountFollowingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFollowingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFollowingsResponse) ProtoMessage() {}

func (x *CountFollowingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFollowingsResponse.ProtoReflect.Descriptor instead.
func (*CountFollowingsResponse) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{10}
}

func (x *CountFollowingsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type IsFollowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserToFollowId string `protobuf:"bytes,2,opt,name=user_to_follow_id,json=userToFollowId,proto3" json:"user_to_follow_id,omitempty"`
}

func (x *IsFollowingRequest) Reset() {
	*x = IsFollowingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFollowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFollowingRequest) ProtoMessage() {}

func (x *IsFollowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFollowingRequest.ProtoReflect.Descriptor instead.
func (*IsFollowingRequest) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{11}
}

func (x *IsFollowingRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IsFollowingRequest) GetUserToFollowId() string {
	if x != nil {
		return x.UserToFollowId
	}
	return ""
}

type IsFollowingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFollow bool `protobuf:"varint,1,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
}

func (x *IsFollowingResponse) Reset() {
	*x = IsFollowingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_followers_v1_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFollowingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFollowingResponse) ProtoMessage() {}

func (x *IsFollowingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_followers_v1_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFollowingResponse.ProtoReflect.Descriptor instead.
func (*IsFollowingResponse) Descriptor() ([]byte, []int) {
	return file_followers_v1_service_proto_rawDescGZIP(), []int{12}
}

func (x *IsFollowingResponse) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

var File_followers_v1_service_proto protoreflect.FileDescriptor

var file_followers_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a,
	0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x74, 0x6f, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x49, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x55, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x0a, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x30, 0x0a, 0x15, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x16, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x16, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f,
	0x0a, 0x17, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x58, 0x0a, 0x12, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29,
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x49, 0x73, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x32, 0xb5, 0x06,
	0x0a, 0x10, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x12,
	0x6a, 0x0a, 0x08, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1d, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x2a, 0x1f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x73, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x12, 0x81, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x23, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x7c, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x54, 0x0a, 0x0b, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x20,
	0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x66, 0x69, 0x76, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_followers_v1_service_proto_rawDescOnce sync.Once
	file_followers_v1_service_proto_rawDescData = file_followers_v1_service_proto_rawDesc
)

func file_followers_v1_service_proto_rawDescGZIP() []byte {
	file_followers_v1_service_proto_rawDescOnce.Do(func() {
		file_followers_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_followers_v1_service_proto_rawDescData)
	})
	return file_followers_v1_service_proto_rawDescData
}

var file_followers_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_followers_v1_service_proto_goTypes = []interface{}{
	(*Follower)(nil),                // 0: followers.v1.Follower
	(*FollowRequest)(nil),           // 1: followers.v1.FollowRequest
	(*UnFollowRequest)(nil),         // 2: followers.v1.UnFollowRequest
	(*ListFollowersRequest)(nil),    // 3: followers.v1.ListFollowersRequest
	(*ListFollowersResponse)(nil),   // 4: followers.v1.ListFollowersResponse
	(*ListFollowingsRequest)(nil),   // 5: followers.v1.ListFollowingsRequest
	(*ListFollowingsResponse)(nil),  // 6: followers.v1.ListFollowingsResponse
	(*CountFollowersRequest)(nil),   // 7: followers.v1.CountFollowersRequest
	(*CountFollowersResponse)(nil),  // 8: followers.v1.CountFollowersResponse
	(*CountFollowingsRequest)(nil),  // 9: followers.v1.CountFollowingsRequest
	(*CountFollowingsResponse)(nil), // 10: followers.v1.CountFollowingsResponse
	(*IsFollowingRequest)(nil),      // 11: followers.v1.IsFollowingRequest
	(*IsFollowingResponse)(nil),     // 12: followers.v1.IsFollowingResponse
	(*emptypb.Empty)(nil),           // 13: google.protobuf.Empty
}
var file_followers_v1_service_proto_depIdxs = []int32{
	0,  // 0: followers.v1.ListFollowersResponse.followers:type_name -> followers.v1.Follower
	0,  // 1: followers.v1.ListFollowingsResponse.followings:type_name -> followers.v1.Follower
	1,  // 2: followers.v1.FollowersService.Follow:input_type -> followers.v1.FollowRequest
	2,  // 3: followers.v1.FollowersService.Unfollow:input_type -> followers.v1.UnFollowRequest
	3,  // 4: followers.v1.FollowersService.ListFollowers:input_type -> followers.v1.ListFollowersRequest
	5,  // 5: followers.v1.FollowersService.ListFollowings:input_type -> followers.v1.ListFollowingsRequest
	7,  // 6: followers.v1.FollowersService.CountFollowers:input_type -> followers.v1.CountFollowersRequest
	9,  // 7: followers.v1.FollowersService.CountFollowings:input_type -> followers.v1.CountFollowingsRequest
	11, // 8: followers.v1.FollowersService.IsFollowing:input_type -> followers.v1.IsFollowingRequest
	13, // 9: followers.v1.FollowersService.Follow:output_type -> google.protobuf.Empty
	13, // 10: followers.v1.FollowersService.Unfollow:output_type -> google.protobuf.Empty
	4,  // 11: followers.v1.FollowersService.ListFollowers:output_type -> followers.v1.ListFollowersResponse
	6,  // 12: followers.v1.FollowersService.ListFollowings:output_type -> followers.v1.ListFollowingsResponse
	8,  // 13: followers.v1.FollowersService.CountFollowers:output_type -> followers.v1.CountFollowersResponse
	10, // 14: followers.v1.FollowersService.CountFollowings:output_type -> followers.v1.CountFollowingsResponse
	12, // 15: followers.v1.FollowersService.IsFollowing:output_type -> followers.v1.IsFollowingResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_followers_v1_service_proto_init() }
func file_followers_v1_service_proto_init() {
	if File_followers_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_followers_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Follower); i {
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
		file_followers_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_followers_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnFollowRequest); i {
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
		file_followers_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowersRequest); i {
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
		file_followers_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowersResponse); i {
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
		file_followers_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowingsRequest); i {
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
		file_followers_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowingsResponse); i {
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
		file_followers_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFollowersRequest); i {
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
		file_followers_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFollowersResponse); i {
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
		file_followers_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFollowingsRequest); i {
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
		file_followers_v1_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFollowingsResponse); i {
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
		file_followers_v1_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFollowingRequest); i {
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
		file_followers_v1_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFollowingResponse); i {
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
			RawDescriptor: file_followers_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_followers_v1_service_proto_goTypes,
		DependencyIndexes: file_followers_v1_service_proto_depIdxs,
		MessageInfos:      file_followers_v1_service_proto_msgTypes,
	}.Build()
	File_followers_v1_service_proto = out.File
	file_followers_v1_service_proto_rawDesc = nil
	file_followers_v1_service_proto_goTypes = nil
	file_followers_v1_service_proto_depIdxs = nil
}
