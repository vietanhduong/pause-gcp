// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: apis/v1/resource.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Specifier:
	//
	//	*Resource_Cluster
	//	*Resource_Sql
	//	*Resource_Vm
	Specifier isResource_Specifier   `protobuf_oneof:"specifier"`
	PausedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=paused_at,json=pausedAt,proto3" json:"paused_at,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_apis_v1_resource_proto_rawDescGZIP(), []int{0}
}

func (m *Resource) GetSpecifier() isResource_Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

func (x *Resource) GetCluster() *Cluster {
	if x, ok := x.GetSpecifier().(*Resource_Cluster); ok {
		return x.Cluster
	}
	return nil
}

func (x *Resource) GetSql() *Sql {
	if x, ok := x.GetSpecifier().(*Resource_Sql); ok {
		return x.Sql
	}
	return nil
}

func (x *Resource) GetVm() *Vm {
	if x, ok := x.GetSpecifier().(*Resource_Vm); ok {
		return x.Vm
	}
	return nil
}

func (x *Resource) GetPausedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PausedAt
	}
	return nil
}

type isResource_Specifier interface {
	isResource_Specifier()
}

type Resource_Cluster struct {
	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type Resource_Sql struct {
	Sql *Sql `protobuf:"bytes,2,opt,name=sql,proto3,oneof"`
}

type Resource_Vm struct {
	Vm *Vm `protobuf:"bytes,3,opt,name=vm,proto3,oneof"`
}

func (*Resource_Cluster) isResource_Specifier() {}

func (*Resource_Sql) isResource_Specifier() {}

func (*Resource_Vm) isResource_Specifier() {}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project   string              `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Name      string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location  string              `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	NodePools []*Cluster_NodePool `protobuf:"bytes,4,rep,name=node_pools,json=nodePools,proto3" json:"node_pools,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_apis_v1_resource_proto_rawDescGZIP(), []int{1}
}

func (x *Cluster) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Cluster) GetNodePools() []*Cluster_NodePool {
	if x != nil {
		return x.NodePools
	}
	return nil
}

type Sql struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Project         string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	Region          string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	GceZone         string `protobuf:"bytes,4,opt,name=gce_zone,json=gceZone,proto3" json:"gce_zone,omitempty"`
	IsRunning       bool   `protobuf:"varint,5,opt,name=is_running,json=isRunning,proto3" json:"is_running,omitempty"`
	DatabaseVersion string `protobuf:"bytes,6,opt,name=database_version,json=databaseVersion,proto3" json:"database_version,omitempty"`
}

func (x *Sql) Reset() {
	*x = Sql{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sql) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sql) ProtoMessage() {}

func (x *Sql) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sql.ProtoReflect.Descriptor instead.
func (*Sql) Descriptor() ([]byte, []int) {
	return file_apis_v1_resource_proto_rawDescGZIP(), []int{2}
}

func (x *Sql) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sql) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Sql) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Sql) GetGceZone() string {
	if x != nil {
		return x.GceZone
	}
	return ""
}

func (x *Sql) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

func (x *Sql) GetDatabaseVersion() string {
	if x != nil {
		return x.DatabaseVersion
	}
	return ""
}

type Vm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Zone    string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Project string `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *Vm) Reset() {
	*x = Vm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vm) ProtoMessage() {}

func (x *Vm) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vm.ProtoReflect.Descriptor instead.
func (*Vm) Descriptor() ([]byte, []int) {
	return file_apis_v1_resource_proto_rawDescGZIP(), []int{3}
}

func (x *Vm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Vm) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Vm) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Vm) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type Cluster_NodePool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InstanceGroups   []string                      `protobuf:"bytes,2,rep,name=instance_groups,json=instanceGroups,proto3" json:"instance_groups,omitempty"`
	Locations        []string                      `protobuf:"bytes,3,rep,name=locations,proto3" json:"locations,omitempty"`
	InitialNodeCount int32                         `protobuf:"varint,4,opt,name=initial_node_count,json=initialNodeCount,proto3" json:"initial_node_count,omitempty"`
	CurrentSize      int32                         `protobuf:"varint,5,opt,name=current_size,json=currentSize,proto3" json:"current_size,omitempty"`
	Autoscaling      *Cluster_NodePool_AutoScaling `protobuf:"bytes,6,opt,name=autoscaling,proto3" json:"autoscaling,omitempty"`
	Spot             bool                          `protobuf:"varint,7,opt,name=spot,proto3" json:"spot,omitempty"`
	Preemptible      bool                          `protobuf:"varint,8,opt,name=preemptible,proto3" json:"preemptible,omitempty"`
}

func (x *Cluster_NodePool) Reset() {
	*x = Cluster_NodePool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster_NodePool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster_NodePool) ProtoMessage() {}

func (x *Cluster_NodePool) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster_NodePool.ProtoReflect.Descriptor instead.
func (*Cluster_NodePool) Descriptor() ([]byte, []int) {
	return file_apis_v1_resource_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Cluster_NodePool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster_NodePool) GetInstanceGroups() []string {
	if x != nil {
		return x.InstanceGroups
	}
	return nil
}

func (x *Cluster_NodePool) GetLocations() []string {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *Cluster_NodePool) GetInitialNodeCount() int32 {
	if x != nil {
		return x.InitialNodeCount
	}
	return 0
}

func (x *Cluster_NodePool) GetCurrentSize() int32 {
	if x != nil {
		return x.CurrentSize
	}
	return 0
}

func (x *Cluster_NodePool) GetAutoscaling() *Cluster_NodePool_AutoScaling {
	if x != nil {
		return x.Autoscaling
	}
	return nil
}

func (x *Cluster_NodePool) GetSpot() bool {
	if x != nil {
		return x.Spot
	}
	return false
}

func (x *Cluster_NodePool) GetPreemptible() bool {
	if x != nil {
		return x.Preemptible
	}
	return false
}

type Cluster_NodePool_AutoScaling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled           bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	MinNodeCount      int32  `protobuf:"varint,2,opt,name=min_node_count,json=minNodeCount,proto3" json:"min_node_count,omitempty"`
	MaxNodeCount      int32  `protobuf:"varint,3,opt,name=max_node_count,json=maxNodeCount,proto3" json:"max_node_count,omitempty"`
	Autoprovisioned   bool   `protobuf:"varint,4,opt,name=autoprovisioned,proto3" json:"autoprovisioned,omitempty"`
	LocationPolicy    string `protobuf:"bytes,5,opt,name=location_policy,json=locationPolicy,proto3" json:"location_policy,omitempty"`
	TotalMinNodeCount int32  `protobuf:"varint,6,opt,name=total_min_node_count,json=totalMinNodeCount,proto3" json:"total_min_node_count,omitempty"`
	TotalMaxNodeCount int32  `protobuf:"varint,7,opt,name=total_max_node_count,json=totalMaxNodeCount,proto3" json:"total_max_node_count,omitempty"`
}

func (x *Cluster_NodePool_AutoScaling) Reset() {
	*x = Cluster_NodePool_AutoScaling{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_v1_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster_NodePool_AutoScaling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster_NodePool_AutoScaling) ProtoMessage() {}

func (x *Cluster_NodePool_AutoScaling) ProtoReflect() protoreflect.Message {
	mi := &file_apis_v1_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster_NodePool_AutoScaling.ProtoReflect.Descriptor instead.
func (*Cluster_NodePool_AutoScaling) Descriptor() ([]byte, []int) {
	return file_apis_v1_resource_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *Cluster_NodePool_AutoScaling) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Cluster_NodePool_AutoScaling) GetMinNodeCount() int32 {
	if x != nil {
		return x.MinNodeCount
	}
	return 0
}

func (x *Cluster_NodePool_AutoScaling) GetMaxNodeCount() int32 {
	if x != nil {
		return x.MaxNodeCount
	}
	return 0
}

func (x *Cluster_NodePool_AutoScaling) GetAutoprovisioned() bool {
	if x != nil {
		return x.Autoprovisioned
	}
	return false
}

func (x *Cluster_NodePool_AutoScaling) GetLocationPolicy() string {
	if x != nil {
		return x.LocationPolicy
	}
	return ""
}

func (x *Cluster_NodePool_AutoScaling) GetTotalMinNodeCount() int32 {
	if x != nil {
		return x.TotalMinNodeCount
	}
	return 0
}

func (x *Cluster_NodePool_AutoScaling) GetTotalMaxNodeCount() int32 {
	if x != nil {
		return x.TotalMaxNodeCount
	}
	return 0
}

var File_apis_v1_resource_proto protoreflect.FileDescriptor

var file_apis_v1_resource_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x03, 0x73, 0x71, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x71, 0x6c, 0x48, 0x00, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x12,
	0x1d, 0x0a, 0x02, 0x76, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6d, 0x48, 0x00, 0x52, 0x02, 0x76, 0x6d, 0x12, 0x37,
	0x0a, 0x09, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x22, 0xf0, 0x05, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0a, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x50,
	0x6f, 0x6f, 0x6c, 0x73, 0x1a, 0xe0, 0x04, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x12,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x47, 0x0a,
	0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x6f, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x70, 0x6f, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x65, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x1a, 0xa8, 0x02, 0x0a,
	0x0b, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6d, 0x69, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x75, 0x74,
	0x6f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2f, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d,
	0x69, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x69, 0x6e, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x61, 0x78, 0x4e, 0x6f,
	0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x03, 0x53, 0x71, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x63, 0x65, 0x5f, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x63, 0x65, 0x5a, 0x6f, 0x6e, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x29, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x02, 0x56, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x65, 0x74, 0x61, 0x6e, 0x68, 0x64, 0x75,
	0x6f, 0x6e, 0x67, 0x2f, 0x70, 0x61, 0x75, 0x73, 0x65, 0x2d, 0x67, 0x6b, 0x65, 0x2d, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_v1_resource_proto_rawDescOnce sync.Once
	file_apis_v1_resource_proto_rawDescData = file_apis_v1_resource_proto_rawDesc
)

func file_apis_v1_resource_proto_rawDescGZIP() []byte {
	file_apis_v1_resource_proto_rawDescOnce.Do(func() {
		file_apis_v1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_v1_resource_proto_rawDescData)
	})
	return file_apis_v1_resource_proto_rawDescData
}

var file_apis_v1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apis_v1_resource_proto_goTypes = []interface{}{
	(*Resource)(nil),                     // 0: apis.v1.Resource
	(*Cluster)(nil),                      // 1: apis.v1.Cluster
	(*Sql)(nil),                          // 2: apis.v1.Sql
	(*Vm)(nil),                           // 3: apis.v1.Vm
	(*Cluster_NodePool)(nil),             // 4: apis.v1.Cluster.NodePool
	(*Cluster_NodePool_AutoScaling)(nil), // 5: apis.v1.Cluster.NodePool.AutoScaling
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_apis_v1_resource_proto_depIdxs = []int32{
	1, // 0: apis.v1.Resource.cluster:type_name -> apis.v1.Cluster
	2, // 1: apis.v1.Resource.sql:type_name -> apis.v1.Sql
	3, // 2: apis.v1.Resource.vm:type_name -> apis.v1.Vm
	6, // 3: apis.v1.Resource.paused_at:type_name -> google.protobuf.Timestamp
	4, // 4: apis.v1.Cluster.node_pools:type_name -> apis.v1.Cluster.NodePool
	5, // 5: apis.v1.Cluster.NodePool.autoscaling:type_name -> apis.v1.Cluster.NodePool.AutoScaling
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_apis_v1_resource_proto_init() }
func file_apis_v1_resource_proto_init() {
	if File_apis_v1_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_v1_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_apis_v1_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_apis_v1_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sql); i {
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
		file_apis_v1_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vm); i {
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
		file_apis_v1_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster_NodePool); i {
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
		file_apis_v1_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster_NodePool_AutoScaling); i {
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
	file_apis_v1_resource_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Resource_Cluster)(nil),
		(*Resource_Sql)(nil),
		(*Resource_Vm)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_v1_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_v1_resource_proto_goTypes,
		DependencyIndexes: file_apis_v1_resource_proto_depIdxs,
		MessageInfos:      file_apis_v1_resource_proto_msgTypes,
	}.Build()
	File_apis_v1_resource_proto = out.File
	file_apis_v1_resource_proto_rawDesc = nil
	file_apis_v1_resource_proto_goTypes = nil
	file_apis_v1_resource_proto_depIdxs = nil
}
