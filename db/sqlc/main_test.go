package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
