// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/athomecomar/athome/pb/pbimages (interfaces: Images_CreateImageClient,ImagesClient)

// Package pbimagestest is a generated GoMock package.
package pbimagestest

import (
	context "context"
	pbimages "github.com/athomecomar/athome/pb/pbimages"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

// MockImages_CreateImageClient is a mock of Images_CreateImageClient interface
type MockImages_CreateImageClient struct {
	ctrl     *gomock.Controller
	recorder *MockImages_CreateImageClientMockRecorder
}

// MockImages_CreateImageClientMockRecorder is the mock recorder for MockImages_CreateImageClient
type MockImages_CreateImageClientMockRecorder struct {
	mock *MockImages_CreateImageClient
}

// NewMockImages_CreateImageClient creates a new mock instance
func NewMockImages_CreateImageClient(ctrl *gomock.Controller) *MockImages_CreateImageClient {
	mock := &MockImages_CreateImageClient{ctrl: ctrl}
	mock.recorder = &MockImages_CreateImageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImages_CreateImageClient) EXPECT() *MockImages_CreateImageClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method
func (m *MockImages_CreateImageClient) CloseAndRecv() (*pbimages.CreateImageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*pbimages.CreateImageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv
func (mr *MockImages_CreateImageClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockImages_CreateImageClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method
func (m *MockImages_CreateImageClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockImages_CreateImageClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockImages_CreateImageClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockImages_CreateImageClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockImages_CreateImageClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockImages_CreateImageClient)(nil).Context))
}

// Header mocks base method
func (m *MockImages_CreateImageClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockImages_CreateImageClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockImages_CreateImageClient)(nil).Header))
}

// RecvMsg mocks base method
func (m *MockImages_CreateImageClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockImages_CreateImageClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockImages_CreateImageClient)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockImages_CreateImageClient) Send(arg0 *pbimages.CreateImageRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockImages_CreateImageClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockImages_CreateImageClient)(nil).Send), arg0)
}

// SendMsg mocks base method
func (m *MockImages_CreateImageClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockImages_CreateImageClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockImages_CreateImageClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockImages_CreateImageClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockImages_CreateImageClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockImages_CreateImageClient)(nil).Trailer))
}

// MockImagesClient is a mock of ImagesClient interface
type MockImagesClient struct {
	ctrl     *gomock.Controller
	recorder *MockImagesClientMockRecorder
}

// MockImagesClientMockRecorder is the mock recorder for MockImagesClient
type MockImagesClientMockRecorder struct {
	mock *MockImagesClient
}

// NewMockImagesClient creates a new mock instance
func NewMockImagesClient(ctrl *gomock.Controller) *MockImagesClient {
	mock := &MockImagesClient{ctrl: ctrl}
	mock.recorder = &MockImagesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImagesClient) EXPECT() *MockImagesClientMockRecorder {
	return m.recorder
}

// ChangeEntityImages mocks base method
func (m *MockImagesClient) ChangeEntityImages(arg0 context.Context, arg1 *pbimages.ChangeEntityImagesRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeEntityImages", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeEntityImages indicates an expected call of ChangeEntityImages
func (mr *MockImagesClientMockRecorder) ChangeEntityImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEntityImages", reflect.TypeOf((*MockImagesClient)(nil).ChangeEntityImages), varargs...)
}

// CloneImages mocks base method
func (m *MockImagesClient) CloneImages(arg0 context.Context, arg1 *pbimages.CloneImagesRequest, arg2 ...grpc.CallOption) (*pbimages.CloneImagesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CloneImages", varargs...)
	ret0, _ := ret[0].(*pbimages.CloneImagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneImages indicates an expected call of CloneImages
func (mr *MockImagesClientMockRecorder) CloneImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneImages", reflect.TypeOf((*MockImagesClient)(nil).CloneImages), varargs...)
}

// CreateImage mocks base method
func (m *MockImagesClient) CreateImage(arg0 context.Context, arg1 ...grpc.CallOption) (pbimages.Images_CreateImageClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateImage", varargs...)
	ret0, _ := ret[0].(pbimages.Images_CreateImageClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateImage indicates an expected call of CreateImage
func (mr *MockImagesClientMockRecorder) CreateImage(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImage", reflect.TypeOf((*MockImagesClient)(nil).CreateImage), varargs...)
}

// DeleteImages mocks base method
func (m *MockImagesClient) DeleteImages(arg0 context.Context, arg1 *pbimages.DeleteImagesRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteImages", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteImages indicates an expected call of DeleteImages
func (mr *MockImagesClientMockRecorder) DeleteImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImages", reflect.TypeOf((*MockImagesClient)(nil).DeleteImages), varargs...)
}

// RetrieveImages mocks base method
func (m *MockImagesClient) RetrieveImages(arg0 context.Context, arg1 *pbimages.RetrieveImagesRequest, arg2 ...grpc.CallOption) (*pbimages.RetrieveImagesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveImages", varargs...)
	ret0, _ := ret[0].(*pbimages.RetrieveImagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveImages indicates an expected call of RetrieveImages
func (mr *MockImagesClientMockRecorder) RetrieveImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveImages", reflect.TypeOf((*MockImagesClient)(nil).RetrieveImages), varargs...)
}
