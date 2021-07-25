// Code generated by MockGen. DO NOT EDIT.
// Source: sampling.go

// Package mock_sampling is a generated GoMock package.
package mock_sampling

import (
	v1alpha1 "github.com/alibaba/morphling/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSampling is a mock of Sampling interface
type MockSampling struct {
	ctrl     *gomock.Controller
	recorder *MockSamplingMockRecorder
}

// MockSamplingMockRecorder is the mock recorder for MockSampling
type MockSamplingMockRecorder struct {
	mock *MockSampling
}

// NewMockSampling creates a new mock instance
func NewMockSampling(ctrl *gomock.Controller) *MockSampling {
	mock := &MockSampling{ctrl: ctrl}
	mock.recorder = &MockSamplingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSampling) EXPECT() *MockSamplingMockRecorder {
	return m.recorder
}

// GetOrCreateSampling mocks base method
func (m *MockSampling) GetOrCreateSampling(suggestionRequests int32, instance *v1alpha1.ProfilingExperiment, samplingRequests *v1alpha1.ObjectiveSpec) (*v1alpha1.Sampling, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateSampling", suggestionRequests, instance, samplingRequests)
	ret0, _ := ret[0].(*v1alpha1.Sampling)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreateSampling indicates an expected call of GetOrCreateSampling
func (mr *MockSamplingMockRecorder) GetOrCreateSampling(suggestionRequests, instance, samplingRequests interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateSampling", reflect.TypeOf((*MockSampling)(nil).GetOrCreateSampling), suggestionRequests, instance, samplingRequests)
}

// UpdateSampling mocks base method
func (m *MockSampling) UpdateSampling(sampling *v1alpha1.Sampling) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSampling", sampling)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSampling indicates an expected call of UpdateSampling
func (mr *MockSamplingMockRecorder) UpdateSampling(sampling interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSampling", reflect.TypeOf((*MockSampling)(nil).UpdateSampling), sampling)
}

// UpdateSamplingStatus mocks base method
func (m *MockSampling) UpdateSamplingStatus(sampling *v1alpha1.Sampling) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSamplingStatus", sampling)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSamplingStatus indicates an expected call of UpdateSamplingStatus
func (mr *MockSamplingMockRecorder) UpdateSamplingStatus(sampling interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSamplingStatus", reflect.TypeOf((*MockSampling)(nil).UpdateSamplingStatus), sampling)
}

// GetSamplings mocks base method
func (m *MockSampling) GetSamplings(numRequests int32, instance *v1alpha1.ProfilingExperiment, currentCount int32, trials []v1alpha1.Trial) ([]v1alpha1.TrialAssignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSamplings", numRequests, instance, currentCount, trials)
	ret0, _ := ret[0].([]v1alpha1.TrialAssignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamplings indicates an expected call of GetSamplings
func (mr *MockSamplingMockRecorder) GetSamplings(numRequests, instance, currentCount, trials interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamplings", reflect.TypeOf((*MockSampling)(nil).GetSamplings), numRequests, instance, currentCount, trials)
}
