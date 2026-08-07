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

func NewAddHabitUseCase(repository domain.HabitRepository) *AddHabitUseCase {
	return &AddHabitUseCase{repository}
}

func (ah *AddHabitUseCase) Execute(input *AddHabitInputDTO) error {
	habit, err := domain.NewHabit(input.Name)
	if err != nil {
		return err
	}

	if err := ah.repository.Save(habit); err != nil {
		return err
	}

	return nil
}
