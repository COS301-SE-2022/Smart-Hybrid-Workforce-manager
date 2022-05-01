package db

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

//////////////////////////////////////////////////
// Structures and Variables

// database connection
var database *sql.DB

// Access is a wrapper for a db and its transaction state
type Access struct {
	DataBase    *sql.DB
	Transaction *sql.Tx
}

/////////////////////////////////////////////
// Functions

// RegisterAccess creates a connection pool
func RegisterAccess() error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_DSN"))
	if err != nil {
		return err
	}
	maxIdleConn, err := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		return err
	}
	maxOpenConn, err := strconv.Atoi(os.Getenv("DATABASE_MAX_OPEN_CONNECTIONS"))
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)
	duration, err := time.ParseDuration("30m")
	db.SetConnMaxLifetime(duration)
	err = db.Ping()
	if err != nil {
		return err
	}
	database = db
	return nil
}

// UnregisterAccess disposes of the connection pool
func UnregisterAccess() error {
	err := database.Close()
	if err != nil {
		return err
	}
	return nil
}

// Open creates a new Access structure
func Open() (*Access, error) {
	access := Access{database, nil}
	transaction, err := access.DataBase.Begin()
	if err != nil {
		return nil, err
	}
	access.Transaction = transaction
	return &access, nil
}

// Query the Database
func (access *Access) Query(query string, mapping func(*sql.Rows) (interface{}, error), args ...interface{}) ([]interface{}, error) {
	rows, err := access.Transaction.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if mapping == nil {
		return nil, nil
	}
	result := make([]interface{}, 0)
	for rows.Next() {
		row, err := mapping(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

// Commit transaction and start anew
func (access *Access) Commit() error {
	err := access.Transaction.Commit()
	if err != nil {
		return err
	}
	transaction, err := access.DataBase.Begin()
	if err != nil {
		return err
	}
	access.Transaction = transaction
	return nil
}

// Rollback transaction and start anew
func (access *Access) Rollback() error {
	err := access.Transaction.Rollback()
	if err != nil {
		return err
	}
	transaction, err := access.DataBase.Begin()
	if err != nil {
		return err
	}
	access.Transaction = transaction
	return nil
}

// Close closes the current Access
func (access *Access) Close() error {
	if access.Transaction == nil {
		return nil
	}
	err := access.Transaction.Rollback()
	if err != nil {
		return err
	}
	access.Transaction = nil
	access.DataBase = nil
	return nil
}
