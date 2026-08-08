package domain

import (
	"errors"
	"strings"
	"time"
)

type Habit struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	IsActive  bool
}

var ErrEmptyHabitName = errors.New("Habit name cannot be empty")

func NewHabit(name string) (*Habit, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, ErrEmptyHabitName
	}

	return &Habit{
		Name:      trimmedName,
		IsActive:  false,
		CreatedAt: time.Now(),
	}, nil
}
