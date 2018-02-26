// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rtapi/realtime.proto

/*
Package rtapi is a generated protocol buffer package.

It is generated from these files:
	rtapi/realtime.proto

It has these top-level messages:
	Envelope
	Error
	Notifications
	Stream
	StreamData
	StreamPresence
	StreamPresenceEvent
*/
package rtapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import nakama_api "github.com/heroiclabs/nakama/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The selection of possible error codes.
type Error_Code int32

const (
	// An unexpected result from the server.
	Error_RUNTIME_EXCEPTION Error_Code = 0
	// The server received a message which is not recognised.
	Error_UNRECOGNIZED_PAYLOAD Error_Code = 1
	// A message was expected but contains no content.
	Error_MISSING_PAYLOAD Error_Code = 2
	// Fields in the message have an invalid format.
	Error_BAD_INPUT Error_Code = 3
	// The match id was not found.
	Error_MATCH_NOT_FOUND Error_Code = 4
	// The runtime function does not exist on the server.
	Error_RUNTIME_FUNCTION_NOT_FOUND Error_Code = 5
	// The runtime function executed with an error.
	Error_RUNTIME_FUNCTION_EXCEPTION Error_Code = 6
)

var Error_Code_name = map[int32]string{
	0: "RUNTIME_EXCEPTION",
	1: "UNRECOGNIZED_PAYLOAD",
	2: "MISSING_PAYLOAD",
	3: "BAD_INPUT",
	4: "MATCH_NOT_FOUND",
	5: "RUNTIME_FUNCTION_NOT_FOUND",
	6: "RUNTIME_FUNCTION_EXCEPTION",
}
var Error_Code_value = map[string]int32{
	"RUNTIME_EXCEPTION":          0,
	"UNRECOGNIZED_PAYLOAD":       1,
	"MISSING_PAYLOAD":            2,
	"BAD_INPUT":                  3,
	"MATCH_NOT_FOUND":            4,
	"RUNTIME_FUNCTION_NOT_FOUND": 5,
	"RUNTIME_FUNCTION_EXCEPTION": 6,
}

func (x Error_Code) String() string {
	return proto.EnumName(Error_Code_name, int32(x))
}
func (Error_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// An envelope for a realtime message.
type Envelope struct {
	Cid string `protobuf:"bytes,1,opt,name=cid" json:"cid,omitempty"`
	// Types that are valid to be assigned to Message:
	//	*Envelope_Error
	//	*Envelope_Rpc
	//	*Envelope_StreamData
	//	*Envelope_StreamPresenceEvent
	//	*Envelope_Notifications
	Message isEnvelope_Message `protobuf_oneof:"message"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isEnvelope_Message interface {
	isEnvelope_Message()
}

type Envelope_Error struct {
	Error *Error `protobuf:"bytes,2,opt,name=error,oneof"`
}
type Envelope_Rpc struct {
	Rpc *nakama_api.Rpc `protobuf:"bytes,3,opt,name=rpc,oneof"`
}
type Envelope_StreamData struct {
	StreamData *StreamData `protobuf:"bytes,4,opt,name=stream_data,json=streamData,oneof"`
}
type Envelope_StreamPresenceEvent struct {
	StreamPresenceEvent *StreamPresenceEvent `protobuf:"bytes,5,opt,name=stream_presence_event,json=streamPresenceEvent,oneof"`
}
type Envelope_Notifications struct {
	Notifications *Notifications `protobuf:"bytes,6,opt,name=notifications,oneof"`
}

func (*Envelope_Error) isEnvelope_Message()               {}
func (*Envelope_Rpc) isEnvelope_Message()                 {}
func (*Envelope_StreamData) isEnvelope_Message()          {}
func (*Envelope_StreamPresenceEvent) isEnvelope_Message() {}
func (*Envelope_Notifications) isEnvelope_Message()       {}

func (m *Envelope) GetMessage() isEnvelope_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *Envelope) GetError() *Error {
	if x, ok := m.GetMessage().(*Envelope_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Envelope) GetRpc() *nakama_api.Rpc {
	if x, ok := m.GetMessage().(*Envelope_Rpc); ok {
		return x.Rpc
	}
	return nil
}

func (m *Envelope) GetStreamData() *StreamData {
	if x, ok := m.GetMessage().(*Envelope_StreamData); ok {
		return x.StreamData
	}
	return nil
}

func (m *Envelope) GetStreamPresenceEvent() *StreamPresenceEvent {
	if x, ok := m.GetMessage().(*Envelope_StreamPresenceEvent); ok {
		return x.StreamPresenceEvent
	}
	return nil
}

func (m *Envelope) GetNotifications() *Notifications {
	if x, ok := m.GetMessage().(*Envelope_Notifications); ok {
		return x.Notifications
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Envelope) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Envelope_OneofMarshaler, _Envelope_OneofUnmarshaler, _Envelope_OneofSizer, []interface{}{
		(*Envelope_Error)(nil),
		(*Envelope_Rpc)(nil),
		(*Envelope_StreamData)(nil),
		(*Envelope_StreamPresenceEvent)(nil),
		(*Envelope_Notifications)(nil),
	}
}

func _Envelope_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Envelope_Rpc:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rpc); err != nil {
			return err
		}
	case *Envelope_StreamData:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StreamData); err != nil {
			return err
		}
	case *Envelope_StreamPresenceEvent:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StreamPresenceEvent); err != nil {
			return err
		}
	case *Envelope_Notifications:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Notifications); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Envelope.Message has unexpected type %T", x)
	}
	return nil
}

func _Envelope_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Envelope)
	switch tag {
	case 2: // message.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Error{msg}
		return true, err
	case 3: // message.rpc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(nakama_api.Rpc)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Rpc{msg}
		return true, err
	case 4: // message.stream_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamData)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_StreamData{msg}
		return true, err
	case 5: // message.stream_presence_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamPresenceEvent)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_StreamPresenceEvent{msg}
		return true, err
	case 6: // message.notifications
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Notifications)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Notifications{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Envelope_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Rpc:
		s := proto.Size(x.Rpc)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_StreamData:
		s := proto.Size(x.StreamData)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_StreamPresenceEvent:
		s := proto.Size(x.StreamPresenceEvent)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Notifications:
		s := proto.Size(x.Notifications)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A logical error which may occur on the server.
type Error struct {
	// The error code which should be one of "Error.Code" enums.
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	// A message in English to help developers debug the response.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// Additional error details which may be different for each response.
	Context map[string]string `protobuf:"bytes,3,rep,name=context" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

// A collection of zero or more notifications.
type Notifications struct {
	// Collection of notifications.
	Notifications []*nakama_api.Notification `protobuf:"bytes,1,rep,name=notifications" json:"notifications,omitempty"`
}

func (m *Notifications) Reset()                    { *m = Notifications{} }
func (m *Notifications) String() string            { return proto.CompactTextString(m) }
func (*Notifications) ProtoMessage()               {}
func (*Notifications) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Notifications) GetNotifications() []*nakama_api.Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

// Represents identifying information for a stream.
type Stream struct {
	// Mode identifies the type of stream.
	Mode int32 `protobuf:"varint,1,opt,name=mode" json:"mode,omitempty"`
	// Subject is the primary identifier, if any.
	Subject string `protobuf:"bytes,2,opt,name=subject" json:"subject,omitempty"`
	// Descriptor is a secondary identifier, if any.
	Descriptor_ string `protobuf:"bytes,3,opt,name=descriptor" json:"descriptor,omitempty"`
	// The label is an arbitrary identifying string, if the stream has one.
	Label string `protobuf:"bytes,4,opt,name=label" json:"label,omitempty"`
}

func (m *Stream) Reset()                    { *m = Stream{} }
func (m *Stream) String() string            { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()               {}
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Stream) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *Stream) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Stream) GetDescriptor_() string {
	if m != nil {
		return m.Descriptor_
	}
	return ""
}

func (m *Stream) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// A data message delivered over a stream.
type StreamData struct {
	// The stream this data message relates to.
	Stream *Stream `protobuf:"bytes,1,opt,name=stream" json:"stream,omitempty"`
	// The sender, if any.
	Sender *StreamPresence `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	// Arbitrary contents of the data message.
	Data string `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *StreamData) Reset()                    { *m = StreamData{} }
func (m *StreamData) String() string            { return proto.CompactTextString(m) }
func (*StreamData) ProtoMessage()               {}
func (*StreamData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StreamData) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

func (m *StreamData) GetSender() *StreamPresence {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *StreamData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// A user session associated to a stream, usually through a list operation or a join/leave event.
type StreamPresence struct {
	// The user this presence belongs to.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// A unique session ID identifying the particular connection, because the user may have many.
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	// The username for display purposes.
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	// Whether this presence generates persistent data/messages, if applicable for the stream type.
	Persistence bool `protobuf:"varint,4,opt,name=persistence" json:"persistence,omitempty"`
	// A user-set status message for this stream, if applicable.
	Status string `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *StreamPresence) Reset()                    { *m = StreamPresence{} }
func (m *StreamPresence) String() string            { return proto.CompactTextString(m) }
func (*StreamPresence) ProtoMessage()               {}
func (*StreamPresence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StreamPresence) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *StreamPresence) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *StreamPresence) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *StreamPresence) GetPersistence() bool {
	if m != nil {
		return m.Persistence
	}
	return false
}

func (m *StreamPresence) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// A set of joins and leaves on a particular stream.
type StreamPresenceEvent struct {
	// The stream this event relates to.
	Stream *Stream `protobuf:"bytes,1,opt,name=stream" json:"stream,omitempty"`
	// Presences joining the stream as part of this event, if any.
	Joins []*StreamPresence `protobuf:"bytes,2,rep,name=joins" json:"joins,omitempty"`
	// Presences leaving the stream as part of this event, if any.
	Leaves []*StreamPresence `protobuf:"bytes,3,rep,name=leaves" json:"leaves,omitempty"`
}

func (m *StreamPresenceEvent) Reset()                    { *m = StreamPresenceEvent{} }
func (m *StreamPresenceEvent) String() string            { return proto.CompactTextString(m) }
func (*StreamPresenceEvent) ProtoMessage()               {}
func (*StreamPresenceEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StreamPresenceEvent) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

func (m *StreamPresenceEvent) GetJoins() []*StreamPresence {
	if m != nil {
		return m.Joins
	}
	return nil
}

func (m *StreamPresenceEvent) GetLeaves() []*StreamPresence {
	if m != nil {
		return m.Leaves
	}
	return nil
}

func init() {
	proto.RegisterType((*Envelope)(nil), "nakama.realtime.Envelope")
	proto.RegisterType((*Error)(nil), "nakama.realtime.Error")
	proto.RegisterType((*Notifications)(nil), "nakama.realtime.Notifications")
	proto.RegisterType((*Stream)(nil), "nakama.realtime.Stream")
	proto.RegisterType((*StreamData)(nil), "nakama.realtime.StreamData")
	proto.RegisterType((*StreamPresence)(nil), "nakama.realtime.StreamPresence")
	proto.RegisterType((*StreamPresenceEvent)(nil), "nakama.realtime.StreamPresenceEvent")
	proto.RegisterEnum("nakama.realtime.Error_Code", Error_Code_name, Error_Code_value)
}

func init() { proto.RegisterFile("rtapi/realtime.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0x9d, 0xd8, 0xad, 0x4f, 0xe8, 0x36, 0x4c, 0xdb, 0x5d, 0x2b, 0x88, 0x12, 0x79, 0xb9,
	0xa8, 0xb8, 0x70, 0xa4, 0x22, 0x04, 0x5a, 0x89, 0x15, 0xf9, 0x71, 0x37, 0x16, 0xac, 0x13, 0x4d,
	0x13, 0x09, 0x7a, 0x63, 0x4d, 0x9c, 0x61, 0xeb, 0xd6, 0x7f, 0xf2, 0x4c, 0x22, 0xfa, 0x0a, 0x3c,
	0x04, 0x17, 0x5c, 0x72, 0xc5, 0x15, 0xcf, 0x83, 0xc4, 0x8b, 0xa0, 0x19, 0xdb, 0xa9, 0xfb, 0xa7,
	0x6a, 0xaf, 0x3c, 0xe7, 0x9c, 0xef, 0x9c, 0x33, 0xfe, 0xbe, 0x73, 0x06, 0x0e, 0x73, 0x4e, 0xb2,
	0xb0, 0x9f, 0x53, 0x12, 0xf1, 0x30, 0xa6, 0x76, 0x96, 0xa7, 0x3c, 0x45, 0xfb, 0x09, 0xb9, 0x26,
	0x31, 0xb1, 0x2b, 0x77, 0xf7, 0xab, 0x0f, 0x21, 0xbf, 0x5c, 0x2f, 0xed, 0x20, 0x8d, 0xfb, 0x97,
	0x34, 0x4f, 0xc3, 0x20, 0x22, 0x4b, 0xd6, 0x2f, 0x60, 0x7d, 0x51, 0x81, 0x64, 0x61, 0x91, 0x6c,
	0xfd, 0xab, 0xc2, 0xae, 0x93, 0x6c, 0x68, 0x94, 0x66, 0x14, 0x75, 0xa0, 0x19, 0x84, 0x2b, 0x53,
	0xe9, 0x29, 0x27, 0x06, 0x16, 0x47, 0x64, 0x83, 0x46, 0xf3, 0x3c, 0xcd, 0x4d, 0xb5, 0xa7, 0x9c,
	0xb4, 0x4f, 0x5f, 0xda, 0xf7, 0x7a, 0xd9, 0x8e, 0x88, 0x4e, 0x1a, 0xb8, 0x80, 0xa1, 0xd7, 0xd0,
	0xcc, 0xb3, 0xc0, 0x6c, 0x4a, 0xf4, 0x7e, 0x85, 0x16, 0xed, 0x70, 0x16, 0x4c, 0x1a, 0x58, 0x44,
	0xd1, 0x5b, 0x68, 0x33, 0x9e, 0x53, 0x12, 0xfb, 0x2b, 0xc2, 0x89, 0xd9, 0x92, 0xe0, 0xcf, 0x1e,
	0x94, 0x3e, 0x97, 0x98, 0x31, 0xe1, 0x64, 0xd2, 0xc0, 0xc0, 0xb6, 0x16, 0xba, 0x80, 0xa3, 0x32,
	0x3f, 0xcb, 0x29, 0xa3, 0x49, 0x40, 0x7d, 0xba, 0xa1, 0x09, 0x37, 0x35, 0x59, 0xe9, 0xcb, 0x27,
	0x2a, 0xcd, 0x4a, 0xb0, 0x23, 0xb0, 0x93, 0x06, 0x3e, 0x60, 0x0f, 0xdd, 0xe8, 0x0c, 0xf6, 0x92,
	0x94, 0x87, 0xbf, 0x86, 0x01, 0xe1, 0x61, 0x9a, 0x30, 0x53, 0x97, 0x35, 0x8f, 0x1f, 0xd4, 0xf4,
	0xea, 0xa8, 0x49, 0x03, 0xdf, 0x4d, 0x1b, 0x1a, 0xb0, 0x13, 0x53, 0xc6, 0xc8, 0x07, 0x6a, 0xfd,
	0xa7, 0x82, 0x26, 0x69, 0x42, 0x08, 0x5a, 0x41, 0xba, 0xa2, 0x92, 0x60, 0x0d, 0xcb, 0x33, 0x32,
	0xb7, 0x40, 0xc9, 0xb1, 0x81, 0x2b, 0x13, 0x7d, 0x0f, 0x3b, 0x41, 0x9a, 0x70, 0xfa, 0x1b, 0x37,
	0x9b, 0xbd, 0xe6, 0x49, 0xfb, 0xf4, 0xf5, 0xe3, 0xec, 0xdb, 0xa3, 0x02, 0xe5, 0x24, 0x3c, 0xbf,
	0xc1, 0x55, 0x4e, 0xf7, 0x0d, 0x7c, 0x52, 0x0f, 0x08, 0x71, 0xaf, 0xe9, 0x4d, 0x25, 0xee, 0x35,
	0xbd, 0x41, 0x87, 0xa0, 0x6d, 0x48, 0xb4, 0xae, 0x1a, 0x17, 0xc6, 0x1b, 0xf5, 0x3b, 0xc5, 0xfa,
	0x5b, 0x81, 0xd6, 0x48, 0xdc, 0xee, 0x08, 0x3e, 0xc5, 0x0b, 0x6f, 0xee, 0xbe, 0x77, 0x7c, 0xe7,
	0xe7, 0x91, 0x33, 0x9b, 0xbb, 0x53, 0xaf, 0xd3, 0x40, 0x26, 0x1c, 0x2e, 0x3c, 0xec, 0x8c, 0xa6,
	0xef, 0x3c, 0xf7, 0xc2, 0x19, 0xfb, 0xb3, 0xc1, 0x2f, 0x3f, 0x4d, 0x07, 0xe3, 0x8e, 0x82, 0x0e,
	0x60, 0xff, 0xbd, 0x7b, 0x7e, 0xee, 0x7a, 0xef, 0xb6, 0x4e, 0x15, 0xed, 0x81, 0x31, 0x1c, 0x8c,
	0x7d, 0xd7, 0x9b, 0x2d, 0xe6, 0x9d, 0xa6, 0xc4, 0x0c, 0xe6, 0xa3, 0x89, 0xef, 0x4d, 0xe7, 0xfe,
	0xd9, 0x74, 0xe1, 0x8d, 0x3b, 0x2d, 0x74, 0x0c, 0xdd, 0xaa, 0xd3, 0xd9, 0xc2, 0x1b, 0x89, 0x46,
	0xb5, 0xb8, 0xf6, 0x68, 0xfc, 0xf6, 0x4a, 0xba, 0x35, 0x85, 0xbd, 0x3b, 0x92, 0xa0, 0xb7, 0xf7,
	0x95, 0x54, 0x24, 0x89, 0x66, 0x7d, 0x28, 0xeb, 0x19, 0xf7, 0x14, 0xb4, 0x22, 0xd0, 0x8b, 0xb9,
	0x11, 0xb2, 0xc5, 0x35, 0xd9, 0xe2, 0x52, 0x36, 0xb6, 0x5e, 0x5e, 0xd1, 0x80, 0x57, 0xb2, 0x95,
	0x26, 0x3a, 0x06, 0x58, 0x51, 0x16, 0xe4, 0x61, 0xc6, 0xd3, 0x5c, 0x6e, 0x82, 0x81, 0x6b, 0x1e,
	0xc1, 0x7a, 0x44, 0x96, 0x34, 0x92, 0x73, 0x6f, 0xe0, 0xc2, 0xb0, 0x7e, 0x57, 0x00, 0x6e, 0x07,
	0x1e, 0xf5, 0x41, 0x2f, 0xa6, 0x53, 0x36, 0x6d, 0x9f, 0xbe, 0x7a, 0x62, 0xa6, 0x71, 0x09, 0x43,
	0xdf, 0x82, 0xce, 0x68, 0xb2, 0xa2, 0xd5, 0xa6, 0x7e, 0xf1, 0xcc, 0x12, 0xe0, 0x12, 0x2e, 0x7e,
	0x4e, 0x6e, 0x61, 0x71, 0x51, 0x79, 0xb6, 0xfe, 0x50, 0xe0, 0xc5, 0x5d, 0x38, 0x7a, 0x05, 0x3b,
	0x6b, 0x46, 0x73, 0x7f, 0xfb, 0x3c, 0xe8, 0xc2, 0x74, 0x57, 0xe8, 0x73, 0x00, 0x46, 0x19, 0x0b,
	0xd3, 0x44, 0xc4, 0x0a, 0x2e, 0x8c, 0xd2, 0xe3, 0xae, 0x50, 0x17, 0x76, 0x05, 0x30, 0x21, 0x31,
	0x2d, 0x5b, 0x6c, 0x6d, 0xd4, 0x83, 0x76, 0x46, 0x73, 0x16, 0x32, 0x2e, 0x5a, 0x48, 0x3e, 0x76,
	0x71, 0xdd, 0x85, 0x5e, 0x0a, 0x1a, 0x08, 0x5f, 0x33, 0xb9, 0xda, 0x06, 0x2e, 0x2d, 0xeb, 0x1f,
	0x05, 0x0e, 0x1e, 0x59, 0xea, 0x8f, 0xa7, 0xed, 0x1b, 0xd0, 0xae, 0xd2, 0x30, 0x61, 0xa6, 0x2a,
	0x87, 0xe3, 0x59, 0xd6, 0x0a, 0xb4, 0x60, 0x3b, 0xa2, 0x64, 0x43, 0x59, 0xb9, 0x99, 0xcf, 0xb3,
	0x5d, 0xc0, 0x87, 0x3f, 0xc0, 0x51, 0x90, 0xc6, 0xf6, 0xed, 0xab, 0x5c, 0x26, 0x0e, 0x5f, 0x78,
	0xf2, 0x8b, 0xcb, 0xfc, 0x99, 0x72, 0xa1, 0xc9, 0xc7, 0xfe, 0x4f, 0xb5, 0xe5, 0xfd, 0x38, 0x1b,
	0xfe, 0xa5, 0xea, 0x05, 0x60, 0xa9, 0xcb, 0x77, 0xfb, 0xeb, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xc9, 0x14, 0x69, 0x59, 0x0c, 0x06, 0x00, 0x00,
}