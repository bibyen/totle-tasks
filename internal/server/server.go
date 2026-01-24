package server

import (
	totletasksv1 "github.com/bibyen/totle-tasks/internal/pb/totle_tasks/v1"
)

// Server implements the gRPC server for GoalService and BingoService.
type Server struct {
	totletasksv1.UnimplementedGoalServiceServer
	totletasksv1.UnimplementedBingoServiceServer
}
