package cmd

import (
	"fmt"
	"strings"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [nome do hábito]",
	Short: "Cadastra um novo hábito",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		allArgs := args[0:]
		habit := strings.Join(allArgs, " ")
		habitRepository := repository.NewHabitRepository(db.DB)
		addHabitUseCase := usecase.NewAddHabitUseCase(habitRepository)
		addHabitInputDTO := usecase.AddHabitInputDTO{Name: habit}

		output, err := addHabitUseCase.Execute(&addHabitInputDTO)
		if err != nil {
			return err
		}

		fmt.Printf("✔ Habit '%s' (ID %d) added!\n", output.Name, output.ID)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
