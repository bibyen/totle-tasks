package domain

import (
	"context"
)

// GoalRepoProvider defines the interface for goal repository operations.
type GoalRepoProvider interface {
	// Create creates a new goal.
	Create(ctx context.Context, g *Goal) error
	// List lists all active goals for a user.
	List(ctx context.Context, userID string) ([]*Goal, error)
	// Update updates an existing goal.
	// Returns an error if the goal does not exist or is not active.
	Update(ctx context.Context, g *Goal) error
	// Archive "soft-deletes" a goal by marking it as inactive.
	// Returns an error if the goal does not exist or is already inactive.
	Archive(ctx context.Context, id string) error
	// GetByID fetches a goal by ID.
	GetByID(ctx context.Context, id string) (*Goal, error)
}
