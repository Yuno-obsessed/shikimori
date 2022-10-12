// Messages package is responsible for messages handling and reacting
package messages

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	ds "github.com/bwmarrin/discordgo"
	"github.com/yuno-obsessed/shikimori/internal/logs"
)

func MessageCreate(s *ds.Session, m *ds.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	//s.ChannelPermissionSet("825200883087573003",s.State.User.ID,1,true,false)
	err := s.UpdateGameStatus(0, "Waifuborn")
	if err != nil {
		fmt.Println(err)
	}

	// Logging messages to standard output
	if m.ChannelID == "825200883087573003" {
		logs.LogProviding(m.Message.Author.Username + " глаголит истину: " + m.Message.Content)
	}

	// randomize output
	rand.Seed(time.Now().Unix())
	// convert string in lowercase
	m.Content = strings.ToLower(m.Content)

	if m.Author.ID == "369852581972803584" {
		reference := m.Reference()
		_, err := s.ChannelMessageSendReply(m.ChannelID, "some cool phrase", reference)
		if err != nil {
			fmt.Println(err)
		}
	}
}
