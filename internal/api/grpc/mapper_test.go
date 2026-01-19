package grpc

import (
	"testing"
	"time"

	"github.com/bibyen/totle-tasks/internal/domain"
	pb "github.com/bibyen/totle-tasks/internal/pb/totle_tasks/v1"
	"github.com/stretchr/testify/assert"
)

func TestGoalToProto(t *testing.T) {
	now := time.Now()
	id := "550e8400-e29b-41d4-a716-446655440000"

	dGoal := &domain.Goal{
		ID:         id,
		Title:      "Test Goal",
		Completed:  true,
		Visibility: domain.GoalVisibilityPublic,
		IsActive:   true,
		CreateTime: now,
		UpdateTime: now,
	}

	pGoal := GoalToProto(dGoal)

	assert.Equal(t, "goals/"+id, pGoal.Name)
	assert.Equal(t, pb.Goal_VISIBILITY_PUBLIC, pGoal.Visibility)
	assert.False(t, pGoal.IsArchived)
}

func TestGetIDFromResourceName(t *testing.T) {
	assert.Equal(t, "123", GetUUIDFromResourceName("goals/123"))
	assert.Equal(t, "abc", GetUUIDFromResourceName("bingoCards/abc"))
	assert.Equal(t, "", GetUUIDFromResourceName("invalid"))
}
