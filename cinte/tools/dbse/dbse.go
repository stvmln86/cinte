// Package dbse implements low-level database handling functions.
package dbse

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Connect returns a new SQLite database connection with an executed query.
func Connect(path, code string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("cannot connect database %q - %w", path, err)
	}

	if _, err := db.Exec(code); err != nil {
		return nil, fmt.Errorf("cannot initialise database %q - %w", path, err)
	}

	return db, nil
}
