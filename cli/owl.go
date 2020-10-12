package main

import (
	"github.com/gsxhnd/owl/cli/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Run(os.Args); err != nil {
		panic(err)
	}
}
