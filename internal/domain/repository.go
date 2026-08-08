package domain

type HabitRepository interface {
	Save(*Habit) error
	List() ([]Habit, error)
}
