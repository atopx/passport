package service

import (
	"context"
	"template/internal/model"
	"template/protocol"
)

type Action struct {
	model.Action
}

func NewActionWithContext(ctx context.Context, header *protocol.Header) *Action {
	action := Action{}
	action.SetContext(ctx, header)
	return &action
}
