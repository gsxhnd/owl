package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	version      = "0.2.0"
	gitTag       string
	gitCommit    string
	gitTreeState string
	buildDate    string
	goVersion    = runtime.Version()
	compiler     = runtime.Compiler
	platform     = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: `Show version`,
	Long:  `Show version of owl`,
	Run: func(cmd *cobra.Command, args []string) {
		if gitTag == "" {
			fmt.Println("owl version: ", version)
		} else {
			fmt.Println("owl version: ", gitTag)
		}
		fmt.Println("owl commit: ", gitCommit)
		fmt.Println("owl build date: ", gitTreeState)
		fmt.Println("owl build date: ", buildDate)
		fmt.Println("go version: ", goVersion)
		fmt.Println("go compiler: ", compiler)
		fmt.Println("platform: ", platform)
	},
}
