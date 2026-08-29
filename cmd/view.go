package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [Habit id]",
	Short: "View habit details",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Error viewing habit %v", ID)
		}
		habitRepository := repository.NewHabitRepository(db.DB)
		habitLogRepository := repository.NewHabitLogRepository(db.DB)
		input := &usecase.ViewHabitUseCaseInputDTO{HabitID: ID}
		viewHabitUseCase := usecase.NewViewHabitUseCase(habitRepository, habitLogRepository)
		output, err := viewHabitUseCase.Execute(input)
		if err != nil {
			return fmt.Errorf("Error viewing habit %v. Error %v", ID, err)
		}

		fmt.Println(RenderHabitsExecution(output, time.Now()))
		return nil
	},
}

func RenderHabitsExecution(output *usecase.ViewHabitUseCaseOutputDTO, now time.Time) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s (#%d) — 🔥 %d-%s streak\n\n",
		output.HabitName,
		output.HabitID,
		output.Streak,
		getDaysLabel(output.Streak),
	))

	executionsMap := make(map[string]bool)
	lastIdx := len(output.Executions) - 1
	for i := lastIdx; i >= 0; i-- {
		if len(executionsMap) == 7 {
			break
		}
		dt := output.Executions[i].Format(layoutBR)
		executionsMap[dt] = true
	}
	var headers, statuses []string
	for i := 6; i >= 0; i-- {
		day := now.AddDate(0, 0, -i)
		ptWeekDay := convertWeekToPT(day.Weekday().String())
		headers = append(headers, fmt.Sprintf("%-3s", ptWeekDay))
		if _, exists := executionsMap[day.Format(layoutBR)]; exists {
			statuses = append(statuses, " ✔ ")
		} else {
			statuses = append(statuses, " ✘ ")
		}
	}

	sb.WriteString(strings.Join(headers, "  ") + "\n")
	sb.WriteString(strings.Join(statuses, "  ") + "\n")

	return sb.String()
}

func getDaysLabel(streak int64) string {
	if streak == 1 {
		return "day"
	}

	return "days"
}

func convertWeekToPT(week string) string {
	if len(week) > 3 {
		return week[:3]
	}
	return week
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
