// Code generated by mockery v2.1.0. DO NOT EDIT.

package jobservice

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	scheduler "github.com/goharbor/harbor/src/pkg/scheduler"
)

// SchedulerController is an autogenerated mock type for the SchedulerController type
type SchedulerController struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, vendorType, cronType, cron, callbackFuncName, policy, extrasParam
func (_m *SchedulerController) Create(ctx context.Context, vendorType string, cronType string, cron string, callbackFuncName string, policy interface{}, extrasParam map[string]interface{}) (int64, error) {
	ret := _m.Called(ctx, vendorType, cronType, cron, callbackFuncName, policy, extrasParam)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, interface{}, map[string]interface{}) int64); ok {
		r0 = rf(ctx, vendorType, cronType, cron, callbackFuncName, policy, extrasParam)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, interface{}, map[string]interface{}) error); ok {
		r1 = rf(ctx, vendorType, cronType, cron, callbackFuncName, policy, extrasParam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, vendorType
func (_m *SchedulerController) Delete(ctx context.Context, vendorType string) error {
	ret := _m.Called(ctx, vendorType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, vendorType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, vendorType
func (_m *SchedulerController) Get(ctx context.Context, vendorType string) (*scheduler.Schedule, error) {
	ret := _m.Called(ctx, vendorType)

	var r0 *scheduler.Schedule
	if rf, ok := ret.Get(0).(func(context.Context, string) *scheduler.Schedule); ok {
		r0 = rf(ctx, vendorType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scheduler.Schedule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, vendorType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
