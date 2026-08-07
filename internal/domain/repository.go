package domain

type HabitRepository interface {
	Save(*Habit) error
}
