package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [Habit id] [New habit name]",
	Short: "Edit habit's name",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		habitID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Could not edit habit")
		}

		newNameArg := args[1:]
		newHabitName := strings.Join(newNameArg, " ")

		habitRepository := repository.NewHabitRepository(db.DB)

		editHabitUseCase := usecase.NewEditHabitUseCase(habitRepository)
		input := &usecase.EditHabitInputDTO{HabitID: habitID, Name: newHabitName}
		output, err := editHabitUseCase.Execute(input)
		if err != nil {
			return fmt.Errorf("Error editing habit %v.", habitID)
		}

		fmt.Printf("\033[32m✔\033[0m Habit '%s' renamed to '%s'!\n", output.OldName, output.NewName)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
