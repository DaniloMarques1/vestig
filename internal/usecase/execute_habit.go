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

type ExecuteHabitOutputDTO struct {
	HabitName string
}

func (e *ExecuteHabitUseCase) Execute(input ExecuteHabitInputDTO) (*ExecuteHabitOutputDTO, error) {
	habit, err := e.habitRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}

	habitLog := domain.NewHabitLog(habit.ID, input.ExecutedAt)
	if err := e.habitLogRepository.Save(habitLog); err != nil {
		return nil, err
	}

	output := &ExecuteHabitOutputDTO{HabitName: habit.Name}
	return output, nil
}
