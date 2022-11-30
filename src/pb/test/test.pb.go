// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: test.proto

package testpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TimeToDo           int32  `protobuf:"varint,2,opt,name=time_to_do,json=timeToDo,proto3" json:"time_to_do,omitempty"`
	TimeStart          string `protobuf:"bytes,3,opt,name=time_start,json=timeStart,proto3" json:"time_start,omitempty"`
	TimeEnd            string `protobuf:"bytes,4,opt,name=time_end,json=timeEnd,proto3" json:"time_end,omitempty"`
	DoOnce             bool   `protobuf:"varint,5,opt,name=do_once,json=doOnce,proto3" json:"do_once,omitempty"`
	Password           string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	PreventCheat       uint32 `protobuf:"varint,7,opt,name=prevent_cheat,json=preventCheat,proto3" json:"prevent_cheat,omitempty"`
	IsAuthenticateUser bool   `protobuf:"varint,8,opt,name=is_authenticate_user,json=isAuthenticateUser,proto3" json:"is_authenticate_user,omitempty"`
	ShowMark           uint32 `protobuf:"varint,9,opt,name=show_mark,json=showMark,proto3" json:"show_mark,omitempty"`
	ShowAnswer         uint32 `protobuf:"varint,10,opt,name=show_answer,json=showAnswer,proto3" json:"show_answer,omitempty"`
}

func (x *CreateTestRequest) Reset() {
	*x = CreateTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestRequest) ProtoMessage() {}

func (x *CreateTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestRequest.ProtoReflect.Descriptor instead.
func (*CreateTestRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTestRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTestRequest) GetTimeToDo() int32 {
	if x != nil {
		return x.TimeToDo
	}
	return 0
}

func (x *CreateTestRequest) GetTimeStart() string {
	if x != nil {
		return x.TimeStart
	}
	return ""
}

func (x *CreateTestRequest) GetTimeEnd() string {
	if x != nil {
		return x.TimeEnd
	}
	return ""
}

func (x *CreateTestRequest) GetDoOnce() bool {
	if x != nil {
		return x.DoOnce
	}
	return false
}

func (x *CreateTestRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateTestRequest) GetPreventCheat() uint32 {
	if x != nil {
		return x.PreventCheat
	}
	return 0
}

func (x *CreateTestRequest) GetIsAuthenticateUser() bool {
	if x != nil {
		return x.IsAuthenticateUser
	}
	return false
}

func (x *CreateTestRequest) GetShowMark() uint32 {
	if x != nil {
		return x.ShowMark
	}
	return 0
}

func (x *CreateTestRequest) GetShowAnswer() uint32 {
	if x != nil {
		return x.ShowAnswer
	}
	return 0
}

type CreateTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateTestResponse) Reset() {
	*x = CreateTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestResponse) ProtoMessage() {}

func (x *CreateTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestResponse.ProtoReflect.Descriptor instead.
func (*CreateTestResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MultipleChoiceAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestMultipleChoiceId int64   `protobuf:"varint,1,opt,name=test_multiple_choice_id,json=testMultipleChoiceId,proto3" json:"test_multiple_choice_id,omitempty"`
	Answer               string  `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Score                float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	Type                 int32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *MultipleChoiceAnswer) Reset() {
	*x = MultipleChoiceAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultipleChoiceAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultipleChoiceAnswer) ProtoMessage() {}

func (x *MultipleChoiceAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultipleChoiceAnswer.ProtoReflect.Descriptor instead.
func (*MultipleChoiceAnswer) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *MultipleChoiceAnswer) GetTestMultipleChoiceId() int64 {
	if x != nil {
		return x.TestMultipleChoiceId
	}
	return 0
}

func (x *MultipleChoiceAnswer) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *MultipleChoiceAnswer) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *MultipleChoiceAnswer) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type TestMultipleChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string                  `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Score    float32                 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	Answers  []*MultipleChoiceAnswer `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *TestMultipleChoice) Reset() {
	*x = TestMultipleChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMultipleChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMultipleChoice) ProtoMessage() {}

func (x *TestMultipleChoice) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMultipleChoice.ProtoReflect.Descriptor instead.
func (*TestMultipleChoice) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *TestMultipleChoice) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *TestMultipleChoice) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *TestMultipleChoice) GetAnswers() []*MultipleChoiceAnswer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type CreateTestContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestId         int64               `protobuf:"varint,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	Typeable       int64               `protobuf:"varint,2,opt,name=typeable,proto3" json:"typeable,omitempty"`
	MultipleChoice *TestMultipleChoice `protobuf:"bytes,3,opt,name=multiple_choice,json=multipleChoice,proto3" json:"multiple_choice,omitempty"`
}

func (x *CreateTestContentRequest) Reset() {
	*x = CreateTestContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestContentRequest) ProtoMessage() {}

func (x *CreateTestContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestContentRequest.ProtoReflect.Descriptor instead.
func (*CreateTestContentRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTestContentRequest) GetTestId() int64 {
	if x != nil {
		return x.TestId
	}
	return 0
}

func (x *CreateTestContentRequest) GetTypeable() int64 {
	if x != nil {
		return x.Typeable
	}
	return 0
}

func (x *CreateTestContentRequest) GetMultipleChoice() *TestMultipleChoice {
	if x != nil {
		return x.MultipleChoice
	}
	return nil
}

type CreateTestContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateTestContentResponse) Reset() {
	*x = CreateTestContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestContentResponse) ProtoMessage() {}

func (x *CreateTestContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestContentResponse.ProtoReflect.Descriptor instead.
func (*CreateTestContentResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTestContentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc9, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x45,
	0x6e, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x6f, 0x5f, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x6f, 0x4f, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x70, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x65, 0x61, 0x74, 0x12, 0x30, 0x0a, 0x14,
	0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x68, 0x6f, 0x77, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8f, 0x01, 0x0a,
	0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x17, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7d,
	0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x68,
	0x6f, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x92, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x41, 0x0a, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x52, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x22, 0x35, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd8, 0x01, 0x0a, 0x0b, 0x54, 0x65,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x72, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x3a, 0x01, 0x2a, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_test_proto_goTypes = []interface{}{
	(*CreateTestRequest)(nil),         // 0: test.CreateTestRequest
	(*CreateTestResponse)(nil),        // 1: test.CreateTestResponse
	(*MultipleChoiceAnswer)(nil),      // 2: test.MultipleChoiceAnswer
	(*TestMultipleChoice)(nil),        // 3: test.TestMultipleChoice
	(*CreateTestContentRequest)(nil),  // 4: test.CreateTestContentRequest
	(*CreateTestContentResponse)(nil), // 5: test.CreateTestContentResponse
}
var file_test_proto_depIdxs = []int32{
	2, // 0: test.TestMultipleChoice.answers:type_name -> test.MultipleChoiceAnswer
	3, // 1: test.CreateTestContentRequest.multiple_choice:type_name -> test.TestMultipleChoice
	0, // 2: test.TestService.Create:input_type -> test.CreateTestRequest
	4, // 3: test.TestService.CreateContent:input_type -> test.CreateTestContentRequest
	1, // 4: test.TestService.Create:output_type -> test.CreateTestResponse
	5, // 5: test.TestService.CreateContent:output_type -> test.CreateTestContentResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestRequest); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestResponse); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultipleChoiceAnswer); i {
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
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMultipleChoice); i {
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
		file_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestContentRequest); i {
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
		file_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestContentResponse); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
