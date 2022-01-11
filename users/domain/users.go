package domain

import "context"

type Users struct {
	ID      int    `json:"id"`
	PwdHash string `json:"pwd_hash"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	//ParentID  int
}

type UserRepository interface {
	CreateUser(ctx context.Context, user Users) (int, error)
}
