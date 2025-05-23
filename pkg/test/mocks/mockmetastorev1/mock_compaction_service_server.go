// Code generated by mockery. DO NOT EDIT.

package mockmetastorev1

import (
	context "context"

	metastorev1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"
	mock "github.com/stretchr/testify/mock"
)

// MockCompactionServiceServer is an autogenerated mock type for the CompactionServiceServer type
type MockCompactionServiceServer struct {
	mock.Mock
}

type MockCompactionServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCompactionServiceServer) EXPECT() *MockCompactionServiceServer_Expecter {
	return &MockCompactionServiceServer_Expecter{mock: &_m.Mock}
}

// PollCompactionJobs provides a mock function with given fields: _a0, _a1
func (_m *MockCompactionServiceServer) PollCompactionJobs(_a0 context.Context, _a1 *metastorev1.PollCompactionJobsRequest) (*metastorev1.PollCompactionJobsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PollCompactionJobs")
	}

	var r0 *metastorev1.PollCompactionJobsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.PollCompactionJobsRequest) (*metastorev1.PollCompactionJobsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.PollCompactionJobsRequest) *metastorev1.PollCompactionJobsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metastorev1.PollCompactionJobsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *metastorev1.PollCompactionJobsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCompactionServiceServer_PollCompactionJobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PollCompactionJobs'
type MockCompactionServiceServer_PollCompactionJobs_Call struct {
	*mock.Call
}

// PollCompactionJobs is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *metastorev1.PollCompactionJobsRequest
func (_e *MockCompactionServiceServer_Expecter) PollCompactionJobs(_a0 interface{}, _a1 interface{}) *MockCompactionServiceServer_PollCompactionJobs_Call {
	return &MockCompactionServiceServer_PollCompactionJobs_Call{Call: _e.mock.On("PollCompactionJobs", _a0, _a1)}
}

func (_c *MockCompactionServiceServer_PollCompactionJobs_Call) Run(run func(_a0 context.Context, _a1 *metastorev1.PollCompactionJobsRequest)) *MockCompactionServiceServer_PollCompactionJobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*metastorev1.PollCompactionJobsRequest))
	})
	return _c
}

func (_c *MockCompactionServiceServer_PollCompactionJobs_Call) Return(_a0 *metastorev1.PollCompactionJobsResponse, _a1 error) *MockCompactionServiceServer_PollCompactionJobs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCompactionServiceServer_PollCompactionJobs_Call) RunAndReturn(run func(context.Context, *metastorev1.PollCompactionJobsRequest) (*metastorev1.PollCompactionJobsResponse, error)) *MockCompactionServiceServer_PollCompactionJobs_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedCompactionServiceServer provides a mock function with given fields:
func (_m *MockCompactionServiceServer) mustEmbedUnimplementedCompactionServiceServer() {
	_m.Called()
}

// MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedCompactionServiceServer'
type MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedCompactionServiceServer is a helper method to define mock.On call
func (_e *MockCompactionServiceServer_Expecter) mustEmbedUnimplementedCompactionServiceServer() *MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call {
	return &MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedCompactionServiceServer")}
}

func (_c *MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call) Run(run func()) *MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call) Return() *MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call) RunAndReturn(run func()) *MockCompactionServiceServer_mustEmbedUnimplementedCompactionServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCompactionServiceServer creates a new instance of MockCompactionServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCompactionServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCompactionServiceServer {
	mock := &MockCompactionServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
