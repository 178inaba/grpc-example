// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package file

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

type FileRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FileResponse struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileResponse) Reset()         { *m = FileResponse{} }
func (m *FileResponse) String() string { return proto.CompactTextString(m) }
func (*FileResponse) ProtoMessage()    {}
func (*FileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}

func (m *FileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileResponse.Unmarshal(m, b)
}
func (m *FileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileResponse.Marshal(b, m, deterministic)
}
func (m *FileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileResponse.Merge(m, src)
}
func (m *FileResponse) XXX_Size() int {
	return xxx_messageInfo_FileResponse.Size(m)
}
func (m *FileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileResponse proto.InternalMessageInfo

func (m *FileResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*FileRequest)(nil), "file.FileRequest")
	proto.RegisterType((*FileResponse)(nil), "file.FileResponse")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x4c, 0xb9, 0xb8, 0xdd, 0x32,
	0x73, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x4a, 0x62, 0x49, 0xa2,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0xad, 0xa4, 0xc4, 0xc5, 0x03, 0xd1, 0x56, 0x5c,
	0x90, 0x9f, 0x57, 0x0c, 0x56, 0x53, 0x9c, 0x59, 0x05, 0xd1, 0xc7, 0x1c, 0x04, 0x66, 0x1b, 0x39,
	0x40, 0x8c, 0x0e, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x32, 0xe4, 0x62, 0x0b, 0x2d, 0xc8,
	0xc9, 0x4f, 0x4c, 0x11, 0x12, 0xd4, 0x03, 0x3b, 0x03, 0xc9, 0x5e, 0x29, 0x21, 0x64, 0x21, 0x88,
	0x99, 0x1a, 0x8c, 0x49, 0x6c, 0x60, 0x97, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x76,
	0x75, 0x64, 0xb7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadClient, error)
}

type fileServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileServiceClient(cc *grpc.ClientConn) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileService_serviceDesc.Streams[0], "/file.FileService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadClient{stream}
	return x, nil
}

type FileService_UploadClient interface {
	Send(*FileRequest) error
	CloseAndRecv() (*FileResponse, error)
	grpc.ClientStream
}

type fileServiceUploadClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadClient) Send(m *FileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadClient) CloseAndRecv() (*FileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	Upload(FileService_UploadServer) error
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) Upload(srv FileService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).Upload(&fileServiceUploadServer{stream})
}

type FileService_UploadServer interface {
	SendAndClose(*FileResponse) error
	Recv() (*FileRequest, error)
	grpc.ServerStream
}

type fileServiceUploadServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadServer) SendAndClose(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadServer) Recv() (*FileRequest, error) {
	m := new(FileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "file.proto",
}
