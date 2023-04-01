package tests

import (
	"bytes"
	fsBotAPI "github.com/electricbubble/feishu-bot-api"
	"log"
	"testing"
)

func TestNotifyToFeiShu(t *testing.T) {
	webhook := ""

	// 开启签名校验时使用
	secretKey := ""

	bot := fsBotAPI.NewBot(webhook, fsBotAPI.WithSecretKey(secretKey))

	{ // 发送普通文本消息
		buf := bytes.NewBufferString("💰套利成功提醒！\n")
		buf.WriteString("🤓一笔新收入到账！" + "\n")
		err := bot.PushText(buf.String())
		if err != nil {
			log.Fatalln(err)
		}
	}

}
