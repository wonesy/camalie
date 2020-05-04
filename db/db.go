package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connection controls the database interaction
type Connection struct {
	url string
	db  *sql.DB
}

// NewConnection constructor for Connection
func NewConnection(un, pw, host, db string) (*Connection, error) {
	c := &Connection{
		url: fmt.Sprintf("postgres://%s:%s@%s/%s", un, pw, host, db),
	}
	return c, c.open()
}

func (c *Connection) open() (err error) {
	c.db, err = sql.Open("postgres", c.url)
	return err
}

// Query executes a query on the database
func (c *Connection) Query(q string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(q, args)
}
