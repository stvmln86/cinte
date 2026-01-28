package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockDB(t *testing.T) {
	// success
	db := MockDB()
	assert.NotNil(t, db)
}
