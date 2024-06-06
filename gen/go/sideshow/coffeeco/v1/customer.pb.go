// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sideshow/coffeeco/v1/customer.proto

package coffeecov1

import (
	_ "github.com/datacircus/sideshow/gen/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Customer_CustomerType int32

const (
	Customer_CUSTOMER_TYPE_UNSPECIFIED Customer_CustomerType = 0
	Customer_CUSTOMER_TYPE_GUEST       Customer_CustomerType = 1
	Customer_CUSTOMER_TYPE_MEMBER      Customer_CustomerType = 2
)

// Enum value maps for Customer_CustomerType.
var (
	Customer_CustomerType_name = map[int32]string{
		0: "CUSTOMER_TYPE_UNSPECIFIED",
		1: "CUSTOMER_TYPE_GUEST",
		2: "CUSTOMER_TYPE_MEMBER",
	}
	Customer_CustomerType_value = map[string]int32{
		"CUSTOMER_TYPE_UNSPECIFIED": 0,
		"CUSTOMER_TYPE_GUEST":       1,
		"CUSTOMER_TYPE_MEMBER":      2,
	}
)

func (x Customer_CustomerType) Enum() *Customer_CustomerType {
	p := new(Customer_CustomerType)
	*p = x
	return p
}

func (x Customer_CustomerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Customer_CustomerType) Descriptor() protoreflect.EnumDescriptor {
	return file_sideshow_coffeeco_v1_customer_proto_enumTypes[0].Descriptor()
}

func (Customer_CustomerType) Type() protoreflect.EnumType {
	return &file_sideshow_coffeeco_v1_customer_proto_enumTypes[0]
}

func (x Customer_CustomerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Customer_CustomerType.Descriptor instead.
func (Customer_CustomerType) EnumDescriptor() ([]byte, []int) {
	return file_sideshow_coffeeco_v1_customer_proto_rawDescGZIP(), []int{0, 0}
}

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stores the name, could be a nickname, or real name
	// We don't assume to capture more data than necessary
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Stores the unique identifier for a given customer
	// Note: the `murmur3(name:uuid)` can be used to generate a strong hash if needed
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// Stores the pointer to the 'first time' we've encountered a Customer
	FirstSeen *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=first_seen,json=firstSeen,proto3" json:"first_seen,omitempty"`
	// Stores the value of Guest or Member
	CustomerType Customer_CustomerType `protobuf:"varint,4,opt,name=customer_type,json=customerType,proto3,enum=sideshow.coffeeco.v1.Customer_CustomerType" json:"customer_type,omitempty"`
	// Stores the optional loyalty_member_id. This value can identify the
	// customer without the need to lookup any personally identifiable information
	// can only be associated with known Member - since collecting points requires some compute and dollar spend
	LoyaltyMemberId string `protobuf:"bytes,5,opt,name=loyalty_member_id,json=loyaltyMemberId,proto3" json:"loyalty_member_id,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sideshow_coffeeco_v1_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_sideshow_coffeeco_v1_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_sideshow_coffeeco_v1_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Customer) GetFirstSeen() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstSeen
	}
	return nil
}

func (x *Customer) GetCustomerType() Customer_CustomerType {
	if x != nil {
		return x.CustomerType
	}
	return Customer_CUSTOMER_TYPE_UNSPECIFIED
}

func (x *Customer) GetLoyaltyMemberId() string {
	if x != nil {
		return x.LoyaltyMemberId
	}
	return ""
}

var File_sideshow_coffeeco_v1_customer_proto protoreflect.FileDescriptor

var file_sideshow_coffeeco_v1_customer_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65,
	0x65, 0x63, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x2e,
	0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x03, 0x0a, 0x08, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x72, 0x05, 0x10, 0x04,
	0x18, 0x80, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0xad, 0x01, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x72, 0xba, 0x48, 0x6f,
	0xba, 0x01, 0x6c, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x68,
	0x65, 0x5f, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x12, 0x48, 0x77, 0x65, 0x20, 0x63, 0x61, 0x6e,
	0x27, 0x74, 0x20, 0x66, 0x69, 0x72, 0x73, 0x74, 0x20, 0x73, 0x65, 0x65, 0x20, 0x61, 0x20, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x20, 0x4e, 0x6f, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x74,
	0x72, 0x61, 0x76, 0x65, 0x6c, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x20, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x0b, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3c, 0x3d, 0x20, 0x6e, 0x6f, 0x77, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x12, 0x5e, 0x0a, 0x0d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x66,
	0x66, 0x65, 0x65, 0x63, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0c,
	0xba, 0x48, 0x09, 0xc8, 0x01, 0x01, 0x82, 0x01, 0x03, 0x22, 0x01, 0x00, 0x52, 0x0c, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x11, 0x6c, 0x6f,
	0x79, 0x61, 0x6c, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52,
	0x0f, 0x6c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x60, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x19, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x17, 0x0a, 0x13, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x02, 0x42, 0xe2, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x73,
	0x68, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x6f, 0x2e, 0x76, 0x31, 0x42,
	0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x69, 0x72, 0x63, 0x75, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77,
	0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x66,
	0x66, 0x65, 0x65, 0x63, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x43, 0x58, 0xaa, 0x02, 0x14,
	0x53, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63,
	0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x53, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x5c,
	0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x20, 0x53, 0x69,
	0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x5c, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x6f, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x16, 0x53, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x6f, 0x66, 0x66, 0x65,
	0x65, 0x63, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sideshow_coffeeco_v1_customer_proto_rawDescOnce sync.Once
	file_sideshow_coffeeco_v1_customer_proto_rawDescData = file_sideshow_coffeeco_v1_customer_proto_rawDesc
)

func file_sideshow_coffeeco_v1_customer_proto_rawDescGZIP() []byte {
	file_sideshow_coffeeco_v1_customer_proto_rawDescOnce.Do(func() {
		file_sideshow_coffeeco_v1_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_sideshow_coffeeco_v1_customer_proto_rawDescData)
	})
	return file_sideshow_coffeeco_v1_customer_proto_rawDescData
}

var file_sideshow_coffeeco_v1_customer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sideshow_coffeeco_v1_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sideshow_coffeeco_v1_customer_proto_goTypes = []interface{}{
	(Customer_CustomerType)(0),    // 0: sideshow.coffeeco.v1.Customer.CustomerType
	(*Customer)(nil),              // 1: sideshow.coffeeco.v1.Customer
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_sideshow_coffeeco_v1_customer_proto_depIdxs = []int32{
	2, // 0: sideshow.coffeeco.v1.Customer.first_seen:type_name -> google.protobuf.Timestamp
	0, // 1: sideshow.coffeeco.v1.Customer.customer_type:type_name -> sideshow.coffeeco.v1.Customer.CustomerType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sideshow_coffeeco_v1_customer_proto_init() }
func file_sideshow_coffeeco_v1_customer_proto_init() {
	if File_sideshow_coffeeco_v1_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sideshow_coffeeco_v1_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
			RawDescriptor: file_sideshow_coffeeco_v1_customer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sideshow_coffeeco_v1_customer_proto_goTypes,
		DependencyIndexes: file_sideshow_coffeeco_v1_customer_proto_depIdxs,
		EnumInfos:         file_sideshow_coffeeco_v1_customer_proto_enumTypes,
		MessageInfos:      file_sideshow_coffeeco_v1_customer_proto_msgTypes,
	}.Build()
	File_sideshow_coffeeco_v1_customer_proto = out.File
	file_sideshow_coffeeco_v1_customer_proto_rawDesc = nil
	file_sideshow_coffeeco_v1_customer_proto_goTypes = nil
	file_sideshow_coffeeco_v1_customer_proto_depIdxs = nil
}
