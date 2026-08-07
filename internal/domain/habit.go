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
	Done      bool
}

var ErrEmptyHabitName = errors.New("Habit name cannot be empty")

func NewHabit(name string) (*Habit, error) {
	trimmedName := strings.TrimSpace(name)
	if name == "" {
		return nil, ErrEmptyHabitName
	}

	return &Habit{
		Name:      trimmedName,
		Done:      false,
		CreatedAt: time.Now(),
	}, nil
}
