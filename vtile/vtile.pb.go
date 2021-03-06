// Code generated by protoc-gen-go.
// source: protob/vtile/2.1/vtile.proto
// DO NOT EDIT!

/*
Package vtile is a generated protocol buffer package.

It is generated from these files:
	protob/vtile/2.1/vtile.proto

It has these top-level messages:
	Tile
*/
package vtile

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// GeomType is described in section 4.3.4 of the specification
type Tile_GeomType int32

const (
	Tile_UNKNOWN    Tile_GeomType = 0
	Tile_POINT      Tile_GeomType = 1
	Tile_LINESTRING Tile_GeomType = 2
	Tile_POLYGON    Tile_GeomType = 3
)

var Tile_GeomType_name = map[int32]string{
	0: "UNKNOWN",
	1: "POINT",
	2: "LINESTRING",
	3: "POLYGON",
}
var Tile_GeomType_value = map[string]int32{
	"UNKNOWN":    0,
	"POINT":      1,
	"LINESTRING": 2,
	"POLYGON":    3,
}

func (x Tile_GeomType) Enum() *Tile_GeomType {
	p := new(Tile_GeomType)
	*p = x
	return p
}
func (x Tile_GeomType) String() string {
	return proto.EnumName(Tile_GeomType_name, int32(x))
}
func (x *Tile_GeomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Tile_GeomType_value, data, "Tile_GeomType")
	if err != nil {
		return err
	}
	*x = Tile_GeomType(value)
	return nil
}
func (Tile_GeomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Tile struct {
	Layers           []*Tile_Layer             `protobuf:"bytes,3,rep,name=layers" json:"layers,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Tile) Reset()                    { *m = Tile{} }
func (m *Tile) String() string            { return proto.CompactTextString(m) }
func (*Tile) ProtoMessage()               {}
func (*Tile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

var extRange_Tile = []proto.ExtensionRange{
	{16, 8191},
}

func (*Tile) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Tile
}
func (m *Tile) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Tile) GetLayers() []*Tile_Layer {
	if m != nil {
		return m.Layers
	}
	return nil
}

// Variant type encoding
// The use of values is described in section 4.1 of the specification
type Tile_Value struct {
	// Exactly one of these values must be present in a valid message
	StringValue      *string                   `protobuf:"bytes,1,opt,name=string_value" json:"string_value,omitempty"`
	FloatValue       *float32                  `protobuf:"fixed32,2,opt,name=float_value" json:"float_value,omitempty"`
	DoubleValue      *float64                  `protobuf:"fixed64,3,opt,name=double_value" json:"double_value,omitempty"`
	IntValue         *int64                    `protobuf:"varint,4,opt,name=int_value" json:"int_value,omitempty"`
	UintValue        *uint64                   `protobuf:"varint,5,opt,name=uint_value" json:"uint_value,omitempty"`
	SintValue        *int64                    `protobuf:"zigzag64,6,opt,name=sint_value" json:"sint_value,omitempty"`
	BoolValue        *bool                     `protobuf:"varint,7,opt,name=bool_value" json:"bool_value,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Tile_Value) Reset()                    { *m = Tile_Value{} }
func (m *Tile_Value) String() string            { return proto.CompactTextString(m) }
func (*Tile_Value) ProtoMessage()               {}
func (*Tile_Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

var extRange_Tile_Value = []proto.ExtensionRange{
	{8, 536870911},
}

func (*Tile_Value) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Tile_Value
}
func (m *Tile_Value) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Tile_Value) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Tile_Value) GetFloatValue() float32 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Tile_Value) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *Tile_Value) GetIntValue() int64 {
	if m != nil && m.IntValue != nil {
		return *m.IntValue
	}
	return 0
}

func (m *Tile_Value) GetUintValue() uint64 {
	if m != nil && m.UintValue != nil {
		return *m.UintValue
	}
	return 0
}

func (m *Tile_Value) GetSintValue() int64 {
	if m != nil && m.SintValue != nil {
		return *m.SintValue
	}
	return 0
}

func (m *Tile_Value) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

// Features are described in section 4.2 of the specification
type Tile_Feature struct {
	Id *uint64 `protobuf:"varint,1,opt,name=id,def=0" json:"id,omitempty"`
	// Tags of this feature are encoded as repeated pairs of
	// integers.
	// A detailed description of tags is located in sections
	// 4.2 and 4.4 of the specification
	Tags []uint32 `protobuf:"varint,2,rep,packed,name=tags" json:"tags,omitempty"`
	// The type of geometry stored in this feature.
	Type *Tile_GeomType `protobuf:"varint,3,opt,name=type,enum=vector_tile.Tile_GeomType,def=0" json:"type,omitempty"`
	// Contains a stream of commands and parameters (vertices).
	// A detailed description on geometry encoding is located in
	// section 4.3 of the specification.
	Geometry         []uint32 `protobuf:"varint,4,rep,packed,name=geometry" json:"geometry,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Tile_Feature) Reset()                    { *m = Tile_Feature{} }
func (m *Tile_Feature) String() string            { return proto.CompactTextString(m) }
func (*Tile_Feature) ProtoMessage()               {}
func (*Tile_Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

const Default_Tile_Feature_Id uint64 = 0
const Default_Tile_Feature_Type Tile_GeomType = Tile_UNKNOWN

func (m *Tile_Feature) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return Default_Tile_Feature_Id
}

func (m *Tile_Feature) GetTags() []uint32 {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Tile_Feature) GetType() Tile_GeomType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Tile_Feature_Type
}

func (m *Tile_Feature) GetGeometry() []uint32 {
	if m != nil {
		return m.Geometry
	}
	return nil
}

// Layers are described in section 4.1 of the specification
type Tile_Layer struct {
	// Any compliant implementation must first read the version
	// number encoded in this message and choose the correct
	// implementation for this version number before proceeding to
	// decode other parts of this message.
	Version *uint32 `protobuf:"varint,15,req,name=version,def=1" json:"version,omitempty"`
	Name    *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The actual features in this tile.
	Features []*Tile_Feature `protobuf:"bytes,2,rep,name=features" json:"features,omitempty"`
	// Dictionary encoding for keys
	Keys []string `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
	// Dictionary encoding for values
	Values []*Tile_Value `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
	// Although this is an "optional" field it is required by the specification.
	// See https://github.com/mapbox/vector-tile-spec/issues/47
	Extent           *uint32                   `protobuf:"varint,5,opt,name=extent,def=4096" json:"extent,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Tile_Layer) Reset()                    { *m = Tile_Layer{} }
func (m *Tile_Layer) String() string            { return proto.CompactTextString(m) }
func (*Tile_Layer) ProtoMessage()               {}
func (*Tile_Layer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

var extRange_Tile_Layer = []proto.ExtensionRange{
	{16, 536870911},
}

func (*Tile_Layer) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Tile_Layer
}
func (m *Tile_Layer) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_Tile_Layer_Version uint32 = 1
const Default_Tile_Layer_Extent uint32 = 4096

func (m *Tile_Layer) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Default_Tile_Layer_Version
}

func (m *Tile_Layer) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Tile_Layer) GetFeatures() []*Tile_Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *Tile_Layer) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Tile_Layer) GetValues() []*Tile_Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Tile_Layer) GetExtent() uint32 {
	if m != nil && m.Extent != nil {
		return *m.Extent
	}
	return Default_Tile_Layer_Extent
}

func init() {
	proto.RegisterType((*Tile)(nil), "vector_tile.Tile")
	proto.RegisterType((*Tile_Value)(nil), "vector_tile.Tile.Value")
	proto.RegisterType((*Tile_Feature)(nil), "vector_tile.Tile.Feature")
	proto.RegisterType((*Tile_Layer)(nil), "vector_tile.Tile.Layer")
	proto.RegisterEnum("vector_tile.Tile_GeomType", Tile_GeomType_name, Tile_GeomType_value)
}

var fileDescriptor0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6e, 0x94, 0x40,
	0x18, 0xc5, 0x65, 0x18, 0x76, 0xe1, 0xa3, 0x5b, 0x71, 0x34, 0x71, 0x25, 0x5e, 0x6c, 0x7a, 0x23,
	0xa9, 0x09, 0xed, 0x12, 0x63, 0x22, 0x37, 0x9a, 0x26, 0xba, 0x6e, 0xdc, 0x40, 0xa3, 0xa8, 0xd1,
	0x9b, 0x86, 0xb5, 0xd3, 0x0d, 0x91, 0x32, 0x0d, 0x0c, 0x44, 0xee, 0x7c, 0x15, 0x1f, 0xc0, 0x77,
	0xf0, 0xd1, 0xfc, 0x86, 0x3f, 0xa9, 0xc9, 0xca, 0xd5, 0xf0, 0xe3, 0x30, 0x73, 0xce, 0xf9, 0x06,
	0x1e, 0xdf, 0x94, 0x42, 0x8a, 0xed, 0x49, 0x23, 0xb3, 0x9c, 0x9f, 0x04, 0xfe, 0xb2, 0x5f, 0xf9,
	0x1d, 0x66, 0x76, 0xc3, 0xbf, 0x49, 0x51, 0x5e, 0x28, 0x74, 0xf4, 0x87, 0x02, 0x4d, 0x70, 0xc1,
	0x9e, 0xc0, 0x24, 0x4f, 0x5b, 0x5e, 0x56, 0x73, 0x7d, 0xa1, 0x7b, 0x76, 0xf0, 0xd0, 0xff, 0x47,
	0xe6, 0x2b, 0x89, 0xbf, 0x51, 0xdf, 0xdd, 0x5f, 0x1a, 0x18, 0x9f, 0xd2, 0xbc, 0xe6, 0xec, 0x01,
	0x1c, 0x54, 0xb2, 0xcc, 0x8a, 0xdd, 0x45, 0xa3, 0xde, 0xe7, 0xda, 0x42, 0xf3, 0x2c, 0x76, 0x1f,
	0xec, 0xab, 0x5c, 0xa4, 0x72, 0x80, 0x04, 0x21, 0x51, 0xd2, 0x4b, 0x51, 0x6f, 0x73, 0x3e, 0x50,
	0x1d, 0xa9, 0xc6, 0xee, 0x81, 0x95, 0x15, 0xa3, 0x90, 0x22, 0xd2, 0x19, 0x03, 0xa8, 0x6f, 0x99,
	0x81, 0x8c, 0x2a, 0x56, 0xdd, 0xb2, 0x09, 0x32, 0xa6, 0xd8, 0x56, 0x88, 0x7c, 0x60, 0x53, 0x64,
	0xe6, 0xb1, 0x69, 0x9a, 0xce, 0x4f, 0x7c, 0x88, 0x5b, 0xc1, 0xf4, 0x0d, 0x4f, 0x65, 0x5d, 0x72,
	0x36, 0x03, 0x92, 0x5d, 0x76, 0xd6, 0x68, 0xa8, 0x9d, 0x32, 0x07, 0xa8, 0x4c, 0x77, 0x15, 0xda,
	0xd2, 0xbd, 0xd9, 0x19, 0x71, 0x34, 0xb6, 0x44, 0xd2, 0xde, 0xf4, 0x96, 0x0e, 0x03, 0x77, 0x3f,
	0xf6, 0x8a, 0x8b, 0xeb, 0x04, 0x15, 0xe1, 0xf4, 0x63, 0xf4, 0x2e, 0x8a, 0x3f, 0x47, 0x98, 0xc6,
	0xdc, 0x21, 0xe4, 0xb2, 0x6c, 0xd1, 0xf6, 0xb0, 0x91, 0xfb, 0x1b, 0x8b, 0xe9, 0x2a, 0x42, 0x73,
	0xd3, 0x06, 0x9b, 0xcc, 0x44, 0x31, 0xbf, 0xbb, 0x20, 0xde, 0x2c, 0xd4, 0x96, 0xec, 0x00, 0x68,
	0x91, 0x5e, 0xab, 0x92, 0x08, 0x96, 0xf4, 0x14, 0xcc, 0xab, 0xde, 0x60, 0x6f, 0xc5, 0x0e, 0x1e,
	0xed, 0x1f, 0x3c, 0x46, 0xc0, 0x5f, 0xbf, 0xf3, 0xb6, 0x1f, 0x8c, 0xa5, 0x06, 0xd5, 0x85, 0xae,
	0xba, 0xa3, 0xff, 0x3b, 0xa8, 0x71, 0x3c, 0x13, 0xfe, 0x43, 0xf2, 0x42, 0x76, 0x35, 0xce, 0x42,
	0xfa, 0xec, 0xf4, 0xc5, 0x73, 0x2c, 0xc9, 0xe9, 0x4b, 0x3a, 0x7a, 0x09, 0xe6, 0x18, 0x8d, 0xd9,
	0x30, 0x86, 0x73, 0xee, 0x30, 0x0b, 0x8c, 0xf3, 0x78, 0x1d, 0x25, 0x58, 0xce, 0x21, 0xc0, 0x66,
	0x1d, 0xbd, 0xfe, 0x90, 0xbc, 0x5f, 0x47, 0x2b, 0x87, 0x28, 0xdd, 0x79, 0xbc, 0xf9, 0xb2, 0x8a,
	0x23, 0x47, 0x3f, 0x36, 0xd4, 0x56, 0xaf, 0xce, 0xac, 0xb7, 0xfa, 0x57, 0xa3, 0xbb, 0x60, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x73, 0x35, 0xd5, 0x33, 0x79, 0x02, 0x00, 0x00,
}
