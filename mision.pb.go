// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: mision.proto

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

type PreparacionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PreparacionRequest) Reset() {
	*x = PreparacionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparacionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparacionRequest) ProtoMessage() {}

func (x *PreparacionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparacionRequest.ProtoReflect.Descriptor instead.
func (*PreparacionRequest) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{0}
}

func (x *PreparacionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PreparacionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PreparacionResponse) Reset() {
	*x = PreparacionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparacionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparacionResponse) ProtoMessage() {}

func (x *PreparacionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparacionResponse.ProtoReflect.Descriptor instead.
func (*PreparacionResponse) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{1}
}

func (x *PreparacionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DecisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Piso     int32 `protobuf:"varint,2,opt,name=piso,proto3" json:"piso,omitempty"`
	Decision int32 `protobuf:"varint,3,opt,name=decision,proto3" json:"decision,omitempty"`
}

func (x *DecisionRequest) Reset() {
	*x = DecisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionRequest) ProtoMessage() {}

func (x *DecisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionRequest.ProtoReflect.Descriptor instead.
func (*DecisionRequest) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{2}
}

func (x *DecisionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DecisionRequest) GetPiso() int32 {
	if x != nil {
		return x.Piso
	}
	return 0
}

func (x *DecisionRequest) GetDecision() int32 {
	if x != nil {
		return x.Decision
	}
	return 0
}

type DecisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataNodeIp string `protobuf:"bytes,1,opt,name=dataNodeIp,proto3" json:"dataNodeIp,omitempty"`
}

func (x *DecisionResponse) Reset() {
	*x = DecisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionResponse) ProtoMessage() {}

func (x *DecisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionResponse.ProtoReflect.Descriptor instead.
func (*DecisionResponse) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{3}
}

func (x *DecisionResponse) GetDataNodeIp() string {
	if x != nil {
		return x.DataNodeIp
	}
	return ""
}

type EliminacionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Piso int32 `protobuf:"varint,2,opt,name=piso,proto3" json:"piso,omitempty"`
}

func (x *EliminacionRequest) Reset() {
	*x = EliminacionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EliminacionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EliminacionRequest) ProtoMessage() {}

func (x *EliminacionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EliminacionRequest.ProtoReflect.Descriptor instead.
func (*EliminacionRequest) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{4}
}

func (x *EliminacionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EliminacionRequest) GetPiso() int32 {
	if x != nil {
		return x.Piso
	}
	return 0
}

type EliminacionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *EliminacionResponse) Reset() {
	*x = EliminacionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EliminacionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EliminacionResponse) ProtoMessage() {}

func (x *EliminacionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EliminacionResponse.ProtoReflect.Descriptor instead.
func (*EliminacionResponse) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{5}
}

func (x *EliminacionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RegistroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Piso int32 `protobuf:"varint,2,opt,name=piso,proto3" json:"piso,omitempty"`
}

func (x *RegistroRequest) Reset() {
	*x = RegistroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistroRequest) ProtoMessage() {}

func (x *RegistroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistroRequest.ProtoReflect.Descriptor instead.
func (*RegistroRequest) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{6}
}

func (x *RegistroRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegistroRequest) GetPiso() int32 {
	if x != nil {
		return x.Piso
	}
	return 0
}

type RegistroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MontoActual int64 `protobuf:"varint,1,opt,name=montoActual,proto3" json:"montoActual,omitempty"`
}

func (x *RegistroResponse) Reset() {
	*x = RegistroResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistroResponse) ProtoMessage() {}

func (x *RegistroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistroResponse.ProtoReflect.Descriptor instead.
func (*RegistroResponse) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{7}
}

func (x *RegistroResponse) GetMontoActual() int64 {
	if x != nil {
		return x.MontoActual
	}
	return 0
}

type MontoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MontoRequest) Reset() {
	*x = MontoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontoRequest) ProtoMessage() {}

func (x *MontoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontoRequest.ProtoReflect.Descriptor instead.
func (*MontoRequest) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{8}
}

func (x *MontoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MontoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monto int64 `protobuf:"varint,1,opt,name=monto,proto3" json:"monto,omitempty"`
}

func (x *MontoResponse) Reset() {
	*x = MontoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontoResponse) ProtoMessage() {}

func (x *MontoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontoResponse.ProtoReflect.Descriptor instead.
func (*MontoResponse) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{9}
}

func (x *MontoResponse) GetMonto() int64 {
	if x != nil {
		return x.Monto
	}
	return 0
}

type GuardarDecisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Piso     int32 `protobuf:"varint,2,opt,name=piso,proto3" json:"piso,omitempty"`
	Decision int32 `protobuf:"varint,3,opt,name=decision,proto3" json:"decision,omitempty"`
}

func (x *GuardarDecisionRequest) Reset() {
	*x = GuardarDecisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardarDecisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardarDecisionRequest) ProtoMessage() {}

func (x *GuardarDecisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardarDecisionRequest.ProtoReflect.Descriptor instead.
func (*GuardarDecisionRequest) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{10}
}

func (x *GuardarDecisionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GuardarDecisionRequest) GetPiso() int32 {
	if x != nil {
		return x.Piso
	}
	return 0
}

func (x *GuardarDecisionRequest) GetDecision() int32 {
	if x != nil {
		return x.Decision
	}
	return 0
}

type GuardarDecisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *GuardarDecisionResponse) Reset() {
	*x = GuardarDecisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardarDecisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardarDecisionResponse) ProtoMessage() {}

func (x *GuardarDecisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardarDecisionResponse.ProtoReflect.Descriptor instead.
func (*GuardarDecisionResponse) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{11}
}

func (x *GuardarDecisionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ObtenerDecisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Piso int32 `protobuf:"varint,2,opt,name=piso,proto3" json:"piso,omitempty"`
}

func (x *ObtenerDecisionRequest) Reset() {
	*x = ObtenerDecisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObtenerDecisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObtenerDecisionRequest) ProtoMessage() {}

func (x *ObtenerDecisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObtenerDecisionRequest.ProtoReflect.Descriptor instead.
func (*ObtenerDecisionRequest) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{12}
}

func (x *ObtenerDecisionRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObtenerDecisionRequest) GetPiso() int32 {
	if x != nil {
		return x.Piso
	}
	return 0
}

type ObtenerDecisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision int32 `protobuf:"varint,1,opt,name=decision,proto3" json:"decision,omitempty"`
}

func (x *ObtenerDecisionResponse) Reset() {
	*x = ObtenerDecisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mision_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObtenerDecisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObtenerDecisionResponse) ProtoMessage() {}

func (x *ObtenerDecisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mision_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObtenerDecisionResponse.ProtoReflect.Descriptor instead.
func (*ObtenerDecisionResponse) Descriptor() ([]byte, []int) {
	return file_mision_proto_rawDescGZIP(), []int{13}
}

func (x *ObtenerDecisionResponse) GetDecision() int32 {
	if x != nil {
		return x.Decision
	}
	return 0
}

var File_mision_proto protoreflect.FileDescriptor

var file_mision_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a,
	0x13, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x51,
	0x0a, 0x0f, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x73, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x69, 0x73, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x32, 0x0a, 0x10, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x70, 0x22, 0x38, 0x0a, 0x12, 0x45, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61,
	0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x69, 0x73, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x69, 0x73, 0x6f, 0x22,
	0x2f, 0x0a, 0x13, 0x45, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x35, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x73, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x69, 0x73, 0x6f, 0x22, 0x34, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x6f, 0x6e, 0x74, 0x6f, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x22, 0x1e, 0x0a,
	0x0c, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a,
	0x0d, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x16, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x44,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x69, 0x73, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x69,
	0x73, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x33,
	0x0a, 0x17, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x3c, 0x0a, 0x16, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x44, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x69, 0x73, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x69, 0x73,
	0x6f, 0x22, 0x35, 0x0a, 0x17, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xed, 0x01, 0x0a, 0x08, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61,
	0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0e, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x13, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x72, 0x45, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x96, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x73,
	0x68, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x4b, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x72, 0x45, 0x6c, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x4d, 0x6f, 0x6e,
	0x74, 0x6f, 0x12, 0x15, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x6e,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xaa, 0x01, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x48,
	0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x4f, 0x62, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x44, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb6,
	0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x47,
	0x75, 0x61, 0x72, 0x64, 0x61, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61, 0x72,
	0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x61,
	0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x54, 0x0a, 0x0f, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f,
	0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa0, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x73, 0x12, 0x4b, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x61, 0x72, 0x4d, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x54, 0x6f, 0x6d, 0x61, 0x72, 0x44, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mision_proto_rawDescOnce sync.Once
	file_mision_proto_rawDescData = file_mision_proto_rawDesc
)

func file_mision_proto_rawDescGZIP() []byte {
	file_mision_proto_rawDescOnce.Do(func() {
		file_mision_proto_rawDescData = protoimpl.X.CompressGZIP(file_mision_proto_rawDescData)
	})
	return file_mision_proto_rawDescData
}

var file_mision_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_mision_proto_goTypes = []interface{}{
	(*PreparacionRequest)(nil),      // 0: mission.PreparacionRequest
	(*PreparacionResponse)(nil),     // 1: mission.PreparacionResponse
	(*DecisionRequest)(nil),         // 2: mission.DecisionRequest
	(*DecisionResponse)(nil),        // 3: mission.DecisionResponse
	(*EliminacionRequest)(nil),      // 4: mission.EliminacionRequest
	(*EliminacionResponse)(nil),     // 5: mission.EliminacionResponse
	(*RegistroRequest)(nil),         // 6: mission.RegistroRequest
	(*RegistroResponse)(nil),        // 7: mission.RegistroResponse
	(*MontoRequest)(nil),            // 8: mission.MontoRequest
	(*MontoResponse)(nil),           // 9: mission.MontoResponse
	(*GuardarDecisionRequest)(nil),  // 10: mission.GuardarDecisionRequest
	(*GuardarDecisionResponse)(nil), // 11: mission.GuardarDecisionResponse
	(*ObtenerDecisionRequest)(nil),  // 12: mission.ObtenerDecisionRequest
	(*ObtenerDecisionResponse)(nil), // 13: mission.ObtenerDecisionResponse
}
var file_mision_proto_depIdxs = []int32{
	0,  // 0: mission.Director.Preparacion:input_type -> mission.PreparacionRequest
	2,  // 1: mission.Director.EnviarDecision:input_type -> mission.DecisionRequest
	4,  // 2: mission.Director.InformarEliminacion:input_type -> mission.EliminacionRequest
	6,  // 3: mission.DoshBank.RegistrarEliminacion:input_type -> mission.RegistroRequest
	8,  // 4: mission.DoshBank.ObtenerMonto:input_type -> mission.MontoRequest
	2,  // 5: mission.NameNode.RegistrarDecision:input_type -> mission.DecisionRequest
	12, // 6: mission.NameNode.ObtenerDecision:input_type -> mission.ObtenerDecisionRequest
	10, // 7: mission.DataNode.GuardarDecision:input_type -> mission.GuardarDecisionRequest
	12, // 8: mission.DataNode.ObtenerDecision:input_type -> mission.ObtenerDecisionRequest
	0,  // 9: mission.Mercenarios.PrepararMision:input_type -> mission.PreparacionRequest
	2,  // 10: mission.Mercenarios.TomarDecision:input_type -> mission.DecisionRequest
	1,  // 11: mission.Director.Preparacion:output_type -> mission.PreparacionResponse
	3,  // 12: mission.Director.EnviarDecision:output_type -> mission.DecisionResponse
	5,  // 13: mission.Director.InformarEliminacion:output_type -> mission.EliminacionResponse
	7,  // 14: mission.DoshBank.RegistrarEliminacion:output_type -> mission.RegistroResponse
	9,  // 15: mission.DoshBank.ObtenerMonto:output_type -> mission.MontoResponse
	3,  // 16: mission.NameNode.RegistrarDecision:output_type -> mission.DecisionResponse
	13, // 17: mission.NameNode.ObtenerDecision:output_type -> mission.ObtenerDecisionResponse
	11, // 18: mission.DataNode.GuardarDecision:output_type -> mission.GuardarDecisionResponse
	13, // 19: mission.DataNode.ObtenerDecision:output_type -> mission.ObtenerDecisionResponse
	1,  // 20: mission.Mercenarios.PrepararMision:output_type -> mission.PreparacionResponse
	3,  // 21: mission.Mercenarios.TomarDecision:output_type -> mission.DecisionResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mision_proto_init() }
func file_mision_proto_init() {
	if File_mision_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mision_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparacionRequest); i {
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
		file_mision_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparacionResponse); i {
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
		file_mision_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecisionRequest); i {
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
		file_mision_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecisionResponse); i {
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
		file_mision_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EliminacionRequest); i {
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
		file_mision_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EliminacionResponse); i {
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
		file_mision_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistroRequest); i {
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
		file_mision_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistroResponse); i {
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
		file_mision_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontoRequest); i {
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
		file_mision_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontoResponse); i {
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
		file_mision_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardarDecisionRequest); i {
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
		file_mision_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardarDecisionResponse); i {
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
		file_mision_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObtenerDecisionRequest); i {
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
		file_mision_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObtenerDecisionResponse); i {
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
			RawDescriptor: file_mision_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_mision_proto_goTypes,
		DependencyIndexes: file_mision_proto_depIdxs,
		MessageInfos:      file_mision_proto_msgTypes,
	}.Build()
	File_mision_proto = out.File
	file_mision_proto_rawDesc = nil
	file_mision_proto_goTypes = nil
	file_mision_proto_depIdxs = nil
}
