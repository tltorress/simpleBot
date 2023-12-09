package main

import (
	"github.com/tltorress/simpleBot/internal/bot"
	"github.com/tltorress/simpleBot/internal/platform/logger"
)

func main() {
	logger.InitLogger()

	bot.StartBot()
}
