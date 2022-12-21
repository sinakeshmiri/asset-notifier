package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Adapter struct {
	bot *tgbotapi.BotAPI
	ids []int64
}

// NewAdapter creates a new Adapter
func NewAdapter(apiToken string, ids []int64) (*Adapter, error) {
	// connect
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Fatalf("failed to connect to Telegram: %v", err)
		return nil, err
	}

	return &Adapter{bot: bot, ids: ids}, nil
}

func (ta Adapter) SendNotif(messages []string) error {
	for _, id := range ta.ids {
		for _, m := range messages {
			msg := tgbotapi.NewMessage(id, m)
			if _, err := ta.bot.Send(msg); err != nil {
				return err
			}
		}
	}
	return nil
}
