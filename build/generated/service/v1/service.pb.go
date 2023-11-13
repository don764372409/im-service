// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0--rc2
// source: service.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResultDTOCode int32

const (
	ResultDTOCode_DEFAULT ResultDTOCode = 0
	ResultDTOCode_SUCCESS ResultDTOCode = 200
	ResultDTOCode_ERROR   ResultDTOCode = 500
	// * 跳转2级密码输入页
	ResultDTOCode_TO_INPUT_PWD2 ResultDTOCode = 2000
)

// Enum value maps for ResultDTOCode.
var (
	ResultDTOCode_name = map[int32]string{
		0:    "DEFAULT",
		200:  "SUCCESS",
		500:  "ERROR",
		2000: "TO_INPUT_PWD2",
	}
	ResultDTOCode_value = map[string]int32{
		"DEFAULT":       0,
		"SUCCESS":       200,
		"ERROR":         500,
		"TO_INPUT_PWD2": 2000,
	}
)

func (x ResultDTOCode) Enum() *ResultDTOCode {
	p := new(ResultDTOCode)
	*p = x
	return p
}

func (x ResultDTOCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultDTOCode) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[0].Descriptor()
}

func (ResultDTOCode) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[0]
}

func (x ResultDTOCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultDTOCode.Descriptor instead.
func (ResultDTOCode) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type ConfigReq_LogSwitch int32

const (
	ConfigReq_LogSwitchUNKNOWN ConfigReq_LogSwitch = 0
	ConfigReq_CLOSE            ConfigReq_LogSwitch = 2 //关闭日志
	ConfigReq_CONSOLE          ConfigReq_LogSwitch = 3 //控制台日志
	ConfigReq_FILE             ConfigReq_LogSwitch = 4 //文件日志
	ConfigReq_CONSOLE_FILE     ConfigReq_LogSwitch = 5 //控制台+文件
)

// Enum value maps for ConfigReq_LogSwitch.
var (
	ConfigReq_LogSwitch_name = map[int32]string{
		0: "LogSwitchUNKNOWN",
		2: "CLOSE",
		3: "CONSOLE",
		4: "FILE",
		5: "CONSOLE_FILE",
	}
	ConfigReq_LogSwitch_value = map[string]int32{
		"LogSwitchUNKNOWN": 0,
		"CLOSE":            2,
		"CONSOLE":          3,
		"FILE":             4,
		"CONSOLE_FILE":     5,
	}
)

func (x ConfigReq_LogSwitch) Enum() *ConfigReq_LogSwitch {
	p := new(ConfigReq_LogSwitch)
	*p = x
	return p
}

func (x ConfigReq_LogSwitch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigReq_LogSwitch) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[1].Descriptor()
}

func (ConfigReq_LogSwitch) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[1]
}

func (x ConfigReq_LogSwitch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigReq_LogSwitch.Descriptor instead.
func (ConfigReq_LogSwitch) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9, 0}
}

type ConfigReq_DeviceType int32

const (
	ConfigReq_Unknown ConfigReq_DeviceType = 0
	ConfigReq_PC      ConfigReq_DeviceType = 1 // 电脑端
	ConfigReq_Android ConfigReq_DeviceType = 2 // 手机端
	ConfigReq_IOS     ConfigReq_DeviceType = 3 // IOS
	ConfigReq_H5      ConfigReq_DeviceType = 4 // H5
)

// Enum value maps for ConfigReq_DeviceType.
var (
	ConfigReq_DeviceType_name = map[int32]string{
		0: "Unknown",
		1: "PC",
		2: "Android",
		3: "IOS",
		4: "H5",
	}
	ConfigReq_DeviceType_value = map[string]int32{
		"Unknown": 0,
		"PC":      1,
		"Android": 2,
		"IOS":     3,
		"H5":      4,
	}
)

func (x ConfigReq_DeviceType) Enum() *ConfigReq_DeviceType {
	p := new(ConfigReq_DeviceType)
	*p = x
	return p
}

func (x ConfigReq_DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigReq_DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_proto_enumTypes[2].Descriptor()
}

func (ConfigReq_DeviceType) Type() protoreflect.EnumType {
	return &file_service_proto_enumTypes[2]
}

func (x ConfigReq_DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigReq_DeviceType.Descriptor instead.
func (ConfigReq_DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9, 1}
}

type MsgPageDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`      //聊天类型 friend、group
	Target uint64 `protobuf:"varint,2,opt,name=target,proto3" json:"target,omitempty"` //聊天目标 用户ID或群ID
	Time   uint64 `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`     //消息发送时间 查询此时间之前的消息
}

func (x *MsgPageDTO) Reset() {
	*x = MsgPageDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPageDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPageDTO) ProtoMessage() {}

func (x *MsgPageDTO) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPageDTO.ProtoReflect.Descriptor instead.
func (*MsgPageDTO) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *MsgPageDTO) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MsgPageDTO) GetTarget() uint64 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *MsgPageDTO) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

// * 客户端发送消息状态回调
type SendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	No   string `protobuf:"bytes,1,opt,name=no,proto3" json:"no,omitempty"`      //消息ID
	Send int32  `protobuf:"varint,2,opt,name=send,proto3" json:"send,omitempty"` //发送状态 1-发送中 2-发送成功 -1-发送失败
}

func (x *SendResp) Reset() {
	*x = SendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResp) ProtoMessage() {}

func (x *SendResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResp.ProtoReflect.Descriptor instead.
func (*SendResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendResp) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *SendResp) GetSend() int32 {
	if x != nil {
		return x.Send
	}
	return 0
}

type ChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`       // 聊天类型 friend, group
	Target  uint64 `protobuf:"varint,2,opt,name=target,proto3" json:"target,omitempty"`  // 聊天目标 用户ID或群ID
	No      string `protobuf:"bytes,3,opt,name=no,proto3" json:"no,omitempty"`           //消息ID 客户端通过UUID生成 发送成功失败时根据此ID获取消息客户端并修改状态
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"` // 聊天内容 仅发生消息时传
}

func (x *ChatReq) Reset() {
	*x = ChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatReq) ProtoMessage() {}

func (x *ChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatReq.ProtoReflect.Descriptor instead.
func (*ChatReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *ChatReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ChatReq) GetTarget() uint64 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *ChatReq) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *ChatReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// * 好友详情页根据ID获取好友数据
type FriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    //好友ID
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` //好友备注 查询/删除时留空，修改好友备注时传入
}

func (x *FriendReq) Reset() {
	*x = FriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendReq) ProtoMessage() {}

func (x *FriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendReq.ProtoReflect.Descriptor instead.
func (*FriendReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *FriendReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FriendReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// * 搜索用户、群聊得到的请求参数
type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"` // 搜索关键字
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *SearchReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// * 好友请求参数
type FriendApplyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`        //添加好友时是用户ID 同意或拒绝好友时是新朋友记录ID
	Remark string `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"` //备注 同意或拒绝时留空
	State  int32  `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`  //-1拒绝 2同意 操作时传，添加时留空
}

func (x *FriendApplyReq) Reset() {
	*x = FriendApplyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendApplyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendApplyReq) ProtoMessage() {}

func (x *FriendApplyReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendApplyReq.ProtoReflect.Descriptor instead.
func (*FriendApplyReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *FriendApplyReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FriendApplyReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *FriendApplyReq) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

// * 修改用户信息的请求参数
type UpdateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    //用户ID
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` //可以是昵称、签名、邮箱、头像
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type UploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"` // 文件路径 包含路径和文件名 列如 C:\\Users\\Administrator\\Desktop\\result.png
}

func (x *UploadReq) Reset() {
	*x = UploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReq) ProtoMessage() {}

func (x *UploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReq.ProtoReflect.Descriptor instead.
func (*UploadReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *UploadReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// * 注册
type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *UserReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseDir    string               `protobuf:"bytes,1,opt,name=baseDir,proto3" json:"baseDir,omitempty"`                                  //配置根目录
	DeviceType ConfigReq_DeviceType `protobuf:"varint,2,opt,name=deviceType,proto3,enum=ConfigReq_DeviceType" json:"deviceType,omitempty"` //设备类型
	LogSwitch  ConfigReq_LogSwitch  `protobuf:"varint,3,opt,name=logSwitch,proto3,enum=ConfigReq_LogSwitch" json:"logSwitch,omitempty"`    //日志开关
	ApiHost    string               `protobuf:"bytes,4,opt,name=ApiHost,proto3" json:"ApiHost,omitempty"`
	WsHost     string               `protobuf:"bytes,5,opt,name=WsHost,proto3" json:"WsHost,omitempty"`
}

func (x *ConfigReq) Reset() {
	*x = ConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigReq) ProtoMessage() {}

func (x *ConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigReq.ProtoReflect.Descriptor instead.
func (*ConfigReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9}
}

func (x *ConfigReq) GetBaseDir() string {
	if x != nil {
		return x.BaseDir
	}
	return ""
}

func (x *ConfigReq) GetDeviceType() ConfigReq_DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return ConfigReq_Unknown
}

func (x *ConfigReq) GetLogSwitch() ConfigReq_LogSwitch {
	if x != nil {
		return x.LogSwitch
	}
	return ConfigReq_LogSwitchUNKNOWN
}

func (x *ConfigReq) GetApiHost() string {
	if x != nil {
		return x.ApiHost
	}
	return ""
}

func (x *ConfigReq) GetWsHost() string {
	if x != nil {
		return x.WsHost
	}
	return ""
}

type ResultDTOResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ResultDTOResp) Reset() {
	*x = ResultDTOResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultDTOResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultDTOResp) ProtoMessage() {}

func (x *ResultDTOResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultDTOResp.ProtoReflect.Descriptor instead.
func (*ResultDTOResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{10}
}

func (x *ResultDTOResp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResultDTOResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ResultDTOResp) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4c, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x50, 0x61, 0x67, 0x65, 0x44, 0x54, 0x4f, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x2e, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x22, 0x5f, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2f,
	0x0a, 0x09, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x25, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4e, 0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x33, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1f, 0x0a, 0x09, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x51, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0xda, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x73, 0x65, 0x44, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x73, 0x65, 0x44, 0x69, 0x72, 0x12, 0x35, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x2e, 0x4c, 0x6f,
	0x67, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x69, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x57, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x73,
	0x48, 0x6f, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x53, 0x4f, 0x4c, 0x45, 0x10, 0x03, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e,
	0x53, 0x4f, 0x4c, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x05, 0x22, 0x3f, 0x0a, 0x0a, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x43, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x49,
	0x4f, 0x53, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x35, 0x10, 0x04, 0x22, 0x49, 0x0a, 0x0d,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x2a, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x44, 0x54, 0x4f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0xc8, 0x01, 0x12, 0x0a, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf4, 0x03, 0x12,
	0x12, 0x0a, 0x0d, 0x54, 0x4f, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x50, 0x57, 0x44, 0x32,
	0x10, 0xd0, 0x0f, 0x42, 0x42, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x62, 0x65, 0x61, 0x6e, 0x42, 0x0f, 0x42, 0x74,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0xa2,
	0x02, 0x05, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_proto_goTypes = []interface{}{
	(ResultDTOCode)(0),        // 0: ResultDTOCode
	(ConfigReq_LogSwitch)(0),  // 1: ConfigReq.LogSwitch
	(ConfigReq_DeviceType)(0), // 2: ConfigReq.DeviceType
	(*MsgPageDTO)(nil),        // 3: MsgPageDTO
	(*SendResp)(nil),          // 4: SendResp
	(*ChatReq)(nil),           // 5: ChatReq
	(*FriendReq)(nil),         // 6: FriendReq
	(*SearchReq)(nil),         // 7: SearchReq
	(*FriendApplyReq)(nil),    // 8: FriendApplyReq
	(*UpdateUserReq)(nil),     // 9: UpdateUserReq
	(*UploadReq)(nil),         // 10: UploadReq
	(*UserReq)(nil),           // 11: UserReq
	(*ConfigReq)(nil),         // 12: ConfigReq
	(*ResultDTOResp)(nil),     // 13: ResultDTOResp
}
var file_service_proto_depIdxs = []int32{
	2, // 0: ConfigReq.deviceType:type_name -> ConfigReq.DeviceType
	1, // 1: ConfigReq.logSwitch:type_name -> ConfigReq.LogSwitch
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPageDTO); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendApplyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultDTOResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		EnumInfos:         file_service_proto_enumTypes,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
