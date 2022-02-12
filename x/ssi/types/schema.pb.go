// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ssi/v1/schema.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Schema struct {
	Type         string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	ModelVersion string          `protobuf:"bytes,2,opt,name=modelVersion,proto3" json:"modelVersion,omitempty"`
	Id           string          `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Name         string          `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Author       string          `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Authored     string          `protobuf:"bytes,6,opt,name=authored,proto3" json:"authored,omitempty"`
	Schema       *SchemaProperty `protobuf:"bytes,7,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d55333c8301c7bb, []int{0}
}
func (m *Schema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return m.Size()
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Schema) GetModelVersion() string {
	if m != nil {
		return m.ModelVersion
	}
	return ""
}

func (m *Schema) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Schema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Schema) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Schema) GetAuthored() string {
	if m != nil {
		return m.Authored
	}
	return ""
}

func (m *Schema) GetSchema() *SchemaProperty {
	if m != nil {
		return m.Schema
	}
	return nil
}

type SchemaProperty struct {
	Schema               string   `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Properties           string   `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	Required             []string `protobuf:"bytes,5,rep,name=required,proto3" json:"required,omitempty"`
	AdditionalProperties bool     `protobuf:"varint,6,opt,name=additionalProperties,proto3" json:"additionalProperties,omitempty"`
}

func (m *SchemaProperty) Reset()         { *m = SchemaProperty{} }
func (m *SchemaProperty) String() string { return proto.CompactTextString(m) }
func (*SchemaProperty) ProtoMessage()    {}
func (*SchemaProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d55333c8301c7bb, []int{1}
}
func (m *SchemaProperty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SchemaProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SchemaProperty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SchemaProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaProperty.Merge(m, src)
}
func (m *SchemaProperty) XXX_Size() int {
	return m.Size()
}
func (m *SchemaProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaProperty.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaProperty proto.InternalMessageInfo

func (m *SchemaProperty) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *SchemaProperty) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SchemaProperty) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SchemaProperty) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func (m *SchemaProperty) GetRequired() []string {
	if m != nil {
		return m.Required
	}
	return nil
}

func (m *SchemaProperty) GetAdditionalProperties() bool {
	if m != nil {
		return m.AdditionalProperties
	}
	return false
}

func init() {
	proto.RegisterType((*Schema)(nil), "hypersignprotocol.hidnode.ssi.Schema")
	proto.RegisterType((*SchemaProperty)(nil), "hypersignprotocol.hidnode.ssi.SchemaProperty")
}

func init() { proto.RegisterFile("ssi/v1/schema.proto", fileDescriptor_3d55333c8301c7bb) }

var fileDescriptor_3d55333c8301c7bb = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0xeb, 0x7e, 0xe4, 0xdf, 0xbf, 0x8b, 0x3a, 0x18, 0x84, 0x4c, 0x25, 0xac, 0xa8, 0x53,
	0x97, 0xc6, 0xa2, 0xbc, 0x01, 0x12, 0x1b, 0x43, 0x55, 0x24, 0x06, 0xb6, 0x34, 0x36, 0x89, 0xa5,
	0x24, 0x0e, 0x76, 0x82, 0xc8, 0x5b, 0xf0, 0x58, 0x4c, 0xa8, 0x23, 0x1b, 0xa8, 0x7d, 0x11, 0x64,
	0x3b, 0xfd, 0x92, 0x10, 0xdb, 0x3d, 0xf7, 0xc8, 0x47, 0xf7, 0x77, 0x64, 0x78, 0xaa, 0xb5, 0xa0,
	0x2f, 0x57, 0x54, 0x47, 0x09, 0xcf, 0xc2, 0xa0, 0x50, 0xb2, 0x94, 0xe8, 0x32, 0xa9, 0x0b, 0xae,
	0xb4, 0x88, 0x73, 0xab, 0x23, 0x99, 0x06, 0x89, 0x60, 0xb9, 0x64, 0x3c, 0xd0, 0x5a, 0x8c, 0x2e,
	0x62, 0x29, 0xe3, 0x94, 0x53, 0x6b, 0x2e, 0xab, 0x27, 0x1a, 0xe6, 0xb5, 0x7b, 0x39, 0xfe, 0x02,
	0xd0, 0xbb, 0xb7, 0x51, 0x08, 0xc1, 0x6e, 0x59, 0x17, 0x1c, 0x03, 0x1f, 0x4c, 0xfe, 0x2f, 0xec,
	0x8c, 0xc6, 0xf0, 0x24, 0x93, 0x8c, 0xa7, 0x0f, 0x26, 0x5d, 0xe6, 0xb8, 0x6d, 0xbd, 0xa3, 0x1d,
	0x1a, 0xc2, 0xb6, 0x60, 0xb8, 0x63, 0x9d, 0xb6, 0x60, 0x26, 0x27, 0x0f, 0x33, 0x8e, 0xbb, 0x2e,
	0xc7, 0xcc, 0xe8, 0x1c, 0x7a, 0x61, 0x55, 0x26, 0x52, 0xe1, 0x9e, 0xdd, 0x36, 0x0a, 0x8d, 0x60,
	0xdf, 0x4d, 0x9c, 0x61, 0xcf, 0x3a, 0x3b, 0x8d, 0x6e, 0xa1, 0xe7, 0x20, 0xf1, 0x3f, 0x1f, 0x4c,
	0x06, 0xb3, 0x69, 0xf0, 0x27, 0x65, 0xe0, 0x30, 0xe6, 0x4a, 0x16, 0x5c, 0x95, 0xf5, 0xa2, 0x79,
	0x3c, 0xfe, 0x00, 0x70, 0x78, 0x6c, 0x99, 0x6b, 0x9a, 0x64, 0xc7, 0xda, 0x28, 0xe4, 0xc3, 0x01,
	0xe3, 0x3a, 0x52, 0xa2, 0x28, 0xf7, 0xb0, 0x87, 0xab, 0x5d, 0x47, 0x9d, 0x83, 0x8e, 0x08, 0x84,
	0x85, 0x4b, 0x16, 0x5c, 0x37, 0xd4, 0x07, 0x1b, 0xc3, 0xa8, 0xf8, 0x73, 0x25, 0x0c, 0x63, 0xcf,
	0xef, 0x18, 0xc6, 0xad, 0x46, 0x33, 0x78, 0x16, 0x32, 0x26, 0x4c, 0x76, 0x98, 0xce, 0xf7, 0x29,
	0xa6, 0x8b, 0xfe, 0xe2, 0x57, 0xef, 0xe6, 0xee, 0x7d, 0x4d, 0xc0, 0x6a, 0x4d, 0xc0, 0xf7, 0x9a,
	0x80, 0xb7, 0x0d, 0x69, 0xad, 0x36, 0xa4, 0xf5, 0xb9, 0x21, 0xad, 0xc7, 0x59, 0x2c, 0xca, 0xa4,
	0x5a, 0x06, 0x91, 0xcc, 0xe8, 0xae, 0xab, 0xe9, 0xb6, 0x2c, 0x9a, 0x08, 0x36, 0x35, 0x6d, 0xd1,
	0x57, 0x6a, 0x3e, 0x91, 0x39, 0x5e, 0x2f, 0x3d, 0x6b, 0x5f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x8c, 0x93, 0x9a, 0x3f, 0x58, 0x02, 0x00, 0x00,
}

func (m *Schema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Schema != nil {
		{
			size, err := m.Schema.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSchema(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Authored) > 0 {
		i -= len(m.Authored)
		copy(dAtA[i:], m.Authored)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Authored)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Author) > 0 {
		i -= len(m.Author)
		copy(dAtA[i:], m.Author)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Author)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ModelVersion) > 0 {
		i -= len(m.ModelVersion)
		copy(dAtA[i:], m.ModelVersion)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.ModelVersion)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SchemaProperty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaProperty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SchemaProperty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AdditionalProperties {
		i--
		if m.AdditionalProperties {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Required) > 0 {
		for iNdEx := len(m.Required) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Required[iNdEx])
			copy(dAtA[i:], m.Required[iNdEx])
			i = encodeVarintSchema(dAtA, i, uint64(len(m.Required[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Properties) > 0 {
		i -= len(m.Properties)
		copy(dAtA[i:], m.Properties)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Properties)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Schema) > 0 {
		i -= len(m.Schema)
		copy(dAtA[i:], m.Schema)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Schema)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Schema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.ModelVersion)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Authored)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.Schema != nil {
		l = m.Schema.Size()
		n += 1 + l + sovSchema(uint64(l))
	}
	return n
}

func (m *SchemaProperty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Properties)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if len(m.Required) > 0 {
		for _, s := range m.Required {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.AdditionalProperties {
		n += 2
	}
	return n
}

func sovSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Schema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Schema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModelVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authored", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authored = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Schema == nil {
				m.Schema = &SchemaProperty{}
			}
			if err := m.Schema.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
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
func (m *SchemaProperty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaProperty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaProperty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Required", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Required = append(m.Required, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalProperties", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AdditionalProperties = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
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
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
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
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
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
			if length < 0 {
				return 0, ErrInvalidLengthSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchema = fmt.Errorf("proto: unexpected end of group")
)