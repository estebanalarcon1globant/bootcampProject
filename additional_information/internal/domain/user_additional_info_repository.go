package domain

import "context"

type AdditionalInfoRepository interface {
	CreateAdditionalInfo(ctx context.Context, info UserAdditionalInfo) (string, error)
}
