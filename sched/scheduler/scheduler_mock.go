// Automatically generated by MockGen. DO NOT EDIT!
// Source: scheduler.go

package scheduler

import (
	gomock "github.com/golang/mock/gomock"
	sched "github.com/scootdev/scoot/sched"
)

// Mock of Scheduler interface
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *_MockSchedulerRecorder
}

// Recorder for MockScheduler (not exported)
type _MockSchedulerRecorder struct {
	mock *MockScheduler
}

func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &_MockSchedulerRecorder{mock}
	return mock
}

func (_m *MockScheduler) EXPECT() *_MockSchedulerRecorder {
	return _m.recorder
}

func (_m *MockScheduler) ScheduleJob(job sched.Job) error {
	ret := _m.ctrl.Call(_m, "ScheduleJob", job)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSchedulerRecorder) ScheduleJob(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ScheduleJob", arg0)
}