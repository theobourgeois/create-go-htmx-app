package routes

import (
	"log"
	"net/http"
	"os"

	"project/app/templates/layout"

	"github.com/a-h/templ"
)

type Route struct {
	Name string
	Component templ.Component	
}

var routes []Route

func CreateRoute(name string, component templ.Component) {
	routes = append(routes, Route{name, component})
}

func SetupRoutes() {
	// render each route, wrapped in the layout. 
	for _, route := range routes {
		http.Handle(route.Name, templ.Handler(layout.Layout(route.Component)))
	}

	// serve static files
	fullPath := "node/styles"
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatalln("Static file does not exist", err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(fullPath))))
}

func RenderHTMLPage(filename string, w http.ResponseWriter, r *http.Request) {
	fullPath := "../../app/pages/" + filename
  // check if file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatalln("File does not exist", err)
	}
	http.ServeFile(w, r, fullPath) 
}