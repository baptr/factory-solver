// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package configpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProductionType int32

const (
	ProductionType_ProductionType_UNKNOWN ProductionType = 0
	ProductionType_HARVESTED              ProductionType = 1
	ProductionType_SMELTED                ProductionType = 2
	ProductionType_ASSEMBLED              ProductionType = 3
	ProductionType_REFINED                ProductionType = 4
	ProductionType_CHEMICAL               ProductionType = 5
	ProductionType_BURNED                 ProductionType = 6
	ProductionType_RESEARCHED             ProductionType = 7
)

var ProductionType_name = map[int32]string{
	0: "ProductionType_UNKNOWN",
	1: "HARVESTED",
	2: "SMELTED",
	3: "ASSEMBLED",
	4: "REFINED",
	5: "CHEMICAL",
	6: "BURNED",
	7: "RESEARCHED",
}

var ProductionType_value = map[string]int32{
	"ProductionType_UNKNOWN": 0,
	"HARVESTED":              1,
	"SMELTED":                2,
	"ASSEMBLED":              3,
	"REFINED":                4,
	"CHEMICAL":               5,
	"BURNED":                 6,
	"RESEARCHED":             7,
}

func (x ProductionType) String() string {
	return proto.EnumName(ProductionType_name, int32(x))
}

func (ProductionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

type Recipe struct {
	Name   string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Result []*ItemQuantity `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	Input  []*ItemQuantity `protobuf:"bytes,3,rep,name=input,proto3" json:"input,omitempty"`
	// Types that are valid to be assigned to Timing:
	//	*Recipe_Duration
	//	*Recipe_PerMinute
	//	*Recipe_PerSecond
	Timing               isRecipe_Timing `protobuf_oneof:"timing"`
	Type                 ProductionType  `protobuf:"varint,7,opt,name=type,proto3,enum=factorysolver.ProductionType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Recipe) Reset()         { *m = Recipe{} }
func (m *Recipe) String() string { return proto.CompactTextString(m) }
func (*Recipe) ProtoMessage()    {}
func (*Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipe.Unmarshal(m, b)
}
func (m *Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipe.Marshal(b, m, deterministic)
}
func (m *Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipe.Merge(m, src)
}
func (m *Recipe) XXX_Size() int {
	return xxx_messageInfo_Recipe.Size(m)
}
func (m *Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_Recipe proto.InternalMessageInfo

func (m *Recipe) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Recipe) GetResult() []*ItemQuantity {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Recipe) GetInput() []*ItemQuantity {
	if m != nil {
		return m.Input
	}
	return nil
}

type isRecipe_Timing interface {
	isRecipe_Timing()
}

type Recipe_Duration struct {
	Duration *Duration `protobuf:"bytes,4,opt,name=duration,proto3,oneof"`
}

type Recipe_PerMinute struct {
	PerMinute float64 `protobuf:"fixed64,5,opt,name=per_minute,json=perMinute,proto3,oneof"`
}

type Recipe_PerSecond struct {
	PerSecond float64 `protobuf:"fixed64,6,opt,name=per_second,json=perSecond,proto3,oneof"`
}

func (*Recipe_Duration) isRecipe_Timing() {}

func (*Recipe_PerMinute) isRecipe_Timing() {}

func (*Recipe_PerSecond) isRecipe_Timing() {}

func (m *Recipe) GetTiming() isRecipe_Timing {
	if m != nil {
		return m.Timing
	}
	return nil
}

func (m *Recipe) GetDuration() *Duration {
	if x, ok := m.GetTiming().(*Recipe_Duration); ok {
		return x.Duration
	}
	return nil
}

func (m *Recipe) GetPerMinute() float64 {
	if x, ok := m.GetTiming().(*Recipe_PerMinute); ok {
		return x.PerMinute
	}
	return 0
}

func (m *Recipe) GetPerSecond() float64 {
	if x, ok := m.GetTiming().(*Recipe_PerSecond); ok {
		return x.PerSecond
	}
	return 0
}

func (m *Recipe) GetType() ProductionType {
	if m != nil {
		return m.Type
	}
	return ProductionType_ProductionType_UNKNOWN
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Recipe) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Recipe_Duration)(nil),
		(*Recipe_PerMinute)(nil),
		(*Recipe_PerSecond)(nil),
	}
}

type ItemQuantity struct {
	Item                 string   `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Quantity             int32    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemQuantity) Reset()         { *m = ItemQuantity{} }
func (m *ItemQuantity) String() string { return proto.CompactTextString(m) }
func (*ItemQuantity) ProtoMessage()    {}
func (*ItemQuantity) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *ItemQuantity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemQuantity.Unmarshal(m, b)
}
func (m *ItemQuantity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemQuantity.Marshal(b, m, deterministic)
}
func (m *ItemQuantity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemQuantity.Merge(m, src)
}
func (m *ItemQuantity) XXX_Size() int {
	return xxx_messageInfo_ItemQuantity.Size(m)
}
func (m *ItemQuantity) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemQuantity.DiscardUnknown(m)
}

var xxx_messageInfo_ItemQuantity proto.InternalMessageInfo

func (m *ItemQuantity) GetItem() string {
	if m != nil {
		return m.Item
	}
	return ""
}

func (m *ItemQuantity) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type Bonus struct {
	Type                 []ProductionType `protobuf:"varint,1,rep,packed,name=type,proto3,enum=factorysolver.ProductionType" json:"type,omitempty"`
	Multiplier           float64          `protobuf:"fixed64,2,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Bonus) Reset()         { *m = Bonus{} }
func (m *Bonus) String() string { return proto.CompactTextString(m) }
func (*Bonus) ProtoMessage()    {}
func (*Bonus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *Bonus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bonus.Unmarshal(m, b)
}
func (m *Bonus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bonus.Marshal(b, m, deterministic)
}
func (m *Bonus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bonus.Merge(m, src)
}
func (m *Bonus) XXX_Size() int {
	return xxx_messageInfo_Bonus.Size(m)
}
func (m *Bonus) XXX_DiscardUnknown() {
	xxx_messageInfo_Bonus.DiscardUnknown(m)
}

var xxx_messageInfo_Bonus proto.InternalMessageInfo

func (m *Bonus) GetType() []ProductionType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Bonus) GetMultiplier() float64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type Config struct {
	Recipe               []*Recipe `protobuf:"bytes,1,rep,name=recipe,proto3" json:"recipe,omitempty"`
	Bonus                []*Bonus  `protobuf:"bytes,2,rep,name=bonus,proto3" json:"bonus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetRecipe() []*Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

func (m *Config) GetBonus() []*Bonus {
	if m != nil {
		return m.Bonus
	}
	return nil
}

type Duration struct {
	Seconds              int32    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Millis               int32    `protobuf:"varint,2,opt,name=millis,proto3" json:"millis,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Duration) Reset()         { *m = Duration{} }
func (m *Duration) String() string { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()    {}
func (*Duration) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{4}
}

func (m *Duration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Duration.Unmarshal(m, b)
}
func (m *Duration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Duration.Marshal(b, m, deterministic)
}
func (m *Duration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duration.Merge(m, src)
}
func (m *Duration) XXX_Size() int {
	return xxx_messageInfo_Duration.Size(m)
}
func (m *Duration) XXX_DiscardUnknown() {
	xxx_messageInfo_Duration.DiscardUnknown(m)
}

var xxx_messageInfo_Duration proto.InternalMessageInfo

func (m *Duration) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Duration) GetMillis() int32 {
	if m != nil {
		return m.Millis
	}
	return 0
}

func init() {
	proto.RegisterEnum("factorysolver.ProductionType", ProductionType_name, ProductionType_value)
	proto.RegisterType((*Recipe)(nil), "factorysolver.Recipe")
	proto.RegisterType((*ItemQuantity)(nil), "factorysolver.ItemQuantity")
	proto.RegisterType((*Bonus)(nil), "factorysolver.Bonus")
	proto.RegisterType((*Config)(nil), "factorysolver.Config")
	proto.RegisterType((*Duration)(nil), "factorysolver.Duration")
}

func init() {
	proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4)
}

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdf, 0x8b, 0xda, 0x40,
	0x10, 0x76, 0xd5, 0x44, 0x1d, 0x3d, 0x09, 0x43, 0x7b, 0x0d, 0x57, 0xda, 0x06, 0x9f, 0x82, 0x70,
	0xca, 0x79, 0xf4, 0xad, 0x14, 0xfc, 0xb1, 0x45, 0xa9, 0xda, 0x76, 0x73, 0xd7, 0xc2, 0xbd, 0x1c,
	0x1a, 0xf7, 0xec, 0x42, 0x7e, 0x35, 0xd9, 0x14, 0xfc, 0x27, 0xfa, 0x07, 0xf5, 0xaf, 0x2b, 0xd9,
	0x44, 0x51, 0x5f, 0x7a, 0x6f, 0x3b, 0xf3, 0x7d, 0x33, 0x3b, 0xdf, 0x7c, 0x0c, 0xb4, 0xdc, 0x30,
	0x78, 0x12, 0xdb, 0x5e, 0x14, 0x87, 0x32, 0xc4, 0x8b, 0xa7, 0x95, 0x2b, 0xc3, 0x78, 0x97, 0x84,
	0xde, 0x6f, 0x1e, 0x77, 0xfe, 0x96, 0x41, 0x67, 0xdc, 0x15, 0x11, 0x47, 0x84, 0x6a, 0xb0, 0xf2,
	0xb9, 0x49, 0x2c, 0x62, 0x37, 0x98, 0x7a, 0xe3, 0x2d, 0xe8, 0x31, 0x4f, 0x52, 0x4f, 0x9a, 0x65,
	0xab, 0x62, 0x37, 0x07, 0xaf, 0x7b, 0x27, 0xe5, 0xbd, 0x99, 0xe4, 0xfe, 0xb7, 0x74, 0x15, 0x48,
	0x21, 0x77, 0xac, 0xa0, 0xe2, 0x0d, 0x68, 0x22, 0x88, 0x52, 0x69, 0x56, 0xfe, 0x5f, 0x93, 0x33,
	0xf1, 0x3d, 0xd4, 0x37, 0x69, 0xbc, 0x92, 0x22, 0x0c, 0xcc, 0xaa, 0x45, 0xec, 0xe6, 0xe0, 0xd5,
	0x59, 0xd5, 0xa4, 0x80, 0xa7, 0x25, 0x76, 0xa0, 0xe2, 0x3b, 0x80, 0x88, 0xc7, 0x8f, 0xbe, 0x08,
	0x52, 0xc9, 0x4d, 0xcd, 0x22, 0x36, 0x99, 0x96, 0x58, 0x23, 0xe2, 0xf1, 0x42, 0xa5, 0xf6, 0x84,
	0x84, 0xbb, 0x61, 0xb0, 0x31, 0xf5, 0x23, 0x82, 0xa3, 0x52, 0x78, 0x03, 0x55, 0xb9, 0x8b, 0xb8,
	0x59, 0xb3, 0x88, 0xdd, 0x1e, 0xbc, 0x39, 0xfb, 0xf4, 0x6b, 0x1c, 0x6e, 0x52, 0x37, 0xfb, 0xea,
	0x6e, 0x17, 0x71, 0xa6, 0xa8, 0xa3, 0x3a, 0xe8, 0x52, 0xf8, 0x22, 0xd8, 0x76, 0x3e, 0x42, 0xeb,
	0x58, 0x4c, 0xb6, 0x41, 0x21, 0xb9, 0xbf, 0xdf, 0x60, 0xf6, 0xc6, 0x2b, 0xa8, 0xff, 0x2a, 0x70,
	0xb3, 0x6c, 0x11, 0x5b, 0x63, 0x87, 0xb8, 0xf3, 0x00, 0xda, 0x28, 0x0c, 0xd2, 0xe4, 0x30, 0x05,
	0xb1, 0x2a, 0xcf, 0x9c, 0x02, 0xdf, 0x02, 0xf8, 0xa9, 0x27, 0x45, 0xe4, 0x09, 0x1e, 0xab, 0xce,
	0x84, 0x1d, 0x65, 0x3a, 0x2e, 0xe8, 0x63, 0xe5, 0x3b, 0x5e, 0x67, 0x1e, 0x66, 0x0e, 0xab, 0xf6,
	0xcd, 0xc1, 0xcb, 0xb3, 0xf6, 0xb9, 0xfd, 0xac, 0x20, 0x61, 0x17, 0xb4, 0x75, 0x36, 0x54, 0xe1,
	0xf8, 0x8b, 0x33, 0xb6, 0x1a, 0x98, 0xe5, 0x94, 0xce, 0x07, 0xa8, 0xef, 0x7d, 0x41, 0x13, 0x6a,
	0xf9, 0x9a, 0x13, 0xa5, 0x5f, 0x63, 0xfb, 0x10, 0x2f, 0x41, 0xf7, 0x85, 0xe7, 0x89, 0xa4, 0x58,
	0x40, 0x11, 0x75, 0xff, 0x10, 0x68, 0x9f, 0x6a, 0xc3, 0x2b, 0xb8, 0x3c, 0xcd, 0x3c, 0xde, 0x2f,
	0x3f, 0x2f, 0xbf, 0xfc, 0x58, 0x1a, 0x25, 0xbc, 0x80, 0xc6, 0x74, 0xc8, 0xbe, 0x53, 0xe7, 0x8e,
	0x4e, 0x0c, 0x82, 0x4d, 0xa8, 0x39, 0x0b, 0x3a, 0xcf, 0x82, 0x72, 0x86, 0x0d, 0x1d, 0x87, 0x2e,
	0x46, 0x73, 0x3a, 0x31, 0x2a, 0x19, 0xc6, 0xe8, 0xa7, 0xd9, 0x92, 0x4e, 0x8c, 0x2a, 0xb6, 0xa0,
	0x3e, 0x9e, 0xd2, 0xc5, 0x6c, 0x3c, 0x9c, 0x1b, 0x1a, 0x02, 0xe8, 0xa3, 0x7b, 0x96, 0x21, 0x3a,
	0xb6, 0x01, 0x18, 0x75, 0xe8, 0x90, 0x8d, 0xa7, 0x74, 0x62, 0xd4, 0x46, 0xdd, 0x07, 0x7b, 0x2b,
	0xe4, 0xcf, 0x74, 0xdd, 0x73, 0x43, 0xbf, 0xbf, 0x5e, 0x45, 0x32, 0xee, 0x17, 0xea, 0xaf, 0x73,
	0xf9, 0xfd, 0xfc, 0x96, 0xa2, 0xf5, 0x5a, 0x57, 0xe7, 0x74, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x0b, 0xe1, 0xeb, 0x23, 0x5e, 0x03, 0x00, 0x00,
}
