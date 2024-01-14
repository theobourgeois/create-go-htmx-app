package routes

import (
	"database/sql"
	"net/http"
	"project/internal/routes"
)

func InitRoutes(db *sql.DB) {
	routes.CreateRoute("/", func(w http.ResponseWriter, r *http.Request) {
				// do something
	})

}
