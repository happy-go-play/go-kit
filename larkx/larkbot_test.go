package larkx

import (
	"fmt"
	"github.com/go-lark/lark"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const (
	webhook = ""
	secret  = ""
)

func TestSendTextMessage(t *testing.T) {
	larkBotConf := LarkBotConfig{
		Webhook: webhook,
		Secret:  secret,
	}
	bot := NewLarkBot(larkBotConf)
	err := bot.SendTextMessage("Hello, World!")
	require.NoError(t, err)
}

func TestNewLarkBotWithLimiter(t *testing.T) {
	larkBotConf := LarkBotConfig{
		Webhook: webhook,
		Secret:  secret,
	}
	bot := NewLarkBotWithLimiter(larkBotConf, 1, 2)
	failedCount := 0
	for i := 0; i < 10; i++ {
		err := bot.SendTextMessage(fmt.Sprintf("[%s]Hello, World!", time.Now().Format(time.RFC3339Nano)))
		if err != nil {
			failedCount++
			require.ErrorIs(t, err, ErrLarkBotRateLimitExceeded, "expected rate limit error, got: %v", err)
		}
	}
	t.Logf("Failed count: %d", failedCount)
}

func TestSendMessageCard(t *testing.T) {
	b := lark.NewCardBuilder()
	c := b.Card(
		b.Markdown("🔥 **抖音文创“双十一”全年最优惠最低价活动今日开启** 🔥 \n🔥跨店每满300-30（可无限叠加）\n🔥店铺优惠可以和平台满减叠加：满199-20（叠加跨店满减，可以满300-50哦）").
			Href("urlVal",
				b.URL().Href("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
			),
		b.Img("img_v2_dae0a058-ca49-4a69-911c-2b27984f66eg"),
		b.Hr(),
		b.Markdown("**🌟 特别福利：**\n下单加1元就送价值11元抖音热门梗文件夹，数量有限，每ID限1件").
			Href("urlVal", b.URL().Href("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN")),
		b.Action(
			b.Button(b.Text("立即抢购")).Primary().URL("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
			b.Button(b.Text("查看更多优惠券")).URL("https://open.feishu.cn/document/ukTMukTMukTM/uczM3QjL3MzN04yNzcDN"),
		),
		b.Hr(),
		b.Note().
			AddText(b.Text("活动时间：2021年11月1日~2021年11月20日")),
	).Title("\U0001F973 抖音文创“双十一”年度大促").Purple()

	conf := LarkBotConfig{
		Webhook: webhook,
		Secret:  secret,
	}
	client := NewLarkBot(conf)
	err := client.SendMessageCard(c)
	require.NoError(t, err, "error: %v", err)
}

func TestSendMarkdownMessage(t *testing.T) {
	title := "如果不需要显示标题的话，标题可以省略"
	msg := `
🔥 **抖音文创“双十一”全年最优惠最低价活动今日开启** 🔥 
🔥跨店每满300-30（可无限叠加）
🔥店铺优惠可以和平台满减叠加：满199-20（叠加跨店满减，可以满300-50哦）
`
	conf := LarkBotConfig{
		Webhook: webhook,
		Secret:  secret,
	}
	client := NewLarkBot(conf)
	err := client.SendMarkdownMessageCard(msg, title)
	require.NoError(t, err, "error: %v", err)
}
