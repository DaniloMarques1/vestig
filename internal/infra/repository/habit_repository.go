package repository

import (
	"danilo.marques/vestig/internal/domain"
	"database/sql"
)

type HabitRepository struct {
	db *sql.DB
}

func NewHabitRepository(db *sql.DB) *HabitRepository {
	return &HabitRepository{db}
}

func (hr *HabitRepository) Save(habit *domain.Habit) error {
	query := `insert into habits(name, done, created_at) values (?, ?, ?)`
	if _, err := hr.db.Exec(query, habit.Name, habit.Done, habit.CreatedAt); err != nil {
		return err
	}

	return nil
}
