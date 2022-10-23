package commands

import (
	"log"
	"math/rand"
	"regexp"
	"time"

	ds "github.com/bwmarrin/discordgo"
	"github.com/yuno-obsessed/shikimori/internal/logs"
)

var Commands []string = []string{
	"!tags",
	"!advice",
	"!avatar",
	"!info",
	"!tagadd",
}

type LoveInfo struct {
	sender   string
	receiver string
	timeout  bool
}

func (l *LoveInfo) CheckIf(is bool) {

}

// To invoke a function:
// if string(str[1]) != "/" && string(str[2]) != "" {
func NewCommand(str string) {

	if CountWords(str) == 2 {
		Commands = append(Commands, str)
	}
	log.Printf("Command %v was successfully added\n", str)
}

func ListTags() string {
	var result string
	for _, val := range Commands {
		result += ("> " + val + "\n")
	}
	return result
}
func LoveIndicator(sender string, receiver string) int {
	timer := time.NewTimer(time.Hour * 24)

	<-timer.C
	love := rand.Intn(101)
	return love
}

// Function that returns avatar's
// URL with size specified
func InsertAvatar(m *ds.MessageCreate) string {
	var imageURL string
	switch CountWords(m.Content) {
	case 1:
		imageURL = "https://cdn.discordapp.com/avatars/" + m.Author.ID + "/" + m.Author.Avatar + ".png?size=1024"

	case 2:
		l := m.Mentions[0]
		imageURL = "https://cdn.discordapp.com/avatars/" + l.ID + "/" + l.Avatar + ".png?size=1024"

	default:
		l := m.Mentions[0]
		imageURL = "https://cdn.discordapp.com/avatars/" + l.ID + "/" + l.Avatar + ".png?size=1024"
		logs.LogErr(logs.ErrTooMuchFlags, "InsertAvatar function")
	}
	return imageURL
}

func CountWords(str string) int {
	re := regexp.MustCompile(`[\S]+`)

	// Find all matches and return count.
	stringQuantity := re.FindAllString(str, -1)
	return len(stringQuantity)
}
