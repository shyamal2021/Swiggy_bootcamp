// Copyright 2020 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v3/services/customer_negative_criterion_service.proto

package services

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for
// [CustomerNegativeCriterionService.GetCustomerNegativeCriterion][google.ads.googleads.v3.services.CustomerNegativeCriterionService.GetCustomerNegativeCriterion].
type GetCustomerNegativeCriterionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the criterion to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetCustomerNegativeCriterionRequest) Reset() {
	*x = GetCustomerNegativeCriterionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerNegativeCriterionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerNegativeCriterionRequest) ProtoMessage() {}

func (x *GetCustomerNegativeCriterionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerNegativeCriterionRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerNegativeCriterionRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerNegativeCriterionRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for
// [CustomerNegativeCriterionService.MutateCustomerNegativeCriteria][google.ads.googleads.v3.services.CustomerNegativeCriterionService.MutateCustomerNegativeCriteria].
type MutateCustomerNegativeCriteriaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual criteria.
	Operations []*CustomerNegativeCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateCustomerNegativeCriteriaRequest) Reset() {
	*x = MutateCustomerNegativeCriteriaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerNegativeCriteriaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerNegativeCriteriaRequest) ProtoMessage() {}

func (x *MutateCustomerNegativeCriteriaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerNegativeCriteriaRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerNegativeCriteriaRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateCustomerNegativeCriteriaRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerNegativeCriteriaRequest) GetOperations() []*CustomerNegativeCriterionOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateCustomerNegativeCriteriaRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateCustomerNegativeCriteriaRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// A single operation (create or remove) on a customer level negative criterion.
type CustomerNegativeCriterionOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*CustomerNegativeCriterionOperation_Create
	//	*CustomerNegativeCriterionOperation_Remove
	Operation isCustomerNegativeCriterionOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CustomerNegativeCriterionOperation) Reset() {
	*x = CustomerNegativeCriterionOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerNegativeCriterionOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerNegativeCriterionOperation) ProtoMessage() {}

func (x *CustomerNegativeCriterionOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerNegativeCriterionOperation.ProtoReflect.Descriptor instead.
func (*CustomerNegativeCriterionOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescGZIP(), []int{2}
}

func (m *CustomerNegativeCriterionOperation) GetOperation() isCustomerNegativeCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CustomerNegativeCriterionOperation) GetCreate() *resources.CustomerNegativeCriterion {
	if x, ok := x.GetOperation().(*CustomerNegativeCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *CustomerNegativeCriterionOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*CustomerNegativeCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isCustomerNegativeCriterionOperation_Operation interface {
	isCustomerNegativeCriterionOperation_Operation()
}

type CustomerNegativeCriterionOperation_Create struct {
	// Create operation: No resource name is expected for the new criterion.
	Create *resources.CustomerNegativeCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerNegativeCriterionOperation_Remove struct {
	// Remove operation: A resource name for the removed criterion is expected,
	// in this format:
	//
	// `customers/{customer_id}/customerNegativeCriteria/{criterion_id}`
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CustomerNegativeCriterionOperation_Create) isCustomerNegativeCriterionOperation_Operation() {}

func (*CustomerNegativeCriterionOperation_Remove) isCustomerNegativeCriterionOperation_Operation() {}

// Response message for customer negative criterion mutate.
type MutateCustomerNegativeCriteriaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateCustomerNegativeCriteriaResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateCustomerNegativeCriteriaResponse) Reset() {
	*x = MutateCustomerNegativeCriteriaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerNegativeCriteriaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerNegativeCriteriaResponse) ProtoMessage() {}

func (x *MutateCustomerNegativeCriteriaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerNegativeCriteriaResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerNegativeCriteriaResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerNegativeCriteriaResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateCustomerNegativeCriteriaResponse) GetResults() []*MutateCustomerNegativeCriteriaResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCustomerNegativeCriteriaResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCustomerNegativeCriteriaResult) Reset() {
	*x = MutateCustomerNegativeCriteriaResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerNegativeCriteriaResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerNegativeCriteriaResult) ProtoMessage() {}

func (x *MutateCustomerNegativeCriteriaResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerNegativeCriteriaResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerNegativeCriteriaResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateCustomerNegativeCriteriaResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v3_services_customer_negative_criterion_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x43,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x86, 0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x34, 0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x86, 0x02, 0x0a, 0x25, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x69, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c,
	0x79, 0x22, 0xa3, 0x01, 0x0a, 0x22, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd2, 0x01, 0x0a, 0x26, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x60, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x4b, 0x0a, 0x24,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xd2, 0x04, 0x0a, 0x20, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf7,
	0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x12,
	0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x12, 0x3a, 0x2f, 0x76,
	0x33, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x96, 0x02, 0x0a, 0x1e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x47, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x42, 0x22, 0x3d, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x3a, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x1b, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x8c,
	0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x25, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x33, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x33, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescData = file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDesc
)

func file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDescData
}

var file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_goTypes = []interface{}{
	(*GetCustomerNegativeCriterionRequest)(nil),    // 0: google.ads.googleads.v3.services.GetCustomerNegativeCriterionRequest
	(*MutateCustomerNegativeCriteriaRequest)(nil),  // 1: google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaRequest
	(*CustomerNegativeCriterionOperation)(nil),     // 2: google.ads.googleads.v3.services.CustomerNegativeCriterionOperation
	(*MutateCustomerNegativeCriteriaResponse)(nil), // 3: google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaResponse
	(*MutateCustomerNegativeCriteriaResult)(nil),   // 4: google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaResult
	(*resources.CustomerNegativeCriterion)(nil),    // 5: google.ads.googleads.v3.resources.CustomerNegativeCriterion
	(*status.Status)(nil),                          // 6: google.rpc.Status
}
var file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaRequest.operations:type_name -> google.ads.googleads.v3.services.CustomerNegativeCriterionOperation
	5, // 1: google.ads.googleads.v3.services.CustomerNegativeCriterionOperation.create:type_name -> google.ads.googleads.v3.resources.CustomerNegativeCriterion
	6, // 2: google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaResponse.partial_failure_error:type_name -> google.rpc.Status
	4, // 3: google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaResponse.results:type_name -> google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaResult
	0, // 4: google.ads.googleads.v3.services.CustomerNegativeCriterionService.GetCustomerNegativeCriterion:input_type -> google.ads.googleads.v3.services.GetCustomerNegativeCriterionRequest
	1, // 5: google.ads.googleads.v3.services.CustomerNegativeCriterionService.MutateCustomerNegativeCriteria:input_type -> google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaRequest
	5, // 6: google.ads.googleads.v3.services.CustomerNegativeCriterionService.GetCustomerNegativeCriterion:output_type -> google.ads.googleads.v3.resources.CustomerNegativeCriterion
	3, // 7: google.ads.googleads.v3.services.CustomerNegativeCriterionService.MutateCustomerNegativeCriteria:output_type -> google.ads.googleads.v3.services.MutateCustomerNegativeCriteriaResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_init() }
func file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_init() {
	if File_google_ads_googleads_v3_services_customer_negative_criterion_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerNegativeCriterionRequest); i {
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
		file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerNegativeCriteriaRequest); i {
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
		file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerNegativeCriterionOperation); i {
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
		file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerNegativeCriteriaResponse); i {
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
		file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerNegativeCriteriaResult); i {
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
	file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CustomerNegativeCriterionOperation_Create)(nil),
		(*CustomerNegativeCriterionOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v3_services_customer_negative_criterion_service_proto = out.File
	file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_rawDesc = nil
	file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_goTypes = nil
	file_google_ads_googleads_v3_services_customer_negative_criterion_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerNegativeCriterionServiceClient is the client API for CustomerNegativeCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerNegativeCriterionServiceClient interface {
	// Returns the requested criterion in full detail.
	GetCustomerNegativeCriterion(ctx context.Context, in *GetCustomerNegativeCriterionRequest, opts ...grpc.CallOption) (*resources.CustomerNegativeCriterion, error)
	// Creates or removes criteria. Operation statuses are returned.
	MutateCustomerNegativeCriteria(ctx context.Context, in *MutateCustomerNegativeCriteriaRequest, opts ...grpc.CallOption) (*MutateCustomerNegativeCriteriaResponse, error)
}

type customerNegativeCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerNegativeCriterionServiceClient(cc grpc.ClientConnInterface) CustomerNegativeCriterionServiceClient {
	return &customerNegativeCriterionServiceClient{cc}
}

func (c *customerNegativeCriterionServiceClient) GetCustomerNegativeCriterion(ctx context.Context, in *GetCustomerNegativeCriterionRequest, opts ...grpc.CallOption) (*resources.CustomerNegativeCriterion, error) {
	out := new(resources.CustomerNegativeCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerNegativeCriterionService/GetCustomerNegativeCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerNegativeCriterionServiceClient) MutateCustomerNegativeCriteria(ctx context.Context, in *MutateCustomerNegativeCriteriaRequest, opts ...grpc.CallOption) (*MutateCustomerNegativeCriteriaResponse, error) {
	out := new(MutateCustomerNegativeCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CustomerNegativeCriterionService/MutateCustomerNegativeCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerNegativeCriterionServiceServer is the server API for CustomerNegativeCriterionService service.
type CustomerNegativeCriterionServiceServer interface {
	// Returns the requested criterion in full detail.
	GetCustomerNegativeCriterion(context.Context, *GetCustomerNegativeCriterionRequest) (*resources.CustomerNegativeCriterion, error)
	// Creates or removes criteria. Operation statuses are returned.
	MutateCustomerNegativeCriteria(context.Context, *MutateCustomerNegativeCriteriaRequest) (*MutateCustomerNegativeCriteriaResponse, error)
}

// UnimplementedCustomerNegativeCriterionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerNegativeCriterionServiceServer struct {
}

func (*UnimplementedCustomerNegativeCriterionServiceServer) GetCustomerNegativeCriterion(context.Context, *GetCustomerNegativeCriterionRequest) (*resources.CustomerNegativeCriterion, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCustomerNegativeCriterion not implemented")
}
func (*UnimplementedCustomerNegativeCriterionServiceServer) MutateCustomerNegativeCriteria(context.Context, *MutateCustomerNegativeCriteriaRequest) (*MutateCustomerNegativeCriteriaResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCustomerNegativeCriteria not implemented")
}

func RegisterCustomerNegativeCriterionServiceServer(s *grpc.Server, srv CustomerNegativeCriterionServiceServer) {
	s.RegisterService(&_CustomerNegativeCriterionService_serviceDesc, srv)
}

func _CustomerNegativeCriterionService_GetCustomerNegativeCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerNegativeCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNegativeCriterionServiceServer).GetCustomerNegativeCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerNegativeCriterionService/GetCustomerNegativeCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNegativeCriterionServiceServer).GetCustomerNegativeCriterion(ctx, req.(*GetCustomerNegativeCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerNegativeCriterionService_MutateCustomerNegativeCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerNegativeCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNegativeCriterionServiceServer).MutateCustomerNegativeCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CustomerNegativeCriterionService/MutateCustomerNegativeCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNegativeCriterionServiceServer).MutateCustomerNegativeCriteria(ctx, req.(*MutateCustomerNegativeCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerNegativeCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CustomerNegativeCriterionService",
	HandlerType: (*CustomerNegativeCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerNegativeCriterion",
			Handler:    _CustomerNegativeCriterionService_GetCustomerNegativeCriterion_Handler,
		},
		{
			MethodName: "MutateCustomerNegativeCriteria",
			Handler:    _CustomerNegativeCriterionService_MutateCustomerNegativeCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/customer_negative_criterion_service.proto",
}
