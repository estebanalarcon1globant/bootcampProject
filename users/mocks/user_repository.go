package mocks

import (
	"bootcampProject/users/domain"
	"context"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *UserRepoMock) CreateUser(_a0 context.Context, _a1 domain.Users) (int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, domain.Users) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Users) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetUsers provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserRepoMock) GetUsers(_a0 context.Context, _a1 int, _a2 int) ([]domain.Users, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []domain.Users
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []domain.Users); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).([]domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *UserRepoMock) GetUserByEmail(_a0 context.Context, _a1 string) (domain.Users, error) {
	ret := _m.Called(_a0, _a1)

	var r0 domain.Users
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Users); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(domain.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}
