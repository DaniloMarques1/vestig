package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var executedDate string

var executeCmd = &cobra.Command{
	Use:   "execute [Habit id]",
	Short: "Record a habit execution",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		argumentId := args[0]
		ID, err := strconv.ParseInt(argumentId, 10, 64)
		if err != nil {
			return fmt.Errorf("Error executing habit %v", ID)
		}

		executedAt, err := getExecutedDate(executedDate)
		if err != nil {
			return err
		}

		habitRepository := repository.NewHabitRepository(db.DB)
		habitLogRepository := repository.NewHabitLogRepository(db.DB)
		executeHabitUseCase := usecase.NewExecuteHabitUseCase(habitRepository, habitLogRepository)
		input := usecase.ExecuteHabitInputDTO{ID: ID, ExecutedAt: executedAt}

		output, err := executeHabitUseCase.Execute(input)
		if err != nil {
			return fmt.Errorf("Error executing habit %v. Error: %v", ID, err)
		}

		fmt.Printf("\033[32m✔\033[0m Habit '%s' marked as executed for today %v!\n", output.HabitName, getDateAsPTBR(output.ExecutedAt))
		return nil
	},
}

func getExecutedDate(executedDate string) (time.Time, error) {
	if len(executedDate) == 0 {
		return time.Now(), nil
	}

	t, err := time.ParseInLocation(layoutBR, executedDate, time.Local)
	if err != nil {
		return time.Time{}, errors.New("Invalid date format. Expected format: DD/MM/YYYY (e.g., 27/03/2006)")
	}

	return t, nil
}

func getDateAsPTBR(t time.Time) string {
	return t.Format(layoutBR)
}

func init() {
	rootCmd.AddCommand(executeCmd)
	executeCmd.Flags().StringVarP(&executedDate, "include-date", "i", "", "Add execution to a specific date")
}
