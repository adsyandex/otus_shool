// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: todo.proto

package todo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Completed     bool                   `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_todo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *Task) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Task) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_todo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_todo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{2}
}

func (x *GetTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Completed     bool                   `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	mi := &file_todo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTaskRequest) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_todo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	mi := &file_todo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type TaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	mi := &file_todo_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{6}
}

func (x *TaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type ListTasksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksRequest) Reset() {
	*x = ListTasksRequest{}
	mi := &file_todo_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksRequest) ProtoMessage() {}

func (x *ListTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksRequest.ProtoReflect.Descriptor instead.
func (*ListTasksRequest) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{7}
}

func (x *ListTasksRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTasksRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListTasksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksResponse) Reset() {
	*x = ListTasksResponse{}
	mi := &file_todo_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksResponse) ProtoMessage() {}

func (x *ListTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksResponse.ProtoReflect.Descriptor instead.
func (*ListTasksResponse) Descriptor() ([]byte, []int) {
	return file_todo_proto_rawDescGZIP(), []int{8}
}

func (x *ListTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *ListTasksResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_todo_proto protoreflect.FileDescriptor

const file_todo_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"todo.proto\x12\x04todo\"\xaa\x01\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1c\n" +
	"\tcompleted\x18\x04 \x01(\bR\tcompleted\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\tR\tupdatedAt\"K\n" +
	"\x11CreateTaskRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\" \n" +
	"\x0eGetTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"y\n" +
	"\x11UpdateTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1c\n" +
	"\tcompleted\x18\x04 \x01(\bR\tcompleted\"#\n" +
	"\x11DeleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\".\n" +
	"\x12DeleteTaskResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\".\n" +
	"\fTaskResponse\x12\x1e\n" +
	"\x04task\x18\x01 \x01(\v2\n" +
	".todo.TaskR\x04task\"@\n" +
	"\x10ListTasksRequest\x12\x14\n" +
	"\x05limit\x18\x01 \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\x02 \x01(\x05R\x06offset\"K\n" +
	"\x11ListTasksResponse\x12 \n" +
	"\x05tasks\x18\x01 \x03(\v2\n" +
	".todo.TaskR\x05tasks\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total2\xb7\x02\n" +
	"\vTodoService\x129\n" +
	"\n" +
	"CreateTask\x12\x17.todo.CreateTaskRequest\x1a\x12.todo.TaskResponse\x123\n" +
	"\aGetTask\x12\x14.todo.GetTaskRequest\x1a\x12.todo.TaskResponse\x129\n" +
	"\n" +
	"UpdateTask\x12\x17.todo.UpdateTaskRequest\x1a\x12.todo.TaskResponse\x12?\n" +
	"\n" +
	"DeleteTask\x12\x17.todo.DeleteTaskRequest\x1a\x18.todo.DeleteTaskResponse\x12<\n" +
	"\tListTasks\x12\x16.todo.ListTasksRequest\x1a\x17.todo.ListTasksResponseB\bZ\x06.;todob\x06proto3"

var (
	file_todo_proto_rawDescOnce sync.Once
	file_todo_proto_rawDescData []byte
)

func file_todo_proto_rawDescGZIP() []byte {
	file_todo_proto_rawDescOnce.Do(func() {
		file_todo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_todo_proto_rawDesc), len(file_todo_proto_rawDesc)))
	})
	return file_todo_proto_rawDescData
}

var file_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_todo_proto_goTypes = []any{
	(*Task)(nil),               // 0: todo.Task
	(*CreateTaskRequest)(nil),  // 1: todo.CreateTaskRequest
	(*GetTaskRequest)(nil),     // 2: todo.GetTaskRequest
	(*UpdateTaskRequest)(nil),  // 3: todo.UpdateTaskRequest
	(*DeleteTaskRequest)(nil),  // 4: todo.DeleteTaskRequest
	(*DeleteTaskResponse)(nil), // 5: todo.DeleteTaskResponse
	(*TaskResponse)(nil),       // 6: todo.TaskResponse
	(*ListTasksRequest)(nil),   // 7: todo.ListTasksRequest
	(*ListTasksResponse)(nil),  // 8: todo.ListTasksResponse
}
var file_todo_proto_depIdxs = []int32{
	0, // 0: todo.TaskResponse.task:type_name -> todo.Task
	0, // 1: todo.ListTasksResponse.tasks:type_name -> todo.Task
	1, // 2: todo.TodoService.CreateTask:input_type -> todo.CreateTaskRequest
	2, // 3: todo.TodoService.GetTask:input_type -> todo.GetTaskRequest
	3, // 4: todo.TodoService.UpdateTask:input_type -> todo.UpdateTaskRequest
	4, // 5: todo.TodoService.DeleteTask:input_type -> todo.DeleteTaskRequest
	7, // 6: todo.TodoService.ListTasks:input_type -> todo.ListTasksRequest
	6, // 7: todo.TodoService.CreateTask:output_type -> todo.TaskResponse
	6, // 8: todo.TodoService.GetTask:output_type -> todo.TaskResponse
	6, // 9: todo.TodoService.UpdateTask:output_type -> todo.TaskResponse
	5, // 10: todo.TodoService.DeleteTask:output_type -> todo.DeleteTaskResponse
	8, // 11: todo.TodoService.ListTasks:output_type -> todo.ListTasksResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_todo_proto_init() }
func file_todo_proto_init() {
	if File_todo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_todo_proto_rawDesc), len(file_todo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_proto_goTypes,
		DependencyIndexes: file_todo_proto_depIdxs,
		MessageInfos:      file_todo_proto_msgTypes,
	}.Build()
	File_todo_proto = out.File
	file_todo_proto_goTypes = nil
	file_todo_proto_depIdxs = nil
}
