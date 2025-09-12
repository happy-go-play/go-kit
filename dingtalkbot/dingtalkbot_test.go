package dingtalkbot

import (
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	webhook   = "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxx"
	secret    = "SECyyyyyyyy"
	atUserIds []string
	atMobiles []string
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func TestSendTextMessage(t *testing.T) {
	larkBotConf := DingtalkBotConfig{
		Webhook: webhook,
		Secret:  secret,
	}
	bot := NewDingtalkBot(larkBotConf)
	{
		err := bot.SendTextMessage("Hello, World!")
		require.NoError(t, err)
	}

	time.Sleep(time.Millisecond * 500)

	{
		atOptUserIds := AtUserIds(atUserIds...)
		atOptMobiles := AtMobiles(atMobiles...)
		err := bot.SendTextMessage("Hello, World!", atOptMobiles, atOptUserIds)
		require.NoError(t, err)
	}

	time.Sleep(time.Millisecond * 500)

	{
		atOptAll := AtAll()
		err := bot.SendTextMessage("Hello, World!", atOptAll)
		require.NoError(t, err)
	}
}

func TestSendMarkdownMessage(t *testing.T) {
	larkBotConf := DingtalkBotConfig{
		Webhook: webhook,
		Secret:  secret,
	}
	bot := NewDingtalkBot(larkBotConf)
	err := bot.SendMarkdownMessage("Title", "**Hello**, World!\n- 111\n- 222")
	require.NoError(t, err)
}
