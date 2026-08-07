package cmd

import (
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
		habit := args[0]

		habitRepository := repository.NewHabitRepository(db.DB)
		addHabitUseCase := usecase.NewAddHabitUseCase(habitRepository)
		addHabitInputDTO := usecase.AddHabitInputDTO{Name: habit}

		if err := addHabitUseCase.Execute(&addHabitInputDTO); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
