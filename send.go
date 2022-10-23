package serverChan

import (
	"encoding/json"
	"github.com/zhangyiming748/serverChan/util/log"
	"github.com/zhangyiming748/serverChan/util/net"
	"strings"
)

type Req struct {
	Title   string `json:"title"` // 标题
	Desp    string `json:"desp"`  // 描述
	Short   string `json:"short"` //这是摘要
	Channel string `json:"channel"`
}
type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Pushid  string `json:"pushid"`
		Readkey string `json:"readkey"`
		Error   string `json:"error"`
		Errno   int    `json:"errno"`
	} `json:"data"`
}

func Send(content *Req, SendKey string) (Res, error) {
	header := map[string]string{
		"User-Agent":   "apifox/1.0.0 (https://www.apifox.cn)",
		"Content-Type": "application/json",
	}
	data := map[string]string{
		"title":   content.Title,
		"desp":    content.Desp,
		"short":   content.Short,
		"channel": content.Channel,
	}
	url := strings.Join([]string{"https://sctapi.ftqq.com/", SendKey, ".send"}, "")
	res, err := net.HttpPostJson(header, data, url)
	if err != nil {
		log.Debug.Printf("此次请求发送失败:%+v\n", content)
		return Res{}, err
	}
	var result Res
	if err := json.Unmarshal(res, &result); err != nil {
		log.Debug.Printf("此次请求获取返回值失败:%+v\n", content)
		return Res{}, err
	} else {
		log.Debug.Printf("此次请求返回的状态信息:%+v\n", result)
		return result, nil
	}
}
