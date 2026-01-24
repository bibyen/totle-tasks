package domain

import (
	"context"
)

// Contained in the server in internal/server/server.go
type BingoService struct{}

// CreateBingoCard explicitly creates a new bingo card for a specific period.
func (s *BingoService) CreateBingoCard(ctx context.Context, newCard *BingoCard) (*BingoCard, error) {
	return &BingoCard{}, nil
}

// GetBingoCard retrieves the bingo card for a specific year and month.
func (s *BingoService) GetBingoCard(ctx context.Context, year int32, month int32) (*BingoCard, error) {
	return &BingoCard{}, nil
}

// UpdateBingoCard updates the layout or goal assignments within a bingo card.
func (s *BingoService) UpdateBingoCard(ctx context.Context, cardID string, card *BingoCard, update map[string]any) (*BingoCard, error) {
	return &BingoCard{}, nil
}
