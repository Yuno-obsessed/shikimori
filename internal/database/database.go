// Package database is responsible for database connection and storing data
package database

import (
	"Shikimori/internal/guilds"
	"fmt"
	"log"

	ds "github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server ds.Guild

func MongoConnect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("<ATLAS_URI_HERE>"))
	// Implement a full connection + func later when is needed
	if err != nil {
		log.Println(err)
	}
	fmt.Println(client)
}

// TODO: Make a function to output guild struct values
// when bot is running at the server(to copy the output)

// Function that checks if the server is in a database
func CheckServerIs(session *ds.Session, message *ds.Message) bool {
	server := guilds.DefineGuild(session, message)
	// Add a new function to check if this server exists in database
	serverIs := true
	if server == nil {
		serverIs = false
		AddNewServer(server)
	}
	return serverIs
}

// Function that takes structure with server info
// and passes it into the server info in database
func AddNewServer(server *ds.Guild) Server {
	return Server{
		ID:      "825185921359413278",
		Name:    "Gallina Blanca Crusaders",
		Region:  "Russia",
		OwnerID: "438476530302189579",
	}
}
