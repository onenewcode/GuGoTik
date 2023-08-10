// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: feed.proto

package feed

import (
	user "GuGoTik/src/rpc/user"
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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author        *user.User `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	PlayUrl       string     `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`
	CoverUrl      string     `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	FavoriteCount uint32     `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`
	CommentCount  uint32     `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	IsFavorite    bool       `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
	Title         string     `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *user.User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() uint32 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() uint32 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ListFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime *string `protobuf:"bytes,1,opt,name=latest_time,json=latestTime,proto3,oneof" json:"latest_time,omitempty"`
	ActorId    *uint32 `protobuf:"varint,2,opt,name=actor_id,json=actorId,proto3,oneof" json:"actor_id,omitempty"`
}

func (x *ListFeedRequest) Reset() {
	*x = ListFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedRequest) ProtoMessage() {}

func (x *ListFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedRequest.ProtoReflect.Descriptor instead.
func (*ListFeedRequest) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{1}
}

func (x *ListFeedRequest) GetLatestTime() string {
	if x != nil && x.LatestTime != nil {
		return *x.LatestTime
	}
	return ""
}

func (x *ListFeedRequest) GetActorId() uint32 {
	if x != nil && x.ActorId != nil {
		return *x.ActorId
	}
	return 0
}

type ListFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	NextTime   *uint32  `protobuf:"varint,3,opt,name=next_time,json=nextTime,proto3,oneof" json:"next_time,omitempty"`
	VideoList  []*Video `protobuf:"bytes,4,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *ListFeedResponse) Reset() {
	*x = ListFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeedResponse) ProtoMessage() {}

func (x *ListFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeedResponse.ProtoReflect.Descriptor instead.
func (*ListFeedResponse) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{2}
}

func (x *ListFeedResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListFeedResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ListFeedResponse) GetNextTime() uint32 {
	if x != nil && x.NextTime != nil {
		return *x.NextTime
	}
	return 0
}

func (x *ListFeedResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type QueryVideosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId  uint32   `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	VideoIds []uint32 `protobuf:"varint,2,rep,packed,name=video_ids,json=videoIds,proto3" json:"video_ids,omitempty"`
}

func (x *QueryVideosRequest) Reset() {
	*x = QueryVideosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryVideosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryVideosRequest) ProtoMessage() {}

func (x *QueryVideosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryVideosRequest.ProtoReflect.Descriptor instead.
func (*QueryVideosRequest) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{3}
}

func (x *QueryVideosRequest) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *QueryVideosRequest) GetVideoIds() []uint32 {
	if x != nil {
		return x.VideoIds
	}
	return nil
}

type QueryVideosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (x *QueryVideosResponse) Reset() {
	*x = QueryVideosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryVideosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryVideosResponse) ProtoMessage() {}

func (x *QueryVideosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryVideosResponse.ProtoReflect.Descriptor instead.
func (*QueryVideosResponse) Descriptor() ([]byte, []int) {
	return file_feed_proto_rawDescGZIP(), []int{4}
}

func (x *QueryVideosResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *QueryVideosResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *QueryVideosResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_feed_proto protoreflect.FileDescriptor

var file_feed_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x70,
	0x63, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69,
	0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22,
	0x74, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x22, 0xb2, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x0a,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x12, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x12, 0x2e, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x32, 0x9e, 0x01, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x19,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x16, 0x5a, 0x14, 0x47, 0x75, 0x47, 0x6f, 0x54, 0x69, 0x6b, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_feed_proto_rawDescOnce sync.Once
	file_feed_proto_rawDescData = file_feed_proto_rawDesc
)

func file_feed_proto_rawDescGZIP() []byte {
	file_feed_proto_rawDescOnce.Do(func() {
		file_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_feed_proto_rawDescData)
	})
	return file_feed_proto_rawDescData
}

var file_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_feed_proto_goTypes = []interface{}{
	(*Video)(nil),               // 0: rpc.feed.Video
	(*ListFeedRequest)(nil),     // 1: rpc.feed.ListFeedRequest
	(*ListFeedResponse)(nil),    // 2: rpc.feed.ListFeedResponse
	(*QueryVideosRequest)(nil),  // 3: rpc.feed.QueryVideosRequest
	(*QueryVideosResponse)(nil), // 4: rpc.feed.QueryVideosResponse
	(*user.User)(nil),           // 5: rpc.user.User
}
var file_feed_proto_depIdxs = []int32{
	5, // 0: rpc.feed.Video.author:type_name -> rpc.user.User
	0, // 1: rpc.feed.ListFeedResponse.video_list:type_name -> rpc.feed.Video
	0, // 2: rpc.feed.QueryVideosResponse.video_list:type_name -> rpc.feed.Video
	1, // 3: rpc.feed.FeedService.ListVideos:input_type -> rpc.feed.ListFeedRequest
	3, // 4: rpc.feed.FeedService.QueryVideos:input_type -> rpc.feed.QueryVideosRequest
	2, // 5: rpc.feed.FeedService.ListVideos:output_type -> rpc.feed.ListFeedResponse
	4, // 6: rpc.feed.FeedService.QueryVideos:output_type -> rpc.feed.QueryVideosResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_feed_proto_init() }
func file_feed_proto_init() {
	if File_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeedRequest); i {
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
		file_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeedResponse); i {
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
		file_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryVideosRequest); i {
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
		file_feed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryVideosResponse); i {
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
	file_feed_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_feed_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feed_proto_goTypes,
		DependencyIndexes: file_feed_proto_depIdxs,
		MessageInfos:      file_feed_proto_msgTypes,
	}.Build()
	File_feed_proto = out.File
	file_feed_proto_rawDesc = nil
	file_feed_proto_goTypes = nil
	file_feed_proto_depIdxs = nil
}
