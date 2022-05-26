package common

import (
	"context"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"passport/protocol"
)

type Action struct {
	ctx     context.Context
	traceId int64
	userId  int64
	client  string

	db *gorm.DB
}

func (action *Action) SetContext(ctx context.Context, header *protocol.RequestHeader) {
	action.ctx = ctx
	action.traceId = header.GetTraceId()
	action.userId = header.GetUserId()
	action.client = header.GetClient()
}

func (action *Action) SetDatabase(db *gorm.DB) {
	action.db = db.WithContext(action.ctx)
}

func (action *Action) GetTraceId() int64 {
	return action.traceId
}

func (action *Action) GetUserId() int64 {
	return action.userId
}

func (action *Action) GetClient() string {
	return action.client
}

func (action *Action) GetDatabase() *gorm.DB {
	return action.db
}

func (action *Action) NewOkResponseHeader() *protocol.ResponseHeader {
	return &protocol.ResponseHeader{
		TraceId: action.GetTraceId(),
		Code:    codes.OK.String(),
		Message: "success",
	}
}

func (action *Action) NewResponseHeader(code codes.Code, message string) *protocol.ResponseHeader {
	return &protocol.ResponseHeader{
		TraceId: action.GetTraceId(),
		Code:    code.String(),
		Message: message,
	}
}
