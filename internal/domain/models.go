package domain

import "time"

// GoalVisibility levels
const (
	GoalVisibilityPrivate = 1
	GoalVisibilityPublic  = 3
)

type Goal struct {
	ID         string
	UserID     string
	Title      string
	Completed  bool
	Visibility int
	IsActive   bool
	IsAssigned bool
	CreateTime time.Time
	UpdateTime time.Time
}

type BingoCard struct {
	ID         string
	UserID     string
	Title      string
	Columns    int
	Rows       int
	Year       int
	Month      int
	IsActive   bool
	Slots      []Slot
	CreateTime time.Time // Added
	UpdateTime time.Time
}

type Slot struct {
	CardID      string
	GoalID      string
	RowIndex    int
	ColumnIndex int
	Goal        *Goal
}
