package service

import (
	pb "bootcampProject/additional_information/proto"
	"bootcampProject/users/domain"
	"context"
	"google.golang.org/grpc"
)

type additionalInformationClient struct {
	additionalInfoClient pb.AdditionalInformationServiceClient
	address              string
	conn                 *grpc.ClientConn
}

func NewAdditionalInformationClient(address string) (domain.AdditionalInformationClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	client := pb.NewAdditionalInformationServiceClient(conn)
	return &additionalInformationClient{
		additionalInfoClient: client,
		address:              address,
		conn:                 conn,
	}, nil
}

func (c *additionalInformationClient) Close() {
	c.conn.Close()
}

// CreateAdditionalInfo is only a service to testing integration between 2 microservices
func (c *additionalInformationClient) CreateAdditionalInfo(ctx context.Context, in *pb.CreateAdditionalInfoReq, opts ...grpc.CallOption) (*pb.CreateAdditionalInfoResp, error) {
	return c.additionalInfoClient.CreateAdditionalInfo(ctx, in, opts...)
}
