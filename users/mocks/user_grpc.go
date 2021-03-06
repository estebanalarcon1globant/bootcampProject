package mocks

import (
	pb "bootcampProject/proto"
	"context"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type UserGrpcMock struct {
	mock.Mock
}

func (_m *UserGrpcMock) CreateUser(_a0 context.Context, _a1 *pb.CreateUserReq, _ ...grpc.CallOption) (*pb.CreateUserResp, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.CreateUserResp
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreateUserReq) *pb.CreateUserResp); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(*pb.CreateUserResp)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreateUserReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetUsers provides a mock function with given fields: _a0, _a1, _a2
func (_m *UserGrpcMock) GetUsers(_a0 context.Context, _a1 *pb.GetUsersReq, _ ...grpc.CallOption) (*pb.GetUsersResp, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.GetUsersResp
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetUsersReq) *pb.GetUsersResp); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(*pb.GetUsersResp)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetUsersReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *UserGrpcMock) ServeGRPC(_a0 context.Context, _a1 interface{}) (context.Context, interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) context.Context); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(context.Context)
	}

	var r1 interface{}
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) interface{}); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(2)
	}
	return r0, r1, r2
}
