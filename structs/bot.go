package structs

import "github.com/raymondbernard/gowebsocket"

type bot struct {
	bot_name         string
	webex_token      string
	approved_rooms   []string
	approved_domains []string
}

// add config object
func Start() {
	ip := "localhost"
	port := "3999"
	server := gowebsocket.New(ip, port)
	server.Start()
}
