package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)


var testQuires *Queries
var conn *sql.DB

func TestMain(m *testing.M) {
	var err error
	conn, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db",err)
	}
	testQuires = New(conn)
	println(testQuires)

	os.Exit(m.Run())
}
