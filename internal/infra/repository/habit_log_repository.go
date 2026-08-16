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

func (r *habitLogRepository) Find(habitID int64) ([]domain.HabitLog, error) {
	query := `
	select id, habit_id, executed_at
	from habit_logs
	where habit_id = ?
	order by executed_at asc
	`
	rows, err := r.db.Query(query, habitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	logs := make([]domain.HabitLog, 0)
	for rows.Next() {
		log := domain.HabitLog{}
		if err := rows.Scan(&log.ID, &log.HabitID, &log.ExecutedAt); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return logs, nil
}
