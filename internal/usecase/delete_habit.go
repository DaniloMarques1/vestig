package usecase

import (
	"danilo.marques/vestig/internal/domain"
)

type deleteHabitUseCase struct {
	habitRepository domain.HabitRepository
}

type DeleteHabitUseCaseInputDTO struct {
	HabitID int64
}

type DeleteHabitUseCaseOutputDTO struct {
	HabitID   int64
	HabitName string
}

func NewDeleteHabitUseCase(habitRepository domain.HabitRepository) *deleteHabitUseCase {
	return &deleteHabitUseCase{habitRepository}
}

func (d *deleteHabitUseCase) Execute(input DeleteHabitUseCaseInputDTO) (*DeleteHabitUseCaseOutputDTO, error) {
	habit, err := d.habitRepository.FindById(input.HabitID)
	if err != nil {
		return nil, err
	}

	habit.MarkAsInactive()
	if err := d.habitRepository.Update(habit); err != nil {
		return nil, err
	}

	output := &DeleteHabitUseCaseOutputDTO{HabitID: habit.ID, HabitName: habit.Name}
	return output, nil
}
