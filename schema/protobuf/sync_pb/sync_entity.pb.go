// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: sync_entity.proto

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

type SyncEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This item's identifier.  In a commit of a new item, this will be a
	// client-generated ID.  If the commit succeeds, the server will generate
	// a globally unique ID and return it to the committing client in the
	// CommitResponse.EntryResponse.  In the context of a GetUpdatesResponse,
	// |id_string| is always the server generated ID.  The original
	// client-generated ID is preserved in the |originator_client_id| field.
	// Present in both GetUpdatesResponse and CommitMessage.
	IdString *string `protobuf:"bytes,1,opt,name=id_string,json=idString" json:"id_string,omitempty"`
	// An id referencing this item's parent in the hierarchy.  In a
	// CommitMessage, it is accepted for this to be a client-generated temporary
	// ID if there was a new created item with that ID appearing earlier
	// in the message.  In all other situations, it is a server ID.
	// Present in both GetUpdatesResponse and CommitMessage.
	//
	// Starting with M99, this field is optional and used for legacy bookmarks
	// only:
	// 1. When processing GetUpdatesResponse, it is unused for modern data created
	//    or reuploaded by M94 or above, which populates the parent's GUID in
	//    BookmarkSpecifics (which is sufficient).
	// 2. When issuing CommitMessage, the field is populated for compatibility
	//    with clients before M99.
	ParentIdString *string `protobuf:"bytes,2,opt,name=parent_id_string,json=parentIdString" json:"parent_id_string,omitempty"`
	// The version of this item -- a monotonically increasing value that is
	// maintained by for each item.  If zero in a CommitMessage, the server
	// will interpret this entity as a newly-created item and generate a
	// new server ID and an initial version number.  If nonzero in a
	// CommitMessage, this item is treated as an update to an existing item, and
	// the server will use |id_string| to locate the item.  Then, if the item's
	// current version on the server does not match |version|, the commit will
	// fail for that item.  The server will not update it, and will return
	// a result code of CONFLICT.  In a GetUpdatesResponse, |version| is
	// always positive and indentifies the revision of the item data being sent
	// to the client.
	// Present in both GetUpdatesResponse and CommitMessage.
	// WARNING: This field used to be required before M60. Any client before this
	// will fail to deserialize if this field is missing.
	Version *int64 `protobuf:"varint,4,opt,name=version" json:"version,omitempty"`
	// Last modification time, in milliseconds since Unix epoch.
	// Present in both GetUpdatesResponse and CommitMessage.
	Mtime *int64 `protobuf:"varint,5,opt,name=mtime" json:"mtime,omitempty"`
	// Creation time, in milliseconds since Unix epoch.
	// Present in both GetUpdatesResponse and CommitMessage.
	Ctime *int64 `protobuf:"varint,6,opt,name=ctime" json:"ctime,omitempty"`
	// The name of this item.
	// Historical note:
	//   Since November 2010, this value is no different from non_unique_name.
	//   Before then, server implementations would maintain a unique-within-parent
	//   value separate from its base, "non-unique" value.  Clients had not
	//   depended on the uniqueness of the property since November 2009; it was
	//   removed from Chromium by http://codereview.chromium.org/371029 .
	// Present in both GetUpdatesResponse and CommitMessage.
	// WARNING: This field used to be required before M60. Any client before this
	// will fail to deserialize if this field is missing.
	Name *string `protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	// The name of this item.  Same as |name|.
	// |non_unique_name| should take precedence over the |name| value if both
	// are supplied.  For efficiency, clients and servers should avoid setting
	// this redundant value.
	// Present in both GetUpdatesResponse and CommitMessage.
	NonUniqueName *string `protobuf:"bytes,8,opt,name=non_unique_name,json=nonUniqueName" json:"non_unique_name,omitempty"`
	// If present, this tag identifies this item as being a uniquely
	// instanced item.  The server ensures that there is never more
	// than one entity in a user's store with the same tag value.
	// This value is used to identify and find e.g. the "Google Chrome" settings
	// folder without relying on it existing at a particular path, or having
	// a particular name, in the data store.
	//
	// This variant of the tag is created by the server, so clients can't create
	// an item with a tag using this field.
	//
	// Use client_defined_unique_tag if you want to create one from the client.
	//
	// An item can't have both a client_defined_unique_tag and
	// a server_defined_unique_tag.
	//
	// Present only in GetUpdatesResponse.
	ServerDefinedUniqueTag *string `protobuf:"bytes,10,opt,name=server_defined_unique_tag,json=serverDefinedUniqueTag" json:"server_defined_unique_tag,omitempty"`
	// Ancient fields, predecessors for |unique_position|, deprecated with M26 and
	// still supported to deal with old incoming data. See field |unique_position|
	// for details as well as the data-upgrading implementation in
	// GetUniquePositionFromSyncEntity().
	//
	// Deprecated: Do not use.
	PositionInParent *int64 `protobuf:"varint,15,opt,name=position_in_parent,json=positionInParent" json:"position_in_parent,omitempty"`
	// Deprecated: Do not use.
	InsertAfterItemId *string `protobuf:"bytes,16,opt,name=insert_after_item_id,json=insertAfterItemId" json:"insert_after_item_id,omitempty"`
	// If true, indicates that this item has been (or should be) deleted.
	// Present in both GetUpdatesResponse and CommitMessage.
	Deleted *bool `protobuf:"varint,18,opt,name=deleted,def=0" json:"deleted,omitempty"`
	// A unique ID that identifies the the sync client who initially committed
	// this entity.  This value corresponds to |cache_guid| in CommitMessage.
	// This field, along with |originator_client_item_id|, can be used to
	// reunite the original with its official committed version in the case
	// where a client does not receive or process the commit response for
	// some reason.
	//
	// Present only in GetUpdatesResponse.
	//
	// This field is also used in determining the unique identifier used in
	// bookmarks' unique_position field.
	OriginatorCacheGuid *string `protobuf:"bytes,19,opt,name=originator_cache_guid,json=originatorCacheGuid" json:"originator_cache_guid,omitempty"`
	// Item ID as generated by the client that initially created this entity. Used
	// exclusively for bookmarks (other datatypes use client_defined_unique_tag).
	// There are three generation of bookmarks that have populated this field
	// differently, depending on which version of the browser created the
	// bookmark:
	// 1. For bookmarks created before M44 (2015), the field got populated with an
	//    ID that is locally unique, but not globally unique (usually a negative
	//    number).
	// 2. For bookmarks created between M45 and M51, both inclusive, the field got
	//    populated with a globally unique GUID in uppercase form.
	// 3. For bookmarks created with M52 or above, the field gets populated with
	//    a globally unique GUID in lowercase form.
	//
	// Present only in GetUpdatesResponse.
	OriginatorClientItemId *string `protobuf:"bytes,20,opt,name=originator_client_item_id,json=originatorClientItemId" json:"originator_client_item_id,omitempty"`
	// Extensible container for datatype-specific data.
	// This became available in version 23 of the protocol.
	Specifics *EntitySpecifics `protobuf:"bytes,21,opt,name=specifics" json:"specifics,omitempty"`
	// Indicate whether this is a folder or not. Available in version 23+.
	Folder *bool `protobuf:"varint,22,opt,name=folder,def=0" json:"folder,omitempty"`
	// A client defined unique hash for this entity.
	// Similar to server_defined_unique_tag.
	//
	// When initially committing an entity, a client can request that the entity
	// is unique per that account. To do so, the client should specify a
	// client_defined_unique_tag. At most one entity per tag value may exist.
	// per account. The server will enforce uniqueness on this tag
	// and fail attempts to create duplicates of this tag.
	// Will be returned in any updates for this entity.
	//
	// The difference between server_defined_unique_tag and
	// client_defined_unique_tag is the creator of the entity. Server defined
	// tags are entities created by the server at account creation,
	// while client defined tags are entities created by the client at any time.
	//
	// During GetUpdates, a sync entity update will come back with ONE of:
	// a) Originator and cache id - If client committed the item as non "unique"
	// b) Server tag - If server committed the item as unique
	// c) Client tag - If client committed the item as unique
	//
	// May be present in CommitMessages for the initial creation of an entity.
	// If present in Commit updates for the entity, it will be ignored.
	//
	// May be returned in GetUpdatesMessage and sent up in CommitMessage.
	ClientDefinedUniqueTag *string `protobuf:"bytes,23,opt,name=client_defined_unique_tag,json=clientDefinedUniqueTag" json:"client_defined_unique_tag,omitempty"`
	// Introduced in M26, represents ordering among entities, in practice used for
	// bookmarks only. Clients should not assume it is always populated in
	// GetUpdatesMessage due to the following caveats:
	// 1. Tombstones and permanent folders do not populate it (ordering is
	//    irrelevant).
	// 2. It may remain unset by future versions of the client, as long as the
	//    field with the same name is populated inside BookmarkSpecifics. M94 and
	//    above populate both for backward compatibility reasons, but when support
	//    for M93 is retired, modern clients at the time may stop populating this
	//    field.
	// 3. Very old data (last committed by M25 or below, before the field was
	//    introduced) does not include this field, and in that case the legacy
	//    fields |position_in_parent| or |insert_after_item_id| must be honored
	//    instead.
	//
	// May be returned in GetUpdatesMessage and sent up in CommitMessage.
	UniquePosition *UniquePosition `protobuf:"bytes,25,opt,name=unique_position,json=uniquePosition" json:"unique_position,omitempty"`
}

// Default values for SyncEntity fields.
const (
	Default_SyncEntity_Deleted = bool(false)
	Default_SyncEntity_Folder  = bool(false)
)

func (x *SyncEntity) Reset() {
	*x = SyncEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncEntity) ProtoMessage() {}

func (x *SyncEntity) ProtoReflect() protoreflect.Message {
	mi := &file_sync_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncEntity.ProtoReflect.Descriptor instead.
func (*SyncEntity) Descriptor() ([]byte, []int) {
	return file_sync_entity_proto_rawDescGZIP(), []int{0}
}

func (x *SyncEntity) GetIdString() string {
	if x != nil && x.IdString != nil {
		return *x.IdString
	}
	return ""
}

func (x *SyncEntity) GetParentIdString() string {
	if x != nil && x.ParentIdString != nil {
		return *x.ParentIdString
	}
	return ""
}

func (x *SyncEntity) GetVersion() int64 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *SyncEntity) GetMtime() int64 {
	if x != nil && x.Mtime != nil {
		return *x.Mtime
	}
	return 0
}

func (x *SyncEntity) GetCtime() int64 {
	if x != nil && x.Ctime != nil {
		return *x.Ctime
	}
	return 0
}

func (x *SyncEntity) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SyncEntity) GetNonUniqueName() string {
	if x != nil && x.NonUniqueName != nil {
		return *x.NonUniqueName
	}
	return ""
}

func (x *SyncEntity) GetServerDefinedUniqueTag() string {
	if x != nil && x.ServerDefinedUniqueTag != nil {
		return *x.ServerDefinedUniqueTag
	}
	return ""
}

// Deprecated: Do not use.
func (x *SyncEntity) GetPositionInParent() int64 {
	if x != nil && x.PositionInParent != nil {
		return *x.PositionInParent
	}
	return 0
}

// Deprecated: Do not use.
func (x *SyncEntity) GetInsertAfterItemId() string {
	if x != nil && x.InsertAfterItemId != nil {
		return *x.InsertAfterItemId
	}
	return ""
}

func (x *SyncEntity) GetDeleted() bool {
	if x != nil && x.Deleted != nil {
		return *x.Deleted
	}
	return Default_SyncEntity_Deleted
}

func (x *SyncEntity) GetOriginatorCacheGuid() string {
	if x != nil && x.OriginatorCacheGuid != nil {
		return *x.OriginatorCacheGuid
	}
	return ""
}

func (x *SyncEntity) GetOriginatorClientItemId() string {
	if x != nil && x.OriginatorClientItemId != nil {
		return *x.OriginatorClientItemId
	}
	return ""
}

func (x *SyncEntity) GetSpecifics() *EntitySpecifics {
	if x != nil {
		return x.Specifics
	}
	return nil
}

func (x *SyncEntity) GetFolder() bool {
	if x != nil && x.Folder != nil {
		return *x.Folder
	}
	return Default_SyncEntity_Folder
}

func (x *SyncEntity) GetClientDefinedUniqueTag() string {
	if x != nil && x.ClientDefinedUniqueTag != nil {
		return *x.ClientDefinedUniqueTag
	}
	return ""
}

func (x *SyncEntity) GetUniquePosition() *UniquePosition {
	if x != nil {
		return x.UniquePosition
	}
	return nil
}

var File_sync_entity_proto protoreflect.FileDescriptor

var file_sync_entity_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x1a, 0x16, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x07, 0x0a, 0x0a,
	0x53, 0x79, 0x6e, 0x63, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x74, 0x61, 0x67,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x54, 0x61, 0x67, 0x12, 0x30,
	0x0a, 0x12, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18, 0x01, 0x52, 0x10,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x33, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x75, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x73, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x0a,
	0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66,
	0x61, 0x6c, 0x73, 0x65, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x19,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x55, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x54, 0x61, 0x67, 0x12, 0x40, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a,
	0x04, 0x08, 0x09, 0x10, 0x0a, 0x4a, 0x04, 0x08, 0x0b, 0x10, 0x0c, 0x4a, 0x04, 0x08, 0x0c, 0x10,
	0x0d, 0x4a, 0x04, 0x08, 0x0d, 0x10, 0x0e, 0x4a, 0x04, 0x08, 0x0e, 0x10, 0x0f, 0x4a, 0x04, 0x08,
	0x11, 0x10, 0x12, 0x4a, 0x04, 0x08, 0x18, 0x10, 0x19, 0x4a, 0x04, 0x08, 0x1a, 0x10, 0x1b, 0x52,
	0x0d, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x0e,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c,
	0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x62, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x62,
	0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x52, 0x10, 0x62, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x52, 0x13, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x11, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x5f, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50,
	0x01,
}

var (
	file_sync_entity_proto_rawDescOnce sync.Once
	file_sync_entity_proto_rawDescData = file_sync_entity_proto_rawDesc
)

func file_sync_entity_proto_rawDescGZIP() []byte {
	file_sync_entity_proto_rawDescOnce.Do(func() {
		file_sync_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_entity_proto_rawDescData)
	})
	return file_sync_entity_proto_rawDescData
}

var file_sync_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sync_entity_proto_goTypes = []interface{}{
	(*SyncEntity)(nil),      // 0: sync_pb.SyncEntity
	(*EntitySpecifics)(nil), // 1: sync_pb.EntitySpecifics
	(*UniquePosition)(nil),  // 2: sync_pb.UniquePosition
}
var file_sync_entity_proto_depIdxs = []int32{
	1, // 0: sync_pb.SyncEntity.specifics:type_name -> sync_pb.EntitySpecifics
	2, // 1: sync_pb.SyncEntity.unique_position:type_name -> sync_pb.UniquePosition
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sync_entity_proto_init() }
func file_sync_entity_proto_init() {
	if File_sync_entity_proto != nil {
		return
	}
	file_entity_specifics_proto_init()
	file_unique_position_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sync_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncEntity); i {
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
			RawDescriptor: file_sync_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_entity_proto_goTypes,
		DependencyIndexes: file_sync_entity_proto_depIdxs,
		MessageInfos:      file_sync_entity_proto_msgTypes,
	}.Build()
	File_sync_entity_proto = out.File
	file_sync_entity_proto_rawDesc = nil
	file_sync_entity_proto_goTypes = nil
	file_sync_entity_proto_depIdxs = nil
}
