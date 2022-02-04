package domain

type Users struct {
	ID      int    `json:"id"`
	PwdHash string `json:"pwd_hash"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	//ParentID  int
}
