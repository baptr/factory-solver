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
	ProductionType_COLLIDED               ProductionType = 8
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
	8: "COLLIDED",
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
	"COLLIDED":               8,
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
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0xc6, 0x1d, 0x35, 0x31, 0x1e, 0x5d, 0x09, 0x43, 0xbb, 0x0d, 0x5b, 0xda, 0x06, 0xaf, 0x82,
	0xb0, 0xca, 0xba, 0xf4, 0xae, 0x14, 0xd4, 0x4c, 0x51, 0xaa, 0x6e, 0x3b, 0xd9, 0x6d, 0x61, 0x6f,
	0x16, 0x8d, 0xb3, 0x76, 0x20, 0xff, 0x9a, 0x4c, 0x0a, 0xbe, 0x4b, 0x9f, 0xa4, 0x4f, 0x57, 0x66,
	0x12, 0x45, 0xbd, 0x69, 0xef, 0xe6, 0x9c, 0xf3, 0x9b, 0x33, 0xe7, 0x7c, 0x1f, 0x03, 0x6d, 0x3f,
	0x8e, 0x9e, 0xf9, 0xb6, 0x9f, 0xa4, 0xb1, 0x88, 0xf1, 0xc5, 0xf3, 0xca, 0x17, 0x71, 0xba, 0xcb,
	0xe2, 0xe0, 0x17, 0x4b, 0xbb, 0x7f, 0xaa, 0xa0, 0x53, 0xe6, 0xf3, 0x84, 0x61, 0x0c, 0xf5, 0x68,
	0x15, 0x32, 0x0b, 0xd9, 0xc8, 0x69, 0x52, 0x75, 0xc6, 0xb7, 0xa0, 0xa7, 0x2c, 0xcb, 0x03, 0x61,
	0x55, 0xed, 0x9a, 0xd3, 0x1a, 0xbe, 0xee, 0x9f, 0x5c, 0xef, 0xcf, 0x04, 0x0b, 0xbf, 0xe6, 0xab,
	0x48, 0x70, 0xb1, 0xa3, 0x25, 0x8a, 0x6f, 0x40, 0xe3, 0x51, 0x92, 0x0b, 0xab, 0xf6, 0xef, 0x3b,
	0x05, 0x89, 0xdf, 0x83, 0xb1, 0xc9, 0xd3, 0x95, 0xe0, 0x71, 0x64, 0xd5, 0x6d, 0xe4, 0xb4, 0x86,
	0xaf, 0xce, 0x6e, 0xb9, 0x65, 0x79, 0x5a, 0xa1, 0x07, 0x14, 0xbf, 0x03, 0x48, 0x58, 0xfa, 0x14,
	0xf2, 0x28, 0x17, 0xcc, 0xd2, 0x6c, 0xe4, 0xa0, 0x69, 0x85, 0x36, 0x13, 0x96, 0x2e, 0x54, 0x6a,
	0x0f, 0x64, 0xcc, 0x8f, 0xa3, 0x8d, 0xa5, 0x1f, 0x01, 0x9e, 0x4a, 0xe1, 0x1b, 0xa8, 0x8b, 0x5d,
	0xc2, 0xac, 0x86, 0x8d, 0x9c, 0xce, 0xf0, 0xcd, 0xd9, 0xa3, 0x5f, 0xd2, 0x78, 0x93, 0xfb, 0xf2,
	0xa9, 0xfb, 0x5d, 0xc2, 0xa8, 0x42, 0xc7, 0x06, 0xe8, 0x82, 0x87, 0x3c, 0xda, 0x76, 0x3f, 0x42,
	0xfb, 0x78, 0x19, 0xa9, 0x20, 0x17, 0x2c, 0xdc, 0x2b, 0x28, 0xcf, 0xf8, 0x0a, 0x8c, 0x9f, 0x65,
	0xdd, 0xaa, 0xda, 0xc8, 0xd1, 0xe8, 0x21, 0xee, 0x3e, 0x82, 0x36, 0x8e, 0xa3, 0x3c, 0x3b, 0x4c,
	0x81, 0xec, 0xda, 0x7f, 0x4e, 0x81, 0xdf, 0x02, 0x84, 0x79, 0x20, 0x78, 0x12, 0x70, 0x96, 0xaa,
	0xce, 0x88, 0x1e, 0x65, 0xba, 0x3e, 0xe8, 0x13, 0xe5, 0x3b, 0xbe, 0x96, 0x1e, 0x4a, 0x87, 0x55,
	0xfb, 0xd6, 0xf0, 0xe5, 0x59, 0xfb, 0xc2, 0x7e, 0x5a, 0x42, 0xb8, 0x07, 0xda, 0x5a, 0x0e, 0x55,
	0x3a, 0xfe, 0xe2, 0x8c, 0x56, 0x03, 0xd3, 0x02, 0xe9, 0x7e, 0x00, 0x63, 0xef, 0x0b, 0xb6, 0xa0,
	0x51, 0xc8, 0x9c, 0xa9, 0xfd, 0x35, 0xba, 0x0f, 0xf1, 0x25, 0xe8, 0x21, 0x0f, 0x02, 0x9e, 0x95,
	0x02, 0x94, 0x51, 0xef, 0x37, 0x82, 0xce, 0xe9, 0x6e, 0xf8, 0x0a, 0x2e, 0x4f, 0x33, 0x4f, 0x0f,
	0xcb, 0xcf, 0xcb, 0xbb, 0xef, 0x4b, 0xb3, 0x82, 0x2f, 0xa0, 0x39, 0x1d, 0xd1, 0x6f, 0xc4, 0xbb,
	0x27, 0xae, 0x89, 0x70, 0x0b, 0x1a, 0xde, 0x82, 0xcc, 0x65, 0x50, 0x95, 0xb5, 0x91, 0xe7, 0x91,
	0xc5, 0x78, 0x4e, 0x5c, 0xb3, 0x26, 0x6b, 0x94, 0x7c, 0x9a, 0x2d, 0x89, 0x6b, 0xd6, 0x71, 0x1b,
	0x8c, 0xc9, 0x94, 0x2c, 0x66, 0x93, 0xd1, 0xdc, 0xd4, 0x30, 0x80, 0x3e, 0x7e, 0xa0, 0xb2, 0xa2,
	0xe3, 0x0e, 0x00, 0x25, 0x1e, 0x19, 0xd1, 0xc9, 0x94, 0xb8, 0x66, 0x43, 0x91, 0x77, 0xf3, 0xf9,
	0xcc, 0x25, 0xae, 0x69, 0x8c, 0x7b, 0x8f, 0xce, 0x96, 0x8b, 0x1f, 0xf9, 0xba, 0xef, 0xc7, 0xe1,
	0x60, 0xbd, 0x4a, 0x44, 0x3a, 0x28, 0xb5, 0xb8, 0x2e, 0xc4, 0x18, 0x14, 0x3f, 0x2b, 0x59, 0xaf,
	0x75, 0xf5, 0xb9, 0x6e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xee, 0xf9, 0x5a, 0x76, 0x6c, 0x03,
	0x00, 0x00,
}
