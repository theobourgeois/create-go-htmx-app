package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"project/internal/routes"

	"github.com/joho/godotenv"
)

func Start() {	
	routes.SetupRoutes()
	err := godotenv.Load()
	if err != nil {
			log.Fatalln("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	fmt.Println("Server starting on port", port, "...") 
	go func() {
		if err := http.ListenAndServe(":" + port, nil); err != nil {
			log.Fatalln(err)
		}
	}()

	select {}
}

