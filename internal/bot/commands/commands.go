package commands

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}

func Random(s *discordgo.Session, m *discordgo.MessageCreate) {

	num := rand.Intn(1000)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("El numero random es: %d", num))
}
