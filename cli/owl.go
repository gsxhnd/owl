package main

import (
	"os"

	"github.com/heart-dance-x/owl/cli/cmd"
)

func main() {
	if err := cmd.RootCmd.Run(os.Args); err != nil {
		panic(err)
	}
}
