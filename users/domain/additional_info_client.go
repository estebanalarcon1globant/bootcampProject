package domain

import (
	pb "bootcampProject/additional_information/proto"
	"context"
	"google.golang.org/grpc"
)

type AdditionalInformationClient interface {
	Close()
	CreateAdditionalInfo(ctx context.Context,
		in *pb.CreateAdditionalInfoReq,
		opts ...grpc.CallOption) (*pb.CreateAdditionalInfoResp, error)
}
