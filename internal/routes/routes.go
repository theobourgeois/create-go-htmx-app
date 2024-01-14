package routes

import (
	"log"
	"net/http"
	"os"

	"project/app/templates/layout"

	"github.com/a-h/templ"
)

type Route struct {
	Name      string
	Component templ.Component
}

type ApiRouteHandler func(w http.ResponseWriter, r *http.Request) templ.Component
type ApiRoute struct {
	Name    string
	Handler ApiRouteHandler
}

var routes []Route
var apiRoutes []ApiRoute

func CreateRoute(name string, component templ.Component) {
	routes = append(routes, Route{name, component})
}

func CreateApiRoute(name string, handler ApiRouteHandler) {
	apiRoutes = append(apiRoutes, ApiRoute{name, handler})
}

func SetupRoutes() {
	for _, route := range routes {
		http.Handle(route.Name, templ.Handler(layout.Layout(route.Component)))
	}

	for _, apiRoute := range apiRoutes {
		http.HandleFunc(apiRoute.Name, func(w http.ResponseWriter, r *http.Request) {
			component := apiRoute.Handler(w, r)
			templ.Handler(component).ServeHTTP(w, r)
		})
	}

	serveStaticFiles()
}

func serveStaticFiles() {
	// Tailwind files
	tailwindCssDir := "styles"
	if _, err := os.Stat(tailwindCssDir); os.IsNotExist(err) {
		log.Fatalln("Static file does not exist", err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(tailwindCssDir))))
}
