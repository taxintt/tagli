package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
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

func checkGitRepositoryExists(RepositoryPath string) {
	_, err := git.PlainOpen(RepositoryPath)
	if errors.Is(err, git.ErrRepositoryNotExists) {
		fmt.Println("Git Repository not exists")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Failed to open repository")
		os.Exit(1)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
