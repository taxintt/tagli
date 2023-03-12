package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		listTags()
	},
	Short: "List all tags of Git repository",
}

func listTags() {
	appPath := "."
	r, err := git.PlainOpen(appPath)

	if err != nil {
		log.Fatal(err)
	}
	hashCommit, err := r.Head()
	if err != nil {
		log.Fatal("head", err)
	}
	fmt.Println("hash", hashCommit)

	tags := make(map[plumbing.Hash]string)

	iter, err := r.Tags()
	if err != nil {
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
		tags[ref.Hash()] = ref.Name().Short()

	}

	cIter, err := r.Log(&git.LogOptions{From: hashCommit.Hash()})
	for {
		commit, err := cIter.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if str, ok := tags[commit.Hash]; ok {
			fmt.Printf("%s-%s\n", str, hashCommit.Hash().String()[:8])
			return
		}
	}

	fmt.Println(hashCommit.Hash().String()[:8])
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
