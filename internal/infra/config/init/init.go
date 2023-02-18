package inits

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/FedorLap2006/disgolf"
	ds "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/yuno-obsessed/shikimori/internal/commands"
	"github.com/yuno-obsessed/shikimori/internal/logs"
	"github.com/yuno-obsessed/shikimori/internal/messages"
)

// Reading bot token from .env
func ReadBotToken() string {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println(err)
	}
	token := os.Getenv("BOT_TOKEN")
	return token
}

type logger logs.Logger

// Function that takes token value and creates a new discord session
func InitializeBot(token string) *disgolf.Bot {
	discordSession, err := disgolf.New(token)
	if err != nil {
		logger.Error("Error initializing a bot", err.Error())
	}
	// Here we add our commands(create a function to wrap all commands in
	// one to be able to easily pass it from commands to init package)
	discordSession.AddHandler(func(session *ds.Session, message *ds.MessageCreate) {
		messages.MessageCreate(session, message)
	})
	commands.InitializeCommands(discordSession)
	discordSession.AddHandler(func(session *ds.Session, r *ds.Ready) {
		logger.Info("Shikimori is up and running", "")
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
	err = discordSession.Router.Sync(discordSession.Session, "1000845128317022249", "931186431215435807")

	if err != nil {
		logger.Error(logs.ErrPublishingCommands, err.Error())
	}
	return discordSession
}

// Function that takes session and starts it
func StartBot(discordSession *disgolf.Bot) {
	err := discordSession.Open()
	if err != nil {
		logger.Error(logs.ErrSessionOpening, err.Error())
	}

	err = discordSession.UpdateGameStatus(0, "Waifuborn")
	defer discordSession.Close()
	if err != nil {
		logger.Error(logs.ErrStatusUpdate, err.Error())
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}

// Function to gather all the functions that start the bot
func Init() {
	botToken := ReadBotToken()
	botSession := InitializeBot(botToken)

	StartBot(botSession)
}
