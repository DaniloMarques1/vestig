package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [nome do hábito]",
	Short: "Cadastra um novo hábito",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Inserir hábito no banco SQLite
		fmt.Printf("Opa mano tu quer adicionar ne %v\n", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
