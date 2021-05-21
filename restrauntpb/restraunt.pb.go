// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: restraunt.proto

package __

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

// The request message containing the result for inserting new data.
type InsertResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InsertResult) Reset() {
	*x = InsertResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restraunt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertResult) ProtoMessage() {}

func (x *InsertResult) ProtoReflect() protoreflect.Message {
	mi := &file_restraunt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertResult.ProtoReflect.Descriptor instead.
func (*InsertResult) Descriptor() ([]byte, []int) {
	return file_restraunt_proto_rawDescGZIP(), []int{0}
}

func (x *InsertResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

// The response message containing the greetings
type Restraunt_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Restraunt_Request) Reset() {
	*x = Restraunt_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restraunt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restraunt_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restraunt_Request) ProtoMessage() {}

func (x *Restraunt_Request) ProtoReflect() protoreflect.Message {
	mi := &file_restraunt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restraunt_Request.ProtoReflect.Descriptor instead.
func (*Restraunt_Request) Descriptor() ([]byte, []int) {
	return file_restraunt_proto_rawDescGZIP(), []int{1}
}

func (x *Restraunt_Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Restraunt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *Restraunt) Reset() {
	*x = Restraunt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restraunt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restraunt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restraunt) ProtoMessage() {}

func (x *Restraunt) ProtoReflect() protoreflect.Message {
	mi := &file_restraunt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restraunt.ProtoReflect.Descriptor instead.
func (*Restraunt) Descriptor() ([]byte, []int) {
	return file_restraunt_proto_rawDescGZIP(), []int{2}
}

func (x *Restraunt) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restraunt) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Restraunt) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_restraunt_proto protoreflect.FileDescriptor

var file_restraunt_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x0d,
	0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75,
	0x6e, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32, 0xa4, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x5f, 0x52, 0x65,
	0x63, 0x69, 0x65, 0x76, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x0d, 0x41, 0x64,
	0x64, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e,
	0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x2e, 0x69, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4d, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x12, 0x1c,
	0x2e, 0x72, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72,
	0x61, 0x75, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x75, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x72, 0x61, 0x75,
	0x6e, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restraunt_proto_rawDescOnce sync.Once
	file_restraunt_proto_rawDescData = file_restraunt_proto_rawDesc
)

func file_restraunt_proto_rawDescGZIP() []byte {
	file_restraunt_proto_rawDescOnce.Do(func() {
		file_restraunt_proto_rawDescData = protoimpl.X.CompressGZIP(file_restraunt_proto_rawDescData)
	})
	return file_restraunt_proto_rawDescData
}

var file_restraunt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_restraunt_proto_goTypes = []interface{}{
	(*InsertResult)(nil),      // 0: restraunt.insert_result
	(*Restraunt_Request)(nil), // 1: restraunt.Restraunt_Request
	(*Restraunt)(nil),         // 2: restraunt.restraunt
}
var file_restraunt_proto_depIdxs = []int32{
	2, // 0: restraunt.Add_Recieve_data.Add_restraunt:input_type -> restraunt.restraunt
	1, // 1: restraunt.Add_Recieve_data.Get_restraunt:input_type -> restraunt.Restraunt_Request
	0, // 2: restraunt.Add_Recieve_data.Add_restraunt:output_type -> restraunt.insert_result
	1, // 3: restraunt.Add_Recieve_data.Get_restraunt:output_type -> restraunt.Restraunt_Request
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_restraunt_proto_init() }
func file_restraunt_proto_init() {
	if File_restraunt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restraunt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertResult); i {
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
		file_restraunt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restraunt_Request); i {
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
		file_restraunt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restraunt); i {
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
			RawDescriptor: file_restraunt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_restraunt_proto_goTypes,
		DependencyIndexes: file_restraunt_proto_depIdxs,
		MessageInfos:      file_restraunt_proto_msgTypes,
	}.Build()
	File_restraunt_proto = out.File
	file_restraunt_proto_rawDesc = nil
	file_restraunt_proto_goTypes = nil
	file_restraunt_proto_depIdxs = nil
}
