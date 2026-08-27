package usecase

import (
	"errors"
	"testing"
	"time"

	"danilo.marques/vestig/internal/domain"
)

func TestViewHabitUseCase_Execute(t *testing.T) {
	t.Run("should view a habit and its executions successfully", func(t *testing.T) {
		now := time.Now()
		nowUTC := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, time.UTC)

		habitRepo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				if id != 1 {
					t.Errorf("expected ID 1, got %d", id)
				}
				return &domain.Habit{
					ID:        1,
					Name:      "Beber Água",
					IsActive:  true,
					CreatedAt: nowUTC,
				}, nil
			},
		}

		logRepo := &mockHabitLogRepository{
			findFn: func(habitID int64) ([]domain.HabitLog, error) {
				if habitID != 1 {
					t.Errorf("expected habitID 1, got %d", habitID)
				}
				return []domain.HabitLog{
					{ID: 10, HabitID: 1, ExecutedAt: nowUTC},
				}, nil
			},
		}

		uc := NewViewHabitUseCase(habitRepo, logRepo)
		input := &ViewHabitUseCaseInputDTO{HabitID: 1}

		output, err := uc.Execute(input)

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}

		if !habitRepo.called {
			t.Error("expected habit repository to have been called")
		}

		if !logRepo.called {
			t.Error("expected habit log repository to have been called")
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

		if output.Streak != 1 {
			t.Errorf("expected streak 1, got %d", output.Streak)
		}

		if len(output.Executions) != 1 {
			t.Fatalf("expected 1 execution, got %d", len(output.Executions))
		}

		if !output.Executions[0].Equal(nowUTC) {
			t.Errorf("expected execution %v, got %v", nowUTC, output.Executions[0])
		}
	})

	t.Run("should return habit with no executions and zero streak when history is empty", func(t *testing.T) {
		habitRepo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				return &domain.Habit{
					ID:       1,
					Name:     "Beber Água",
					IsActive: true,
				}, nil
			},
		}

		logRepo := &mockHabitLogRepository{
			findFn: func(habitID int64) ([]domain.HabitLog, error) {
				return []domain.HabitLog{}, nil
			},
		}

		uc := NewViewHabitUseCase(habitRepo, logRepo)
		input := &ViewHabitUseCaseInputDTO{HabitID: 1}

		output, err := uc.Execute(input)

		if err != nil {
			t.Fatalf("expected nil error, got: %v", err)
		}

		if output == nil {
			t.Fatal("output should not be nil")
		}

		if output.Streak != 0 {
			t.Errorf("expected streak 0, got %d", output.Streak)
		}

		if len(output.Executions) != 0 {
			t.Errorf("expected 0 executions, got %d", len(output.Executions))
		}
	})

	t.Run("should return an error when habit is not found in repository", func(t *testing.T) {
		habitRepo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				return nil, domain.ErrHabitNotFound
			},
		}

		logRepo := &mockHabitLogRepository{}

		uc := NewViewHabitUseCase(habitRepo, logRepo)
		input := &ViewHabitUseCaseInputDTO{HabitID: 99}

		output, err := uc.Execute(input)

		if !errors.Is(err, domain.ErrHabitNotFound) {
			t.Errorf("expected error ErrHabitNotFound, got: %v", err)
		}

		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}

		if !habitRepo.called {
			t.Error("expected habit repository to have been called")
		}

		if logRepo.called {
			t.Error("habit log repository should not have been called when habit is not found")
		}
	})

	t.Run("should return an error when repository fails to fetch logs", func(t *testing.T) {
		errDb := errors.New("database connection failure")

		habitRepo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				return &domain.Habit{
					ID:       1,
					Name:     "Ler Livro",
					IsActive: true,
				}, nil
			},
		}

		logRepo := &mockHabitLogRepository{
			findFn: func(habitID int64) ([]domain.HabitLog, error) {
				return nil, errDb
			},
		}

		uc := NewViewHabitUseCase(habitRepo, logRepo)
		input := &ViewHabitUseCaseInputDTO{HabitID: 1}

		output, err := uc.Execute(input)

		if !errors.Is(err, errDb) {
			t.Errorf("expected error '%v', got: '%v'", errDb, err)
		}

		if output != nil {
			t.Errorf("expected nil output, got: %+v", output)
		}

		if !habitRepo.called {
			t.Error("expected habit repository to have been called")
		}

		if !logRepo.called {
			t.Error("expected habit log repository to have been called")
		}
	})
}
