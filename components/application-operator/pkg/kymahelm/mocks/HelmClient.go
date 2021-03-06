// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import services "k8s.io/helm/pkg/proto/hapi/services"

// HelmClient is an autogenerated mock type for the HelmClient type
type HelmClient struct {
	mock.Mock
}

// DeleteRelease provides a mock function with given fields: releaseName
func (_m *HelmClient) DeleteRelease(releaseName string) (*services.UninstallReleaseResponse, error) {
	ret := _m.Called(releaseName)

	var r0 *services.UninstallReleaseResponse
	if rf, ok := ret.Get(0).(func(string) *services.UninstallReleaseResponse); ok {
		r0 = rf(releaseName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.UninstallReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(releaseName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallReleaseFromChart provides a mock function with given fields: chartDir, ns, releaseName, overrides
func (_m *HelmClient) InstallReleaseFromChart(chartDir string, ns string, releaseName string, overrides string) (*services.InstallReleaseResponse, error) {
	ret := _m.Called(chartDir, ns, releaseName, overrides)

	var r0 *services.InstallReleaseResponse
	if rf, ok := ret.Get(0).(func(string, string, string, string) *services.InstallReleaseResponse); ok {
		r0 = rf(chartDir, ns, releaseName, overrides)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.InstallReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(chartDir, ns, releaseName, overrides)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReleases provides a mock function with given fields: ns
func (_m *HelmClient) ListReleases(ns string) (*services.ListReleasesResponse, error) {
	ret := _m.Called(ns)

	var r0 *services.ListReleasesResponse
	if rf, ok := ret.Get(0).(func(string) *services.ListReleasesResponse); ok {
		r0 = rf(ns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.ListReleasesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseStatus provides a mock function with given fields: rlsName
func (_m *HelmClient) ReleaseStatus(rlsName string) (*services.GetReleaseStatusResponse, error) {
	ret := _m.Called(rlsName)

	var r0 *services.GetReleaseStatusResponse
	if rf, ok := ret.Get(0).(func(string) *services.GetReleaseStatusResponse); ok {
		r0 = rf(rlsName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.GetReleaseStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(rlsName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateReleaseFromChart provides a mock function with given fields: chartDir, releaseName, overrides
func (_m *HelmClient) UpdateReleaseFromChart(chartDir string, releaseName string, overrides string) (*services.UpdateReleaseResponse, error) {
	ret := _m.Called(chartDir, releaseName, overrides)

	var r0 *services.UpdateReleaseResponse
	if rf, ok := ret.Get(0).(func(string, string, string) *services.UpdateReleaseResponse); ok {
		r0 = rf(chartDir, releaseName, overrides)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.UpdateReleaseResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(chartDir, releaseName, overrides)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
