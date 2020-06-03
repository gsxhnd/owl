package cmd

import (
	"github.com/spf13/cobra"
)

var (
	endPoint string
	rootCmd  = &cobra.Command{
		Use:   "owl",
		Short: "A generator for Cobra based Applications",
		Long:  `owl is a CLI  applications.`,
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(putComd)
	rootCmd.AddCommand(getCmd)
}
