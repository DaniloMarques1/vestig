package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vestig",
	Short: "Vestig - Gerenciador de hábitos e rastros diários via terminal",
	Long:  `Vestig (do latim 'vestigium': rastro/pegada) é uma CLI para rastrear seus hábitos diários.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Inicialização do SQLite e diretórios da aplicação
}
