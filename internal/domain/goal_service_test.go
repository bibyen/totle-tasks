package domain_test

import (
	"context"
	"testing"

	"github.com/bibyen/totle-tasks/internal/domain"
	"github.com/stretchr/testify/require"
)

// DummyGoalRepo is a mock implementation of the GoalRepoProvider interface for testing.
type DummyGoalRepo struct{}

func (r *DummyGoalRepo) Create(ctx context.Context, g *domain.Goal) error {
	// Mock implementation: Assume success
	return nil
}

func (r *DummyGoalRepo) List(ctx context.Context, userID string) ([]*domain.Goal, error) {
	// Mock implementation: Return empty list
	return []*domain.Goal{}, nil
}

func (r *DummyGoalRepo) Update(ctx context.Context, g *domain.Goal) error {
	// Mock implementation: Assume success
	return nil
}

func (r *DummyGoalRepo) Archive(ctx context.Context, id string) error {
	// Mock implementation: Assume success
	return nil
}

func (r *DummyGoalRepo) GetByID(ctx context.Context, id string) (*domain.Goal, error) {
	// Mock implementation: Return a dummy goal
	return &domain.Goal{
		ID:         id,
		UserID:     "user-123",
		Title:      "Dummy Goal",
		Completed:  false,
		Visibility: 1,
		IsActive:   true,
		IsAssigned: false,
	}, nil
}

func TestGoalService_CreateGoal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		newGoal *domain.Goal
		want    *domain.Goal
		wantErr bool
	}{
		// Success cases
		{
			name: "Create valid goal",
			newGoal: &domain.Goal{
				ID:         "goal-123",
				UserID:     "user-456",
				Title:      "Learn Go Testing",
				Completed:  false,
				Visibility: 1,
				IsActive:   true,
				IsAssigned: false,
			},
			want: &domain.Goal{
				ID:         "goal-123",
				UserID:     "user-456",
				Title:      "Learn Go Testing",
				Completed:  false,
				Visibility: 1,
				IsActive:   true,
				IsAssigned: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s domain.GoalService
			dummyRepo := DummyGoalRepo{}
			s.GoalRepoProvider = &dummyRepo

			got, gotErr := s.CreateGoal(context.Background(), tt.newGoal)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateGoal() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateGoal() succeeded unexpectedly")
			}

			require.Equal(t, tt.want, got, "CreateGoal() got = %v, want %v", got, tt.want)
		})
	}
}
