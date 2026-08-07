package usecase

import (
	"danilo.marques/vestig/internal/domain"
)

type AddHabitUseCase struct {
	repository domain.HabitRepository
}

type AddHabitInputDTO struct {
	Name string
}

type AddHabitOutputDTO struct {
	ID   int64
	Name string
}

func NewAddHabitUseCase(repository domain.HabitRepository) *AddHabitUseCase {
	return &AddHabitUseCase{repository}
}

func (ah *AddHabitUseCase) Execute(input *AddHabitInputDTO) (*AddHabitOutputDTO, error) {
	habit, err := domain.NewHabit(input.Name)
	if err != nil {
		return nil, err
	}

	if err := ah.repository.Save(habit); err != nil {
		return nil, err
	}

	output := &AddHabitOutputDTO{
		ID:   habit.ID,
		Name: habit.Name,
	}
	return output, nil
}
