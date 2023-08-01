package cmd

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var version string

func getVersion(version string) string {
	return fmt.Sprintf(`tagli

Version: %s
OS: %s
Arch: %s`, version, runtime.GOOS, runtime.GOARCH)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			version = "unknown"
		}
		fmt.Println(getVersion(info.Main.Version))
	},
	Short: "Show cli version info",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
