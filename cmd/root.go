package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	endPoint string
	rootCmd  = &cobra.Command{
		Use:   "gecko",
		Short: "A generator for Cobra based Applications",
		Long:  `gecko is a CLI  applications.`,
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	} else {
		fmt.Println("end")
	}
}

func init() {
	rootCmd.AddCommand(versionCmd, putComd)
}
