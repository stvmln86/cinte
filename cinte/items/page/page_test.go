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
	assert.EqualExportedValues(t, &Page{
		DB:   db,
		ID:   int64(4),
		Init: time.Now().Unix(),
		Note: int64(1),
		Body: "Body.\n",
	}, page)
	assert.NoError(t, err)

	// confirm - page created
	var body string
	db.Get(&body, "select body from Pages where id=4")
	assert.Equal(t, "Body.\n", body)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB()

	// success - existing page
	page, err := Get(db, 1)
	assert.EqualExportedValues(t, &Page{
		DB:   db,
		ID:   int64(1),
		Init: int64(1767232800),
		Note: int64(1),
		Body: "Alpha old.\n",
	}, page)
	assert.NoError(t, err)

	// success - nonexistent page
	page, err = Get(db, -1)
	assert.Nil(t, page)
	assert.NoError(t, err)
}

func TestGetLatest(t *testing.T) {
	// setup
	db := test.MockDB()

	// success - existing page
	page, err := GetLatest(db, 1)
	assert.EqualExportedValues(t, &Page{
		DB:   db,
		ID:   int64(2),
		Init: int64(1767236400),
		Note: int64(1),
		Body: "Alpha new.\n",
	}, page)
	assert.NoError(t, err)

	// success - nonexistent page
	page, err = GetLatest(db, -1)
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
