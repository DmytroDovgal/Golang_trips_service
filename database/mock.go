// Code generated by MockGen. DO NOT EDIT.
// Source: team-project/database (interfaces: TicketRepository)

// Package database is a generated GoMock package.
package database

import (
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
	data "team-project/services/data"
)

// MockTicketRepository is a mock of TicketRepository interface
type MockTicketRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRepositoryMockRecorder
}

// MockTicketRepositoryMockRecorder is the mock recorder for MockTicketRepository
type MockTicketRepositoryMockRecorder struct {
	mock *MockTicketRepository
}

// NewMockTicketRepository creates a new mock instance
func NewMockTicketRepository(ctrl *gomock.Controller) *MockTicketRepository {
	mock := &MockTicketRepository{ctrl: ctrl}
	mock.recorder = &MockTicketRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTicketRepository) EXPECT() *MockTicketRepositoryMockRecorder {
	return m.recorder
}

// AllTickets mocks base method
func (m *MockTicketRepository) AllTickets() ([]data.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllTickets")
	ret0, _ := ret[0].([]data.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllTickets indicates an expected call of AllTickets
func (mr *MockTicketRepositoryMockRecorder) AllTickets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllTickets", reflect.TypeOf((*MockTicketRepository)(nil).AllTickets))
}

// CreateTicket mocks base method
func (m *MockTicketRepository) CreateTicket(arg0 data.Ticket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicket", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTicket indicates an expected call of CreateTicket
func (mr *MockTicketRepositoryMockRecorder) CreateTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockTicketRepository)(nil).CreateTicket), arg0)
}

// DeleteTicket mocks base method
func (m *MockTicketRepository) DeleteTicket(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicket", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTicket indicates an expected call of DeleteTicket
func (mr *MockTicketRepositoryMockRecorder) DeleteTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicket", reflect.TypeOf((*MockTicketRepository)(nil).DeleteTicket), arg0)
}

// GetTicket mocks base method
func (m *MockTicketRepository) GetTicket(arg0 uuid.UUID) (data.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicket", arg0)
	ret0, _ := ret[0].(data.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket
func (mr *MockTicketRepositoryMockRecorder) GetTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*MockTicketRepository)(nil).GetTicket), arg0)
}

// UpdateTicket mocks base method
func (m *MockTicketRepository) UpdateTicket(arg0 data.Ticket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicket", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTicket indicates an expected call of UpdateTicket
func (mr *MockTicketRepositoryMockRecorder) UpdateTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicket", reflect.TypeOf((*MockTicketRepository)(nil).UpdateTicket), arg0)
}
