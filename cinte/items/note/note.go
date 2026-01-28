// Package note implements the Note type and methods.
package note

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/cinte/cinte/items/page"
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
	noteDelete     = "delete from Notes where id=?"
	noteInsert     = "insert into Notes (name) values (?) returning *"
	noteRename     = "update Notes set name=? where id=?"
	noteSelectName = "select * from Notes where name=? limit 1"
	noteUpdate     = "insert into Pages (note, body) values (?, ?) returning *"
)

// Create creates and returns a new Note in a database.
func Create(db *sqlx.DB, name, body string) (*Note, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("cannot create note - %w", err)
	}

	note := &Note{DB: db}
	if err := tx.Get(note, noteInsert, name); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("cannot create note - %w", err)
	}

	page := &page.Page{DB: db}
	if err := tx.Get(page, noteUpdate, note.ID, body); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("cannot create note - %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("cannot create note - %w", err)
	}

	return note, nil
}

// Get returns an existing Note from a database, or nil.
func Get(db *sqlx.DB, name string) (*Note, error) {
	note := &Note{DB: db}
	err := db.Get(note, noteSelectName, name)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get note - %w", err)
	default:
		return note, nil
	}
}

// Delete deletes the Note from the database.
func (n *Note) Delete() error {
	if _, err := n.DB.Exec(noteDelete, n.ID); err != nil {
		return fmt.Errorf("cannot delete note - %w", err)
	}

	return nil
}

// Latest returns the Note's latest Page from the database.
func (n *Note) Latest() (*page.Page, error) {
	return page.GetLatest(n.DB, n.ID)
}

// Rename renames the Note in the database.
// Returns an error if name is empty after sanitization.
func (n *Note) Rename(name string) error {
	if _, err := n.DB.Exec(noteRename, name, n.ID); err != nil {
		return fmt.Errorf("cannot rename note - %w", err)
	}

	n.Name = name
	return nil
}

// Time returns the Note's creation time.
func (n *Note) Time() time.Time {
	return neat.Time(n.Init)
}
