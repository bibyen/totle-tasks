package domain

import (
	"context"
	"errors"
	"fmt"
	"log"

	"connectrpc.com/connect"
	"github.com/bibyen/totle-tasks/internal/auth"
)

// Contained in the server in internal/server/server.go
type GoalService struct {
	GoalRepoProvider GoalRepoProvider
}

// CreateGoal creates a new Goal for the authenticated user.
// Users can only create goals for themself.
func (s *GoalService) CreateGoal(ctx context.Context, newGoal *Goal) (*Goal, error) {
	switch {
	case newGoal == nil:
		return nil, fmt.Errorf("goal cannot be nil")
	case newGoal.ID == "":
		return nil, fmt.Errorf("goal ID cannot be empty")
	case newGoal.UserID == "":
		return nil, fmt.Errorf("user ID cannot be empty")
	case newGoal.Title == "":
		return nil, fmt.Errorf("goal title cannot be empty")
	}
	log.Default().Println("GoalService.CreateGoal called with goal:", newGoal)

	// // Extract user id
	// userID, ok := auth.GetIdentityFromContext(ctx)
	// if !ok {
	// 	return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("identity not found"))
	// }

	// // Ensure user is creating goal for self
	// if userID != newGoal.UserID {
	// 	return nil, connect.NewError(connect.CodePermissionDenied, errors.New("cannot create goal for another user"))
	// }

	// Attempt to save the new goal
	log.Default().Println("Saving new goal for user:", newGoal.UserID)
	err := s.GoalRepoProvider.Create(ctx, newGoal)
	if err != nil {
		// TODO: Consider using a structured logger for better log management. https://github.com/bibyen/totle-tasks/issues/10
		// Track metadata about the request + error
		// e.g., timestamp, userID, goal details (excluding sensitive info), error message
		// log.Printf("Internal error creating goal (%v) for user %s: %v", *newGoal, userID, err)
		return nil, fmt.Errorf(("unable to save goal, please try again later"))
	}
	log.Default().Println("Successfully created goal with ID:", newGoal.ID)

	return newGoal, nil
}

// GetGoal retrieves a specific Goal by its unique resource name.
// Checks the caller has permission to access the specified Goal.
func (s *GoalService) GetGoal(ctx context.Context, goalID string) (*Goal, error) {
	// Extract user id
	userID, ok := auth.GetIdentityFromContext(ctx)
	if !ok {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("identity not found"))
	}

	// Ensure user is fetching goal for self
	// Fetch the goal
	goal, err := s.GoalRepoProvider.GetByID(ctx, goalID)
	if err != nil { // Log the error for internal tracking
		// TODO: Consider using a structured logger for better log management. https://github.com/bibyen/totle-tasks/issues/10
		// Track metadata about the request + error
		// e.g., timestamp, userID, goal details (excluding sensitive info), error message
		log.Default().Printf("Internal error fetching goal with id %s: %v", goalID, err)
		return nil, connect.NewError(connect.CodeInternal, errors.New("unable to get goal, please try again later"))
	}

	if goal.UserID != userID {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("cannot fetch goal belonging to another user"))
	}

	return goal, nil
}

// ListGoals returns a paginated list of Goals belonging to the parent resource.
func (s *GoalService) ListGoals(ctx context.Context, parent string, pageSize int32, pageToken string) ([]*Goal, string, error) {
	// Extract user id
	_, ok := auth.GetIdentityFromContext(ctx)
	if !ok {
		return nil, "", connect.NewError(connect.CodeUnauthenticated, errors.New("identity not found"))
	}
	return []*Goal{}, "", nil
}

// UpdateGoal updates specific fields of an existing Goal using a FieldMask.
func (s *GoalService) UpdateGoal(ctx context.Context, goalID string, update map[string]any) (*Goal, error) {
	// Extract user id
	_, ok := auth.GetIdentityFromContext(ctx)
	if !ok {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("identity not found"))
	}
	return &Goal{}, nil
}

// DeleteGoal archives a Goal from the system.
func (s *GoalService) DeleteGoal(ctx context.Context, goalID string) error {
	// Extract user id
	_, ok := auth.GetIdentityFromContext(ctx)
	if !ok {
		return connect.NewError(connect.CodeUnauthenticated, errors.New("identity not found"))
	}

	return nil
}
