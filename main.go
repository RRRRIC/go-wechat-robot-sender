package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	Key    string
	Msg    string
	client *http.Client
)

func init() {
	client = &http.Client{
		Timeout: time.Second * 5,
	}
}

func main() {
	flag.StringVar(&Key, "k", "", "the robot key")

	flag.StringVar(&Msg, "m", "test-msg", "the plain msg sent to wechat-enterprise")
	flag.Parse()

	if Key == "" {
		log.Println("Must assign robot key")
	}

	sentMsgToWechat(Key, Msg)
}

type WechatMsg struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func sentMsgToWechat(key, msg string) {
	robotUrl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + key
	sentMsg := &WechatMsg{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{Content: msg},
	}

	var err error
	defer func() {
		if err != nil {
			log.Println("Can't send msg to wechat, err :", err)
		}
	}()

	jsoned, err := json.Marshal(sentMsg)
	request, err := http.NewRequest("POST", robotUrl, bytes.NewReader(jsoned))

	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)
	defer func() {
		if err != nil {
			log.Println("Can't send msg to wechat, err :", err)
		} else {
			res.Body.Close()
		}
	}()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println("Sent Msg wrong status , body :", string(body))
		}
	}
	log.Println("Succeed sent msg to wechat , key : ", key, " , msg : ", msg)
}
