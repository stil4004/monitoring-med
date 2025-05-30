// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api_m/operation.proto

package api_m

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

// TODO access then userId
type WithdrawCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestTag string `protobuf:"bytes,1,opt,name=destTag,proto3" json:"destTag,omitempty"` // Monero XRP, destTag from client side
	// string callbackId = 2; // delete nahui
	InternalId string  `protobuf:"bytes,2,opt,name=internalId,proto3" json:"internalId,omitempty"`
	TotpToken  string  `protobuf:"bytes,3,opt,name=totpToken,proto3" json:"totpToken,omitempty"`
	Receiver   string  `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount     float64 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Token      string  `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	Code       string  `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"` // seed phrase or email code
}

func (x *WithdrawCryptoRequest) Reset() {
	*x = WithdrawCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_m_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawCryptoRequest) ProtoMessage() {}

func (x *WithdrawCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_m_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawCryptoRequest.ProtoReflect.Descriptor instead.
func (*WithdrawCryptoRequest) Descriptor() ([]byte, []int) {
	return file_api_m_operation_proto_rawDescGZIP(), []int{0}
}

func (x *WithdrawCryptoRequest) GetDestTag() string {
	if x != nil {
		return x.DestTag
	}
	return ""
}

func (x *WithdrawCryptoRequest) GetInternalId() string {
	if x != nil {
		return x.InternalId
	}
	return ""
}

func (x *WithdrawCryptoRequest) GetTotpToken() string {
	if x != nil {
		return x.TotpToken
	}
	return ""
}

func (x *WithdrawCryptoRequest) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *WithdrawCryptoRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *WithdrawCryptoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *WithdrawCryptoRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type WithdrawCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackerId string `protobuf:"bytes,1,opt,name=trackerId,proto3" json:"trackerId,omitempty"`
}

func (x *WithdrawCryptoResponse) Reset() {
	*x = WithdrawCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_m_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawCryptoResponse) ProtoMessage() {}

func (x *WithdrawCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_m_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawCryptoResponse.ProtoReflect.Descriptor instead.
func (*WithdrawCryptoResponse) Descriptor() ([]byte, []int) {
	return file_api_m_operation_proto_rawDescGZIP(), []int{1}
}

func (x *WithdrawCryptoResponse) GetTrackerId() string {
	if x != nil {
		return x.TrackerId
	}
	return ""
}

type GetUserBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnCreate bool   `protobuf:"varint,1,opt,name=onCreate,proto3" json:"onCreate,omitempty"` // for front balance - comission when creating offer
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`        //  string access = 3;
}

func (x *GetUserBalanceRequest) Reset() {
	*x = GetUserBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_m_operation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBalanceRequest) ProtoMessage() {}

func (x *GetUserBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_m_operation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetUserBalanceRequest) Descriptor() ([]byte, []int) {
	return file_api_m_operation_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserBalanceRequest) GetOnCreate() bool {
	if x != nil {
		return x.OnCreate
	}
	return false
}

func (x *GetUserBalanceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetUserBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Balance `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"` //  int64 userId = 2; // delete nahui
}

func (x *GetUserBalanceResponse) Reset() {
	*x = GetUserBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_m_operation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBalanceResponse) ProtoMessage() {}

func (x *GetUserBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_m_operation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetUserBalanceResponse) Descriptor() ([]byte, []int) {
	return file_api_m_operation_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserBalanceResponse) GetData() []*Balance {
	if x != nil {
		return x.Data
	}
	return nil
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available float64 `protobuf:"fixed64,1,opt,name=available,proto3" json:"available,omitempty"`
	Frozen    float64 `protobuf:"fixed64,2,opt,name=frozen,proto3" json:"frozen,omitempty"`
	// double usdt = 3; // <>
	// double bonusBalance = 4; // <>
	All   float64 `protobuf:"fixed64,3,opt,name=all,proto3" json:"all,omitempty"`
	Token string  `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_m_operation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_api_m_operation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_api_m_operation_proto_rawDescGZIP(), []int{4}
}

func (x *Balance) GetAvailable() float64 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *Balance) GetFrozen() float64 {
	if x != nil {
		return x.Frozen
	}
	return 0
}

func (x *Balance) GetAll() float64 {
	if x != nil {
		return x.All
	}
	return 0
}

func (x *Balance) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_api_m_operation_proto protoreflect.FileDescriptor

var file_api_m_operation_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x22, 0xcd,
	0x01, 0x0a, 0x15, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x54,
	0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x36,
	0x0a, 0x16, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x5f,
	0x6d, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x67, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x7a,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61,
	0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa9, 0x01, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x0e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6d,
	0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x2e, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_m_operation_proto_rawDescOnce sync.Once
	file_api_m_operation_proto_rawDescData = file_api_m_operation_proto_rawDesc
)

func file_api_m_operation_proto_rawDescGZIP() []byte {
	file_api_m_operation_proto_rawDescOnce.Do(func() {
		file_api_m_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_m_operation_proto_rawDescData)
	})
	return file_api_m_operation_proto_rawDescData
}

var file_api_m_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_m_operation_proto_goTypes = []interface{}{
	(*WithdrawCryptoRequest)(nil),  // 0: api_m.WithdrawCryptoRequest
	(*WithdrawCryptoResponse)(nil), // 1: api_m.WithdrawCryptoResponse
	(*GetUserBalanceRequest)(nil),  // 2: api_m.GetUserBalanceRequest
	(*GetUserBalanceResponse)(nil), // 3: api_m.GetUserBalanceResponse
	(*Balance)(nil),                // 4: api_m.Balance
}
var file_api_m_operation_proto_depIdxs = []int32{
	4, // 0: api_m.GetUserBalanceResponse.data:type_name -> api_m.Balance
	0, // 1: api_m.operation.WithdrawCrypto:input_type -> api_m.WithdrawCryptoRequest
	2, // 2: api_m.operation.GetUserBalance:input_type -> api_m.GetUserBalanceRequest
	1, // 3: api_m.operation.WithdrawCrypto:output_type -> api_m.WithdrawCryptoResponse
	3, // 4: api_m.operation.GetUserBalance:output_type -> api_m.GetUserBalanceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_m_operation_proto_init() }
func file_api_m_operation_proto_init() {
	if File_api_m_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_m_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawCryptoRequest); i {
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
		file_api_m_operation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawCryptoResponse); i {
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
		file_api_m_operation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBalanceRequest); i {
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
		file_api_m_operation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBalanceResponse); i {
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
		file_api_m_operation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
			RawDescriptor: file_api_m_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_m_operation_proto_goTypes,
		DependencyIndexes: file_api_m_operation_proto_depIdxs,
		MessageInfos:      file_api_m_operation_proto_msgTypes,
	}.Build()
	File_api_m_operation_proto = out.File
	file_api_m_operation_proto_rawDesc = nil
	file_api_m_operation_proto_goTypes = nil
	file_api_m_operation_proto_depIdxs = nil
}
