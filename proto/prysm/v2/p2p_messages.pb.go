// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: proto/prysm/v2/p2p_messages.proto

package v2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	github_com_prysmaticlabs_eth2_types "github.com/prysmaticlabs/eth2-types"
	github_com_prysmaticlabs_go_bitfield "github.com/prysmaticlabs/go-bitfield"
	_ "github.com/prysmaticlabs/prysm/proto/eth/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForkDigest     []byte                                    `protobuf:"bytes,1,opt,name=fork_digest,json=forkDigest,proto3" json:"fork_digest,omitempty" ssz-size:"4"`
	FinalizedRoot  []byte                                    `protobuf:"bytes,2,opt,name=finalized_root,json=finalizedRoot,proto3" json:"finalized_root,omitempty" ssz-size:"32"`
	FinalizedEpoch github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,3,opt,name=finalized_epoch,json=finalizedEpoch,proto3" json:"finalized_epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
	HeadRoot       []byte                                    `protobuf:"bytes,4,opt,name=head_root,json=headRoot,proto3" json:"head_root,omitempty" ssz-size:"32"`
	HeadSlot       github_com_prysmaticlabs_eth2_types.Slot  `protobuf:"varint,5,opt,name=head_slot,json=headSlot,proto3" json:"head_slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_prysm_v2_p2p_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetForkDigest() []byte {
	if x != nil {
		return x.ForkDigest
	}
	return nil
}

func (x *Status) GetFinalizedRoot() []byte {
	if x != nil {
		return x.FinalizedRoot
	}
	return nil
}

func (x *Status) GetFinalizedEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.FinalizedEpoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

func (x *Status) GetHeadRoot() []byte {
	if x != nil {
		return x.HeadRoot
	}
	return nil
}

func (x *Status) GetHeadSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.HeadSlot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

type BeaconBlocksByRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartSlot github_com_prysmaticlabs_eth2_types.Slot `protobuf:"varint,1,opt,name=start_slot,json=startSlot,proto3" json:"start_slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	Count     uint64                                   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Step      uint64                                   `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *BeaconBlocksByRangeRequest) Reset() {
	*x = BeaconBlocksByRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconBlocksByRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconBlocksByRangeRequest) ProtoMessage() {}

func (x *BeaconBlocksByRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconBlocksByRangeRequest.ProtoReflect.Descriptor instead.
func (*BeaconBlocksByRangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_prysm_v2_p2p_messages_proto_rawDescGZIP(), []int{1}
}

func (x *BeaconBlocksByRangeRequest) GetStartSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.StartSlot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *BeaconBlocksByRangeRequest) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *BeaconBlocksByRangeRequest) GetStep() uint64 {
	if x != nil {
		return x.Step
	}
	return 0
}

type ENRForkID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentForkDigest []byte                                    `protobuf:"bytes,1,opt,name=current_fork_digest,json=currentForkDigest,proto3" json:"current_fork_digest,omitempty" ssz-size:"4"`
	NextForkVersion   []byte                                    `protobuf:"bytes,2,opt,name=next_fork_version,json=nextForkVersion,proto3" json:"next_fork_version,omitempty" ssz-size:"4"`
	NextForkEpoch     github_com_prysmaticlabs_eth2_types.Epoch `protobuf:"varint,3,opt,name=next_fork_epoch,json=nextForkEpoch,proto3" json:"next_fork_epoch,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Epoch"`
}

func (x *ENRForkID) Reset() {
	*x = ENRForkID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ENRForkID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ENRForkID) ProtoMessage() {}

func (x *ENRForkID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ENRForkID.ProtoReflect.Descriptor instead.
func (*ENRForkID) Descriptor() ([]byte, []int) {
	return file_proto_prysm_v2_p2p_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ENRForkID) GetCurrentForkDigest() []byte {
	if x != nil {
		return x.CurrentForkDigest
	}
	return nil
}

func (x *ENRForkID) GetNextForkVersion() []byte {
	if x != nil {
		return x.NextForkVersion
	}
	return nil
}

func (x *ENRForkID) GetNextForkEpoch() github_com_prysmaticlabs_eth2_types.Epoch {
	if x != nil {
		return x.NextForkEpoch
	}
	return github_com_prysmaticlabs_eth2_types.Epoch(0)
}

type MetaDataV0 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNumber uint64                                           `protobuf:"varint,1,opt,name=seq_number,json=seqNumber,proto3" json:"seq_number,omitempty"`
	Attnets   github_com_prysmaticlabs_go_bitfield.Bitvector64 `protobuf:"bytes,2,opt,name=attnets,proto3" json:"attnets,omitempty" cast-type:"github.com/prysmaticlabs/go-bitfield.Bitvector64" ssz-size:"8"`
}

func (x *MetaDataV0) Reset() {
	*x = MetaDataV0{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaDataV0) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaDataV0) ProtoMessage() {}

func (x *MetaDataV0) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaDataV0.ProtoReflect.Descriptor instead.
func (*MetaDataV0) Descriptor() ([]byte, []int) {
	return file_proto_prysm_v2_p2p_messages_proto_rawDescGZIP(), []int{3}
}

func (x *MetaDataV0) GetSeqNumber() uint64 {
	if x != nil {
		return x.SeqNumber
	}
	return 0
}

func (x *MetaDataV0) GetAttnets() github_com_prysmaticlabs_go_bitfield.Bitvector64 {
	if x != nil {
		return x.Attnets
	}
	return github_com_prysmaticlabs_go_bitfield.Bitvector64(nil)
}

type MetaDataV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNumber uint64                                            `protobuf:"varint,1,opt,name=seq_number,json=seqNumber,proto3" json:"seq_number,omitempty"`
	Attnets   github_com_prysmaticlabs_go_bitfield.Bitvector64  `protobuf:"bytes,2,opt,name=attnets,proto3" json:"attnets,omitempty" cast-type:"github.com/prysmaticlabs/go-bitfield.Bitvector64" ssz-size:"8"`
	Syncnets  github_com_prysmaticlabs_go_bitfield.Bitvector512 `protobuf:"bytes,3,opt,name=syncnets,proto3" json:"syncnets,omitempty" cast-type:"github.com/prysmaticlabs/go-bitfield.Bitvector512" ssz-size:"64"`
}

func (x *MetaDataV1) Reset() {
	*x = MetaDataV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaDataV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaDataV1) ProtoMessage() {}

func (x *MetaDataV1) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_v2_p2p_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaDataV1.ProtoReflect.Descriptor instead.
func (*MetaDataV1) Descriptor() ([]byte, []int) {
	return file_proto_prysm_v2_p2p_messages_proto_rawDescGZIP(), []int{4}
}

func (x *MetaDataV1) GetSeqNumber() uint64 {
	if x != nil {
		return x.SeqNumber
	}
	return 0
}

func (x *MetaDataV1) GetAttnets() github_com_prysmaticlabs_go_bitfield.Bitvector64 {
	if x != nil {
		return x.Attnets
	}
	return github_com_prysmaticlabs_go_bitfield.Bitvector64(nil)
}

func (x *MetaDataV1) GetSyncnets() github_com_prysmaticlabs_go_bitfield.Bitvector512 {
	if x != nil {
		return x.Syncnets
	}
	return github_com_prysmaticlabs_go_bitfield.Bitvector512(nil)
}

var File_proto_prysm_v2_p2p_messages_proto protoreflect.FileDescriptor

var file_proto_prysm_v2_p2p_messages_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x70, 0x32, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x2e, 0x76, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x26, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x6b, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x05, 0x8a, 0xb5, 0x18, 0x01, 0x34, 0x52, 0x0a, 0x66, 0x6f,
	0x72, 0x6b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0e, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0d, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x56, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x2d, 0x82, 0xb5, 0x18, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65,
	0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x52,
	0x0e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12,
	0x23, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64,
	0x52, 0x6f, 0x6f, 0x74, 0x12, 0x49, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x6c, 0x6f,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69,
	0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x53, 0x6c, 0x6f, 0x74, 0x22,
	0x93, 0x01, 0x0a, 0x1a, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x42, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0xcc, 0x01, 0x0a, 0x09, 0x45, 0x4e, 0x52, 0x46, 0x6f, 0x72,
	0x6b, 0x49, 0x44, 0x12, 0x35, 0x0a, 0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66,
	0x6f, 0x72, 0x6b, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x05, 0x8a, 0xb5, 0x18, 0x01, 0x34, 0x52, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x46, 0x6f, 0x72, 0x6b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x11, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x05, 0x8a, 0xb5, 0x18, 0x01, 0x34, 0x52, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x46, 0x6f, 0x72, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6b, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2d, 0x82, 0xb5, 0x18, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6b, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x22, 0x80, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x56, 0x30, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x53, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x39, 0x8a, 0xb5, 0x18, 0x01, 0x38, 0x82, 0xb5, 0x18, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x74, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x2e, 0x42, 0x69, 0x74, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x36, 0x34, 0x52, 0x07,
	0x61, 0x74, 0x74, 0x6e, 0x65, 0x74, 0x73, 0x22, 0xd9, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x56, 0x31, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x71, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x6e, 0x65, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x39, 0x8a, 0xb5, 0x18, 0x01, 0x38, 0x82, 0xb5, 0x18,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73,
	0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x74,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x69, 0x74, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x36,
	0x34, 0x52, 0x07, 0x61, 0x74, 0x74, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x57, 0x0a, 0x08, 0x73, 0x79,
	0x6e, 0x63, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x3b, 0x8a, 0xb5,
	0x18, 0x02, 0x36, 0x34, 0x82, 0xb5, 0x18, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x69, 0x74, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x42, 0x69, 0x74,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x35, 0x31, 0x32, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x6e,
	0x65, 0x74, 0x73, 0x42, 0x85, 0x01, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2e, 0x76, 0x32, 0x42, 0x10, 0x50,
	0x32, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x32,
	0x3b, 0x76, 0x32, 0xaa, 0x02, 0x11, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x50,
	0x72, 0x79, 0x73, 0x6d, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x11, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x5c, 0x50, 0x72, 0x79, 0x73, 0x6d, 0x5c, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_prysm_v2_p2p_messages_proto_rawDescOnce sync.Once
	file_proto_prysm_v2_p2p_messages_proto_rawDescData = file_proto_prysm_v2_p2p_messages_proto_rawDesc
)

func file_proto_prysm_v2_p2p_messages_proto_rawDescGZIP() []byte {
	file_proto_prysm_v2_p2p_messages_proto_rawDescOnce.Do(func() {
		file_proto_prysm_v2_p2p_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_prysm_v2_p2p_messages_proto_rawDescData)
	})
	return file_proto_prysm_v2_p2p_messages_proto_rawDescData
}

var file_proto_prysm_v2_p2p_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_prysm_v2_p2p_messages_proto_goTypes = []interface{}{
	(*Status)(nil),                     // 0: ethereum.prysm.v2.Status
	(*BeaconBlocksByRangeRequest)(nil), // 1: ethereum.prysm.v2.BeaconBlocksByRangeRequest
	(*ENRForkID)(nil),                  // 2: ethereum.prysm.v2.ENRForkID
	(*MetaDataV0)(nil),                 // 3: ethereum.prysm.v2.MetaDataV0
	(*MetaDataV1)(nil),                 // 4: ethereum.prysm.v2.MetaDataV1
}
var file_proto_prysm_v2_p2p_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_prysm_v2_p2p_messages_proto_init() }
func file_proto_prysm_v2_p2p_messages_proto_init() {
	if File_proto_prysm_v2_p2p_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_prysm_v2_p2p_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_proto_prysm_v2_p2p_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconBlocksByRangeRequest); i {
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
		file_proto_prysm_v2_p2p_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ENRForkID); i {
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
		file_proto_prysm_v2_p2p_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaDataV0); i {
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
		file_proto_prysm_v2_p2p_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaDataV1); i {
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
			RawDescriptor: file_proto_prysm_v2_p2p_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_prysm_v2_p2p_messages_proto_goTypes,
		DependencyIndexes: file_proto_prysm_v2_p2p_messages_proto_depIdxs,
		MessageInfos:      file_proto_prysm_v2_p2p_messages_proto_msgTypes,
	}.Build()
	File_proto_prysm_v2_p2p_messages_proto = out.File
	file_proto_prysm_v2_p2p_messages_proto_rawDesc = nil
	file_proto_prysm_v2_p2p_messages_proto_goTypes = nil
	file_proto_prysm_v2_p2p_messages_proto_depIdxs = nil
}
