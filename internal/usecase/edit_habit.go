package usecase

import (
	"strings"

	"danilo.marques/vestig/internal/domain"
)

type EditHabitUseCase struct {
	repository domain.HabitRepository
}

type EditHabitInputDTO struct {
	HabitID int64
	Name    string
}

type EditHabitOutputDTO struct {
	HabitID int64
	OldName string
	NewName string
}

func NewEditHabitUseCase(repository domain.HabitRepository) *EditHabitUseCase {
	return &EditHabitUseCase{repository}
}

func (e *EditHabitUseCase) Execute(input *EditHabitInputDTO) (*EditHabitOutputDTO, error) {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return nil, domain.ErrEmptyHabitName
	}

	habit, err := e.repository.FindById(input.HabitID)
	if err != nil {
		return nil, err
	}

	oldName := habit.Name
	habit.UpdateName(name)
	if err := e.repository.Update(habit); err != nil {
		return nil, err
	}

	output := &EditHabitOutputDTO{
		HabitID: habit.ID,
		OldName: oldName,
		NewName: habit.Name,
	}
	return output, nil
}
