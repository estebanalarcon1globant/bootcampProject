package domain

type Users struct {
	ID      int    `json:"id" validate:"-"`
	PwdHash string `json:" pwd_hash" validate:"nonzero"`
	Name    string `json:"name" validate:"nonzero"`
	Age     int    `json:"age" validate:"nonzero"`
	Email   string `json:"email" validate:"nonzero, regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	//ParentID  int
}
