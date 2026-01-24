package server

import (
	"fmt"

	"github.com/bibyen/totle-tasks/internal/domain"
	"github.com/bibyen/totle-tasks/internal/pb/totle_tasks/v1/totletasksv1connect"
)

// Server is the main server struct that holds all service handlers.
type Server struct {
	totletasksv1connect.UnimplementedGoalServiceHandler
	totletasksv1connect.UnimplementedBingoServiceHandler

	GoalService  *domain.GoalService
	BingoService *domain.BingoService
}

// NewServer returns a new Server instance.
func NewServer(goalService *domain.GoalService, bingoService *domain.BingoService) (*Server, error) {
	if goalService == nil {
		return nil, fmt.Errorf("goalService cannot be nil")
	}
	if bingoService == nil {
		return nil, fmt.Errorf("bingoService cannot be nil")
	}

	return &Server{
		GoalService:  goalService,
		BingoService: bingoService,
	}, nil
}
