package server

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/bibyen/totle-tasks/internal/domain"
	totletasksv1 "github.com/bibyen/totle-tasks/internal/pb/totle_tasks/v1"
	"github.com/bibyen/totle-tasks/internal/pb/totle_tasks/v1/totletasksv1connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Server is the main server struct that holds all service handlers.
type Server struct {
	totletasksv1connect.UnimplementedGoalServiceHandler
	totletasksv1connect.UnimplementedBingoServiceHandler
	goalService  GoalProvider
	bingoService BingoProvider
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
		goalService:  goalService,
		bingoService: bingoService,
	}, nil
}

// CreateGoal handles the CreateGoal RPC.
func (s Server) CreateGoal(ctx context.Context, req *connect.Request[totletasksv1.CreateGoalRequest]) (*connect.Response[totletasksv1.CreateGoalResponse], error) {
	switch {
		case req == nil:
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("request cannot be nil"))
		case req.Msg == nil:
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("request message cannot be nil"))
		case req.Msg.Goal == nil:
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("goal cannot be nil"))
	}

	// Transform request to domain Goal
	goal := domain.Goal{
		ID:         "placeholder-id", // Generate a new ID for the goal
		UserID:     "placeholder-user-id", // TODO: Extract from auth context
		Title:      req.Msg.Goal.Title,
		Completed:  req.Msg.Goal.Completed,
		Visibility: int(req.Msg.Goal.Visibility),
		IsActive:   true,
		IsAssigned: false,
		CreateTime: timestamppb.Now().AsTime(),
		UpdateTime: timestamppb.Now().AsTime(),
	}

	createdGoal, err := s.goalService.CreateGoal(ctx, &goal)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create goal: %w", err))
	}

	return connect.NewResponse(&totletasksv1.CreateGoalResponse{
		Goal: &totletasksv1.Goal{
		Name:       createdGoal.ID,
		Title:      createdGoal.Title,
		Completed:  createdGoal.Completed,
		Visibility: totletasksv1.Goal_Visibility(createdGoal.Visibility),
		CreateTime: timestamppb.New(createdGoal.CreateTime),
		UpdateTime: timestamppb.New(createdGoal.UpdateTime),
		},
	}), nil
}
