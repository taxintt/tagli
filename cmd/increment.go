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
		checkGitRepositoryExists(RepositoryPath)

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

		versionType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println("Failed to parse --type option")
			os.Exit(1)
		}
		incrementGitTag(tagName, incrementValue, versionType)
	},
	Short: "Increment a tag of Git repository",
}

func incrementGitTag(tagName string, incrementValue int, versionType string) {
	repo, err := git.PlainOpen(RepositoryPath)
	head, err := repo.Head()
	if err != nil {
		fmt.Printf("get HEAD error: %s", err)
		os.Exit(1)
	}

	incrementTag, err := incrementVersion(tagName, incrementValue, versionType)
	if err != nil {
		fmt.Printf("increment version error: %s", err)
		os.Exit(1)
	}

	_, err = repo.CreateTag(incrementTag, head.Hash(), nil)
	if err != nil {
		fmt.Printf("create tag error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Tag %s was created\n", incrementTag)
	os.Exit(0)
}

func incrementVersion(tagName string, incrementValue int, versionType string) (string, error) {
	versionObj, err := semver.NewVersion(tagName)
	if err != nil {
		return "", fmt.Errorf("Invalid version: %s", tagName)
	}

	switch versionType {
	case "major":
		return incrementMajor(versionObj, incrementValue), nil
	case "minor":
		return incrementMinor(versionObj, incrementValue), nil
	case "patch":
		return incrementPatch(versionObj, incrementValue), nil
	default:
		return "", fmt.Errorf("Invalid version type: %s", versionType)
	}
}

func containsVprefix(tagName string) bool {
	return tagName[0] == 'v'
}

func incrementMajor(version *semver.Version, incrementValue int) string {
	m := version.IncMajor()
	if containsVprefix(version.Original()) {
		return "v" + m.String()
	}
	return m.String()
}

func incrementMinor(version *semver.Version, incrementValue int) string {
	m := version.IncMinor()
	if containsVprefix(version.Original()) {
		return "v" + m.String()
	}
	return m.String()
}

func incrementPatch(version *semver.Version, incrementValue int) string {
	m := version.IncPatch()
	if containsVprefix(version.Original()) {
		return "v" + m.String()
	}
	return m.String()
}

func init() {
	incrementCmd.Flags().StringP("tag", "t", "", "tag name")
	incrementCmd.Flags().IntP("increment", "i", 1, "increment value")
	incrementCmd.Flags().StringP("repository", "r", ".", "repository path")
	incrementCmd.Flags().StringP("version", "v", "", "version (patch, minor, major))")

	incrementCmd.MarkFlagRequired("tag")
	incrementCmd.MarkFlagRequired("version")
	rootCmd.AddCommand(incrementCmd)
}
