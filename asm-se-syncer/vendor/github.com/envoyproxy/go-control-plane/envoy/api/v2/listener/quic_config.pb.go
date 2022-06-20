// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/api/v2/listener/quic_config.proto

package envoy_api_v2_listener

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration specific to the QUIC protocol.
// Next id: 4
type QuicProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maximum number of streams that the client can negotiate per connection. 100
	// if not specified.
	MaxConcurrentStreams *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	// Maximum number of milliseconds that connection will be alive when there is
	// no network activity. 300000ms if not specified.
	IdleTimeout *duration.Duration `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// Connection timeout in milliseconds before the crypto handshake is finished.
	// 20000ms if not specified.
	CryptoHandshakeTimeout *duration.Duration `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
}

func (x *QuicProtocolOptions) Reset() {
	*x = QuicProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_listener_quic_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuicProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuicProtocolOptions) ProtoMessage() {}

func (x *QuicProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_listener_quic_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuicProtocolOptions.ProtoReflect.Descriptor instead.
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_listener_quic_config_proto_rawDescGZIP(), []int{0}
}

func (x *QuicProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxConcurrentStreams
	}
	return nil
}

func (x *QuicProtocolOptions) GetIdleTimeout() *duration.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *QuicProtocolOptions) GetCryptoHandshakeTimeout() *duration.Duration {
	if x != nil {
		return x.CryptoHandshakeTimeout
	}
	return nil
}

var File_envoy_api_v2_listener_quic_config_proto protoreflect.FileDescriptor

var file_envoy_api_v2_listener_quic_config_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfc, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x52, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x69,
	0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x64,
	0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x53, 0x0a, 0x18, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x48, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x94,
	0x01, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x42, 0x0f, 0x51, 0x75, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xaa, 0x02, 0x17, 0x45, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x4e, 0x53, 0xea, 0x02, 0x17, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x56, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x4e, 0x53, 0xf2, 0x98, 0xfe,
	0x8f, 0x05, 0x1a, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_api_v2_listener_quic_config_proto_rawDescOnce sync.Once
	file_envoy_api_v2_listener_quic_config_proto_rawDescData = file_envoy_api_v2_listener_quic_config_proto_rawDesc
)

func file_envoy_api_v2_listener_quic_config_proto_rawDescGZIP() []byte {
	file_envoy_api_v2_listener_quic_config_proto_rawDescOnce.Do(func() {
		file_envoy_api_v2_listener_quic_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_api_v2_listener_quic_config_proto_rawDescData)
	})
	return file_envoy_api_v2_listener_quic_config_proto_rawDescData
}

var file_envoy_api_v2_listener_quic_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_api_v2_listener_quic_config_proto_goTypes = []interface{}{
	(*QuicProtocolOptions)(nil),  // 0: envoy.api.v2.listener.QuicProtocolOptions
	(*wrappers.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
	(*duration.Duration)(nil),    // 2: google.protobuf.Duration
}
var file_envoy_api_v2_listener_quic_config_proto_depIdxs = []int32{
	1, // 0: envoy.api.v2.listener.QuicProtocolOptions.max_concurrent_streams:type_name -> google.protobuf.UInt32Value
	2, // 1: envoy.api.v2.listener.QuicProtocolOptions.idle_timeout:type_name -> google.protobuf.Duration
	2, // 2: envoy.api.v2.listener.QuicProtocolOptions.crypto_handshake_timeout:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_api_v2_listener_quic_config_proto_init() }
func file_envoy_api_v2_listener_quic_config_proto_init() {
	if File_envoy_api_v2_listener_quic_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_api_v2_listener_quic_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuicProtocolOptions); i {
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
			RawDescriptor: file_envoy_api_v2_listener_quic_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_api_v2_listener_quic_config_proto_goTypes,
		DependencyIndexes: file_envoy_api_v2_listener_quic_config_proto_depIdxs,
		MessageInfos:      file_envoy_api_v2_listener_quic_config_proto_msgTypes,
	}.Build()
	File_envoy_api_v2_listener_quic_config_proto = out.File
	file_envoy_api_v2_listener_quic_config_proto_rawDesc = nil
	file_envoy_api_v2_listener_quic_config_proto_goTypes = nil
	file_envoy_api_v2_listener_quic_config_proto_depIdxs = nil
}
