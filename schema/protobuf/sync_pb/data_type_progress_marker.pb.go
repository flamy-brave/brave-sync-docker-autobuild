// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: data_type_progress_marker.proto

package sync_pb

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

type GarbageCollectionDirective_Type int32

const (
	GarbageCollectionDirective_UNKNOWN           GarbageCollectionDirective_Type = 0
	GarbageCollectionDirective_VERSION_WATERMARK GarbageCollectionDirective_Type = 1
	GarbageCollectionDirective_AGE_WATERMARK     GarbageCollectionDirective_Type = 2
	// Deprecated: Do not use.
	GarbageCollectionDirective_DEPRECATED_MAX_ITEM_COUNT GarbageCollectionDirective_Type = 3
)

// Enum value maps for GarbageCollectionDirective_Type.
var (
	GarbageCollectionDirective_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "VERSION_WATERMARK",
		2: "AGE_WATERMARK",
		3: "DEPRECATED_MAX_ITEM_COUNT",
	}
	GarbageCollectionDirective_Type_value = map[string]int32{
		"UNKNOWN":                   0,
		"VERSION_WATERMARK":         1,
		"AGE_WATERMARK":             2,
		"DEPRECATED_MAX_ITEM_COUNT": 3,
	}
)

func (x GarbageCollectionDirective_Type) Enum() *GarbageCollectionDirective_Type {
	p := new(GarbageCollectionDirective_Type)
	*p = x
	return p
}

func (x GarbageCollectionDirective_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GarbageCollectionDirective_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_data_type_progress_marker_proto_enumTypes[0].Descriptor()
}

func (GarbageCollectionDirective_Type) Type() protoreflect.EnumType {
	return &file_data_type_progress_marker_proto_enumTypes[0]
}

func (x GarbageCollectionDirective_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GarbageCollectionDirective_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GarbageCollectionDirective_Type(num)
	return nil
}

// Deprecated: Use GarbageCollectionDirective_Type.Descriptor instead.
func (GarbageCollectionDirective_Type) EnumDescriptor() ([]byte, []int) {
	return file_data_type_progress_marker_proto_rawDescGZIP(), []int{2, 0}
}

type DataTypeProgressMarker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An integer identifying the data type whose progress is tracked by this
	// marker.  The legitimate values of this field correspond to the protobuf
	// field numbers of all EntitySpecifics fields supported by the server.
	// These values are externally declared in per-datatype .proto files.
	DataTypeId *int32 `protobuf:"varint,1,opt,name=data_type_id,json=dataTypeId" json:"data_type_id,omitempty"`
	// An opaque-to-the-client sequence of bytes that the server may interpret
	// as an indicator of the client's knowledge state.  If this is empty or
	// omitted by the client, it indicates that the client is initiating a
	// a first-time sync of this datatype.  Otherwise, clients must supply a
	// value previously returned by the server in an earlier GetUpdatesResponse.
	// These values are not comparable or generable on the client.
	//
	// The opaque semantics of this field are to afford server implementations
	// some flexibility in implementing progress tracking.  For instance,
	// a server implementation built on top of a distributed storage service --
	// or multiple heterogenous such services -- might need to supply a vector
	// of totally ordered monotonic update timestamps, rather than a single
	// monotonically increasing value.  Other optimizations may also be
	// possible if the server is allowed to embed arbitrary information in
	// the progress token.
	//
	// Server implementations should keep the size of these tokens relatively
	// small, on the order of tens of bytes, and they should remain small
	// regardless of the number of items synchronized.  (A possible bad server
	// implementation would be for progress_token to contain a list of all the
	// items ever sent to the client.  Servers shouldn't do this.)
	Token []byte `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	// This field will be included only in GetUpdates with origin GU_TRIGGER.
	GetUpdateTriggers *GetUpdateTriggers `protobuf:"bytes,5,opt,name=get_update_triggers,json=getUpdateTriggers" json:"get_update_triggers,omitempty"`
	// The garbage collection directive for this data type.  The client should
	// purge items locally based on this directive.  Since this directive is
	// designed to be sent from server only, the client should persist it locally
	// as needed and avoid sending it to the server.
	GcDirective *GarbageCollectionDirective `protobuf:"bytes,6,opt,name=gc_directive,json=gcDirective" json:"gc_directive,omitempty"`
}

func (x *DataTypeProgressMarker) Reset() {
	*x = DataTypeProgressMarker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_type_progress_marker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataTypeProgressMarker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTypeProgressMarker) ProtoMessage() {}

func (x *DataTypeProgressMarker) ProtoReflect() protoreflect.Message {
	mi := &file_data_type_progress_marker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTypeProgressMarker.ProtoReflect.Descriptor instead.
func (*DataTypeProgressMarker) Descriptor() ([]byte, []int) {
	return file_data_type_progress_marker_proto_rawDescGZIP(), []int{0}
}

func (x *DataTypeProgressMarker) GetDataTypeId() int32 {
	if x != nil && x.DataTypeId != nil {
		return *x.DataTypeId
	}
	return 0
}

func (x *DataTypeProgressMarker) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *DataTypeProgressMarker) GetGetUpdateTriggers() *GetUpdateTriggers {
	if x != nil {
		return x.GetUpdateTriggers
	}
	return nil
}

func (x *DataTypeProgressMarker) GetGcDirective() *GarbageCollectionDirective {
	if x != nil {
		return x.GcDirective
	}
	return nil
}

// A single datatype's sync context. Allows the datatype to pass along
// datatype specific information with its own server backend.
type DataTypeContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type this context is associated with.
	DataTypeId *int32 `protobuf:"varint,1,opt,name=data_type_id,json=dataTypeId" json:"data_type_id,omitempty"`
	// The context for the datatype.
	Context []byte `protobuf:"bytes,2,opt,name=context" json:"context,omitempty"`
	// The version of the context.
	Version *int64 `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
}

func (x *DataTypeContext) Reset() {
	*x = DataTypeContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_type_progress_marker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataTypeContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTypeContext) ProtoMessage() {}

func (x *DataTypeContext) ProtoReflect() protoreflect.Message {
	mi := &file_data_type_progress_marker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTypeContext.ProtoReflect.Descriptor instead.
func (*DataTypeContext) Descriptor() ([]byte, []int) {
	return file_data_type_progress_marker_proto_rawDescGZIP(), []int{1}
}

func (x *DataTypeContext) GetDataTypeId() int32 {
	if x != nil && x.DataTypeId != nil {
		return *x.DataTypeId
	}
	return 0
}

func (x *DataTypeContext) GetContext() []byte {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *DataTypeContext) GetVersion() int64 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

type GarbageCollectionDirective struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *GarbageCollectionDirective_Type `protobuf:"varint,1,opt,name=type,enum=sync_pb.GarbageCollectionDirective_Type,def=0" json:"type,omitempty"`
	// This field specifies the watermark for the versions which should get
	// garbage collected.  The client should purge all sync entities when
	// receiving any value of this.  This is a change from previous behavior,
	// where the client would only be required to purge items older than the
	// specified watermark.
	// TODO(crbug.com/877951): Rename this to make clear that whenever it's set,
	// the client will delete ALL data, regardless of its value.
	VersionWatermark *int64 `protobuf:"varint,2,opt,name=version_watermark,json=versionWatermark" json:"version_watermark,omitempty"`
	// This field specifies the watermark in terms of age in days.  The client
	// should purge all sync entities which are older than this specific value
	// based on last modified time.
	AgeWatermarkInDays *int32 `protobuf:"varint,3,opt,name=age_watermark_in_days,json=ageWatermarkInDays" json:"age_watermark_in_days,omitempty"`
}

// Default values for GarbageCollectionDirective fields.
const (
	Default_GarbageCollectionDirective_Type = GarbageCollectionDirective_UNKNOWN
)

func (x *GarbageCollectionDirective) Reset() {
	*x = GarbageCollectionDirective{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_type_progress_marker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarbageCollectionDirective) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarbageCollectionDirective) ProtoMessage() {}

func (x *GarbageCollectionDirective) ProtoReflect() protoreflect.Message {
	mi := &file_data_type_progress_marker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarbageCollectionDirective.ProtoReflect.Descriptor instead.
func (*GarbageCollectionDirective) Descriptor() ([]byte, []int) {
	return file_data_type_progress_marker_proto_rawDescGZIP(), []int{2}
}

func (x *GarbageCollectionDirective) GetType() GarbageCollectionDirective_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_GarbageCollectionDirective_Type
}

func (x *GarbageCollectionDirective) GetVersionWatermark() int64 {
	if x != nil && x.VersionWatermark != nil {
		return *x.VersionWatermark
	}
	return 0
}

func (x *GarbageCollectionDirective) GetAgeWatermarkInDays() int32 {
	if x != nil && x.AgeWatermarkInDays != nil {
		return *x.AgeWatermarkInDays
	}
	return 0
}

// This message communicates additional per-type information related to
// requests with origin GU_TRIGGER.  This message is not relevant when any
// other origin value is used.
// Introduced in M29.
type GetUpdateTriggers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An opaque-to-the-client string of bytes, received through a notification,
	// that the server may interpret as a hint about the location of the latest
	// version of the data for this type.
	//
	// Note that this will eventually replace the 'optional' field of the same
	// name defined in the progress marker, but the client and server should
	// support both until it's safe to deprecate the old one.
	//
	// This field was introduced in M29.
	NotificationHint []string `protobuf:"bytes,1,rep,name=notification_hint,json=notificationHint" json:"notification_hint,omitempty"`
	// This flag is set if the client was forced to drop hints because the number
	// of queued hints exceeded its limit.  The oldest hints will be discarded
	// first.  Introduced in M29.
	ClientDroppedHints *bool `protobuf:"varint,2,opt,name=client_dropped_hints,json=clientDroppedHints" json:"client_dropped_hints,omitempty"`
	// This flag is set when the client suspects that its list of invalidation
	// hints may be incomplete.  This may be the case if:
	// - The client is syncing for the first time.
	// - The client has just restarted and it was unable to keep track of
	//   invalidations that were received prior to the restart.
	// - The client's connection to the invalidation server is currently or
	//   was recently broken.
	//
	// It's difficult to provide more details here.  This is implemented by
	// setting the flag to false whenever anything that might adversely affect
	// notifications happens (eg. a crash, restart on a platform that doesn't
	// support invalidation ack-tracking, transient invalidation error) and is
	// unset only after we've experienced one successful sync cycle while
	// notifications were enabled.
	//
	// This flag was introduced in M29.
	InvalidationsOutOfSync *bool `protobuf:"varint,3,opt,name=invalidations_out_of_sync,json=invalidationsOutOfSync" json:"invalidations_out_of_sync,omitempty"`
	// This counts the number of times the syncer has been asked to commit
	// changes for this type since the last successful sync cycle.  The number of
	// nudges may not be related to the actual number of items modified.  It
	// often correlates with the number of user actions, but that's not always
	// the case.
	// Introduced in M29.
	LocalModificationNudges *int64 `protobuf:"varint,4,opt,name=local_modification_nudges,json=localModificationNudges" json:"local_modification_nudges,omitempty"`
	// This counts the number of times the syncer has been explicitly asked to
	// fetch updates for this type since the last successful sync cycle.  These
	// explicit refresh requests should be relatively rare on most platforms, and
	// associated with user actions.  For example, at the time of this writing
	// the most common (only?) source of refresh requests is when a user opens
	// the new tab page on a platform that does not support sessions
	// invalidations.
	// Introduced in M29.
	DatatypeRefreshNudges *int64 `protobuf:"varint,5,opt,name=datatype_refresh_nudges,json=datatypeRefreshNudges" json:"datatype_refresh_nudges,omitempty"`
	// This flag is set if the invalidation server reports that it may have
	// dropped some invalidations at some point.  Introduced in M33.
	ServerDroppedHints *bool `protobuf:"varint,6,opt,name=server_dropped_hints,json=serverDroppedHints" json:"server_dropped_hints,omitempty"`
	// This flag is set if this GetUpdate request is due at least in part due
	// to the fact that this type has not finished initial sync yet, and the
	// client would like to initialize itself with the server data.
	//
	// Only some types support performing an initial sync as part of a normal
	// GetUpdate request.  Many types must be in configure mode when fetching
	// initial sync data.
	//
	// Introduced in M38.
	InitialSyncInProgress *bool `protobuf:"varint,7,opt,name=initial_sync_in_progress,json=initialSyncInProgress" json:"initial_sync_in_progress,omitempty"`
	// This flag is set if this GetUpdate request is due to client receiving
	// conflict response from server, so client needs to sync and then resolve
	// conflict locally, and then commit again.
	//
	// Introduced in M42.
	SyncForResolveConflictInProgress *bool `protobuf:"varint,8,opt,name=sync_for_resolve_conflict_in_progress,json=syncForResolveConflictInProgress" json:"sync_for_resolve_conflict_in_progress,omitempty"`
}

func (x *GetUpdateTriggers) Reset() {
	*x = GetUpdateTriggers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_type_progress_marker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdateTriggers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdateTriggers) ProtoMessage() {}

func (x *GetUpdateTriggers) ProtoReflect() protoreflect.Message {
	mi := &file_data_type_progress_marker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdateTriggers.ProtoReflect.Descriptor instead.
func (*GetUpdateTriggers) Descriptor() ([]byte, []int) {
	return file_data_type_progress_marker_proto_rawDescGZIP(), []int{3}
}

func (x *GetUpdateTriggers) GetNotificationHint() []string {
	if x != nil {
		return x.NotificationHint
	}
	return nil
}

func (x *GetUpdateTriggers) GetClientDroppedHints() bool {
	if x != nil && x.ClientDroppedHints != nil {
		return *x.ClientDroppedHints
	}
	return false
}

func (x *GetUpdateTriggers) GetInvalidationsOutOfSync() bool {
	if x != nil && x.InvalidationsOutOfSync != nil {
		return *x.InvalidationsOutOfSync
	}
	return false
}

func (x *GetUpdateTriggers) GetLocalModificationNudges() int64 {
	if x != nil && x.LocalModificationNudges != nil {
		return *x.LocalModificationNudges
	}
	return 0
}

func (x *GetUpdateTriggers) GetDatatypeRefreshNudges() int64 {
	if x != nil && x.DatatypeRefreshNudges != nil {
		return *x.DatatypeRefreshNudges
	}
	return 0
}

func (x *GetUpdateTriggers) GetServerDroppedHints() bool {
	if x != nil && x.ServerDroppedHints != nil {
		return *x.ServerDroppedHints
	}
	return false
}

func (x *GetUpdateTriggers) GetInitialSyncInProgress() bool {
	if x != nil && x.InitialSyncInProgress != nil {
		return *x.InitialSyncInProgress
	}
	return false
}

func (x *GetUpdateTriggers) GetSyncForResolveConflictInProgress() bool {
	if x != nil && x.SyncForResolveConflictInProgress != nil {
		return *x.SyncForResolveConflictInProgress
	}
	return false
}

var File_data_type_progress_marker_proto protoreflect.FileDescriptor

var file_data_type_progress_marker_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0xa2, 0x02, 0x0a, 0x16, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4a, 0x0a,
	0x13, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x11, 0x67, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x46, 0x0a, 0x0c, 0x67, 0x63, 0x5f,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x72, 0x62, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x52, 0x0b, 0x67, 0x63, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x52, 0x1d, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x66,
	0x6f, 0x72, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x22,
	0x67, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc0, 0x02, 0x0a, 0x1a, 0x47, 0x61, 0x72,
	0x62, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e,
	0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x3a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x31, 0x0a, 0x15, 0x61,
	0x67, 0x65, 0x5f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x5f,
	0x64, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x61, 0x67, 0x65, 0x57,
	0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x22, 0x60,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x57,
	0x41, 0x54, 0x45, 0x52, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x47,
	0x45, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x02, 0x12, 0x21, 0x0a,
	0x19, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x4d, 0x41, 0x58, 0x5f,
	0x49, 0x54, 0x45, 0x4d, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x03, 0x1a, 0x02, 0x08, 0x01,
	0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xdd, 0x03, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x12, 0x2b, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x30,
	0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64,
	0x5f, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x44, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x48, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x39, 0x0a, 0x19, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x3a, 0x0a, 0x19, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x75, 0x64, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x75, 0x64, 0x67, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x6e, 0x75, 0x64, 0x67,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4e, 0x75, 0x64, 0x67, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65,
	0x64, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x48, 0x69, 0x6e, 0x74,
	0x73, 0x12, 0x37, 0x0a, 0x18, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x79, 0x6e, 0x63,
	0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4f, 0x0a, 0x25, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x46,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63,
	0x74, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x2b, 0x0a, 0x25, 0x6f,
	0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_data_type_progress_marker_proto_rawDescOnce sync.Once
	file_data_type_progress_marker_proto_rawDescData = file_data_type_progress_marker_proto_rawDesc
)

func file_data_type_progress_marker_proto_rawDescGZIP() []byte {
	file_data_type_progress_marker_proto_rawDescOnce.Do(func() {
		file_data_type_progress_marker_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_type_progress_marker_proto_rawDescData)
	})
	return file_data_type_progress_marker_proto_rawDescData
}

var file_data_type_progress_marker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_data_type_progress_marker_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_type_progress_marker_proto_goTypes = []interface{}{
	(GarbageCollectionDirective_Type)(0), // 0: sync_pb.GarbageCollectionDirective.Type
	(*DataTypeProgressMarker)(nil),       // 1: sync_pb.DataTypeProgressMarker
	(*DataTypeContext)(nil),              // 2: sync_pb.DataTypeContext
	(*GarbageCollectionDirective)(nil),   // 3: sync_pb.GarbageCollectionDirective
	(*GetUpdateTriggers)(nil),            // 4: sync_pb.GetUpdateTriggers
}
var file_data_type_progress_marker_proto_depIdxs = []int32{
	4, // 0: sync_pb.DataTypeProgressMarker.get_update_triggers:type_name -> sync_pb.GetUpdateTriggers
	3, // 1: sync_pb.DataTypeProgressMarker.gc_directive:type_name -> sync_pb.GarbageCollectionDirective
	0, // 2: sync_pb.GarbageCollectionDirective.type:type_name -> sync_pb.GarbageCollectionDirective.Type
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_data_type_progress_marker_proto_init() }
func file_data_type_progress_marker_proto_init() {
	if File_data_type_progress_marker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_type_progress_marker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataTypeProgressMarker); i {
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
		file_data_type_progress_marker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataTypeContext); i {
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
		file_data_type_progress_marker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarbageCollectionDirective); i {
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
		file_data_type_progress_marker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpdateTriggers); i {
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
			RawDescriptor: file_data_type_progress_marker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_type_progress_marker_proto_goTypes,
		DependencyIndexes: file_data_type_progress_marker_proto_depIdxs,
		EnumInfos:         file_data_type_progress_marker_proto_enumTypes,
		MessageInfos:      file_data_type_progress_marker_proto_msgTypes,
	}.Build()
	File_data_type_progress_marker_proto = out.File
	file_data_type_progress_marker_proto_rawDesc = nil
	file_data_type_progress_marker_proto_goTypes = nil
	file_data_type_progress_marker_proto_depIdxs = nil
}
