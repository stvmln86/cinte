// Package test implements unit testing data and functions.
package test

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stvmln86/cinte/cinte/tools/sqls"
)

// MockSchema is additional mock database schema for unit testing.
const MockSchema = `
	insert into Notes (init, name) values
		(1767232800, 'alpha'), -- 2026-01-01 12:00
		(1767319200, 'bravo'); -- 2026-01-02 12:00

	insert into Pages (init, note, body) values
		(1767232800, 1, 'Alpha old.' || char(10)), -- 2026-01-01 12:00
		(1767236400, 1, 'Alpha new.' || char(10)), -- 2026-01-01 13:00
		(1767319200, 2, 'Bravo.'     || char(10)); -- 2026-01-02 12:00
`

// MockDB returns a new in-memory database initialised with MockSchema.
func MockDB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(sqls.Pragma + sqls.Schema + MockSchema)
	return db
}
