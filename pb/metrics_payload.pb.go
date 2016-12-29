// Code generated by protoc-gen-go.
// source: metrics_payload.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	metrics_payload.proto

It has these top-level messages:
	MetricsPayload
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MetricsPayload struct {
	Timeseries []*MetricsPayload_Timeserie `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *MetricsPayload) Reset()                    { *m = MetricsPayload{} }
func (m *MetricsPayload) String() string            { return proto.CompactTextString(m) }
func (*MetricsPayload) ProtoMessage()               {}
func (*MetricsPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetricsPayload) GetTimeseries() []*MetricsPayload_Timeserie {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

type MetricsPayload_Timeserie struct {
	Metric   string                            `protobuf:"bytes,1,opt,name=metric" json:"metric,omitempty"`
	Type     string                            `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Host     string                            `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	Points   []*MetricsPayload_Timeserie_Point `protobuf:"bytes,4,rep,name=points" json:"points,omitempty"`
	Tags     []string                          `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
	Interval int32                             `protobuf:"varint,6,opt,name=interval" json:"interval,omitempty"`
}

func (m *MetricsPayload_Timeserie) Reset()                    { *m = MetricsPayload_Timeserie{} }
func (m *MetricsPayload_Timeserie) String() string            { return proto.CompactTextString(m) }
func (*MetricsPayload_Timeserie) ProtoMessage()               {}
func (*MetricsPayload_Timeserie) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *MetricsPayload_Timeserie) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *MetricsPayload_Timeserie) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetricsPayload_Timeserie) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *MetricsPayload_Timeserie) GetPoints() []*MetricsPayload_Timeserie_Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *MetricsPayload_Timeserie) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MetricsPayload_Timeserie) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

type MetricsPayload_Timeserie_Point struct {
	Ts    int64   `protobuf:"varint,1,opt,name=ts" json:"ts,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *MetricsPayload_Timeserie_Point) Reset()         { *m = MetricsPayload_Timeserie_Point{} }
func (m *MetricsPayload_Timeserie_Point) String() string { return proto.CompactTextString(m) }
func (*MetricsPayload_Timeserie_Point) ProtoMessage()    {}
func (*MetricsPayload_Timeserie_Point) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *MetricsPayload_Timeserie_Point) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *MetricsPayload_Timeserie_Point) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*MetricsPayload)(nil), "pb.MetricsPayload")
	proto.RegisterType((*MetricsPayload_Timeserie)(nil), "pb.MetricsPayload.Timeserie")
	proto.RegisterType((*MetricsPayload_Timeserie_Point)(nil), "pb.MetricsPayload.Timeserie.Point")
}

func init() { proto.RegisterFile("metrics_payload.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0x26, 0xd9, 0xee, 0x62, 0x47, 0xe8, 0x61, 0x50, 0x09, 0x8b, 0x87, 0xa5, 0xa7, 0xbd, 0x98,
	0x83, 0xde, 0xc4, 0x57, 0x10, 0x4a, 0xf0, 0x2e, 0x59, 0x0d, 0x35, 0xb0, 0x6d, 0x42, 0x66, 0x2c,
	0xf4, 0xc9, 0x7c, 0x23, 0x9f, 0x43, 0x36, 0xa9, 0x8b, 0x5e, 0x7a, 0xfb, 0xfe, 0x92, 0xef, 0x63,
	0xe0, 0x7a, 0xe7, 0x38, 0xf9, 0x37, 0x7a, 0x8d, 0xf6, 0x38, 0x06, 0xfb, 0xae, 0x63, 0x0a, 0x1c,
	0x50, 0xc6, 0x61, 0xfd, 0x25, 0x61, 0xf5, 0x5c, 0xdc, 0x4d, 0x31, 0xf1, 0x09, 0x80, 0xfd, 0xce,
	0x91, 0x4b, 0xde, 0x91, 0x12, 0x5d, 0xd5, 0x5f, 0xde, 0xdf, 0xea, 0x38, 0xe8, 0xff, 0x39, 0xfd,
	0xf2, 0x1b, 0x32, 0x7f, 0xf2, 0xed, 0xb7, 0x80, 0xe5, 0xec, 0xe0, 0x0d, 0x34, 0xa5, 0x5b, 0x89,
	0x4e, 0xf4, 0x4b, 0x73, 0x62, 0x88, 0xb0, 0xe0, 0x63, 0x74, 0x4a, 0x66, 0x35, 0xe3, 0x49, 0xfb,
	0x08, 0xc4, 0xaa, 0x2a, 0xda, 0x84, 0xf1, 0x11, 0x9a, 0x18, 0xfc, 0x9e, 0x49, 0x2d, 0xf2, 0x8e,
	0xf5, 0xb9, 0x1d, 0x7a, 0x33, 0x45, 0xcd, 0xe9, 0x45, 0xee, 0xb0, 0x5b, 0x52, 0x75, 0x57, 0xe5,
	0x0e, 0xbb, 0x25, 0x6c, 0xe1, 0xc2, 0xef, 0xd9, 0xa5, 0x83, 0x1d, 0x55, 0xd3, 0x89, 0xbe, 0x36,
	0x33, 0x6f, 0xef, 0xa0, 0xce, 0x1f, 0xe0, 0x0a, 0x24, 0x53, 0x1e, 0x5c, 0x19, 0xc9, 0x84, 0x57,
	0x50, 0x1f, 0xec, 0xf8, 0x59, 0xd6, 0x0a, 0x53, 0xc8, 0xd0, 0xe4, 0x23, 0x3e, 0xfc, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x47, 0xda, 0x02, 0xe7, 0x5d, 0x01, 0x00, 0x00,
}