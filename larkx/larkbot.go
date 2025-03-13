package larkx

import (
	"fmt"
	"github.com/go-lark/lark"
	"time"
)

type LarkBot struct {
	webhook string
	secret  string
}

func NewLarkBot(larkBotConf LarkBotConfig) *LarkBot {
	return &LarkBot{
		webhook: larkBotConf.Webhook,
		secret:  larkBotConf.Secret,
	}
}

func (a LarkBot) SendTextMessage(text string) error {
	// https://open.larksuite.com/document/client-docs/bot-v3/add-custom-bot
	// The frequency control of the customized robot is different from the normal application, which is 100 times/minute and 5 times/second.
	// 自定义机器人的频率控制和普通应用不同，为 100 次/分钟，5 次/秒。
	bot := lark.NewNotificationBot(a.webhook)
	msgBuffer := lark.NewMsgBuffer(lark.MsgText)
	if a.secret != "" {
		msgBuffer = msgBuffer.WithSign(a.secret, time.Now().Unix())
	}
	msgBuffer = msgBuffer.Text(text)
	resp, err := bot.PostNotificationV2(msgBuffer.Build())
	if err != nil {
		return fmt.Errorf("lark bot.PostNotificationV2 error: %w", err)
	}
	if resp.Code != 0 {
		return fmt.Errorf("resp.Code: %v, resp.Msg: %s, resp.StatusCode: %v, resp.StatusMessage: %s", resp.Code, resp.Msg, resp.StatusCode, resp.StatusMessage)
	}
	return nil
}
