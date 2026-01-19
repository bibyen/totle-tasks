package domain

import (
	"context"
)

// GoalService defines the business logic and orchestration for Goals.
type GoalService interface {
	CreateGoal(ctx context.Context, userID string, title string, visibility int) (*Goal, error)
	GetGoal(ctx context.Context, id string) (*Goal, error)
	ListGoals(ctx context.Context, userID string) ([]*Goal, error)
	UpdateGoal(ctx context.Context, goal *Goal) error
	ArchiveGoal(ctx context.Context, id string) error
}

// BingoService defines the logic for managing 5x5 Bingo Cards.
type BingoService interface {
	// CreateBingoCard explicitly initializes a new 5x5 grid for a user.
	// This ensures the year/month requirements are met before a card is born.
	CreateBingoCard(ctx context.Context, userID string, title string, year int, month int) (*BingoCard, error)

	// GetMonthlyCard finds the standard 5x5 card for a specific period.
	GetMonthlyCard(ctx context.Context, userID string, year int, month int) (*BingoCard, error)

	// AssignGoalToSlot places a specific goal on a Bingo Card coordinate.
	AssignGoalToSlot(ctx context.Context, cardID string, goalID string, row int, col int) error

	// ArchiveBingoCard marks a card as inactive.
	ArchiveBingoCard(ctx context.Context, id string) error
}
