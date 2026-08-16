package usecase

import (
	"time"

	"danilo.marques/vestig/internal/domain"
	"danilo.marques/vestig/internal/domain/service"
)

type viewHabitUseCase struct {
	habitRepository    domain.HabitRepository
	habitLogRepository domain.HabitLogRepository
}

type ViewHabitUseCaseInputDTO struct {
	HabitID int64
}

type ViewHabitUseCaseOutputDTO struct {
	HabitID    int64
	HabitName  string
	Streak     int64
	Executions []time.Time
}

func NewViewHabitUseCase(habitRepository domain.HabitRepository, habitLogRepository domain.HabitLogRepository) *viewHabitUseCase {
	return &viewHabitUseCase{habitRepository, habitLogRepository}
}

func (v *viewHabitUseCase) Execute(input *ViewHabitUseCaseInputDTO) (*ViewHabitUseCaseOutputDTO, error) {
	habit, err := v.habitRepository.FindById(input.HabitID)
	if err != nil {
		return nil, err
	}
	logs, err := v.habitLogRepository.Find(habit.ID)
	if err != nil {
		return nil, err
	}
	executions := make([]time.Time, 0, len(logs))
	for _, log := range logs {
		executions = append(executions, log.ExecutedAt)
	}
	streak := service.CalculateStreak(executions, time.Now())
	output := &ViewHabitUseCaseOutputDTO{
		HabitID:    habit.ID,
		HabitName:  habit.Name,
		Streak:     streak,
		Executions: executions,
	}

	return output, nil
}
