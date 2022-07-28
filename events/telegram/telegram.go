package telegram

import "github.com/Includeoyasi/tgbot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}

func New()
