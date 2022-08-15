package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token    string
	tags     = []string{"!panda", "!anime", "!banwords", "!rules", "!links", "!pinned", "!tags", "ответ неврита", "ответ токсика"}
	banwords = []string{"трап", "рфаб", "ванилла", "ваниль", "хуй", "далбаеб", "хуйня"}
	images   = []string{"akame-shocked.gif", "akame-sword.gif", "cringe.png", "moe.gif", "nisekoi-chitoge.gif", "nisekoi-smug.gif"}
	links    = []string{"1)Полезная инфа: https://discord.com/channels/825185921359413278/825197106460753941/1005026874977693748", "2)Таблица: https://docs.google.com/spreadsheets/d/1XsKJBINxQxzXa2TtUoSLqt1Kp0-03Sz2tZ65PlJY94M/edit#gid=1846372233", "3)Видос о сборке: https://youtu.be/g-dUqkDT6wQ"}
)

const ImageURL = "https://raw.githubusercontent.com/Yuno-obsessed/shikimori/main/images/"

func init() {
	file, err := os.ReadFile("token.txt")
	Token = string(file)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Shikimori is ready for her job.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

type Shiki struct {
	Name string `json: "name"`
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.ChannelID == "827608911020163102" {

		if m.Author.ID == s.State.User.ID {
			return
		}
		err := s.UpdateGameStatus(0, "Your waifu")
		if err != nil {
			fmt.Println(err)
		}
		rand.Seed(time.Now().Unix())
		var n int
		for n = 0; n < len(banwords); n++ {
			if strings.Contains(m.Content, banwords[n]) {
				_, err := s.ChannelMessageSend(m.ChannelID, ImageURL+images[rand.Intn(len(images))])
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		if strings.Contains(m.Content, "как ") || strings.Contains(m.Content, "Как ") || strings.Contains(m.Content, "Каким образом") {
			_, err := s.ChannelMessageSend(m.ChannelID, ImageURL+"how-to.jpg")
			if err != nil {
				fmt.Println(err)
			}
		} else if m.Content == tags[0] {
			_, err := s.ChannelMessageSend(m.ChannelID, "https://tenor.com/view/gfg-gif-22720654")
			if err != nil {
				fmt.Println(err)
			}
		} else if m.Content == tags[1] {
			_, err = s.ChannelMessageSend(m.ChannelID, "Иди нахуй дрочер!")
			if err != nil {
				fmt.Println(err)
			}
			//have to group anime links and randomize their output!
		} else if m.Content == tags[2] {
			var banwordies string
			for n = 0; n < len(banwords); n++ {
				banwordies += banwords[n] + ", "
			}
			_, err := s.ChannelMessageSend(m.ChannelID, banwordies)
			if err != nil {
				fmt.Println(err)
			}
		} else if m.Content == tags[4] {
			_, err := s.ChannelMessageSend(m.ChannelID, links[0]+"\n"+links[1])
			if err != nil {
				fmt.Println(err)
			}
		} else if m.Content == tags[5] {
			_, err := s.ChannelMessagesPinned(m.ChannelID)
			if err != nil {
				fmt.Println(err)
			}
		} else if m.Content == tags[6] {
			var taggies string
			for n = 0; n < len(tags); n++ {
				taggies += tags[n] + ", "
			}
			_, err := s.ChannelMessageSend(m.ChannelID, taggies)
			if err != nil {
				fmt.Println(err)
			}
		} else if m.Content == tags[7] {
			_, err := s.ChannelMessageSend(m.ChannelID, ImageURL+"ответ_неврита.png")
			if err != nil {
				fmt.Println(err)
			}
		} else if m.Content == tags[8] {
			_, err := s.ChannelMessageSend(m.ChannelID, ImageURL+"ответ_токсика.jpg")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
