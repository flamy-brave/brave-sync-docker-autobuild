// Copyright 2024 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for shared tab groups and tabs.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: shared_tab_group_data_specifics.proto

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

// These colors map to tab_groups::TabGroupColorId. They DO NOT match the enum
// integer values due to "kGrey" being in the "Unspecified" field.
type SharedTabGroup_Color int32

const (
	SharedTabGroup_UNSPECIFIED SharedTabGroup_Color = 0
	SharedTabGroup_GREY        SharedTabGroup_Color = 1
	SharedTabGroup_BLUE        SharedTabGroup_Color = 2
	SharedTabGroup_RED         SharedTabGroup_Color = 3
	SharedTabGroup_YELLOW      SharedTabGroup_Color = 4
	SharedTabGroup_GREEN       SharedTabGroup_Color = 5
	SharedTabGroup_PINK        SharedTabGroup_Color = 6
	SharedTabGroup_PURPLE      SharedTabGroup_Color = 7
	SharedTabGroup_CYAN        SharedTabGroup_Color = 8
	SharedTabGroup_ORANGE      SharedTabGroup_Color = 9
)

// Enum value maps for SharedTabGroup_Color.
var (
	SharedTabGroup_Color_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "GREY",
		2: "BLUE",
		3: "RED",
		4: "YELLOW",
		5: "GREEN",
		6: "PINK",
		7: "PURPLE",
		8: "CYAN",
		9: "ORANGE",
	}
	SharedTabGroup_Color_value = map[string]int32{
		"UNSPECIFIED": 0,
		"GREY":        1,
		"BLUE":        2,
		"RED":         3,
		"YELLOW":      4,
		"GREEN":       5,
		"PINK":        6,
		"PURPLE":      7,
		"CYAN":        8,
		"ORANGE":      9,
	}
)

func (x SharedTabGroup_Color) Enum() *SharedTabGroup_Color {
	p := new(SharedTabGroup_Color)
	*p = x
	return p
}

func (x SharedTabGroup_Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SharedTabGroup_Color) Descriptor() protoreflect.EnumDescriptor {
	return file_shared_tab_group_data_specifics_proto_enumTypes[0].Descriptor()
}

func (SharedTabGroup_Color) Type() protoreflect.EnumType {
	return &file_shared_tab_group_data_specifics_proto_enumTypes[0]
}

func (x SharedTabGroup_Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SharedTabGroup_Color) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SharedTabGroup_Color(num)
	return nil
}

// Deprecated: Use SharedTabGroup_Color.Descriptor instead.
func (SharedTabGroup_Color) EnumDescriptor() ([]byte, []int) {
	return file_shared_tab_group_data_specifics_proto_rawDescGZIP(), []int{1, 0}
}

type SharedTabGroupDataSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for this entity.
	Guid *string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	// TODO(b/319521964): consider to rename and add a comment to clarify the
	// content.
	LastModificationAuthor *string `protobuf:"bytes,2,opt,name=last_modification_author,json=lastModificationAuthor" json:"last_modification_author,omitempty"`
	// Types that are assignable to Entity:
	//	*SharedTabGroupDataSpecifics_TabGroup
	//	*SharedTabGroupDataSpecifics_Tab
	Entity isSharedTabGroupDataSpecifics_Entity `protobuf_oneof:"entity"`
}

func (x *SharedTabGroupDataSpecifics) Reset() {
	*x = SharedTabGroupDataSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_tab_group_data_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedTabGroupDataSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedTabGroupDataSpecifics) ProtoMessage() {}

func (x *SharedTabGroupDataSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_shared_tab_group_data_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedTabGroupDataSpecifics.ProtoReflect.Descriptor instead.
func (*SharedTabGroupDataSpecifics) Descriptor() ([]byte, []int) {
	return file_shared_tab_group_data_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *SharedTabGroupDataSpecifics) GetGuid() string {
	if x != nil && x.Guid != nil {
		return *x.Guid
	}
	return ""
}

func (x *SharedTabGroupDataSpecifics) GetLastModificationAuthor() string {
	if x != nil && x.LastModificationAuthor != nil {
		return *x.LastModificationAuthor
	}
	return ""
}

func (m *SharedTabGroupDataSpecifics) GetEntity() isSharedTabGroupDataSpecifics_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *SharedTabGroupDataSpecifics) GetTabGroup() *SharedTabGroup {
	if x, ok := x.GetEntity().(*SharedTabGroupDataSpecifics_TabGroup); ok {
		return x.TabGroup
	}
	return nil
}

func (x *SharedTabGroupDataSpecifics) GetTab() *SharedTab {
	if x, ok := x.GetEntity().(*SharedTabGroupDataSpecifics_Tab); ok {
		return x.Tab
	}
	return nil
}

type isSharedTabGroupDataSpecifics_Entity interface {
	isSharedTabGroupDataSpecifics_Entity()
}

type SharedTabGroupDataSpecifics_TabGroup struct {
	TabGroup *SharedTabGroup `protobuf:"bytes,3,opt,name=tab_group,json=tabGroup,oneof"`
}

type SharedTabGroupDataSpecifics_Tab struct {
	Tab *SharedTab `protobuf:"bytes,4,opt,name=tab,oneof"`
}

func (*SharedTabGroupDataSpecifics_TabGroup) isSharedTabGroupDataSpecifics_Entity() {}

func (*SharedTabGroupDataSpecifics_Tab) isSharedTabGroupDataSpecifics_Entity() {}

type SharedTabGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The displayed title of the group, shown on the tab group.
	Title *string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	// The color of the tab group, mapped to tab_groups::TabGroupColorId.
	Color *SharedTabGroup_Color `protobuf:"varint,2,opt,name=color,enum=sync_pb.SharedTabGroup_Color" json:"color,omitempty"`
}

func (x *SharedTabGroup) Reset() {
	*x = SharedTabGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_tab_group_data_specifics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedTabGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedTabGroup) ProtoMessage() {}

func (x *SharedTabGroup) ProtoReflect() protoreflect.Message {
	mi := &file_shared_tab_group_data_specifics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedTabGroup.ProtoReflect.Descriptor instead.
func (*SharedTabGroup) Descriptor() ([]byte, []int) {
	return file_shared_tab_group_data_specifics_proto_rawDescGZIP(), []int{1}
}

func (x *SharedTabGroup) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *SharedTabGroup) GetColor() SharedTabGroup_Color {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return SharedTabGroup_UNSPECIFIED
}

type SharedTab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The URL of the page.
	Url *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// The title of the page.
	Title *string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// The URL to the favicon of the page.
	FaviconUrl *string `protobuf:"bytes,3,opt,name=favicon_url,json=faviconUrl" json:"favicon_url,omitempty"`
	// The GUID of the SharedTabGroup this is a member of.
	SharedTabGroupGuid *string `protobuf:"bytes,4,opt,name=shared_tab_group_guid,json=sharedTabGroupGuid" json:"shared_tab_group_guid,omitempty"`
	// The UniquePosition within a SharedTabGroup.
	UniquePosition *UniquePosition `protobuf:"bytes,5,opt,name=unique_position,json=uniquePosition" json:"unique_position,omitempty"`
}

func (x *SharedTab) Reset() {
	*x = SharedTab{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_tab_group_data_specifics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedTab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedTab) ProtoMessage() {}

func (x *SharedTab) ProtoReflect() protoreflect.Message {
	mi := &file_shared_tab_group_data_specifics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedTab.ProtoReflect.Descriptor instead.
func (*SharedTab) Descriptor() ([]byte, []int) {
	return file_shared_tab_group_data_specifics_proto_rawDescGZIP(), []int{2}
}

func (x *SharedTab) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *SharedTab) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *SharedTab) GetFaviconUrl() string {
	if x != nil && x.FaviconUrl != nil {
		return *x.FaviconUrl
	}
	return ""
}

func (x *SharedTab) GetSharedTabGroupGuid() string {
	if x != nil && x.SharedTabGroupGuid != nil {
		return *x.SharedTabGroupGuid
	}
	return ""
}

func (x *SharedTab) GetUniquePosition() *UniquePosition {
	if x != nil {
		return x.UniquePosition
	}
	return nil
}

var File_shared_tab_group_data_specifics_proto protoreflect.FileDescriptor

var file_shared_tab_group_data_specifics_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x62, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x1a, 0x15, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x1b, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x54, 0x61, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x61, 0x62, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x48, 0x00, 0x52, 0x08, 0x74, 0x61, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x0a,
	0x03, 0x74, 0x61, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x61, 0x62, 0x48, 0x00,
	0x52, 0x03, 0x74, 0x61, 0x62, 0x42, 0x08, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0xd5, 0x01, 0x0a, 0x0e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x61, 0x62, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70,
	0x62, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x61, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x78, 0x0a,
	0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x52, 0x45, 0x59, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x52,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x04,
	0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x49, 0x4e, 0x4b, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x52, 0x50, 0x4c, 0x45, 0x10,
	0x07, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x59, 0x41, 0x4e, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x4f,
	0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x09, 0x22, 0xc9, 0x01, 0x0a, 0x09, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x54, 0x61, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x31,
	0x0a, 0x15, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x62, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x61, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x75, 0x69,
	0x64, 0x12, 0x40, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x36, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
	0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
}

var (
	file_shared_tab_group_data_specifics_proto_rawDescOnce sync.Once
	file_shared_tab_group_data_specifics_proto_rawDescData = file_shared_tab_group_data_specifics_proto_rawDesc
)

func file_shared_tab_group_data_specifics_proto_rawDescGZIP() []byte {
	file_shared_tab_group_data_specifics_proto_rawDescOnce.Do(func() {
		file_shared_tab_group_data_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_tab_group_data_specifics_proto_rawDescData)
	})
	return file_shared_tab_group_data_specifics_proto_rawDescData
}

var file_shared_tab_group_data_specifics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shared_tab_group_data_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shared_tab_group_data_specifics_proto_goTypes = []interface{}{
	(SharedTabGroup_Color)(0),           // 0: sync_pb.SharedTabGroup.Color
	(*SharedTabGroupDataSpecifics)(nil), // 1: sync_pb.SharedTabGroupDataSpecifics
	(*SharedTabGroup)(nil),              // 2: sync_pb.SharedTabGroup
	(*SharedTab)(nil),                   // 3: sync_pb.SharedTab
	(*UniquePosition)(nil),              // 4: sync_pb.UniquePosition
}
var file_shared_tab_group_data_specifics_proto_depIdxs = []int32{
	2, // 0: sync_pb.SharedTabGroupDataSpecifics.tab_group:type_name -> sync_pb.SharedTabGroup
	3, // 1: sync_pb.SharedTabGroupDataSpecifics.tab:type_name -> sync_pb.SharedTab
	0, // 2: sync_pb.SharedTabGroup.color:type_name -> sync_pb.SharedTabGroup.Color
	4, // 3: sync_pb.SharedTab.unique_position:type_name -> sync_pb.UniquePosition
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_shared_tab_group_data_specifics_proto_init() }
func file_shared_tab_group_data_specifics_proto_init() {
	if File_shared_tab_group_data_specifics_proto != nil {
		return
	}
	file_unique_position_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shared_tab_group_data_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedTabGroupDataSpecifics); i {
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
		file_shared_tab_group_data_specifics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedTabGroup); i {
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
		file_shared_tab_group_data_specifics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedTab); i {
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
	file_shared_tab_group_data_specifics_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SharedTabGroupDataSpecifics_TabGroup)(nil),
		(*SharedTabGroupDataSpecifics_Tab)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_tab_group_data_specifics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_tab_group_data_specifics_proto_goTypes,
		DependencyIndexes: file_shared_tab_group_data_specifics_proto_depIdxs,
		EnumInfos:         file_shared_tab_group_data_specifics_proto_enumTypes,
		MessageInfos:      file_shared_tab_group_data_specifics_proto_msgTypes,
	}.Build()
	File_shared_tab_group_data_specifics_proto = out.File
	file_shared_tab_group_data_specifics_proto_rawDesc = nil
	file_shared_tab_group_data_specifics_proto_goTypes = nil
	file_shared_tab_group_data_specifics_proto_depIdxs = nil
}
