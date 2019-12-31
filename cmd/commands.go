package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "gecko",
		Short: "A generator for Cobra based Applications",
		Long:  `gecko is a CLI  applications.`,
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("start pre run")
			fmt.Println(cmd)
			fmt.Println(args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			fmt.Println(cmd)
			fmt.Println(args)
			fmt.Println("config: ", viper.GetString("config"))
			fmt.Println("author: ", viper.GetString("author"))
			fmt.Println("license: ", viper.GetString("license"))
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	_ = viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	_ = viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}

	viper.AutomaticEnv()
}
