package server

import (
	"context"

	"github.com/bibyen/totle-tasks/internal/domain" // Your domain models
)

// GoalProvider manages business logic using pure domain models.
// The Server implementation is responsible for converting Proto -> Domain.
type GoalProvider interface {
	// CreateGoal handles the logic of saving a new goal.
	CreateGoal(ctx context.Context, goal *domain.Goal) (*domain.Goal, error)
	// GetGoal retrieves a domain goal by its internal ID.
	GetGoal(ctx context.Context, id string) (*domain.Goal, error)
	// ListGoals returns a slice of domain goals and pagination info.
	ListGoals(ctx context.Context, parentID string, limit int, offset int) ([]*domain.Goal, error)
	// UpdateGoal applies updates to a domain model.
	UpdateGoal(ctx context.Context, id string, updates map[string]any) (*domain.Goal, error)
	// DeleteGoal archives a goal.
	DeleteGoal(ctx context.Context, id string) error
}

// BingoProvider manages the bingo grid logic using domain models.
// The Server implementation is responsible for converting Proto -> Domain.
type BingoProvider interface {
	// CreateBingoCard creates a new bingo card for a specific period.
	CreateBingoCard(ctx context.Context, card *domain.BingoCard) (*domain.BingoCard, error)
	// GetBingoCard retrieves a bingo card by its internal ID.
	GetBingoCard(ctx context.Context, id string) (*domain.BingoCard, error)
}
