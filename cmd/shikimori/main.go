package main

import (
	botInit "Shikimori/internal/init"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
	tags  = []string{"!panda", "!anime", "!ответ неврита", "!ch1h", "!felya", "!newrite", "!warealdok", "bzik", "!auf", "!links", "!пепега", "!как играть?", "!ответ чиха", "!pinned", "!tags"}
	links = []string{"1)Полезная инфа: <https://discord.com/channels/825185921359413278/825197106460753941/1005026874977693748>", "2)Таблица: <https://docs.google.com/spreadsheets/d/1XsKJBINxQxzXa2TtUoSLqt1Kp0-03Sz2tZ65PlJY94M/edit#gid=1846372233>", "3)Видос о сборке: <https://youtu.be/g-dUqkDT6wQ>"}
	//warealdok = []string{"warealdok_1.png", "warealdok_2.png", "warealdok_3.png", "warealdok_4.png", "warealdok_5.png", "warealdok_6.png", "warealdok_7.png", "warealdok_8.png", "warealdok_9.png", "warealdok_10.png", "warealdok_11.png"}
	bz1k    = []string{"bz1k_1.png", "bz1k_2.png"}
	ch1h    = []string{"ch1h_1.png", "ch1h_2.png", "ch1h_3.jpg"}
	felya   = []string{"felya_1.jpg", "felya_2.jpg", "felya_3.png"}
	newrite = []string{"newrite_1.png", "newrite_2.png", "newrite_3.png", "newrite_4.png", "newrite_5.png"}
)

const ImageURL = "https://raw.githubusercontent.com/Yuno-obsessed/shikimori/main/images/"

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + botInit.ReadBotToken())
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

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//if m.ChannelID == "827608911020163102" {

	if m.Author.ID == s.State.User.ID {
		return
	}

	err := s.UpdateGameStatus(0, "Waifuborn")
	if err != nil {
		fmt.Println(err)
	}

	rand.Seed(time.Now().Unix())
	m.Content = strings.ToLower(m.Content)

	if strings.Contains(m.Content, "как играть за ") {
		_, err := s.ChannelMessageSend(m.ChannelID, ImageURL+"gameplay.png")
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
	} else if m.Content == tags[2] {
		_, err = s.ChannelMessageSend(m.ChannelID, ImageURL+"/newrite/"+"newrite_6.png")
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[3] {
		_, err = s.ChannelMessageSend(m.ChannelID, ImageURL+"/ch1h/"+ch1h[rand.Intn(len(ch1h))])
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[4] {
		_, err = s.ChannelMessageSend(m.ChannelID, ImageURL+"/felya/"+felya[rand.Intn(len(felya))])
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[5] {
		_, err = s.ChannelMessageSend(m.ChannelID, ImageURL+"/newrite/"+newrite[rand.Intn(len(newrite))])
		if err != nil {
			fmt.Println(err)
		}

	} else if m.Content == tags[6] {
		_, err = s.ChannelMessageSend(m.ChannelID, "***Я милая няша, я няша-стесняша***")
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[7] {
		_, err = s.ChannelMessageSend(m.ChannelID, ImageURL+bz1k[rand.Intn(len(bz1k))])
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[8] {
		_, err = s.ChannelMessageSend(m.ChannelID, ImageURL+"/auf.png")
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[9] {
		_, err := s.ChannelMessageSend(m.ChannelID, links[0]+"\n"+links[1]+"\n"+links[2])
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[10] {
		_, err = s.ChannelMessageSend(m.ChannelID, "https://docs.google.com/spreadsheets/d/1XsKJBINxQxzXa2TtUoSLqt1Kp0-03Sz2tZ65PlJY94M/edit#gid=299200314")
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[11] {
		_, err = s.ChannelMessageSend(m.ChannelID, "```бандиты -> изгои -> фалмеры -> драугры -> вампиры```")
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[12] {
		_, err = s.ChannelMessageSend(m.ChannelID, ImageURL+"/ch1h/"+"ch1h_2.png")
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[13] {
		_, err = s.ChannelMessageSend(m.ChannelID, ImageURL+"pinned.png")
		if err != nil {
			fmt.Println(err)
		}
	} else if m.Content == tags[14] {
		var taggies string
		for n := 0; n < len(tags); n++ {
			taggies += strconv.Itoa(n+1) + ") " + tags[n] + "\n "
		}
		_, err := s.ChannelMessageSend(m.ChannelID, "``` "+taggies+"```")
		if err != nil {
			fmt.Println(err)
		}
	}
	//}
}
