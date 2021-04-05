// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/goreleases.proto

package proto

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

type GetReleaseInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetReleaseInfoRequest) Reset() {
	*x = GetReleaseInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goreleases_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReleaseInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReleaseInfoRequest) ProtoMessage() {}

func (x *GetReleaseInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goreleases_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReleaseInfoRequest.ProtoReflect.Descriptor instead.
func (*GetReleaseInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_goreleases_proto_rawDescGZIP(), []int{0}
}

func (x *GetReleaseInfoRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ListReleasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListReleasesRequest) Reset() {
	*x = ListReleasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goreleases_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReleasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReleasesRequest) ProtoMessage() {}

func (x *ListReleasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goreleases_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReleasesRequest.ProtoReflect.Descriptor instead.
func (*ListReleasesRequest) Descriptor() ([]byte, []int) {
	return file_proto_goreleases_proto_rawDescGZIP(), []int{1}
}

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goreleases_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goreleases_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_goreleases_proto_rawDescGZIP(), []int{2}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListReleasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Releases []*ReleaseInfo `protobuf:"bytes,1,rep,name=releases,proto3" json:"releases,omitempty"`
}

func (x *ListReleasesResponse) Reset() {
	*x = ListReleasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goreleases_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReleasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReleasesResponse) ProtoMessage() {}

func (x *ListReleasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goreleases_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReleasesResponse.ProtoReflect.Descriptor instead.
func (*ListReleasesResponse) Descriptor() ([]byte, []int) {
	return file_proto_goreleases_proto_rawDescGZIP(), []int{3}
}

func (x *ListReleasesResponse) GetReleases() []*ReleaseInfo {
	if x != nil {
		return x.Releases
	}
	return nil
}

type ReleaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version         string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ReleaseDate     string `protobuf:"bytes,2,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	ReleaseNotesUrl string `protobuf:"bytes,3,opt,name=release_notes_url,json=releaseNotesUrl,proto3" json:"release_notes_url,omitempty"`
}

func (x *ReleaseInfo) Reset() {
	*x = ReleaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goreleases_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseInfo) ProtoMessage() {}

func (x *ReleaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goreleases_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseInfo.ProtoReflect.Descriptor instead.
func (*ReleaseInfo) Descriptor() ([]byte, []int) {
	return file_proto_goreleases_proto_rawDescGZIP(), []int{4}
}

func (x *ReleaseInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ReleaseInfo) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *ReleaseInfo) GetReleaseNotesUrl() string {
	if x != nil {
		return x.ReleaseNotesUrl
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goreleases_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goreleases_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_goreleases_proto_rawDescGZIP(), []int{5}
}

func (x *SayHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type InsertUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname   string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Othername string `protobuf:"bytes,4,opt,name=othername,proto3" json:"othername,omitempty"`
}

func (x *InsertUserRequest) Reset() {
	*x = InsertUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goreleases_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserRequest) ProtoMessage() {}

func (x *InsertUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goreleases_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserRequest.ProtoReflect.Descriptor instead.
func (*InsertUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_goreleases_proto_rawDescGZIP(), []int{6}
}

func (x *InsertUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InsertUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsertUserRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *InsertUserRequest) GetOthername() string {
	if x != nil {
		return x.Othername
	}
	return ""
}

type InsertUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InsertUserResponse) Reset() {
	*x = InsertUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_goreleases_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserResponse) ProtoMessage() {}

func (x *InsertUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_goreleases_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserResponse.ProtoReflect.Descriptor instead.
func (*InsertUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_goreleases_proto_rawDescGZIP(), []int{7}
}

func (x *InsertUserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_goreleases_proto protoreflect.FileDescriptor

var file_proto_goreleases_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x31, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x46, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x22, 0x76, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x55, 0x72, 0x6c,
	0x22, 0x2c, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6f,
	0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x24, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa7, 0x02, 0x0a, 0x10, 0x47, 0x6f, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x73,
	0x73, 0x65, 0x72, 0x67, 0x65, 0x69, 0x2f, 0x42, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_goreleases_proto_rawDescOnce sync.Once
	file_proto_goreleases_proto_rawDescData = file_proto_goreleases_proto_rawDesc
)

func file_proto_goreleases_proto_rawDescGZIP() []byte {
	file_proto_goreleases_proto_rawDescOnce.Do(func() {
		file_proto_goreleases_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_goreleases_proto_rawDescData)
	})
	return file_proto_goreleases_proto_rawDescData
}

var file_proto_goreleases_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_goreleases_proto_goTypes = []interface{}{
	(*GetReleaseInfoRequest)(nil), // 0: proto.GetReleaseInfoRequest
	(*ListReleasesRequest)(nil),   // 1: proto.ListReleasesRequest
	(*SayHelloRequest)(nil),       // 2: proto.SayHelloRequest
	(*ListReleasesResponse)(nil),  // 3: proto.ListReleasesResponse
	(*ReleaseInfo)(nil),           // 4: proto.ReleaseInfo
	(*SayHelloResponse)(nil),      // 5: proto.SayHelloResponse
	(*InsertUserRequest)(nil),     // 6: proto.InsertUserRequest
	(*InsertUserResponse)(nil),    // 7: proto.InsertUserResponse
}
var file_proto_goreleases_proto_depIdxs = []int32{
	4, // 0: proto.ListReleasesResponse.releases:type_name -> proto.ReleaseInfo
	0, // 1: proto.GoReleaseService.GetReleaseInfo:input_type -> proto.GetReleaseInfoRequest
	1, // 2: proto.GoReleaseService.ListReleases:input_type -> proto.ListReleasesRequest
	2, // 3: proto.GoReleaseService.SayHello:input_type -> proto.SayHelloRequest
	6, // 4: proto.GoReleaseService.InsertUser:input_type -> proto.InsertUserRequest
	4, // 5: proto.GoReleaseService.GetReleaseInfo:output_type -> proto.ReleaseInfo
	3, // 6: proto.GoReleaseService.ListReleases:output_type -> proto.ListReleasesResponse
	5, // 7: proto.GoReleaseService.SayHello:output_type -> proto.SayHelloResponse
	7, // 8: proto.GoReleaseService.InsertUser:output_type -> proto.InsertUserResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_goreleases_proto_init() }
func file_proto_goreleases_proto_init() {
	if File_proto_goreleases_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_goreleases_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReleaseInfoRequest); i {
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
		file_proto_goreleases_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReleasesRequest); i {
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
		file_proto_goreleases_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_proto_goreleases_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReleasesResponse); i {
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
		file_proto_goreleases_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseInfo); i {
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
		file_proto_goreleases_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
		file_proto_goreleases_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUserRequest); i {
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
		file_proto_goreleases_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUserResponse); i {
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
			RawDescriptor: file_proto_goreleases_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_goreleases_proto_goTypes,
		DependencyIndexes: file_proto_goreleases_proto_depIdxs,
		MessageInfos:      file_proto_goreleases_proto_msgTypes,
	}.Build()
	File_proto_goreleases_proto = out.File
	file_proto_goreleases_proto_rawDesc = nil
	file_proto_goreleases_proto_goTypes = nil
	file_proto_goreleases_proto_depIdxs = nil
}
