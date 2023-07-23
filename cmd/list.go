package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var RepositoryPath string

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		isJsonOutput, err := cmd.Flags().GetBool("json")
		if err != nil {
			fmt.Println("Failed to parse --json option")
			os.Exit(1)
		}

		listGitTags(isJsonOutput)
	},
	Short: "List all tags of Git repository",
}

func listGitTags(isJsonOutput bool) {
	repo, err := git.PlainOpen(RepositoryPath)
	if errors.Is(err, git.ErrRepositoryNotExists) {
		fmt.Println("Repository not exists")
	} else if err != nil {
		fmt.Println("Failed to open repository")
		os.Exit(1)
	}

	if isJsonOutput {
		printJsonFormat(repo)
	} else {
		printPlainFormat(repo)
	}
	os.Exit(0)
}

func printJsonFormat(repo *git.Repository) {
	tags := make(map[string]string)
	iter, err := repo.Tags()
	if err != nil {
		fmt.Println("Failed to get tags")
		os.Exit(1)
	}
	for {
		ref, err := iter.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			os.Exit(1)
		}
		tags[ref.Name().Short()] = ref.Hash().String()
	}

	bytes, err := json.Marshal(tags)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return
	}

	fmt.Println(string(bytes))
}

func printPlainFormat(repo *git.Repository) {
	iter, err := repo.Tags()
	if err != nil {
		fmt.Println("Failed to get tags")
		os.Exit(1)
	}
	for {
		ref, err := iter.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(ref.Name().Short())
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVar(&RepositoryPath, "repo", ".", "Path to Git repository")
	listCmd.Flags().BoolP("json", "j", false, "Print output in JSON format")
}
