package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		checkGitRepositoryExists(RepositoryPath)
		tagName, err := cmd.Flags().GetString("tag")
		if err != nil {
			fmt.Println("Failed to parse --json option")
			os.Exit(1)
		}

		addGitTag(tagName)
	},
	Short: "Add a tag to Git repository",
}

func tagExists(tag string, repo *git.Repository) bool {
	tags, err := repo.TagObjects()
	if err != nil {
		return false
	}
	res := false
	err = tags.ForEach(func(t *object.Tag) error {
		if t.Name == tag {
			res = true
			return fmt.Errorf("Error: tag was found")
		}
		return nil
	})
	if err != nil && err.Error() != "Error: tag was found" {
		return false
	}
	return res
}

func addGitTag(tagName string) {
	repo, err := git.PlainOpen(RepositoryPath)
	if tagExists(tagName, repo) {
		fmt.Printf("Error: tag %s already exists", tagName)
		os.Exit(1)
	}

	head, err := repo.Head()
	if err != nil {
		fmt.Printf("get HEAD error: %s", err)
		os.Exit(1)
	}
	_, err = repo.CreateTag(tagName, head.Hash(), nil)
	if err != nil {
		fmt.Printf("create tag error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Tag %s was created", tagName)
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&RepositoryPath, "repo", "r", ".", "Path to Git repository")
	addCmd.Flags().StringVarP(&tagName, "tag", "t", "", "Git tag")

	addCmd.MarkFlagRequired("tag")
}
