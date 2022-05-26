package model

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
)

type Password struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Keyword   string `json:"keyword"`
	Value     string `json:"value"`
	Domain    string `json:"domain"`
	Encrypted bool   `json:"encrypted"`
	Deleter   int64  `json:"deleter"`
	Creator   int64  `json:"creator"`
	Updater   int64  `json:"updater"`
	DeleteAt  int64  `json:"delete_at"`
	CreateAt  int64  `json:"create_at"`
	UpdateAt  int64  `json:"update_at"`
}

func (*Password) TableName() string {
	return "password"
}

func (p *Password) BeforeCreate(*gorm.DB) error {
	if p.Encrypted {
		ctx := md5.New()
		ctx.Write([]byte(p.Value))
		p.Value = hex.EncodeToString(ctx.Sum(nil))
	}
	return nil
}
