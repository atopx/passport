package action

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"passport/internal/model"
	"passport/protocol"
)

func (action *Action) CreateUser(param *protocol.CreateUserParam) (*protocol.CreateUserReply, error) {
	db := action.GetDatabase().Begin()
	user := model.User{Account: param.Username}
	if err := db.Model(&user).Where(&user).First(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		// user already exists
		header := action.NewResponseHeader(codes.AlreadyExists, "user already exists")
		return &protocol.CreateUserReply{Header: header}, nil
	}
	user.Rule = protocol.Rule_RULE_NONE.String()
	if err := db.Create(&user); err != nil {
		header := action.NewResponseHeader(codes.Internal, "server internal error")
		return &protocol.CreateUserReply{Header: header}, nil
	}
	password := model.Password{
		UserId:    user.Id,
		Keyword:   param.Keyword,
		Domain:    param.Domain,
		Value:     param.Password,
		Encrypted: param.Encrypted,
		Creator:   action.GetUserId(),
		Updater:   action.GetUserId(),
	}
	if err := db.Create(&password); err != nil {
		header := action.NewResponseHeader(codes.Internal, "server internal error")
		return &protocol.CreateUserReply{Header: header}, nil
	}
	db.Commit()
	header := action.NewOkResponseHeader()
	return &protocol.CreateUserReply{Header: header, User: &protocol.User{
		Id:       user.Id,
		Account:  user.Account,
		Rule:     param.Rule,
		LoginAt:  user.LoginAt,
		DeleteAt: user.DeleteAt,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}}, nil
}
