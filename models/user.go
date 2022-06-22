package models

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

var userMetadata = table.Metadata{
	Name:    "users",
	Columns: []string{"id", "firstname", "lastname", "age", "email", "city"},
	PartKey: []string{"id"},
	SortKey: nil,
}

var UserTable = table.New(userMetadata)

type User struct {
	ID        gocql.UUID
	Firstname string
	Lastname  string
	Age       int
	Email     string
	City      string
}
