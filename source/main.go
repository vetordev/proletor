package main

import (
	"proletor/db"
	"proletor/server"
)

func main() {
	db.Connect()
	server.Serve()
}
