package transport

import (
	"bootcampProject/additional_information/internal/domain"
	pb "bootcampProject/additional_information/proto"
	"bootcampProject/utils"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type AdditionalInfoEnpointsGRPC struct {
	CreateAdditionalInfo endpoint.Endpoint
}

func MakeEndpointsGRPC(s domain.UserAdditionalInfoService) AdditionalInfoEnpointsGRPC {
	return AdditionalInfoEnpointsGRPC{
		CreateAdditionalInfo: makeCreateAdditionalInfoGRPCEndpoints(s),
	}
}

func makeCreateAdditionalInfoGRPCEndpoints(s domain.UserAdditionalInfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(*pb.CreateAdditionalInfoReq); ok {
			message, err := s.CreateAdditionalInfo(ctx, domain.UserAdditionalInfo{
				UserId:         int(req.GetUserId()),
				AdditionalInfo: req.GetAdditionalInfo(),
			})
			return CreateAdditionalInfoResponse{
				Message: message,
			}, err
		}
		return CreateAdditionalInfoResponse{
			Err: utils.NewErrBadRequest(),
		}, utils.NewErrBadRequest()
	}
}
