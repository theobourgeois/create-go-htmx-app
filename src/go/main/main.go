package main

import (
	"db"
	"server"
)

func main() {
	db.InitDB()
	server.Start()
}