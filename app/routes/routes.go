package routes

import (
	"database/sql"
	"net/http"
	"project/app/api/posts"
	"project/app/templates/pages/index"
	"project/internal/router"
)

func InitRoutes(db *sql.DB) {
	// page routes
	homePage := index.Index()
	router.CreateRoute("/", homePage)

	// api routes
	router.CreateApiRoute("/api/posts", http.MethodGet, posts.GetPosts(db))
	router.CreateApiRoute("/api/posts", http.MethodPost, posts.PostPosts(db))

	router.CreateApiRoute("/api/posts/{id}", http.MethodDelete, posts.DeletePost(db))
}
