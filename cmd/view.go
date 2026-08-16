package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"danilo.marques/vestig/internal/infra/db"
	"danilo.marques/vestig/internal/infra/repository"
	"danilo.marques/vestig/internal/usecase"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [id do habito]",
	Short: "Lista todos os hábitos e o status do dia",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err) // TODO
		}
		habitRepository := repository.NewHabitRepository(db.DB)
		habitLogRepository := repository.NewHabitLogRepository(db.DB)
		input := &usecase.ViewHabitUseCaseInputDTO{HabitID: ID}
		viewHabitUseCase := usecase.NewViewHabitUseCase(habitRepository, habitLogRepository)
		output, err := viewHabitUseCase.Execute(input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(RenderHabitsExecution(output, time.Now()))
		return nil
	},
}

func RenderHabitsExecution(output *usecase.ViewHabitUseCaseOutputDTO, now time.Time) string {
	var sb strings.Builder

	// TODO: fix dias, it depends on streak value
	sb.WriteString(fmt.Sprintf("%s (#%d) — 🔥 %d dias seguidos\n\n", output.HabitName, output.HabitID, output.Streak))

	executionsMap := make(map[string]bool)
	lastIdx := len(output.Executions) - 1
	for i := lastIdx; i >= 0; i-- {
		if len(executionsMap) == 7 {
			break
		}
		dt := output.Executions[i].Format("02/01/2006")
		executionsMap[dt] = true
	}
	var headers, statuses []string
	for i := 6; i >= 0; i-- {
		day := now.AddDate(0, 0, -i)
		ptWeekDay := convertWeekToPT(day.Weekday().String())
		headers = append(headers, fmt.Sprintf("%-3s", ptWeekDay))
		if _, exists := executionsMap[day.Format("02/01/2006")]; exists {
			statuses = append(statuses, " ✔ ")
		} else {
			statuses = append(statuses, " ✘ ")
		}
	}

	sb.WriteString(strings.Join(headers, "  ") + "\n")
	sb.WriteString(strings.Join(statuses, "  ") + "\n")

	return sb.String()
}

func convertWeekToPT(week string) string {
	switch week {
	case "Monday":
		return "Seg"
	case "Tuesday":
		return "Ter"
	case "Wednesday":
		return "Qua"
	case "Thursday":
		return "Qui"
	case "Friday":
		return "Sex"
	case "Saturday":
		return "Sab"
	case "Sunday":
		return "Dom"
	default:
		return week
	}
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
