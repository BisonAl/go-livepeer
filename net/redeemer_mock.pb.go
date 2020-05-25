// Code generated by MockGen. DO NOT EDIT.
// Source: redeemer.pb.go

// Package net is a generated GoMock package.
package net

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockTicketRedeemerClient is a mock of TicketRedeemerClient interface.
type MockTicketRedeemerClient struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRedeemerClientMockRecorder
}

// MockTicketRedeemerClientMockRecorder is the mock recorder for MockTicketRedeemerClient.
type MockTicketRedeemerClientMockRecorder struct {
	mock *MockTicketRedeemerClient
}

// NewMockTicketRedeemerClient creates a new mock instance.
func NewMockTicketRedeemerClient(ctrl *gomock.Controller) *MockTicketRedeemerClient {
	mock := &MockTicketRedeemerClient{ctrl: ctrl}
	mock.recorder = &MockTicketRedeemerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketRedeemerClient) EXPECT() *MockTicketRedeemerClientMockRecorder {
	return m.recorder
}

// QueueTicket mocks base method.
func (m *MockTicketRedeemerClient) QueueTicket(ctx context.Context, in *Ticket, opts ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueueTicket", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueTicket indicates an expected call of QueueTicket.
func (mr *MockTicketRedeemerClientMockRecorder) QueueTicket(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueTicket", reflect.TypeOf((*MockTicketRedeemerClient)(nil).QueueTicket), varargs...)
}

// MonitorMaxFloat mocks base method.
func (m *MockTicketRedeemerClient) MonitorMaxFloat(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (TicketRedeemer_MonitorMaxFloatClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MonitorMaxFloat", varargs...)
	ret0, _ := ret[0].(TicketRedeemer_MonitorMaxFloatClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MonitorMaxFloat indicates an expected call of MonitorMaxFloat.
func (mr *MockTicketRedeemerClientMockRecorder) MonitorMaxFloat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorMaxFloat", reflect.TypeOf((*MockTicketRedeemerClient)(nil).MonitorMaxFloat), varargs...)
}

// MockTicketRedeemer_MonitorMaxFloatClient is a mock of TicketRedeemer_MonitorMaxFloatClient interface.
type MockTicketRedeemer_MonitorMaxFloatClient struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder
}

// MockTicketRedeemer_MonitorMaxFloatClientMockRecorder is the mock recorder for MockTicketRedeemer_MonitorMaxFloatClient.
type MockTicketRedeemer_MonitorMaxFloatClientMockRecorder struct {
	mock *MockTicketRedeemer_MonitorMaxFloatClient
}

// NewMockTicketRedeemer_MonitorMaxFloatClient creates a new mock instance.
func NewMockTicketRedeemer_MonitorMaxFloatClient(ctrl *gomock.Controller) *MockTicketRedeemer_MonitorMaxFloatClient {
	mock := &MockTicketRedeemer_MonitorMaxFloatClient{ctrl: ctrl}
	mock.recorder = &MockTicketRedeemer_MonitorMaxFloatClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketRedeemer_MonitorMaxFloatClient) EXPECT() *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder {
	return m.recorder
}

// Recv mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatClient) Recv() (*MaxFloatUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*MaxFloatUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatClient)(nil).Recv))
}

// Header mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatClient)(nil).Header))
}

// Trailer mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatClient)(nil).Trailer))
}

// CloseSend mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatClient)(nil).Context))
}

// SendMsg mocks base method.
func (m_2 *MockTicketRedeemer_MonitorMaxFloatClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method.
func (m_2 *MockTicketRedeemer_MonitorMaxFloatClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockTicketRedeemer_MonitorMaxFloatClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatClient)(nil).RecvMsg), m)
}

// MockTicketRedeemerServer is a mock of TicketRedeemerServer interface.
type MockTicketRedeemerServer struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRedeemerServerMockRecorder
}

// MockTicketRedeemerServerMockRecorder is the mock recorder for MockTicketRedeemerServer.
type MockTicketRedeemerServerMockRecorder struct {
	mock *MockTicketRedeemerServer
}

// NewMockTicketRedeemerServer creates a new mock instance.
func NewMockTicketRedeemerServer(ctrl *gomock.Controller) *MockTicketRedeemerServer {
	mock := &MockTicketRedeemerServer{ctrl: ctrl}
	mock.recorder = &MockTicketRedeemerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketRedeemerServer) EXPECT() *MockTicketRedeemerServerMockRecorder {
	return m.recorder
}

// QueueTicket mocks base method.
func (m *MockTicketRedeemerServer) QueueTicket(arg0 context.Context, arg1 *Ticket) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueTicket", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueTicket indicates an expected call of QueueTicket.
func (mr *MockTicketRedeemerServerMockRecorder) QueueTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueTicket", reflect.TypeOf((*MockTicketRedeemerServer)(nil).QueueTicket), arg0, arg1)
}

// MonitorMaxFloat mocks base method.
func (m *MockTicketRedeemerServer) MonitorMaxFloat(arg0 *empty.Empty, arg1 TicketRedeemer_MonitorMaxFloatServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitorMaxFloat", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MonitorMaxFloat indicates an expected call of MonitorMaxFloat.
func (mr *MockTicketRedeemerServerMockRecorder) MonitorMaxFloat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorMaxFloat", reflect.TypeOf((*MockTicketRedeemerServer)(nil).MonitorMaxFloat), arg0, arg1)
}

// MockTicketRedeemer_MonitorMaxFloatServer is a mock of TicketRedeemer_MonitorMaxFloatServer interface.
type MockTicketRedeemer_MonitorMaxFloatServer struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder
}

// MockTicketRedeemer_MonitorMaxFloatServerMockRecorder is the mock recorder for MockTicketRedeemer_MonitorMaxFloatServer.
type MockTicketRedeemer_MonitorMaxFloatServerMockRecorder struct {
	mock *MockTicketRedeemer_MonitorMaxFloatServer
}

// NewMockTicketRedeemer_MonitorMaxFloatServer creates a new mock instance.
func NewMockTicketRedeemer_MonitorMaxFloatServer(ctrl *gomock.Controller) *MockTicketRedeemer_MonitorMaxFloatServer {
	mock := &MockTicketRedeemer_MonitorMaxFloatServer{ctrl: ctrl}
	mock.recorder = &MockTicketRedeemer_MonitorMaxFloatServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketRedeemer_MonitorMaxFloatServer) EXPECT() *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatServer) Send(arg0 *MaxFloatUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatServer)(nil).Send), arg0)
}

// SetHeader mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatServer)(nil).SetTrailer), arg0)
}

// Context mocks base method.
func (m *MockTicketRedeemer_MonitorMaxFloatServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatServer)(nil).Context))
}

// SendMsg mocks base method.
func (m_2 *MockTicketRedeemer_MonitorMaxFloatServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method.
func (m_2 *MockTicketRedeemer_MonitorMaxFloatServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockTicketRedeemer_MonitorMaxFloatServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTicketRedeemer_MonitorMaxFloatServer)(nil).RecvMsg), m)
}
