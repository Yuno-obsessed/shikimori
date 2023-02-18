package chat

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

type ChatData struct {
	ChatId         string
	MessageChannel chan *discordgo.Message
}

func ReadMessages(session *discordgo.Session, m *discordgo.MessageCreate) {
	chatDataList := []ChatData{
		{"827608911020163102", make(chan *discordgo.Message, 1)},
		{"825200883087573003", make(chan *discordgo.Message, 1)},
	}

	for _, chatData := range chatDataList {
		go readMessagesFromChat(session, chatData.ChatId, chatData.MessageChannel)
	}

	for _, chatData := range chatDataList {
		go func(ch *ChatData) {
			for message := range ch.MessageChannel {
				fmt.Printf("Received message '%s' from channel %s\n", message.Content, ch.ChatId)
			}
		}(&chatData)
	}
}

func readMessagesFromChat(session *discordgo.Session, chatId string, messageChannel chan *discordgo.Message) {
	for {
		ch, err := session.Channel(chatId)
		if err != nil {
			log.Println(err)
			continue
		}
		message, err := session.ChannelMessage(chatId, ch.LastMessageID)
		if err != nil {
			log.Println(err)
			continue
		}
		for {
			select {
			case <-time.After(10 * time.Second):
				message, err = session.ChannelMessage(chatId, ch.LastMessageID)
				if err != nil {
					log.Println(err)
					continue
				}
			default:
				if message.ID != ch.LastMessageID {
					messageChannel <- message
					break
				}
			}
		}
	}
}
