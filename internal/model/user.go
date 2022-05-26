package model

type User struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`
	Rule     string `json:"rule"`
	LoginAt  int64  `json:"login_at"`
	DeleteAt int64  `json:"delete_at"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

func (*User) TableName() string {
	return "user"
}
