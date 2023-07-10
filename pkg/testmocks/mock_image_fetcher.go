// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/pack/pkg/client (interfaces: ImageFetcher)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	context "context"
	reflect "reflect"

	imgutil "github.com/buildpacks/imgutil"
	gomock "github.com/golang/mock/gomock"

	image "github.com/buildpacks/pack/pkg/image"
)

// MockImageFetcher is a mock of ImageFetcher interface.
type MockImageFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockImageFetcherMockRecorder
}

// MockImageFetcherMockRecorder is the mock recorder for MockImageFetcher.
type MockImageFetcherMockRecorder struct {
	mock *MockImageFetcher
}

// NewMockImageFetcher creates a new mock instance.
func NewMockImageFetcher(ctrl *gomock.Controller) *MockImageFetcher {
	mock := &MockImageFetcher{ctrl: ctrl}
	mock.recorder = &MockImageFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageFetcher) EXPECT() *MockImageFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockImageFetcher) Fetch(arg0 context.Context, arg1 string, arg2 image.FetchOptions) (imgutil.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0, arg1, arg2)
	ret0, _ := ret[0].(imgutil.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockImageFetcherMockRecorder) Fetch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockImageFetcher)(nil).Fetch), arg0, arg1, arg2)
}