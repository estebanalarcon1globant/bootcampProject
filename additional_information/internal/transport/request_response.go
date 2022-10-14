package transport

import "bootcampProject/additional_information/internal/domain"

type CreateAdditionalInfoRequest struct {
	AdditionalInfo domain.UserAdditionalInfo
}

type CreateAdditionalInfoResponse struct {
	Message string `json:"message"`
	Err     error  `json:"error,omitempty"`
}
