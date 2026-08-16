package domain

type HabitRepository interface {
	Save(*Habit) error
	List(showAll bool) ([]Habit, error)
	FindById(id int64) (*Habit, error)
	Update(*Habit) error
}

type HabitLogRepository interface {
	Save(*HabitLog) error
	Find(int64) ([]HabitLog, error)
}
