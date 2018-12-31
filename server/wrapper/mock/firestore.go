// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-hermes/hermes-app/server/wrapper (interfaces: DBClientInterface,CollectionRefInterface,DocRefInterface,DocumentInteratorInterface)

// Package mock_wrapper is a generated GoMock package.
package mock_wrapper

import (
	firestore "cloud.google.com/go/firestore"
	context "context"
	gomock "github.com/golang/mock/gomock"
	wrapper "github.com/project-hermes/hermes-app/server/wrapper"
	reflect "reflect"
)

// MockDBClientInterface is a mock of DBClientInterface interface
type MockDBClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientInterfaceMockRecorder
}

// MockDBClientInterfaceMockRecorder is the mock recorder for MockDBClientInterface
type MockDBClientInterfaceMockRecorder struct {
	mock *MockDBClientInterface
}

// NewMockDBClientInterface creates a new mock instance
func NewMockDBClientInterface(ctrl *gomock.Controller) *MockDBClientInterface {
	mock := &MockDBClientInterface{ctrl: ctrl}
	mock.recorder = &MockDBClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBClientInterface) EXPECT() *MockDBClientInterfaceMockRecorder {
	return m.recorder
}

// Collection mocks base method
func (m *MockDBClientInterface) Collection(arg0 string) wrapper.CollectionRefInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collection", arg0)
	ret0, _ := ret[0].(wrapper.CollectionRefInterface)
	return ret0
}

// Collection indicates an expected call of Collection
func (mr *MockDBClientInterfaceMockRecorder) Collection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockDBClientInterface)(nil).Collection), arg0)
}

// MockCollectionRefInterface is a mock of CollectionRefInterface interface
type MockCollectionRefInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionRefInterfaceMockRecorder
}

// MockCollectionRefInterfaceMockRecorder is the mock recorder for MockCollectionRefInterface
type MockCollectionRefInterfaceMockRecorder struct {
	mock *MockCollectionRefInterface
}

// NewMockCollectionRefInterface creates a new mock instance
func NewMockCollectionRefInterface(ctrl *gomock.Controller) *MockCollectionRefInterface {
	mock := &MockCollectionRefInterface{ctrl: ctrl}
	mock.recorder = &MockCollectionRefInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollectionRefInterface) EXPECT() *MockCollectionRefInterfaceMockRecorder {
	return m.recorder
}

// Doc mocks base method
func (m *MockCollectionRefInterface) Doc(arg0 string) wrapper.DocRefInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Doc", arg0)
	ret0, _ := ret[0].(wrapper.DocRefInterface)
	return ret0
}

// Doc indicates an expected call of Doc
func (mr *MockCollectionRefInterfaceMockRecorder) Doc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Doc", reflect.TypeOf((*MockCollectionRefInterface)(nil).Doc), arg0)
}

// Documents mocks base method
func (m *MockCollectionRefInterface) Documents(arg0 context.Context) wrapper.DocumentInteratorInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Documents", arg0)
	ret0, _ := ret[0].(wrapper.DocumentInteratorInterface)
	return ret0
}

// Documents indicates an expected call of Documents
func (mr *MockCollectionRefInterfaceMockRecorder) Documents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Documents", reflect.TypeOf((*MockCollectionRefInterface)(nil).Documents), arg0)
}

// NewDoc mocks base method
func (m *MockCollectionRefInterface) NewDoc() wrapper.DocRefInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDoc")
	ret0, _ := ret[0].(wrapper.DocRefInterface)
	return ret0
}

// NewDoc indicates an expected call of NewDoc
func (mr *MockCollectionRefInterfaceMockRecorder) NewDoc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDoc", reflect.TypeOf((*MockCollectionRefInterface)(nil).NewDoc))
}

// MockDocRefInterface is a mock of DocRefInterface interface
type MockDocRefInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDocRefInterfaceMockRecorder
}

// MockDocRefInterfaceMockRecorder is the mock recorder for MockDocRefInterface
type MockDocRefInterfaceMockRecorder struct {
	mock *MockDocRefInterface
}

// NewMockDocRefInterface creates a new mock instance
func NewMockDocRefInterface(ctrl *gomock.Controller) *MockDocRefInterface {
	mock := &MockDocRefInterface{ctrl: ctrl}
	mock.recorder = &MockDocRefInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocRefInterface) EXPECT() *MockDocRefInterfaceMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockDocRefInterface) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockDocRefInterfaceMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockDocRefInterface)(nil).ID))
}

// Set mocks base method
func (m *MockDocRefInterface) Set(arg0 context.Context, arg1 interface{}, arg2 ...firestore.SetOption) (*firestore.WriteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Set", varargs...)
	ret0, _ := ret[0].(*firestore.WriteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockDocRefInterfaceMockRecorder) Set(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDocRefInterface)(nil).Set), varargs...)
}

// MockDocumentInteratorInterface is a mock of DocumentInteratorInterface interface
type MockDocumentInteratorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDocumentInteratorInterfaceMockRecorder
}

// MockDocumentInteratorInterfaceMockRecorder is the mock recorder for MockDocumentInteratorInterface
type MockDocumentInteratorInterfaceMockRecorder struct {
	mock *MockDocumentInteratorInterface
}

// NewMockDocumentInteratorInterface creates a new mock instance
func NewMockDocumentInteratorInterface(ctrl *gomock.Controller) *MockDocumentInteratorInterface {
	mock := &MockDocumentInteratorInterface{ctrl: ctrl}
	mock.recorder = &MockDocumentInteratorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDocumentInteratorInterface) EXPECT() *MockDocumentInteratorInterfaceMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockDocumentInteratorInterface) Next() (*firestore.DocumentSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*firestore.DocumentSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockDocumentInteratorInterfaceMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockDocumentInteratorInterface)(nil).Next))
}
