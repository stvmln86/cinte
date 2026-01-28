package dbse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// setup
	db, err := Connect(":memory:", "create table Mock (a)")
	assert.NotNil(t, db)
	assert.NoError(t, err)

	// confirm - database
	var size int
	err = db.Get(&size, "select count(*) from SQLITE_SCHEMA")
	assert.Equal(t, 1, size)
	assert.NoError(t, err)

	// error - cannot connect
	db, err = Connect("/nope/nope.db", "")
	assert.Nil(t, db)
	assert.ErrorContains(t, err, "cannot connect database")

	// error - cannot initialise
	db, err = Connect(":memory:", "nope")
	assert.Nil(t, db)
	assert.ErrorContains(t, err, "cannot initialise database")
}
