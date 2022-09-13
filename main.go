package main

import (
	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
	"github.com/raymondbernard/gowebsocket"
)

func main() {
	var client = webexteams.NewClient()

	ip := "localhost"
	port := "3999"
	server := gowebsocket.New(ip, port)
	server.Start()
}
