// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: telem.proto

package telem

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

type TelemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Telemetry     []byte `protobuf:"bytes,1,opt,name=telemetry,proto3" json:"telemetry,omitempty"`
	Address       string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	TelemetryType string `protobuf:"bytes,3,opt,name=telemetry_type,json=telemetryType,proto3" json:"telemetry_type,omitempty"`
	SentAt        int64  `protobuf:"varint,4,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *TelemRequest) Reset() {
	*x = TelemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemRequest) ProtoMessage() {}

func (x *TelemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemRequest.ProtoReflect.Descriptor instead.
func (*TelemRequest) Descriptor() ([]byte, []int) {
	return file_telem_proto_rawDescGZIP(), []int{0}
}

func (x *TelemRequest) GetTelemetry() []byte {
	if x != nil {
		return x.Telemetry
	}
	return nil
}

func (x *TelemRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TelemRequest) GetTelemetryType() string {
	if x != nil {
		return x.TelemetryType
	}
	return ""
}

func (x *TelemRequest) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

type TelemBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId    string   `protobuf:"bytes,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	Telemetry     [][]byte `protobuf:"bytes,2,rep,name=telemetry,proto3" json:"telemetry,omitempty"`
	TelemetryType string   `protobuf:"bytes,3,opt,name=telemetry_type,json=telemetryType,proto3" json:"telemetry_type,omitempty"`
	SentAt        int64    `protobuf:"varint,4,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *TelemBatchRequest) Reset() {
	*x = TelemBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemBatchRequest) ProtoMessage() {}

func (x *TelemBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemBatchRequest.ProtoReflect.Descriptor instead.
func (*TelemBatchRequest) Descriptor() ([]byte, []int) {
	return file_telem_proto_rawDescGZIP(), []int{1}
}

func (x *TelemBatchRequest) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *TelemBatchRequest) GetTelemetry() [][]byte {
	if x != nil {
		return x.Telemetry
	}
	return nil
}

func (x *TelemBatchRequest) GetTelemetryType() string {
	if x != nil {
		return x.TelemetryType
	}
	return ""
}

func (x *TelemBatchRequest) GetSentAt() int64 {
	if x != nil {
		return x.SentAt
	}
	return 0
}

type TelemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *TelemResponse) Reset() {
	*x = TelemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemResponse) ProtoMessage() {}

func (x *TelemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemResponse.ProtoReflect.Descriptor instead.
func (*TelemResponse) Descriptor() ([]byte, []int) {
	return file_telem_proto_rawDescGZIP(), []int{2}
}

func (x *TelemResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_telem_proto protoreflect.FileDescriptor

var file_telem_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x92, 0x01,
	0x0a, 0x11, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x22, 0x23, 0x0a, 0x0d, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x79, 0x0a, 0x05, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x12, 0x32, 0x0a, 0x05, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x12, 0x13, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69,
	0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72,
	0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telem_proto_rawDescOnce sync.Once
	file_telem_proto_rawDescData = file_telem_proto_rawDesc
)

func file_telem_proto_rawDescGZIP() []byte {
	file_telem_proto_rawDescOnce.Do(func() {
		file_telem_proto_rawDescData = protoimpl.X.CompressGZIP(file_telem_proto_rawDescData)
	})
	return file_telem_proto_rawDescData
}

var file_telem_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_telem_proto_goTypes = []interface{}{
	(*TelemRequest)(nil),      // 0: telem.TelemRequest
	(*TelemBatchRequest)(nil), // 1: telem.TelemBatchRequest
	(*TelemResponse)(nil),     // 2: telem.TelemResponse
}
var file_telem_proto_depIdxs = []int32{
	0, // 0: telem.Telem.Telem:input_type -> telem.TelemRequest
	1, // 1: telem.Telem.TelemBatch:input_type -> telem.TelemBatchRequest
	2, // 2: telem.Telem.Telem:output_type -> telem.TelemResponse
	2, // 3: telem.Telem.TelemBatch:output_type -> telem.TelemResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_telem_proto_init() }
func file_telem_proto_init() {
	if File_telem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemRequest); i {
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
		file_telem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemBatchRequest); i {
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
		file_telem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemResponse); i {
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
			RawDescriptor: file_telem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_telem_proto_goTypes,
		DependencyIndexes: file_telem_proto_depIdxs,
		MessageInfos:      file_telem_proto_msgTypes,
	}.Build()
	File_telem_proto = out.File
	file_telem_proto_rawDesc = nil
	file_telem_proto_goTypes = nil
	file_telem_proto_depIdxs = nil
}
