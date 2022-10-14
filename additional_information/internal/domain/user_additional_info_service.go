package domain

import "context"

type UserAdditionalInfoService interface {
	CreateAdditionalInfo(ctx context.Context, additionalInfo UserAdditionalInfo) (string, error)
}
