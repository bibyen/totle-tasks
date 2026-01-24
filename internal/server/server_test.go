package server

// test new server
import (
	"testing"

	"github.com/bibyen/totle-tasks/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServer_Success(t *testing.T) {
	t.Run("server created successfully", func(t *testing.T) {
		goalService := &service.GoalService{}
		bingoService := &service.BingoService{}

		_, gotErr := NewServer(goalService, bingoService)
		require.NoError(t, gotErr)
	})
}

func TestNewServer_Failure(t *testing.T) {
	tests := []struct {
		name         string
		goalService  *service.GoalService
		bingoService *service.BingoService
		errMsg       string
	}{
		{
			name:         "Nil GoalService",
			goalService:  nil,
			bingoService: &service.BingoService{},
			errMsg:       "goalService cannot be nil",
		},
		{
			name:         "Nil BingoService",
			goalService:  &service.GoalService{},
			bingoService: nil,
			errMsg:       "bingoService cannot be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := NewServer(tt.goalService, tt.bingoService)
			if gotErr != nil {
				assert.EqualError(t, gotErr, tt.errMsg)
			}
		})
	}
}
