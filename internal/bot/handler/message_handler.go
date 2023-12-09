package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tltorress/simpleBot/cmd/config"
	"github.com/tltorress/simpleBot/internal/bot/commands"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(m.Content, " ")

	if args[0] != config.PREFIX {
		return
	}

	if strings.ToLower(args[1]) == "ping" {
		commands.Ping(s, m)
	}

}
