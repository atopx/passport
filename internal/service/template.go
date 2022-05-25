package service

import (
	"context"
	"gorm.io/gorm"
	"template/protocol"
)

type TemplateService struct {
	db *gorm.DB
}

func NewTemplateService(db *gorm.DB) *TemplateService {
	return &TemplateService{db: db}
}

func (srv TemplateService) Create(ctx context.Context, param *protocol.CreateTemplateParam) (*protocol.CreateTemplateReply, error) {
	return NewActionWithContext(ctx, param.Header).CreateTemplate(param)
}

func (srv TemplateService) Get(ctx context.Context, param *protocol.GetTemplateParam) (*protocol.GetTemplateReply, error) {
	return NewActionWithContext(ctx, param.Header).GetTemplate(param)
}
