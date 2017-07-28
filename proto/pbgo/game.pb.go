// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: game.proto

/*
	Package pbgo is a generated protocol buffer package.

	It is generated from these files:
		game.proto

	It has these top-level messages:
		Connect
		Connected
		SayRequest
		SayResponse
		NickRequest
		NickResponse
*/
package pbgo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import actor "github.com/AsynkronIT/protoactor-go/actor"

import strings "strings"
import reflect "reflect"

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

// 玩家连接
type Connect struct {
	Sender *actor.PID `protobuf:"bytes,1,opt,name=Sender" json:"Sender,omitempty"`
}

func (m *Connect) Reset()                    { *m = Connect{} }
func (*Connect) ProtoMessage()               {}
func (*Connect) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{0} }

func (m *Connect) GetSender() *actor.PID {
	if m != nil {
		return m.Sender
	}
	return nil
}

type Connected struct {
	Message string     `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Server  *actor.PID `protobuf:"bytes,2,opt,name=Server" json:"Server,omitempty"`
}

func (m *Connected) Reset()                    { *m = Connected{} }
func (*Connected) ProtoMessage()               {}
func (*Connected) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{1} }

func (m *Connected) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Connected) GetServer() *actor.PID {
	if m != nil {
		return m.Server
	}
	return nil
}

type SayRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (m *SayRequest) Reset()                    { *m = SayRequest{} }
func (*SayRequest) ProtoMessage()               {}
func (*SayRequest) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{2} }

func (m *SayRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SayRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SayResponse struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (m *SayResponse) Reset()                    { *m = SayResponse{} }
func (*SayResponse) ProtoMessage()               {}
func (*SayResponse) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{3} }

func (m *SayResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SayResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type NickRequest struct {
	OldUserName string `protobuf:"bytes,1,opt,name=OldUserName,proto3" json:"OldUserName,omitempty"`
	NewUserName string `protobuf:"bytes,2,opt,name=NewUserName,proto3" json:"NewUserName,omitempty"`
}

func (m *NickRequest) Reset()                    { *m = NickRequest{} }
func (*NickRequest) ProtoMessage()               {}
func (*NickRequest) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{4} }

func (m *NickRequest) GetOldUserName() string {
	if m != nil {
		return m.OldUserName
	}
	return ""
}

func (m *NickRequest) GetNewUserName() string {
	if m != nil {
		return m.NewUserName
	}
	return ""
}

type NickResponse struct {
	OldUserName string `protobuf:"bytes,1,opt,name=OldUserName,proto3" json:"OldUserName,omitempty"`
	NewUserName string `protobuf:"bytes,2,opt,name=NewUserName,proto3" json:"NewUserName,omitempty"`
}

func (m *NickResponse) Reset()                    { *m = NickResponse{} }
func (*NickResponse) ProtoMessage()               {}
func (*NickResponse) Descriptor() ([]byte, []int) { return fileDescriptorGame, []int{5} }

func (m *NickResponse) GetOldUserName() string {
	if m != nil {
		return m.OldUserName
	}
	return ""
}

func (m *NickResponse) GetNewUserName() string {
	if m != nil {
		return m.NewUserName
	}
	return ""
}

func init() {
	proto.RegisterType((*Connect)(nil), "pbgo.Connect")
	proto.RegisterType((*Connected)(nil), "pbgo.Connected")
	proto.RegisterType((*SayRequest)(nil), "pbgo.SayRequest")
	proto.RegisterType((*SayResponse)(nil), "pbgo.SayResponse")
	proto.RegisterType((*NickRequest)(nil), "pbgo.NickRequest")
	proto.RegisterType((*NickResponse)(nil), "pbgo.NickResponse")
}
func (this *Connect) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Connect)
	if !ok {
		that2, ok := that.(Connect)
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
	if !this.Sender.Equal(that1.Sender) {
		return false
	}
	return true
}
func (this *Connected) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Connected)
	if !ok {
		that2, ok := that.(Connected)
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
	if this.Message != that1.Message {
		return false
	}
	if !this.Server.Equal(that1.Server) {
		return false
	}
	return true
}
func (this *SayRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SayRequest)
	if !ok {
		that2, ok := that.(SayRequest)
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
	if this.UserName != that1.UserName {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *SayResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SayResponse)
	if !ok {
		that2, ok := that.(SayResponse)
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
	if this.UserName != that1.UserName {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *NickRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NickRequest)
	if !ok {
		that2, ok := that.(NickRequest)
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
	if this.OldUserName != that1.OldUserName {
		return false
	}
	if this.NewUserName != that1.NewUserName {
		return false
	}
	return true
}
func (this *NickResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NickResponse)
	if !ok {
		that2, ok := that.(NickResponse)
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
	if this.OldUserName != that1.OldUserName {
		return false
	}
	if this.NewUserName != that1.NewUserName {
		return false
	}
	return true
}
func (this *Connect) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pbgo.Connect{")
	if this.Sender != nil {
		s = append(s, "Sender: "+fmt.Sprintf("%#v", this.Sender)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Connected) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pbgo.Connected{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	if this.Server != nil {
		s = append(s, "Server: "+fmt.Sprintf("%#v", this.Server)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SayRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pbgo.SayRequest{")
	s = append(s, "UserName: "+fmt.Sprintf("%#v", this.UserName)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SayResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pbgo.SayResponse{")
	s = append(s, "UserName: "+fmt.Sprintf("%#v", this.UserName)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NickRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pbgo.NickRequest{")
	s = append(s, "OldUserName: "+fmt.Sprintf("%#v", this.OldUserName)+",\n")
	s = append(s, "NewUserName: "+fmt.Sprintf("%#v", this.NewUserName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NickResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pbgo.NickResponse{")
	s = append(s, "OldUserName: "+fmt.Sprintf("%#v", this.OldUserName)+",\n")
	s = append(s, "NewUserName: "+fmt.Sprintf("%#v", this.NewUserName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGame(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Connect) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Connect) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sender != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.Sender.Size()))
		n1, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Connected) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Connected) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Server != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGame(dAtA, i, uint64(m.Server.Size()))
		n2, err := m.Server.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *SayRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.UserName)))
		i += copy(dAtA[i:], m.UserName)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *SayResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.UserName)))
		i += copy(dAtA[i:], m.UserName)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func (m *NickRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NickRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.OldUserName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.OldUserName)))
		i += copy(dAtA[i:], m.OldUserName)
	}
	if len(m.NewUserName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.NewUserName)))
		i += copy(dAtA[i:], m.NewUserName)
	}
	return i, nil
}

func (m *NickResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NickResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.OldUserName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.OldUserName)))
		i += copy(dAtA[i:], m.OldUserName)
	}
	if len(m.NewUserName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGame(dAtA, i, uint64(len(m.NewUserName)))
		i += copy(dAtA[i:], m.NewUserName)
	}
	return i, nil
}

func encodeFixed64Game(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Game(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGame(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Connect) Size() (n int) {
	var l int
	_ = l
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *Connected) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	if m.Server != nil {
		l = m.Server.Size()
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *SayRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *SayResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *NickRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.OldUserName)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.NewUserName)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *NickResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.OldUserName)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.NewUserName)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func sovGame(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGame(x uint64) (n int) {
	return sovGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Connect) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Connect{`,
		`Sender:` + strings.Replace(fmt.Sprintf("%v", this.Sender), "PID", "actor.PID", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Connected) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Connected{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`Server:` + strings.Replace(fmt.Sprintf("%v", this.Server), "PID", "actor.PID", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SayRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SayRequest{`,
		`UserName:` + fmt.Sprintf("%v", this.UserName) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SayResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SayResponse{`,
		`UserName:` + fmt.Sprintf("%v", this.UserName) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NickRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NickRequest{`,
		`OldUserName:` + fmt.Sprintf("%v", this.OldUserName) + `,`,
		`NewUserName:` + fmt.Sprintf("%v", this.NewUserName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NickResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NickResponse{`,
		`OldUserName:` + fmt.Sprintf("%v", this.OldUserName) + `,`,
		`NewUserName:` + fmt.Sprintf("%v", this.NewUserName) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGame(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Connect) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: Connect: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Connect: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &actor.PID{}
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *Connected) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: Connected: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Connected: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Server == nil {
				m.Server = &actor.PID{}
			}
			if err := m.Server.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *SayRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: SayRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *SayResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: SayResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *NickRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: NickRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NickRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldUserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldUserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewUserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewUserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func (m *NickResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
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
			return fmt.Errorf("proto: NickResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NickResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldUserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldUserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewUserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
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
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewUserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
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
func skipGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
					return 0, ErrIntOverflowGame
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
				return 0, ErrInvalidLengthGame
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGame
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
				next, err := skipGame(dAtA[start:])
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
	ErrInvalidLengthGame = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGame   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("game.proto", fileDescriptorGame) }

var fileDescriptorGame = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x48, 0x4a, 0xcf, 0x97, 0x32, 0x4b, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x2c, 0xae, 0xcc, 0xcb, 0x2e, 0xca,
	0xcf, 0xf3, 0x0c, 0xd1, 0x07, 0x2b, 0x49, 0x4c, 0x2e, 0xc9, 0x2f, 0xd2, 0x4d, 0xcf, 0xd7, 0x07,
	0x33, 0x20, 0x62, 0xc5, 0x10, 0xdd, 0x4a, 0xba, 0x5c, 0xec, 0xce, 0xf9, 0x79, 0x79, 0xa9, 0xc9,
	0x25, 0x42, 0x4a, 0x5c, 0x6c, 0xc1, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x5c, 0x7a, 0x60, 0xf5, 0x7a, 0x01, 0x9e, 0x2e, 0x41, 0x50, 0x19, 0x25, 0x4f, 0x2e,
	0x4e, 0xa8, 0xf2, 0xd4, 0x14, 0x21, 0x09, 0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54,
	0xb0, 0x0e, 0xce, 0x20, 0x18, 0x17, 0x62, 0x54, 0x51, 0x59, 0x6a, 0x91, 0x04, 0x13, 0x36, 0xa3,
	0x40, 0x32, 0x4a, 0x4e, 0x5c, 0x5c, 0xc1, 0x89, 0x95, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x52, 0x5c, 0x1c, 0xa1, 0xc5, 0xa9, 0x45, 0x7e, 0x89, 0xb9, 0x30, 0xc3, 0xe0, 0x7c, 0x64,
	0x7b, 0x98, 0x50, 0xec, 0x51, 0x72, 0xe6, 0xe2, 0x06, 0x9b, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x4a, 0xa6, 0x21, 0x81, 0x5c, 0xdc, 0x7e, 0x99, 0xc9, 0xd9, 0x30, 0x97, 0x28, 0x70, 0x71, 0xfb,
	0xe7, 0xa4, 0xa0, 0x99, 0x83, 0x2c, 0x04, 0x52, 0xe1, 0x97, 0x5a, 0x0e, 0x57, 0x01, 0x31, 0x0e,
	0x59, 0x48, 0x29, 0x88, 0x8b, 0x07, 0x62, 0x24, 0xd4, 0x61, 0x54, 0x30, 0xd3, 0x49, 0xe7, 0xc2,
	0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8,
	0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0xf8, 0xe2, 0x91, 0x1c, 0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81,
	0xa3, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x74, 0x60, 0xd6, 0x6b, 0x2a, 0x02, 0x00, 0x00,
}