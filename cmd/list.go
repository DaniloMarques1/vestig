package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todos os hábitos e o status do dia",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Consultar hábitos no SQLite e exibir no terminal
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
