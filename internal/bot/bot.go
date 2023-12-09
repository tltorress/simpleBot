package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/tltorress/simpleBot/cmd/config"
	"github.com/tltorress/simpleBot/internal/bot/handler"
	"github.com/tltorress/simpleBot/internal/platform/logger"
)

func StartBot() error {
	dg, err := discordgo.New(fmt.Sprintf("Bot %s", config.BOT_TOKEN))
	if err != nil {
		logger.MyLog.Error("Can't create the bot: ", err)
		return err
	}

	dg.AddHandler(handler.MessageHandler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		logger.MyLog.Error("Can't stablish connection with bot: ", err)
	}
	defer dg.Close()

	logger.MyLog.Info("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return nil
}
