package main

import (
	"net/http"
	"os"

	"github.com/wonesy/camalie/chat"
	"github.com/wonesy/camalie/db"
	pbchat "github.com/wonesy/camalie/proto/chat"
)

func main() {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASS")
	database := os.Getenv("POSTGRES_DB")

	conn, err := db.NewConnection(user, pass, "0.0.0.0", database)
	if err != nil {
		// TODO log
		panic("shit")
	}

	server := chat.NewServer(conn)
	twpHandler := pbchat.NewServiceServer(server, nil)

	http.ListenAndServe(":8080", twpHandler)
}
