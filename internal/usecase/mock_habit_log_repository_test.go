package usecase

import (
	"time"

	"danilo.marques/vestig/internal/domain"
)

type mockHabitLogRepository struct {
	saveFn            func(*domain.HabitLog) error
	findFn            func(int64) ([]domain.HabitLog, error)
	findExecutionAtFn func(time.Time) (bool, error)
	called            bool
}

func (mock *mockHabitLogRepository) Save(log *domain.HabitLog) error {
	if mock.saveFn == nil {
		return nil
	}
	mock.called = true
	return mock.saveFn(log)
}

func (mock *mockHabitLogRepository) Find(habitID int64) ([]domain.HabitLog, error) {
	if mock.findFn == nil {
		return nil, nil
	}
	mock.called = true
	return mock.findFn(habitID)
}

func (mock *mockHabitLogRepository) FindExecutionAt(dt time.Time) (bool, error) {
	if mock.findExecutionAtFn == nil {
		return false, nil
	}
	mock.called = true
	return mock.findExecutionAtFn(dt)
}
