package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getTemplateHtmlFilename(page string) string {
	switch page {
		case "/":
			return "index"
		case "/about":
			return "about" 
		default:
			return "index" 
	}
}

func getTemplateHtmlFromPageName(page string) string {
	filename := getTemplateHtmlFilename(page)
	fullFilename := "../templates/" + filename + ".html"

	body, err := os.ReadFile(fullFilename)

	if err != nil {
		log.Fatalln("Error reading file", err)
	}

	return string(body)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlContent := getTemplateHtmlFromPageName("/") 
		fmt.Fprint(w, htmlContent)
	})

	// serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../templates"))))
}

func setupApiRoutes() {
}

func Start() {	
	setupRoutes()
	setupApiRoutes()

	fmt.Println("Server starting on port 8080...")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
				fmt.Println(err)
		}
	}()

	// This will now print out after the server starts.
	fmt.Println("Server started.")

	// Prevent the program from exiting immediately.
	select {}
}

