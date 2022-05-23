package data

import "database/sql"

//////////////////////////////////////////////////
// Mappers

// mapString
func mapString(rows *sql.Rows) (interface{}, error) {
	var s string
	err := rows.Scan(
		&s,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}
