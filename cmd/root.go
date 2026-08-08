package cmd

import (
	"log"
	"os"
	"path/filepath"

	"danilo.marques/vestig/internal/infra/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vestig",
	Short: "Vestig - Gerenciador de hábitos e rastros diários via terminal",
	Long:  `Vestig (do latim 'vestigium': rastro/pegada) é uma CLI para rastrear seus hábitos diários.`,
}

func Execute() error {
	defer db.Close()
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Inicialização do SQLite e diretórios da aplicação
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	appDir := filepath.Join(configDir, "vestig")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(appDir, "vestig.db")
	if err := db.InitDB(dbPath); err != nil {
		log.Fatal(err)
	}
}
