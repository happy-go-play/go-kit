package larkx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	larkBotConf := LarkBotConfig{
		Webhook: "...",
		Secret:  "",
	}
	bot := NewLarkBot(larkBotConf)
	err := bot.SendTextMessage("Hello, World!")
	require.NoError(t, err)
}
