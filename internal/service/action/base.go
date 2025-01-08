package action

import (
	"context"
	"passport/internal/common"
	"passport/protocol"

	"gorm.io/gorm"
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
