// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: apis/factorioapi/v1/v1.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_apis_factorioapi_v1_v1_proto protoreflect.FileDescriptor

var file_apis_factorioapi_v1_v1_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x42, 0xd4, 0x01, 0x92, 0x41, 0x97, 0x01, 0x12, 0x19, 0x0a, 0x12, 0x46, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x20, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x20, 0x41, 0x50, 0x49,
	0x32, 0x03, 0x31, 0x2e, 0x30, 0x52, 0x3d, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x36, 0x0a, 0x15,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x1b, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x3b, 0x0a, 0x03, 0x35, 0x30, 0x33, 0x12, 0x34, 0x0a, 0x13, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x1b, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65,
	0x6b, 0x6f, 0x6d, 0x65, 0x6f, 0x77, 0x77, 0x77, 0x2f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x69,
	0x6f, 0x2d, 0x61, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x69, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_apis_factorioapi_v1_v1_proto_goTypes = []interface{}{}
var file_apis_factorioapi_v1_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_factorioapi_v1_v1_proto_init() }
func file_apis_factorioapi_v1_v1_proto_init() {
	if File_apis_factorioapi_v1_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_factorioapi_v1_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_factorioapi_v1_v1_proto_goTypes,
		DependencyIndexes: file_apis_factorioapi_v1_v1_proto_depIdxs,
	}.Build()
	File_apis_factorioapi_v1_v1_proto = out.File
	file_apis_factorioapi_v1_v1_proto_rawDesc = nil
	file_apis_factorioapi_v1_v1_proto_goTypes = nil
	file_apis_factorioapi_v1_v1_proto_depIdxs = nil
}
