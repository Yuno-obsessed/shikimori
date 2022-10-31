package commands

import (
	"github.com/FedorLap2006/disgolf"
)

func InitializeCommands(session *disgolf.Bot) {
	AvatarCommand(session)
	AdviceCommand(session)
	InfoCommand(session)
	WeatherCommand(session)
}
