package serverChan

import (
	"testing"
)

func TestSend(t *testing.T) {
	msg := &Req{
		Title:   "通过go请求2次",
		Desp:    "这是描述",
		Short:   "这是摘要",
		Channel: "9",
	}
	key := "xxxx"
	recive, err := Send(msg, key)
	if err != nil {
		return
	}
	t.Log(recive)
}
