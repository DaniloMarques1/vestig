package domain

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrHabitNotFound        = errors.New("Habit not found")
	ErrEmptyHabitName       = errors.New("Habit name cannot be empty")
	ErrHabitAlreadyExecuted = errors.New("Habit already executed today")
)

type Habit struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	IsActive  bool
}

func NewHabit(name string) (*Habit, error) {
	trimmedName, err := validateHabitName(name)
	if err != nil {
		return nil, err
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

func (h *Habit) UpdateName(name string) error {
	trimmedName, err := validateHabitName(name)
	if err != nil {
		return err
	}

	h.Name = trimmedName
	return nil
}

func validateHabitName(name string) (string, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return "", ErrEmptyHabitName
	}

	return trimmedName, nil
}
