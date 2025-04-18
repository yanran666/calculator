// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: calculator.proto

package calculatorv1

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

type CalculatorRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number1       float64                `protobuf:"fixed64,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2       float64                `protobuf:"fixed64,2,opt,name=number2,proto3" json:"number2,omitempty"`
	Operator      string                 `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	mi := &file_calculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculatorRequest) GetNumber1() float64 {
	if x != nil {
		return x.Number1
	}
	return 0
}

func (x *CalculatorRequest) GetNumber2() float64 {
	if x != nil {
		return x.Number2
	}
	return 0
}

func (x *CalculatorRequest) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

type CalculatorResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float64                `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	mi := &file_calculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_proto protoreflect.FileDescriptor

const file_calculator_proto_rawDesc = "" +
	"\n" +
	"\x10calculator.proto\x12\rcalculator.v1\"c\n" +
	"\x11CalculatorRequest\x12\x18\n" +
	"\anumber1\x18\x01 \x01(\x01R\anumber1\x12\x18\n" +
	"\anumber2\x18\x02 \x01(\x01R\anumber2\x12\x1a\n" +
	"\boperator\x18\x03 \x01(\tR\boperator\",\n" +
	"\x12CalculatorResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\x01R\x06result2e\n" +
	"\x11CalculatorService\x12P\n" +
	"\tCalculate\x12 .calculator.v1.CalculatorRequest\x1a!.calculator.v1.CalculatorResponseB5Z3github.com/yanran666/calculator/server/calculatorv1b\x06proto3"

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData []byte
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calculator_proto_rawDesc), len(file_calculator_proto_rawDesc)))
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_proto_goTypes = []any{
	(*CalculatorRequest)(nil),  // 0: calculator.v1.CalculatorRequest
	(*CalculatorResponse)(nil), // 1: calculator.v1.CalculatorResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.v1.CalculatorService.Calculate:input_type -> calculator.v1.CalculatorRequest
	1, // 1: calculator.v1.CalculatorService.Calculate:output_type -> calculator.v1.CalculatorResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculator_proto_rawDesc), len(file_calculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}
