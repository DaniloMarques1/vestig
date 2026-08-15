package domain

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrHabitNotFound  = errors.New("Habit not found")
	ErrEmptyHabitName = errors.New("Habit name cannot be empty")
)

type Habit struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	IsActive  bool
}

func NewHabit(name string) (*Habit, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return nil, ErrEmptyHabitName
	}

	return &Habit{
		Name:      trimmedName,
		IsActive:  true,
		CreatedAt: time.Now(),
	}, nil
}

func (h *Habit) MarkAsInactive() {
	h.IsActive = false
}
