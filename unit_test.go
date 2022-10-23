package serverChan

import (
	"testing"
)

func TestSend(t *testing.T) {
	msg := &Req{
		Title:   "任务成功完成",
		Desp:    "任务共节省了32PB的空间",
		Short:   "任务摘要", // 看起来不支持
		Channel: "9",    // 看起来不支持两个通道
	}
	key := "xxxx"
	recive, err := Send(msg, key)
	if err != nil {
		return
	}
	t.Log(recive)
}
