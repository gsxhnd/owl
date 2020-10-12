package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"runtime"
)

var (
	version      = "1.2.0"
	gitTag       string
	gitCommit    string
	gitTreeState string
	buildDate    string
	goVersion    = runtime.Version()
	compiler     = runtime.Compiler
	platform     = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)
var versionCmd = &cli.Command{
	Name:        "version",
	Usage:       "show version",
	UsageText:   "",
	Description: "",
	Action: func(c *cli.Context) error {
		if gitTag == "" {
			fmt.Println("owl version: ", version)
		} else {
			fmt.Println("owl version: ", gitTag)
		}
		fmt.Println("owl commit: ", gitCommit)
		fmt.Println("owl tree state: ", gitTreeState)
		fmt.Println("owl build date: ", buildDate)
		fmt.Println("go version: ", goVersion)
		fmt.Println("go compiler: ", compiler)
		fmt.Println("platform: ", platform)
		return nil
	},
}
