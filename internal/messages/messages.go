// Messages package is responsible for messages handling and reacting
package messages

import (
	"math/rand"
	"strings"
	"time"

	ds "github.com/bwmarrin/discordgo"
	"github.com/yuno-obsessed/shikimori/internal/commands"
	"github.com/yuno-obsessed/shikimori/internal/functions"
	"github.com/yuno-obsessed/shikimori/internal/logs"
)

func MessageCreate(s *ds.Session, m *ds.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	//s.ChannelPermissionSet("825200883087573003",s.State.User.ID,1,true,false)
	// err := s.UpdateGameStatus(0, "Waifuborn")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// randomize output
	rand.Seed(time.Now().Unix())
	// convert string in lowercase
	m.Content = strings.ToLower(m.Content)
	// Logging messages to standard output
	if m.ChannelID == "825200883087573003" {
		msgInfo := m.Message.Author.Username + ": " + m.Message.Content
		logs.Log(msgInfo)
	}
	if m.Content == commands.Commands[0] {
		// SelectMenu struct
		err, _ := s.ChannelMessageSend(m.ChannelID, commands.ListTags())
		if err != nil {
			logs.LogErr(logs.ErrUnableToSendMessage, "https://github.com/Yuno-obsessed/shikimori/blob/main/internal/messages/messages.go#L39")
		}
	} else if m.Content == commands.Commands[1] {
		_, err := s.ChannelMessageSend(m.ChannelID, AdviceMessage(m))
		if err != nil {
			logs.LogErr(logs.ErrUnableToSendMessage, "https://github.com/Yuno-obsessed/shikimori/blob/main/internal/messages/messages.go#L44")
		}
	} else if strings.Contains(m.Content, commands.Commands[2]) && commands.CountWords(m.Content) < 3 {
		av := m.Member.Avatar
		_, err := s.ChannelMessageSendReply(m.ChannelID, av, m.Reference())
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

func AdviceMessage(m *ds.MessageCreate) string {
	response := functions.NiceAdvice()
	return response
}

func InsertAvatar(string) {

}
