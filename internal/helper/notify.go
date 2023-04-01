package helper

import (
	"bytes"
	"log"

	"arb-template/config"

	fsBotAPI "github.com/electricbubble/feishu-bot-api"
)

// webhook

const (
	MsgTypeSuccess = iota
	MsgTypeWarning
	MsgTypeError
)

type NotifyMsg struct {
	MsgType  int
	Title    string
	SubTitle string
	Content  string
}

// NotifyToFeiShu send message to feishu
func NotifyToFeiShu(msg NotifyMsg) {
	bot := fsBotAPI.NewBot(config.Conf.FeishuWebhook, fsBotAPI.WithSecretKey(config.Conf.FeishuSecretKey))
	buf := bytes.NewBufferString(msg.Title + "\n")
	buf.WriteString(msg.SubTitle + "\n")
	buf.WriteString(msg.Content + "\n")
	err := bot.PushText(buf.String())
	if err != nil {
		log.Fatalln("Notify Failed!", err)
	}
}

// BuildSuccessMsg build success message
func BuildSuccessMsg(msgType int, title string, content string) NotifyMsg {
	return NotifyMsg{
		MsgType: msgType,
		Title:   title,
		Content: content,
	}
}
