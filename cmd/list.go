package cmd

import (
	"fmt"

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
	fmt.Println("This is the list of repository :)")
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
