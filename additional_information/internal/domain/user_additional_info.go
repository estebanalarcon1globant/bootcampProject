package domain

type UserAdditionalInfo struct {
	UserId         int    `json:"user_id"`
	AdditionalInfo string `json:"additional_info"`
}
