package usecase

import (
	"errors"
	"testing"

	"danilo.marques/vestig/internal/domain"
)

func TestDeleteHabitUseCase_Execute(t *testing.T) {
	t.Run("should mark a habit as inactive successfully", func(t *testing.T) {
		repo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				if id != 1 {
					t.Errorf("expected ID 1, got %d", id)
				}
				return &domain.Habit{ID: 1, Name: "Beber Água", IsActive: true}, nil
			},
			updateFn: func(habit *domain.Habit) error {
				if habit.IsActive {
					t.Error("expected habit to be inactive before update")
				}
				return nil
			},
		}

		uc := NewDeleteHabitUseCase(repo)
		input := DeleteHabitUseCaseInputDTO{HabitID: 1}

		output, err := uc.Execute(input)

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}

		if !repo.called {
			t.Error("expected repository to have been called")
		}

		if output == nil {
			t.Fatal("output should not be nil")
		}

		if output.HabitID != 1 {
			t.Errorf("expected HabitID 1, got %d", output.HabitID)
		}

		if output.HabitName != "Beber Água" {
			t.Errorf("expected name 'Beber Água', got '%s'", output.HabitName)
		}
	})

	t.Run("should return an error when habit is not found", func(t *testing.T) {
		repo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				return nil, domain.ErrHabitNotFound
			},
		}

		uc := NewDeleteHabitUseCase(repo)
		input := DeleteHabitUseCaseInputDTO{HabitID: 99}

		output, err := uc.Execute(input)

		if !errors.Is(err, domain.ErrHabitNotFound) {
			t.Errorf("expected error ErrHabitNotFound, got: %v", err)
		}

		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}
	})

	t.Run("should return an error when repository fails to update the habit", func(t *testing.T) {
		errDB := errors.New("database connection failure")
		repo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				return &domain.Habit{ID: 1, Name: "Ler Livro", IsActive: true}, nil
			},
			updateFn: func(habit *domain.Habit) error {
				return errDB
			},
		}

		uc := NewDeleteHabitUseCase(repo)
		input := DeleteHabitUseCaseInputDTO{HabitID: 1}

		output, err := uc.Execute(input)

		if !errors.Is(err, errDB) {
			t.Errorf("expected error '%v', got: '%v'", errDB, err)
		}

		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}

		if !repo.called {
			t.Error("expected repository to have been called")
		}
	})
}
