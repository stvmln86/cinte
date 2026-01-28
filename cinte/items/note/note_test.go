package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/cinte/cinte/tools/test"
)

func mockNote() *Note {
	note, _ := Get(test.MockDB(), "alpha")
	return note
}

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB()

	// success
	note, err := Create(db, "name")
	assert.NotNil(t, note.DB)
	assert.Equal(t, int64(3), note.ID)
	assert.Equal(t, time.Now().Unix(), note.Init)
	assert.Equal(t, "name", note.Name)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB()

	// success - existing note
	note, err := Get(db, "alpha")
	assert.NotNil(t, note.DB)
	assert.Equal(t, int64(1), note.ID)
	assert.Equal(t, int64(1767232800), note.Init)
	assert.Equal(t, "alpha", note.Name)
	assert.NoError(t, err)

	// success - nonexistent note
	note, err = Get(db, "nope")
	assert.Nil(t, note)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	note := mockNote()

	// success
	err := note.Delete()
	assert.NoError(t, err)

	// confirm - note deleted
	var size bool
	note.DB.Get(&size, "select count(*) from Notes where id=1")
	assert.Zero(t, size)

	// confirm - pages deleted
	note.DB.Get(&size, "select count(*) from Pages where note=1")
	assert.Zero(t, size)
}

func TestRename(t *testing.T) {
	// setup
	note := mockNote()

	// success
	err := note.Rename("name")
	assert.Equal(t, "name", note.Name)
	assert.NoError(t, err)

	// confirm - note updated
	var name string
	note.DB.Get(&name, "select name from Notes where id=1")
	assert.Equal(t, "name", name)
}

func TestTime(t *testing.T) {
	// setup
	note := mockNote()
	want := time.Unix(1767232800, 0).Local()

	// success
	tobj := note.Time()
	assert.Equal(t, want, tobj)
}
