package routes

import (
	"database/sql"
	"project/app/api/files"
	"project/app/templates/pages/index"
	"project/internal/routes"
)

func InitRoutes(db *sql.DB) {
	homePage := index.Index("World")
	routes.CreateRoute("/", homePage)
	routes.CreateApiRoute("/api/test", files.FilesHandler(db))
}
