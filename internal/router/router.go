package router

import (
	"log"
	"net/http"

	"project/app/templates/layout"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

type Route struct {
	Name      string
	Component templ.Component
}

type ApiRouteHandler = func(w http.ResponseWriter, r *http.Request) templ.Component

type ApiRoute struct {
	Name    string
	Handler ApiRouteHandler
	Method  string
}

var routes []Route
var apiRoutes []ApiRoute

func CreateRoute(name string, component templ.Component) {
	routes = append(routes, Route{name, component})
}

func CreateApiRoute(name string, method string, handler ApiRouteHandler) {
	apiRoutes = append(apiRoutes, ApiRoute{name, handler, method})
}

func SetupRoutes() {
	r := mux.NewRouter()

	for _, route := range routes {
		r.Handle(route.Name, templ.Handler(layout.Layout(route.Component)))
	}

	for _, apiRoute := range apiRoutes {
		r.HandleFunc(apiRoute.Name, makeHandler(apiRoute)).Methods(apiRoute.Method)
	}

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)
}

func makeHandler(apiRoute ApiRoute) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != apiRoute.Method {
			log.Println("Invalid request method, expected", apiRoute.Method, "got", r.Method, "for", apiRoute.Name, "route")
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		component := apiRoute.Handler(w, r)
		templ.Handler(component).ServeHTTP(w, r)
	}
}
