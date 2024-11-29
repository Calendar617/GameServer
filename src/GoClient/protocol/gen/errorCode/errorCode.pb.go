// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.17.3
// source: src/errorCode.proto

package errorCode

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

type ErrorCode int32

const (
	ErrorCode_SUCCESS    ErrorCode = 0    //成功
	ErrorCode_ARGS_ERROR ErrorCode = 1000 //参数错误
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:    "SUCCESS",
		1000: "ARGS_ERROR",
	}
	ErrorCode_value = map[string]int32{
		"SUCCESS":    0,
		"ARGS_ERROR": 1000,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_src_errorCode_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_src_errorCode_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_src_errorCode_proto_rawDescGZIP(), []int{0}
}

var File_src_errorCode_proto protoreflect.FileDescriptor

var file_src_errorCode_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x72, 0x63, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x2a, 0x29, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0a, 0x41, 0x52,
	0x47, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xe8, 0x07, 0x42, 0x0f, 0x5a, 0x0d, 0x67,
	0x65, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_errorCode_proto_rawDescOnce sync.Once
	file_src_errorCode_proto_rawDescData = file_src_errorCode_proto_rawDesc
)

func file_src_errorCode_proto_rawDescGZIP() []byte {
	file_src_errorCode_proto_rawDescOnce.Do(func() {
		file_src_errorCode_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_errorCode_proto_rawDescData)
	})
	return file_src_errorCode_proto_rawDescData
}

var file_src_errorCode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_errorCode_proto_goTypes = []any{
	(ErrorCode)(0), // 0: errorCode.errorCode
}
var file_src_errorCode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_errorCode_proto_init() }
func file_src_errorCode_proto_init() {
	if File_src_errorCode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_errorCode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_errorCode_proto_goTypes,
		DependencyIndexes: file_src_errorCode_proto_depIdxs,
		EnumInfos:         file_src_errorCode_proto_enumTypes,
	}.Build()
	File_src_errorCode_proto = out.File
	file_src_errorCode_proto_rawDesc = nil
	file_src_errorCode_proto_goTypes = nil
	file_src_errorCode_proto_depIdxs = nil
}