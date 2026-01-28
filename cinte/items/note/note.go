// Package note implements the Note type and methods.
package note

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/cinte/cinte/tools/neat"
)

// Note is a single named note in a database.
type Note struct {
	DB   *sqlx.DB `db:"-"`
	ID   int64    `db:"id"`
	Init int64    `db:"init"`
	Name string   `db:"name"`
}

const (
	delete     = "delete from Notes where id=?"
	insert     = "insert into Notes (name) values (?) returning *"
	selectName = "select * from Notes where name=? limit 1"
	updateName = "update Notes set name=? where id=?"
)

// Create creates and returns a new Note in a database.
func Create(db *sqlx.DB, name string) (*Note, error) {
	note := &Note{DB: db}
	if err := db.Get(note, insert, name); err != nil {
		return nil, fmt.Errorf("cannot create note %q - %w", name, err)
	}

	return note, nil
}

// Get returns an existing Note from a database.
func Get(db *sqlx.DB, name string) (*Note, error) {
	note := &Note{DB: db}
	err := db.Get(note, selectName, name)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get note %q - %w", name, err)
	default:
		return note, nil
	}
}

// Delete deletes the Note from the database.
func (n *Note) Delete() error {
	if _, err := n.DB.Exec(delete, n.ID); err != nil {
		return fmt.Errorf("cannot delete note %q - %w", n.Name, err)
	}

	return nil
}

// Rename renames the Note in the database.
func (n *Note) Rename(name string) error {
	if _, err := n.DB.Exec(updateName, name, n.ID); err != nil {
		return fmt.Errorf("cannot rename note %q - %w", n.Name, err)
	}

	n.Name = name
	return nil
}

// Time returns the Note's creation time.
func (n *Note) Time() time.Time {
	return neat.Time(n.Init)
}
