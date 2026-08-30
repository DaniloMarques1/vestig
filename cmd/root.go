package cmd

import (
	"log"
	"os"
	"path/filepath"

	"danilo.marques/vestig/internal/infra/db"
	"github.com/spf13/cobra"
)

const layoutBR = "02/01/2006"

var rootCmd = &cobra.Command{
	Use:   "vestig",
	Short: "Vestig - Daily habit and trail manager for the terminal",
	Long:  "Vestig (from Latin 'vestigium': trail/footprint) is a CLI for tracking your daily habits.",
}

func Execute() error {
	defer db.Close()
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Init sqlite dir
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
