package db

import (
	"database/sql"
	"fmt"
)

// Connection controls the database interaction
type Connection struct {
	url string
	db  *sql.DB
}

// NewConnection constructor for Connection
func NewConnection(un, pw, host, db string) (*Connection, error) {
	c := &Connection{
		url: fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", un, pw, host, db),
	}
	return c, c.open()
}

func (c *Connection) open() (err error) {
	c.db, err = sql.Open("postgres", c.url)
	return err
}

// Query executes a query on the database
func (c *Connection) Query(q string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(q, args...)
}

// Execute will execute an SQL query that doesn't return any rows
func (c *Connection) Execute(s string, args ...interface{}) (sql.Result, error) {
	return c.db.Exec(s, args...)
}
