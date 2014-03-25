// Code generated by protoc-gen-go.
// source: DataFrame.proto
// DO NOT EDIT!

package dataframe

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type DataFrame_Type int32

const (
	DataFrame_EMPTY  DataFrame_Type = 0
	DataFrame_NUMBER DataFrame_Type = 1
	DataFrame_REAL   DataFrame_Type = 2
	DataFrame_TEXT   DataFrame_Type = 3
	DataFrame_BINARY DataFrame_Type = 4
)

var DataFrame_Type_name = map[int32]string{
	0: "EMPTY",
	1: "NUMBER",
	2: "REAL",
	3: "TEXT",
	4: "BINARY",
}
var DataFrame_Type_value = map[string]int32{
	"EMPTY":  0,
	"NUMBER": 1,
	"REAL":   2,
	"TEXT":   3,
	"BINARY": 4,
}

func (x DataFrame_Type) Enum() *DataFrame_Type {
	p := new(DataFrame_Type)
	*p = x
	return p
}
func (x DataFrame_Type) String() string {
	return proto.EnumName(DataFrame_Type_name, int32(x))
}
func (x DataFrame_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DataFrame_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataFrame_Type_value, data, "DataFrame_Type")
	if err != nil {
		return err
	}
	*x = DataFrame_Type(value)
	return nil
}

type DataFrame struct {
	Source           []*DataFrame_Tag `protobuf:"bytes,1,rep,name=source" json:"source,omitempty"`
	Timestamp        *uint64          `protobuf:"fixed64,2,req,name=timestamp" json:"timestamp,omitempty"`
	Payload          *DataFrame_Type  `protobuf:"varint,3,req,name=payload,enum=dataframe.DataFrame_Type" json:"payload,omitempty"`
	ValueNumeric     *int64           `protobuf:"varint,4,opt,name=value_numeric" json:"value_numeric,omitempty"`
	ValueMeasurement *float64         `protobuf:"fixed64,5,opt,name=value_measurement" json:"value_measurement,omitempty"`
	ValueTextual     *string          `protobuf:"bytes,6,opt,name=value_textual" json:"value_textual,omitempty"`
	ValueBlob        []byte           `protobuf:"bytes,7,opt,name=value_blob" json:"value_blob,omitempty"`
	Origin           []byte           `protobuf:"bytes,8,opt,name=origin" json:"origin,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DataFrame) Reset()         { *m = DataFrame{} }
func (m *DataFrame) String() string { return proto.CompactTextString(m) }
func (*DataFrame) ProtoMessage()    {}

func (m *DataFrame) GetSource() []*DataFrame_Tag {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *DataFrame) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *DataFrame) GetPayload() DataFrame_Type {
	if m != nil && m.Payload != nil {
		return *m.Payload
	}
	return DataFrame_EMPTY
}

func (m *DataFrame) GetValueNumeric() int64 {
	if m != nil && m.ValueNumeric != nil {
		return *m.ValueNumeric
	}
	return 0
}

func (m *DataFrame) GetValueMeasurement() float64 {
	if m != nil && m.ValueMeasurement != nil {
		return *m.ValueMeasurement
	}
	return 0
}

func (m *DataFrame) GetValueTextual() string {
	if m != nil && m.ValueTextual != nil {
		return *m.ValueTextual
	}
	return ""
}

func (m *DataFrame) GetValueBlob() []byte {
	if m != nil {
		return m.ValueBlob
	}
	return nil
}

func (m *DataFrame) GetOrigin() []byte {
	if m != nil {
		return m.Origin
	}
	return nil
}

type DataFrame_Tag struct {
	Field            *string `protobuf:"bytes,1,req,name=field" json:"field,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DataFrame_Tag) Reset()         { *m = DataFrame_Tag{} }
func (m *DataFrame_Tag) String() string { return proto.CompactTextString(m) }
func (*DataFrame_Tag) ProtoMessage()    {}

func (m *DataFrame_Tag) GetField() string {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return ""
}

func (m *DataFrame_Tag) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("dataframe.DataFrame_Type", DataFrame_Type_name, DataFrame_Type_value)
}