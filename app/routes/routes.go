package routes

import (
	"database/sql"
	"net/http"
	"project/app/api/posts"
	"project/app/templates/pages/index"
	"project/app/templates/pages/postpage"
	"project/internal/router"

	"github.com/a-h/templ"
)

func InitRoutes(db *sql.DB) {
	// page routes
	router.CreateRoute("/", index.Index())
	router.CreateDynamicRoute("/posts/{id}", func(vars router.Vars) templ.Component {
		id := vars["id"]
		return postpage.Posts(id)
	})

	// api routes
	router.CreateApiRoute("/api/posts", http.MethodGet, posts.GetPosts(db))
	router.CreateApiRoute("/api/posts", http.MethodPost, posts.PostPosts(db))

	router.CreateApiRoute("/api/posts/{id}", http.MethodGet, posts.GetPost(db))
	router.CreateApiRoute("/api/posts/{id}", http.MethodDelete, posts.DeletePost(db))
	router.CreateApiRoute("/api/posts/{id}", http.MethodPatch, posts.ToggleEditPost(db))
	router.CreateApiRoute("/api/posts/{id}", http.MethodPut, posts.UpdatePost(db))
}
