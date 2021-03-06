// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/architectv/networking-course-project/backend/pkg/repositories (interfaces: Board)

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"
	models "github.com/architectv/networking-course-project/backend/pkg/models"

	gomock "github.com/golang/mock/gomock"
)

// MockBoard is a mock of Board interface.
type MockBoard struct {
	ctrl     *gomock.Controller
	recorder *MockBoardMockRecorder
}

// MockBoardMockRecorder is the mock recorder for MockBoard.
type MockBoardMockRecorder struct {
	mock *MockBoard
}

// NewMockBoard creates a new mock instance.
func NewMockBoard(ctrl *gomock.Controller) *MockBoard {
	mock := &MockBoard{ctrl: ctrl}
	mock.recorder = &MockBoardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBoard) EXPECT() *MockBoardMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBoard) Create(arg0 int, arg1 *models.Board) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBoardMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBoard)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockBoard) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBoardMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBoard)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockBoard) GetAll(arg0, arg1 int) ([]*models.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]*models.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBoardMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBoard)(nil).GetAll), arg0, arg1)
}

// GetBoardsCountByOwnerId mocks base method.
func (m *MockBoard) GetBoardsCountByOwnerId(arg0, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoardsCountByOwnerId", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoardsCountByOwnerId indicates an expected call of GetBoardsCountByOwnerId.
func (mr *MockBoardMockRecorder) GetBoardsCountByOwnerId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoardsCountByOwnerId", reflect.TypeOf((*MockBoard)(nil).GetBoardsCountByOwnerId), arg0, arg1)
}

// GetById mocks base method.
func (m *MockBoard) GetById(arg0 int) (*models.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*models.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockBoardMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockBoard)(nil).GetById), arg0)
}

// GetMembers mocks base method.
func (m *MockBoard) GetMembers(arg0 int) ([]*models.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembers", arg0)
	ret0, _ := ret[0].([]*models.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockBoardMockRecorder) GetMembers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockBoard)(nil).GetMembers), arg0)
}

// GetPermissions mocks base method.
func (m *MockBoard) GetPermissions(arg0, arg1 int) (*models.Permission, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPermissions", arg0, arg1)
	ret0, _ := ret[0].(*models.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermissions indicates an expected call of GetPermissions.
func (mr *MockBoardMockRecorder) GetPermissions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermissions", reflect.TypeOf((*MockBoard)(nil).GetPermissions), arg0, arg1)
}

// Update mocks base method.
func (m *MockBoard) Update(arg0 int, arg1 *models.UpdateBoard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBoardMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBoard)(nil).Update), arg0, arg1)
}
