package usecase

import (
	"errors"
	"testing"
	"time"

	"danilo.marques/vestig/internal/domain"
)

func TestExecuteHabitUseCase_Execute(t *testing.T) {
	executedAt := time.Date(2026, 8, 26, 10, 30, 0, 0, time.UTC)

	t.Run("does not save when the habit was already executed", func(t *testing.T) {
		saved := false
		habitRepo := &mockHabitRepository{findByIdFn: func(id int64) (*domain.Habit, error) {
			return &domain.Habit{ID: id, Name: "Beber Água"}, nil
		}}
		logRepo := &mockHabitLogRepository{
			findExecutionAtFn: func(habitId int64, dt time.Time) (bool, error) {
				if !dt.Equal(executedAt) {
					t.Errorf("expected execution date %v, got %v", executedAt, dt)
				}
				return true, nil
			},
			saveFn: func(*domain.HabitLog) error {
				saved = true
				return nil
			},
		}

		output, err := NewExecuteHabitUseCase(habitRepo, logRepo).Execute(ExecuteHabitInputDTO{
			ID: 1, ExecutedAt: executedAt,
		})

		if !errors.Is(err, domain.ErrHabitAlreadyExecuted) {
			t.Fatalf("expected error ErrHabitAlreadyExecuted, got %v", err)
		}
		if output != nil {
			t.Fatalf("expected nil output, got %+v", output)
		}
		if saved {
			t.Error("habit log should not have been saved")
		}
	})

	t.Run("returns an error when checking execution fails", func(t *testing.T) {
		errCheck := errors.New("database connection failure")
		habitRepo := &mockHabitRepository{findByIdFn: func(int64) (*domain.Habit, error) {
			return &domain.Habit{ID: 1, Name: "Ler Livro"}, nil
		}}
		logRepo := &mockHabitLogRepository{findExecutionAtFn: func(habitId int64, dt time.Time) (bool, error) {
			return false, errCheck
		}}

		output, err := NewExecuteHabitUseCase(habitRepo, logRepo).Execute(ExecuteHabitInputDTO{
			ID: 1, ExecutedAt: executedAt,
		})

		if !errors.Is(err, errCheck) {
			t.Fatalf("expected error %v, got %v", errCheck, err)
		}
		if output != nil {
			t.Fatalf("expected nil output, got %+v", output)
		}
	})
}
