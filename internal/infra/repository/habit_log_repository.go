package repository

import (
	"database/sql"
	"time"

	"danilo.marques/vestig/internal/domain"
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

func (r *habitLogRepository) FindExecutionAt(habitID int64, dt time.Time) (bool, error) {
	startOfDay := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1)

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM habit_logs
			WHERE executed_at >= ? AND executed_at < ?
			and habit_id = ?
		)
	`

	var executionExists bool
	if err := r.db.QueryRow(query, startOfDay, endOfDay, habitID).Scan(&executionExists); err != nil {
		return false, err
	}

	return executionExists, nil
}
