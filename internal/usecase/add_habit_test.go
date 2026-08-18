package usecase

import (
	"errors"
	"testing"

	"danilo.marques/vestig/internal/domain"
)

func TestAddHabitUseCase_Execute(t *testing.T) {
	t.Run("deve criar um hábito com sucesso", func(t *testing.T) {
		repo := &mockHabitRepository{
			saveFn: func(h *domain.Habit) error {
				if h.Name != "Beber Água" {
					t.Errorf("esperava nome 'Beber Água', recebeu '%s'", h.Name)
				}
				// Simula o banco gerando o ID no struct
				h.ID = 1
				return nil
			},
		}

		uc := NewAddHabitUseCase(repo)
		input := &AddHabitInputDTO{Name: "Beber Água"}

		output, err := uc.Execute(input)

		if err != nil {
			t.Fatalf("esperava erro nil, recebeu: %v", err)
		}

		if !repo.called {
			t.Error("esperava que o repositório tivesse sido chamado")
		}

		if output == nil {
			t.Fatal("output não deveria ser nil")
		}

		if output.ID != 1 {
			t.Errorf("esperava ID 1, recebeu %d", output.ID)
		}

		if output.Name != "Beber Água" {
			t.Errorf("esperava nome 'Beber Água', recebeu '%s'", output.Name)
		}
	})

	t.Run("deve retornar erro quando o nome for vazio", func(t *testing.T) {
		repo := &mockHabitRepository{}
		uc := NewAddHabitUseCase(repo)

		input := &AddHabitInputDTO{Name: "   "}

		output, err := uc.Execute(input)

		if !errors.Is(err, domain.ErrEmptyHabitName) {
			t.Errorf("esperava erro ErrEmptyHabitName, recebeu: %v", err)
		}

		if output != nil {
			t.Errorf("esperava output nil, recebeu: %+v", output)
		}

		if repo.called {
			t.Error("repositório não deveria ter sido chamado quando a validação falha")
		}
	})

	t.Run("deve retornar erro quando falhar ao salvar no repositório", func(t *testing.T) {
		errDb := errors.New("falha de conexão com o banco")

		repo := &mockHabitRepository{
			saveFn: func(h *domain.Habit) error {
				return errDb
			},
		}

		uc := NewAddHabitUseCase(repo)
		input := &AddHabitInputDTO{Name: "Ler Livro"}

		output, err := uc.Execute(input)

		if !errors.Is(err, errDb) {
			t.Errorf("esperava erro '%v', recebeu: '%v'", errDb, err)
		}

		if output != nil {
			t.Errorf("esperava output nil, recebeu: %+v", output)
		}

		if !repo.called {
			t.Error("esperava que o repositório tivesse sido chamado")
		}
	})
}
