// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozonva/ova-entertainment-api/internal/models"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddEntertainments mocks base method.
func (m *MockRepo) AddEntertainments(models []models.Entertainment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEntertainments", models)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEntertainments indicates an expected call of AddEntertainments.
func (mr *MockRepoMockRecorder) AddEntertainments(models interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntertainments", reflect.TypeOf((*MockRepo)(nil).AddEntertainments), models)
}

// DescribeEntertainment mocks base method.
func (m *MockRepo) UpdateEntertainment(model models.Entertainment) (*models.Entertainment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEntertainment", model)
	ret0, _ := ret[0].(*models.Entertainment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntertainment indicates an expected call of DescribeEntertainment.
func (mr *MockRepoMockRecorder) DescribeEntertainment(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEntertainment", reflect.TypeOf((*MockRepo)(nil).UpdateEntertainment), model)
}

// FindEntertainment mocks base method.
func (m *MockRepo) FindEntertainment(ID uint64, title string) (*models.Entertainment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntertainment", ID, title)
	ret0, _ := ret[0].(*models.Entertainment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntertainment indicates an expected call of FindEntertainment.
func (mr *MockRepoMockRecorder) FindEntertainment(ID, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntertainment", reflect.TypeOf((*MockRepo)(nil).FindEntertainment), ID, title)
}

// ListEntertainments mocks base method.
func (m *MockRepo) ListEntertainments(limit, offset uint32) ([]models.Entertainment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntertainments", limit, offset)
	ret0, _ := ret[0].([]models.Entertainment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntertainments indicates an expected call of ListEntertainments.
func (mr *MockRepoMockRecorder) ListEntertainments(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntertainments", reflect.TypeOf((*MockRepo)(nil).ListEntertainments), limit, offset)
}

// RemoveEntertainment mocks base method.
func (m *MockRepo) RemoveEntertainment(ID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEntertainment", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEntertainment indicates an expected call of RemoveEntertainment.
func (mr *MockRepoMockRecorder) RemoveEntertainment(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEntertainment", reflect.TypeOf((*MockRepo)(nil).RemoveEntertainment), ID)
}
