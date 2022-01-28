package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	Key         string
	Msg         string
	Member_list []string
	Mobile_list []string
	client      *http.Client
)

func init() {
	client = &http.Client{
		Timeout: time.Second * 5,
	}
}

func main() {
	var mobileString string
	var memberString string

	flag.StringVar(&Key, "k", "", "the robot key")
	flag.StringVar(&Msg, "m", "test-msg", "the plain msg sent to wechat-enterprise")
	flag.StringVar(&mobileString, "mobile", "", "the metioned users' mobile , multi-user split by ','")
	flag.StringVar(&memberString, "member", "", "the metioned users' id , multi-user split by ','")
	flag.Parse()

	Mobile_list = strings.Split(mobileString, ",")
	Member_list = strings.Split(memberString, ",")

	if Key == "" {
		log.Println("Must assign robot key")
	}

	sentMsgToWechat()
}

type WechatMsg struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content               string   `json:"content"`
		Mentioned_list        []string `json:"mentioned_list"`
		Mentioned_mobile_list []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

func sentMsgToWechat() {
	robotUrl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + Key
	sentMsg := &WechatMsg{
		MsgType: "text",
		Text: struct {
			Content               string   `json:"content"`
			Mentioned_list        []string `json:"mentioned_list"`
			Mentioned_mobile_list []string `json:"mentioned_mobile_list"`
		}{Content: Msg, Mentioned_list: Member_list, Mentioned_mobile_list: Mobile_list},
	}

	var err error
	defer func() {
		if err != nil {
			log.Println("Can't send msg to wechat, err :", err)
		}
	}()

	jsoned, err := json.Marshal(sentMsg)
	if err != nil {
		return
	}
	request, err := http.NewRequest("POST", robotUrl, bytes.NewReader(jsoned))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)

	if err != nil {
		return
	} else {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println("Sent Msg wrong status , body :", string(body))
		}
	}
	log.Println("Succeed sent msg to wechat, key : ", Key, " , msg : ", Msg)
}
