// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client.proto

package pbgo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import actor "github.com/AsynkronIT/protoactor-go/actor"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 客服端
type ClientMessageReq struct {
	Pid  *actor.PID `protobuf:"bytes,1,opt,name=Pid" json:"Pid,omitempty"`
	Body []byte     `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *ClientMessageReq) Reset()                    { *m = ClientMessageReq{} }
func (*ClientMessageReq) ProtoMessage()               {}
func (*ClientMessageReq) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

func (m *ClientMessageReq) GetPid() *actor.PID {
	if m != nil {
		return m.Pid
	}
	return nil
}

func (m *ClientMessageReq) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type ClientMessageRsp struct {
	Pid  *actor.PID `protobuf:"bytes,1,opt,name=Pid" json:"Pid,omitempty"`
	Body []byte     `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *ClientMessageRsp) Reset()                    { *m = ClientMessageRsp{} }
func (*ClientMessageRsp) ProtoMessage()               {}
func (*ClientMessageRsp) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{1} }

func (m *ClientMessageRsp) GetPid() *actor.PID {
	if m != nil {
		return m.Pid
	}
	return nil
}

func (m *ClientMessageRsp) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientMessageReq)(nil), "pbgo.ClientMessageReq")
	proto.RegisterType((*ClientMessageRsp)(nil), "pbgo.ClientMessageRsp")
}
func (this *ClientMessageReq) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClientMessageReq)
	if !ok {
		that2, ok := that.(ClientMessageReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Pid.Equal(that1.Pid) {
		return false
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	return true
}
func (this *ClientMessageRsp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClientMessageRsp)
	if !ok {
		that2, ok := that.(ClientMessageRsp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Pid.Equal(that1.Pid) {
		return false
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	return true
}
func (this *ClientMessageReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pbgo.ClientMessageReq{")
	if this.Pid != nil {
		s = append(s, "Pid: "+fmt.Sprintf("%#v", this.Pid)+",\n")
	}
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClientMessageRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pbgo.ClientMessageRsp{")
	if this.Pid != nil {
		s = append(s, "Pid: "+fmt.Sprintf("%#v", this.Pid)+",\n")
	}
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringClient(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ClientMessageReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientMessageReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pid != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.Pid.Size()))
		n1, err := m.Pid.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	return i, nil
}

func (m *ClientMessageRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientMessageRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pid != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.Pid.Size()))
		n2, err := m.Pid.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	return i, nil
}

func encodeFixed64Client(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Client(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintClient(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClientMessageReq) Size() (n int) {
	var l int
	_ = l
	if m.Pid != nil {
		l = m.Pid.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	return n
}

func (m *ClientMessageRsp) Size() (n int) {
	var l int
	_ = l
	if m.Pid != nil {
		l = m.Pid.Size()
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	return n
}

func sovClient(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClient(x uint64) (n int) {
	return sovClient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClientMessageReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClientMessageReq{`,
		`Pid:` + strings.Replace(fmt.Sprintf("%v", this.Pid), "PID", "actor.PID", 1) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClientMessageRsp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClientMessageRsp{`,
		`Pid:` + strings.Replace(fmt.Sprintf("%v", this.Pid), "PID", "actor.PID", 1) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringClient(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClientMessageReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: ClientMessageReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientMessageReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pid == nil {
				m.Pid = &actor.PID{}
			}
			if err := m.Pid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func (m *ClientMessageRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: ClientMessageRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientMessageRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pid == nil {
				m.Pid = &actor.PID{}
			}
			if err := m.Pid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func skipClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
				return 0, ErrInvalidLengthClient
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClient
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
				next, err := skipClient(dAtA[start:])
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
	ErrInvalidLengthClient = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClient   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("client.proto", fileDescriptorClient) }

var fileDescriptorClient = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x48, 0x4a, 0xcf, 0x97, 0x32,
	0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x2c, 0xae, 0xcc, 0xcb,
	0x2e, 0xca, 0xcf, 0xf3, 0x0c, 0xd1, 0x07, 0x2b, 0x49, 0x4c, 0x2e, 0xc9, 0x2f, 0xd2, 0x4d, 0xcf,
	0xd7, 0x07, 0x33, 0x20, 0x62, 0xc5, 0x10, 0xdd, 0x4a, 0x2e, 0x5c, 0x02, 0xce, 0x60, 0xd3, 0x7c,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x83, 0x52, 0x0b, 0x85, 0x64, 0xb8, 0x98, 0x03, 0x32, 0x53,
	0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xf4, 0xc0, 0xba, 0xf4, 0x02, 0x3c, 0x5d, 0x82,
	0x40, 0xc2, 0x42, 0x42, 0x5c, 0x2c, 0x4e, 0xf9, 0x29, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x60, 0x36, 0xa6, 0x29, 0xc5, 0x05, 0xa4, 0x9b, 0xe2, 0xa4, 0x73, 0xe1, 0xa1, 0x1c, 0xc3,
	0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e,
	0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48,
	0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x1e, 0x30, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x40, 0xbf, 0x5f, 0x7b, 0x0e, 0x01, 0x00, 0x00,
}