// Init package is responsible for initializing and making bot accessible
package inits

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/FedorLap2006/disgolf"
	ds "github.com/bwmarrin/discordgo"
	"github.com/yuno-obsessed/shikimori/internal/commands"
	"github.com/yuno-obsessed/shikimori/internal/logs"
	"github.com/yuno-obsessed/shikimori/internal/messages"
)

//go:embed token.txt
var tokenTxt embed.FS

// Function that returns bot's token from token.txt file as a string
func ReadBotToken() string {
	tokenInBytes, err := tokenTxt.ReadFile("token.txt")
	if err != nil {
		log.Println(err)
	}
	token := string(tokenInBytes)
	return token
}

// Function that takes token value and creates a new discord session
func InitializeBot(token string) *disgolf.Bot {
	discordSession, err := disgolf.New(token)
	if err != nil {
		logs.LogErr(logs.ErrSessionOpening, "InitializeBot")
	}
	// Here we add our commands(create a function to wrap all commands in
	// one to be able to easily pass it from commands to init package)
	discordSession.AddHandler(func(session *ds.Session, message *ds.MessageCreate) {
		messages.MessageCreate(session, message)
	})
	commands.InitializeCommands(discordSession)
	discordSession.AddHandler(func(session *ds.Session, r *ds.Ready) {
		log.Println("Shikimori is ready for her job.")
	})
	discordSession.AddHandler(discordSession.Router.HandleInteraction)
	discordSession.AddHandler(discordSession.Router.MakeMessageHandler(&disgolf.MessageHandlerConfig{
		Prefixes:      []string{"d.", "dis.", "disgolf."},
		MentionPrefix: true,
	}))
	discordSession.StateEnabled = true
	// discordSession.Identify.Intents = ds.IntentsGuildMessages
	// discordSession.Identify.Intents |= ds.IntentGuildMembers
	// discordSession.Identify.Intents |= ds.IntentMessageContent
	// discordSession.Identify.Intents |= ds.IntentGuildPresences
	// discordSession.Identify.Intents |= ds.IntentGuildIntegrations
	discordSession.Identify.Intents = ds.IntentsAll
	err = discordSession.Router.Sync(discordSession.Session, "1000845128317022249", "1000850818406293526")
	err = discordSession.Router.Sync(discordSession.Session, "1000845128317022249", "825185921359413278")

	if err != nil {
		log.Println(err)
	}
	return discordSession
}

// Function that takes session and starts it
func StartBot(discordSession *disgolf.Bot) {
	err := discordSession.Open()
	if err != nil {
		logs.LogErr(logs.ErrSessionOpening, "")
	}

	discordSession.UpdateGameStatus(0, "Waifuborn")
	defer discordSession.Close()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot publish commands: %w", err))
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}

func Init() {
	botToken := ReadBotToken()
	botSession := InitializeBot(botToken)

	StartBot(botSession)

}
