package posts

import (
	"database/sql"
	"log"
	"net/http"
	tposts "project/app/templates/components/posts"
	"project/app/utils/dbtypes"
	"project/app/utils/dbutils"
	"project/internal/router"
	"time"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func getPosts(db *sql.DB) ([]dbtypes.Post, error) {
	query := "SELECT id, title, body FROM posts"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	posts, err := dbutils.GetQueryRows[dbtypes.Post](rows, func(rowData *dbtypes.Post) error {
		err := rows.Scan(&rowData.Id, &rowData.Title, &rowData.Body)
		rowData.CreatedAt = time.Now()
		return err
	})
	if err != nil {
		return nil, err
	}
	log.Println("Success getting posts with query", query)
	return posts, nil
}

func PostPosts(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		_, err := db.Exec("INSERT INTO posts (title, body) VALUES (?, ?)", r.FormValue("name"), r.FormValue("body"))
		if err != nil {
			http.Error(w, "Error inserting post", http.StatusInternalServerError)
			log.Println("Error inserting post", err)
			return templ.NopComponent
		}

		posts, err := getPosts(db)
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			log.Println("Error getting posts", err)
			return templ.NopComponent
		}
		return tposts.Posts(posts)
	}
}

func GetPosts(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		posts, err := getPosts(db)
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			log.Println("Error getting posts", err)
			return templ.NopComponent
		}
		return tposts.Posts(posts)
	}
}

func DeletePost(db *sql.DB) router.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		id := mux.Vars(r)["id"]

		_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Error deleting post", http.StatusInternalServerError)
			log.Println("Error deleting post", err)
		}

		posts, err := getPosts(db)
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			log.Println("Error getting posts", err)
			return templ.NopComponent
		}
		return tposts.Posts(posts)
	}
}
