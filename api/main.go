package main

import (
	"topicpad-api/db"
	"topicpad-api/server"
)

func main() {
  db.Init()

	defer db.Close()
	server.Init()
}
