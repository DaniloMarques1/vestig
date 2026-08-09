package repository

import (
	"danilo.marques/vestig/internal/domain"
	"database/sql"
)

type habitLogRepository struct {
	db *sql.DB
}

func NewHabitLogRepository(db *sql.DB) domain.HabitLogRepository {
	return &habitLogRepository{db}
}

func (r *habitLogRepository) Save(habitLog *domain.HabitLog) error {
	query := `insert into habit_logs(habit_id, executed_at) values(?, ?)`
	result, err := r.db.Exec(query, habitLog.HabitID, habitLog.ExecutedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	habitLog.ID = id
	return nil
}
