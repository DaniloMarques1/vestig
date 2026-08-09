package cmd

import (
	"log"
	"strconv"
	"time"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id do hábito]",
	Short: "Marca a execução de um hábito",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		argumentId := args[0]
		ID, err := strconv.ParseInt(argumentId, 10, 64)
		if err != nil {
			log.Fatal(err) // TODO: message explaining id should be a number
		}

		habitRepository := repository.NewHabitRepository(db.DB)
		habitLogRepository := repository.NewHabitLogRepository(db.DB)
		executeHabitUseCase := usecase.NewExecuteHabitUseCase(habitRepository, habitLogRepository)
		executeHabitInputDTO := usecase.ExecuteHabitInputDTO{ID: ID, ExecutedAt: time.Now()}

		if err := executeHabitUseCase.Execute(executeHabitInputDTO); err != nil {
			log.Fatal(err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
