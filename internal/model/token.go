package model

type Token struct {
	ID      int64  `json:"id"`
	UserId  int64  `json:"userId"`
	Client  string `json:"client"`
	Invalid bool   `json:"invalid"`
	Expire  int64  `json:"expire"`
	CreatAt int64  `json:"creat_at"`
}

func (*Token) TableName() string {
	return "token"
}
