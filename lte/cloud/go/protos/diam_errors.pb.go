//
//Copyright 2022 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: lte/protos/diam_errors.proto

package protos

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

// ErrorCode reflects Experimental-Result values which are 3GPP failures
// to be processed by EPC. Diameter Base Protocol errors are reflected in gRPC status code
type ErrorCode int32

const (
	ErrorCode_UNDEFINED ErrorCode = 0
	// Default success code
	ErrorCode_MULTI_ROUND_AUTH        ErrorCode = 1001
	ErrorCode_SUCCESS                 ErrorCode = 2001
	ErrorCode_LIMITED_SUCCESS         ErrorCode = 2002
	ErrorCode_COMMAND_UNSUPORTED      ErrorCode = 3001
	ErrorCode_UNABLE_TO_DELIVER       ErrorCode = 3002
	ErrorCode_REALM_NOT_SERVED        ErrorCode = 3003
	ErrorCode_TOO_BUSY                ErrorCode = 3004
	ErrorCode_LOOP_DETECTED           ErrorCode = 3005
	ErrorCode_REDIRECT_INDICATION     ErrorCode = 3006
	ErrorCode_APPLICATION_UNSUPPORTED ErrorCode = 3007
	ErrorCode_INVALIDH_DR_BITS        ErrorCode = 3008
	ErrorCode_INVALID_AVP_BITS        ErrorCode = 3009
	ErrorCode_UNKNOWN_PEER            ErrorCode = 3010
	ErrorCode_AUTHENTICATION_REJECTED ErrorCode = 4001
	ErrorCode_OUT_OF_SPACE            ErrorCode = 4002
	ErrorCode_ELECTION_LOST           ErrorCode = 4003
	ErrorCode_AUTHORIZATION_REJECTED  ErrorCode = 5003
	// Permanent Failures 7.4.3
	ErrorCode_USER_UNKNOWN             ErrorCode = 5001
	ErrorCode_UNKNOWN_SESSION_ID       ErrorCode = 5002
	ErrorCode_UNKNOWN_EPS_SUBSCRIPTION ErrorCode = 5420
	ErrorCode_RAT_NOT_ALLOWED          ErrorCode = 5421
	ErrorCode_ROAMING_NOT_ALLOWED      ErrorCode = 5004
	ErrorCode_EQUIPMENT_UNKNOWN        ErrorCode = 5422
	ErrorCode_UNKNOWN_SERVING_NODE     ErrorCode = 5423
	// Transient Failures 7.4.4
	ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE ErrorCode = 4181
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:    "UNDEFINED",
		1001: "MULTI_ROUND_AUTH",
		2001: "SUCCESS",
		2002: "LIMITED_SUCCESS",
		3001: "COMMAND_UNSUPORTED",
		3002: "UNABLE_TO_DELIVER",
		3003: "REALM_NOT_SERVED",
		3004: "TOO_BUSY",
		3005: "LOOP_DETECTED",
		3006: "REDIRECT_INDICATION",
		3007: "APPLICATION_UNSUPPORTED",
		3008: "INVALIDH_DR_BITS",
		3009: "INVALID_AVP_BITS",
		3010: "UNKNOWN_PEER",
		4001: "AUTHENTICATION_REJECTED",
		4002: "OUT_OF_SPACE",
		4003: "ELECTION_LOST",
		5003: "AUTHORIZATION_REJECTED",
		5001: "USER_UNKNOWN",
		5002: "UNKNOWN_SESSION_ID",
		5420: "UNKNOWN_EPS_SUBSCRIPTION",
		5421: "RAT_NOT_ALLOWED",
		5004: "ROAMING_NOT_ALLOWED",
		5422: "EQUIPMENT_UNKNOWN",
		5423: "UNKNOWN_SERVING_NODE",
		4181: "AUTHENTICATION_DATA_UNAVAILABLE",
	}
	ErrorCode_value = map[string]int32{
		"UNDEFINED":                       0,
		"MULTI_ROUND_AUTH":                1001,
		"SUCCESS":                         2001,
		"LIMITED_SUCCESS":                 2002,
		"COMMAND_UNSUPORTED":              3001,
		"UNABLE_TO_DELIVER":               3002,
		"REALM_NOT_SERVED":                3003,
		"TOO_BUSY":                        3004,
		"LOOP_DETECTED":                   3005,
		"REDIRECT_INDICATION":             3006,
		"APPLICATION_UNSUPPORTED":         3007,
		"INVALIDH_DR_BITS":                3008,
		"INVALID_AVP_BITS":                3009,
		"UNKNOWN_PEER":                    3010,
		"AUTHENTICATION_REJECTED":         4001,
		"OUT_OF_SPACE":                    4002,
		"ELECTION_LOST":                   4003,
		"AUTHORIZATION_REJECTED":          5003,
		"USER_UNKNOWN":                    5001,
		"UNKNOWN_SESSION_ID":              5002,
		"UNKNOWN_EPS_SUBSCRIPTION":        5420,
		"RAT_NOT_ALLOWED":                 5421,
		"ROAMING_NOT_ALLOWED":             5004,
		"EQUIPMENT_UNKNOWN":               5422,
		"UNKNOWN_SERVING_NODE":            5423,
		"AUTHENTICATION_DATA_UNAVAILABLE": 4181,
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
	return file_lte_protos_diam_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_lte_protos_diam_errors_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_lte_protos_diam_errors_proto_rawDescGZIP(), []int{0}
}

var File_lte_protos_diam_errors_proto protoreflect.FileDescriptor

var file_lte_protos_diam_errors_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x61,
	0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2a, 0xef, 0x04, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f,
	0x52, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x10, 0xe9, 0x07, 0x12, 0x0c, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0xd1, 0x0f, 0x12, 0x14, 0x0a, 0x0f, 0x4c,
	0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0xd2,
	0x0f, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x55, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0xb9, 0x17, 0x12, 0x16, 0x0a, 0x11, 0x55, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x10,
	0xba, 0x17, 0x12, 0x15, 0x0a, 0x10, 0x52, 0x45, 0x41, 0x4c, 0x4d, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x10, 0xbb, 0x17, 0x12, 0x0d, 0x0a, 0x08, 0x54, 0x4f, 0x4f,
	0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0xbc, 0x17, 0x12, 0x12, 0x0a, 0x0d, 0x4c, 0x4f, 0x4f, 0x50,
	0x5f, 0x44, 0x45, 0x54, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0xbd, 0x17, 0x12, 0x18, 0x0a, 0x13,
	0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x49, 0x4e, 0x44, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0xbe, 0x17, 0x12, 0x1c, 0x0a, 0x17, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45,
	0x44, 0x10, 0xbf, 0x17, 0x12, 0x15, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x48,
	0x5f, 0x44, 0x52, 0x5f, 0x42, 0x49, 0x54, 0x53, 0x10, 0xc0, 0x17, 0x12, 0x15, 0x0a, 0x10, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x56, 0x50, 0x5f, 0x42, 0x49, 0x54, 0x53, 0x10,
	0xc1, 0x17, 0x12, 0x11, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x45,
	0x45, 0x52, 0x10, 0xc2, 0x17, 0x12, 0x1c, 0x0a, 0x17, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x10, 0xa1, 0x1f, 0x12, 0x11, 0x0a, 0x0c, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x53, 0x50,
	0x41, 0x43, 0x45, 0x10, 0xa2, 0x1f, 0x12, 0x12, 0x0a, 0x0d, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x53, 0x54, 0x10, 0xa3, 0x1f, 0x12, 0x1b, 0x0a, 0x16, 0x41, 0x55,
	0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4a, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x8b, 0x27, 0x12, 0x11, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x89, 0x27, 0x12, 0x17, 0x0a, 0x12, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44,
	0x10, 0x8a, 0x27, 0x12, 0x1d, 0x0a, 0x18, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45,
	0x50, 0x53, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0xac, 0x2a, 0x12, 0x14, 0x0a, 0x0f, 0x52, 0x41, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0xad, 0x2a, 0x12, 0x18, 0x0a, 0x13, 0x52, 0x4f, 0x41, 0x4d,
	0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10,
	0x8c, 0x27, 0x12, 0x16, 0x0a, 0x11, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xae, 0x2a, 0x12, 0x19, 0x0a, 0x14, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x4f,
	0x44, 0x45, 0x10, 0xaf, 0x2a, 0x12, 0x24, 0x0a, 0x1f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0xd5, 0x20, 0x42, 0x1b, 0x5a, 0x19, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6c, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lte_protos_diam_errors_proto_rawDescOnce sync.Once
	file_lte_protos_diam_errors_proto_rawDescData = file_lte_protos_diam_errors_proto_rawDesc
)

func file_lte_protos_diam_errors_proto_rawDescGZIP() []byte {
	file_lte_protos_diam_errors_proto_rawDescOnce.Do(func() {
		file_lte_protos_diam_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_lte_protos_diam_errors_proto_rawDescData)
	})
	return file_lte_protos_diam_errors_proto_rawDescData
}

var file_lte_protos_diam_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lte_protos_diam_errors_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: magma.lte.ErrorCode
}
var file_lte_protos_diam_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lte_protos_diam_errors_proto_init() }
func file_lte_protos_diam_errors_proto_init() {
	if File_lte_protos_diam_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lte_protos_diam_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lte_protos_diam_errors_proto_goTypes,
		DependencyIndexes: file_lte_protos_diam_errors_proto_depIdxs,
		EnumInfos:         file_lte_protos_diam_errors_proto_enumTypes,
	}.Build()
	File_lte_protos_diam_errors_proto = out.File
	file_lte_protos_diam_errors_proto_rawDesc = nil
	file_lte_protos_diam_errors_proto_goTypes = nil
	file_lte_protos_diam_errors_proto_depIdxs = nil
}
