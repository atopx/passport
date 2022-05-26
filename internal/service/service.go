package service

import (
	"context"
	"gorm.io/gorm"
	"passport/internal/service/action"
	"passport/protocol"
)

type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (srv Service) Health(ctx context.Context, param *protocol.HealthParam) (*protocol.HealthParamReply, error) {
	header := action.NewActionWithContext(ctx, srv.db, param.Header).NewOkResponseHeader()
	return &protocol.HealthParamReply{Header: header}, nil
}

func (srv Service) CreateUser(ctx context.Context, param *protocol.CreateUserParam) (*protocol.CreateUserReply, error) {

	return action.NewActionWithContext(ctx, srv.db, param.Header).CreateUser(param)
}

func (srv Service) DeleteUser(ctx context.Context, param *protocol.DeleteUserParam) (*protocol.DeleteUserReply, error) {
	return action.NewActionWithContext(ctx, srv.db, param.Header).DeleteUser(param)
}

func (srv Service) UpdateUser(ctx context.Context, param *protocol.UpdateUserParam) (*protocol.UpdateUserReply, error) {
	return action.NewActionWithContext(ctx, srv.db, param.Header).UpdateUser(param)
}

func (srv Service) SearchUser(ctx context.Context, param *protocol.SearchUserParam) (*protocol.SearchUserReply, error) {
	return action.NewActionWithContext(ctx, srv.db, param.Header).SearchUser(param)
}

func (srv Service) GetUserById(ctx context.Context, param *protocol.GetUserByIdParam) (*protocol.GetUserByIdReply, error) {
	return action.NewActionWithContext(ctx, srv.db, param.Header).GetUserById(param)
}

func (srv Service) GetUserByToken(ctx context.Context, param *protocol.GetUserByTokenParam) (*protocol.GetUserByTokenReply, error) {
	return action.NewActionWithContext(ctx, srv.db, param.Header).GetUserByToken(param)
}

func (srv Service) SignIn(ctx context.Context, param *protocol.SignInParam) (*protocol.SignInReply, error) {
	return action.NewActionWithContext(ctx, srv.db, param.Header).SignIn(param)
}

func (srv Service) SignOut(ctx context.Context, param *protocol.SignOutParam) (*protocol.SignOutReply, error) {
	return action.NewActionWithContext(ctx, srv.db, param.Header).SignOut(param)
}
