package usecase

import (
	"danilo.marques/vestig/internal/domain"
)

type ListHabitUseCase struct {
	repository domain.HabitRepository
}

func NewListHabitUseCase(repository domain.HabitRepository) *ListHabitUseCase {
	return &ListHabitUseCase{repository}
}

type ListHabitInputDTO struct {
	ShowAll bool
}

type ListHabitOutputDTO struct {
	Habits []HabitDTO
}

type HabitDTO struct {
	ID       int64
	Name     string
	IsActive bool
}

func (lh *ListHabitUseCase) Execute(input ListHabitInputDTO) (*ListHabitOutputDTO, error) {
	isActive := !input.ShowAll
	habits, err := lh.repository.List(isActive)
	if err != nil {
		return nil, err
	}
	habitsDTO := make([]HabitDTO, 0, len(habits))
	for _, habit := range habits {
		habitDTO := HabitDTO{
			ID:       habit.ID,
			Name:     habit.Name,
			IsActive: habit.IsActive,
		}
		habitsDTO = append(habitsDTO, habitDTO)
	}

	output := &ListHabitOutputDTO{
		Habits: habitsDTO,
	}

	return output, nil
}
