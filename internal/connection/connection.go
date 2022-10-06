// Connection package is responsible for configuring bot connection
package connection

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type BotConnection discordgo.Session

func ConnectionSettings() *BotConnection {
	settings := &BotConnection{
		ShouldReconnectOnError: true,
		StateEnabled:           true,
		SyncEvents:             false,
		MaxRestRetries:         5,
		Client:                 *&http.DefaultClient,
	}
	return settings
}

