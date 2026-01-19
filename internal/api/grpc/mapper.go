package grpc

import (
	"fmt"
	"strings"

	"github.com/bibyen/totle-tasks/internal/domain"
	pb "github.com/bibyen/totle-tasks/internal/pb/totle_tasks/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GoalToProto converts a domain.Goal to a pb.Goal message.
func GoalToProto(g *domain.Goal) *pb.Goal {
	if g == nil {
		return nil
	}

	return &pb.Goal{
		Name:      fmt.Sprintf("goals/%s", g.ID),
		Title:     g.Title,
		Completed: g.Completed,
		// Casts domain int (1, 3) to Proto Enum (Private, Public)
		Visibility: pb.Goal_Visibility(g.Visibility),
		IsArchived: !g.IsActive,
		CreateTime: timestamppb.New(g.CreateTime),
		UpdateTime: timestamppb.New(g.UpdateTime),
	}
}

// BingoCardToProto converts a domain.BingoCard to a pb.BingoCard message.
func BingoCardToProto(c *domain.BingoCard) *pb.BingoCard {
	if c == nil {
		return nil
	}

	// Transform flat slots into GridRow structure (Strict 5x5)
	rows := make([]*pb.BingoCard_GridRow, 5)
	for i := range rows {
		rows[i] = &pb.BingoCard_GridRow{
			Slots: make([]*pb.BingoCard_Slot, 5),
		}
	}

	for _, slot := range c.Slots {
		// Ensure coordinate is within 5x5 bounds
		if slot.RowIndex >= 0 && slot.RowIndex < 5 && slot.ColumnIndex >= 0 && slot.ColumnIndex < 5 {
			rows[slot.RowIndex].Slots[slot.ColumnIndex] = &pb.BingoCard_Slot{
				Goal:      fmt.Sprintf("goals/%s", slot.GoalID),
				GoalValue: GoalToProto(slot.Goal),
			}
		}
	}

	return &pb.BingoCard{
		Name:       fmt.Sprintf("bingoCards/%s", c.ID),
		Grid:       rows,
		Year:       int32(c.Year),
		Month:      int32(c.Month),
		CreateTime: timestamppb.New(c.CreateTime),
		UpdateTime: timestamppb.New(c.UpdateTime),
	}
}

// GetUUIDFromResourceName extracts the UUID from "prefix/{uuid}"
func GetUUIDFromResourceName(name string) string {
	parts := strings.Split(name, "/")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-1]
}
