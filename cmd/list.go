package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

var (
	RepositoryPath = "."
)

var encodeCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		listTags()
	},
	Short: "List all tags of Git repository",
}

func listTags() {
	repo, err := git.PlainOpen(RepositoryPath)
	if err != nil {
		log.Fatal(err)
	}

	iterator, err := repo.Tags()
	if err != nil {
		os.Exit(1)
	}

	if err := iterator.ForEach(func(ref *plumbing.Reference) error {
		obj, err := repo.TagObject(ref.Hash())
		switch err {
		case nil:
			fmt.Println(obj.Name)
		case plumbing.ErrObjectNotFound:
			fmt.Println("there is no tags")
		default:
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
