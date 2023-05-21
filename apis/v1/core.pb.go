// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: apis/v1/core.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Repeat_Day int32

const (
	Repeat_SUNDAY    Repeat_Day = 0
	Repeat_MONDAY    Repeat_Day = 1
	Repeat_TUESDAY   Repeat_Day = 2
	Repeat_WEDNESDAY Repeat_Day = 3
	Repeat_THURSDAY  Repeat_Day = 4
	Repeat_FRIDAY    Repeat_Day = 5
	Repeat_SATURDAY  Repeat_Day = 6
)

// Enum value maps for Repeat_Day.
var (
	Repeat_Day_name = map[int32]string{
		0: "SUNDAY",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
	}
	Repeat_Day_value = map[string]int32{
		"SUNDAY":    0,
		"MONDAY":    1,
		"TUESDAY":   2,
		"WEDNESDAY": 3,
		"THURSDAY":  4,
		"FRIDAY":    5,
		"SATURDAY":  6,
	}
)

func (x Repeat_Day) Enum() *Repeat_Day {
	p := new(Repeat_Day)
	*p = x
	return p
}

func (x Repeat_Day) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Repeat_Day) Descriptor() protoreflect.EnumDescriptor {
	return file_apis_v1_core_proto_enumTypes[0].Descriptor()
}

func (Repeat_Day) Type() protoreflect.EnumType {
	return &file_apis_v1_core_proto_enumTypes[0]
}

func (x Repeat_Day) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Repeat_Day.Descriptor instead.
func (Repeat_Day) EnumDescriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{3, 0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateOutput string      `protobuf:"bytes,1,opt,name=state_output,json=stateOutput,proto3" json:"state_output,omitempty"`
	Schedules   []*Schedule `protobuf:"bytes,2,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetStateOutput() string {
	if x != nil {
		return x.StateOutput
	}
	return ""
}

func (x *Config) GetSchedules() []*Schedule {
	if x != nil {
		return x.Schedules
	}
	return nil
}

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string    `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	StopAt  string    `protobuf:"bytes,2,opt,name=stop_at,json=stopAt,proto3" json:"stop_at,omitempty"`
	StartAt string    `protobuf:"bytes,3,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	Repeat  *Repeat   `protobuf:"bytes,4,opt,name=repeat,proto3" json:"repeat,omitempty"`
	Except  []*Except `protobuf:"bytes,5,rep,name=except,proto3" json:"except,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{1}
}

func (x *Schedule) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Schedule) GetStopAt() string {
	if x != nil {
		return x.StopAt
	}
	return ""
}

func (x *Schedule) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *Schedule) GetRepeat() *Repeat {
	if x != nil {
		return x.Repeat
	}
	return nil
}

func (x *Schedule) GetExcept() []*Except {
	if x != nil {
		return x.Except
	}
	return nil
}

type Except struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Specifier:
	//
	//	*Except_Cluster_
	//	*Except_Sql_
	//	*Except_Vm_
	Specifier isExcept_Specifier `protobuf_oneof:"specifier"`
}

func (x *Except) Reset() {
	*x = Except{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Except) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Except) ProtoMessage() {}

func (x *Except) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Except.ProtoReflect.Descriptor instead.
func (*Except) Descriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{2}
}

func (m *Except) GetSpecifier() isExcept_Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

func (x *Except) GetCluster() *Except_Cluster {
	if x, ok := x.GetSpecifier().(*Except_Cluster_); ok {
		return x.Cluster
	}
	return nil
}

func (x *Except) GetSql() *Except_Sql {
	if x, ok := x.GetSpecifier().(*Except_Sql_); ok {
		return x.Sql
	}
	return nil
}

func (x *Except) GetVm() *Except_Vm {
	if x, ok := x.GetSpecifier().(*Except_Vm_); ok {
		return x.Vm
	}
	return nil
}

type isExcept_Specifier interface {
	isExcept_Specifier()
}

type Except_Cluster_ struct {
	Cluster *Except_Cluster `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type Except_Sql_ struct {
	Sql *Except_Sql `protobuf:"bytes,2,opt,name=sql,proto3,oneof"`
}

type Except_Vm_ struct {
	Vm *Except_Vm `protobuf:"bytes,3,opt,name=vm,proto3,oneof"`
}

func (*Except_Cluster_) isExcept_Specifier() {}

func (*Except_Sql_) isExcept_Specifier() {}

func (*Except_Vm_) isExcept_Specifier() {}

type Repeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Specifier:
	//
	//	*Repeat_EveryDay
	//	*Repeat_WeekDays
	//	*Repeat_Weekends
	//	*Repeat_Other_
	Specifier isRepeat_Specifier `protobuf_oneof:"specifier"`
}

func (x *Repeat) Reset() {
	*x = Repeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repeat) ProtoMessage() {}

func (x *Repeat) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repeat.ProtoReflect.Descriptor instead.
func (*Repeat) Descriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{3}
}

func (m *Repeat) GetSpecifier() isRepeat_Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

func (x *Repeat) GetEveryDay() bool {
	if x, ok := x.GetSpecifier().(*Repeat_EveryDay); ok {
		return x.EveryDay
	}
	return false
}

func (x *Repeat) GetWeekDays() bool {
	if x, ok := x.GetSpecifier().(*Repeat_WeekDays); ok {
		return x.WeekDays
	}
	return false
}

func (x *Repeat) GetWeekends() bool {
	if x, ok := x.GetSpecifier().(*Repeat_Weekends); ok {
		return x.Weekends
	}
	return false
}

func (x *Repeat) GetOther() *Repeat_Other {
	if x, ok := x.GetSpecifier().(*Repeat_Other_); ok {
		return x.Other
	}
	return nil
}

type isRepeat_Specifier interface {
	isRepeat_Specifier()
}

type Repeat_EveryDay struct {
	EveryDay bool `protobuf:"varint,1,opt,name=every_day,json=everyDay,proto3,oneof"`
}

type Repeat_WeekDays struct {
	WeekDays bool `protobuf:"varint,2,opt,name=week_days,json=weekDays,proto3,oneof"`
}

type Repeat_Weekends struct {
	Weekends bool `protobuf:"varint,3,opt,name=weekends,proto3,oneof"`
}

type Repeat_Other_ struct {
	Other *Repeat_Other `protobuf:"bytes,4,opt,name=other,proto3,oneof"`
}

func (*Repeat_EveryDay) isRepeat_Specifier() {}

func (*Repeat_WeekDays) isRepeat_Specifier() {}

func (*Repeat_Weekends) isRepeat_Specifier() {}

func (*Repeat_Other_) isRepeat_Specifier() {}

type Except_Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pools  []string `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools,omitempty"`
	Zone   string   `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	Region string   `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *Except_Cluster) Reset() {
	*x = Except_Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Except_Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Except_Cluster) ProtoMessage() {}

func (x *Except_Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Except_Cluster.ProtoReflect.Descriptor instead.
func (*Except_Cluster) Descriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Except_Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Except_Cluster) GetPools() []string {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *Except_Cluster) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Except_Cluster) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type Except_Sql struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Except_Sql) Reset() {
	*x = Except_Sql{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Except_Sql) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Except_Sql) ProtoMessage() {}

func (x *Except_Sql) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Except_Sql.ProtoReflect.Descriptor instead.
func (*Except_Sql) Descriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Except_Sql) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Except_Vm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Zone string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
}

func (x *Except_Vm) Reset() {
	*x = Except_Vm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Except_Vm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Except_Vm) ProtoMessage() {}

func (x *Except_Vm) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Except_Vm.ProtoReflect.Descriptor instead.
func (*Except_Vm) Descriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{2, 2}
}

func (x *Except_Vm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Except_Vm) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

type Repeat_Other struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Days []Repeat_Day `protobuf:"varint,1,rep,packed,name=days,proto3,enum=v1.Repeat_Day" json:"days,omitempty"`
}

func (x *Repeat_Other) Reset() {
	*x = Repeat_Other{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repeat_Other) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repeat_Other) ProtoMessage() {}

func (x *Repeat_Other) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repeat_Other.ProtoReflect.Descriptor instead.
func (*Repeat_Other) Descriptor() ([]byte, []int) {
	return file_apis_v1_core_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Repeat_Other) GetDays() []Repeat_Day {
	if x != nil {
		return x.Days
	}
	return nil
}

var File_apis_v1_core_proto protoreflect.FileDescriptor

var file_apis_v1_core_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x57, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2a,
	0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xfd, 0x01, 0x0a, 0x08, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x73, 0x74,
	0x6f, 0x70, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xfa, 0x42, 0x25,
	0x72, 0x23, 0x32, 0x21, 0x5e, 0x28, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5b, 0x30, 0x2d, 0x39,
	0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x33, 0x5d, 0x29, 0x3a, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x5b,
	0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x70, 0x41, 0x74, 0x12, 0x43, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x28, 0xfa, 0x42, 0x25, 0x72, 0x23, 0x32, 0x21, 0x5e, 0x28, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f,
	0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x33, 0x5d, 0x29, 0x3a, 0x5b, 0x30,
	0x2d, 0x35, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x41, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x65,
	0x70, 0x74, 0x52, 0x06, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x22, 0xb4, 0x02, 0x0a, 0x06, 0x45,
	0x78, 0x63, 0x65, 0x70, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x65,
	0x70, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x03, 0x73, 0x71, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x53,
	0x71, 0x6c, 0x48, 0x00, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x12, 0x1f, 0x0a, 0x02, 0x76, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x2e, 0x56, 0x6d, 0x48, 0x00, 0x52, 0x02, 0x76, 0x6d, 0x1a, 0x5f, 0x0a, 0x07, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x6f,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x7a,
	0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x0a, 0x03, 0x53,
	0x71, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x2c, 0x0a, 0x02, 0x56, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x7a, 0x6f, 0x6e, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x22, 0xab, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x09,
	0x65, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x00, 0x52, 0x08, 0x65, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x79, 0x12, 0x1d, 0x0a, 0x09, 0x77,
	0x65, 0x65, 0x6b, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x08, 0x77, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1c, 0x0a, 0x08, 0x77, 0x65,
	0x65, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08,
	0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x1a, 0x2b, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x2e, 0x44, 0x61, 0x79, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x22,
	0x61, 0x0a, 0x03, 0x44, 0x61, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x54, 0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57,
	0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x48,
	0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x49, 0x44,
	0x41, 0x59, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59,
	0x10, 0x06, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69,
	0x65, 0x74, 0x61, 0x6e, 0x68, 0x64, 0x75, 0x6f, 0x6e, 0x67, 0x2f, 0x70, 0x61, 0x75, 0x73, 0x65,
	0x2d, 0x67, 0x6b, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_v1_core_proto_rawDescOnce sync.Once
	file_apis_v1_core_proto_rawDescData = file_apis_v1_core_proto_rawDesc
)

func file_apis_v1_core_proto_rawDescGZIP() []byte {
	file_apis_v1_core_proto_rawDescOnce.Do(func() {
		file_apis_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_v1_core_proto_rawDescData)
	})
	return file_apis_v1_core_proto_rawDescData
}

var file_apis_v1_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apis_v1_core_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apis_v1_core_proto_goTypes = []interface{}{
	(Repeat_Day)(0),        // 0: v1.Repeat.Day
	(*Config)(nil),         // 1: v1.Config
	(*Schedule)(nil),       // 2: v1.Schedule
	(*Except)(nil),         // 3: v1.Except
	(*Repeat)(nil),         // 4: v1.Repeat
	(*Except_Cluster)(nil), // 5: v1.Except.Cluster
	(*Except_Sql)(nil),     // 6: v1.Except.Sql
	(*Except_Vm)(nil),      // 7: v1.Except.Vm
	(*Repeat_Other)(nil),   // 8: v1.Repeat.Other
}
var file_apis_v1_core_proto_depIdxs = []int32{
	2, // 0: v1.Config.schedules:type_name -> v1.Schedule
	4, // 1: v1.Schedule.repeat:type_name -> v1.Repeat
	3, // 2: v1.Schedule.except:type_name -> v1.Except
	5, // 3: v1.Except.cluster:type_name -> v1.Except.Cluster
	6, // 4: v1.Except.sql:type_name -> v1.Except.Sql
	7, // 5: v1.Except.vm:type_name -> v1.Except.Vm
	8, // 6: v1.Repeat.other:type_name -> v1.Repeat.Other
	0, // 7: v1.Repeat.Other.days:type_name -> v1.Repeat.Day
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_apis_v1_core_proto_init() }
func file_apis_v1_core_proto_init() {
	if File_apis_v1_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_v1_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_apis_v1_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
		file_apis_v1_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Except); i {
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
		file_apis_v1_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repeat); i {
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
		file_apis_v1_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Except_Cluster); i {
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
		file_apis_v1_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Except_Sql); i {
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
		file_apis_v1_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Except_Vm); i {
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
		file_apis_v1_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repeat_Other); i {
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
	file_apis_v1_core_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Except_Cluster_)(nil),
		(*Except_Sql_)(nil),
		(*Except_Vm_)(nil),
	}
	file_apis_v1_core_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Repeat_EveryDay)(nil),
		(*Repeat_WeekDays)(nil),
		(*Repeat_Weekends)(nil),
		(*Repeat_Other_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_v1_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_v1_core_proto_goTypes,
		DependencyIndexes: file_apis_v1_core_proto_depIdxs,
		EnumInfos:         file_apis_v1_core_proto_enumTypes,
		MessageInfos:      file_apis_v1_core_proto_msgTypes,
	}.Build()
	File_apis_v1_core_proto = out.File
	file_apis_v1_core_proto_rawDesc = nil
	file_apis_v1_core_proto_goTypes = nil
	file_apis_v1_core_proto_depIdxs = nil
}
