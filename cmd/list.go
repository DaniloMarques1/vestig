package cmd

import (
	"fmt"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all habits and today's status",
	RunE: func(cmd *cobra.Command, args []string) error {
		repository := repository.NewHabitRepository(db.DB)
		listHabitsUseCase := usecase.NewListHabitUseCase(repository)

		input := usecase.ListHabitInputDTO{ShowAll: showAll}
		output, err := listHabitsUseCase.Execute(input)
		if err != nil {
			return err
		}

		if len(output.Habits) == 0 {
			fmt.Println("No habits tracked yet. Run 'vestig add <name>' to get started")
			return nil
		}
		// Cabeçalho da tabela
		fmt.Printf("%-4s | %-25s | %-6s\n", "ID", "Habit", "Active")
		fmt.Println("-----+---------------------------+--------")

		for _, h := range output.Habits {
			activeStatus := "Yes"
			if !h.IsActive {
				activeStatus = "No"
			}
			fmt.Printf("%-4d | %-25s | %-6s\n", h.ID, h.Name, activeStatus)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Include finished habits")
}
