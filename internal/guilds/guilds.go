// Guilds package manages guilds (AKA servers) and defines functions
// to use according to guilds part
package guilds

import (
	ds "github.com/bwmarrin/discordgo"
)

// Idk
func checkGuild(session *ds.Session, serverID string) {

}

// Function that returns a struct with server info
// of a server message was sent in
func DefineGuild(session *ds.Session, message *ds.Message) *ds.Guild {
	discordServerID := message.GuildID
	serverStruct, _ := session.Guild(discordServerID)
	return serverStruct
}
