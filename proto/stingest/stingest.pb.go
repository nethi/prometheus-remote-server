// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stingest.proto

package stingestproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StatsIngestRequestStatus int32

const (
	StatsIngestRequestStatus_STINGEST_REQ_STATUS_UNKNOWN StatsIngestRequestStatus = 0
	StatsIngestRequestStatus_STINGEST_REQ_STATUS_QUEUED  StatsIngestRequestStatus = 1
	StatsIngestRequestStatus_STINGEST_REQ_STATUS_ACTIVE  StatsIngestRequestStatus = 2
)

var StatsIngestRequestStatus_name = map[int32]string{
	0: "STINGEST_REQ_STATUS_UNKNOWN",
	1: "STINGEST_REQ_STATUS_QUEUED",
	2: "STINGEST_REQ_STATUS_ACTIVE",
}

var StatsIngestRequestStatus_value = map[string]int32{
	"STINGEST_REQ_STATUS_UNKNOWN": 0,
	"STINGEST_REQ_STATUS_QUEUED":  1,
	"STINGEST_REQ_STATUS_ACTIVE":  2,
}

func (x StatsIngestRequestStatus) String() string {
	return proto.EnumName(StatsIngestRequestStatus_name, int32(x))
}

func (StatsIngestRequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{0}
}

type StatsIngestErrorCode int32

const (
	StatsIngestErrorCode_OK              StatsIngestErrorCode = 0
	StatsIngestErrorCode_STINGEST_FAILED StatsIngestErrorCode = 1
	StatsIngestErrorCode_STINVALID_REQ   StatsIngestErrorCode = 2
	StatsIngestErrorCode_STNOT_READY     StatsIngestErrorCode = 3
)

var StatsIngestErrorCode_name = map[int32]string{
	0: "OK",
	1: "STINGEST_FAILED",
	2: "STINVALID_REQ",
	3: "STNOT_READY",
}

var StatsIngestErrorCode_value = map[string]int32{
	"OK":              0,
	"STINGEST_FAILED": 1,
	"STINVALID_REQ":   2,
	"STNOT_READY":     3,
}

func (x StatsIngestErrorCode) String() string {
	return proto.EnumName(StatsIngestErrorCode_name, int32(x))
}

func (StatsIngestErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{1}
}

type StatsFlushErrorCode int32

const (
	StatsFlushErrorCode_STFLUSH_OK     StatsFlushErrorCode = 0
	StatsFlushErrorCode_STFLUSH_FAILED StatsFlushErrorCode = 1
)

var StatsFlushErrorCode_name = map[int32]string{
	0: "STFLUSH_OK",
	1: "STFLUSH_FAILED",
}

var StatsFlushErrorCode_value = map[string]int32{
	"STFLUSH_OK":     0,
	"STFLUSH_FAILED": 1,
}

func (x StatsFlushErrorCode) String() string {
	return proto.EnumName(StatsFlushErrorCode_name, int32(x))
}

func (StatsFlushErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{2}
}

type StatsIngestStatusErrorCode int32

const (
	StatsIngestStatusErrorCode_STSTATUS_OK     StatsIngestStatusErrorCode = 0
	StatsIngestStatusErrorCode_STSTATUS_FAILED StatsIngestStatusErrorCode = 1
)

var StatsIngestStatusErrorCode_name = map[int32]string{
	0: "STSTATUS_OK",
	1: "STSTATUS_FAILED",
}

var StatsIngestStatusErrorCode_value = map[string]int32{
	"STSTATUS_OK":     0,
	"STSTATUS_FAILED": 1,
}

func (x StatsIngestStatusErrorCode) String() string {
	return proto.EnumName(StatsIngestStatusErrorCode_name, int32(x))
}

func (StatsIngestStatusErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{3}
}

type StatsIngestRequest struct {
	Account              string                   `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	DFpath               string                   `protobuf:"bytes,2,opt,name=DFpath,proto3" json:"DFpath,omitempty"`
	MFpath               string                   `protobuf:"bytes,3,opt,name=MFpath,proto3" json:"MFpath,omitempty"`
	InstanceId           string                   `protobuf:"bytes,4,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	Type                 string                   `protobuf:"bytes,5,opt,name=Type,proto3" json:"Type,omitempty"`
	DTs                  int64                    `protobuf:"varint,6,opt,name=DTs,proto3" json:"DTs,omitempty"`
	Status               StatsIngestRequestStatus `protobuf:"varint,7,opt,name=Status,proto3,enum=stingestproto.StatsIngestRequestStatus" json:"Status,omitempty"`
	Epoch                int64                    `protobuf:"varint,8,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *StatsIngestRequest) Reset()         { *m = StatsIngestRequest{} }
func (m *StatsIngestRequest) String() string { return proto.CompactTextString(m) }
func (*StatsIngestRequest) ProtoMessage()    {}
func (*StatsIngestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{0}
}

func (m *StatsIngestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsIngestRequest.Unmarshal(m, b)
}
func (m *StatsIngestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsIngestRequest.Marshal(b, m, deterministic)
}
func (m *StatsIngestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsIngestRequest.Merge(m, src)
}
func (m *StatsIngestRequest) XXX_Size() int {
	return xxx_messageInfo_StatsIngestRequest.Size(m)
}
func (m *StatsIngestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsIngestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsIngestRequest proto.InternalMessageInfo

func (m *StatsIngestRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *StatsIngestRequest) GetDFpath() string {
	if m != nil {
		return m.DFpath
	}
	return ""
}

func (m *StatsIngestRequest) GetMFpath() string {
	if m != nil {
		return m.MFpath
	}
	return ""
}

func (m *StatsIngestRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *StatsIngestRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StatsIngestRequest) GetDTs() int64 {
	if m != nil {
		return m.DTs
	}
	return 0
}

func (m *StatsIngestRequest) GetStatus() StatsIngestRequestStatus {
	if m != nil {
		return m.Status
	}
	return StatsIngestRequestStatus_STINGEST_REQ_STATUS_UNKNOWN
}

func (m *StatsIngestRequest) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

type StatsIngestResponse struct {
	Err                  StatsIngestErrorCode `protobuf:"varint,1,opt,name=Err,proto3,enum=stingestproto.StatsIngestErrorCode" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StatsIngestResponse) Reset()         { *m = StatsIngestResponse{} }
func (m *StatsIngestResponse) String() string { return proto.CompactTextString(m) }
func (*StatsIngestResponse) ProtoMessage()    {}
func (*StatsIngestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{1}
}

func (m *StatsIngestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsIngestResponse.Unmarshal(m, b)
}
func (m *StatsIngestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsIngestResponse.Marshal(b, m, deterministic)
}
func (m *StatsIngestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsIngestResponse.Merge(m, src)
}
func (m *StatsIngestResponse) XXX_Size() int {
	return xxx_messageInfo_StatsIngestResponse.Size(m)
}
func (m *StatsIngestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsIngestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatsIngestResponse proto.InternalMessageInfo

func (m *StatsIngestResponse) GetErr() StatsIngestErrorCode {
	if m != nil {
		return m.Err
	}
	return StatsIngestErrorCode_OK
}

type StatsFlushRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsFlushRequest) Reset()         { *m = StatsFlushRequest{} }
func (m *StatsFlushRequest) String() string { return proto.CompactTextString(m) }
func (*StatsFlushRequest) ProtoMessage()    {}
func (*StatsFlushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{2}
}

func (m *StatsFlushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsFlushRequest.Unmarshal(m, b)
}
func (m *StatsFlushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsFlushRequest.Marshal(b, m, deterministic)
}
func (m *StatsFlushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsFlushRequest.Merge(m, src)
}
func (m *StatsFlushRequest) XXX_Size() int {
	return xxx_messageInfo_StatsFlushRequest.Size(m)
}
func (m *StatsFlushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsFlushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsFlushRequest proto.InternalMessageInfo

func (m *StatsFlushRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type StatsFlushResponse struct {
	Err                  StatsFlushErrorCode `protobuf:"varint,1,opt,name=Err,proto3,enum=stingestproto.StatsFlushErrorCode" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StatsFlushResponse) Reset()         { *m = StatsFlushResponse{} }
func (m *StatsFlushResponse) String() string { return proto.CompactTextString(m) }
func (*StatsFlushResponse) ProtoMessage()    {}
func (*StatsFlushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{3}
}

func (m *StatsFlushResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsFlushResponse.Unmarshal(m, b)
}
func (m *StatsFlushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsFlushResponse.Marshal(b, m, deterministic)
}
func (m *StatsFlushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsFlushResponse.Merge(m, src)
}
func (m *StatsFlushResponse) XXX_Size() int {
	return xxx_messageInfo_StatsFlushResponse.Size(m)
}
func (m *StatsFlushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsFlushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatsFlushResponse proto.InternalMessageInfo

func (m *StatsFlushResponse) GetErr() StatsFlushErrorCode {
	if m != nil {
		return m.Err
	}
	return StatsFlushErrorCode_STFLUSH_OK
}

type StatsIngestStatusRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsIngestStatusRequest) Reset()         { *m = StatsIngestStatusRequest{} }
func (m *StatsIngestStatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatsIngestStatusRequest) ProtoMessage()    {}
func (*StatsIngestStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{4}
}

func (m *StatsIngestStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsIngestStatusRequest.Unmarshal(m, b)
}
func (m *StatsIngestStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsIngestStatusRequest.Marshal(b, m, deterministic)
}
func (m *StatsIngestStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsIngestStatusRequest.Merge(m, src)
}
func (m *StatsIngestStatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatsIngestStatusRequest.Size(m)
}
func (m *StatsIngestStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsIngestStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsIngestStatusRequest proto.InternalMessageInfo

func (m *StatsIngestStatusRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type StatsIngestStatusResponse struct {
	Requests             []*StatsIngestRequest      `protobuf:"bytes,1,rep,name=Requests,proto3" json:"Requests,omitempty"`
	Err                  StatsIngestStatusErrorCode `protobuf:"varint,2,opt,name=Err,proto3,enum=stingestproto.StatsIngestStatusErrorCode" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *StatsIngestStatusResponse) Reset()         { *m = StatsIngestStatusResponse{} }
func (m *StatsIngestStatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatsIngestStatusResponse) ProtoMessage()    {}
func (*StatsIngestStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_302ef5ec6fbb31ff, []int{5}
}

func (m *StatsIngestStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsIngestStatusResponse.Unmarshal(m, b)
}
func (m *StatsIngestStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsIngestStatusResponse.Marshal(b, m, deterministic)
}
func (m *StatsIngestStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsIngestStatusResponse.Merge(m, src)
}
func (m *StatsIngestStatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatsIngestStatusResponse.Size(m)
}
func (m *StatsIngestStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsIngestStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatsIngestStatusResponse proto.InternalMessageInfo

func (m *StatsIngestStatusResponse) GetRequests() []*StatsIngestRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *StatsIngestStatusResponse) GetErr() StatsIngestStatusErrorCode {
	if m != nil {
		return m.Err
	}
	return StatsIngestStatusErrorCode_STSTATUS_OK
}

func init() {
	proto.RegisterEnum("stingestproto.StatsIngestRequestStatus", StatsIngestRequestStatus_name, StatsIngestRequestStatus_value)
	proto.RegisterEnum("stingestproto.StatsIngestErrorCode", StatsIngestErrorCode_name, StatsIngestErrorCode_value)
	proto.RegisterEnum("stingestproto.StatsFlushErrorCode", StatsFlushErrorCode_name, StatsFlushErrorCode_value)
	proto.RegisterEnum("stingestproto.StatsIngestStatusErrorCode", StatsIngestStatusErrorCode_name, StatsIngestStatusErrorCode_value)
	proto.RegisterType((*StatsIngestRequest)(nil), "stingestproto.StatsIngestRequest")
	proto.RegisterType((*StatsIngestResponse)(nil), "stingestproto.StatsIngestResponse")
	proto.RegisterType((*StatsFlushRequest)(nil), "stingestproto.StatsFlushRequest")
	proto.RegisterType((*StatsFlushResponse)(nil), "stingestproto.StatsFlushResponse")
	proto.RegisterType((*StatsIngestStatusRequest)(nil), "stingestproto.StatsIngestStatusRequest")
	proto.RegisterType((*StatsIngestStatusResponse)(nil), "stingestproto.StatsIngestStatusResponse")
}

func init() {
	proto.RegisterFile("stingest.proto", fileDescriptor_302ef5ec6fbb31ff)
}

var fileDescriptor_302ef5ec6fbb31ff = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xed, 0xc6, 0x2d, 0x53, 0xea, 0xba, 0xd3, 0x0a, 0x2d, 0x41, 0x2a, 0xa9, 0x39, 0x34,
	0x44, 0x22, 0x87, 0x50, 0x0e, 0x08, 0x21, 0x64, 0x1a, 0x07, 0x4c, 0x83, 0xa3, 0xd8, 0x4e, 0x2b,
	0x4e, 0x51, 0x70, 0x2d, 0x82, 0x84, 0x6c, 0xe3, 0x75, 0x90, 0x10, 0x1f, 0x02, 0xff, 0xca, 0x05,
	0x79, 0x77, 0x1d, 0x6c, 0x88, 0xeb, 0xdb, 0xbe, 0x99, 0x37, 0xef, 0xcd, 0xdb, 0x5d, 0xd0, 0x68,
	0xf6, 0x39, 0xfa, 0x14, 0xd2, 0xac, 0x9f, 0xa4, 0x71, 0x16, 0xe3, 0x7e, 0x81, 0x19, 0x34, 0x7e,
	0x4b, 0x80, 0x5e, 0xb6, 0xc8, 0xa8, 0xcd, 0x8a, 0x6e, 0xf8, 0x75, 0x15, 0xd2, 0x0c, 0x09, 0xec,
	0x98, 0x41, 0x10, 0xaf, 0xa2, 0x8c, 0x48, 0x1d, 0xa9, 0x7b, 0xc7, 0x2d, 0x20, 0xde, 0x03, 0x75,
	0x38, 0x4a, 0x16, 0xd9, 0x92, 0xc8, 0xac, 0x21, 0x50, 0x5e, 0x7f, 0xcf, 0xeb, 0x0a, 0xaf, 0x73,
	0x84, 0x27, 0x00, 0x76, 0x44, 0xb3, 0x45, 0x14, 0x84, 0xf6, 0x0d, 0xd9, 0x66, 0xbd, 0x52, 0x05,
	0x11, 0xb6, 0xfd, 0xef, 0x49, 0x48, 0x5a, 0xac, 0xc3, 0xce, 0xa8, 0x83, 0x32, 0xf4, 0x29, 0x51,
	0x3b, 0x52, 0x57, 0x71, 0xf3, 0x23, 0xbe, 0x02, 0x35, 0xdf, 0x72, 0x45, 0xc9, 0x4e, 0x47, 0xea,
	0x6a, 0x83, 0xb3, 0x7e, 0x25, 0x46, 0xff, 0xff, 0x08, 0x9c, 0xee, 0x8a, 0x31, 0x3c, 0x86, 0x96,
	0x95, 0xc4, 0xc1, 0x92, 0xec, 0x32, 0x51, 0x0e, 0x8c, 0x31, 0x1c, 0x55, 0x26, 0x69, 0x12, 0x47,
	0x34, 0xc4, 0x67, 0xa0, 0x58, 0x69, 0xca, 0x92, 0x6b, 0x83, 0x47, 0xf5, 0x56, 0x56, 0x9a, 0xc6,
	0xe9, 0x45, 0x7c, 0x13, 0xba, 0x39, 0xdf, 0x78, 0x02, 0x87, 0xac, 0x39, 0xfa, 0xb2, 0xa2, 0xcb,
	0xc6, 0x9b, 0x34, 0xde, 0x89, 0x9b, 0x17, 0x74, 0xe1, 0x7d, 0x5e, 0xf6, 0x36, 0x36, 0x79, 0x33,
	0xfe, 0x3f, 0xd6, 0xe7, 0x40, 0x4a, 0x7b, 0x89, 0xec, 0x8d, 0x1b, 0xfc, 0x94, 0xe0, 0xfe, 0x86,
	0x31, 0xb1, 0xc9, 0x4b, 0xd8, 0x15, 0x12, 0x94, 0x48, 0x1d, 0xa5, 0xbb, 0x37, 0x38, 0x6d, 0xbc,
	0x75, 0x77, 0x3d, 0x82, 0x2f, 0x78, 0x10, 0x99, 0x05, 0x79, 0x5c, 0x3f, 0xc9, 0x5d, 0xab, 0x79,
	0x7a, 0x3f, 0x2a, 0x79, 0x2a, 0x4f, 0x8a, 0x0f, 0xe1, 0x81, 0xe7, 0xdb, 0xce, 0x1b, 0xcb, 0xf3,
	0xe7, 0xae, 0x35, 0x9d, 0x7b, 0xbe, 0xe9, 0xcf, 0xbc, 0xf9, 0xcc, 0xb9, 0x74, 0x26, 0xd7, 0x8e,
	0xbe, 0x85, 0x27, 0xd0, 0xde, 0x44, 0x98, 0xce, 0xac, 0x99, 0x35, 0xd4, 0xa5, 0xba, 0xbe, 0x79,
	0xe1, 0xdb, 0x57, 0x96, 0x2e, 0xf7, 0xae, 0xe1, 0x78, 0xd3, 0x23, 0xa3, 0x0a, 0xf2, 0xe4, 0x52,
	0xdf, 0xc2, 0x23, 0x38, 0x58, 0xcf, 0x8f, 0x4c, 0x7b, 0xcc, 0x44, 0x0f, 0x61, 0x3f, 0x2f, 0x5e,
	0x99, 0x63, 0x7b, 0x98, 0xab, 0xea, 0x32, 0x1e, 0xc0, 0x9e, 0xe7, 0x3b, 0x93, 0xdc, 0xc4, 0x1c,
	0x7e, 0xd0, 0x95, 0xde, 0x73, 0xf1, 0xdd, 0xaa, 0x2f, 0x88, 0x1a, 0x80, 0xe7, 0x8f, 0xc6, 0x33,
	0xef, 0xed, 0x9c, 0xe9, 0x23, 0x68, 0x05, 0x2e, 0xe4, 0x7b, 0xaf, 0xa1, 0x5d, 0x7f, 0x67, 0xdc,
	0x49, 0xc4, 0xf8, 0xbb, 0xa2, 0x28, 0x14, 0x1a, 0x83, 0x5f, 0x32, 0x68, 0x65, 0x91, 0x6f, 0x01,
	0x4e, 0x41, 0xe5, 0x00, 0x9b, 0xdf, 0xb6, 0x6d, 0xdc, 0x46, 0xe1, 0x9f, 0xc6, 0xd8, 0x42, 0x07,
	0x5a, 0x2c, 0x1f, 0x76, 0x6a, 0x3f, 0x6f, 0x21, 0x78, 0x7a, 0x0b, 0x63, 0xad, 0x17, 0xc0, 0xdd,
	0x72, 0x68, 0x3c, 0x6b, 0xfa, 0x4a, 0x85, 0x7a, 0xb7, 0x99, 0x58, 0x98, 0x7c, 0x54, 0x19, 0xe5,
	0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xbc, 0x9d, 0x50, 0x2e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StatsIngestSvcClient is the client API for StatsIngestSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsIngestSvcClient interface {
	// StatsIngest API
	Ingest(ctx context.Context, in *StatsIngestRequest, opts ...grpc.CallOption) (*StatsIngestResponse, error)
	// StatsFlush API
	Flush(ctx context.Context, in *StatsFlushRequest, opts ...grpc.CallOption) (*StatsFlushResponse, error)
	// StatsStatus API
	IngestStatus(ctx context.Context, in *StatsIngestStatusRequest, opts ...grpc.CallOption) (*StatsIngestStatusResponse, error)
}

type statsIngestSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsIngestSvcClient(cc grpc.ClientConnInterface) StatsIngestSvcClient {
	return &statsIngestSvcClient{cc}
}

func (c *statsIngestSvcClient) Ingest(ctx context.Context, in *StatsIngestRequest, opts ...grpc.CallOption) (*StatsIngestResponse, error) {
	out := new(StatsIngestResponse)
	err := c.cc.Invoke(ctx, "/stingestproto.StatsIngestSvc/Ingest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsIngestSvcClient) Flush(ctx context.Context, in *StatsFlushRequest, opts ...grpc.CallOption) (*StatsFlushResponse, error) {
	out := new(StatsFlushResponse)
	err := c.cc.Invoke(ctx, "/stingestproto.StatsIngestSvc/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsIngestSvcClient) IngestStatus(ctx context.Context, in *StatsIngestStatusRequest, opts ...grpc.CallOption) (*StatsIngestStatusResponse, error) {
	out := new(StatsIngestStatusResponse)
	err := c.cc.Invoke(ctx, "/stingestproto.StatsIngestSvc/IngestStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsIngestSvcServer is the server API for StatsIngestSvc service.
type StatsIngestSvcServer interface {
	// StatsIngest API
	Ingest(context.Context, *StatsIngestRequest) (*StatsIngestResponse, error)
	// StatsFlush API
	Flush(context.Context, *StatsFlushRequest) (*StatsFlushResponse, error)
	// StatsStatus API
	IngestStatus(context.Context, *StatsIngestStatusRequest) (*StatsIngestStatusResponse, error)
}

// UnimplementedStatsIngestSvcServer can be embedded to have forward compatible implementations.
type UnimplementedStatsIngestSvcServer struct {
}

func (*UnimplementedStatsIngestSvcServer) Ingest(ctx context.Context, req *StatsIngestRequest) (*StatsIngestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ingest not implemented")
}
func (*UnimplementedStatsIngestSvcServer) Flush(ctx context.Context, req *StatsFlushRequest) (*StatsFlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (*UnimplementedStatsIngestSvcServer) IngestStatus(ctx context.Context, req *StatsIngestStatusRequest) (*StatsIngestStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IngestStatus not implemented")
}

func RegisterStatsIngestSvcServer(s *grpc.Server, srv StatsIngestSvcServer) {
	s.RegisterService(&_StatsIngestSvc_serviceDesc, srv)
}

func _StatsIngestSvc_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsIngestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsIngestSvcServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stingestproto.StatsIngestSvc/Ingest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsIngestSvcServer).Ingest(ctx, req.(*StatsIngestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsIngestSvc_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsFlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsIngestSvcServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stingestproto.StatsIngestSvc/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsIngestSvcServer).Flush(ctx, req.(*StatsFlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsIngestSvc_IngestStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsIngestStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsIngestSvcServer).IngestStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stingestproto.StatsIngestSvc/IngestStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsIngestSvcServer).IngestStatus(ctx, req.(*StatsIngestStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatsIngestSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stingestproto.StatsIngestSvc",
	HandlerType: (*StatsIngestSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ingest",
			Handler:    _StatsIngestSvc_Ingest_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _StatsIngestSvc_Flush_Handler,
		},
		{
			MethodName: "IngestStatus",
			Handler:    _StatsIngestSvc_IngestStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stingest.proto",
}
