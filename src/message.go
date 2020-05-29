package src

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	cmd := m.Content

	if existName := strings.Index(cmd, "AHOXA"); existName == -1 {
		return
	}

	_, Err := s.ChannelMessageSend(m.ChannelID, "こんにちは、あほくさですwwwwww")
	if Err != nil {
		log.Println("Cannot send message:", Err)
		return
	}
}
