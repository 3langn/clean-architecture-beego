// Code generated by mockery v1.0.0
package mocks

import (
	context "context"

	"clean-architecture-beego/models"

	mock "github.com/stretchr/testify/mock"
)

// repository is an autogenerated mock type for the repository type
type Repository struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Repository) List(ctx context.Context, limit, offset int) (res []models.SampleModule, err error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []models.SampleModule
	if rf, ok := ret.Get(0).(func(context.Context,  int, int) []models.SampleModule); ok {
		r0 = rf(ctx,  limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.SampleModule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context,  int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) Count(ctx context.Context) (res int, err error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
