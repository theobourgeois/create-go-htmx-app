package routes

import (
	"database/sql"
	"project/app/templates/index"
	"project/internal/routes"
)

func InitRoutes(db *sql.DB) {
	homePage := index.Index("World")
	routes.CreateRoute("/", homePage)
}
