package service

import (
	"fmt"
	"template/protocol"
)

func (action *Action) CreateTemplate(param *protocol.CreateTemplateParam) (*protocol.CreateTemplateReply, error) {
	fmt.Println(param)
	return nil, nil
}
