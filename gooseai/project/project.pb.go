// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.5
// source: project.proto

package project

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

type ProjectAccess int32

const (
	ProjectAccess_PROJECT_ACCESS_PRIVATE ProjectAccess = 0 // Private access, only owner organization can access
	ProjectAccess_PROJECT_ACCESS_PUBLIC  ProjectAccess = 1 // Public access, anyone can access
)

// Enum value maps for ProjectAccess.
var (
	ProjectAccess_name = map[int32]string{
		0: "PROJECT_ACCESS_PRIVATE",
		1: "PROJECT_ACCESS_PUBLIC",
	}
	ProjectAccess_value = map[string]int32{
		"PROJECT_ACCESS_PRIVATE": 0,
		"PROJECT_ACCESS_PUBLIC":  1,
	}
)

func (x ProjectAccess) Enum() *ProjectAccess {
	p := new(ProjectAccess)
	*p = x
	return p
}

func (x ProjectAccess) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectAccess) Descriptor() protoreflect.EnumDescriptor {
	return file_project_proto_enumTypes[0].Descriptor()
}

func (ProjectAccess) Type() protoreflect.EnumType {
	return &file_project_proto_enumTypes[0]
}

func (x ProjectAccess) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectAccess.Descriptor instead.
func (ProjectAccess) EnumDescriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{0}
}

type ProjectStatus int32

const (
	ProjectStatus_PROJECT_STATUS_INACTIVE ProjectStatus = 0 // Inactive project
	ProjectStatus_PROJECT_STATUS_ACTIVE   ProjectStatus = 1 // Active project
	ProjectStatus_PROJECT_STATUS_DELETED  ProjectStatus = 2 // Deleted project, only set by delete endpoint
)

// Enum value maps for ProjectStatus.
var (
	ProjectStatus_name = map[int32]string{
		0: "PROJECT_STATUS_INACTIVE",
		1: "PROJECT_STATUS_ACTIVE",
		2: "PROJECT_STATUS_DELETED",
	}
	ProjectStatus_value = map[string]int32{
		"PROJECT_STATUS_INACTIVE": 0,
		"PROJECT_STATUS_ACTIVE":   1,
		"PROJECT_STATUS_DELETED":  2,
	}
)

func (x ProjectStatus) Enum() *ProjectStatus {
	p := new(ProjectStatus)
	*p = x
	return p
}

func (x ProjectStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_project_proto_enumTypes[1].Descriptor()
}

func (ProjectStatus) Type() protoreflect.EnumType {
	return &file_project_proto_enumTypes[1]
}

func (x ProjectStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectStatus.Descriptor instead.
func (ProjectStatus) EnumDescriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{1}
}

type ProjectAssetUse int32

const (
	ProjectAssetUse_PROJECT_ASSET_USE_UNDEFINED    ProjectAssetUse = 0 // Asset does not have use defined
	ProjectAssetUse_PROJECT_ASSET_USE_INPUT        ProjectAssetUse = 1 // Asset is used as an input for the project
	ProjectAssetUse_PROJECT_ASSET_USE_OUTPUT       ProjectAssetUse = 2 // Asset is an output from the project
	ProjectAssetUse_PROJECT_ASSET_USE_INTERMEDIATE ProjectAssetUse = 3 // Asset is an output from an intermediate step of the project
	ProjectAssetUse_PROJECT_ASSET_USE_PROJECT      ProjectAssetUse = 4 // Asset is used as the project file for the project
)

// Enum value maps for ProjectAssetUse.
var (
	ProjectAssetUse_name = map[int32]string{
		0: "PROJECT_ASSET_USE_UNDEFINED",
		1: "PROJECT_ASSET_USE_INPUT",
		2: "PROJECT_ASSET_USE_OUTPUT",
		3: "PROJECT_ASSET_USE_INTERMEDIATE",
		4: "PROJECT_ASSET_USE_PROJECT",
	}
	ProjectAssetUse_value = map[string]int32{
		"PROJECT_ASSET_USE_UNDEFINED":    0,
		"PROJECT_ASSET_USE_INPUT":        1,
		"PROJECT_ASSET_USE_OUTPUT":       2,
		"PROJECT_ASSET_USE_INTERMEDIATE": 3,
		"PROJECT_ASSET_USE_PROJECT":      4,
	}
)

func (x ProjectAssetUse) Enum() *ProjectAssetUse {
	p := new(ProjectAssetUse)
	*p = x
	return p
}

func (x ProjectAssetUse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectAssetUse) Descriptor() protoreflect.EnumDescriptor {
	return file_project_proto_enumTypes[2].Descriptor()
}

func (ProjectAssetUse) Type() protoreflect.EnumType {
	return &file_project_proto_enumTypes[2]
}

func (x ProjectAssetUse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectAssetUse.Descriptor instead.
func (ProjectAssetUse) EnumDescriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{2}
}

type ProjectAsset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                 // ID of the asset, UUIDv4
	Uri string          `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`                               // The URI to the asset
	Use ProjectAssetUse `protobuf:"varint,3,opt,name=use,proto3,enum=gooseai.ProjectAssetUse" json:"use,omitempty"` // The use of the asset with respect to the project
}

func (x *ProjectAsset) Reset() {
	*x = ProjectAsset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAsset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAsset) ProtoMessage() {}

func (x *ProjectAsset) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAsset.ProtoReflect.Descriptor instead.
func (*ProjectAsset) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectAsset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProjectAsset) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *ProjectAsset) GetUse() ProjectAssetUse {
	if x != nil {
		return x.Use
	}
	return ProjectAssetUse_PROJECT_ASSET_USE_UNDEFINED
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                     // ID of the project, UUIDv4
	Title     string          `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                               // Title of the project
	OwnerId   string          `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`            // The ID of the organization owning the project
	Access    ProjectAccess   `protobuf:"varint,4,opt,name=access,proto3,enum=gooseai.ProjectAccess" json:"access,omitempty"` // The access of the project (such as public vs private)
	Status    ProjectStatus   `protobuf:"varint,5,opt,name=status,proto3,enum=gooseai.ProjectStatus" json:"status,omitempty"` // The status of the project (such as active vs inactive)
	Size      uint64          `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`                                // The size of the project in bytes
	File      *ProjectAsset   `protobuf:"bytes,7,opt,name=file,proto3" json:"file,omitempty"`                                 // Project file for the project
	CreatedAt uint64          `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`     // Time of project creation (UTC seconds epoch)
	UpdatedAt uint64          `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`     // Time of last project update (UTC seconds epoch)
	Assets    []*ProjectAsset `protobuf:"bytes,10,rep,name=assets,proto3" json:"assets,omitempty"`                            // The listing of all assets associated with the project
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{1}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Project) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Project) GetAccess() ProjectAccess {
	if x != nil {
		return x.Access
	}
	return ProjectAccess_PROJECT_ACCESS_PRIVATE
}

func (x *Project) GetStatus() ProjectStatus {
	if x != nil {
		return x.Status
	}
	return ProjectStatus_PROJECT_STATUS_INACTIVE
}

func (x *Project) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Project) GetFile() *ProjectAsset {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Project) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Project) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Project) GetAssets() []*ProjectAsset {
	if x != nil {
		return x.Assets
	}
	return nil
}

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`                               // title of the project
	OwnerId *string       `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3,oneof" json:"owner_id,omitempty"`      // The ID of the organization owning the project, empty for default org for user
	Access  ProjectAccess `protobuf:"varint,3,opt,name=access,proto3,enum=gooseai.ProjectAccess" json:"access,omitempty"` // The access of the project (such as public vs private)
	Status  ProjectStatus `protobuf:"varint,4,opt,name=status,proto3,enum=gooseai.ProjectStatus" json:"status,omitempty"` // The status of the project (such as active vs inactive)
	File    *ProjectAsset `protobuf:"bytes,5,opt,name=file,proto3,oneof" json:"file,omitempty"`                           // Project file for the project (if existing, else will create)
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProjectRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateProjectRequest) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

func (x *CreateProjectRequest) GetAccess() ProjectAccess {
	if x != nil {
		return x.Access
	}
	return ProjectAccess_PROJECT_ACCESS_PRIVATE
}

func (x *CreateProjectRequest) GetStatus() ProjectStatus {
	if x != nil {
		return x.Status
	}
	return ProjectStatus_PROJECT_STATUS_INACTIVE
}

func (x *CreateProjectRequest) GetFile() *ProjectAsset {
	if x != nil {
		return x.File
	}
	return nil
}

type UpdateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                           // ID of the project, UUIDv4
	OwnerId *string        `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3,oneof" json:"owner_id,omitempty"`            // The ID of the organization owning the project, empty for default org for user
	Title   *string        `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`                               // Title of the project
	Access  *ProjectAccess `protobuf:"varint,4,opt,name=access,proto3,enum=gooseai.ProjectAccess,oneof" json:"access,omitempty"` // The access of the project (such as public vs private)
	Status  *ProjectStatus `protobuf:"varint,5,opt,name=status,proto3,enum=gooseai.ProjectStatus,oneof" json:"status,omitempty"` // The status of the project (such as active vs inactive)
	File    *ProjectAsset  `protobuf:"bytes,6,opt,name=file,proto3,oneof" json:"file,omitempty"`                                 // Project file for the project
}

func (x *UpdateProjectRequest) Reset() {
	*x = UpdateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectRequest) ProtoMessage() {}

func (x *UpdateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProjectRequest) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

func (x *UpdateProjectRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateProjectRequest) GetAccess() ProjectAccess {
	if x != nil && x.Access != nil {
		return *x.Access
	}
	return ProjectAccess_PROJECT_ACCESS_PRIVATE
}

func (x *UpdateProjectRequest) GetStatus() ProjectStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ProjectStatus_PROJECT_STATUS_INACTIVE
}

func (x *UpdateProjectRequest) GetFile() *ProjectAsset {
	if x != nil {
		return x.File
	}
	return nil
}

type ListProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId *string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3,oneof" json:"owner_id,omitempty"` // The ID of the organization owning the project, empty for default org for user
}

func (x *ListProjectRequest) Reset() {
	*x = ListProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectRequest) ProtoMessage() {}

func (x *ListProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectRequest.ProtoReflect.Descriptor instead.
func (*ListProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{4}
}

func (x *ListProjectRequest) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

type GetProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                // ID of the project to request, UUIDv4
	OwnerId *string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3,oneof" json:"owner_id,omitempty"` // The ID of the organization owning the project, empty for default org for user
}

func (x *GetProjectRequest) Reset() {
	*x = GetProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectRequest) ProtoMessage() {}

func (x *GetProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectRequest.ProtoReflect.Descriptor instead.
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{5}
}

func (x *GetProjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProjectRequest) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

type DeleteProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                // ID of the project to request, UUIDv4
	OwnerId *string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3,oneof" json:"owner_id,omitempty"` // The ID of the organization owning the project, empty for default org for user
}

func (x *DeleteProjectRequest) Reset() {
	*x = DeleteProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProjectRequest) ProtoMessage() {}

func (x *DeleteProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteProjectRequest) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

var File_project_proto protoreflect.FileDescriptor

var file_project_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x22, 0x5c, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x2a, 0x0a, 0x03, 0x75, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61,
	0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x52, 0x03, 0x75, 0x73, 0x65, 0x22, 0xd6, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x2d, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22,
	0xf2, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e,
	0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2e,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x48, 0x01, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0xb1, 0x02, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65,
	0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x48, 0x02, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x48, 0x04, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x41, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x53, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x2a, 0x46, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x2a, 0x63, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x50,
	0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x2a,
	0xb0, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41,
	0x53, 0x53, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10,
	0x01, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x53, 0x53,
	0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x10, 0x02, 0x12,
	0x22, 0x0a, 0x1e, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54,
	0x5f, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54,
	0x45, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41,
	0x53, 0x53, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54,
	0x10, 0x04, 0x32, 0xb9, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x6f,
	0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x67, 0x6f, 0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x73, 0x65, 0x61, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x67, 0x6f, 0x6f,
	0x73, 0x65, 0x61, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_proto_rawDescOnce sync.Once
	file_project_proto_rawDescData = file_project_proto_rawDesc
)

func file_project_proto_rawDescGZIP() []byte {
	file_project_proto_rawDescOnce.Do(func() {
		file_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_proto_rawDescData)
	})
	return file_project_proto_rawDescData
}

var file_project_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_project_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_project_proto_goTypes = []interface{}{
	(ProjectAccess)(0),           // 0: gooseai.ProjectAccess
	(ProjectStatus)(0),           // 1: gooseai.ProjectStatus
	(ProjectAssetUse)(0),         // 2: gooseai.ProjectAssetUse
	(*ProjectAsset)(nil),         // 3: gooseai.ProjectAsset
	(*Project)(nil),              // 4: gooseai.Project
	(*CreateProjectRequest)(nil), // 5: gooseai.CreateProjectRequest
	(*UpdateProjectRequest)(nil), // 6: gooseai.UpdateProjectRequest
	(*ListProjectRequest)(nil),   // 7: gooseai.ListProjectRequest
	(*GetProjectRequest)(nil),    // 8: gooseai.GetProjectRequest
	(*DeleteProjectRequest)(nil), // 9: gooseai.DeleteProjectRequest
}
var file_project_proto_depIdxs = []int32{
	2,  // 0: gooseai.ProjectAsset.use:type_name -> gooseai.ProjectAssetUse
	0,  // 1: gooseai.Project.access:type_name -> gooseai.ProjectAccess
	1,  // 2: gooseai.Project.status:type_name -> gooseai.ProjectStatus
	3,  // 3: gooseai.Project.file:type_name -> gooseai.ProjectAsset
	3,  // 4: gooseai.Project.assets:type_name -> gooseai.ProjectAsset
	0,  // 5: gooseai.CreateProjectRequest.access:type_name -> gooseai.ProjectAccess
	1,  // 6: gooseai.CreateProjectRequest.status:type_name -> gooseai.ProjectStatus
	3,  // 7: gooseai.CreateProjectRequest.file:type_name -> gooseai.ProjectAsset
	0,  // 8: gooseai.UpdateProjectRequest.access:type_name -> gooseai.ProjectAccess
	1,  // 9: gooseai.UpdateProjectRequest.status:type_name -> gooseai.ProjectStatus
	3,  // 10: gooseai.UpdateProjectRequest.file:type_name -> gooseai.ProjectAsset
	5,  // 11: gooseai.ProjectService.Create:input_type -> gooseai.CreateProjectRequest
	6,  // 12: gooseai.ProjectService.Update:input_type -> gooseai.UpdateProjectRequest
	7,  // 13: gooseai.ProjectService.List:input_type -> gooseai.ListProjectRequest
	8,  // 14: gooseai.ProjectService.Get:input_type -> gooseai.GetProjectRequest
	9,  // 15: gooseai.ProjectService.Delete:input_type -> gooseai.DeleteProjectRequest
	4,  // 16: gooseai.ProjectService.Create:output_type -> gooseai.Project
	4,  // 17: gooseai.ProjectService.Update:output_type -> gooseai.Project
	4,  // 18: gooseai.ProjectService.List:output_type -> gooseai.Project
	4,  // 19: gooseai.ProjectService.Get:output_type -> gooseai.Project
	4,  // 20: gooseai.ProjectService.Delete:output_type -> gooseai.Project
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_project_proto_init() }
func file_project_proto_init() {
	if File_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAsset); i {
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
		file_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRequest); i {
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
		file_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProjectRequest); i {
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
		file_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectRequest); i {
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
		file_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectRequest); i {
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
		file_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProjectRequest); i {
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
	file_project_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_project_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_project_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_project_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_project_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_proto_goTypes,
		DependencyIndexes: file_project_proto_depIdxs,
		EnumInfos:         file_project_proto_enumTypes,
		MessageInfos:      file_project_proto_msgTypes,
	}.Build()
	File_project_proto = out.File
	file_project_proto_rawDesc = nil
	file_project_proto_goTypes = nil
	file_project_proto_depIdxs = nil
}
