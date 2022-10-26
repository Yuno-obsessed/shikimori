// Messages package is responsible for messages handling and reacting
package messages

import (
	"math/rand"
	"strings"
	"time"

	ds "github.com/bwmarrin/discordgo"
	"github.com/yuno-obsessed/shikimori/internal/commands"
	"github.com/yuno-obsessed/shikimori/internal/logs"
)

func MessageCreate(s *ds.Session, m *ds.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	StartLogging(m)

	if m.Content == commands.Commands[0] {
		// SelectMenu struct
		err, _ := s.ChannelMessageSend(m.ChannelID, commands.ListTags())
		if err != nil {
			logs.LogErr(logs.ErrUnableToSendMessage, "https://github.com/Yuno-obsessed/shikimori/blob/main/internal/messages/messages.go#L39")
		}
	} else if strings.Contains(m.Content, commands.Commands[2]) {
		avatarURL := commands.InsertAvatar(m)
		_, err := s.ChannelMessageSendReply(m.ChannelID, avatarURL, m.Reference())
		if err != nil {
			logs.LogErr(logs.ErrUnableToSendMessage, "https://github.com/Yuno-obsessed/shikimori/blob/main/internal/messages/messages.go#L50")
		}
	} else if m.Content == commands.Commands[3] {
		logs.LogErr(logs.ErrFuncUnavailable, "https://github.com/Yuno-obsessed/shikimori/blob/main/internal/messages/messages.go#L53")
	} else if strings.Contains(m.Content, commands.Commands[4]) && string(m.Content[9]) != "!" && string(m.Content[10]) != "!" {
		commands.NewCommand(m.Content[9:])
		logs.LogErr(logs.ErrFuncUnavailable, "https://github.com/Yuno-obsessed/shikimori/blob/main/internal/messages/messages.go#L56")
		//_, err := s.ChannelMessageSendReply(m.ChannelID, "Your tag was successfully added!", m.Reference())
		//if err != nil {
		//	log.Println(err)
		//}
	}

	// if m.Author.ID == "369852581972803584" {
	// 	reference := m.Reference()
	// 	_, err := s.ChannelMessageSendReply(m.ChannelID, "some cool phrase", reference)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
}

func StartLogging(m *ds.MessageCreate) {
	// randomize output
	rand.Seed(time.Now().Unix())
	// convert string in lowercase
	m.Content = strings.ToLower(m.Content)
	// Logging messages to logfile
	if m.ChannelID == "825200883087573003" {
		msgInfo := m.Message.Author.Username + ": " + m.Message.Content
		logs.Log(msgInfo)
	}
}
