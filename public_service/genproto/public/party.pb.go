// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: party.proto

package public

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

type CreatePartyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,2,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,3,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreatePartyReq) Reset() {
	*x = CreatePartyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePartyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartyReq) ProtoMessage() {}

func (x *CreatePartyReq) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePartyReq.ProtoReflect.Descriptor instead.
func (*CreatePartyReq) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePartyReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePartyReq) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *CreatePartyReq) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *CreatePartyReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetPartyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slogan      string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	OpenedDate  string `protobuf:"bytes,4,opt,name=opened_date,json=openedDate,proto3" json:"opened_date,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetPartyRes) Reset() {
	*x = GetPartyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartyRes) ProtoMessage() {}

func (x *GetPartyRes) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartyRes.ProtoReflect.Descriptor instead.
func (*GetPartyRes) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{1}
}

func (x *GetPartyRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPartyRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPartyRes) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *GetPartyRes) GetOpenedDate() string {
	if x != nil {
		return x.OpenedDate
	}
	return ""
}

func (x *GetPartyRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetAllPartyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parties []*GetPartyRes `protobuf:"bytes,1,rep,name=Parties,proto3" json:"Parties,omitempty"`
}

func (x *GetAllPartyRes) Reset() {
	*x = GetAllPartyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPartyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPartyRes) ProtoMessage() {}

func (x *GetAllPartyRes) ProtoReflect() protoreflect.Message {
	mi := &file_party_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPartyRes.ProtoReflect.Descriptor instead.
func (*GetAllPartyRes) Descriptor() ([]byte, []int) {
	return file_party_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllPartyRes) GetParties() []*GetPartyRes {
	if x != nil {
		return x.Parties
	}
	return nil
}

var File_party_proto protoreflect.FileDescriptor

var file_party_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x1a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f,
	0x67, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x67,
	0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x67, 0x61, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x52, 0x07, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x32, 0xfb, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x1a, 0x16, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x1a, 0x0c, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x79,
	0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_proto_rawDescOnce sync.Once
	file_party_proto_rawDescData = file_party_proto_rawDesc
)

func file_party_proto_rawDescGZIP() []byte {
	file_party_proto_rawDescOnce.Do(func() {
		file_party_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_proto_rawDescData)
	})
	return file_party_proto_rawDescData
}

var file_party_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_party_proto_goTypes = []interface{}{
	(*CreatePartyReq)(nil), // 0: public.CreatePartyReq
	(*GetPartyRes)(nil),    // 1: public.GetPartyRes
	(*GetAllPartyRes)(nil), // 2: public.GetAllPartyRes
	(*ById)(nil),           // 3: public.ById
	(*Filter)(nil),         // 4: public.Filter
	(*Void)(nil),           // 5: public.Void
}
var file_party_proto_depIdxs = []int32{
	1, // 0: public.GetAllPartyRes.Parties:type_name -> public.GetPartyRes
	0, // 1: public.PartyService.Create:input_type -> public.CreatePartyReq
	3, // 2: public.PartyService.GetById:input_type -> public.ById
	4, // 3: public.PartyService.GetAll:input_type -> public.Filter
	1, // 4: public.PartyService.Update:input_type -> public.GetPartyRes
	3, // 5: public.PartyService.Delete:input_type -> public.ById
	5, // 6: public.PartyService.Create:output_type -> public.Void
	1, // 7: public.PartyService.GetById:output_type -> public.GetPartyRes
	2, // 8: public.PartyService.GetAll:output_type -> public.GetAllPartyRes
	5, // 9: public.PartyService.Update:output_type -> public.Void
	5, // 10: public.PartyService.Delete:output_type -> public.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_party_proto_init() }
func file_party_proto_init() {
	if File_party_proto != nil {
		return
	}
	file_public_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_party_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePartyReq); i {
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
		file_party_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartyRes); i {
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
		file_party_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPartyRes); i {
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
			RawDescriptor: file_party_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_party_proto_goTypes,
		DependencyIndexes: file_party_proto_depIdxs,
		MessageInfos:      file_party_proto_msgTypes,
	}.Build()
	File_party_proto = out.File
	file_party_proto_rawDesc = nil
	file_party_proto_goTypes = nil
	file_party_proto_depIdxs = nil
}
