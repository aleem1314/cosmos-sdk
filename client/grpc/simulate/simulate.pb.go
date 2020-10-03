// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/base/simulate/v1beta1/simulate.proto

package simulate

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SimulateRequest is the request type for the SimulateServiceService.Simulate
// RPC method.
type SimulateRequest struct {
	// tx is the transaction to simulate.
	Tx *tx.Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *SimulateRequest) Reset()         { *m = SimulateRequest{} }
func (m *SimulateRequest) String() string { return proto.CompactTextString(m) }
func (*SimulateRequest) ProtoMessage()    {}
func (*SimulateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_503c836d80bb2d47, []int{0}
}
func (m *SimulateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimulateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimulateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimulateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulateRequest.Merge(m, src)
}
func (m *SimulateRequest) XXX_Size() int {
	return m.Size()
}
func (m *SimulateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimulateRequest proto.InternalMessageInfo

func (m *SimulateRequest) GetTx() *tx.Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

// SimulateResponse is the response type for the
// SimulateServiceService.SimulateRPC method.
type SimulateResponse struct {
	// gas_info is the information about gas used in the simulation.
	GasInfo *types.GasInfo `protobuf:"bytes,1,opt,name=gas_info,json=gasInfo,proto3" json:"gas_info,omitempty"`
	// result is the result of the simulation.
	Result *types.Result `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *SimulateResponse) Reset()         { *m = SimulateResponse{} }
func (m *SimulateResponse) String() string { return proto.CompactTextString(m) }
func (*SimulateResponse) ProtoMessage()    {}
func (*SimulateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_503c836d80bb2d47, []int{1}
}
func (m *SimulateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimulateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimulateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimulateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulateResponse.Merge(m, src)
}
func (m *SimulateResponse) XXX_Size() int {
	return m.Size()
}
func (m *SimulateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimulateResponse proto.InternalMessageInfo

func (m *SimulateResponse) GetGasInfo() *types.GasInfo {
	if m != nil {
		return m.GasInfo
	}
	return nil
}

func (m *SimulateResponse) GetResult() *types.Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*SimulateRequest)(nil), "cosmos.base.simulate.v1beta1.SimulateRequest")
	proto.RegisterType((*SimulateResponse)(nil), "cosmos.base.simulate.v1beta1.SimulateResponse")
}

func init() {
	proto.RegisterFile("cosmos/base/simulate/v1beta1/simulate.proto", fileDescriptor_503c836d80bb2d47)
}

var fileDescriptor_503c836d80bb2d47 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x7b, 0x19, 0xfa, 0x96, 0x7b, 0x87, 0xf7, 0x25, 0x20, 0x94, 0x50, 0x42, 0x8d, 0x28,
	0x05, 0xe9, 0x1d, 0xad, 0x4b, 0x07, 0x27, 0x17, 0x11, 0xb7, 0xd4, 0xc9, 0x45, 0x2e, 0xf1, 0x1a,
	0x0f, 0xd3, 0xbb, 0x98, 0x7b, 0x52, 0x32, 0x3b, 0x3a, 0x09, 0x4e, 0xfe, 0x11, 0xfe, 0x1f, 0x8e,
	0x05, 0x17, 0x47, 0x69, 0xfd, 0x43, 0xa4, 0xc9, 0x25, 0x2d, 0x82, 0xd2, 0x29, 0xb9, 0xe7, 0xf9,
	0x7c, 0xbf, 0xcf, 0x2f, 0x7c, 0x18, 0x2a, 0x3d, 0x55, 0x9a, 0x06, 0x4c, 0x73, 0xaa, 0xc5, 0x34,
	0x8b, 0x19, 0x70, 0x3a, 0x1b, 0x04, 0x1c, 0xd8, 0xa0, 0x0e, 0x90, 0x24, 0x55, 0xa0, 0xec, 0x4e,
	0x09, 0x93, 0x15, 0x4c, 0xea, 0x9c, 0x81, 0x9d, 0x4e, 0xa4, 0x54, 0x14, 0x73, 0xca, 0x12, 0x41,
	0x99, 0x94, 0x0a, 0x18, 0x08, 0x25, 0x75, 0xa9, 0x75, 0xf6, 0x36, 0x0b, 0xb1, 0x20, 0x14, 0x75,
	0x91, 0xd5, 0xc3, 0x40, 0x8e, 0x81, 0x20, 0xaf, 0xb3, 0x90, 0x97, 0x39, 0x6f, 0x84, 0xff, 0x8d,
	0x4d, 0x49, 0x9f, 0xdf, 0x65, 0x5c, 0x83, 0xbd, 0x8f, 0x2d, 0xc8, 0xdb, 0xa8, 0x8b, 0x7a, 0x7f,
	0x87, 0x3b, 0xc4, 0x34, 0x07, 0x79, 0xd5, 0x11, 0xb9, 0xc8, 0x7d, 0x0b, 0x72, 0xef, 0x01, 0xe1,
	0xff, 0x6b, 0xa9, 0x4e, 0x94, 0xd4, 0xdc, 0x3e, 0xc6, 0xad, 0x88, 0xe9, 0x2b, 0x21, 0x27, 0xca,
	0x38, 0xec, 0x92, 0xcd, 0xf1, 0x8a, 0xae, 0x2a, 0xa3, 0x53, 0xa6, 0xcf, 0xe4, 0x44, 0xf9, 0x7f,
	0xa2, 0xf2, 0xc7, 0x1e, 0xe1, 0x66, 0xca, 0x75, 0x16, 0x43, 0xdb, 0x2a, 0xb4, 0xdd, 0x9f, 0xb5,
	0x7e, 0xc1, 0xf9, 0x86, 0x1f, 0xbe, 0xa0, 0xf5, 0x1c, 0x63, 0x9e, 0xce, 0x44, 0xc8, 0xed, 0x67,
	0x84, 0x5b, 0x55, 0xcc, 0xee, 0x93, 0xdf, 0xb6, 0x4c, 0xbe, 0xed, 0xc0, 0x21, 0xdb, 0xe2, 0xe5,
	0xdc, 0x1e, 0xb9, 0x7f, 0xfb, 0x7c, 0xb2, 0x7a, 0xde, 0x01, 0xdd, 0xea, 0xf2, 0x27, 0xe7, 0xaf,
	0x0b, 0x17, 0xcd, 0x17, 0x2e, 0xfa, 0x58, 0xb8, 0xe8, 0x71, 0xe9, 0x36, 0xe6, 0x4b, 0xb7, 0xf1,
	0xbe, 0x74, 0x1b, 0x97, 0x83, 0x48, 0xc0, 0x4d, 0x16, 0x90, 0x50, 0x4d, 0x2b, 0xaf, 0xf2, 0xd3,
	0xd7, 0xd7, 0xb7, 0x34, 0x8c, 0x05, 0x97, 0x40, 0xa3, 0x34, 0x09, 0x6b, 0xb3, 0xa0, 0x59, 0x9c,
	0xf2, 0xe8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x16, 0xed, 0x90, 0x76, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimulateServiceClient is the client API for SimulateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimulateServiceClient interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error)
}

type simulateServiceClient struct {
	cc grpc1.ClientConn
}

func NewSimulateServiceClient(cc grpc1.ClientConn) SimulateServiceClient {
	return &simulateServiceClient{cc}
}

func (c *simulateServiceClient) Simulate(ctx context.Context, in *SimulateRequest, opts ...grpc.CallOption) (*SimulateResponse, error) {
	out := new(SimulateResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.simulate.v1beta1.SimulateService/Simulate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimulateServiceServer is the server API for SimulateService service.
type SimulateServiceServer interface {
	// Simulate simulates executing a transaction for estimating gas usage.
	Simulate(context.Context, *SimulateRequest) (*SimulateResponse, error)
}

// UnimplementedSimulateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSimulateServiceServer struct {
}

func (*UnimplementedSimulateServiceServer) Simulate(ctx context.Context, req *SimulateRequest) (*SimulateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simulate not implemented")
}

func RegisterSimulateServiceServer(s grpc1.Server, srv SimulateServiceServer) {
	s.RegisterService(&_SimulateService_serviceDesc, srv)
}

func _SimulateService_Simulate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulateServiceServer).Simulate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.simulate.v1beta1.SimulateService/Simulate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulateServiceServer).Simulate(ctx, req.(*SimulateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimulateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.simulate.v1beta1.SimulateService",
	HandlerType: (*SimulateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simulate",
			Handler:    _SimulateService_Simulate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/simulate/v1beta1/simulate.proto",
}

func (m *SimulateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimulateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimulateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSimulate(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimulateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimulateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimulateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSimulate(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.GasInfo != nil {
		{
			size, err := m.GasInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSimulate(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSimulate(dAtA []byte, offset int, v uint64) int {
	offset -= sovSimulate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SimulateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovSimulate(uint64(l))
	}
	return n
}

func (m *SimulateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GasInfo != nil {
		l = m.GasInfo.Size()
		n += 1 + l + sovSimulate(uint64(l))
	}
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovSimulate(uint64(l))
	}
	return n
}

func sovSimulate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSimulate(x uint64) (n int) {
	return sovSimulate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SimulateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSimulate
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SimulateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimulateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimulate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSimulate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSimulate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &tx.Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSimulate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSimulate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSimulate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SimulateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSimulate
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SimulateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimulateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimulate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSimulate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSimulate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GasInfo == nil {
				m.GasInfo = &types.GasInfo{}
			}
			if err := m.GasInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSimulate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSimulate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSimulate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &types.Result{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSimulate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSimulate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSimulate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSimulate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSimulate
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSimulate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSimulate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSimulate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSimulate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSimulate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSimulate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSimulate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSimulate = fmt.Errorf("proto: unexpected end of group")
)
