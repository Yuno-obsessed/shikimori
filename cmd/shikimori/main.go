package main

import (
	inits "github.com/yuno-obsessed/shikimori/internal/init"
)

func main() {

	token := inits.ReadBotToken()
	session := inits.InitializeBot(token)
	inits.StartBot(session)

}
