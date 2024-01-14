package main

import (
	"project/app/routes"
	"project/internal/db"
	"project/internal/server"
)

func main() {
	db.InitDB()
	routes.InitRoutes(db.GetDB())
	server.Start()
}