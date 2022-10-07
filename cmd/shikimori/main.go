package main

import (
	inits "Shikimori/internal/init"
	"Shikimori/internal/logs"
	"fmt"
	"math/rand"
	"strings"
	"time"

	ds "github.com/bwmarrin/discordgo"
)

func main() {

	// // Create a new Discord session using the provided bot token.
	// dg, err := discordgo.New("Bot " + botInit.ReadBotToken())
	// if err != nil {
	// 	fmt.Println("error creating Discord session,", err)
	// 	return
	// }

	// // Register the messageCreate func as a callback for MessageCreate events.
	// dg.AddHandler(messageCreate)

	// // In this example, we only care about receiving message events.
	// dg.Identify.Intents = discordgo.IntentsGuildMessages

	// // Open a websocket connection to Discord and begin listening.
	// err = dg.Open()
	// if err != nil {
	// 	fmt.Println("error opening connection,", err)
	// 	return
	// }
	token := inits.ReadBotToken()
	session := inits.InitializeBot(token)
	inits.StartBot(session)

	// Wait here until CTRL-C or other term signal is received.

	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// <-sc

	// // Cleanly close down the Discord session.
	// session.Close()
}

func messageCreate(s *ds.Session, m *ds.MessageCreate) {

	//if m.ChannelID == "827608911020163102" {
	if m.Author.ID == s.State.User.ID {
		return
	}

	err := s.UpdateGameStatus(0, "Waifuborn")
	if err != nil {
		fmt.Println(err)
	}
	/*content, err := m.ContentWithMoreMentionsReplaced(s)
	if err != nil {
		log.Println(err)
	}
	*/

	if m.ChannelID == "825200883087573003" {
		logs.LogProviding(m.Message.Author.Username + " глаголит истину: " + m.Message.Content)
	}

	// randomize output
	rand.Seed(time.Now().Unix())
	// convert string in lowercase
	m.Content = strings.ToLower(m.Content)

	if m.Author.ID == "369852581972803584" {
		reference := m.Reference()
		_, err := s.ChannelMessageSendReply(m.ChannelID, "Пошёл нахуй, пошёл нахуй, пошёл нахуй, пидарас!", reference)
		if err != nil {
			fmt.Println(err)
		}
	}

}
