package service

import (
	"context"

	domain "github.com/bibyen/totle-tasks/internal/domain"
)

// Called by the GoalServiceHandler in internal/server/server.go
type GoalService struct{}

// CreateGoal creates a new Goal for the authenticated user.
func (s *GoalService) CreateGoal(ctx context.Context, newGoal domain.Goal) (*domain.Goal, error) {
	return &domain.Goal{}, nil
}

// GetGoal retrieves a specific Goal by its unique resource name.
// Checks the caller has permission to access the specified Goal.
func (s *GoalService) GetGoal(ctx context.Context, goalID string) (*domain.Goal, error) {
	return &domain.Goal{}, nil
}

// ListGoals returns a paginated list of Goals belonging to the parent resource.
func (s *GoalService) ListGoals(ctx context.Context, parentID string) ([]*domain.Goal, error) {
	return []*domain.Goal{}, nil
}

// UpdateGoal updates specific fields of an existing Goal using a FieldMask.
func (s *GoalService) UpdateGoal(ctx context.Context, goalID string, goal domain.Goal, update map[string]any) (*domain.Goal, error) {
	return &domain.Goal{}, nil
}

// DeleteGoal archives a Goal from the system.
func (s *GoalService) DeleteGoal(ctx context.Context, goalID string) error {
	return nil
}
