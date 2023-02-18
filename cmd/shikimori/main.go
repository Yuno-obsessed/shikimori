package main

import (
	"github.com/yuno-obsessed/shikimori/internal/infra/config/init"
)

func main() {
	token := inits.ReadBotToken()
	session := inits.InitializeBot(token)
	inits.StartBot(session)
}
