//go:build integration

package postgres

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/bibyen/totle-tasks/internal/domain"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// setupTestDB handles the connection and provides a cleanup function.
// It relies on the database already being created and schema-loaded by the Makefile.
func setupTestDB(t *testing.T) (*sql.DB, func()) {
	t.Helper()

	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		// Default fallback for local IDE testing
		dsn = "postgres://postgres:postgres@localhost:5432/totle_tasks_test?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}

	// Verify the physical connection exists
	if err := db.Ping(); err != nil {
		t.Fatalf("Could not ping database at %s. Did you run 'make db/init-test'? Error: %v", dsn, err)
	}

	// Teardown: Wipes data between tests but keeps the tables
	teardown := func() {
		// CASCADE handles foreign key dependencies (e.g. slots -> goals)
		_, err := db.Exec("TRUNCATE goals, bingo_cards, bingo_slots, users RESTART IDENTITY CASCADE")
		if err != nil {
			t.Logf("Cleanup warning: %v", err)
		}
		db.Close()
	}

	return db, teardown
}

func TestGoalRepo_Integration(t *testing.T) {
	db, teardown := setupTestDB(t)
	defer teardown()

	repo := NewGoalRepo(db)
	ctx := context.Background()

	t.Run("Create and Retrieve", func(t *testing.T) {
		// Use valid UUID strings
		gID := "550e8400-e29b-41d4-a716-446655440000"

		g := &domain.Goal{
			ID:         gID,
			UserID:     "user-1",
			Title:      "Finish the MVP",
			Completed:  false,
			Visibility: 1,
			IsActive:   true,
			IsAssigned: false,
		}

		err := repo.Create(ctx, g)
		assert.NoError(t, err)

		// ... rest of test
	})

	t.Run("Archive (Soft Delete)", func(t *testing.T) {
		// MUST be a valid UUID format
		gID := "72161491-7210-449a-9e19-07f7c4627407"

		_ = repo.Create(ctx, &domain.Goal{
			ID: gID, UserID: "user-1", Title: "To be archived", IsActive: true,
		})

		err := repo.Archive(ctx, gID)
		assert.NoError(t, err)

		fetched, err := repo.GetByID(ctx, gID)
		assert.NoError(t, err)
		assert.False(t, fetched.IsActive, "Goal should be logically deleted (is_active = false)")
	})

	t.Run("List Active Goals", func(t *testing.T) {
		// Different UUIDs for different goals
		id1 := "11111111-1111-1111-1111-111111111111"
		id2 := "22222222-2222-2222-2222-222222222222"

		_ = repo.Create(ctx, &domain.Goal{ID: id1, UserID: "user-list", Title: "Active", IsActive: true})
		_ = repo.Create(ctx, &domain.Goal{ID: id2, UserID: "user-list", Title: "Archived", IsActive: false})

		goals, err := repo.List(ctx, "user-list")
		assert.NoError(t, err)
		assert.Len(t, goals, 1, "List should only return goals where is_active = true")
		assert.Equal(t, id1, goals[0].ID)
	})
}
