// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: workflows/wasm/pb/wasm.proto

package pb

import (
	pb "github.com/smartcontractkit/chainlink-common/pkg/capabilities/pb"
	pb1 "github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ComputeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *pb.CapabilityRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *ComputeRequest) Reset() {
	*x = ComputeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeRequest) ProtoMessage() {}

func (x *ComputeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeRequest.ProtoReflect.Descriptor instead.
func (*ComputeRequest) Descriptor() ([]byte, []int) {
	return file_workflows_wasm_pb_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeRequest) GetRequest() *pb.CapabilityRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Config []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// Types that are assignable to Message:
	//
	//	*Request_ComputeRequest
	//	*Request_SpecRequest
	Message isRequest_Message `protobuf_oneof:"message"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_workflows_wasm_pb_wasm_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Request) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

func (m *Request) GetMessage() isRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *Request) GetComputeRequest() *ComputeRequest {
	if x, ok := x.GetMessage().(*Request_ComputeRequest); ok {
		return x.ComputeRequest
	}
	return nil
}

func (x *Request) GetSpecRequest() *emptypb.Empty {
	if x, ok := x.GetMessage().(*Request_SpecRequest); ok {
		return x.SpecRequest
	}
	return nil
}

type isRequest_Message interface {
	isRequest_Message()
}

type Request_ComputeRequest struct {
	ComputeRequest *ComputeRequest `protobuf:"bytes,3,opt,name=computeRequest,proto3,oneof"`
}

type Request_SpecRequest struct {
	SpecRequest *emptypb.Empty `protobuf:"bytes,4,opt,name=specRequest,proto3,oneof"`
}

func (*Request_ComputeRequest) isRequest_Message() {}

func (*Request_SpecRequest) isRequest_Message() {}

type ComputeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *pb.CapabilityResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ComputeResponse) Reset() {
	*x = ComputeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeResponse) ProtoMessage() {}

func (x *ComputeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeResponse.ProtoReflect.Descriptor instead.
func (*ComputeResponse) Descriptor() ([]byte, []int) {
	return file_workflows_wasm_pb_wasm_proto_rawDescGZIP(), []int{2}
}

func (x *ComputeResponse) GetResponse() *pb.CapabilityResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type StepInputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputRef string   `protobuf:"bytes,1,opt,name=outputRef,proto3" json:"outputRef,omitempty"`
	Mapping   *pb1.Map `protobuf:"bytes,2,opt,name=mapping,proto3" json:"mapping,omitempty"`
}

func (x *StepInputs) Reset() {
	*x = StepInputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepInputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepInputs) ProtoMessage() {}

func (x *StepInputs) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepInputs.ProtoReflect.Descriptor instead.
func (*StepInputs) Descriptor() ([]byte, []int) {
	return file_workflows_wasm_pb_wasm_proto_rawDescGZIP(), []int{3}
}

func (x *StepInputs) GetOutputRef() string {
	if x != nil {
		return x.OutputRef
	}
	return ""
}

func (x *StepInputs) GetMapping() *pb1.Map {
	if x != nil {
		return x.Mapping
	}
	return nil
}

type StepDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ref            string      `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	Inputs         *StepInputs `protobuf:"bytes,3,opt,name=inputs,proto3" json:"inputs,omitempty"`
	Config         *pb1.Map    `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	CapabilityType string      `protobuf:"bytes,5,opt,name=capabilityType,proto3" json:"capabilityType,omitempty"`
}

func (x *StepDefinition) Reset() {
	*x = StepDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepDefinition) ProtoMessage() {}

func (x *StepDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepDefinition.ProtoReflect.Descriptor instead.
func (*StepDefinition) Descriptor() ([]byte, []int) {
	return file_workflows_wasm_pb_wasm_proto_rawDescGZIP(), []int{4}
}

func (x *StepDefinition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StepDefinition) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *StepDefinition) GetInputs() *StepInputs {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *StepDefinition) GetConfig() *pb1.Map {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *StepDefinition) GetCapabilityType() string {
	if x != nil {
		return x.CapabilityType
	}
	return ""
}

type WorkflowSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner     string            `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Triggers  []*StepDefinition `protobuf:"bytes,3,rep,name=triggers,proto3" json:"triggers,omitempty"`
	Actions   []*StepDefinition `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	Consensus []*StepDefinition `protobuf:"bytes,5,rep,name=consensus,proto3" json:"consensus,omitempty"`
	Targets   []*StepDefinition `protobuf:"bytes,6,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *WorkflowSpec) Reset() {
	*x = WorkflowSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowSpec) ProtoMessage() {}

func (x *WorkflowSpec) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowSpec.ProtoReflect.Descriptor instead.
func (*WorkflowSpec) Descriptor() ([]byte, []int) {
	return file_workflows_wasm_pb_wasm_proto_rawDescGZIP(), []int{5}
}

func (x *WorkflowSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowSpec) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *WorkflowSpec) GetTriggers() []*StepDefinition {
	if x != nil {
		return x.Triggers
	}
	return nil
}

func (x *WorkflowSpec) GetActions() []*StepDefinition {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *WorkflowSpec) GetConsensus() []*StepDefinition {
	if x != nil {
		return x.Consensus
	}
	return nil
}

func (x *WorkflowSpec) GetTargets() []*StepDefinition {
	if x != nil {
		return x.Targets
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	// Types that are assignable to Message:
	//
	//	*Response_ComputeResponse
	//	*Response_SpecResponse
	Message isResponse_Message `protobuf_oneof:"message"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_wasm_pb_wasm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_workflows_wasm_pb_wasm_proto_rawDescGZIP(), []int{6}
}

func (x *Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Response) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (m *Response) GetMessage() isResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *Response) GetComputeResponse() *ComputeResponse {
	if x, ok := x.GetMessage().(*Response_ComputeResponse); ok {
		return x.ComputeResponse
	}
	return nil
}

func (x *Response) GetSpecResponse() *WorkflowSpec {
	if x, ok := x.GetMessage().(*Response_SpecResponse); ok {
		return x.SpecResponse
	}
	return nil
}

type isResponse_Message interface {
	isResponse_Message()
}

type Response_ComputeResponse struct {
	ComputeResponse *ComputeResponse `protobuf:"bytes,3,opt,name=computeResponse,proto3,oneof"`
}

type Response_SpecResponse struct {
	SpecResponse *WorkflowSpec `protobuf:"bytes,4,opt,name=specResponse,proto3,oneof"`
}

func (*Response_ComputeResponse) isResponse_Message() {}

func (*Response_SpecResponse) isResponse_Message() {}

var File_workflows_wasm_pb_wasm_proto protoreflect.FileDescriptor

var file_workflows_wasm_pb_wasm_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d,
	0x2f, 0x70, 0x62, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x73, 0x64, 0x6b, 0x1a, 0x22, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2f,
	0x70, 0x62, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb7, 0x01, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b,
	0x73, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x70, 0x65,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x0a, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x66,
	0x12, 0x25, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x07,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x65, 0x70,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x27, 0x0a, 0x06,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x52, 0x06, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x70, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x2d,
	0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x12, 0x2d, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x22,
	0xb8, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x4d, 0x73, 0x67, 0x12, 0x40, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x70, 0x65, 0x63, 0x48,
	0x00, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflows_wasm_pb_wasm_proto_rawDescOnce sync.Once
	file_workflows_wasm_pb_wasm_proto_rawDescData = file_workflows_wasm_pb_wasm_proto_rawDesc
)

func file_workflows_wasm_pb_wasm_proto_rawDescGZIP() []byte {
	file_workflows_wasm_pb_wasm_proto_rawDescOnce.Do(func() {
		file_workflows_wasm_pb_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflows_wasm_pb_wasm_proto_rawDescData)
	})
	return file_workflows_wasm_pb_wasm_proto_rawDescData
}

var file_workflows_wasm_pb_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_workflows_wasm_pb_wasm_proto_goTypes = []interface{}{
	(*ComputeRequest)(nil),        // 0: sdk.ComputeRequest
	(*Request)(nil),               // 1: sdk.Request
	(*ComputeResponse)(nil),       // 2: sdk.ComputeResponse
	(*StepInputs)(nil),            // 3: sdk.StepInputs
	(*StepDefinition)(nil),        // 4: sdk.StepDefinition
	(*WorkflowSpec)(nil),          // 5: sdk.WorkflowSpec
	(*Response)(nil),              // 6: sdk.Response
	(*pb.CapabilityRequest)(nil),  // 7: capabilities.CapabilityRequest
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
	(*pb.CapabilityResponse)(nil), // 9: capabilities.CapabilityResponse
	(*pb1.Map)(nil),               // 10: values.Map
}
var file_workflows_wasm_pb_wasm_proto_depIdxs = []int32{
	7,  // 0: sdk.ComputeRequest.request:type_name -> capabilities.CapabilityRequest
	0,  // 1: sdk.Request.computeRequest:type_name -> sdk.ComputeRequest
	8,  // 2: sdk.Request.specRequest:type_name -> google.protobuf.Empty
	9,  // 3: sdk.ComputeResponse.response:type_name -> capabilities.CapabilityResponse
	10, // 4: sdk.StepInputs.mapping:type_name -> values.Map
	3,  // 5: sdk.StepDefinition.inputs:type_name -> sdk.StepInputs
	10, // 6: sdk.StepDefinition.config:type_name -> values.Map
	4,  // 7: sdk.WorkflowSpec.triggers:type_name -> sdk.StepDefinition
	4,  // 8: sdk.WorkflowSpec.actions:type_name -> sdk.StepDefinition
	4,  // 9: sdk.WorkflowSpec.consensus:type_name -> sdk.StepDefinition
	4,  // 10: sdk.WorkflowSpec.targets:type_name -> sdk.StepDefinition
	2,  // 11: sdk.Response.computeResponse:type_name -> sdk.ComputeResponse
	5,  // 12: sdk.Response.specResponse:type_name -> sdk.WorkflowSpec
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_workflows_wasm_pb_wasm_proto_init() }
func file_workflows_wasm_pb_wasm_proto_init() {
	if File_workflows_wasm_pb_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflows_wasm_pb_wasm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeRequest); i {
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
		file_workflows_wasm_pb_wasm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_workflows_wasm_pb_wasm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeResponse); i {
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
		file_workflows_wasm_pb_wasm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepInputs); i {
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
		file_workflows_wasm_pb_wasm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepDefinition); i {
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
		file_workflows_wasm_pb_wasm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowSpec); i {
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
		file_workflows_wasm_pb_wasm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_workflows_wasm_pb_wasm_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Request_ComputeRequest)(nil),
		(*Request_SpecRequest)(nil),
	}
	file_workflows_wasm_pb_wasm_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Response_ComputeResponse)(nil),
		(*Response_SpecResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workflows_wasm_pb_wasm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workflows_wasm_pb_wasm_proto_goTypes,
		DependencyIndexes: file_workflows_wasm_pb_wasm_proto_depIdxs,
		MessageInfos:      file_workflows_wasm_pb_wasm_proto_msgTypes,
	}.Build()
	File_workflows_wasm_pb_wasm_proto = out.File
	file_workflows_wasm_pb_wasm_proto_rawDesc = nil
	file_workflows_wasm_pb_wasm_proto_goTypes = nil
	file_workflows_wasm_pb_wasm_proto_depIdxs = nil
}
