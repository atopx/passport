package service

import (
	"fmt"
	"template/protocol"
)

func (action *Action) GetTemplate(param *protocol.GetTemplateParam) (*protocol.GetTemplateReply, error) {
	fmt.Println(param)
	return nil, nil
}
