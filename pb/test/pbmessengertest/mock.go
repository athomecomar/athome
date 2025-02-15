// Code generated by MockGen. DO NOT EDIT.
// Source: ../pbmessenger/messenger.pb.go

// Package pbmessengertest is a generated GoMock package.
package pbmessengertest

import (
	context "context"
	pbmessenger "github.com/athomecomar/athome/pb/pbmessenger"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockMessagesClient is a mock of MessagesClient interface
type MockMessagesClient struct {
	ctrl     *gomock.Controller
	recorder *MockMessagesClientMockRecorder
}

// MockMessagesClientMockRecorder is the mock recorder for MockMessagesClient
type MockMessagesClientMockRecorder struct {
	mock *MockMessagesClient
}

// NewMockMessagesClient creates a new mock instance
func NewMockMessagesClient(ctrl *gomock.Controller) *MockMessagesClient {
	mock := &MockMessagesClient{ctrl: ctrl}
	mock.recorder = &MockMessagesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessagesClient) EXPECT() *MockMessagesClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMessagesClient) Create(ctx context.Context, in *pbmessenger.CreateRequest, opts ...grpc.CallOption) (*pbmessenger.CreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*pbmessenger.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockMessagesClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessagesClient)(nil).Create), varargs...)
}

// Retrieve mocks base method
func (m *MockMessagesClient) Retrieve(ctx context.Context, in *pbmessenger.RetrieveRequest, opts ...grpc.CallOption) (*pbmessenger.Message, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Retrieve", varargs...)
	ret0, _ := ret[0].(*pbmessenger.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve
func (mr *MockMessagesClientMockRecorder) Retrieve(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockMessagesClient)(nil).Retrieve), varargs...)
}

// RetrieveMany mocks base method
func (m *MockMessagesClient) RetrieveMany(ctx context.Context, in *pbmessenger.RetrieveManyRequest, opts ...grpc.CallOption) (*pbmessenger.RetrieveManyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveMany", varargs...)
	ret0, _ := ret[0].(*pbmessenger.RetrieveManyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMany indicates an expected call of RetrieveMany
func (mr *MockMessagesClientMockRecorder) RetrieveMany(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMany", reflect.TypeOf((*MockMessagesClient)(nil).RetrieveMany), varargs...)
}

// SetReceived mocks base method
func (m *MockMessagesClient) SetReceived(ctx context.Context, in *pbmessenger.UpdateStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetReceived", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetReceived indicates an expected call of SetReceived
func (mr *MockMessagesClientMockRecorder) SetReceived(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReceived", reflect.TypeOf((*MockMessagesClient)(nil).SetReceived), varargs...)
}

// SetSeen mocks base method
func (m *MockMessagesClient) SetSeen(ctx context.Context, in *pbmessenger.UpdateStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetSeen", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSeen indicates an expected call of SetSeen
func (mr *MockMessagesClientMockRecorder) SetSeen(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSeen", reflect.TypeOf((*MockMessagesClient)(nil).SetSeen), varargs...)
}

// Reply mocks base method
func (m *MockMessagesClient) Reply(ctx context.Context, in *pbmessenger.ReplyRequest, opts ...grpc.CallOption) (*pbmessenger.CreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Reply", varargs...)
	ret0, _ := ret[0].(*pbmessenger.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reply indicates an expected call of Reply
func (mr *MockMessagesClientMockRecorder) Reply(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reply", reflect.TypeOf((*MockMessagesClient)(nil).Reply), varargs...)
}

// MockMessagesServer is a mock of MessagesServer interface
type MockMessagesServer struct {
	ctrl     *gomock.Controller
	recorder *MockMessagesServerMockRecorder
}

// MockMessagesServerMockRecorder is the mock recorder for MockMessagesServer
type MockMessagesServerMockRecorder struct {
	mock *MockMessagesServer
}

// NewMockMessagesServer creates a new mock instance
func NewMockMessagesServer(ctrl *gomock.Controller) *MockMessagesServer {
	mock := &MockMessagesServer{ctrl: ctrl}
	mock.recorder = &MockMessagesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessagesServer) EXPECT() *MockMessagesServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockMessagesServer) Create(arg0 context.Context, arg1 *pbmessenger.CreateRequest) (*pbmessenger.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*pbmessenger.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockMessagesServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessagesServer)(nil).Create), arg0, arg1)
}

// Retrieve mocks base method
func (m *MockMessagesServer) Retrieve(arg0 context.Context, arg1 *pbmessenger.RetrieveRequest) (*pbmessenger.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retrieve", arg0, arg1)
	ret0, _ := ret[0].(*pbmessenger.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Retrieve indicates an expected call of Retrieve
func (mr *MockMessagesServerMockRecorder) Retrieve(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retrieve", reflect.TypeOf((*MockMessagesServer)(nil).Retrieve), arg0, arg1)
}

// RetrieveMany mocks base method
func (m *MockMessagesServer) RetrieveMany(arg0 context.Context, arg1 *pbmessenger.RetrieveManyRequest) (*pbmessenger.RetrieveManyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveMany", arg0, arg1)
	ret0, _ := ret[0].(*pbmessenger.RetrieveManyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveMany indicates an expected call of RetrieveMany
func (mr *MockMessagesServerMockRecorder) RetrieveMany(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveMany", reflect.TypeOf((*MockMessagesServer)(nil).RetrieveMany), arg0, arg1)
}

// SetReceived mocks base method
func (m *MockMessagesServer) SetReceived(arg0 context.Context, arg1 *pbmessenger.UpdateStatusRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReceived", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetReceived indicates an expected call of SetReceived
func (mr *MockMessagesServerMockRecorder) SetReceived(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReceived", reflect.TypeOf((*MockMessagesServer)(nil).SetReceived), arg0, arg1)
}

// SetSeen mocks base method
func (m *MockMessagesServer) SetSeen(arg0 context.Context, arg1 *pbmessenger.UpdateStatusRequest) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSeen", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSeen indicates an expected call of SetSeen
func (mr *MockMessagesServerMockRecorder) SetSeen(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSeen", reflect.TypeOf((*MockMessagesServer)(nil).SetSeen), arg0, arg1)
}

// Reply mocks base method
func (m *MockMessagesServer) Reply(arg0 context.Context, arg1 *pbmessenger.ReplyRequest) (*pbmessenger.CreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reply", arg0, arg1)
	ret0, _ := ret[0].(*pbmessenger.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reply indicates an expected call of Reply
func (mr *MockMessagesServerMockRecorder) Reply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reply", reflect.TypeOf((*MockMessagesServer)(nil).Reply), arg0, arg1)
}
