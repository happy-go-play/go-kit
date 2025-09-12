package dingtalkbot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/happy-go-play/go-kit/dingtalkbot/internal"
)

// DingtalkBot
// https://open.dingtalk.com/document/orgapp/obtain-the-webhook-address-of-a-custom-robot
// 每个机器人每分钟最多发送20条消息到群里，如果超过20条，会限流10分钟。
type DingtalkBot struct {
	webhook string
	secret  string
}

func NewDingtalkBot(dingtalkBotConfig DingtalkBotConfig) *DingtalkBot {
	return &DingtalkBot{
		webhook: dingtalkBotConfig.Webhook,
		secret:  dingtalkBotConfig.Secret,
	}
}

func generateSign(timestamp int64, secret string) string {
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (b *DingtalkBot) SendTextMessage(text string, opts ...AtOpt) error {
	textMsg := internal.Message{
		MsgType: "text",
		Text: &internal.Text{
			Content: text,
		},
	}
	//if at != nil {
	//	textMsg.At = &internal.At{
	//		IsAtAll:   at.IsAtAll,
	//		AtUserIds: at.AtUserIds,
	//		AtMobiles: at.AtMobiles,
	//	}
	//}
	if at := buildAt(opts); at != nil {
		textMsg.At = at
	}
	return b.sendMessage(textMsg)
}

func (b *DingtalkBot) SendMarkdownMessage(title, text string, opts ...AtOpt) error {
	markdownMsg := internal.Message{
		MsgType: "markdown",
		Markdown: &internal.Markdown{
			Title: title,
			Text:  text,
		},
	}
	//if at != nil {
	//	markdownMsg.At = &internal.At{
	//		IsAtAll:   at.IsAtAll,
	//		AtUserIds: at.AtUserIds,
	//		AtMobiles: at.AtMobiles,
	//	}
	//}
	if at := buildAt(opts); at != nil {
		markdownMsg.At = at
	}
	return b.sendMessage(markdownMsg)
}

func (b *DingtalkBot) sendMessage(message internal.Message) error {
	timestamp := time.Now().UnixMilli()
	sign := generateSign(timestamp, b.secret)

	// "https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s"
	sep := "?"
	if strings.Contains(b.webhook, "?") {
		sep = "&"
	}
	webhookURL := fmt.Sprintf("%s%s&timestamp=%d&sign=%s", b.webhook, sep, timestamp, sign)

	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed. statusCode:%d", resp.StatusCode)
	}

	return nil
}
