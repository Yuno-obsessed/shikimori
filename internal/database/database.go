// Package database is responsible for database connection and storing data
package database

import (
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
