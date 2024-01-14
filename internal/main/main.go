package main

import (
	"project/app/routes"
	"project/internal/db"
	"project/internal/server"
)

func main() {
	routes.InitRoutes(db.GetDB())
	db.InitDB()
	server.Start()
}