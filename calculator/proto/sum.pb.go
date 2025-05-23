// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: sum.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SumRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FirstNumber   int32                  `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber  int32                  `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	mi := &file_sum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *SumRequest) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sum           int32                  `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	mi := &file_sum_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

var File_sum_proto protoreflect.FileDescriptor

const file_sum_proto_rawDesc = "" +
	"\n" +
	"\tsum.proto\x12\n" +
	"calculator\"T\n" +
	"\n" +
	"SumRequest\x12!\n" +
	"\ffirst_number\x18\x01 \x01(\x05R\vfirstNumber\x12#\n" +
	"\rsecond_number\x18\x02 \x01(\x05R\fsecondNumber\"\x1f\n" +
	"\vSumResponse\x12\x10\n" +
	"\x03sum\x18\x01 \x01(\x05R\x03sumB.Z,github.com/kangs029/grpc-go/calculator/protob\x06proto3"

var (
	file_sum_proto_rawDescOnce sync.Once
	file_sum_proto_rawDescData []byte
)

func file_sum_proto_rawDescGZIP() []byte {
	file_sum_proto_rawDescOnce.Do(func() {
		file_sum_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sum_proto_rawDesc), len(file_sum_proto_rawDesc)))
	})
	return file_sum_proto_rawDescData
}

var file_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sum_proto_goTypes = []any{
	(*SumRequest)(nil),  // 0: calculator.SumRequest
	(*SumResponse)(nil), // 1: calculator.SumResponse
}
var file_sum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sum_proto_init() }
func file_sum_proto_init() {
	if File_sum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sum_proto_rawDesc), len(file_sum_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sum_proto_goTypes,
		DependencyIndexes: file_sum_proto_depIdxs,
		MessageInfos:      file_sum_proto_msgTypes,
	}.Build()
	File_sum_proto = out.File
	file_sum_proto_goTypes = nil
	file_sum_proto_depIdxs = nil
}
