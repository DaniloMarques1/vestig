package cmd

import (
	"fmt"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todos os hábitos e o status do dia",
	RunE: func(cmd *cobra.Command, args []string) error {
		repository := repository.NewHabitRepository(db.DB)
		listHabitsUseCase := usecase.NewListHabitUseCase(repository)

		output, err := listHabitsUseCase.Execute()
		if err != nil {
			return err
		}

		if len(output.Habits) == 0 {
			fmt.Println("Nenhum hábito cadastrado ainda. Use 'vestig add <nome>' para começar.")
			return nil
		}
		// Cabeçalho da tabela
		fmt.Printf("%-4s | %-25s | %-6s\n", "ID", "Hábito", "Ativo")
		fmt.Println("-----+---------------------------+--------")

		for _, h := range output.Habits {
			activeStatus := "Sim"
			if !h.IsActive {
				activeStatus = "Não"
			}
			fmt.Printf("%-4d | %-25s | %-6s\n", h.ID, h.Name, activeStatus)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
