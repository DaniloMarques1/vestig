package usecase

import (
	"errors"
	"testing"
	"time"

	"danilo.marques/vestig/internal/domain"
)

func TestViewHabitUseCase_Execute(t *testing.T) {
	t.Run("deve visualizar um hábito e suas execuções com sucesso", func(t *testing.T) {
		now := time.Now()
		nowUTC := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, time.UTC)

		habitRepo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				if id != 1 {
					t.Errorf("esperava ID 1, recebeu %d", id)
				}
				return &domain.Habit{
					ID:        1,
					Name:      "Beber Água",
					IsActive:  true,
					CreatedAt: nowUTC,
				}, nil
			},
		}

		logRepo := &mockHabitLogRepository{
			findFn: func(habitID int64) ([]domain.HabitLog, error) {
				if habitID != 1 {
					t.Errorf("esperava habitID 1, recebeu %d", habitID)
				}
				return []domain.HabitLog{
					{ID: 10, HabitID: 1, ExecutedAt: nowUTC},
				}, nil
			},
		}

		uc := NewViewHabitUseCase(habitRepo, logRepo)
		input := &ViewHabitUseCaseInputDTO{HabitID: 1}

		output, err := uc.Execute(input)

		if err != nil {
			t.Fatalf("esperava erro nil, recebeu: %v", err)
		}

		if !habitRepo.called {
			t.Error("esperava que o repositório de hábitos tivesse sido chamado")
		}

		if !logRepo.called {
			t.Error("esperava que o repositório de logs tivesse sido chamado")
		}

		if output == nil {
			t.Fatal("output não deveria ser nil")
		}

		if output.HabitID != 1 {
			t.Errorf("esperava HabitID 1, recebeu %d", output.HabitID)
		}

		if output.HabitName != "Beber Água" {
			t.Errorf("esperava nome 'Beber Água', recebeu '%s'", output.HabitName)
		}

		if output.Streak != 1 {
			t.Errorf("esperava streak 1, recebeu %d", output.Streak)
		}

		if len(output.Executions) != 1 {
			t.Fatalf("esperava 1 execução, recebeu %d", len(output.Executions))
		}

		if !output.Executions[0].Equal(nowUTC) {
			t.Errorf("esperava execução %v, recebeu %v", nowUTC, output.Executions[0])
		}
	})

	t.Run("deve retornar hábito sem execuções e streak zero quando não houver histórico", func(t *testing.T) {
		habitRepo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				return &domain.Habit{
					ID:       1,
					Name:     "Beber Água",
					IsActive: true,
				}, nil
			},
		}

		logRepo := &mockHabitLogRepository{
			findFn: func(habitID int64) ([]domain.HabitLog, error) {
				return []domain.HabitLog{}, nil
			},
		}

		uc := NewViewHabitUseCase(habitRepo, logRepo)
		input := &ViewHabitUseCaseInputDTO{HabitID: 1}

		output, err := uc.Execute(input)

		if err != nil {
			t.Fatalf("esperava erro nil, recebeu: %v", err)
		}

		if output == nil {
			t.Fatal("output não deveria ser nil")
		}

		if output.Streak != 0 {
			t.Errorf("esperava streak 0, recebeu %d", output.Streak)
		}

		if len(output.Executions) != 0 {
			t.Errorf("esperava 0 execuções, recebeu %d", len(output.Executions))
		}
	})

	t.Run("deve retornar erro quando o hábito não for encontrado no repositório", func(t *testing.T) {
		habitRepo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				return nil, domain.ErrHabitNotFound
			},
		}

		logRepo := &mockHabitLogRepository{}

		uc := NewViewHabitUseCase(habitRepo, logRepo)
		input := &ViewHabitUseCaseInputDTO{HabitID: 99}

		output, err := uc.Execute(input)

		if !errors.Is(err, domain.ErrHabitNotFound) {
			t.Errorf("esperava erro ErrHabitNotFound, recebeu: %v", err)
		}

		if output != nil {
			t.Errorf("esperava output nil, recebeu: %+v", output)
		}

		if !habitRepo.called {
			t.Error("esperava que o repositório de hábitos tivesse sido chamado")
		}

		if logRepo.called {
			t.Error("repositório de logs não deveria ter sido chamado quando o hábito não é encontrado")
		}
	})

	t.Run("deve retornar erro quando falhar ao buscar os logs no repositório", func(t *testing.T) {
		errDb := errors.New("falha de conexão com o banco de dados")

		habitRepo := &mockHabitRepository{
			findByIdFn: func(id int64) (*domain.Habit, error) {
				return &domain.Habit{
					ID:       1,
					Name:     "Ler Livro",
					IsActive: true,
				}, nil
			},
		}

		logRepo := &mockHabitLogRepository{
			findFn: func(habitID int64) ([]domain.HabitLog, error) {
				return nil, errDb
			},
		}

		uc := NewViewHabitUseCase(habitRepo, logRepo)
		input := &ViewHabitUseCaseInputDTO{HabitID: 1}

		output, err := uc.Execute(input)

		if !errors.Is(err, errDb) {
			t.Errorf("esperava erro '%v', recebeu: '%v'", errDb, err)
		}

		if output != nil {
			t.Errorf("esperava output nil, recebeu: %+v", output)
		}

		if !habitRepo.called {
			t.Error("esperava que o repositório de hábitos tivesse sido chamado")
		}

		if !logRepo.called {
			t.Error("esperava que o repositório de logs tivesse sido chamado")
		}
	})
}
