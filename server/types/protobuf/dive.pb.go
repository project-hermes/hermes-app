// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dive.proto

package protobuf // import "github.com/project-hermes/hermes-app/server/types/protobuf"

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

type Dive struct {
	SensorId             uint64     `protobuf:"varint,1,opt,name=sensorId,proto3" json:"sensorId,omitempty"`
	TimeStart            int32      `protobuf:"varint,2,opt,name=timeStart,proto3" json:"timeStart,omitempty"`
	TimeEnd              int32      `protobuf:"varint,3,opt,name=timeEnd,proto3" json:"timeEnd,omitempty"`
	DataCount            int32      `protobuf:"varint,4,opt,name=dataCount,proto3" json:"dataCount,omitempty"`
	NmeaStart            string     `protobuf:"bytes,5,opt,name=nmeaStart,proto3" json:"nmeaStart,omitempty"`
	NmeaEnd              string     `protobuf:"bytes,6,opt,name=nmeaEnd,proto3" json:"nmeaEnd,omitempty"`
	DiveData             *DataPoint `protobuf:"bytes,7,opt,name=diveData,proto3" json:"diveData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Dive) Reset()         { *m = Dive{} }
func (m *Dive) String() string { return proto.CompactTextString(m) }
func (*Dive) ProtoMessage()    {}
func (*Dive) Descriptor() ([]byte, []int) {
	return fileDescriptor_dive_821c774cb1d976ed, []int{0}
}
func (m *Dive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dive.Unmarshal(m, b)
}
func (m *Dive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dive.Marshal(b, m, deterministic)
}
func (dst *Dive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dive.Merge(dst, src)
}
func (m *Dive) XXX_Size() int {
	return xxx_messageInfo_Dive.Size(m)
}
func (m *Dive) XXX_DiscardUnknown() {
	xxx_messageInfo_Dive.DiscardUnknown(m)
}

var xxx_messageInfo_Dive proto.InternalMessageInfo

func (m *Dive) GetSensorId() uint64 {
	if m != nil {
		return m.SensorId
	}
	return 0
}

func (m *Dive) GetTimeStart() int32 {
	if m != nil {
		return m.TimeStart
	}
	return 0
}

func (m *Dive) GetTimeEnd() int32 {
	if m != nil {
		return m.TimeEnd
	}
	return 0
}

func (m *Dive) GetDataCount() int32 {
	if m != nil {
		return m.DataCount
	}
	return 0
}

func (m *Dive) GetNmeaStart() string {
	if m != nil {
		return m.NmeaStart
	}
	return ""
}

func (m *Dive) GetNmeaEnd() string {
	if m != nil {
		return m.NmeaEnd
	}
	return ""
}

func (m *Dive) GetDiveData() *DataPoint {
	if m != nil {
		return m.DiveData
	}
	return nil
}

type DataPoint struct {
	Temp1                float32  `protobuf:"fixed32,1,opt,name=temp1,proto3" json:"temp1,omitempty"`
	Temp2                float32  `protobuf:"fixed32,2,opt,name=temp2,proto3" json:"temp2,omitempty"`
	Depth                float32  `protobuf:"fixed32,3,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataPoint) Reset()         { *m = DataPoint{} }
func (m *DataPoint) String() string { return proto.CompactTextString(m) }
func (*DataPoint) ProtoMessage()    {}
func (*DataPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_dive_821c774cb1d976ed, []int{1}
}
func (m *DataPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataPoint.Unmarshal(m, b)
}
func (m *DataPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataPoint.Marshal(b, m, deterministic)
}
func (dst *DataPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPoint.Merge(dst, src)
}
func (m *DataPoint) XXX_Size() int {
	return xxx_messageInfo_DataPoint.Size(m)
}
func (m *DataPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPoint.DiscardUnknown(m)
}

var xxx_messageInfo_DataPoint proto.InternalMessageInfo

func (m *DataPoint) GetTemp1() float32 {
	if m != nil {
		return m.Temp1
	}
	return 0
}

func (m *DataPoint) GetTemp2() float32 {
	if m != nil {
		return m.Temp2
	}
	return 0
}

func (m *DataPoint) GetDepth() float32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func init() {
	proto.RegisterType((*Dive)(nil), "types.Dive")
	proto.RegisterType((*DataPoint)(nil), "types.DataPoint")
}

func init() { proto.RegisterFile("dive.proto", fileDescriptor_dive_821c774cb1d976ed) }

var fileDescriptor_dive_821c774cb1d976ed = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x31, 0x4f, 0x84, 0x30,
	0x14, 0xc7, 0xd3, 0x13, 0xee, 0x8e, 0xba, 0x98, 0xc6, 0xa1, 0x31, 0x0e, 0xe4, 0x26, 0x06, 0x0f,
	0x22, 0x6e, 0xc6, 0x49, 0xcf, 0xc1, 0xc1, 0xc4, 0xd4, 0xcd, 0xad, 0x1c, 0x4f, 0xc1, 0x84, 0xb6,
	0x69, 0x1f, 0x24, 0x7e, 0x59, 0x3f, 0x8b, 0x69, 0x2b, 0xdc, 0x04, 0xbf, 0xdf, 0xeb, 0xfb, 0xa7,
	0xff, 0x52, 0xda, 0xf6, 0x13, 0x94, 0xc6, 0x6a, 0xd4, 0x2c, 0xc5, 0x1f, 0x03, 0x6e, 0xf7, 0x4b,
	0x68, 0x72, 0xe8, 0x27, 0x60, 0x57, 0x74, 0xeb, 0x40, 0x39, 0x6d, 0x5f, 0x5a, 0x4e, 0x72, 0x52,
	0x24, 0x62, 0x61, 0x76, 0x4d, 0x33, 0xec, 0x07, 0x78, 0x47, 0x69, 0x91, 0xaf, 0x72, 0x52, 0xa4,
	0xe2, 0x24, 0x18, 0xa7, 0x1b, 0x0f, 0xcf, 0xaa, 0xe5, 0x67, 0x61, 0x36, 0xa3, 0xdf, 0x6b, 0x25,
	0xca, 0x27, 0x3d, 0x2a, 0xe4, 0x49, 0xdc, 0x5b, 0x84, 0x9f, 0xaa, 0x01, 0x64, 0x4c, 0x4d, 0x73,
	0x52, 0x64, 0xe2, 0x24, 0x7c, 0xaa, 0x07, 0x9f, 0xba, 0x0e, 0xb3, 0x19, 0xd9, 0x0d, 0xdd, 0xfa,
	0x1e, 0x07, 0x89, 0x92, 0x6f, 0x72, 0x52, 0x9c, 0xd7, 0x17, 0x65, 0x28, 0x53, 0x7a, 0xf5, 0xa6,
	0x7b, 0x85, 0x62, 0x39, 0xb1, 0x7b, 0xa5, 0xd9, 0xa2, 0xd9, 0x25, 0x4d, 0x11, 0x06, 0x73, 0x1b,
	0x1a, 0xae, 0x44, 0x84, 0xd9, 0xd6, 0xa1, 0xda, 0xbf, 0xad, 0xbd, 0x6d, 0xc1, 0x60, 0x17, 0x4a,
	0xad, 0x44, 0x84, 0xc7, 0x87, 0x8f, 0xfb, 0xaf, 0x1e, 0xbb, 0xb1, 0x29, 0x8f, 0x7a, 0xa8, 0x8c,
	0xd5, 0xdf, 0x70, 0xc4, 0x7d, 0x07, 0x76, 0x00, 0x57, 0xc5, 0xcf, 0x5e, 0x1a, 0x53, 0x39, 0xb0,
	0x13, 0xd8, 0x2a, 0xdc, 0xab, 0x0a, 0x2f, 0xde, 0x8c, 0x9f, 0xcd, 0x3a, 0xfc, 0xdd, 0xfd, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x38, 0x1a, 0xb8, 0xb3, 0x89, 0x01, 0x00, 0x00,
}
