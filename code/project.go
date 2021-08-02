package code

import (
	"github.com/rob-woerner/protos/common"
	"github.com/rob-woerner/protos/messages"
)

func MyFunc() messages.MyMessage {
	return messages.MyMessage { state: common.MyState_Invalid }
}
