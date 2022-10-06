// Init package is responsible for initializing and making bot accessible
package init

import (
	"embed"
	"log"

	"github.com/bwmarrin/discordgo"
)

//go:embed token.txt
var tokenTxt embed.FS

// This function will return the bot's token from token.txt file as a string
func ReadBotToken() string {
	tokenInBytes, err := tokenTxt.ReadFile("token.txt")
	if err != nil {
		log.Println(err)
	}
	token := string(tokenInBytes)
	return token
}

func InitializeBot(token string) *discordgo.Session {
	dg, err := discordgo.New("BOT" + token)
	if err != nil {
		log.Println(err)
	}
	return dg
}

func StartBot(discordSession *discordgo.Session) {
	err := discordSession.Open()
	if err != nil {
		log.Println(err)
	}
}
