// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

package pkg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type PayloadHeader struct {
	//@inject_tag: json:"Name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"Name"`
	//@inject_tag: json:"Value"
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayloadHeader) Reset()         { *m = PayloadHeader{} }
func (m *PayloadHeader) String() string { return proto.CompactTextString(m) }
func (*PayloadHeader) ProtoMessage()    {}
func (*PayloadHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
}

func (m *PayloadHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadHeader.Unmarshal(m, b)
}
func (m *PayloadHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadHeader.Marshal(b, m, deterministic)
}
func (m *PayloadHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadHeader.Merge(m, src)
}
func (m *PayloadHeader) XXX_Size() int {
	return xxx_messageInfo_PayloadHeader.Size(m)
}
func (m *PayloadHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadHeader proto.InternalMessageInfo

func (m *PayloadHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PayloadHeader) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PayloadAttachment struct {
	//@inject_tag: json:"Name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"Name"`
	//@inject_tag: json:"Content"
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"Content"`
	//@inject_tag: json:"ContentType"
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"ContentType"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayloadAttachment) Reset()         { *m = PayloadAttachment{} }
func (m *PayloadAttachment) String() string { return proto.CompactTextString(m) }
func (*PayloadAttachment) ProtoMessage()    {}
func (*PayloadAttachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

func (m *PayloadAttachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadAttachment.Unmarshal(m, b)
}
func (m *PayloadAttachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadAttachment.Marshal(b, m, deterministic)
}
func (m *PayloadAttachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadAttachment.Merge(m, src)
}
func (m *PayloadAttachment) XXX_Size() int {
	return xxx_messageInfo_PayloadAttachment.Size(m)
}
func (m *PayloadAttachment) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadAttachment.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadAttachment proto.InternalMessageInfo

func (m *PayloadAttachment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PayloadAttachment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PayloadAttachment) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type Payload struct {
	//@inject_tag: protobuf:"varint,1,opt,name=template_id,json=TemplateId,proto3"
	TemplateId int32 `protobuf:"varint,1,opt,name=template_id,json=TemplateId,proto3" json:"template_id,omitempty"`
	//@inject_tag: `protobuf:"bytes,2,opt,name=template_alias,json=TemplateAlias,proto3"
	TemplateAlias string `protobuf:"bytes,2,opt,name=template_alias,json=TemplateAlias,proto3" json:"template_alias,omitempty"`
	//@inject_tag: protobuf:"bytes,3,rep,name=template_model,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"
	TemplateModel map[string]string `protobuf:"bytes,3,rep,name=template_model,proto3" json:"template_model,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//@inject_tag: protobuf:"varint,4,opt,name=inline_css,json=InlineCss,proto3"
	InlineCss bool `protobuf:"varint,4,opt,name=inline_css,json=InlineCss,proto3" json:"inline_css,omitempty"`
	//@inject_tag: protobuf:"bytes,5,opt,name=from,json=From,proto3"
	From string `protobuf:"bytes,5,opt,name=from,json=From,proto3" json:"from,omitempty"`
	//@inject_tag: protobuf:"bytes,6,opt,name=to,json=To,proto3"
	To string `protobuf:"bytes,6,opt,name=to,json=To,proto3" json:"to,omitempty"`
	//@inject_tag: protobuf:"bytes,7,opt,name=cc,json=Cc,proto3"
	Cc string `protobuf:"bytes,7,opt,name=cc,json=Cc,proto3" json:"cc,omitempty"`
	//@inject_tag: protobuf:"bytes,8,opt,name=bcc,json=Bcc,proto3"
	Bcc string `protobuf:"bytes,8,opt,name=bcc,json=Bcc,proto3" json:"bcc,omitempty"`
	//@inject_tag: protobuf:"bytes,9,opt,name=tag,json=Tag,proto3"
	Tag string `protobuf:"bytes,9,opt,name=tag,json=Tag,proto3" json:"tag,omitempty"`
	//@inject_tag: protobuf:"bytes,10,opt,name=reply_to,json=ReplyTo,proto3"
	ReplyTo string `protobuf:"bytes,10,opt,name=reply_to,json=ReplyTo,proto3" json:"reply_to,omitempty"`
	//@inject_tag: protobuf:"bytes,11,rep,name=headers,json=Headers,proto3"
	Headers []*PayloadHeader `protobuf:"bytes,11,rep,name=headers,json=Headers,proto3" json:"headers,omitempty"`
	//@inject_tag: protobuf:"varint,12,opt,name=track_opens,json=TrackOpens,proto3"
	TrackOpens bool `protobuf:"varint,12,opt,name=track_opens,json=TrackOpens,proto3" json:"track_opens,omitempty"`
	//@inject_tag: protobuf:"bytes,13,opt,name=track_links,json=TrackLinks,proto3"
	TrackLinks string `protobuf:"bytes,13,opt,name=track_links,json=TrackLinks,proto3" json:"track_links,omitempty"`
	//@inject_tag: protobuf:"bytes,14,rep,name=Attachments,json=Attachments,proto3"
	Attachments []*PayloadAttachment `protobuf:"bytes,14,rep,name=Attachments,json=Attachments,proto3" json:"attachments,omitempty"`
	//@inject_tag: protobuf:"bytes,15,rep,name=metadata,json=Metadata,proto3" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"
	Metadata map[string]string `protobuf:"bytes,15,rep,name=metadata,json=Metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//@inject_tag: protobuf:"bytes,16,opt,name=template_object_model,json=TemplateModel,proto3"
	TemplateObjectModel  *_struct.Struct `protobuf:"bytes,16,opt,name=template_object_model,json=TemplateModel,proto3" json:"template_object_model,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{2}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetTemplateId() int32 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *Payload) GetTemplateAlias() string {
	if m != nil {
		return m.TemplateAlias
	}
	return ""
}

func (m *Payload) GetTemplateModel() map[string]string {
	if m != nil {
		return m.TemplateModel
	}
	return nil
}

func (m *Payload) GetInlineCss() bool {
	if m != nil {
		return m.InlineCss
	}
	return false
}

func (m *Payload) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Payload) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Payload) GetCc() string {
	if m != nil {
		return m.Cc
	}
	return ""
}

func (m *Payload) GetBcc() string {
	if m != nil {
		return m.Bcc
	}
	return ""
}

func (m *Payload) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Payload) GetReplyTo() string {
	if m != nil {
		return m.ReplyTo
	}
	return ""
}

func (m *Payload) GetHeaders() []*PayloadHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Payload) GetTrackOpens() bool {
	if m != nil {
		return m.TrackOpens
	}
	return false
}

func (m *Payload) GetTrackLinks() string {
	if m != nil {
		return m.TrackLinks
	}
	return ""
}

func (m *Payload) GetAttachments() []*PayloadAttachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Payload) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Payload) GetTemplateObjectModel() *_struct.Struct {
	if m != nil {
		return m.TemplateObjectModel
	}
	return nil
}

func init() {
	proto.RegisterType((*PayloadHeader)(nil), "pkg.PayloadHeader")
	proto.RegisterType((*PayloadAttachment)(nil), "pkg.PayloadAttachment")
	proto.RegisterType((*Payload)(nil), "pkg.Payload")
	proto.RegisterMapType((map[string]string)(nil), "pkg.Payload.MetadataEntry")
	proto.RegisterMapType((map[string]string)(nil), "pkg.Payload.TemplateModelEntry")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0xc7, 0x49, 0xd2, 0x34, 0xc9, 0x5c, 0x13, 0xeb, 0xfa, 0x6b, 0x0d, 0x4a, 0x63, 0x40, 0xc8,
	0x83, 0x5c, 0xa1, 0x82, 0x54, 0x7d, 0xb1, 0x88, 0xa2, 0x68, 0xa9, 0x9c, 0x79, 0x0f, 0x9b, 0xbd,
	0xed, 0xf5, 0xbc, 0x1f, 0x7b, 0xdc, 0x4e, 0x84, 0xfb, 0x53, 0xfd, 0x6f, 0x64, 0x67, 0xf7, 0x92,
	0x14, 0xfb, 0xe2, 0x4b, 0x98, 0xf9, 0xdc, 0x77, 0x76, 0x76, 0xe7, 0x3b, 0x81, 0xa0, 0xaa, 0x35,
	0xea, 0x90, 0x7e, 0x59, 0xaf, 0xca, 0x92, 0xe9, 0xb3, 0x44, 0xeb, 0x24, 0x57, 0xa7, 0x84, 0xd6,
	0x9b, 0xeb, 0x53, 0x83, 0xf5, 0x46, 0xa2, 0x93, 0xcc, 0xdf, 0xc2, 0xf8, 0x87, 0x68, 0x72, 0x2d,
	0xe2, 0x2f, 0x4a, 0xc4, 0xaa, 0x66, 0x0c, 0x0e, 0x4a, 0x51, 0x28, 0xde, 0x99, 0x75, 0x16, 0xa3,
	0x88, 0x62, 0xf6, 0x10, 0xfa, 0xbf, 0x45, 0xbe, 0x51, 0xbc, 0x4b, 0xd0, 0x25, 0xf3, 0x18, 0xee,
	0xfb, 0xd2, 0x0b, 0x44, 0x21, 0x6f, 0x0a, 0x55, 0xe2, 0x9d, 0xe5, 0x1c, 0x06, 0x52, 0x97, 0xa8,
	0x4a, 0xf4, 0x07, 0xb4, 0x29, 0x7b, 0x01, 0x47, 0x3e, 0x5c, 0x61, 0x53, 0x29, 0xde, 0xa3, 0xcf,
	0x81, 0x67, 0xcb, 0xa6, 0x52, 0xf3, 0x3f, 0x7d, 0x18, 0xf8, 0x36, 0xec, 0x04, 0x02, 0x54, 0x45,
	0x95, 0x0b, 0x54, 0xab, 0x34, 0xa6, 0x1e, 0xfd, 0x08, 0x5a, 0xf4, 0x35, 0x66, 0x2f, 0x61, 0xb2,
	0x15, 0x88, 0x3c, 0x15, 0xc6, 0x37, 0x1c, 0xb7, 0xf4, 0xc2, 0x42, 0xf6, 0x79, 0x4f, 0x56, 0xe8,
	0x58, 0xe5, 0xbc, 0x37, 0xeb, 0x2d, 0x82, 0xb3, 0x93, 0xb0, 0xca, 0x92, 0xd0, 0x77, 0x0b, 0x97,
	0x5e, 0x72, 0x69, 0x15, 0x9f, 0x4a, 0xac, 0x9b, 0xdd, 0x39, 0xc4, 0xd8, 0x73, 0x80, 0xb4, 0xcc,
	0xd3, 0x52, 0xad, 0xa4, 0x31, 0xfc, 0x60, 0xd6, 0x59, 0x0c, 0xa3, 0x91, 0x23, 0x1f, 0x8d, 0xb1,
	0xb3, 0xb8, 0xae, 0x75, 0xc1, 0xfb, 0x6e, 0x16, 0x36, 0x66, 0x13, 0xe8, 0xa2, 0xe6, 0x87, 0x44,
	0xba, 0xa8, 0x6d, 0x2e, 0x25, 0x1f, 0xb8, 0x5c, 0x4a, 0x76, 0x0c, 0xbd, 0xb5, 0x94, 0x7c, 0x48,
	0xc0, 0x86, 0x96, 0xa0, 0x48, 0xf8, 0xc8, 0x11, 0x14, 0x09, 0x7b, 0x0a, 0xc3, 0x5a, 0x55, 0x79,
	0xb3, 0x42, 0xcd, 0xc1, 0x0d, 0x94, 0xf2, 0xa5, 0x66, 0xaf, 0x60, 0x70, 0x43, 0x3e, 0x1a, 0x1e,
	0xd0, 0x93, 0xd8, 0xfe, 0x93, 0x9c, 0xc5, 0x51, 0x2b, 0xa1, 0x79, 0xd6, 0x42, 0x66, 0x2b, 0x5d,
	0xa9, 0xd2, 0xf0, 0x23, 0x7a, 0x00, 0x10, 0xba, 0xb2, 0x64, 0x27, 0xc8, 0xd3, 0x32, 0x33, 0x7c,
	0x4c, 0xcd, 0x9c, 0xe0, 0xbb, 0x25, 0xec, 0x1c, 0x02, 0xb1, 0x35, 0xdf, 0xf0, 0x09, 0xf5, 0x7c,
	0xbc, 0xdf, 0x73, 0xb7, 0x1b, 0xd1, 0xbe, 0x94, 0xbd, 0x81, 0x61, 0xa1, 0x50, 0xc4, 0x02, 0x05,
	0xbf, 0x47, 0x65, 0xd3, 0x5b, 0xd3, 0xbf, 0xf4, 0x1f, 0xdd, 0xe0, 0xb7, 0x5a, 0xf6, 0x0d, 0x1e,
	0x6d, 0xbd, 0xd3, 0xeb, 0x5f, 0x4a, 0xa2, 0xb7, 0xf0, 0x78, 0xd6, 0x59, 0x04, 0x67, 0x4f, 0x42,
	0xb7, 0xee, 0x61, 0xbb, 0xee, 0xe1, 0x4f, 0x5a, 0xf7, 0xe8, 0x41, 0x5b, 0x75, 0x45, 0x45, 0x64,
	0xe0, 0xf4, 0x03, 0xb0, 0x7f, 0x5d, 0xb6, 0x13, 0xcf, 0x54, 0xe3, 0x57, 0xd8, 0x86, 0x77, 0xff,
	0x01, 0xde, 0x75, 0xcf, 0x3b, 0xd3, 0xf7, 0x30, 0xbe, 0x75, 0xd3, 0xff, 0x29, 0x5e, 0x1f, 0xd2,
	0x25, 0x5f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x39, 0x0f, 0x89, 0xb5, 0x03, 0x00, 0x00,
}