package page

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/cinte/cinte/tools/test"
)

func mockPage() *Page {
	page, _ := Get(test.MockDB(), 1)
	return page
}

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB()

	// success
	page, err := Create(db, 1, "Body.\n")
	assert.NotNil(t, page.DB)
	assert.Equal(t, int64(4), page.ID)
	assert.Equal(t, time.Now().Unix(), page.Init)
	assert.Equal(t, int64(1), page.Note)
	assert.Equal(t, "Body.\n", page.Body)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB()

	// success - existing page
	page, err := Get(db, 1)
	assert.NotNil(t, page.DB)
	assert.Equal(t, int64(1), page.ID)
	assert.Equal(t, int64(1767232800), page.Init)
	assert.Equal(t, int64(1), page.Note)
	assert.Equal(t, "Alpha old.\n", page.Body)
	assert.NoError(t, err)

	// success - nonexistent page
	page, err = Get(db, -1)
	assert.Nil(t, page)
	assert.NoError(t, err)
}

func TestGetLatest(t *testing.T) {
	// setup
	db := test.MockDB()

	// success - existing note
	page, err := GetLatest(db, 1)
	assert.NotNil(t, page.DB)
	assert.Equal(t, int64(2), page.ID)
	assert.Equal(t, int64(1767236400), page.Init)
	assert.Equal(t, int64(1), page.Note)
	assert.Equal(t, "Alpha new.\n", page.Body)
	assert.NoError(t, err)

	// success - nonexistent note
	page, err = GetLatest(db, 999)
	assert.Nil(t, page)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	page := mockPage()

	// success
	err := page.Delete()
	assert.NoError(t, err)

	// confirm - page deleted
	var size int
	page.DB.Get(&size, "select count(*) from Pages where id=1")
	assert.Zero(t, size)
}

func TestTime(t *testing.T) {
	// setup
	page := mockPage()
	want := time.Unix(1767232800, 0).Local()

	// success
	tobj := page.Time()
	assert.Equal(t, want, tobj)
}
