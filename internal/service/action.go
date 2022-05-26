package service

import (
	"context"
	"template/internal/system"
	"template/protocol"
)

type Action struct {
	system.Action
}

func NewActionWithContext(ctx context.Context, header *protocol.Header) *Action {
	action := Action{}
	action.SetContext(ctx, header)
	return &action
}
