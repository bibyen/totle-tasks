package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bibyen/totle-tasks/internal/domain"
)

type GoalRepo struct {
	db *sql.DB
}

func NewGoalRepo(db *sql.DB) *GoalRepo {
	return &GoalRepo{db: db}
}

func (r *GoalRepo) Create(ctx context.Context, g *domain.Goal) error {
	query := `
		INSERT INTO goals (goal_id, user_id, title, completed, visibility, is_active, is_assigned)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING create_time, update_time`

	return r.db.QueryRowContext(ctx, query,
		g.ID, g.UserID, g.Title, g.Completed, g.Visibility, g.IsActive, g.IsAssigned,
	).Scan(&g.CreateTime, &g.UpdateTime)
}

func (r *GoalRepo) List(ctx context.Context, userID string) ([]*domain.Goal, error) {
	// Standard List: All active goals for the user
	query := `SELECT goal_id, user_id, title, completed, visibility, is_active, is_assigned, create_time, update_time 
	          FROM goals WHERE user_id = $1 AND is_active = true ORDER BY create_time DESC`
	return r.queryList(ctx, query, userID)
}

// Internal helper to avoid code duplication
func (r *GoalRepo) queryList(ctx context.Context, query string, args ...interface{}) ([]*domain.Goal, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []*domain.Goal
	for rows.Next() {
		var g domain.Goal
		if err := rows.Scan(&g.ID, &g.UserID, &g.Title, &g.Completed, &g.Visibility, &g.IsActive, &g.IsAssigned, &g.CreateTime, &g.UpdateTime); err != nil {
			return nil, err
		}
		goals = append(goals, &g)
	}
	return goals, nil
}

// GetByID retrieves a single goal.
func (r *GoalRepo) GetByID(ctx context.Context, id string) (*domain.Goal, error) {
	query := `
		SELECT goal_id, user_id, title, completed, visibility, is_active, create_time, update_time
		FROM goals
		WHERE goal_id = $1`

	var g domain.Goal
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&g.ID, &g.UserID, &g.Title, &g.Completed, &g.Visibility, &g.IsActive, &g.CreateTime, &g.UpdateTime,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Return nil if not found
		}
		return nil, fmt.Errorf("failed to get goal: %w", err)
	}
	return &g, nil
}

// Update modifies an existing goal if it is still active.
func (r *GoalRepo) Update(ctx context.Context, g *domain.Goal) error {
	query := `
		UPDATE goals
		SET title = $1,
		    completed = $2,
		    visibility = $3,
		    update_time = CURRENT_TIMESTAMP
		WHERE goal_id = $4 AND is_active = true
		RETURNING update_time`

	err := r.db.QueryRowContext(ctx, query,
		g.Title, g.Completed, g.Visibility, g.ID,
	).Scan(&g.UpdateTime)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("goal not found or archived")
		}
		return fmt.Errorf("failed to update goal: %w", err)
	}
	return nil
}

// Archive performs a "Soft Delete" by setting is_active to false.
func (r *GoalRepo) Archive(ctx context.Context, id string) error {
	query := `
		UPDATE goals 
		SET is_active = false, update_time = CURRENT_TIMESTAMP 
		WHERE goal_id = $1 AND is_active = true`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to archive goal: %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("goal not found or already archived")
	}
	return nil
}
