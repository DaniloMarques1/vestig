package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var executeCmd = &cobra.Command{
	Use:   "execute [id do hábito]",
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
		input := usecase.ExecuteHabitInputDTO{ID: ID, ExecutedAt: time.Now()}

		output, err := executeHabitUseCase.Execute(input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\033[32m✔\033[0m Hábito '%s' marcado como concluído hoje!\n", output.HabitName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
}
