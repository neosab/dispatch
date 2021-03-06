// Code generated by mockery v1.0.0

package mocks

import context "context"
import filters "github.com/docker/docker/api/types/filters"
import io "io"
import mock "github.com/stretchr/testify/mock"
import registry "github.com/docker/docker/api/types/registry"
import types "github.com/docker/docker/api/types"

// ImageAPIClient is an autogenerated mock type for the ImageAPIClient type
type ImageAPIClient struct {
	mock.Mock
}

// ImageBuild provides a mock function with given fields: ctx, _a1, options
func (_m *ImageAPIClient) ImageBuild(ctx context.Context, _a1 io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	ret := _m.Called(ctx, _a1, options)

	var r0 types.ImageBuildResponse
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, types.ImageBuildOptions) types.ImageBuildResponse); ok {
		r0 = rf(ctx, _a1, options)
	} else {
		r0 = ret.Get(0).(types.ImageBuildResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, io.Reader, types.ImageBuildOptions) error); ok {
		r1 = rf(ctx, _a1, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageCreate provides a mock function with given fields: ctx, parentReference, options
func (_m *ImageAPIClient) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, parentReference, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImageCreateOptions) io.ReadCloser); ok {
		r0 = rf(ctx, parentReference, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImageCreateOptions) error); ok {
		r1 = rf(ctx, parentReference, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageHistory provides a mock function with given fields: ctx, image
func (_m *ImageAPIClient) ImageHistory(ctx context.Context, image string) ([]types.ImageHistory, error) {
	ret := _m.Called(ctx, image)

	var r0 []types.ImageHistory
	if rf, ok := ret.Get(0).(func(context.Context, string) []types.ImageHistory); ok {
		r0 = rf(ctx, image)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ImageHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, image)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageImport provides a mock function with given fields: ctx, source, ref, options
func (_m *ImageAPIClient) ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, source, ref, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, types.ImageImportSource, string, types.ImageImportOptions) io.ReadCloser); ok {
		r0 = rf(ctx, source, ref, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ImageImportSource, string, types.ImageImportOptions) error); ok {
		r1 = rf(ctx, source, ref, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageInspectWithRaw provides a mock function with given fields: ctx, image
func (_m *ImageAPIClient) ImageInspectWithRaw(ctx context.Context, image string) (types.ImageInspect, []byte, error) {
	ret := _m.Called(ctx, image)

	var r0 types.ImageInspect
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ImageInspect); ok {
		r0 = rf(ctx, image)
	} else {
		r0 = ret.Get(0).(types.ImageInspect)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, image)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, image)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ImageList provides a mock function with given fields: ctx, options
func (_m *ImageAPIClient) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.ImageSummary
	if rf, ok := ret.Get(0).(func(context.Context, types.ImageListOptions) []types.ImageSummary); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ImageSummary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ImageListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageLoad provides a mock function with given fields: ctx, input, quiet
func (_m *ImageAPIClient) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error) {
	ret := _m.Called(ctx, input, quiet)

	var r0 types.ImageLoadResponse
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, bool) types.ImageLoadResponse); ok {
		r0 = rf(ctx, input, quiet)
	} else {
		r0 = ret.Get(0).(types.ImageLoadResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, io.Reader, bool) error); ok {
		r1 = rf(ctx, input, quiet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImagePull provides a mock function with given fields: ctx, ref, options
func (_m *ImageAPIClient) ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, ref, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePullOptions) io.ReadCloser); ok {
		r0 = rf(ctx, ref, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImagePullOptions) error); ok {
		r1 = rf(ctx, ref, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImagePush provides a mock function with given fields: ctx, ref, options
func (_m *ImageAPIClient) ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, ref, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePushOptions) io.ReadCloser); ok {
		r0 = rf(ctx, ref, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImagePushOptions) error); ok {
		r1 = rf(ctx, ref, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageRemove provides a mock function with given fields: ctx, image, options
func (_m *ImageAPIClient) ImageRemove(ctx context.Context, image string, options types.ImageRemoveOptions) ([]types.ImageDelete, error) {
	ret := _m.Called(ctx, image, options)

	var r0 []types.ImageDelete
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImageRemoveOptions) []types.ImageDelete); ok {
		r0 = rf(ctx, image, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ImageDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImageRemoveOptions) error); ok {
		r1 = rf(ctx, image, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageSave provides a mock function with given fields: ctx, images
func (_m *ImageAPIClient) ImageSave(ctx context.Context, images []string) (io.ReadCloser, error) {
	ret := _m.Called(ctx, images)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, []string) io.ReadCloser); ok {
		r0 = rf(ctx, images)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, images)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageSearch provides a mock function with given fields: ctx, term, options
func (_m *ImageAPIClient) ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error) {
	ret := _m.Called(ctx, term, options)

	var r0 []registry.SearchResult
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImageSearchOptions) []registry.SearchResult); ok {
		r0 = rf(ctx, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]registry.SearchResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImageSearchOptions) error); ok {
		r1 = rf(ctx, term, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageTag provides a mock function with given fields: ctx, image, ref
func (_m *ImageAPIClient) ImageTag(ctx context.Context, image string, ref string) error {
	ret := _m.Called(ctx, image, ref)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, image, ref)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImagesPrune provides a mock function with given fields: ctx, pruneFilter
func (_m *ImageAPIClient) ImagesPrune(ctx context.Context, pruneFilter filters.Args) (types.ImagesPruneReport, error) {
	ret := _m.Called(ctx, pruneFilter)

	var r0 types.ImagesPruneReport
	if rf, ok := ret.Get(0).(func(context.Context, filters.Args) types.ImagesPruneReport); ok {
		r0 = rf(ctx, pruneFilter)
	} else {
		r0 = ret.Get(0).(types.ImagesPruneReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, filters.Args) error); ok {
		r1 = rf(ctx, pruneFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
