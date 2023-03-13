package cmd

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	Version = "dev"
)

func getVersion() string {
	return fmt.Sprintf(`git-tag-manager

Version: %s
OS: %s
Arch: %s`, Version, runtime.GOOS, runtime.GOARCH)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			Version = "unknown"
		}
		Version = info.Main.Version
		fmt.Println(getVersion())
	},
	Short: "Show cli version info",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
