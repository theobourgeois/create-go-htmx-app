package files

import (
	"database/sql"
	"log"
	"net/http"
	"project/app/templates/components/filelist"
	"project/app/utils/dbutils"
	"project/app/utils/types/dbtypes"
	"project/internal/routes"

	"github.com/a-h/templ"
)

func getFiles(db *sql.DB) []dbtypes.File {
	rows, err := db.Query("SELECT id, name, type, url size FROM File")
	if err != nil {
		log.Fatalln("Error with file query: ", err)
	}
	defer rows.Close()

	var files []dbtypes.File
	for rows.Next() {
		var file dbtypes.File
		err := rows.Scan(&file.Id, &file.Name, &file.Type, &file.Url)
		if err != nil {
			log.Fatalln("Error with scanning file query: ", err)
		}
		files = append(files, file)
	}

	return files
}

func FilesHandler(db *sql.DB) routes.ApiRouteHandler {
	return func(w http.ResponseWriter, r *http.Request) templ.Component {
		filesList, err := dbutils.GetQueryRows[dbtypes.File](db, "SELECT id, name, type, url size FROM File", func(file *dbtypes.File, rows *sql.Rows) {
			err := rows.Scan(&file.Id, &file.Name, &file.Type, &file.Url)
			if err != nil {
				log.Fatalln("Error with scanning file query: ", err)
			}
		})

		if err != nil {
			log.Fatalln("Error with file query: ", err)
		}

		return filelist.FileList(filesList)
	}
}
