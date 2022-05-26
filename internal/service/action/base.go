package action

import (
	"context"
	"gorm.io/gorm"
	"passport/internal/common"
	"passport/protocol"
)

type Action struct {
	common.Action
}

func NewActionWithContext(ctx context.Context, db *gorm.DB, header *protocol.RequestHeader) *Action {
	action := Action{}
	action.SetContext(ctx, header)
	action.SetDatabase(db)
	return &action
}
