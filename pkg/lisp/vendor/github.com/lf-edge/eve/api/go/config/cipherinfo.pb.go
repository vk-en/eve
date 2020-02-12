// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cipherinfo.proto

package config

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

// Security Key Exchange Method
type KeyExchangeScheme int32

const (
	KeyExchangeScheme_KEA_NONE KeyExchangeScheme = 0
	KeyExchangeScheme_KEA_ECDH KeyExchangeScheme = 1
)

var KeyExchangeScheme_name = map[int32]string{
	0: "KEA_NONE",
	1: "KEA_ECDH",
}

var KeyExchangeScheme_value = map[string]int32{
	"KEA_NONE": 0,
	"KEA_ECDH": 1,
}

func (x KeyExchangeScheme) String() string {
	return proto.EnumName(KeyExchangeScheme_name, int32(x))
}

func (KeyExchangeScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{0}
}

// Encryption Scheme for Cipher Payload
type EncryptionScheme int32

const (
	EncryptionScheme_SA_NONE        EncryptionScheme = 0
	EncryptionScheme_SA_AES_256_CFB EncryptionScheme = 1
)

var EncryptionScheme_name = map[int32]string{
	0: "SA_NONE",
	1: "SA_AES_256_CFB",
}

var EncryptionScheme_value = map[string]int32{
	"SA_NONE":        0,
	"SA_AES_256_CFB": 1,
}

func (x EncryptionScheme) String() string {
	return proto.EnumName(EncryptionScheme_name, int32(x))
}

func (EncryptionScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{1}
}

type HashAlgorithm int32

const (
	HashAlgorithm_HASH_NONE           HashAlgorithm = 0
	HashAlgorithm_HASH_SHA256_16bytes HashAlgorithm = 1
)

var HashAlgorithm_name = map[int32]string{
	0: "HASH_NONE",
	1: "HASH_SHA256_16bytes",
}

var HashAlgorithm_value = map[string]int32{
	"HASH_NONE":           0,
	"HASH_SHA256_16bytes": 1,
}

func (x HashAlgorithm) String() string {
	return proto.EnumName(HashAlgorithm_name, int32(x))
}

func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{2}
}

// Cipher information to decrypt Sensitive Data
type CipherContext struct {
	// cipher context id, key to this structure
	ContextId string `protobuf:"bytes,1,opt,name=contextId,proto3" json:"contextId,omitempty"`
	// algorithm used to compute hash for certificates
	HashScheme HashAlgorithm `protobuf:"varint,2,opt,name=hashScheme,proto3,enum=HashAlgorithm" json:"hashScheme,omitempty"`
	// for key exchange scheme, like ECDH etc.
	KeyExchangeScheme KeyExchangeScheme `protobuf:"varint,3,opt,name=keyExchangeScheme,proto3,enum=KeyExchangeScheme" json:"keyExchangeScheme,omitempty"`
	// for encrypting sensitive data, like AES256 etc.
	EncryptionScheme EncryptionScheme `protobuf:"varint,4,opt,name=encryptionScheme,proto3,enum=EncryptionScheme" json:"encryptionScheme,omitempty"`
	// device public certificate hash
	DeviceCertHash []byte `protobuf:"bytes,5,opt,name=deviceCertHash,proto3" json:"deviceCertHash,omitempty"`
	// controller certificate hash
	ControllerCertHash   []byte   `protobuf:"bytes,6,opt,name=controllerCertHash,proto3" json:"controllerCertHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherContext) Reset()         { *m = CipherContext{} }
func (m *CipherContext) String() string { return proto.CompactTextString(m) }
func (*CipherContext) ProtoMessage()    {}
func (*CipherContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{0}
}

func (m *CipherContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherContext.Unmarshal(m, b)
}
func (m *CipherContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherContext.Marshal(b, m, deterministic)
}
func (m *CipherContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherContext.Merge(m, src)
}
func (m *CipherContext) XXX_Size() int {
	return xxx_messageInfo_CipherContext.Size(m)
}
func (m *CipherContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherContext.DiscardUnknown(m)
}

var xxx_messageInfo_CipherContext proto.InternalMessageInfo

func (m *CipherContext) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *CipherContext) GetHashScheme() HashAlgorithm {
	if m != nil {
		return m.HashScheme
	}
	return HashAlgorithm_HASH_NONE
}

func (m *CipherContext) GetKeyExchangeScheme() KeyExchangeScheme {
	if m != nil {
		return m.KeyExchangeScheme
	}
	return KeyExchangeScheme_KEA_NONE
}

func (m *CipherContext) GetEncryptionScheme() EncryptionScheme {
	if m != nil {
		return m.EncryptionScheme
	}
	return EncryptionScheme_SA_NONE
}

func (m *CipherContext) GetDeviceCertHash() []byte {
	if m != nil {
		return m.DeviceCertHash
	}
	return nil
}

func (m *CipherContext) GetControllerCertHash() []byte {
	if m != nil {
		return m.ControllerCertHash
	}
	return nil
}

// Encrypted sensitive data information
type CipherBlock struct {
	// cipher context id
	CipherContextId string `protobuf:"bytes,1,opt,name=cipherContextId,proto3" json:"cipherContextId,omitempty"`
	// Initial Value for Symmetric Key derivation
	InitialValue []byte `protobuf:"bytes,2,opt,name=initialValue,proto3" json:"initialValue,omitempty"`
	// encrypted sensitive data
	CipherData []byte `protobuf:"bytes,3,opt,name=cipherData,proto3" json:"cipherData,omitempty"`
	// sha256 of the plaintext sensitive data
	ClearTextSha256      []byte   `protobuf:"bytes,4,opt,name=clearTextSha256,proto3" json:"clearTextSha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipherBlock) Reset()         { *m = CipherBlock{} }
func (m *CipherBlock) String() string { return proto.CompactTextString(m) }
func (*CipherBlock) ProtoMessage()    {}
func (*CipherBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{1}
}

func (m *CipherBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipherBlock.Unmarshal(m, b)
}
func (m *CipherBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipherBlock.Marshal(b, m, deterministic)
}
func (m *CipherBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipherBlock.Merge(m, src)
}
func (m *CipherBlock) XXX_Size() int {
	return xxx_messageInfo_CipherBlock.Size(m)
}
func (m *CipherBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_CipherBlock.DiscardUnknown(m)
}

var xxx_messageInfo_CipherBlock proto.InternalMessageInfo

func (m *CipherBlock) GetCipherContextId() string {
	if m != nil {
		return m.CipherContextId
	}
	return ""
}

func (m *CipherBlock) GetInitialValue() []byte {
	if m != nil {
		return m.InitialValue
	}
	return nil
}

func (m *CipherBlock) GetCipherData() []byte {
	if m != nil {
		return m.CipherData
	}
	return nil
}

func (m *CipherBlock) GetClearTextSha256() []byte {
	if m != nil {
		return m.ClearTextSha256
	}
	return nil
}

// This message will be filled with the
// credential details and encrypted across
//  wire for data in transit, by the controller
// for encryption
type CredentialBlock struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialBlock) Reset()         { *m = CredentialBlock{} }
func (m *CredentialBlock) String() string { return proto.CompactTextString(m) }
func (*CredentialBlock) ProtoMessage()    {}
func (*CredentialBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_d32c1c7154980027, []int{2}
}

func (m *CredentialBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialBlock.Unmarshal(m, b)
}
func (m *CredentialBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialBlock.Marshal(b, m, deterministic)
}
func (m *CredentialBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialBlock.Merge(m, src)
}
func (m *CredentialBlock) XXX_Size() int {
	return xxx_messageInfo_CredentialBlock.Size(m)
}
func (m *CredentialBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialBlock.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialBlock proto.InternalMessageInfo

func (m *CredentialBlock) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *CredentialBlock) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterEnum("KeyExchangeScheme", KeyExchangeScheme_name, KeyExchangeScheme_value)
	proto.RegisterEnum("EncryptionScheme", EncryptionScheme_name, EncryptionScheme_value)
	proto.RegisterEnum("HashAlgorithm", HashAlgorithm_name, HashAlgorithm_value)
	proto.RegisterType((*CipherContext)(nil), "CipherContext")
	proto.RegisterType((*CipherBlock)(nil), "CipherBlock")
	proto.RegisterType((*CredentialBlock)(nil), "CredentialBlock")
}

func init() { proto.RegisterFile("cipherinfo.proto", fileDescriptor_d32c1c7154980027) }

var fileDescriptor_d32c1c7154980027 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x6e, 0xd3, 0x4c,
	0x10, 0x85, 0x7f, 0xf7, 0x87, 0xd2, 0x4c, 0x9d, 0xd4, 0x59, 0x2e, 0x88, 0x10, 0x82, 0x28, 0x42,
	0x28, 0x8a, 0xc4, 0x5a, 0xa4, 0x6a, 0xb8, 0x42, 0xc2, 0x71, 0x0d, 0xa9, 0x2a, 0x15, 0xc9, 0x46,
	0x5c, 0x70, 0x13, 0x6d, 0xd6, 0x13, 0x7b, 0x55, 0xc7, 0x1b, 0xad, 0x37, 0x25, 0xe1, 0x7d, 0x78,
	0x05, 0x9e, 0x0f, 0x79, 0x9d, 0x86, 0xd4, 0xe1, 0x2e, 0xf3, 0xcd, 0x39, 0x93, 0xf1, 0x99, 0x05,
	0x87, 0x8b, 0x65, 0x8a, 0x4a, 0xe4, 0x73, 0x49, 0x97, 0x4a, 0x6a, 0xd9, 0xfb, 0x7d, 0x04, 0x4d,
	0xdf, 0x40, 0x5f, 0xe6, 0x1a, 0xd7, 0x9a, 0xbc, 0x80, 0x06, 0xaf, 0x7e, 0x5e, 0xc5, 0x1d, 0xab,
	0x6b, 0xf5, 0x1b, 0xe1, 0x5f, 0x40, 0x28, 0x40, 0xca, 0x8a, 0x34, 0xe2, 0x29, 0x2e, 0xb0, 0x73,
	0xd4, 0xb5, 0xfa, 0xad, 0x61, 0x8b, 0x96, 0xc8, 0xcb, 0x12, 0xa9, 0x84, 0x4e, 0x17, 0xe1, 0x9e,
	0x82, 0x7c, 0x84, 0xf6, 0x2d, 0x6e, 0x82, 0x35, 0x4f, 0x59, 0x9e, 0xe0, 0xd6, 0xf6, 0xbf, 0xb1,
	0x11, 0x7a, 0x5d, 0xef, 0x84, 0x87, 0x62, 0xf2, 0x01, 0x1c, 0xcc, 0xb9, 0xda, 0x2c, 0xb5, 0x90,
	0xf9, 0x76, 0xc0, 0x23, 0x33, 0xa0, 0x4d, 0x83, 0x5a, 0x23, 0x3c, 0x90, 0x92, 0x37, 0xd0, 0x8a,
	0xf1, 0x4e, 0x70, 0xf4, 0x51, 0xe9, 0x09, 0x2b, 0xd2, 0xce, 0xe3, 0xae, 0xd5, 0xb7, 0xc3, 0x1a,
	0x25, 0x14, 0x48, 0xf9, 0x95, 0x4a, 0x66, 0x19, 0xaa, 0x9d, 0xf6, 0xd8, 0x68, 0xff, 0xd1, 0xe9,
	0xfd, 0xb2, 0xe0, 0xb4, 0x0a, 0x6e, 0x9c, 0x49, 0x7e, 0x4b, 0xfa, 0x70, 0xc6, 0xf7, 0x73, 0xdc,
	0x85, 0x57, 0xc7, 0xa4, 0x07, 0xb6, 0xc8, 0x85, 0x16, 0x2c, 0xfb, 0xc6, 0xb2, 0x55, 0x15, 0xa2,
	0x1d, 0x3e, 0x60, 0xe4, 0x25, 0x40, 0x65, 0xbb, 0x64, 0x9a, 0x99, 0xbc, 0xec, 0x70, 0x8f, 0x98,
	0x7f, 0xcb, 0x90, 0xa9, 0xaf, 0xb8, 0xd6, 0x51, 0xca, 0x86, 0x17, 0x23, 0x93, 0x89, 0x1d, 0xd6,
	0x71, 0xef, 0x0a, 0xce, 0x7c, 0x85, 0x31, 0xe6, 0xe5, 0xf0, 0x6a, 0xd5, 0xe7, 0x70, 0x22, 0x0c,
	0xd0, 0x9b, 0xed, 0x8e, 0xbb, 0xba, 0xec, 0x2d, 0x59, 0x51, 0xfc, 0x90, 0x2a, 0x36, 0x8b, 0x35,
	0xc2, 0x5d, 0x3d, 0x70, 0xa1, 0x7d, 0x70, 0x31, 0x62, 0xc3, 0xc9, 0x75, 0xe0, 0x4d, 0x6f, 0xbe,
	0xdc, 0x04, 0xce, 0x7f, 0xf7, 0x55, 0xe0, 0x5f, 0x4e, 0x1c, 0x6b, 0x70, 0x0e, 0x4e, 0xfd, 0x42,
	0xe4, 0x14, 0x9e, 0x44, 0x3b, 0x39, 0x81, 0x56, 0xe4, 0x4d, 0xbd, 0x20, 0x9a, 0x0e, 0x2f, 0x46,
	0x53, 0xff, 0xd3, 0xd8, 0xb1, 0x06, 0xef, 0xa1, 0xf9, 0xe0, 0x39, 0x91, 0x26, 0x34, 0x26, 0x5e,
	0x34, 0xb9, 0xf7, 0x3c, 0x83, 0xa7, 0xa6, 0x8c, 0x26, 0x5e, 0x69, 0x7a, 0x37, 0x9a, 0x6d, 0x34,
	0x16, 0x8e, 0x35, 0xfe, 0x0c, 0xaf, 0xb8, 0x5c, 0xd0, 0x9f, 0x18, 0x63, 0xcc, 0x28, 0xcf, 0xe4,
	0x2a, 0xa6, 0xab, 0x02, 0x55, 0x79, 0xe6, 0xea, 0xb5, 0x7f, 0x7f, 0x9d, 0x08, 0x9d, 0xae, 0x66,
	0x94, 0xcb, 0x85, 0x9b, 0xcd, 0xdf, 0x62, 0x9c, 0xa0, 0x8b, 0x77, 0xe8, 0xb2, 0xa5, 0x70, 0x13,
	0xe9, 0x72, 0x99, 0xcf, 0x45, 0x32, 0x3b, 0x36, 0xe2, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0xcd, 0x82, 0x94, 0x2e, 0x03, 0x00, 0x00,
}