package dingtalkbot

import "github.com/happy-go-play/go-kit/dingtalkbot/internal"

type DingtalkBotConfig struct {
	Webhook string
	Secret  string
}

type AtOption struct {
	IsAtAll   bool
	AtUserIds []string
	AtMobiles []string
}

// AtOpt 函数式可选参数
type AtOpt func(a *internal.At)

func AtAll() AtOpt {
	return func(a *internal.At) {
		a.IsAtAll = true
	}
}

func AtUserIds(ids ...string) AtOpt {
	return func(a *internal.At) {
		a.AtUserIds = append(a.AtUserIds, ids...)
	}
}

func AtMobiles(mobiles ...string) AtOpt {
	return func(a *internal.At) {
		a.AtMobiles = append(a.AtMobiles, mobiles...)
	}
}

func buildAt(opts []AtOpt) *internal.At {
	if len(opts) == 0 {
		return nil
	}
	at := &internal.At{}
	for _, opt := range opts {
		opt(at)
	}
	// 若未设置任何字段则返回 nil
	if !at.IsAtAll && len(at.AtUserIds) == 0 && len(at.AtMobiles) == 0 {
		return nil
	}
	return at
}
