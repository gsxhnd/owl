package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	endPoint string
	RootCmd  = cli.NewApp()
)

func init() {

	RootCmd.Usage = "owl"
	RootCmd.Version = ""
	RootCmd.HideVersion = true
	RootCmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "endpoint",
			Aliases:     []string{"e"},
			Value:       "http://127.0.0.1:2379",
			Usage:       "",
			Destination: &endPoint,
		},
	}
	RootCmd.Commands = []*cli.Command{
		getCmd,
		getKeysCmd,
		putCmd,
		versionCmd,
	}
}
