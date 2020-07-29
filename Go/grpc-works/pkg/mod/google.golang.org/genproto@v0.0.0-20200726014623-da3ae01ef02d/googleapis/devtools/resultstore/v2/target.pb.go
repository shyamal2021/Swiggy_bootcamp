// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/devtools/resultstore/v2/target.proto

package resultstore

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
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

// These correspond to the suffix of the rule name. Eg cc_test has type TEST.
type TargetType int32

const (
	// Unspecified by the build system.
	TargetType_TARGET_TYPE_UNSPECIFIED TargetType = 0
	// An application e.g. ios_application.
	TargetType_APPLICATION TargetType = 1
	// A binary target e.g. cc_binary.
	TargetType_BINARY TargetType = 2
	// A library target e.g. java_library
	TargetType_LIBRARY TargetType = 3
	// A package
	TargetType_PACKAGE TargetType = 4
	// Any test target, in bazel that means a rule with a '_test' suffix.
	TargetType_TEST TargetType = 5
)

// Enum value maps for TargetType.
var (
	TargetType_name = map[int32]string{
		0: "TARGET_TYPE_UNSPECIFIED",
		1: "APPLICATION",
		2: "BINARY",
		3: "LIBRARY",
		4: "PACKAGE",
		5: "TEST",
	}
	TargetType_value = map[string]int32{
		"TARGET_TYPE_UNSPECIFIED": 0,
		"APPLICATION":             1,
		"BINARY":                  2,
		"LIBRARY":                 3,
		"PACKAGE":                 4,
		"TEST":                    5,
	}
)

func (x TargetType) Enum() *TargetType {
	p := new(TargetType)
	*p = x
	return p
}

func (x TargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_resultstore_v2_target_proto_enumTypes[0].Descriptor()
}

func (TargetType) Type() protoreflect.EnumType {
	return &file_google_devtools_resultstore_v2_target_proto_enumTypes[0]
}

func (x TargetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetType.Descriptor instead.
func (TargetType) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_target_proto_rawDescGZIP(), []int{0}
}

// Indicates how big the user indicated the test action was.
type TestSize int32

const (
	// Unspecified by the user.
	TestSize_TEST_SIZE_UNSPECIFIED TestSize = 0
	// Unit test taking less than 1 minute.
	TestSize_SMALL TestSize = 1
	// Integration tests taking less than 5 minutes.
	TestSize_MEDIUM TestSize = 2
	// End-to-end tests taking less than 15 minutes.
	TestSize_LARGE TestSize = 3
	// Even bigger than LARGE.
	TestSize_ENORMOUS TestSize = 4
	// Something that doesn't fit into the above categories.
	TestSize_OTHER_SIZE TestSize = 5
)

// Enum value maps for TestSize.
var (
	TestSize_name = map[int32]string{
		0: "TEST_SIZE_UNSPECIFIED",
		1: "SMALL",
		2: "MEDIUM",
		3: "LARGE",
		4: "ENORMOUS",
		5: "OTHER_SIZE",
	}
	TestSize_value = map[string]int32{
		"TEST_SIZE_UNSPECIFIED": 0,
		"SMALL":                 1,
		"MEDIUM":                2,
		"LARGE":                 3,
		"ENORMOUS":              4,
		"OTHER_SIZE":            5,
	}
)

func (x TestSize) Enum() *TestSize {
	p := new(TestSize)
	*p = x
	return p
}

func (x TestSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestSize) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_resultstore_v2_target_proto_enumTypes[1].Descriptor()
}

func (TestSize) Type() protoreflect.EnumType {
	return &file_google_devtools_resultstore_v2_target_proto_enumTypes[1]
}

func (x TestSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestSize.Descriptor instead.
func (TestSize) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_target_proto_rawDescGZIP(), []int{1}
}

// Each Target represents data for a given target in a given Invocation.
// ConfiguredTarget and Action resources under each Target contain the bulk of
// the data.
type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name.  Its format must be:
	// invocations/${INVOCATION_ID}/targets/${url_encode(TARGET_ID)}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the Target. They must match the
	// resource name after proper encoding.
	Id *Target_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// This is the aggregate status of the target.
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// When this target started and its duration.
	Timing *Timing `protobuf:"bytes,4,opt,name=timing,proto3" json:"timing,omitempty"`
	// Attributes that apply to all targets.
	TargetAttributes *TargetAttributes `protobuf:"bytes,5,opt,name=target_attributes,json=targetAttributes,proto3" json:"target_attributes,omitempty"`
	// Attributes that apply to all test actions under this target.
	TestAttributes *TestAttributes `protobuf:"bytes,6,opt,name=test_attributes,json=testAttributes,proto3" json:"test_attributes,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// A list of file references for target level files.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	// Use this field to specify outputs not related to a configuration.
	Files []*File `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
	// Provides a hint to clients as to whether to display the Target to users.
	// If true then clients likely want to display the Target by default.
	// Once set to true, this may not be set back to false.
	Visible bool `protobuf:"varint,10,opt,name=visible,proto3" json:"visible,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_target_proto_rawDescGZIP(), []int{0}
}

func (x *Target) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Target) GetId() *Target_Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Target) GetStatusAttributes() *StatusAttributes {
	if x != nil {
		return x.StatusAttributes
	}
	return nil
}

func (x *Target) GetTiming() *Timing {
	if x != nil {
		return x.Timing
	}
	return nil
}

func (x *Target) GetTargetAttributes() *TargetAttributes {
	if x != nil {
		return x.TargetAttributes
	}
	return nil
}

func (x *Target) GetTestAttributes() *TestAttributes {
	if x != nil {
		return x.TestAttributes
	}
	return nil
}

func (x *Target) GetProperties() []*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Target) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Target) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

// Attributes that apply to all targets.
type TargetAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If known, indicates the type of this target.  In bazel this corresponds
	// to the rule-suffix.
	Type TargetType `protobuf:"varint,1,opt,name=type,proto3,enum=google.devtools.resultstore.v2.TargetType" json:"type,omitempty"`
	// If known, the main language of this target, e.g. java, cc, python, etc.
	Language Language `protobuf:"varint,2,opt,name=language,proto3,enum=google.devtools.resultstore.v2.Language" json:"language,omitempty"`
	// The tags attribute of the build rule. These should be short, descriptive
	// words, and there should only be a few of them.
	// This is implemented as a set. All tags will be unique. Any duplicate tags
	// will be ignored. Tags will be returned in lexicographical order.
	Tags []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *TargetAttributes) Reset() {
	*x = TargetAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_target_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetAttributes) ProtoMessage() {}

func (x *TargetAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_target_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetAttributes.ProtoReflect.Descriptor instead.
func (*TargetAttributes) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_target_proto_rawDescGZIP(), []int{1}
}

func (x *TargetAttributes) GetType() TargetType {
	if x != nil {
		return x.Type
	}
	return TargetType_TARGET_TYPE_UNSPECIFIED
}

func (x *TargetAttributes) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_LANGUAGE_UNSPECIFIED
}

func (x *TargetAttributes) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// Attributes that apply only to test actions under this target.
type TestAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates how big the user indicated the test action was.
	Size TestSize `protobuf:"varint,1,opt,name=size,proto3,enum=google.devtools.resultstore.v2.TestSize" json:"size,omitempty"`
}

func (x *TestAttributes) Reset() {
	*x = TestAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_target_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAttributes) ProtoMessage() {}

func (x *TestAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_target_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAttributes.ProtoReflect.Descriptor instead.
func (*TestAttributes) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_target_proto_rawDescGZIP(), []int{2}
}

func (x *TestAttributes) GetSize() TestSize {
	if x != nil {
		return x.Size
	}
	return TestSize_TEST_SIZE_UNSPECIFIED
}

// The resource ID components that identify the Target.
type Target_Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Invocation ID.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The Target ID.
	TargetId string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
}

func (x *Target_Id) Reset() {
	*x = Target_Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_target_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target_Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target_Id) ProtoMessage() {}

func (x *Target_Id) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_target_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target_Id.ProtoReflect.Descriptor instead.
func (*Target_Id) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_target_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Target_Id) GetInvocationId() string {
	if x != nil {
		return x.InvocationId
	}
	return ""
}

func (x *Target_Id) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

var File_google_devtools_resultstore_v2_target_proto protoreflect.FileDescriptor

var file_google_devtools_resultstore_v2_target_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x2b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x05, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x5d, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x10, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3e,
	0x0a, 0x06, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x54, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x5d,
	0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x10, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x57, 0x0a,
	0x0f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x3a, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x1a, 0x46, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0xac,
	0x01, 0x0a, 0x10, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x4e, 0x0a,
	0x0e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x3c, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x2a, 0x6a, 0x0a,
	0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54,
	0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x50, 0x4c,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e,
	0x41, 0x52, 0x59, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x52, 0x59,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x10, 0x04, 0x12,
	0x08, 0x0a, 0x04, 0x54, 0x45, 0x53, 0x54, 0x10, 0x05, 0x2a, 0x65, 0x0a, 0x08, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x49,
	0x5a, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x41, 0x52, 0x47, 0x45,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x4f, 0x52, 0x4d, 0x4f, 0x55, 0x53, 0x10, 0x04,
	0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x05,
	0x42, 0x71, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_resultstore_v2_target_proto_rawDescOnce sync.Once
	file_google_devtools_resultstore_v2_target_proto_rawDescData = file_google_devtools_resultstore_v2_target_proto_rawDesc
)

func file_google_devtools_resultstore_v2_target_proto_rawDescGZIP() []byte {
	file_google_devtools_resultstore_v2_target_proto_rawDescOnce.Do(func() {
		file_google_devtools_resultstore_v2_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_resultstore_v2_target_proto_rawDescData)
	})
	return file_google_devtools_resultstore_v2_target_proto_rawDescData
}

var file_google_devtools_resultstore_v2_target_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_devtools_resultstore_v2_target_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_devtools_resultstore_v2_target_proto_goTypes = []interface{}{
	(TargetType)(0),          // 0: google.devtools.resultstore.v2.TargetType
	(TestSize)(0),            // 1: google.devtools.resultstore.v2.TestSize
	(*Target)(nil),           // 2: google.devtools.resultstore.v2.Target
	(*TargetAttributes)(nil), // 3: google.devtools.resultstore.v2.TargetAttributes
	(*TestAttributes)(nil),   // 4: google.devtools.resultstore.v2.TestAttributes
	(*Target_Id)(nil),        // 5: google.devtools.resultstore.v2.Target.Id
	(*StatusAttributes)(nil), // 6: google.devtools.resultstore.v2.StatusAttributes
	(*Timing)(nil),           // 7: google.devtools.resultstore.v2.Timing
	(*Property)(nil),         // 8: google.devtools.resultstore.v2.Property
	(*File)(nil),             // 9: google.devtools.resultstore.v2.File
	(Language)(0),            // 10: google.devtools.resultstore.v2.Language
}
var file_google_devtools_resultstore_v2_target_proto_depIdxs = []int32{
	5,  // 0: google.devtools.resultstore.v2.Target.id:type_name -> google.devtools.resultstore.v2.Target.Id
	6,  // 1: google.devtools.resultstore.v2.Target.status_attributes:type_name -> google.devtools.resultstore.v2.StatusAttributes
	7,  // 2: google.devtools.resultstore.v2.Target.timing:type_name -> google.devtools.resultstore.v2.Timing
	3,  // 3: google.devtools.resultstore.v2.Target.target_attributes:type_name -> google.devtools.resultstore.v2.TargetAttributes
	4,  // 4: google.devtools.resultstore.v2.Target.test_attributes:type_name -> google.devtools.resultstore.v2.TestAttributes
	8,  // 5: google.devtools.resultstore.v2.Target.properties:type_name -> google.devtools.resultstore.v2.Property
	9,  // 6: google.devtools.resultstore.v2.Target.files:type_name -> google.devtools.resultstore.v2.File
	0,  // 7: google.devtools.resultstore.v2.TargetAttributes.type:type_name -> google.devtools.resultstore.v2.TargetType
	10, // 8: google.devtools.resultstore.v2.TargetAttributes.language:type_name -> google.devtools.resultstore.v2.Language
	1,  // 9: google.devtools.resultstore.v2.TestAttributes.size:type_name -> google.devtools.resultstore.v2.TestSize
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_devtools_resultstore_v2_target_proto_init() }
func file_google_devtools_resultstore_v2_target_proto_init() {
	if File_google_devtools_resultstore_v2_target_proto != nil {
		return
	}
	file_google_devtools_resultstore_v2_common_proto_init()
	file_google_devtools_resultstore_v2_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_resultstore_v2_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
		file_google_devtools_resultstore_v2_target_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetAttributes); i {
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
		file_google_devtools_resultstore_v2_target_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestAttributes); i {
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
		file_google_devtools_resultstore_v2_target_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target_Id); i {
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
			RawDescriptor: file_google_devtools_resultstore_v2_target_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_resultstore_v2_target_proto_goTypes,
		DependencyIndexes: file_google_devtools_resultstore_v2_target_proto_depIdxs,
		EnumInfos:         file_google_devtools_resultstore_v2_target_proto_enumTypes,
		MessageInfos:      file_google_devtools_resultstore_v2_target_proto_msgTypes,
	}.Build()
	File_google_devtools_resultstore_v2_target_proto = out.File
	file_google_devtools_resultstore_v2_target_proto_rawDesc = nil
	file_google_devtools_resultstore_v2_target_proto_goTypes = nil
	file_google_devtools_resultstore_v2_target_proto_depIdxs = nil
}
