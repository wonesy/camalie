package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
	c, err := NewConnection("un", "pw", "host", "db")
	assert.Equal(t, "postgres://un:pw@host/db", c.url)
	assert.NoError(t, err)
}
