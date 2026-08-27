package usecase

import (
	"errors"
	"testing"

	"danilo.marques/vestig/internal/domain"
)

func TestAddHabitUseCase_Execute(t *testing.T) {
	t.Run("should create a habit successfully", func(t *testing.T) {
		repo := &mockHabitRepository{
			saveFn: func(h *domain.Habit) error {
				if h.Name != "Beber Água" {
					t.Errorf("expected name 'Beber Água', got '%s'", h.Name)
				}
				// Simulates DB generating ID in struct
				h.ID = 1
				return nil
			},
		}

		uc := NewAddHabitUseCase(repo)
		input := &AddHabitInputDTO{Name: "Beber Água"}

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

		if output.ID != 1 {
			t.Errorf("expected ID 1, got %d", output.ID)
		}

		if output.Name != "Beber Água" {
			t.Errorf("expected name 'Beber Água', got '%s'", output.Name)
		}
	})

	t.Run("should return an error when habit name is empty", func(t *testing.T) {
		repo := &mockHabitRepository{}
		uc := NewAddHabitUseCase(repo)

		input := &AddHabitInputDTO{Name: "   "}

		output, err := uc.Execute(input)

		if !errors.Is(err, domain.ErrEmptyHabitName) {
			t.Errorf("expected error ErrEmptyHabitName, got: %v", err)
		}

		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}

		if repo.called {
			t.Error("repository should not have been called when validation fails")
		}
	})

	t.Run("should return an error when repository fails to save", func(t *testing.T) {
		errDb := errors.New("database connection failure")

		repo := &mockHabitRepository{
			saveFn: func(h *domain.Habit) error {
				return errDb
			},
		}

		uc := NewAddHabitUseCase(repo)
		input := &AddHabitInputDTO{Name: "Ler Livro"}

		output, err := uc.Execute(input)

		if !errors.Is(err, errDb) {
			t.Errorf("expected error '%v', got: '%v'", errDb, err)
		}

		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}

		if !repo.called {
			t.Error("expected repository to have been called")
		}
	})
}
