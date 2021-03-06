// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: function.proto

package v1 // import "aif.io/api/networking/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Function struct {
	// function uri eg:/csf/ord_service1 or  /springboot/test/controler1
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// function's own micro-service
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// route to cluster's subset
	Subset string `protobuf:"bytes,3,opt,name=subset,proto3" json:"subset,omitempty"`
	// function's timeout
	Timeout int32 `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// traffic limit config
	Limit                *TrafficLimit `protobuf:"bytes,5,opt,name=limit" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_function_f00e80c034e42a68, []int{0}
}
func (m *Function) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Function.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(dst, src)
}
func (m *Function) XXX_Size() int {
	return m.Size()
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Function) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Function) GetSubset() string {
	if m != nil {
		return m.Subset
	}
	return ""
}

func (m *Function) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Function) GetLimit() *TrafficLimit {
	if m != nil {
		return m.Limit
	}
	return nil
}

type TrafficLimit struct {
	// limit type
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// limit triger value
	Threshhold int32 `protobuf:"varint,2,opt,name=threshhold,proto3" json:"threshhold,omitempty"`
	// limit triger timeWindow
	TimeWindow           int32    `protobuf:"varint,3,opt,name=timeWindow,proto3" json:"timeWindow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficLimit) Reset()         { *m = TrafficLimit{} }
func (m *TrafficLimit) String() string { return proto.CompactTextString(m) }
func (*TrafficLimit) ProtoMessage()    {}
func (*TrafficLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_function_f00e80c034e42a68, []int{1}
}
func (m *TrafficLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrafficLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrafficLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TrafficLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLimit.Merge(dst, src)
}
func (m *TrafficLimit) XXX_Size() int {
	return m.Size()
}
func (m *TrafficLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLimit.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLimit proto.InternalMessageInfo

func (m *TrafficLimit) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TrafficLimit) GetThreshhold() int32 {
	if m != nil {
		return m.Threshhold
	}
	return 0
}

func (m *TrafficLimit) GetTimeWindow() int32 {
	if m != nil {
		return m.TimeWindow
	}
	return 0
}

func init() {
	proto.RegisterType((*Function)(nil), "aif.io.api.networking.v1.Function")
	proto.RegisterType((*TrafficLimit)(nil), "aif.io.api.networking.v1.TrafficLimit")
}
func (m *Function) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Function) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uri) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFunction(dAtA, i, uint64(len(m.Uri)))
		i += copy(dAtA[i:], m.Uri)
	}
	if len(m.Code) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFunction(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if len(m.Subset) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFunction(dAtA, i, uint64(len(m.Subset)))
		i += copy(dAtA[i:], m.Subset)
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintFunction(dAtA, i, uint64(m.Timeout))
	}
	if m.Limit != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintFunction(dAtA, i, uint64(m.Limit.Size()))
		n1, err := m.Limit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TrafficLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrafficLimit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFunction(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Threshhold != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFunction(dAtA, i, uint64(m.Threshhold))
	}
	if m.TimeWindow != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintFunction(dAtA, i, uint64(m.TimeWindow))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFunction(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Function) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	l = len(m.Subset)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovFunction(uint64(m.Timeout))
	}
	if m.Limit != nil {
		l = m.Limit.Size()
		n += 1 + l + sovFunction(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TrafficLimit) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovFunction(uint64(l))
	}
	if m.Threshhold != 0 {
		n += 1 + sovFunction(uint64(m.Threshhold))
	}
	if m.TimeWindow != 0 {
		n += 1 + sovFunction(uint64(m.TimeWindow))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFunction(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFunction(x uint64) (n int) {
	return sovFunction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Function) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Function: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Function: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Limit == nil {
				m.Limit = &TrafficLimit{}
			}
			if err := m.Limit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFunction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TrafficLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFunction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TrafficLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrafficLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFunction
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshhold", wireType)
			}
			m.Threshhold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshhold |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeWindow", wireType)
			}
			m.TimeWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeWindow |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFunction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFunction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFunction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFunction
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
					return 0, ErrIntOverflowFunction
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFunction
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFunction
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFunction
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFunction(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFunction = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFunction   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("function.proto", fileDescriptor_function_f00e80c034e42a68) }

var fileDescriptor_function_f00e80c034e42a68 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0xc6, 0x59, 0xef, 0x72, 0xea, 0x28, 0x22, 0x5b, 0xc8, 0x56, 0x21, 0x5c, 0x71, 0xa4, 0xda,
	0x70, 0xda, 0x5a, 0x59, 0x58, 0x59, 0x05, 0x41, 0xb0, 0xcb, 0x9f, 0x8d, 0x19, 0xbc, 0xec, 0x84,
	0xcd, 0xe4, 0x0e, 0xdf, 0xc6, 0xc7, 0xb1, 0xf4, 0x11, 0x24, 0x4f, 0x22, 0x59, 0xa3, 0x5e, 0x73,
	0xdd, 0x37, 0xbf, 0x99, 0x81, 0x1f, 0x1f, 0x5c, 0x54, 0xbd, 0x2d, 0x18, 0xc9, 0xea, 0xd6, 0x11,
	0x93, 0x54, 0x19, 0x56, 0x1a, 0x49, 0x67, 0x2d, 0x6a, 0x6b, 0x78, 0x47, 0xee, 0x15, 0xed, 0x8b,
	0xde, 0xae, 0x97, 0xef, 0x02, 0x4e, 0xee, 0xa7, 0x63, 0x79, 0x09, 0xb3, 0xde, 0xa1, 0x12, 0x91,
	0x88, 0x4f, 0xd3, 0x31, 0x4a, 0x09, 0xf3, 0x82, 0x4a, 0xa3, 0x8e, 0x3c, 0xf2, 0x59, 0x5e, 0xc1,
	0xa2, 0xeb, 0xf3, 0xce, 0xb0, 0x9a, 0x79, 0x3a, 0x4d, 0x52, 0xc1, 0x31, 0x63, 0x63, 0xa8, 0x67,
	0x35, 0x8f, 0x44, 0x1c, 0xa4, 0xbf, 0xa3, 0xbc, 0x85, 0x60, 0x83, 0x0d, 0xb2, 0x0a, 0x22, 0x11,
	0x9f, 0x5d, 0xaf, 0xf4, 0x21, 0x1d, 0xfd, 0xe8, 0xb2, 0xaa, 0xc2, 0xe2, 0x61, 0xbc, 0x4e, 0x7f,
	0x9e, 0x96, 0x39, 0x9c, 0xef, 0xe3, 0xd1, 0x89, 0xdf, 0x5a, 0x33, 0x69, 0xfa, 0x2c, 0x43, 0x00,
	0xae, 0x9d, 0xe9, 0xea, 0x9a, 0x36, 0xa5, 0xb7, 0x0d, 0xd2, 0x3d, 0xe2, 0xf7, 0xd8, 0x98, 0x27,
	0xb4, 0x25, 0xed, 0xbc, 0xf7, 0xb8, 0xff, 0x23, 0x77, 0xab, 0x8f, 0x21, 0x14, 0x9f, 0x43, 0x28,
	0xbe, 0x86, 0x50, 0x3c, 0x4f, 0x75, 0x25, 0x59, 0x8b, 0xc9, 0xbf, 0x5f, 0xb2, 0x5d, 0xe7, 0x0b,
	0xdf, 0xe7, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x51, 0x23, 0xc8, 0x61, 0x01, 0x00,
	0x00,
}
