package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tltorress/simpleBot/cmd/config"
	"github.com/tltorress/simpleBot/internal/bot/commands"
	"github.com/tltorress/simpleBot/internal/platform/logger"
)

func MessageHandler() func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		logger.MyLog.Info(m.Type)

		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Type == discordgo.MessageTypeGuildMemberJoin {
			logger.MyLog.Info("Se unio un pibardo")
			resp, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				logger.MyLog.Error(err)
			}

			s.ChannelMessageSend(resp.ID, "Llevala!")

		}

		args := strings.Split(m.Content, " ")

		if args[0] != config.PREFIX {
			return
		}

		switch {
		case strings.ToLower(args[1]) == "ping":
			commands.Ping(s, m)
		case strings.ToLower(args[1]) == "rand":
			commands.Random(s, m)
		}
	}
}
