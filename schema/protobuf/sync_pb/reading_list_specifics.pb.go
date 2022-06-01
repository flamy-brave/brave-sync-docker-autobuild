// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for the reading list items.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: reading_list_specifics.proto

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

type ReadingListSpecifics_ReadingListEntryStatus int32

const (
	ReadingListSpecifics_UNREAD ReadingListSpecifics_ReadingListEntryStatus = 0
	ReadingListSpecifics_READ   ReadingListSpecifics_ReadingListEntryStatus = 1
	ReadingListSpecifics_UNSEEN ReadingListSpecifics_ReadingListEntryStatus = 2
)

// Enum value maps for ReadingListSpecifics_ReadingListEntryStatus.
var (
	ReadingListSpecifics_ReadingListEntryStatus_name = map[int32]string{
		0: "UNREAD",
		1: "READ",
		2: "UNSEEN",
	}
	ReadingListSpecifics_ReadingListEntryStatus_value = map[string]int32{
		"UNREAD": 0,
		"READ":   1,
		"UNSEEN": 2,
	}
)

func (x ReadingListSpecifics_ReadingListEntryStatus) Enum() *ReadingListSpecifics_ReadingListEntryStatus {
	p := new(ReadingListSpecifics_ReadingListEntryStatus)
	*p = x
	return p
}

func (x ReadingListSpecifics_ReadingListEntryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReadingListSpecifics_ReadingListEntryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_reading_list_specifics_proto_enumTypes[0].Descriptor()
}

func (ReadingListSpecifics_ReadingListEntryStatus) Type() protoreflect.EnumType {
	return &file_reading_list_specifics_proto_enumTypes[0]
}

func (x ReadingListSpecifics_ReadingListEntryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ReadingListSpecifics_ReadingListEntryStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ReadingListSpecifics_ReadingListEntryStatus(num)
	return nil
}

// Deprecated: Use ReadingListSpecifics_ReadingListEntryStatus.Descriptor instead.
func (ReadingListSpecifics_ReadingListEntryStatus) EnumDescriptor() ([]byte, []int) {
	return file_reading_list_specifics_proto_rawDescGZIP(), []int{0, 0}
}

// Sync Reading list entry. This proto contains the fields synced for a reading
// list entry. It must be kept synced with the reading_list.ReadingListLocal
// protobuf.
type ReadingListSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId                  *string `protobuf:"bytes,1,opt,name=entry_id,json=entryId" json:"entry_id,omitempty"`
	Title                    *string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Url                      *string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	CreationTimeUs           *int64  `protobuf:"varint,4,opt,name=creation_time_us,json=creationTimeUs" json:"creation_time_us,omitempty"`
	UpdateTimeUs             *int64  `protobuf:"varint,5,opt,name=update_time_us,json=updateTimeUs" json:"update_time_us,omitempty"`
	FirstReadTimeUs          *int64  `protobuf:"varint,7,opt,name=first_read_time_us,json=firstReadTimeUs" json:"first_read_time_us,omitempty"`
	UpdateTitleTimeUs        *int64  `protobuf:"varint,8,opt,name=update_title_time_us,json=updateTitleTimeUs" json:"update_title_time_us,omitempty"`
	EstimatedReadTimeSeconds *int32  `protobuf:"varint,9,opt,name=estimated_read_time_seconds,json=estimatedReadTimeSeconds" json:"estimated_read_time_seconds,omitempty"`
	// If the field is not present, it defaults to UNSEEN.
	Status *ReadingListSpecifics_ReadingListEntryStatus `protobuf:"varint,6,opt,name=status,enum=sync_pb.ReadingListSpecifics_ReadingListEntryStatus" json:"status,omitempty"`
}

func (x *ReadingListSpecifics) Reset() {
	*x = ReadingListSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reading_list_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadingListSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadingListSpecifics) ProtoMessage() {}

func (x *ReadingListSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_reading_list_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadingListSpecifics.ProtoReflect.Descriptor instead.
func (*ReadingListSpecifics) Descriptor() ([]byte, []int) {
	return file_reading_list_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *ReadingListSpecifics) GetEntryId() string {
	if x != nil && x.EntryId != nil {
		return *x.EntryId
	}
	return ""
}

func (x *ReadingListSpecifics) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *ReadingListSpecifics) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *ReadingListSpecifics) GetCreationTimeUs() int64 {
	if x != nil && x.CreationTimeUs != nil {
		return *x.CreationTimeUs
	}
	return 0
}

func (x *ReadingListSpecifics) GetUpdateTimeUs() int64 {
	if x != nil && x.UpdateTimeUs != nil {
		return *x.UpdateTimeUs
	}
	return 0
}

func (x *ReadingListSpecifics) GetFirstReadTimeUs() int64 {
	if x != nil && x.FirstReadTimeUs != nil {
		return *x.FirstReadTimeUs
	}
	return 0
}

func (x *ReadingListSpecifics) GetUpdateTitleTimeUs() int64 {
	if x != nil && x.UpdateTitleTimeUs != nil {
		return *x.UpdateTitleTimeUs
	}
	return 0
}

func (x *ReadingListSpecifics) GetEstimatedReadTimeSeconds() int32 {
	if x != nil && x.EstimatedReadTimeSeconds != nil {
		return *x.EstimatedReadTimeSeconds
	}
	return 0
}

func (x *ReadingListSpecifics) GetStatus() ReadingListSpecifics_ReadingListEntryStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ReadingListSpecifics_UNREAD
}

var File_reading_list_specifics_proto protoreflect.FileDescriptor

var file_reading_list_specifics_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0xd0, 0x03, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x24, 0x0a,
	0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x55, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73,
	0x12, 0x2f, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55,
	0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x4c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x34, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a,
	0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x52, 0x45,
	0x41, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x4e, 0x53, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72,
	0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_reading_list_specifics_proto_rawDescOnce sync.Once
	file_reading_list_specifics_proto_rawDescData = file_reading_list_specifics_proto_rawDesc
)

func file_reading_list_specifics_proto_rawDescGZIP() []byte {
	file_reading_list_specifics_proto_rawDescOnce.Do(func() {
		file_reading_list_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_reading_list_specifics_proto_rawDescData)
	})
	return file_reading_list_specifics_proto_rawDescData
}

var file_reading_list_specifics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_reading_list_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_reading_list_specifics_proto_goTypes = []interface{}{
	(ReadingListSpecifics_ReadingListEntryStatus)(0), // 0: sync_pb.ReadingListSpecifics.ReadingListEntryStatus
	(*ReadingListSpecifics)(nil),                     // 1: sync_pb.ReadingListSpecifics
}
var file_reading_list_specifics_proto_depIdxs = []int32{
	0, // 0: sync_pb.ReadingListSpecifics.status:type_name -> sync_pb.ReadingListSpecifics.ReadingListEntryStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reading_list_specifics_proto_init() }
func file_reading_list_specifics_proto_init() {
	if File_reading_list_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reading_list_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadingListSpecifics); i {
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
			RawDescriptor: file_reading_list_specifics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reading_list_specifics_proto_goTypes,
		DependencyIndexes: file_reading_list_specifics_proto_depIdxs,
		EnumInfos:         file_reading_list_specifics_proto_enumTypes,
		MessageInfos:      file_reading_list_specifics_proto_msgTypes,
	}.Build()
	File_reading_list_specifics_proto = out.File
	file_reading_list_specifics_proto_rawDesc = nil
	file_reading_list_specifics_proto_goTypes = nil
	file_reading_list_specifics_proto_depIdxs = nil
}
