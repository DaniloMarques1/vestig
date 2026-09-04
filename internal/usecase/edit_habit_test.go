package usecase

import (
	"errors"
	"testing"

	"danilo.marques/vestig/internal/domain"
)

func TestEditHabitUseCase_Execute(t *testing.T) {
	t.Run("should update a habit name successfully", func(t *testing.T) {
		repo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				if id != 1 {
					t.Errorf("expected ID 1, got %d", id)
				}
				return &domain.Habit{ID: 1, Name: "Beber Água", IsActive: true}, nil
			},
			updateFn: func(habit *domain.Habit) error {
				if habit.Name != "Ler Livro" {
					t.Errorf("expected updated name 'Ler Livro', got '%s'", habit.Name)
				}
				return nil
			},
		}

		output, err := NewEditHabitUseCase(repo).Execute(&EditHabitInputDTO{HabitID: 1, Name: "Ler Livro"})

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}
		if output == nil {
			t.Fatal("output should not be nil")
		}
		if output.HabitID != 1 || output.OldName != "Beber Água" || output.NewName != "Ler Livro" {
			t.Errorf("unexpected output: %+v", output)
		}
		if !repo.called {
			t.Error("expected repository to have been called")
		}
	})

	t.Run("should return an error when the habit is not found", func(t *testing.T) {
		repo := &mockHabitRepository{findByIdFn: func(int64) (*domain.Habit, error) {
			return nil, domain.ErrHabitNotFound
		}}

		output, err := NewEditHabitUseCase(repo).Execute(&EditHabitInputDTO{HabitID: 99, Name: "Ler Livro"})

		if !errors.Is(err, domain.ErrHabitNotFound) {
			t.Errorf("expected ErrHabitNotFound, got: %v", err)
		}
		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}
	})

	t.Run("should reject an empty habit name without updating", func(t *testing.T) {
		repo := &mockHabitRepository{
			findByIdFn: func(int64) (*domain.Habit, error) {
				return &domain.Habit{ID: 1, Name: "Beber Água"}, nil
			},
			updateFn: func(*domain.Habit) error {
				t.Error("repository should not update a habit with an empty name")
				return nil
			},
		}

		output, err := NewEditHabitUseCase(repo).Execute(&EditHabitInputDTO{HabitID: 1, Name: "   "})

		if !errors.Is(err, domain.ErrEmptyHabitName) {
			t.Errorf("expected ErrEmptyHabitName, got: %v", err)
		}
		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}
	})

	t.Run("should return an error when updating the habit fails", func(t *testing.T) {
		errDB := errors.New("database connection failure")
		repo := &mockHabitRepository{
			findByIdFn: func(int64) (*domain.Habit, error) {
				return &domain.Habit{ID: 1, Name: "Beber Água"}, nil
			},
			updateFn: func(*domain.Habit) error { return errDB },
		}

		output, err := NewEditHabitUseCase(repo).Execute(&EditHabitInputDTO{HabitID: 1, Name: "Ler Livro"})

		if !errors.Is(err, errDB) {
			t.Errorf("expected error %v, got: %v", errDB, err)
		}
		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}
	})
}
