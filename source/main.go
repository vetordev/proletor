package main

import (
	"db"
	"server"
)

func main() {
	db.Connect()
	server.Serve()
}
