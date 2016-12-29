// Code generated by protoc-gen-gogo.
// source: events_payload.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EventsPayload struct {
	Events []*EventsPayload_Event `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *EventsPayload) Reset()                    { *m = EventsPayload{} }
func (m *EventsPayload) String() string            { return proto.CompactTextString(m) }
func (*EventsPayload) ProtoMessage()               {}
func (*EventsPayload) Descriptor() ([]byte, []int) { return fileDescriptorEventsPayload, []int{0} }

func (m *EventsPayload) GetEvents() []*EventsPayload_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type EventsPayload_Event struct {
	Title          string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text           string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Ts             int64    `protobuf:"varint,3,opt,name=ts,proto3" json:"ts,omitempty"`
	Priority       string   `protobuf:"bytes,4,opt,name=priority,proto3" json:"priority,omitempty"`
	Host           string   `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Tags           []string `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	AlertType      string   `protobuf:"bytes,7,opt,name=alert_type,json=alertType,proto3" json:"alert_type,omitempty"`
	AggregationKey string   `protobuf:"bytes,8,opt,name=aggregation_key,json=aggregationKey,proto3" json:"aggregation_key,omitempty"`
	SourceTypeName string   `protobuf:"bytes,9,opt,name=source_type_name,json=sourceTypeName,proto3" json:"source_type_name,omitempty"`
}

func (m *EventsPayload_Event) Reset()         { *m = EventsPayload_Event{} }
func (m *EventsPayload_Event) String() string { return proto.CompactTextString(m) }
func (*EventsPayload_Event) ProtoMessage()    {}
func (*EventsPayload_Event) Descriptor() ([]byte, []int) {
	return fileDescriptorEventsPayload, []int{0, 0}
}

func (m *EventsPayload_Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EventsPayload_Event) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *EventsPayload_Event) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *EventsPayload_Event) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

func (m *EventsPayload_Event) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *EventsPayload_Event) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *EventsPayload_Event) GetAlertType() string {
	if m != nil {
		return m.AlertType
	}
	return ""
}

func (m *EventsPayload_Event) GetAggregationKey() string {
	if m != nil {
		return m.AggregationKey
	}
	return ""
}

func (m *EventsPayload_Event) GetSourceTypeName() string {
	if m != nil {
		return m.SourceTypeName
	}
	return ""
}

func init() {
	proto.RegisterType((*EventsPayload)(nil), "pb.EventsPayload")
	proto.RegisterType((*EventsPayload_Event)(nil), "pb.EventsPayload.Event")
}
func (m *EventsPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventsPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0xa
			i++
			i = encodeVarintEventsPayload(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *EventsPayload_Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventsPayload_Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Title) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEventsPayload(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Text) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEventsPayload(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	if m.Ts != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEventsPayload(dAtA, i, uint64(m.Ts))
	}
	if len(m.Priority) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEventsPayload(dAtA, i, uint64(len(m.Priority)))
		i += copy(dAtA[i:], m.Priority)
	}
	if len(m.Host) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEventsPayload(dAtA, i, uint64(len(m.Host)))
		i += copy(dAtA[i:], m.Host)
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.AlertType) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintEventsPayload(dAtA, i, uint64(len(m.AlertType)))
		i += copy(dAtA[i:], m.AlertType)
	}
	if len(m.AggregationKey) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintEventsPayload(dAtA, i, uint64(len(m.AggregationKey)))
		i += copy(dAtA[i:], m.AggregationKey)
	}
	if len(m.SourceTypeName) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintEventsPayload(dAtA, i, uint64(len(m.SourceTypeName)))
		i += copy(dAtA[i:], m.SourceTypeName)
	}
	return i, nil
}

func encodeFixed64EventsPayload(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32EventsPayload(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEventsPayload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EventsPayload) Size() (n int) {
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovEventsPayload(uint64(l))
		}
	}
	return n
}

func (m *EventsPayload_Event) Size() (n int) {
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovEventsPayload(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovEventsPayload(uint64(l))
	}
	if m.Ts != 0 {
		n += 1 + sovEventsPayload(uint64(m.Ts))
	}
	l = len(m.Priority)
	if l > 0 {
		n += 1 + l + sovEventsPayload(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovEventsPayload(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovEventsPayload(uint64(l))
		}
	}
	l = len(m.AlertType)
	if l > 0 {
		n += 1 + l + sovEventsPayload(uint64(l))
	}
	l = len(m.AggregationKey)
	if l > 0 {
		n += 1 + l + sovEventsPayload(uint64(l))
	}
	l = len(m.SourceTypeName)
	if l > 0 {
		n += 1 + l + sovEventsPayload(uint64(l))
	}
	return n
}

func sovEventsPayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEventsPayload(x uint64) (n int) {
	return sovEventsPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventsPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventsPayload
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
			return fmt.Errorf("proto: EventsPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventsPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &EventsPayload_Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventsPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventsPayload
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
func (m *EventsPayload_Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventsPayload
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ts", wireType)
			}
			m.Ts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ts |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Priority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlertType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AlertType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AggregationKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceTypeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventsPayload
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
				return ErrInvalidLengthEventsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceTypeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventsPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventsPayload
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
func skipEventsPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEventsPayload
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
					return 0, ErrIntOverflowEventsPayload
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
					return 0, ErrIntOverflowEventsPayload
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
				return 0, ErrInvalidLengthEventsPayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEventsPayload
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
				next, err := skipEventsPayload(dAtA[start:])
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
	ErrInvalidLengthEventsPayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEventsPayload   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("events_payload.proto", fileDescriptorEventsPayload) }

var fileDescriptorEventsPayload = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xbf, 0x4d, 0xda, 0x7e, 0xed, 0x88, 0xb5, 0x2c, 0x05, 0x97, 0x82, 0x21, 0x78, 0x31,
	0xa7, 0x0a, 0xfa, 0x06, 0x82, 0x27, 0x41, 0xa4, 0x78, 0x0f, 0x1b, 0x1d, 0x62, 0x30, 0xcd, 0x2e,
	0xbb, 0xa3, 0xb8, 0x6f, 0xe2, 0xb3, 0xf8, 0x04, 0x1e, 0x7d, 0x04, 0x89, 0x0f, 0xe1, 0x55, 0x32,
	0x2b, 0xa2, 0xb7, 0xf9, 0xff, 0xe6, 0x37, 0xff, 0xc3, 0xc0, 0x12, 0x1f, 0xb1, 0x23, 0x5f, 0x5a,
	0x1d, 0x5a, 0xa3, 0x6f, 0xd7, 0xd6, 0x19, 0x32, 0x32, 0xb1, 0xd5, 0xe1, 0x4b, 0x02, 0xbb, 0xe7,
	0xbc, 0xbc, 0x8a, 0x3b, 0x79, 0x0c, 0x93, 0x68, 0x2b, 0x91, 0xa7, 0xc5, 0xce, 0xc9, 0xfe, 0xda,
	0x56, 0xeb, 0x3f, 0x4a, 0x4c, 0x9b, 0x6f, 0x6d, 0xf5, 0x29, 0x60, 0xcc, 0x44, 0x2e, 0x61, 0x4c,
	0x0d, 0xb5, 0xa8, 0x44, 0x2e, 0x8a, 0xd9, 0x26, 0x06, 0x29, 0x61, 0x44, 0xf8, 0x44, 0x2a, 0x61,
	0xc8, 0xb3, 0x9c, 0x43, 0x42, 0x5e, 0xa5, 0xb9, 0x28, 0xd2, 0x4d, 0x42, 0x5e, 0xae, 0x60, 0x6a,
	0x5d, 0x63, 0x5c, 0x43, 0x41, 0x8d, 0xd8, 0xfb, 0xc9, 0xc3, 0xfd, 0x9d, 0xf1, 0xa4, 0xc6, 0xf1,
	0x7e, 0x98, 0xb9, 0x53, 0xd7, 0x5e, 0x4d, 0xf2, 0x94, 0x3b, 0x75, 0xed, 0xe5, 0x01, 0x80, 0x6e,
	0xd1, 0x51, 0x49, 0xc1, 0xa2, 0xfa, 0xcf, 0xf6, 0x8c, 0xc9, 0x75, 0xb0, 0x28, 0x8f, 0x60, 0x4f,
	0xd7, 0xb5, 0xc3, 0x5a, 0x53, 0x63, 0xba, 0xf2, 0x1e, 0x83, 0x9a, 0xb2, 0x33, 0xff, 0x85, 0x2f,
	0x30, 0xc8, 0x02, 0x16, 0xde, 0x3c, 0xb8, 0x1b, 0xe4, 0xa2, 0xb2, 0xd3, 0x5b, 0x54, 0xb3, 0x68,
	0x46, 0x3e, 0xd4, 0x5d, 0xea, 0x2d, 0x9e, 0x2d, 0x5e, 0xfb, 0x4c, 0xbc, 0xf5, 0x99, 0x78, 0xef,
	0x33, 0xf1, 0xfc, 0x91, 0xfd, 0xab, 0x26, 0xfc, 0xd9, 0xd3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x8f, 0x3b, 0x36, 0x71, 0x01, 0x00, 0x00,
}