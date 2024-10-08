// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: service/proto/v1/jobworker.proto

package proto

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

//	{
//	    "command_args": [
//	        "-ltra"
//	    ],
//	    "command_name": "ls"
//	}
type StartJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandName string   `protobuf:"bytes,1,opt,name=command_name,json=commandName,proto3" json:"command_name,omitempty"`
	CommandArgs []string `protobuf:"bytes,2,rep,name=command_args,json=commandArgs,proto3" json:"command_args,omitempty"`
}

func (x *StartJobRequest) Reset() {
	*x = StartJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_jobworker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobRequest) ProtoMessage() {}

func (x *StartJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_jobworker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobRequest.ProtoReflect.Descriptor instead.
func (*StartJobRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_jobworker_proto_rawDescGZIP(), []int{0}
}

func (x *StartJobRequest) GetCommandName() string {
	if x != nil {
		return x.CommandName
	}
	return ""
}

func (x *StartJobRequest) GetCommandArgs() []string {
	if x != nil {
		return x.CommandArgs
	}
	return nil
}

//	{
//	    "job_id": "761db04c-0150-4f0b-a6fd-5cab9b9a48bf",
//	}
type StartJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId   string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StartJobResponse) Reset() {
	*x = StartJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_jobworker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobResponse) ProtoMessage() {}

func (x *StartJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_jobworker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobResponse.ProtoReflect.Descriptor instead.
func (*StartJobResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_jobworker_proto_rawDescGZIP(), []int{1}
}

func (x *StartJobResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *StartJobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StopJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *StopJobRequest) Reset() {
	*x = StopJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_jobworker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopJobRequest) ProtoMessage() {}

func (x *StopJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_jobworker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopJobRequest.ProtoReflect.Descriptor instead.
func (*StopJobRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_jobworker_proto_rawDescGZIP(), []int{2}
}

func (x *StopJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

//	{
//	    "success": "true",
//	}
type StopJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StopJobResponse) Reset() {
	*x = StopJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_jobworker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopJobResponse) ProtoMessage() {}

func (x *StopJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_jobworker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopJobResponse.ProtoReflect.Descriptor instead.
func (*StopJobResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_jobworker_proto_rawDescGZIP(), []int{3}
}

func (x *StopJobResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *StopJobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type QueryJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *QueryJobRequest) Reset() {
	*x = QueryJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_jobworker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryJobRequest) ProtoMessage() {}

func (x *QueryJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_jobworker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryJobRequest.ProtoReflect.Descriptor instead.
func (*QueryJobRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_jobworker_proto_rawDescGZIP(), []int{4}
}

func (x *QueryJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type QueryJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                      // Indicates if the job is running or killed|exited
	Pid      int32  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`                           // Process ID of the job
	ExitCode int32  `protobuf:"varint,3,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"` // Exit code if the job has exited
	Signal   int32  `protobuf:"varint,4,opt,name=signal,proto3" json:"signal,omitempty"`                     // Signal used to terminate the job, if applicable
	Message  string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *QueryJobResponse) Reset() {
	*x = QueryJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_jobworker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryJobResponse) ProtoMessage() {}

func (x *QueryJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_jobworker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryJobResponse.ProtoReflect.Descriptor instead.
func (*QueryJobResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_jobworker_proto_rawDescGZIP(), []int{5}
}

func (x *QueryJobResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *QueryJobResponse) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *QueryJobResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *QueryJobResponse) GetSignal() int32 {
	if x != nil {
		return x.Signal
	}
	return 0
}

func (x *QueryJobResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StreamLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *StreamLogsRequest) Reset() {
	*x = StreamLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_jobworker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamLogsRequest) ProtoMessage() {}

func (x *StreamLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_jobworker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamLogsRequest.ProtoReflect.Descriptor instead.
func (*StreamLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_jobworker_proto_rawDescGZIP(), []int{6}
}

func (x *StreamLogsRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

// sample streaming response
//
//	{
//	    "message": "2024-08-26T11:35:21Z [INFO] Server started on port 8080",
//	}
type StreamLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StreamLogsResponse) Reset() {
	*x = StreamLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_v1_jobworker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamLogsResponse) ProtoMessage() {}

func (x *StreamLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_v1_jobworker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamLogsResponse.ProtoReflect.Descriptor instead.
func (*StreamLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_v1_jobworker_proto_rawDescGZIP(), []int{7}
}

func (x *StreamLogsResponse) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_service_proto_v1_jobworker_proto protoreflect.FileDescriptor

var file_service_proto_v1_jobworker_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x6a, 0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0x57, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x41, 0x72, 0x67, 0x73, 0x22, 0x43, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27,
	0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x70, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28,
	0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x10, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62,
	0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xbc, 0x02, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x12, 0x49, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x1d, 0x2e, 0x6a,
	0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6a, 0x6f,
	0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x53,
	0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x12, 0x1c, 0x2e, 0x6a, 0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6a, 0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x12,
	0x1d, 0x2e, 0x6a, 0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6a, 0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x6a,
	0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x6a, 0x6f, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x73, 0x74, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x2d, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_proto_v1_jobworker_proto_rawDescOnce sync.Once
	file_service_proto_v1_jobworker_proto_rawDescData = file_service_proto_v1_jobworker_proto_rawDesc
)

func file_service_proto_v1_jobworker_proto_rawDescGZIP() []byte {
	file_service_proto_v1_jobworker_proto_rawDescOnce.Do(func() {
		file_service_proto_v1_jobworker_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_v1_jobworker_proto_rawDescData)
	})
	return file_service_proto_v1_jobworker_proto_rawDescData
}

var file_service_proto_v1_jobworker_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_proto_v1_jobworker_proto_goTypes = []any{
	(*StartJobRequest)(nil),    // 0: jobworker.v1.StartJobRequest
	(*StartJobResponse)(nil),   // 1: jobworker.v1.StartJobResponse
	(*StopJobRequest)(nil),     // 2: jobworker.v1.StopJobRequest
	(*StopJobResponse)(nil),    // 3: jobworker.v1.StopJobResponse
	(*QueryJobRequest)(nil),    // 4: jobworker.v1.QueryJobRequest
	(*QueryJobResponse)(nil),   // 5: jobworker.v1.QueryJobResponse
	(*StreamLogsRequest)(nil),  // 6: jobworker.v1.StreamLogsRequest
	(*StreamLogsResponse)(nil), // 7: jobworker.v1.StreamLogsResponse
}
var file_service_proto_v1_jobworker_proto_depIdxs = []int32{
	0, // 0: jobworker.v1.JobWorker.StartJob:input_type -> jobworker.v1.StartJobRequest
	2, // 1: jobworker.v1.JobWorker.StopJob:input_type -> jobworker.v1.StopJobRequest
	4, // 2: jobworker.v1.JobWorker.QueryJob:input_type -> jobworker.v1.QueryJobRequest
	6, // 3: jobworker.v1.JobWorker.StreamLogs:input_type -> jobworker.v1.StreamLogsRequest
	1, // 4: jobworker.v1.JobWorker.StartJob:output_type -> jobworker.v1.StartJobResponse
	3, // 5: jobworker.v1.JobWorker.StopJob:output_type -> jobworker.v1.StopJobResponse
	5, // 6: jobworker.v1.JobWorker.QueryJob:output_type -> jobworker.v1.QueryJobResponse
	7, // 7: jobworker.v1.JobWorker.StreamLogs:output_type -> jobworker.v1.StreamLogsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_v1_jobworker_proto_init() }
func file_service_proto_v1_jobworker_proto_init() {
	if File_service_proto_v1_jobworker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_v1_jobworker_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StartJobRequest); i {
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
		file_service_proto_v1_jobworker_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StartJobResponse); i {
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
		file_service_proto_v1_jobworker_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StopJobRequest); i {
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
		file_service_proto_v1_jobworker_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StopJobResponse); i {
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
		file_service_proto_v1_jobworker_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*QueryJobRequest); i {
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
		file_service_proto_v1_jobworker_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*QueryJobResponse); i {
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
		file_service_proto_v1_jobworker_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*StreamLogsRequest); i {
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
		file_service_proto_v1_jobworker_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*StreamLogsResponse); i {
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
			RawDescriptor: file_service_proto_v1_jobworker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_v1_jobworker_proto_goTypes,
		DependencyIndexes: file_service_proto_v1_jobworker_proto_depIdxs,
		MessageInfos:      file_service_proto_v1_jobworker_proto_msgTypes,
	}.Build()
	File_service_proto_v1_jobworker_proto = out.File
	file_service_proto_v1_jobworker_proto_rawDesc = nil
	file_service_proto_v1_jobworker_proto_goTypes = nil
	file_service_proto_v1_jobworker_proto_depIdxs = nil
}
