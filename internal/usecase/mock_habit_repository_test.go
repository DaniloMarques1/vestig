package usecase

import (
	"danilo.marques/vestig/internal/domain"
)

type mockHabitRepository struct {
	saveFn          func(*domain.Habit) error
	findByIdFn      func(int64) (*domain.Habit, error)
	listFn          func(bool) ([]domain.Habit, error)
	updateFn        func(*domain.Habit) error
	called          bool
	lastIsActiveArg bool
}

func (mock *mockHabitRepository) Save(habit *domain.Habit) error {
	if mock.saveFn == nil {
		return nil
	}
	mock.called = true
	return mock.saveFn(habit)
}

func (mock *mockHabitRepository) List(showAll bool) ([]domain.Habit, error) {
	if mock.listFn == nil {
		return nil, nil
	}
	mock.called = true
	mock.lastIsActiveArg = showAll
	return mock.listFn(showAll)
}

func (mock *mockHabitRepository) Update(habit *domain.Habit) error {
	if mock.updateFn == nil {
		return nil
	}
	mock.called = true
	return mock.updateFn(habit)
}

func (mock *mockHabitRepository) FindById(id int64) (*domain.Habit, error) {
	if mock.findByIdFn == nil {
		return nil, nil
	}
	mock.called = true
	return mock.findByIdFn(id)
}
