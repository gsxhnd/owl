package main

import (
	"os"

	"github.com/gsxhnd/owl/cli/cmd"
)

func main() {
	if err := cmd.RootCmd.Run(os.Args); err != nil {
		panic(err)
	}
}
