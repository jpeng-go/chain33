// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

package types

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

type Genesis struct {
	Isrun                bool     `protobuf:"varint,1,opt,name=isrun,proto3" json:"isrun,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0}
}

func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genesis.Unmarshal(m, b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return xxx_messageInfo_Genesis.Size(m)
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetIsrun() bool {
	if m != nil {
		return m.Isrun
	}
	return false
}

type ExecTxList struct {
	StateHash            []byte         `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	ParentHash           []byte         `protobuf:"bytes,7,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	MainHash             []byte         `protobuf:"bytes,8,opt,name=mainHash,proto3" json:"mainHash,omitempty"`
	MainHeight           int64          `protobuf:"varint,9,opt,name=mainHeight,proto3" json:"mainHeight,omitempty"`
	BlockTime            int64          `protobuf:"varint,3,opt,name=blockTime,proto3" json:"blockTime,omitempty"`
	Height               int64          `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Difficulty           uint64         `protobuf:"varint,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	IsMempool            bool           `protobuf:"varint,6,opt,name=isMempool,proto3" json:"isMempool,omitempty"`
	Txs                  []*Transaction `protobuf:"bytes,2,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecTxList) Reset()         { *m = ExecTxList{} }
func (m *ExecTxList) String() string { return proto.CompactTextString(m) }
func (*ExecTxList) ProtoMessage()    {}
func (*ExecTxList) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{1}
}

func (m *ExecTxList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecTxList.Unmarshal(m, b)
}
func (m *ExecTxList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecTxList.Marshal(b, m, deterministic)
}
func (m *ExecTxList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecTxList.Merge(m, src)
}
func (m *ExecTxList) XXX_Size() int {
	return xxx_messageInfo_ExecTxList.Size(m)
}
func (m *ExecTxList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecTxList.DiscardUnknown(m)
}

var xxx_messageInfo_ExecTxList proto.InternalMessageInfo

func (m *ExecTxList) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ExecTxList) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *ExecTxList) GetMainHash() []byte {
	if m != nil {
		return m.MainHash
	}
	return nil
}

func (m *ExecTxList) GetMainHeight() int64 {
	if m != nil {
		return m.MainHeight
	}
	return 0
}

func (m *ExecTxList) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *ExecTxList) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ExecTxList) GetDifficulty() uint64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *ExecTxList) GetIsMempool() bool {
	if m != nil {
		return m.IsMempool
	}
	return false
}

func (m *ExecTxList) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Query struct {
	Execer               []byte   `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	FuncName             string   `protobuf:"bytes,2,opt,name=funcName,proto3" json:"funcName,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{2}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *Query) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *Query) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CreateTxIn struct {
	Execer               []byte   `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	ActionName           string   `protobuf:"bytes,2,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTxIn) Reset()         { *m = CreateTxIn{} }
func (m *CreateTxIn) String() string { return proto.CompactTextString(m) }
func (*CreateTxIn) ProtoMessage()    {}
func (*CreateTxIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{3}
}

func (m *CreateTxIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTxIn.Unmarshal(m, b)
}
func (m *CreateTxIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTxIn.Marshal(b, m, deterministic)
}
func (m *CreateTxIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTxIn.Merge(m, src)
}
func (m *CreateTxIn) XXX_Size() int {
	return xxx_messageInfo_CreateTxIn.Size(m)
}
func (m *CreateTxIn) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTxIn.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTxIn proto.InternalMessageInfo

func (m *CreateTxIn) GetExecer() []byte {
	if m != nil {
		return m.Execer
	}
	return nil
}

func (m *CreateTxIn) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *CreateTxIn) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// 配置修改部分
type ArrayConfig struct {
	Value                []string `protobuf:"bytes,3,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrayConfig) Reset()         { *m = ArrayConfig{} }
func (m *ArrayConfig) String() string { return proto.CompactTextString(m) }
func (*ArrayConfig) ProtoMessage()    {}
func (*ArrayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{4}
}

func (m *ArrayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayConfig.Unmarshal(m, b)
}
func (m *ArrayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayConfig.Marshal(b, m, deterministic)
}
func (m *ArrayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayConfig.Merge(m, src)
}
func (m *ArrayConfig) XXX_Size() int {
	return xxx_messageInfo_ArrayConfig.Size(m)
}
func (m *ArrayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayConfig proto.InternalMessageInfo

func (m *ArrayConfig) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type StringConfig struct {
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringConfig) Reset()         { *m = StringConfig{} }
func (m *StringConfig) String() string { return proto.CompactTextString(m) }
func (*StringConfig) ProtoMessage()    {}
func (*StringConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{5}
}

func (m *StringConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringConfig.Unmarshal(m, b)
}
func (m *StringConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringConfig.Marshal(b, m, deterministic)
}
func (m *StringConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringConfig.Merge(m, src)
}
func (m *StringConfig) XXX_Size() int {
	return xxx_messageInfo_StringConfig.Size(m)
}
func (m *StringConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StringConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StringConfig proto.InternalMessageInfo

func (m *StringConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Int32Config struct {
	Value                int32    `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32Config) Reset()         { *m = Int32Config{} }
func (m *Int32Config) String() string { return proto.CompactTextString(m) }
func (*Int32Config) ProtoMessage()    {}
func (*Int32Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{6}
}

func (m *Int32Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32Config.Unmarshal(m, b)
}
func (m *Int32Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32Config.Marshal(b, m, deterministic)
}
func (m *Int32Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32Config.Merge(m, src)
}
func (m *Int32Config) XXX_Size() int {
	return xxx_messageInfo_Int32Config.Size(m)
}
func (m *Int32Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32Config.DiscardUnknown(m)
}

var xxx_messageInfo_Int32Config proto.InternalMessageInfo

func (m *Int32Config) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ConfigItem struct {
	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*ConfigItem_Arr
	//	*ConfigItem_Str
	//	*ConfigItem_Int
	Value                isConfigItem_Value `protobuf_oneof:"value"`
	Ty                   int32              `protobuf:"varint,11,opt,name=Ty,proto3" json:"Ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ConfigItem) Reset()         { *m = ConfigItem{} }
func (m *ConfigItem) String() string { return proto.CompactTextString(m) }
func (*ConfigItem) ProtoMessage()    {}
func (*ConfigItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{7}
}

func (m *ConfigItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigItem.Unmarshal(m, b)
}
func (m *ConfigItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigItem.Marshal(b, m, deterministic)
}
func (m *ConfigItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigItem.Merge(m, src)
}
func (m *ConfigItem) XXX_Size() int {
	return xxx_messageInfo_ConfigItem.Size(m)
}
func (m *ConfigItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigItem.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigItem proto.InternalMessageInfo

func (m *ConfigItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigItem) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type isConfigItem_Value interface {
	isConfigItem_Value()
}

type ConfigItem_Arr struct {
	Arr *ArrayConfig `protobuf:"bytes,3,opt,name=arr,proto3,oneof"`
}

type ConfigItem_Str struct {
	Str *StringConfig `protobuf:"bytes,4,opt,name=str,proto3,oneof"`
}

type ConfigItem_Int struct {
	Int *Int32Config `protobuf:"bytes,5,opt,name=int,proto3,oneof"`
}

func (*ConfigItem_Arr) isConfigItem_Value() {}

func (*ConfigItem_Str) isConfigItem_Value() {}

func (*ConfigItem_Int) isConfigItem_Value() {}

func (m *ConfigItem) GetValue() isConfigItem_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ConfigItem) GetArr() *ArrayConfig {
	if x, ok := m.GetValue().(*ConfigItem_Arr); ok {
		return x.Arr
	}
	return nil
}

func (m *ConfigItem) GetStr() *StringConfig {
	if x, ok := m.GetValue().(*ConfigItem_Str); ok {
		return x.Str
	}
	return nil
}

func (m *ConfigItem) GetInt() *Int32Config {
	if x, ok := m.GetValue().(*ConfigItem_Int); ok {
		return x.Int
	}
	return nil
}

func (m *ConfigItem) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConfigItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConfigItem_Arr)(nil),
		(*ConfigItem_Str)(nil),
		(*ConfigItem_Int)(nil),
	}
}

type ModifyConfig struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Op                   string   `protobuf:"bytes,3,opt,name=op,proto3" json:"op,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyConfig) Reset()         { *m = ModifyConfig{} }
func (m *ModifyConfig) String() string { return proto.CompactTextString(m) }
func (*ModifyConfig) ProtoMessage()    {}
func (*ModifyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{8}
}

func (m *ModifyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyConfig.Unmarshal(m, b)
}
func (m *ModifyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyConfig.Marshal(b, m, deterministic)
}
func (m *ModifyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyConfig.Merge(m, src)
}
func (m *ModifyConfig) XXX_Size() int {
	return xxx_messageInfo_ModifyConfig.Size(m)
}
func (m *ModifyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyConfig proto.InternalMessageInfo

func (m *ModifyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ModifyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ModifyConfig) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *ModifyConfig) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ReceiptConfig struct {
	Prev                 *ConfigItem `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	Current              *ConfigItem `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReceiptConfig) Reset()         { *m = ReceiptConfig{} }
func (m *ReceiptConfig) String() string { return proto.CompactTextString(m) }
func (*ReceiptConfig) ProtoMessage()    {}
func (*ReceiptConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{9}
}

func (m *ReceiptConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptConfig.Unmarshal(m, b)
}
func (m *ReceiptConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptConfig.Marshal(b, m, deterministic)
}
func (m *ReceiptConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptConfig.Merge(m, src)
}
func (m *ReceiptConfig) XXX_Size() int {
	return xxx_messageInfo_ReceiptConfig.Size(m)
}
func (m *ReceiptConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptConfig proto.InternalMessageInfo

func (m *ReceiptConfig) GetPrev() *ConfigItem {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptConfig) GetCurrent() *ConfigItem {
	if m != nil {
		return m.Current
	}
	return nil
}

type ReplyConfig struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyConfig) Reset()         { *m = ReplyConfig{} }
func (m *ReplyConfig) String() string { return proto.CompactTextString(m) }
func (*ReplyConfig) ProtoMessage()    {}
func (*ReplyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{10}
}

func (m *ReplyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyConfig.Unmarshal(m, b)
}
func (m *ReplyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyConfig.Marshal(b, m, deterministic)
}
func (m *ReplyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyConfig.Merge(m, src)
}
func (m *ReplyConfig) XXX_Size() int {
	return xxx_messageInfo_ReplyConfig.Size(m)
}
func (m *ReplyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyConfig proto.InternalMessageInfo

func (m *ReplyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HistoryCertStore struct {
	Rootcerts            [][]byte `protobuf:"bytes,1,rep,name=rootcerts,proto3" json:"rootcerts,omitempty"`
	IntermediateCerts    [][]byte `protobuf:"bytes,2,rep,name=intermediateCerts,proto3" json:"intermediateCerts,omitempty"`
	RevocationList       [][]byte `protobuf:"bytes,3,rep,name=revocationList,proto3" json:"revocationList,omitempty"`
	CurHeigth            int64    `protobuf:"varint,4,opt,name=curHeigth,proto3" json:"curHeigth,omitempty"`
	NxtHeight            int64    `protobuf:"varint,5,opt,name=nxtHeight,proto3" json:"nxtHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HistoryCertStore) Reset()         { *m = HistoryCertStore{} }
func (m *HistoryCertStore) String() string { return proto.CompactTextString(m) }
func (*HistoryCertStore) ProtoMessage()    {}
func (*HistoryCertStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{11}
}

func (m *HistoryCertStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryCertStore.Unmarshal(m, b)
}
func (m *HistoryCertStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryCertStore.Marshal(b, m, deterministic)
}
func (m *HistoryCertStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryCertStore.Merge(m, src)
}
func (m *HistoryCertStore) XXX_Size() int {
	return xxx_messageInfo_HistoryCertStore.Size(m)
}
func (m *HistoryCertStore) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryCertStore.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryCertStore proto.InternalMessageInfo

func (m *HistoryCertStore) GetRootcerts() [][]byte {
	if m != nil {
		return m.Rootcerts
	}
	return nil
}

func (m *HistoryCertStore) GetIntermediateCerts() [][]byte {
	if m != nil {
		return m.IntermediateCerts
	}
	return nil
}

func (m *HistoryCertStore) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *HistoryCertStore) GetCurHeigth() int64 {
	if m != nil {
		return m.CurHeigth
	}
	return 0
}

func (m *HistoryCertStore) GetNxtHeight() int64 {
	if m != nil {
		return m.NxtHeight
	}
	return 0
}

type AuthorityCfg struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	CryptoPath           string   `protobuf:"bytes,2,opt,name=cryptoPath,proto3" json:"cryptoPath,omitempty"`
	SignType             string   `protobuf:"bytes,3,opt,name=signType,proto3" json:"signType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorityCfg) Reset()         { *m = AuthorityCfg{} }
func (m *AuthorityCfg) String() string { return proto.CompactTextString(m) }
func (*AuthorityCfg) ProtoMessage()    {}
func (*AuthorityCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{12}
}

func (m *AuthorityCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorityCfg.Unmarshal(m, b)
}
func (m *AuthorityCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorityCfg.Marshal(b, m, deterministic)
}
func (m *AuthorityCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorityCfg.Merge(m, src)
}
func (m *AuthorityCfg) XXX_Size() int {
	return xxx_messageInfo_AuthorityCfg.Size(m)
}
func (m *AuthorityCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorityCfg.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorityCfg proto.InternalMessageInfo

func (m *AuthorityCfg) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *AuthorityCfg) GetCryptoPath() string {
	if m != nil {
		return m.CryptoPath
	}
	return ""
}

func (m *AuthorityCfg) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

func init() {
	proto.RegisterType((*Genesis)(nil), "types.Genesis")
	proto.RegisterType((*ExecTxList)(nil), "types.ExecTxList")
	proto.RegisterType((*Query)(nil), "types.Query")
	proto.RegisterType((*CreateTxIn)(nil), "types.CreateTxIn")
	proto.RegisterType((*ArrayConfig)(nil), "types.ArrayConfig")
	proto.RegisterType((*StringConfig)(nil), "types.StringConfig")
	proto.RegisterType((*Int32Config)(nil), "types.Int32Config")
	proto.RegisterType((*ConfigItem)(nil), "types.ConfigItem")
	proto.RegisterType((*ModifyConfig)(nil), "types.ModifyConfig")
	proto.RegisterType((*ReceiptConfig)(nil), "types.ReceiptConfig")
	proto.RegisterType((*ReplyConfig)(nil), "types.ReplyConfig")
	proto.RegisterType((*HistoryCertStore)(nil), "types.HistoryCertStore")
	proto.RegisterType((*AuthorityCfg)(nil), "types.AuthorityCfg")
}

func init() {
	proto.RegisterFile("executor.proto", fileDescriptor_12d1cdcda51e000f)
}

var fileDescriptor_12d1cdcda51e000f = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xe1, 0x6a, 0xdb, 0x48,
	0x10, 0x3e, 0x49, 0x76, 0x1c, 0x8f, 0x7d, 0x21, 0xd9, 0x3b, 0x0e, 0x11, 0xee, 0x12, 0xa3, 0xe4,
	0x72, 0x86, 0x3b, 0x1c, 0xb0, 0xb9, 0x07, 0x48, 0x4c, 0x69, 0x02, 0x4d, 0x69, 0x15, 0xf7, 0x4f,
	0x7e, 0x14, 0xd6, 0xeb, 0xb1, 0xb5, 0xc4, 0xde, 0x15, 0xab, 0x51, 0xb0, 0xde, 0xa6, 0xcf, 0x52,
	0xfa, 0x60, 0x65, 0xd7, 0xb2, 0x25, 0x9a, 0xb4, 0xd0, 0x7f, 0x9a, 0xef, 0xfb, 0x34, 0x3b, 0xdf,
	0xec, 0xce, 0xc0, 0x01, 0xae, 0x51, 0xe4, 0xa4, 0xcd, 0x20, 0x35, 0x9a, 0x34, 0x6b, 0x52, 0x91,
	0x62, 0x76, 0x7c, 0x44, 0x86, 0xab, 0x8c, 0x0b, 0x92, 0x5a, 0x6d, 0x98, 0xe8, 0x14, 0x5a, 0xaf,
	0x51, 0x61, 0x26, 0x33, 0xf6, 0x3b, 0x34, 0x65, 0x66, 0x72, 0x15, 0x7a, 0x3d, 0xaf, 0xbf, 0x1f,
	0x6f, 0x82, 0xe8, 0x93, 0x0f, 0xf0, 0x6a, 0x8d, 0x62, 0xb2, 0x7e, 0x23, 0x33, 0x62, 0x7f, 0x42,
	0x3b, 0x23, 0x4e, 0x78, 0xc3, 0xb3, 0xc4, 0x09, 0xbb, 0x71, 0x05, 0xb0, 0x13, 0x80, 0x94, 0x1b,
	0x54, 0xe4, 0xe8, 0x96, 0xa3, 0x6b, 0x08, 0x3b, 0x86, 0xfd, 0x15, 0x97, 0xca, 0xb1, 0xfb, 0x8e,
	0xdd, 0xc5, 0xf6, 0x5f, 0xf7, 0x8d, 0x72, 0x91, 0x50, 0xd8, 0xee, 0x79, 0xfd, 0x20, 0xae, 0x21,
	0xf6, 0xe4, 0xe9, 0x52, 0x8b, 0xc7, 0x89, 0x5c, 0x61, 0x18, 0x38, 0xba, 0x02, 0xd8, 0x1f, 0xb0,
	0x97, 0x6c, 0xfe, 0x6c, 0x38, 0xaa, 0x8c, 0x6c, 0xd6, 0x99, 0x9c, 0xcf, 0xa5, 0xc8, 0x97, 0x54,
	0x84, 0xcd, 0x9e, 0xd7, 0x6f, 0xc4, 0x35, 0xc4, 0x66, 0x95, 0xd9, 0x1d, 0xae, 0x52, 0xad, 0x97,
	0xe1, 0x9e, 0x33, 0x5e, 0x01, 0xec, 0x1c, 0x02, 0x5a, 0x67, 0xa1, 0xdf, 0x0b, 0xfa, 0x9d, 0x21,
	0x1b, 0xb8, 0x2e, 0x0e, 0x26, 0x55, 0x13, 0x63, 0x4b, 0x47, 0x1f, 0xa0, 0xf9, 0x3e, 0x47, 0x53,
	0xd8, 0x22, 0x6c, 0xe3, 0xd1, 0x94, 0x9d, 0x29, 0x23, 0x6b, 0x7b, 0x9e, 0x2b, 0xf1, 0x96, 0xaf,
	0x30, 0xf4, 0x7b, 0x5e, 0xbf, 0x1d, 0xef, 0x62, 0x16, 0x42, 0x2b, 0xe5, 0xc5, 0x52, 0xf3, 0x99,
	0x33, 0xd5, 0x8d, 0xb7, 0x61, 0xf4, 0x11, 0x60, 0x6c, 0x90, 0x13, 0x4e, 0xd6, 0xb7, 0xea, 0xbb,
	0xb9, 0x4f, 0x00, 0x36, 0xb5, 0xd4, 0xb2, 0xd7, 0x90, 0x1f, 0xe4, 0x3f, 0x83, 0xce, 0x95, 0x31,
	0xbc, 0x18, 0x6b, 0x35, 0x97, 0x0b, 0x7b, 0xfd, 0x4f, 0x7c, 0x99, 0xdb, 0xde, 0x06, 0xfd, 0x76,
	0xbc, 0x09, 0xa2, 0x73, 0xe8, 0xde, 0x93, 0x91, 0x6a, 0xf1, 0x5c, 0xe5, 0x55, 0xaa, 0x33, 0xe8,
	0xdc, 0x2a, 0x1a, 0x0d, 0x5f, 0x12, 0x35, 0xb7, 0xa2, 0x2f, 0x1e, 0xc0, 0x46, 0x70, 0x4b, 0xb8,
	0x62, 0x87, 0x10, 0x3c, 0x62, 0xe1, 0xdc, 0xb4, 0x63, 0xfb, 0xc9, 0x18, 0x34, 0xf8, 0x6c, 0x66,
	0x4a, 0x13, 0xee, 0x9b, 0x5d, 0x40, 0xc0, 0x8d, 0x71, 0x89, 0xaa, 0x1b, 0xa8, 0x95, 0x7d, 0xf3,
	0x4b, 0x6c, 0x05, 0xec, 0x1f, 0x08, 0x32, 0x32, 0xee, 0xf2, 0x3b, 0xc3, 0xdf, 0x4a, 0x5d, 0xbd,
	0x72, 0x2b, 0xcc, 0xc8, 0x25, 0x94, 0x8a, 0xdc, 0x4b, 0xa8, 0x12, 0xd6, 0x8a, 0xb7, 0x3a, 0xa9,
	0x88, 0x1d, 0x80, 0x3f, 0x29, 0xc2, 0x8e, 0x33, 0xe0, 0x4f, 0x8a, 0xeb, 0x56, 0xe9, 0x29, 0x7a,
	0x80, 0xee, 0x9d, 0x9e, 0xc9, 0xf9, 0xb6, 0x6f, 0xcf, 0x7d, 0xec, 0xec, 0xfb, 0xb5, 0x1e, 0xd9,
	0x84, 0x3a, 0x2d, 0xdb, 0xe6, 0xeb, 0x74, 0xe7, 0xb6, 0x51, 0xb9, 0x8d, 0x04, 0xfc, 0x1a, 0xa3,
	0x40, 0x99, 0x52, 0x99, 0xfc, 0x6f, 0x68, 0xa4, 0x06, 0x9f, 0x5c, 0xf6, 0xce, 0xf0, 0xa8, 0x2c,
	0xb7, 0xea, 0x62, 0xec, 0x68, 0xf6, 0x2f, 0xb4, 0x44, 0x6e, 0xec, 0x98, 0xb9, 0x33, 0x5f, 0x54,
	0x6e, 0x15, 0xd1, 0xff, 0xd0, 0x89, 0x31, 0x5d, 0xfe, 0x64, 0xfd, 0xd1, 0x67, 0x0f, 0x0e, 0x6f,
	0x64, 0x46, 0xda, 0x14, 0x63, 0x34, 0x74, 0x4f, 0xda, 0xa0, 0x1d, 0x1f, 0xa3, 0x35, 0x09, 0x34,
	0x94, 0x85, 0x5e, 0x2f, 0xb0, 0xeb, 0x60, 0x07, 0xb0, 0xff, 0xe0, 0x48, 0x2a, 0x42, 0xb3, 0xc2,
	0x99, 0xe4, 0x84, 0x63, 0xa7, 0xf2, 0x9d, 0xea, 0x39, 0xc1, 0x2e, 0xe0, 0xc0, 0xe0, 0x93, 0x16,
	0xdc, 0xbe, 0x5d, 0xbb, 0x6c, 0xdc, 0x4b, 0xec, 0xc6, 0xdf, 0xa0, 0xf6, 0x4c, 0x91, 0x1b, 0xbb,
	0x15, 0x28, 0x29, 0xa7, 0xbd, 0x02, 0x2c, 0xab, 0xd6, 0x54, 0x6e, 0x91, 0xe6, 0x86, 0xdd, 0x01,
	0xd1, 0x14, 0xba, 0x57, 0x39, 0x25, 0xda, 0x48, 0x2a, 0xc6, 0xf3, 0x85, 0x9b, 0x2a, 0xc5, 0xa7,
	0x4b, 0x2c, 0x97, 0x5e, 0x19, 0xd9, 0xa9, 0x12, 0xa6, 0x48, 0x49, 0xbf, 0xe3, 0x94, 0x6c, 0xa7,
	0xaa, 0x42, 0xec, 0x44, 0x67, 0x72, 0xa1, 0x26, 0x45, 0xba, 0x9d, 0x84, 0x5d, 0x7c, 0x7d, 0xfa,
	0xf0, 0xd7, 0x42, 0x52, 0x92, 0x4f, 0x07, 0x42, 0xaf, 0x2e, 0x47, 0x23, 0xa1, 0x2e, 0x45, 0xc2,
	0xa5, 0x1a, 0x8d, 0x2e, 0xdd, 0xa5, 0x4c, 0xf7, 0xdc, 0xea, 0x1d, 0x7d, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xac, 0xc8, 0xee, 0x44, 0xa6, 0x05, 0x00, 0x00,
}
