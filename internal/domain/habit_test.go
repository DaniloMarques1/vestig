package domain

import (
	"errors"
	"testing"
	"time"
)

func TestNewHabit(t *testing.T) {
	t.Run("deve criar um hábito válido com os valores padrão corretos", func(t *testing.T) {
		before := time.Now()

		habit, err := NewHabit("Meditar")

		after := time.Now()

		if err != nil {
			t.Fatalf("esperava erro nil, recebeu: %v", err)
		}

		if habit == nil {
			t.Fatal("habito não deveria ser nil")
		}

		if habit.Name != "Meditar" {
			t.Errorf("esperava nome 'Meditar', recebeu '%s'", habit.Name)
		}

		if habit.IsActive != true {
			t.Errorf("esperava IsActive true, recebeu %t", habit.IsActive)
		}

		if habit.ID != 0 {
			t.Errorf("esperava ID inicial 0, recebeu %d", habit.ID)
		}

		// Valida se o CreatedAt foi preenchido com o momento atual da criação
		if habit.CreatedAt.Before(before) || habit.CreatedAt.After(after) {
			t.Errorf("CreatedAt %v fora do intervalo esperado [%v, %v]", habit.CreatedAt, before, after)
		}
	})

	t.Run("deve remover espaços em branco nas extremidades do nome", func(t *testing.T) {
		habit, err := NewHabit("   Correr no parque   ")

		if err != nil {
			t.Fatalf("esperava erro nil, recebeu: %v", err)
		}

		if habit.Name != "Correr no parque" {
			t.Errorf("esperava nome 'Correr no parque', recebeu '%s'", habit.Name)
		}
	})

	t.Run("deve retornar ErrEmptyHabitName quando o nome for string vazia", func(t *testing.T) {
		habit, err := NewHabit("")

		if !errors.Is(err, ErrEmptyHabitName) {
			t.Errorf("esperava erro ErrEmptyHabitName, recebeu: %v", err)
		}

		if habit != nil {
			t.Errorf("esperava hábito nil ao falhar a criação, recebeu: %+v", habit)
		}
	})

	t.Run("deve retornar ErrEmptyHabitName quando o nome contiver apenas espaços", func(t *testing.T) {
		habit, err := NewHabit("     ")

		if !errors.Is(err, ErrEmptyHabitName) {
			t.Errorf("esperava erro ErrEmptyHabitName, recebeu: %v", err)
		}

		if habit != nil {
			t.Errorf("esperava hábito nil ao falhar a criação, recebeu: %+v", habit)
		}
	})
}
