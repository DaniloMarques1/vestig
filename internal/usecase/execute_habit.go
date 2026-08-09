package usecase

import (
	"time"

	"danilo.marques/vestig/internal/domain"
)

type ExecuteHabitUseCase struct {
	habitRepository    domain.HabitRepository
	habitLogRepository domain.HabitLogRepository
}

func NewExecuteHabitUseCase(habitRepository domain.HabitRepository, habitLogRepository domain.HabitLogRepository) *ExecuteHabitUseCase {
	return &ExecuteHabitUseCase{habitRepository, habitLogRepository}
}

type ExecuteHabitInputDTO struct {
	ID         int64 // TODO: this is the habit id maybe we should change its name
	ExecutedAt time.Time
}

func (e *ExecuteHabitUseCase) Execute(input ExecuteHabitInputDTO) error {
	habit, err := e.habitRepository.FindById(input.ID)
	if err != nil {
		return err
	}

	habitLog := domain.NewHabitLog(habit.ID, input.ExecutedAt)
	if err := e.habitLogRepository.Save(habitLog); err != nil {
		return err
	}

	return nil
}
