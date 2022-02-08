// Code generated by mockery v1.0.0
package mocks

import (
	"clean-architecture-beego/internal/domain"
	"context"
	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

func (_m *Usecase) GetProducts(ctx context.Context,limit, offset int) ([]domain.Product, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []domain.Product); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Usecase) GetProductById(ctx context.Context,id uint) (*domain.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, uint) *domain.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}


	return r0, r1
}

func (_m *Usecase) SaveProduct(ctx context.Context,body domain.ProductStoreRequest) error {
	ret := _m.Called(ctx, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ProductStoreRequest) error); ok {
		r0 = rf(ctx, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Usecase) UpdateProduct(ctx context.Context,body domain.ProductUpdateRequest) error {
	ret := _m.Called(ctx, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ProductUpdateRequest) error); ok {
		r0 = rf(ctx, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
