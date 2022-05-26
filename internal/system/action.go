package system

import (
	"context"
	"gorm.io/gorm"
	"template/protocol"
)

type Action struct {
	ctx     context.Context
	traceId int64
	userId  int64
	client  string

	DB *gorm.DB
}

func (action *Action) SetContext(ctx context.Context, header *protocol.Header) {
	action.ctx = ctx
	action.traceId = header.GetTraceId()
	action.userId = header.GetUserId()
	action.client = header.GetClient()
}

func (action *Action) SetDatabase(db *gorm.DB) {
	action.DB = db.WithContext(action.ctx)
}

func (action *Action) GetTraceId() int64 {
	return action.traceId
}

func (action *Action) GetUserId() int64 {
	return action.userId
}
