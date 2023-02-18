// Package database is responsible for database connection and storing data
package database

import (
	"context"
	"fmt"
	ds "github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yuno-obsessed/shikimori/internal/infra/config/logger"
	"log"
	"os"
)

type Server ds.Guild

type Database struct {
	Pool   *pgxpool.Pool
	Logger logger.Logger
}

func NewDatabase() Database {
	return Database{
		Pool:   DbInit(),
		Logger: logger.NewLogger(),
	}
}

func DbDns() string {
	return fmt.Sprintf("%s://%s:%s@localhost:%s/%s&parseTime=True",
		os.Getenv("POSTGRES_DRIVER"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
}

func DbInit() *pgxpool.Pool {
	dns := DbDns()
	pool, err := pgxpool.New(context.Background(), dns)
	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
		return nil
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("database is unaccessible, %v", err)
		return nil
	}

	return pool
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
