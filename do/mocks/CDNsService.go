// Code generated by mockery v1.0.0. DO NOT EDIT.

// Generated: please do not edit by hand

package mocks

import do "github.com/digitalocean/doctl/do"
import godo "github.com/digitalocean/godo"
import mock "github.com/stretchr/testify/mock"

// CDNsService is an autogenerated mock type for the CDNsService type
type CDNsService struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *CDNsService) Create(_a0 *godo.CDNCreateRequest) (*do.CDN, error) {
	ret := _m.Called(_a0)

	var r0 *do.CDN
	if rf, ok := ret.Get(0).(func(*godo.CDNCreateRequest) *do.CDN); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.CDN)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.CDNCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *CDNsService) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushCache provides a mock function with given fields: _a0, _a1
func (_m *CDNsService) FlushCache(_a0 string, _a1 *godo.CDNFlushCacheRequest) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *godo.CDNFlushCacheRequest) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *CDNsService) Get(_a0 string) (*do.CDN, error) {
	ret := _m.Called(_a0)

	var r0 *do.CDN
	if rf, ok := ret.Get(0).(func(string) *do.CDN); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.CDN)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *CDNsService) List() ([]do.CDN, error) {
	ret := _m.Called()

	var r0 []do.CDN
	if rf, ok := ret.Get(0).(func() []do.CDN); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]do.CDN)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTTL provides a mock function with given fields: _a0, _a1
func (_m *CDNsService) UpdateTTL(_a0 string, _a1 *godo.CDNUpdateRequest) (*do.CDN, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.CDN
	if rf, ok := ret.Get(0).(func(string, *godo.CDNUpdateRequest) *do.CDN); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.CDN)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *godo.CDNUpdateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}