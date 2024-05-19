package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func Connect() {
	connStr := "postgres://v-chat-user:v-chat-password@localhost:5432/v-chat-database"

	var err error

	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to DataBase: %v/n", err)
	}

	err = Conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping to DataBase: %v/n", err)
	} else {
		log.Println("Succesfuly connected to DataBase")
	}
}

func Close() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}