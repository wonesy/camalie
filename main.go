package main

import (
	"net/http"

	"github.com/wonesy/camalie/chat"
	pbchat "github.com/wonesy/camalie/proto/chat"
)

func main() {
	server := chat.NewServer()
	twpHandler := pbchat.NewChatServer(server, nil)

	http.ListenAndServe(":8080", twpHandler)
}
