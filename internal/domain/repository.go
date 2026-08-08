package domain

type HabitRepository interface {
	Save(*Habit) error
	List(showAll bool) ([]Habit, error)
}
