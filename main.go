package main

import (
	chatcmd "github.com/wonesy/camalie/chat/cmd"

	_ "github.com/lib/pq"
)

func main() {
	chatcmd.Start()
}
