package domain

import (
	"errors"
	"testing"
	"time"
)

func TestNewHabit(t *testing.T) {
	t.Run("should create a valid habit with correct default values", func(t *testing.T) {
		before := time.Now()

		habit, err := NewHabit("Meditar")

		after := time.Now()

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}

		if habit == nil {
			t.Fatal("habit should not be nil")
		}

		if habit.Name != "Meditar" {
			t.Errorf("expected name 'Meditar', got '%s'", habit.Name)
		}

		if habit.IsActive != true {
			t.Errorf("expected IsActive true, got %t", habit.IsActive)
		}

		if habit.ID != 0 {
			t.Errorf("expected initial ID 0, got %d", habit.ID)
		}

		if habit.CreatedAt.Before(before) || habit.CreatedAt.After(after) {
			t.Errorf("CreatedAt %v outside expected range [%v, %v]", habit.CreatedAt, before, after)
		}
	})

	t.Run("should trim leading and trailing whitespaces from name", func(t *testing.T) {
		habit, err := NewHabit("   Correr no parque   ")

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}

		if habit.Name != "Correr no parque" {
			t.Errorf("expected name 'Correr no parque', got '%s'", habit.Name)
		}
	})

	t.Run("should return ErrEmptyHabitName when name is an empty string", func(t *testing.T) {
		habit, err := NewHabit("")

		if !errors.Is(err, ErrEmptyHabitName) {
			t.Errorf("expected error ErrEmptyHabitName, got: %v", err)
		}

		if habit != nil {
			t.Errorf("expected nil habit on creation failure, got: %+v", habit)
		}
	})

	t.Run("should return ErrEmptyHabitName when name contains only whitespaces", func(t *testing.T) {
		habit, err := NewHabit("     ")

		if !errors.Is(err, ErrEmptyHabitName) {
			t.Errorf("expected error ErrEmptyHabitName, got: %v", err)
		}

		if habit != nil {
			t.Errorf("expected nil habit on creation failure, got: %+v", habit)
		}
	})
}
