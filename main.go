package main

import (
	"lim/chat"
	"lim/db"
	httpserver "lim/http"
)

func main() {
	db.InitDB()
	webAddr := ":8080"
	go httpserver.StartWebServer(webAddr)
	chatserver.StartChatServer()
}
