package dbutils

import "database/sql"

func GetQueryRows[T any](db *sql.DB, query string, handler func(rowData *T, rows *sql.Rows)) ([]T, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rowsData []T
	for rows.Next() {
		var rowData T
		handler(&rowData, rows)
		rowsData = append(rowsData, rowData)
	}
	return rowsData, nil
}
