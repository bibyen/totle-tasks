package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	// 1. Setup: Connect to a local test database
	// Use an environment variable for the DSN so CI/CD can override it
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/totle_tasks_test?sslmode=disable"
	}

	var err error
	testDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not connect to test database: %v", err)
	}

	// 2. Run the tests
	code := m.Run()

	// 3. Teardown
	err = testDB.Close()
	if err != nil {
		log.Printf("Could not close test database: %v", err)
	}

	// 4. Exit with the proper code
	os.Exit(code)
}
