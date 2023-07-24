package cmd

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var incrementCmd = &cobra.Command{
	Use: "increment",
	Run: func(cmd *cobra.Command, args []string) {
		tagName, err := cmd.Flags().GetString("tag")
		if err != nil {
			fmt.Println("Failed to parse --tag option")
			os.Exit(1)
		}
		incrementValue, err := cmd.Flags().GetInt("increment")
		if err != nil {
			fmt.Println("Failed to parse --increment option")
			os.Exit(1)
		}

		version, err := cmd.Flags().GetString("version")
		if err != nil {
			fmt.Println("Failed to parse --version option")
			os.Exit(1)
		}
		incrementGitTag(tagName, incrementValue, version)
	},
	Short: "Increment a tag of Git repository",
}

func incrementGitTag(tagName string, incrementValue int, version string) {
	repo, err := git.PlainOpen(RepositoryPath)
	if tagExists(tagName, repo) {
		fmt.Printf("tag %s already exists", tagName)
	}

	head, err := repo.Head()
	if err != nil {
		fmt.Printf("get HEAD error: %s", err)
		os.Exit(1)
	}

	incrementTag := increment(tagName, incrementValue, version)
	_, err = repo.CreateTag(incrementTag, head.Hash(), nil)
	if err != nil {
		fmt.Printf("create tag error: %s", err)
		os.Exit(1)
	}
}

func increment(tagName string, incrementValue int, version string) string {
	versionObj, err := semver.NewVersion(tagName)
	if err != nil {
		fmt.Printf("create tag error: %s", err)
		os.Exit(1)
	}

	switch version {
	case "major":
		return incrementMajor(versionObj, incrementValue)
	case "minor":
		return incrementMinor(versionObj, incrementValue)
	case "patch":
		return incrementPatch(versionObj, incrementValue)
	default:
		fmt.Println("Invalid version")
		os.Exit(1)
	}
	return ""
}

func incrementMajor(version *semver.Version, incrementValue int) string {
	m := version.IncMajor()
	return m.String()
}

func incrementMinor(version *semver.Version, incrementValue int) string {
	m := version.IncMinor()
	return m.String()
}

func incrementPatch(version *semver.Version, incrementValue int) string {
	m := version.IncPatch()
	return m.String()
}

func init() {
	incrementCmd.Flags().StringP("tag", "t", "", "tag name")
	incrementCmd.Flags().IntP("increment", "i", 1, "increment value")
	incrementCmd.Flags().StringP("repository", "r", ".", "repository path")
	incrementCmd.Flags().StringP("version", "v", "", "version")

	incrementCmd.MarkFlagRequired("tag")
	incrementCmd.MarkFlagRequired("version")
	rootCmd.AddCommand(incrementCmd)
}
