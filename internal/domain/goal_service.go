package domain

import (
	"context"
)

// Contained in the server in internal/server/server.go
type GoalService struct{}

// CreateGoal creates a new Goal for the authenticated user.
func (s *GoalService) CreateGoal(ctx context.Context, newGoal *Goal) (*Goal, error) {
	return &Goal{}, nil
}

// GetGoal retrieves a specific Goal by its unique resource name.
// Checks the caller has permission to access the specified Goal.
func (s *GoalService) GetGoal(ctx context.Context, goalID string) (*Goal, error) {
	return &Goal{}, nil
}

// ListGoals returns a paginated list of Goals belonging to the parent resource.
func (s *GoalService) ListGoals(ctx context.Context, parentID string) ([]*Goal, error) {
	return []*Goal{}, nil
}

// UpdateGoal updates specific fields of an existing Goal using a FieldMask.
func (s *GoalService) UpdateGoal(ctx context.Context, goalID string, goal *Goal, update map[string]any) (*Goal, error) {
	return &Goal{}, nil
}

// DeleteGoal archives a Goal from the system.
func (s *GoalService) DeleteGoal(ctx context.Context, goalID string) error {
	return nil
}
