package cmd

import (
	"net/http"
	"os"
	"strings"

	"github.com/wonesy/camalie/chat"
	"github.com/wonesy/camalie/db"
	pbchat "github.com/wonesy/camalie/proto/chat"
)

const serviceName = "chat"

// Start kicks off the service
func Start() {
	info := os.Getenv(strings.ToUpper(serviceName) + "_POSTGRES_INFO")
	parts := strings.Split(info, ",")
	if len(parts) != 3 {
		panic("could not get database info for service: chat")
	}

	un, pw, dn := parts[1], parts[2], parts[0]

	conn, err := db.NewConnection(un, pw, "0.0.0.0", dn)
	if err != nil {
		panic("shit")
	}

	server := chat.NewServer(conn)
	twpHandler := pbchat.NewServiceServer(server, nil)

	http.ListenAndServe(":8080", twpHandler)
}
