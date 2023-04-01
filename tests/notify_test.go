package tests

import (
	"bytes"
	fsBotAPI "github.com/electricbubble/feishu-bot-api"
	"log"
	"testing"
)

func TestNotifyToFeiShu(t *testing.T) {
	webhook := ""

	secretKey := ""

	bot := fsBotAPI.NewBot(webhook, fsBotAPI.WithSecretKey(secretKey))

	{
		buf := bytes.NewBufferString("💰套利成功提醒！\n")
		buf.WriteString("🤓一笔新收入到账！" + "\n")
		err := bot.PushText(buf.String())
		if err != nil {
			log.Fatalln(err)
		}
	}

}
