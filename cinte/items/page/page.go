// Package page implements the Page type and methods.
package page

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/cinte/cinte/tools/neat"
)

// Page is a single version of a note in a database.
type Page struct {
	DB   *sqlx.DB `db:"-"`
	ID   int64    `db:"id"`
	Init int64    `db:"init"`
	Note int64    `db:"note"`
	Body string   `db:"body"`
}

const (
	delete       = "delete from Pages where id=?"
	insert       = "insert into Pages (note, body) values (?, ?) returning *"
	selectLatest = "select * from Pages where note=? order by id desc limit 1"
	selectID     = "select * from Pages where id=? limit 1"
)

// Create creates and returns a new Page in a database.
func Create(db *sqlx.DB, note int64, body string) (*Page, error) {
	page := &Page{DB: db}
	if err := db.Get(page, insert, note, body); err != nil {
		return nil, fmt.Errorf("cannot create page - %w", err)
	}

	return page, nil
}

// Get returns an existing Page from a database.
func Get(db *sqlx.DB, id int64) (*Page, error) {
	page := &Page{DB: db}
	err := db.Get(page, selectID, id)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get page - %w", err)
	default:
		return page, nil
	}
}

// GetLatest returns a Note's latest Page from a database.
func GetLatest(db *sqlx.DB, note int64) (*Page, error) {
	page := &Page{DB: db}
	err := db.Get(page, selectLatest, note)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get page - %w", err)
	default:
		return page, nil
	}
}

// Delete deletes the Page from the database.
func (p *Page) Delete() error {
	if _, err := p.DB.Exec(delete, p.ID); err != nil {
		return fmt.Errorf("cannot delete page - %w", err)
	}

	return nil
}

// Time returns the Page's creation time.
func (p *Page) Time() time.Time {
	return neat.Time(p.Init)
}
