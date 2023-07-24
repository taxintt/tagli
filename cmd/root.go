package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	RepositoryPath string
	tagName        string
)

var rootCmd = &cobra.Command{
	Use:   "tagli",
	Short: "CLI tool to handle Git Tag",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
