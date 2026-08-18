package usecase

import (
	"errors"
	"testing"
	"time"

	"danilo.marques/vestig/internal/domain"
)

func TestListHabitUseCase(t *testing.T) {
	t.Run("should request only active habits when ShowAll is false and map DTO correctly", func(t *testing.T) {
		now := time.Now()
		mock := &mockHabitRepository{
			listFn: func(isActive bool) ([]domain.Habit, error) {
				return []domain.Habit{
					{ID: 1, Name: "Meditate", IsActive: true, CreatedAt: now},
					{ID: 2, Name: "Study Go", IsActive: true, CreatedAt: now},
				}, nil
			},
		}

		useCase := NewListHabitUseCase(mock)
		output, err := useCase.Execute(ListHabitInputDTO{ShowAll: false})

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}

		if !mock.called {
			t.Error("expected repository to have been called")
		}

		// Rule inversion: ShowAll=false requires isActive=true on repository call
		if mock.lastIsActiveArg != true {
			t.Errorf("expected repo.List(isActive=true), got: %v", mock.lastIsActiveArg)
		}

		if len(output.Habits) != 2 {
			t.Fatalf("expected 2 habits in DTO, got %d", len(output.Habits))
		}

		// Validate DTO field mapping accuracy
		if output.Habits[0].ID != 1 || output.Habits[0].Name != "Meditate" || !output.Habits[0].IsActive {
			t.Errorf("incorrect mapping for the first habit DTO: %+v", output.Habits[0])
		}
	})

	t.Run("should request all habits (active and inactive) when ShowAll is true", func(t *testing.T) {
		mock := &mockHabitRepository{
			listFn: func(isActive bool) ([]domain.Habit, error) {
				return []domain.Habit{
					{ID: 1, Name: "Meditate", IsActive: true},
					{ID: 2, Name: "Smoke", IsActive: false},
				}, nil
			},
		}

		useCase := NewListHabitUseCase(mock)
		output, err := useCase.Execute(ListHabitInputDTO{ShowAll: true})

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}

		// Rule inversion: ShowAll=true requires isActive=false (no filter) on repository call
		if mock.lastIsActiveArg != false {
			t.Errorf("expected repo.List(isActive=false), got: %v", mock.lastIsActiveArg)
		}

		if len(output.Habits) != 2 {
			t.Fatalf("expected 2 habits in DTO, got %d", len(output.Habits))
		}

		if output.Habits[1].IsActive != false {
			t.Errorf("expected second habit in DTO to be inactive")
		}
	})

	t.Run("should return DTO with empty list without errors when no habits are found", func(t *testing.T) {
		mock := &mockHabitRepository{
			listFn: func(isActive bool) ([]domain.Habit, error) {
				return []domain.Habit{}, nil
			},
		}

		useCase := NewListHabitUseCase(mock)
		output, err := useCase.Execute(ListHabitInputDTO{ShowAll: false})

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}

		if output == nil {
			t.Fatal("output DTO should not be nil")
		}

		if len(output.Habits) != 0 {
			t.Errorf("expected 0 habits in DTO, got %d", len(output.Habits))
		}
	})

	t.Run("should propagate error and return nil DTO when repository fails", func(t *testing.T) {
		expectedErr := errors.New("database query failure")
		mock := &mockHabitRepository{
			listFn: func(isActive bool) ([]domain.Habit, error) {
				return nil, expectedErr
			},
		}

		useCase := NewListHabitUseCase(mock)
		output, err := useCase.Execute(ListHabitInputDTO{ShowAll: false})

		if !errors.Is(err, expectedErr) {
			t.Errorf("expected error '%v', got: '%v'", expectedErr, err)
		}

		if output != nil {
			t.Errorf("expected nil DTO on error, got: %+v", output)
		}
	})
}
