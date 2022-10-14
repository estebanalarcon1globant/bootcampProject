package service

import (
	"bootcampProject/additional_information/internal/domain"
	"context"
)

type userAdditionalInfoService struct {
	additionalInfoRepo domain.AdditionalInfoRepository
}

func NewUserAdditionalInfoService(additionalInfoRepo domain.AdditionalInfoRepository) domain.UserAdditionalInfoService {
	return &userAdditionalInfoService{
		additionalInfoRepo: additionalInfoRepo,
	}
}

func (s *userAdditionalInfoService) CreateAdditionalInfo(ctx context.Context,
	info domain.UserAdditionalInfo) (string, error) {
	return s.additionalInfoRepo.CreateAdditionalInfo(ctx, info)
}
