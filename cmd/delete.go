package cmd

import (
	"fmt"
	"strconv"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [id do habito]",
	Short: "Remove um hábito criado anteriormente",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Error deleting habit %v", ID)
		}

		habitRepository := repository.NewHabitRepository(db.DB)
		deleteHabitUseCase := usecase.NewDeleteHabitUseCase(habitRepository)
		input := usecase.DeleteHabitUseCaseInputDTO{HabitID: ID}

		_, err = deleteHabitUseCase.Execute(input)
		if err != nil {
			return fmt.Errorf("Error deleting habit %v. Error %v", ID, err)
		}
		// TODO: print something nice

		return nil
	},
}

func init() {
	rootCmd.AddCommand(DeleteCmd)
}
