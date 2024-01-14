package routes

import (
	"net/http"
)

type Route struct {
	Name string
	Handler http.HandlerFunc 
}

var routes []Route

func CreateRoute(name string, handler http.HandlerFunc) {
	routes = append(routes, Route{name, handler})
}

func SetupRoutes() {
	setupRoutes()
	setupApiRoutes()
}

func setupRoutes() {
	for _, route := range routes {
		http.HandleFunc(route.Name, func(w http.ResponseWriter, r *http.Request) {
			route.Handler(w, r)
		})
	}
	// serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../node/styles"))))
}

func setupApiRoutes() {
}

func RenderHTMLPage(filename string, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../app/pages/" + filename)
}
