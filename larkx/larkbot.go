package larkx

import (
	"fmt"
	"github.com/go-lark/lark"
	"github.com/go-lark/lark/card"
	"github.com/happy-go-play/go-kit/larkx/internal"
	"time"
)

var ErrLarkBotRateLimitExceeded = fmt.Errorf("lark bot rate limit exceeded, please try again later")
var ErrLarkAPIRateLimitExceeded = fmt.Errorf("lark api rate limit exceeded, please try again later")

type LarkBot struct {
	webhook string
	secret  string

	limiter *internal.RateLimiter
}

func NewLarkBot(larkBotConf LarkBotConfig) *LarkBot {
	// https://open.larksuite.com/document/client-docs/bot-v3/add-custom-bot
	// - 自定义机器人的频率控制和普通应用不同，为 100 次/分钟，5 次/秒。(The frequency control of the customized robot is different from the normal application, which is 100 times/minute and 5 times/second.)
	//   每个自定义机器人的频率控制是独立的。
	// - 请求体大小不能超过 20 K。
	return &LarkBot{
		webhook: larkBotConf.Webhook,
		secret:  larkBotConf.Secret,

		limiter: internal.NewRateLimiter(5, 100), // 5次/秒, 100次/分钟
	}
}

func NewLarkBotWithLimiter(larkBotConf LarkBotConfig, maxSecondsPerRequest, maxRequestsPerMinute int) *LarkBot {
	return &LarkBot{
		webhook: larkBotConf.Webhook,
		secret:  larkBotConf.Secret,

		limiter: internal.NewRateLimiter(maxSecondsPerRequest, maxRequestsPerMinute), // 对于一个自定义机器人被共享使用时，需要自定义速率限制器
	}
}

func (a LarkBot) SendTextMessage(text string) error {
	if !a.limiter.Allow() {
		return ErrLarkBotRateLimitExceeded
	}

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
	err = checkResponseCode(resp)
	if err != nil {
		return err
	}
	return nil
}

// SendMarkdownMessageCard sends a markdown message with a title.
// The body should be in Markdown format.
// The title is the card title, is text only, not Markdown.
func (a LarkBot) SendMarkdownMessageCard(body, title string) error {
	if !a.limiter.Allow() {
		return ErrLarkBotRateLimitExceeded
	}

	b := lark.NewCardBuilder()
	cardBlock := b.Card(
		b.Markdown(body),
	).Title(title).Purple()
	return a.sendMessageCard(cardBlock)
}

func (a LarkBot) SendMessageCard(cardBlock *card.Block) error {
	if !a.limiter.Allow() {
		return ErrLarkBotRateLimitExceeded
	}

	return a.sendMessageCard(cardBlock)
}

func (a LarkBot) sendMessageCard(cardBlock *card.Block) error {
	bot := lark.NewNotificationBot(a.webhook)
	msgV4 := lark.NewMsgBuffer(lark.MsgInteractive)
	msgBuffer := msgV4.Card(cardBlock.String())
	resp, err := bot.PostNotificationV2(msgBuffer.Build())
	if err != nil {
		return fmt.Errorf("lark bot.PostNotificationV2 error: %w", err)
	}
	err = checkResponseCode(resp)
	if err != nil {
		return err
	}
	return nil
}

func checkResponseCode(resp *lark.PostNotificationV2Resp) error {
	if resp.Code != 0 {
		//    resp.Code: 9499, resp.Msg: too many request, resp.StatusCode: 0, resp.StatusMessage:
		//    resp.Code: 11233, resp.Msg: create message chat trigger rate limit, resp.StatusCode: 0, resp.StatusMessage:
		if resp.Code == 9499 || resp.Code == 11233 {
			return ErrLarkAPIRateLimitExceeded
		}
		return fmt.Errorf("resp.Code: %v, resp.Msg: %s, resp.StatusCode: %v, resp.StatusMessage: %s", resp.Code, resp.Msg, resp.StatusCode, resp.StatusMessage)
	}
	return nil
}
