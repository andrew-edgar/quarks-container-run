// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cf-operator/container-run/pkg/containerrun (interfaces: Runner,Checker,Process,OSProcess,ExecCommandContext)

// Package mocks is a generated GoMock package.
package mocks

import (
	containerrun "code.cloudfoundry.org/cf-operator/container-run/pkg/containerrun"
	context "context"
	gomock "github.com/golang/mock/gomock"
	os "os"
	exec "os/exec"
	reflect "reflect"
)

// MockRunner is a mock of Runner interface
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockRunner) Run(arg0 containerrun.Command, arg1 containerrun.Stdio) (containerrun.Process, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(containerrun.Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockRunnerMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunner)(nil).Run), arg0, arg1)
}

// RunContext mocks base method
func (m *MockRunner) RunContext(arg0 context.Context, arg1 containerrun.Command, arg2 containerrun.Stdio) (containerrun.Process, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(containerrun.Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunContext indicates an expected call of RunContext
func (mr *MockRunnerMockRecorder) RunContext(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunContext", reflect.TypeOf((*MockRunner)(nil).RunContext), arg0, arg1, arg2)
}

// MockChecker is a mock of Checker interface
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// Check mocks base method
func (m *MockChecker) Check(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockCheckerMockRecorder) Check(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockChecker)(nil).Check), arg0)
}

// MockProcess is a mock of Process interface
type MockProcess struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMockRecorder
}

// MockProcessMockRecorder is the mock recorder for MockProcess
type MockProcessMockRecorder struct {
	mock *MockProcess
}

// NewMockProcess creates a new mock instance
func NewMockProcess(ctrl *gomock.Controller) *MockProcess {
	mock := &MockProcess{ctrl: ctrl}
	mock.recorder = &MockProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcess) EXPECT() *MockProcessMockRecorder {
	return m.recorder
}

// Signal mocks base method
func (m *MockProcess) Signal(arg0 os.Signal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signal", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal
func (mr *MockProcessMockRecorder) Signal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockProcess)(nil).Signal), arg0)
}

// Wait mocks base method
func (m *MockProcess) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockProcessMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockProcess)(nil).Wait))
}

// MockOSProcess is a mock of OSProcess interface
type MockOSProcess struct {
	ctrl     *gomock.Controller
	recorder *MockOSProcessMockRecorder
}

// MockOSProcessMockRecorder is the mock recorder for MockOSProcess
type MockOSProcessMockRecorder struct {
	mock *MockOSProcess
}

// NewMockOSProcess creates a new mock instance
func NewMockOSProcess(ctrl *gomock.Controller) *MockOSProcess {
	mock := &MockOSProcess{ctrl: ctrl}
	mock.recorder = &MockOSProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOSProcess) EXPECT() *MockOSProcessMockRecorder {
	return m.recorder
}

// Signal mocks base method
func (m *MockOSProcess) Signal(arg0 os.Signal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signal", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal
func (mr *MockOSProcessMockRecorder) Signal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockOSProcess)(nil).Signal), arg0)
}

// Wait mocks base method
func (m *MockOSProcess) Wait() (*os.ProcessState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(*os.ProcessState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wait indicates an expected call of Wait
func (mr *MockOSProcessMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockOSProcess)(nil).Wait))
}

// MockExecCommandContext is a mock of ExecCommandContext interface
type MockExecCommandContext struct {
	ctrl     *gomock.Controller
	recorder *MockExecCommandContextMockRecorder
}

// MockExecCommandContextMockRecorder is the mock recorder for MockExecCommandContext
type MockExecCommandContextMockRecorder struct {
	mock *MockExecCommandContext
}

// NewMockExecCommandContext creates a new mock instance
func NewMockExecCommandContext(ctrl *gomock.Controller) *MockExecCommandContext {
	mock := &MockExecCommandContext{ctrl: ctrl}
	mock.recorder = &MockExecCommandContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecCommandContext) EXPECT() *MockExecCommandContextMockRecorder {
	return m.recorder
}

// CommandContext mocks base method
func (m *MockExecCommandContext) CommandContext(arg0 context.Context, arg1 string, arg2 ...string) *exec.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommandContext", varargs...)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// CommandContext indicates an expected call of CommandContext
func (mr *MockExecCommandContextMockRecorder) CommandContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandContext", reflect.TypeOf((*MockExecCommandContext)(nil).CommandContext), varargs...)
}
