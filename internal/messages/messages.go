// Messages package is responsible for messages handling and reacting
package messages

import (
	"github.com/bwmarrin/discordgo"
	"github.com/yuno-obsessed/shikimori/internal/logs"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// if m.Author.ID == s.State.User.ID {
	// 	return
	// }

	message := m.Message.Author.Username + ": " + m.Message.Content
	// message = strings.ToLower(m.Content)
	logs.Log(message)
}
