// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: stc/auction.proto

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

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   int32  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	BidderId string `protobuf:"bytes,2,opt,name=bidder_id,json=bidderId,proto3" json:"bidder_id,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	mi := &file_stc_auction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stc_auction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_stc_auction_proto_rawDescGZIP(), []int{0}
}

func (x *BidRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BidRequest) GetBidderId() string {
	if x != nil {
		return x.BidderId
	}
	return ""
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	mi := &file_stc_auction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stc_auction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_stc_auction_proto_rawDescGZIP(), []int{1}
}

func (x *BidResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BidResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	mi := &file_stc_auction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stc_auction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_stc_auction_proto_rawDescGZIP(), []int{2}
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighestBid int32  `protobuf:"varint,1,opt,name=highest_bid,json=highestBid,proto3" json:"highest_bid,omitempty"`
	Winner     string `protobuf:"bytes,2,opt,name=winner,proto3" json:"winner,omitempty"`
	IsEnded    bool   `protobuf:"varint,3,opt,name=is_ended,json=isEnded,proto3" json:"is_ended,omitempty"`
	Error      string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	mi := &file_stc_auction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stc_auction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_stc_auction_proto_rawDescGZIP(), []int{3}
}

func (x *ResultResponse) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *ResultResponse) GetWinner() string {
	if x != nil {
		return x.Winner
	}
	return ""
}

func (x *ResultResponse) GetIsEnded() bool {
	if x != nil {
		return x.IsEnded
	}
	return false
}

func (x *ResultResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ReplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount    int32  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	BidderId  string `protobuf:"bytes,2,opt,name=bidder_id,json=bidderId,proto3" json:"bidder_id,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ReplicationRequest) Reset() {
	*x = ReplicationRequest{}
	mi := &file_stc_auction_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationRequest) ProtoMessage() {}

func (x *ReplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stc_auction_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationRequest.ProtoReflect.Descriptor instead.
func (*ReplicationRequest) Descriptor() ([]byte, []int) {
	return file_stc_auction_proto_rawDescGZIP(), []int{4}
}

func (x *ReplicationRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ReplicationRequest) GetBidderId() string {
	if x != nil {
		return x.BidderId
	}
	return ""
}

func (x *ReplicationRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ReplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ReplicationResponse) Reset() {
	*x = ReplicationResponse{}
	mi := &file_stc_auction_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationResponse) ProtoMessage() {}

func (x *ReplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stc_auction_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationResponse.ProtoReflect.Descriptor instead.
func (*ReplicationResponse) Descriptor() ([]byte, []int) {
	return file_stc_auction_proto_rawDescGZIP(), []int{5}
}

func (x *ReplicationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ReplicationResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr      string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	IsLeader  bool   `protobuf:"varint,2,opt,name=is_leader,json=isLeader,proto3" json:"is_leader,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	mi := &file_stc_auction_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stc_auction_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_stc_auction_proto_rawDescGZIP(), []int{6}
}

func (x *HeartbeatRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *HeartbeatRequest) GetIsLeader() bool {
	if x != nil {
		return x.IsLeader
	}
	return false
}

func (x *HeartbeatRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acknowledged bool `protobuf:"varint,1,opt,name=acknowledged,proto3" json:"acknowledged,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	mi := &file_stc_auction_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stc_auction_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_stc_auction_proto_rawDescGZIP(), []int{7}
}

func (x *HeartbeatResponse) GetAcknowledged() bool {
	if x != nil {
		return x.Acknowledged
	}
	return false
}

var File_stc_auction_proto protoreflect.FileDescriptor

var file_stc_auction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x63, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x0a,
	0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3d, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x0f,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x7a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x45, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x67, 0x0a, 0x12, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x45, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x61, 0x0a, 0x10, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x37,
	0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f,
	0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x32, 0x9c, 0x02, 0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x16, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x64, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x19, 0x2e,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stc_auction_proto_rawDescOnce sync.Once
	file_stc_auction_proto_rawDescData = file_stc_auction_proto_rawDesc
)

func file_stc_auction_proto_rawDescGZIP() []byte {
	file_stc_auction_proto_rawDescOnce.Do(func() {
		file_stc_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_stc_auction_proto_rawDescData)
	})
	return file_stc_auction_proto_rawDescData
}

var file_stc_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stc_auction_proto_goTypes = []any{
	(*BidRequest)(nil),          // 0: auction.BidRequest
	(*BidResponse)(nil),         // 1: auction.BidResponse
	(*ResultRequest)(nil),       // 2: auction.ResultRequest
	(*ResultResponse)(nil),      // 3: auction.ResultResponse
	(*ReplicationRequest)(nil),  // 4: auction.ReplicationRequest
	(*ReplicationResponse)(nil), // 5: auction.ReplicationResponse
	(*HeartbeatRequest)(nil),    // 6: auction.HeartbeatRequest
	(*HeartbeatResponse)(nil),   // 7: auction.HeartbeatResponse
}
var file_stc_auction_proto_depIdxs = []int32{
	0, // 0: auction.AuctionService.PlaceBid:input_type -> auction.BidRequest
	2, // 1: auction.AuctionService.GetResult:input_type -> auction.ResultRequest
	4, // 2: auction.AuctionService.ReplicateBid:input_type -> auction.ReplicationRequest
	6, // 3: auction.AuctionService.Heartbeat:input_type -> auction.HeartbeatRequest
	1, // 4: auction.AuctionService.PlaceBid:output_type -> auction.BidResponse
	3, // 5: auction.AuctionService.GetResult:output_type -> auction.ResultResponse
	5, // 6: auction.AuctionService.ReplicateBid:output_type -> auction.ReplicationResponse
	7, // 7: auction.AuctionService.Heartbeat:output_type -> auction.HeartbeatResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stc_auction_proto_init() }
func file_stc_auction_proto_init() {
	if File_stc_auction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stc_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stc_auction_proto_goTypes,
		DependencyIndexes: file_stc_auction_proto_depIdxs,
		MessageInfos:      file_stc_auction_proto_msgTypes,
	}.Build()
	File_stc_auction_proto = out.File
	file_stc_auction_proto_rawDesc = nil
	file_stc_auction_proto_goTypes = nil
	file_stc_auction_proto_depIdxs = nil
}
