// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// LINT: ALLOW_GROUPS

// Benchmark messages for proto2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datasets/google_message2/benchmark_message2.proto

package google_message2

import (
	protoreflect "github.com/ikaiguang/protoc-gen-go/reflect/protoreflect"
	protoimpl "github.com/ikaiguang/protoc-gen-go/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type GoogleMessage2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1   *string                  `protobuf:"bytes,1,opt,name=field1" json:"field1,omitempty"`
	Field3   *int64                   `protobuf:"varint,3,opt,name=field3" json:"field3,omitempty"`
	Field4   *int64                   `protobuf:"varint,4,opt,name=field4" json:"field4,omitempty"`
	Field30  *int64                   `protobuf:"varint,30,opt,name=field30" json:"field30,omitempty"`
	Field75  *bool                    `protobuf:"varint,75,opt,name=field75,def=0" json:"field75,omitempty"`
	Field6   *string                  `protobuf:"bytes,6,opt,name=field6" json:"field6,omitempty"`
	Field2   []byte                   `protobuf:"bytes,2,opt,name=field2" json:"field2,omitempty"`
	Field21  *int32                   `protobuf:"varint,21,opt,name=field21,def=0" json:"field21,omitempty"`
	Field71  *int32                   `protobuf:"varint,71,opt,name=field71" json:"field71,omitempty"`
	Field25  *float32                 `protobuf:"fixed32,25,opt,name=field25" json:"field25,omitempty"`
	Field109 *int32                   `protobuf:"varint,109,opt,name=field109,def=0" json:"field109,omitempty"`
	Field210 *int32                   `protobuf:"varint,210,opt,name=field210,def=0" json:"field210,omitempty"`
	Field211 *int32                   `protobuf:"varint,211,opt,name=field211,def=0" json:"field211,omitempty"`
	Field212 *int32                   `protobuf:"varint,212,opt,name=field212,def=0" json:"field212,omitempty"`
	Field213 *int32                   `protobuf:"varint,213,opt,name=field213,def=0" json:"field213,omitempty"`
	Field216 *int32                   `protobuf:"varint,216,opt,name=field216,def=0" json:"field216,omitempty"`
	Field217 *int32                   `protobuf:"varint,217,opt,name=field217,def=0" json:"field217,omitempty"`
	Field218 *int32                   `protobuf:"varint,218,opt,name=field218,def=0" json:"field218,omitempty"`
	Field220 *int32                   `protobuf:"varint,220,opt,name=field220,def=0" json:"field220,omitempty"`
	Field221 *int32                   `protobuf:"varint,221,opt,name=field221,def=0" json:"field221,omitempty"`
	Field222 *float32                 `protobuf:"fixed32,222,opt,name=field222,def=0" json:"field222,omitempty"`
	Field63  *int32                   `protobuf:"varint,63,opt,name=field63" json:"field63,omitempty"`
	Group1   []*GoogleMessage2_Group1 `protobuf:"group,10,rep,name=Group1,json=group1" json:"group1,omitempty"`
	Field128 []string                 `protobuf:"bytes,128,rep,name=field128" json:"field128,omitempty"`
	Field131 *int64                   `protobuf:"varint,131,opt,name=field131" json:"field131,omitempty"`
	Field127 []string                 `protobuf:"bytes,127,rep,name=field127" json:"field127,omitempty"`
	Field129 *int32                   `protobuf:"varint,129,opt,name=field129" json:"field129,omitempty"`
	Field130 []int64                  `protobuf:"varint,130,rep,name=field130" json:"field130,omitempty"`
	Field205 *bool                    `protobuf:"varint,205,opt,name=field205,def=0" json:"field205,omitempty"`
	Field206 *bool                    `protobuf:"varint,206,opt,name=field206,def=0" json:"field206,omitempty"`
}

// Default values for GoogleMessage2 fields.
const (
	Default_GoogleMessage2_Field75  = bool(false)
	Default_GoogleMessage2_Field21  = int32(0)
	Default_GoogleMessage2_Field109 = int32(0)
	Default_GoogleMessage2_Field210 = int32(0)
	Default_GoogleMessage2_Field211 = int32(0)
	Default_GoogleMessage2_Field212 = int32(0)
	Default_GoogleMessage2_Field213 = int32(0)
	Default_GoogleMessage2_Field216 = int32(0)
	Default_GoogleMessage2_Field217 = int32(0)
	Default_GoogleMessage2_Field218 = int32(0)
	Default_GoogleMessage2_Field220 = int32(0)
	Default_GoogleMessage2_Field221 = int32(0)
	Default_GoogleMessage2_Field222 = float32(0)
	Default_GoogleMessage2_Field205 = bool(false)
	Default_GoogleMessage2_Field206 = bool(false)
)

func (x *GoogleMessage2) Reset() {
	*x = GoogleMessage2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message2_benchmark_message2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleMessage2) ProtoMessage() {}

func (x *GoogleMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message2_benchmark_message2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleMessage2.ProtoReflect.Descriptor instead.
func (*GoogleMessage2) Descriptor() ([]byte, []int) {
	return file_datasets_google_message2_benchmark_message2_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleMessage2) GetField1() string {
	if x != nil && x.Field1 != nil {
		return *x.Field1
	}
	return ""
}

func (x *GoogleMessage2) GetField3() int64 {
	if x != nil && x.Field3 != nil {
		return *x.Field3
	}
	return 0
}

func (x *GoogleMessage2) GetField4() int64 {
	if x != nil && x.Field4 != nil {
		return *x.Field4
	}
	return 0
}

func (x *GoogleMessage2) GetField30() int64 {
	if x != nil && x.Field30 != nil {
		return *x.Field30
	}
	return 0
}

func (x *GoogleMessage2) GetField75() bool {
	if x != nil && x.Field75 != nil {
		return *x.Field75
	}
	return Default_GoogleMessage2_Field75
}

func (x *GoogleMessage2) GetField6() string {
	if x != nil && x.Field6 != nil {
		return *x.Field6
	}
	return ""
}

func (x *GoogleMessage2) GetField2() []byte {
	if x != nil {
		return x.Field2
	}
	return nil
}

func (x *GoogleMessage2) GetField21() int32 {
	if x != nil && x.Field21 != nil {
		return *x.Field21
	}
	return Default_GoogleMessage2_Field21
}

func (x *GoogleMessage2) GetField71() int32 {
	if x != nil && x.Field71 != nil {
		return *x.Field71
	}
	return 0
}

func (x *GoogleMessage2) GetField25() float32 {
	if x != nil && x.Field25 != nil {
		return *x.Field25
	}
	return 0
}

func (x *GoogleMessage2) GetField109() int32 {
	if x != nil && x.Field109 != nil {
		return *x.Field109
	}
	return Default_GoogleMessage2_Field109
}

func (x *GoogleMessage2) GetField210() int32 {
	if x != nil && x.Field210 != nil {
		return *x.Field210
	}
	return Default_GoogleMessage2_Field210
}

func (x *GoogleMessage2) GetField211() int32 {
	if x != nil && x.Field211 != nil {
		return *x.Field211
	}
	return Default_GoogleMessage2_Field211
}

func (x *GoogleMessage2) GetField212() int32 {
	if x != nil && x.Field212 != nil {
		return *x.Field212
	}
	return Default_GoogleMessage2_Field212
}

func (x *GoogleMessage2) GetField213() int32 {
	if x != nil && x.Field213 != nil {
		return *x.Field213
	}
	return Default_GoogleMessage2_Field213
}

func (x *GoogleMessage2) GetField216() int32 {
	if x != nil && x.Field216 != nil {
		return *x.Field216
	}
	return Default_GoogleMessage2_Field216
}

func (x *GoogleMessage2) GetField217() int32 {
	if x != nil && x.Field217 != nil {
		return *x.Field217
	}
	return Default_GoogleMessage2_Field217
}

func (x *GoogleMessage2) GetField218() int32 {
	if x != nil && x.Field218 != nil {
		return *x.Field218
	}
	return Default_GoogleMessage2_Field218
}

func (x *GoogleMessage2) GetField220() int32 {
	if x != nil && x.Field220 != nil {
		return *x.Field220
	}
	return Default_GoogleMessage2_Field220
}

func (x *GoogleMessage2) GetField221() int32 {
	if x != nil && x.Field221 != nil {
		return *x.Field221
	}
	return Default_GoogleMessage2_Field221
}

func (x *GoogleMessage2) GetField222() float32 {
	if x != nil && x.Field222 != nil {
		return *x.Field222
	}
	return Default_GoogleMessage2_Field222
}

func (x *GoogleMessage2) GetField63() int32 {
	if x != nil && x.Field63 != nil {
		return *x.Field63
	}
	return 0
}

func (x *GoogleMessage2) GetGroup1() []*GoogleMessage2_Group1 {
	if x != nil {
		return x.Group1
	}
	return nil
}

func (x *GoogleMessage2) GetField128() []string {
	if x != nil {
		return x.Field128
	}
	return nil
}

func (x *GoogleMessage2) GetField131() int64 {
	if x != nil && x.Field131 != nil {
		return *x.Field131
	}
	return 0
}

func (x *GoogleMessage2) GetField127() []string {
	if x != nil {
		return x.Field127
	}
	return nil
}

func (x *GoogleMessage2) GetField129() int32 {
	if x != nil && x.Field129 != nil {
		return *x.Field129
	}
	return 0
}

func (x *GoogleMessage2) GetField130() []int64 {
	if x != nil {
		return x.Field130
	}
	return nil
}

func (x *GoogleMessage2) GetField205() bool {
	if x != nil && x.Field205 != nil {
		return *x.Field205
	}
	return Default_GoogleMessage2_Field205
}

func (x *GoogleMessage2) GetField206() bool {
	if x != nil && x.Field206 != nil {
		return *x.Field206
	}
	return Default_GoogleMessage2_Field206
}

type GoogleMessage2GroupedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1  *float32 `protobuf:"fixed32,1,opt,name=field1" json:"field1,omitempty"`
	Field2  *float32 `protobuf:"fixed32,2,opt,name=field2" json:"field2,omitempty"`
	Field3  *float32 `protobuf:"fixed32,3,opt,name=field3,def=0" json:"field3,omitempty"`
	Field4  *bool    `protobuf:"varint,4,opt,name=field4" json:"field4,omitempty"`
	Field5  *bool    `protobuf:"varint,5,opt,name=field5" json:"field5,omitempty"`
	Field6  *bool    `protobuf:"varint,6,opt,name=field6,def=1" json:"field6,omitempty"`
	Field7  *bool    `protobuf:"varint,7,opt,name=field7,def=0" json:"field7,omitempty"`
	Field8  *float32 `protobuf:"fixed32,8,opt,name=field8" json:"field8,omitempty"`
	Field9  *bool    `protobuf:"varint,9,opt,name=field9" json:"field9,omitempty"`
	Field10 *float32 `protobuf:"fixed32,10,opt,name=field10" json:"field10,omitempty"`
	Field11 *int64   `protobuf:"varint,11,opt,name=field11" json:"field11,omitempty"`
}

// Default values for GoogleMessage2GroupedMessage fields.
const (
	Default_GoogleMessage2GroupedMessage_Field3 = float32(0)
	Default_GoogleMessage2GroupedMessage_Field6 = bool(true)
	Default_GoogleMessage2GroupedMessage_Field7 = bool(false)
)

func (x *GoogleMessage2GroupedMessage) Reset() {
	*x = GoogleMessage2GroupedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message2_benchmark_message2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleMessage2GroupedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleMessage2GroupedMessage) ProtoMessage() {}

func (x *GoogleMessage2GroupedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message2_benchmark_message2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleMessage2GroupedMessage.ProtoReflect.Descriptor instead.
func (*GoogleMessage2GroupedMessage) Descriptor() ([]byte, []int) {
	return file_datasets_google_message2_benchmark_message2_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleMessage2GroupedMessage) GetField1() float32 {
	if x != nil && x.Field1 != nil {
		return *x.Field1
	}
	return 0
}

func (x *GoogleMessage2GroupedMessage) GetField2() float32 {
	if x != nil && x.Field2 != nil {
		return *x.Field2
	}
	return 0
}

func (x *GoogleMessage2GroupedMessage) GetField3() float32 {
	if x != nil && x.Field3 != nil {
		return *x.Field3
	}
	return Default_GoogleMessage2GroupedMessage_Field3
}

func (x *GoogleMessage2GroupedMessage) GetField4() bool {
	if x != nil && x.Field4 != nil {
		return *x.Field4
	}
	return false
}

func (x *GoogleMessage2GroupedMessage) GetField5() bool {
	if x != nil && x.Field5 != nil {
		return *x.Field5
	}
	return false
}

func (x *GoogleMessage2GroupedMessage) GetField6() bool {
	if x != nil && x.Field6 != nil {
		return *x.Field6
	}
	return Default_GoogleMessage2GroupedMessage_Field6
}

func (x *GoogleMessage2GroupedMessage) GetField7() bool {
	if x != nil && x.Field7 != nil {
		return *x.Field7
	}
	return Default_GoogleMessage2GroupedMessage_Field7
}

func (x *GoogleMessage2GroupedMessage) GetField8() float32 {
	if x != nil && x.Field8 != nil {
		return *x.Field8
	}
	return 0
}

func (x *GoogleMessage2GroupedMessage) GetField9() bool {
	if x != nil && x.Field9 != nil {
		return *x.Field9
	}
	return false
}

func (x *GoogleMessage2GroupedMessage) GetField10() float32 {
	if x != nil && x.Field10 != nil {
		return *x.Field10
	}
	return 0
}

func (x *GoogleMessage2GroupedMessage) GetField11() int64 {
	if x != nil && x.Field11 != nil {
		return *x.Field11
	}
	return 0
}

type GoogleMessage2_Group1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field11 *float32                      `protobuf:"fixed32,11,req,name=field11" json:"field11,omitempty"`
	Field26 *float32                      `protobuf:"fixed32,26,opt,name=field26" json:"field26,omitempty"`
	Field12 *string                       `protobuf:"bytes,12,opt,name=field12" json:"field12,omitempty"`
	Field13 *string                       `protobuf:"bytes,13,opt,name=field13" json:"field13,omitempty"`
	Field14 []string                      `protobuf:"bytes,14,rep,name=field14" json:"field14,omitempty"`
	Field15 *uint64                       `protobuf:"varint,15,req,name=field15" json:"field15,omitempty"`
	Field5  *int32                        `protobuf:"varint,5,opt,name=field5" json:"field5,omitempty"`
	Field27 *string                       `protobuf:"bytes,27,opt,name=field27" json:"field27,omitempty"`
	Field28 *int32                        `protobuf:"varint,28,opt,name=field28" json:"field28,omitempty"`
	Field29 *string                       `protobuf:"bytes,29,opt,name=field29" json:"field29,omitempty"`
	Field16 *string                       `protobuf:"bytes,16,opt,name=field16" json:"field16,omitempty"`
	Field22 []string                      `protobuf:"bytes,22,rep,name=field22" json:"field22,omitempty"`
	Field73 []int32                       `protobuf:"varint,73,rep,name=field73" json:"field73,omitempty"`
	Field20 *int32                        `protobuf:"varint,20,opt,name=field20,def=0" json:"field20,omitempty"`
	Field24 *string                       `protobuf:"bytes,24,opt,name=field24" json:"field24,omitempty"`
	Field31 *GoogleMessage2GroupedMessage `protobuf:"bytes,31,opt,name=field31" json:"field31,omitempty"`
}

// Default values for GoogleMessage2_Group1 fields.
const (
	Default_GoogleMessage2_Group1_Field20 = int32(0)
)

func (x *GoogleMessage2_Group1) Reset() {
	*x = GoogleMessage2_Group1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message2_benchmark_message2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleMessage2_Group1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleMessage2_Group1) ProtoMessage() {}

func (x *GoogleMessage2_Group1) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message2_benchmark_message2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleMessage2_Group1.ProtoReflect.Descriptor instead.
func (*GoogleMessage2_Group1) Descriptor() ([]byte, []int) {
	return file_datasets_google_message2_benchmark_message2_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GoogleMessage2_Group1) GetField11() float32 {
	if x != nil && x.Field11 != nil {
		return *x.Field11
	}
	return 0
}

func (x *GoogleMessage2_Group1) GetField26() float32 {
	if x != nil && x.Field26 != nil {
		return *x.Field26
	}
	return 0
}

func (x *GoogleMessage2_Group1) GetField12() string {
	if x != nil && x.Field12 != nil {
		return *x.Field12
	}
	return ""
}

func (x *GoogleMessage2_Group1) GetField13() string {
	if x != nil && x.Field13 != nil {
		return *x.Field13
	}
	return ""
}

func (x *GoogleMessage2_Group1) GetField14() []string {
	if x != nil {
		return x.Field14
	}
	return nil
}

func (x *GoogleMessage2_Group1) GetField15() uint64 {
	if x != nil && x.Field15 != nil {
		return *x.Field15
	}
	return 0
}

func (x *GoogleMessage2_Group1) GetField5() int32 {
	if x != nil && x.Field5 != nil {
		return *x.Field5
	}
	return 0
}

func (x *GoogleMessage2_Group1) GetField27() string {
	if x != nil && x.Field27 != nil {
		return *x.Field27
	}
	return ""
}

func (x *GoogleMessage2_Group1) GetField28() int32 {
	if x != nil && x.Field28 != nil {
		return *x.Field28
	}
	return 0
}

func (x *GoogleMessage2_Group1) GetField29() string {
	if x != nil && x.Field29 != nil {
		return *x.Field29
	}
	return ""
}

func (x *GoogleMessage2_Group1) GetField16() string {
	if x != nil && x.Field16 != nil {
		return *x.Field16
	}
	return ""
}

func (x *GoogleMessage2_Group1) GetField22() []string {
	if x != nil {
		return x.Field22
	}
	return nil
}

func (x *GoogleMessage2_Group1) GetField73() []int32 {
	if x != nil {
		return x.Field73
	}
	return nil
}

func (x *GoogleMessage2_Group1) GetField20() int32 {
	if x != nil && x.Field20 != nil {
		return *x.Field20
	}
	return Default_GoogleMessage2_Group1_Field20
}

func (x *GoogleMessage2_Group1) GetField24() string {
	if x != nil && x.Field24 != nil {
		return *x.Field24
	}
	return ""
}

func (x *GoogleMessage2_Group1) GetField31() *GoogleMessage2GroupedMessage {
	if x != nil {
		return x.Field31
	}
	return nil
}

var File_datasets_google_message2_benchmark_message2_proto protoreflect.FileDescriptor

var file_datasets_google_message2_benchmark_message2_proto_rawDesc = []byte{
	0x0a, 0x31, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x22, 0x84, 0x0b, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x34, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x12, 0x1f, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x37, 0x35, 0x18, 0x4b, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x35, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x36, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x12, 0x1b, 0x0a, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x37, 0x31, 0x18, 0x47, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x37, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x35, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x35, 0x12, 0x1d, 0x0a,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x39, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x39, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x30, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x30, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x31, 0x18, 0xd3, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x31, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x32, 0x18, 0xd4, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x32, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x33, 0x18, 0xd5, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x33, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x36, 0x18, 0xd8, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x36, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x37, 0x18, 0xd9, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x37, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x38, 0x18, 0xda, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x38, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x30, 0x18, 0xdc, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x30, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x31, 0x18, 0xdd, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x31, 0x12, 0x1e, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x32, 0x18, 0xde, 0x01, 0x20, 0x01, 0x28, 0x02, 0x3a,
	0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x32, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x33, 0x18, 0x3f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x36, 0x33, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x31,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0a, 0x32, 0x28, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x31,
	0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x31, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x32, 0x38, 0x18, 0x80, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x32, 0x38, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33,
	0x31, 0x18, 0x83, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x33, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x37, 0x18, 0x7f,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x37, 0x12, 0x1b,
	0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x39, 0x18, 0x81, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x39, 0x12, 0x1b, 0x0a, 0x08, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x30, 0x18, 0x82, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x30, 0x12, 0x22, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0x30, 0x35, 0x18, 0xcd, 0x01, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c,
	0x73, 0x65, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x35, 0x12, 0x22, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x36, 0x18, 0xce, 0x01, 0x20, 0x01, 0x28, 0x08, 0x3a,
	0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x36,
	0x1a, 0xda, 0x03, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x02, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x36,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x36, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x33, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x18, 0x0f, 0x20, 0x02, 0x28, 0x04, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x38, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0x38, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x39, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x39, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x32, 0x18, 0x16, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x33, 0x18, 0x49, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x33, 0x12, 0x1b, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0x34, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x34, 0x12, 0x49, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x31, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x31, 0x22, 0xba, 0x02,
	0x0a, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x12, 0x19,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01,
	0x30, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x34, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x12, 0x1c, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x12, 0x1d, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x31, 0x42, 0x25, 0x0a, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x48, 0x01, 0xf8, 0x01,
	0x01,
}

var (
	file_datasets_google_message2_benchmark_message2_proto_rawDescOnce sync.Once
	file_datasets_google_message2_benchmark_message2_proto_rawDescData = file_datasets_google_message2_benchmark_message2_proto_rawDesc
)

func file_datasets_google_message2_benchmark_message2_proto_rawDescGZIP() []byte {
	file_datasets_google_message2_benchmark_message2_proto_rawDescOnce.Do(func() {
		file_datasets_google_message2_benchmark_message2_proto_rawDescData = protoimpl.X.CompressGZIP(file_datasets_google_message2_benchmark_message2_proto_rawDescData)
	})
	return file_datasets_google_message2_benchmark_message2_proto_rawDescData
}

var file_datasets_google_message2_benchmark_message2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datasets_google_message2_benchmark_message2_proto_goTypes = []interface{}{
	(*GoogleMessage2)(nil),               // 0: benchmarks.proto2.GoogleMessage2
	(*GoogleMessage2GroupedMessage)(nil), // 1: benchmarks.proto2.GoogleMessage2GroupedMessage
	(*GoogleMessage2_Group1)(nil),        // 2: benchmarks.proto2.GoogleMessage2.Group1
}
var file_datasets_google_message2_benchmark_message2_proto_depIdxs = []int32{
	2, // 0: benchmarks.proto2.GoogleMessage2.group1:type_name -> benchmarks.proto2.GoogleMessage2.Group1
	1, // 1: benchmarks.proto2.GoogleMessage2.Group1.field31:type_name -> benchmarks.proto2.GoogleMessage2GroupedMessage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_datasets_google_message2_benchmark_message2_proto_init() }
func file_datasets_google_message2_benchmark_message2_proto_init() {
	if File_datasets_google_message2_benchmark_message2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datasets_google_message2_benchmark_message2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleMessage2); i {
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
		file_datasets_google_message2_benchmark_message2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleMessage2GroupedMessage); i {
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
		file_datasets_google_message2_benchmark_message2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleMessage2_Group1); i {
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
			RawDescriptor: file_datasets_google_message2_benchmark_message2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datasets_google_message2_benchmark_message2_proto_goTypes,
		DependencyIndexes: file_datasets_google_message2_benchmark_message2_proto_depIdxs,
		MessageInfos:      file_datasets_google_message2_benchmark_message2_proto_msgTypes,
	}.Build()
	File_datasets_google_message2_benchmark_message2_proto = out.File
	file_datasets_google_message2_benchmark_message2_proto_rawDesc = nil
	file_datasets_google_message2_benchmark_message2_proto_goTypes = nil
	file_datasets_google_message2_benchmark_message2_proto_depIdxs = nil
}
