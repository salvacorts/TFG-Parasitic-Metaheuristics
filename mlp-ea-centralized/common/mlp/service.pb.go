// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package mlp

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MLPMsg struct {
	Mlp                  *MultiLayerNetwork `protobuf:"bytes,1,opt,name=mlp,proto3" json:"mlp,omitempty"`
	IndividualID         string             `protobuf:"bytes,2,opt,name=individualID,proto3" json:"individualID,omitempty"`
	Evaluated            bool               `protobuf:"varint,3,opt,name=evaluated,proto3" json:"evaluated,omitempty"`
	Fitness              float64            `protobuf:"fixed64,4,opt,name=Fitness,proto3" json:"Fitness,omitempty"`
	ClientID             string             `protobuf:"bytes,5,opt,name=clientID,proto3" json:"clientID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MLPMsg) Reset()         { *m = MLPMsg{} }
func (m *MLPMsg) String() string { return proto.CompactTextString(m) }
func (*MLPMsg) ProtoMessage()    {}
func (*MLPMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *MLPMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MLPMsg.Unmarshal(m, b)
}
func (m *MLPMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MLPMsg.Marshal(b, m, deterministic)
}
func (m *MLPMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MLPMsg.Merge(m, src)
}
func (m *MLPMsg) XXX_Size() int {
	return xxx_messageInfo_MLPMsg.Size(m)
}
func (m *MLPMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MLPMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MLPMsg proto.InternalMessageInfo

func (m *MLPMsg) GetMlp() *MultiLayerNetwork {
	if m != nil {
		return m.Mlp
	}
	return nil
}

func (m *MLPMsg) GetIndividualID() string {
	if m != nil {
		return m.IndividualID
	}
	return ""
}

func (m *MLPMsg) GetEvaluated() bool {
	if m != nil {
		return m.Evaluated
	}
	return false
}

func (m *MLPMsg) GetFitness() float64 {
	if m != nil {
		return m.Fitness
	}
	return 0
}

func (m *MLPMsg) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

type ProblemDescription struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Epochs               int64    `protobuf:"varint,2,opt,name=epochs,proto3" json:"epochs,omitempty"`
	Folds                int64    `protobuf:"varint,3,opt,name=folds,proto3" json:"folds,omitempty"`
	TrainDataset         string   `protobuf:"bytes,4,opt,name=trainDataset,proto3" json:"trainDataset,omitempty"`
	Classes              []string `protobuf:"bytes,5,rep,name=classes,proto3" json:"classes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProblemDescription) Reset()         { *m = ProblemDescription{} }
func (m *ProblemDescription) String() string { return proto.CompactTextString(m) }
func (*ProblemDescription) ProtoMessage()    {}
func (*ProblemDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}
func (m *ProblemDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProblemDescription.Unmarshal(m, b)
}
func (m *ProblemDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProblemDescription.Marshal(b, m, deterministic)
}
func (m *ProblemDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProblemDescription.Merge(m, src)
}
func (m *ProblemDescription) XXX_Size() int {
	return xxx_messageInfo_ProblemDescription.Size(m)
}
func (m *ProblemDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ProblemDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ProblemDescription proto.InternalMessageInfo

func (m *ProblemDescription) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *ProblemDescription) GetEpochs() int64 {
	if m != nil {
		return m.Epochs
	}
	return 0
}

func (m *ProblemDescription) GetFolds() int64 {
	if m != nil {
		return m.Folds
	}
	return 0
}

func (m *ProblemDescription) GetTrainDataset() string {
	if m != nil {
		return m.TrainDataset
	}
	return ""
}

func (m *ProblemDescription) GetClasses() []string {
	if m != nil {
		return m.Classes
	}
	return nil
}

type Stats struct {
	Evaluations          int64    `protobuf:"varint,2,opt,name=evaluations,proto3" json:"evaluations,omitempty"`
	BestFitness          float64  `protobuf:"fixed64,3,opt,name=bestFitness,proto3" json:"bestFitness,omitempty"`
	AvgFitness           float64  `protobuf:"fixed64,4,opt,name=avgFitness,proto3" json:"avgFitness,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}
func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetEvaluations() int64 {
	if m != nil {
		return m.Evaluations
	}
	return 0
}

func (m *Stats) GetBestFitness() float64 {
	if m != nil {
		return m.BestFitness
	}
	return 0
}

func (m *Stats) GetAvgFitness() float64 {
	if m != nil {
		return m.AvgFitness
	}
	return 0
}

func init() {
	proto.RegisterType((*MLPMsg)(nil), "mlp.MLPMsg")
	proto.RegisterType((*ProblemDescription)(nil), "mlp.ProblemDescription")
	proto.RegisterType((*Stats)(nil), "mlp.Stats")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x6b, 0x12, 0x92, 0x09, 0x95, 0xd0, 0x0a, 0x82, 0x65, 0x21, 0x64, 0xf9, 0xe4, 0x93,
	0x83, 0xda, 0x13, 0x47, 0x90, 0x4b, 0x15, 0xd4, 0xa0, 0xca, 0x7c, 0xc1, 0xda, 0x9e, 0xb8, 0xab,
	0xae, 0xbd, 0xd6, 0xee, 0x38, 0x55, 0xce, 0xfc, 0x06, 0x5f, 0xc0, 0x57, 0x22, 0xaf, 0x93, 0x62,
	0x0b, 0x72, 0xb1, 0xfc, 0xde, 0xbe, 0x37, 0x3b, 0x33, 0xfb, 0xe0, 0xd2, 0xa0, 0xde, 0x8b, 0x1c,
	0xe3, 0x46, 0x2b, 0x52, 0xcc, 0xad, 0x64, 0xe3, 0x5f, 0x97, 0x82, 0x1e, 0xda, 0x2c, 0xce, 0x55,
	0xb5, 0x2e, 0x95, 0xe4, 0x75, 0xb9, 0xb6, 0xa7, 0x59, 0xbb, 0x5b, 0x37, 0x74, 0x68, 0xd0, 0xac,
	0xb1, 0x6a, 0xe8, 0xd0, 0x7f, 0x7b, 0xa7, 0xbf, 0xa8, 0x64, 0xd3, 0xff, 0x86, 0xbf, 0x1d, 0x98,
	0x6d, 0xef, 0xee, 0xb7, 0xa6, 0x64, 0x11, 0x74, 0x15, 0x3d, 0x27, 0x70, 0xa2, 0xe5, 0xd5, 0x2a,
	0xee, 0x34, 0xdb, 0x56, 0x92, 0xb8, 0xe3, 0x07, 0xd4, 0xdf, 0x91, 0x9e, 0x94, 0x7e, 0x4c, 0x3b,
	0x09, 0x0b, 0xe1, 0x95, 0xa8, 0x0b, 0xb1, 0x17, 0x45, 0xcb, 0xe5, 0x26, 0xf1, 0x2e, 0x02, 0x27,
	0x5a, 0xa4, 0x23, 0x8e, 0xbd, 0x87, 0x05, 0xee, 0xb9, 0x6c, 0x39, 0x61, 0xe1, 0xb9, 0x81, 0x13,
	0xcd, 0xd3, 0xbf, 0x04, 0xf3, 0xe0, 0xe5, 0x57, 0x41, 0x35, 0x1a, 0xe3, 0xbd, 0x08, 0x9c, 0xc8,
	0x49, 0x4f, 0x90, 0xf9, 0x30, 0xcf, 0xa5, 0xc0, 0x9a, 0x36, 0x89, 0x37, 0xb5, 0x75, 0x9f, 0x71,
	0xf8, 0xcb, 0x01, 0x76, 0xaf, 0x55, 0x26, 0xb1, 0x4a, 0xd0, 0xe4, 0x5a, 0x34, 0x24, 0x54, 0x3d,
	0xb2, 0x38, 0x63, 0x0b, 0x5b, 0xc1, 0x0c, 0x1b, 0x95, 0x3f, 0x18, 0xdb, 0xa4, 0x9b, 0x1e, 0x11,
	0x7b, 0x03, 0xd3, 0x9d, 0x92, 0x85, 0xb1, 0xad, 0xb9, 0x69, 0x0f, 0xba, 0xc1, 0x48, 0x73, 0x51,
	0x27, 0x9c, 0xb8, 0x41, 0xb2, 0xbd, 0x2d, 0xd2, 0x11, 0xd7, 0xb5, 0x9e, 0x4b, 0x6e, 0x0c, 0x1a,
	0x6f, 0x1a, 0xb8, 0xd1, 0x22, 0x3d, 0xc1, 0xf0, 0x11, 0xa6, 0x3f, 0x88, 0x93, 0x61, 0x01, 0x2c,
	0x8f, 0xa3, 0x0a, 0x55, 0x9f, 0x6e, 0x1e, 0x52, 0x9d, 0x22, 0x43, 0x43, 0xa7, 0x1d, 0xb8, 0x76,
	0x07, 0x43, 0x8a, 0x7d, 0x00, 0xe0, 0xfb, 0x72, 0xbc, 0xa4, 0x01, 0x73, 0xf5, 0xf3, 0x02, 0x2e,
	0x13, 0x61, 0x48, 0x8b, 0xac, 0x25, 0x2c, 0x6e, 0x3e, 0xb3, 0x6f, 0xf0, 0xf6, 0x16, 0xe9, 0x3f,
	0xfb, 0x59, 0xc5, 0xa5, 0x52, 0xa5, 0x3c, 0xe6, 0x26, 0x6b, 0x77, 0xf1, 0x4d, 0x17, 0x06, 0xff,
	0x9d, 0x7d, 0xe3, 0x7f, 0x0d, 0xe1, 0x84, 0x7d, 0x84, 0xf9, 0x2d, 0x52, 0x3f, 0xcd, 0x39, 0x3b,
	0x58, 0xbb, 0xd5, 0x84, 0x13, 0xf6, 0x09, 0x5e, 0x7f, 0x51, 0x5a, 0xab, 0xa7, 0xcd, 0x73, 0x0a,
	0xce, 0x3a, 0x97, 0x7d, 0xb8, 0x6c, 0xec, 0x7a, 0x6b, 0x8a, 0xd4, 0xea, 0x7a, 0x60, 0x1d, 0x4a,
	0xfc, 0x33, 0x75, 0xc2, 0x49, 0x36, 0xb3, 0xcc, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd,
	0xa4, 0xbe, 0x65, 0x1b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DistributedEAClient is the client API for DistributedEA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DistributedEAClient interface {
	GetProblemDescription(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProblemDescription, error)
	GetStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Stats, error)
	BorrowIndividual(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MLPMsg, error)
	ReturnIndividual(ctx context.Context, in *MLPMsg, opts ...grpc.CallOption) (*empty.Empty, error)
}

type distributedEAClient struct {
	cc *grpc.ClientConn
}

func NewDistributedEAClient(cc *grpc.ClientConn) DistributedEAClient {
	return &distributedEAClient{cc}
}

func (c *distributedEAClient) GetProblemDescription(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProblemDescription, error) {
	out := new(ProblemDescription)
	err := c.cc.Invoke(ctx, "/mlp.DistributedEA/GetProblemDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedEAClient) GetStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := c.cc.Invoke(ctx, "/mlp.DistributedEA/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedEAClient) BorrowIndividual(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*MLPMsg, error) {
	out := new(MLPMsg)
	err := c.cc.Invoke(ctx, "/mlp.DistributedEA/BorrowIndividual", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributedEAClient) ReturnIndividual(ctx context.Context, in *MLPMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mlp.DistributedEA/ReturnIndividual", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistributedEAServer is the server API for DistributedEA service.
type DistributedEAServer interface {
	GetProblemDescription(context.Context, *empty.Empty) (*ProblemDescription, error)
	GetStats(context.Context, *empty.Empty) (*Stats, error)
	BorrowIndividual(context.Context, *empty.Empty) (*MLPMsg, error)
	ReturnIndividual(context.Context, *MLPMsg) (*empty.Empty, error)
}

// UnimplementedDistributedEAServer can be embedded to have forward compatible implementations.
type UnimplementedDistributedEAServer struct {
}

func (*UnimplementedDistributedEAServer) GetProblemDescription(ctx context.Context, req *empty.Empty) (*ProblemDescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemDescription not implemented")
}
func (*UnimplementedDistributedEAServer) GetStats(ctx context.Context, req *empty.Empty) (*Stats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (*UnimplementedDistributedEAServer) BorrowIndividual(ctx context.Context, req *empty.Empty) (*MLPMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BorrowIndividual not implemented")
}
func (*UnimplementedDistributedEAServer) ReturnIndividual(ctx context.Context, req *MLPMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnIndividual not implemented")
}

func RegisterDistributedEAServer(s *grpc.Server, srv DistributedEAServer) {
	s.RegisterService(&_DistributedEA_serviceDesc, srv)
}

func _DistributedEA_GetProblemDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedEAServer).GetProblemDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlp.DistributedEA/GetProblemDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedEAServer).GetProblemDescription(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedEA_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedEAServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlp.DistributedEA/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedEAServer).GetStats(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedEA_BorrowIndividual_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedEAServer).BorrowIndividual(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlp.DistributedEA/BorrowIndividual",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedEAServer).BorrowIndividual(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributedEA_ReturnIndividual_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MLPMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributedEAServer).ReturnIndividual(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlp.DistributedEA/ReturnIndividual",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributedEAServer).ReturnIndividual(ctx, req.(*MLPMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistributedEA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mlp.DistributedEA",
	HandlerType: (*DistributedEAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProblemDescription",
			Handler:    _DistributedEA_GetProblemDescription_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _DistributedEA_GetStats_Handler,
		},
		{
			MethodName: "BorrowIndividual",
			Handler:    _DistributedEA_BorrowIndividual_Handler,
		},
		{
			MethodName: "ReturnIndividual",
			Handler:    _DistributedEA_ReturnIndividual_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
