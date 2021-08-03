package code

import (
	common "github.com/rob-woerner/protos/common"
	messages "github.com/rob-woerner/protos/messages"
)

func MyFunc() messages.MyRequest {
	return messages.MyRequest { State: common.MyState_LOGGED_OUT }
}
