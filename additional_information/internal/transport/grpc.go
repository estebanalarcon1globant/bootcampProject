package transport

import (
	pb "bootcampProject/additional_information/proto"
	"bootcampProject/utils"
	"context"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	pb.UnimplementedAdditionalInformationServiceServer
	createAdditionalInfo gt.Handler
}

func NewAdditionalInfoGRPCServer(svcEndpoints AdditionalInfoEnpointsGRPC) pb.AdditionalInformationServiceServer {

	return &gRPCServer{
		createAdditionalInfo: gt.NewServer(
			svcEndpoints.CreateAdditionalInfo,
			decodeCreateAdditionalInfoGRPCRequest,
			encodeCreateAdditionalInfoGRPCResponse,
		),
	}
}

func (s *gRPCServer) CreateAdditionalInfo(ctx context.Context, req *pb.CreateAdditionalInfoReq) (*pb.CreateAdditionalInfoResp, error) {
	_, resp, err := s.createAdditionalInfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateAdditionalInfoResp), nil
}

func decodeCreateAdditionalInfoGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	if req, ok := request.(*pb.CreateAdditionalInfoReq); ok {
		return &pb.CreateAdditionalInfoReq{
			UserId:         req.GetUserId(),
			AdditionalInfo: req.GetAdditionalInfo(),
		}, nil
	}
	return &pb.CreateAdditionalInfoReq{}, utils.NewErrBadRequest()
}

func encodeCreateAdditionalInfoGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	if resp, ok := response.(CreateAdditionalInfoResponse); ok {
		return &pb.CreateAdditionalInfoResp{
			Message: resp.Message,
		}, resp.Err
	}
	return &pb.CreateAdditionalInfoResp{Error: utils.NewErrBadRequest().Error()}, utils.NewErrBadRequest()
}
