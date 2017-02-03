// Code generated by protoc-gen-go.
// source: google/genomics/v1/operations.proto
// DO NOT EDIT!

package genomics

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf5 "github.com/golang/protobuf/ptypes/any"
import google_protobuf6 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Metadata describing an [Operation][google.longrunning.Operation].
type OperationMetadata struct {
	// The Google Cloud Project in which the job is scoped.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	// The time at which the job was submitted to the Genomics service.
	CreateTime *google_protobuf6.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// The time at which the job began to run.
	StartTime *google_protobuf6.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The time at which the job stopped running.
	EndTime *google_protobuf6.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// The original request that started the operation. Note that this will be in
	// current version of the API. If the operation was started with v1beta2 API
	// and a GetOperation is performed on v1 API, a v1 request will be returned.
	Request *google_protobuf5.Any `protobuf:"bytes,5,opt,name=request" json:"request,omitempty"`
	// Optional event messages that were generated during the job's execution.
	// This also contains any warnings that were generated during import
	// or export.
	Events []*OperationEvent `protobuf:"bytes,6,rep,name=events" json:"events,omitempty"`
	// Optionally provided by the caller when submitting the request that creates
	// the operation.
	ClientId string `protobuf:"bytes,7,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Runtime metadata on this Operation.
	RuntimeMetadata *google_protobuf5.Any `protobuf:"bytes,8,opt,name=runtime_metadata,json=runtimeMetadata" json:"runtime_metadata,omitempty"`
}

func (m *OperationMetadata) Reset()                    { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string            { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()               {}
func (*OperationMetadata) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *OperationMetadata) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *OperationMetadata) GetCreateTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetStartTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationMetadata) GetEndTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationMetadata) GetRequest() *google_protobuf5.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *OperationMetadata) GetEvents() []*OperationEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *OperationMetadata) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OperationMetadata) GetRuntimeMetadata() *google_protobuf5.Any {
	if m != nil {
		return m.RuntimeMetadata
	}
	return nil
}

// An event that occurred during an [Operation][google.longrunning.Operation].
type OperationEvent struct {
	// Optional time of when event started.
	StartTime *google_protobuf6.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// Optional time of when event finished. An event can have a start time and no
	// finish time. If an event has a finish time, there must be a start time.
	EndTime *google_protobuf6.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Required description of event.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *OperationEvent) Reset()                    { *m = OperationEvent{} }
func (m *OperationEvent) String() string            { return proto.CompactTextString(m) }
func (*OperationEvent) ProtoMessage()               {}
func (*OperationEvent) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *OperationEvent) GetStartTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationEvent) GetEndTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.genomics.v1.OperationMetadata")
	proto.RegisterType((*OperationEvent)(nil), "google.genomics.v1.OperationEvent")
}

func init() { proto.RegisterFile("google/genomics/v1/operations.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0xab, 0xd3, 0x40,
	0x10, 0xc6, 0xd9, 0x57, 0x6d, 0x9b, 0x09, 0xf8, 0x74, 0x11, 0x89, 0x55, 0x31, 0xd4, 0x4b, 0x4f,
	0x1b, 0xde, 0x13, 0x0f, 0xbe, 0x1e, 0xc4, 0x82, 0x87, 0x1e, 0xc4, 0x12, 0x3c, 0x79, 0x29, 0xdb,
	0x64, 0x0c, 0x5b, 0x9a, 0xdd, 0xb8, 0xbb, 0x2d, 0xf4, 0xff, 0xf1, 0xe6, 0x3f, 0xe8, 0x51, 0xb2,
	0xd9, 0x2d, 0xb5, 0x15, 0x2b, 0xde, 0x92, 0x99, 0xef, 0xb7, 0xf3, 0xf1, 0x0d, 0x03, 0xaf, 0x2a,
	0xa5, 0xaa, 0x0d, 0x66, 0x15, 0x4a, 0x55, 0x8b, 0xc2, 0x64, 0xbb, 0x9b, 0x4c, 0x35, 0xa8, 0xb9,
	0x15, 0x4a, 0x1a, 0xd6, 0x68, 0x65, 0x15, 0xa5, 0x9d, 0x88, 0x05, 0x11, 0xdb, 0xdd, 0x8c, 0x9e,
	0x7b, 0x90, 0x37, 0x22, 0xe3, 0x52, 0x2a, 0x7b, 0x4c, 0x8c, 0x9e, 0xfa, 0xae, 0xfb, 0x5b, 0x6d,
	0xbf, 0x66, 0x5c, 0xee, 0x7d, 0xeb, 0xe5, 0x69, 0xcb, 0x8a, 0x1a, 0x8d, 0xe5, 0x75, 0xd3, 0x09,
	0xc6, 0x3f, 0x7a, 0xf0, 0xe8, 0x53, 0xb0, 0xf0, 0x11, 0x2d, 0x2f, 0xb9, 0xe5, 0xf4, 0x05, 0x40,
	0xa3, 0xd5, 0x1a, 0x0b, 0xbb, 0x14, 0x65, 0x42, 0x52, 0x32, 0x89, 0xf2, 0xc8, 0x57, 0xe6, 0x25,
	0x9d, 0x42, 0x5c, 0x68, 0xe4, 0x16, 0x97, 0xed, 0x73, 0xc9, 0x55, 0x4a, 0x26, 0xf1, 0xed, 0x88,
	0x79, 0xe3, 0x61, 0x16, 0xfb, 0x1c, 0x66, 0xe5, 0xd0, 0xc9, 0xdb, 0x02, 0x7d, 0x0b, 0x60, 0x2c,
	0xd7, 0xb6, 0x63, 0x7b, 0x17, 0xd9, 0xc8, 0xa9, 0x1d, 0xfa, 0x06, 0x86, 0x28, 0xcb, 0x0e, 0xbc,
	0x77, 0x11, 0x1c, 0xa0, 0x2c, 0x1d, 0xc6, 0x60, 0xa0, 0xf1, 0xdb, 0x16, 0x8d, 0x4d, 0xee, 0x3b,
	0xea, 0xf1, 0x19, 0xf5, 0x5e, 0xee, 0xf3, 0x20, 0xa2, 0x77, 0xd0, 0xc7, 0x1d, 0x4a, 0x6b, 0x92,
	0x7e, 0xda, 0x9b, 0xc4, 0xb7, 0x63, 0x76, 0xbe, 0x12, 0x76, 0x08, 0xed, 0x43, 0x2b, 0xcd, 0x3d,
	0x41, 0x9f, 0x41, 0x54, 0x6c, 0x04, 0x4a, 0x17, 0xdc, 0xc0, 0x05, 0x37, 0xec, 0x0a, 0xf3, 0x92,
	0xbe, 0x83, 0x87, 0x7a, 0x2b, 0x5b, 0xfb, 0xcb, 0xda, 0x47, 0x9d, 0x0c, 0xff, 0xe2, 0xe8, 0xda,
	0xab, 0xc3, 0x5e, 0xc6, 0xdf, 0x09, 0x3c, 0xf8, 0x7d, 0xf0, 0x49, 0x9c, 0xe4, 0x7f, 0xe3, 0xbc,
	0xfa, 0xf7, 0x38, 0x53, 0x88, 0x4b, 0x34, 0x85, 0x16, 0x4d, 0xeb, 0xc2, 0x6d, 0x30, 0xca, 0x8f,
	0x4b, 0xb3, 0x35, 0x3c, 0x29, 0x54, 0xfd, 0x87, 0xd4, 0x66, 0xd7, 0x07, 0xf7, 0x66, 0xd1, 0x8e,
	0x58, 0x90, 0x2f, 0x77, 0x41, 0xa6, 0x36, 0x5c, 0x56, 0x4c, 0xe9, 0xaa, 0xbd, 0x0f, 0x67, 0x20,
	0xeb, 0x5a, 0xbc, 0x11, 0xe6, 0xf8, 0x66, 0xa6, 0xe1, 0xfb, 0x27, 0x21, 0xab, 0xbe, 0x53, 0xbe,
	0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x39, 0xad, 0x2f, 0x5c, 0x03, 0x00, 0x00,
}