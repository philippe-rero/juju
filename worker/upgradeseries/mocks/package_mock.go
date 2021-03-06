// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/upgradeseries (interfaces: Facade,Logger,AgentService,ServiceAccess,Upgrader)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/juju/juju/core/model"
	watcher "github.com/juju/juju/core/watcher"
	upgradeseries "github.com/juju/juju/worker/upgradeseries"
	names_v2 "gopkg.in/juju/names.v2"
	reflect "reflect"
)

// MockFacade is a mock of Facade interface
type MockFacade struct {
	ctrl     *gomock.Controller
	recorder *MockFacadeMockRecorder
}

// MockFacadeMockRecorder is the mock recorder for MockFacade
type MockFacadeMockRecorder struct {
	mock *MockFacade
}

// NewMockFacade creates a new mock instance
func NewMockFacade(ctrl *gomock.Controller) *MockFacade {
	mock := &MockFacade{ctrl: ctrl}
	mock.recorder = &MockFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFacade) EXPECT() *MockFacadeMockRecorder {
	return m.recorder
}

// FinishUpgradeSeries mocks base method
func (m *MockFacade) FinishUpgradeSeries(arg0 string) error {
	ret := m.ctrl.Call(m, "FinishUpgradeSeries", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishUpgradeSeries indicates an expected call of FinishUpgradeSeries
func (mr *MockFacadeMockRecorder) FinishUpgradeSeries(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishUpgradeSeries", reflect.TypeOf((*MockFacade)(nil).FinishUpgradeSeries), arg0)
}

// MachineStatus mocks base method
func (m *MockFacade) MachineStatus() (model.UpgradeSeriesStatus, error) {
	ret := m.ctrl.Call(m, "MachineStatus")
	ret0, _ := ret[0].(model.UpgradeSeriesStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MachineStatus indicates an expected call of MachineStatus
func (mr *MockFacadeMockRecorder) MachineStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineStatus", reflect.TypeOf((*MockFacade)(nil).MachineStatus))
}

// PinMachineApplications mocks base method
func (m *MockFacade) PinMachineApplications() (map[string]error, error) {
	ret := m.ctrl.Call(m, "PinMachineApplications")
	ret0, _ := ret[0].(map[string]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PinMachineApplications indicates an expected call of PinMachineApplications
func (mr *MockFacadeMockRecorder) PinMachineApplications() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinMachineApplications", reflect.TypeOf((*MockFacade)(nil).PinMachineApplications))
}

// SetMachineStatus mocks base method
func (m *MockFacade) SetMachineStatus(arg0 model.UpgradeSeriesStatus, arg1 string) error {
	ret := m.ctrl.Call(m, "SetMachineStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMachineStatus indicates an expected call of SetMachineStatus
func (mr *MockFacadeMockRecorder) SetMachineStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMachineStatus", reflect.TypeOf((*MockFacade)(nil).SetMachineStatus), arg0, arg1)
}

// StartUnitCompletion mocks base method
func (m *MockFacade) StartUnitCompletion(arg0 string) error {
	ret := m.ctrl.Call(m, "StartUnitCompletion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartUnitCompletion indicates an expected call of StartUnitCompletion
func (mr *MockFacadeMockRecorder) StartUnitCompletion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartUnitCompletion", reflect.TypeOf((*MockFacade)(nil).StartUnitCompletion), arg0)
}

// TargetSeries mocks base method
func (m *MockFacade) TargetSeries() (string, error) {
	ret := m.ctrl.Call(m, "TargetSeries")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TargetSeries indicates an expected call of TargetSeries
func (mr *MockFacadeMockRecorder) TargetSeries() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetSeries", reflect.TypeOf((*MockFacade)(nil).TargetSeries))
}

// UnitsCompleted mocks base method
func (m *MockFacade) UnitsCompleted() ([]names_v2.UnitTag, error) {
	ret := m.ctrl.Call(m, "UnitsCompleted")
	ret0, _ := ret[0].([]names_v2.UnitTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitsCompleted indicates an expected call of UnitsCompleted
func (mr *MockFacadeMockRecorder) UnitsCompleted() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitsCompleted", reflect.TypeOf((*MockFacade)(nil).UnitsCompleted))
}

// UnitsPrepared mocks base method
func (m *MockFacade) UnitsPrepared() ([]names_v2.UnitTag, error) {
	ret := m.ctrl.Call(m, "UnitsPrepared")
	ret0, _ := ret[0].([]names_v2.UnitTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitsPrepared indicates an expected call of UnitsPrepared
func (mr *MockFacadeMockRecorder) UnitsPrepared() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitsPrepared", reflect.TypeOf((*MockFacade)(nil).UnitsPrepared))
}

// UnpinMachineApplications mocks base method
func (m *MockFacade) UnpinMachineApplications() (map[string]error, error) {
	ret := m.ctrl.Call(m, "UnpinMachineApplications")
	ret0, _ := ret[0].(map[string]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpinMachineApplications indicates an expected call of UnpinMachineApplications
func (mr *MockFacadeMockRecorder) UnpinMachineApplications() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpinMachineApplications", reflect.TypeOf((*MockFacade)(nil).UnpinMachineApplications))
}

// WatchUpgradeSeriesNotifications mocks base method
func (m *MockFacade) WatchUpgradeSeriesNotifications() (watcher.NotifyWatcher, error) {
	ret := m.ctrl.Call(m, "WatchUpgradeSeriesNotifications")
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUpgradeSeriesNotifications indicates an expected call of WatchUpgradeSeriesNotifications
func (mr *MockFacadeMockRecorder) WatchUpgradeSeriesNotifications() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpgradeSeriesNotifications", reflect.TypeOf((*MockFacade)(nil).WatchUpgradeSeriesNotifications))
}

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method
func (m *MockLogger) Errorf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Infof mocks base method
func (m *MockLogger) Infof(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Warningf mocks base method
func (m *MockLogger) Warningf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf
func (mr *MockLoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogger)(nil).Warningf), varargs...)
}

// MockAgentService is a mock of AgentService interface
type MockAgentService struct {
	ctrl     *gomock.Controller
	recorder *MockAgentServiceMockRecorder
}

// MockAgentServiceMockRecorder is the mock recorder for MockAgentService
type MockAgentServiceMockRecorder struct {
	mock *MockAgentService
}

// NewMockAgentService creates a new mock instance
func NewMockAgentService(ctrl *gomock.Controller) *MockAgentService {
	mock := &MockAgentService{ctrl: ctrl}
	mock.recorder = &MockAgentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentService) EXPECT() *MockAgentServiceMockRecorder {
	return m.recorder
}

// Running mocks base method
func (m *MockAgentService) Running() (bool, error) {
	ret := m.ctrl.Call(m, "Running")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Running indicates an expected call of Running
func (mr *MockAgentServiceMockRecorder) Running() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Running", reflect.TypeOf((*MockAgentService)(nil).Running))
}

// Start mocks base method
func (m *MockAgentService) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockAgentServiceMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAgentService)(nil).Start))
}

// Stop mocks base method
func (m *MockAgentService) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockAgentServiceMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockAgentService)(nil).Stop))
}

// MockServiceAccess is a mock of ServiceAccess interface
type MockServiceAccess struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccessMockRecorder
}

// MockServiceAccessMockRecorder is the mock recorder for MockServiceAccess
type MockServiceAccessMockRecorder struct {
	mock *MockServiceAccess
}

// NewMockServiceAccess creates a new mock instance
func NewMockServiceAccess(ctrl *gomock.Controller) *MockServiceAccess {
	mock := &MockServiceAccess{ctrl: ctrl}
	mock.recorder = &MockServiceAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceAccess) EXPECT() *MockServiceAccessMockRecorder {
	return m.recorder
}

// DiscoverService mocks base method
func (m *MockServiceAccess) DiscoverService(arg0 string) (upgradeseries.AgentService, error) {
	ret := m.ctrl.Call(m, "DiscoverService", arg0)
	ret0, _ := ret[0].(upgradeseries.AgentService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverService indicates an expected call of DiscoverService
func (mr *MockServiceAccessMockRecorder) DiscoverService(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverService", reflect.TypeOf((*MockServiceAccess)(nil).DiscoverService), arg0)
}

// ListServices mocks base method
func (m *MockServiceAccess) ListServices() ([]string, error) {
	ret := m.ctrl.Call(m, "ListServices")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockServiceAccessMockRecorder) ListServices() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockServiceAccess)(nil).ListServices))
}

// MockUpgrader is a mock of Upgrader interface
type MockUpgrader struct {
	ctrl     *gomock.Controller
	recorder *MockUpgraderMockRecorder
}

// MockUpgraderMockRecorder is the mock recorder for MockUpgrader
type MockUpgraderMockRecorder struct {
	mock *MockUpgrader
}

// NewMockUpgrader creates a new mock instance
func NewMockUpgrader(ctrl *gomock.Controller) *MockUpgrader {
	mock := &MockUpgrader{ctrl: ctrl}
	mock.recorder = &MockUpgraderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpgrader) EXPECT() *MockUpgraderMockRecorder {
	return m.recorder
}

// PerformUpgrade mocks base method
func (m *MockUpgrader) PerformUpgrade() error {
	ret := m.ctrl.Call(m, "PerformUpgrade")
	ret0, _ := ret[0].(error)
	return ret0
}

// PerformUpgrade indicates an expected call of PerformUpgrade
func (mr *MockUpgraderMockRecorder) PerformUpgrade() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformUpgrade", reflect.TypeOf((*MockUpgrader)(nil).PerformUpgrade))
}
