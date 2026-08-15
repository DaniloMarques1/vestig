package repository

import (
	"danilo.marques/vestig/internal/domain"
	"database/sql"
)

type habitRepository struct {
	db *sql.DB
}

func NewHabitRepository(db *sql.DB) domain.HabitRepository {
	return &habitRepository{db}
}

func (hr *habitRepository) Save(habit *domain.Habit) error {
	query := `insert into habits(name, is_active, created_at) values (?, ?, ?)`
	result, err := hr.db.Exec(query, habit.Name, habit.IsActive, habit.CreatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	habit.ID = id
	return nil
}

func (hr *habitRepository) List(showAll bool) ([]domain.Habit, error) {
	query := `select id, name, is_active, created_at
	from habits`
	var args []interface{}
	if !showAll {
		query += ` where is_active = ?`
		args = append(args, showAll)
	}
	query += ` order by created_at`

	rows, err := hr.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	habits := make([]domain.Habit, 0)
	for rows.Next() {
		var habit domain.Habit
		if err := rows.Scan(&habit.ID, &habit.Name, &habit.IsActive, &habit.CreatedAt); err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return habits, nil
}

func (hr *habitRepository) FindById(id int64) (*domain.Habit, error) {
	query := `
		select id, name, is_active, created_at
		from habits
		where id = ?
	`
	var args []interface{}
	args = append(args, id)

	var habit domain.Habit
	err := hr.db.QueryRow(query, args...).Scan(&habit.ID, &habit.Name, &habit.IsActive, &habit.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrHabitNotFound
		}

		return nil, err
	}

	return &habit, nil
}
