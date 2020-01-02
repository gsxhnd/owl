package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type MysqlDbInfo struct {
	Host     string `validate:"required"` // Host. 地址
	Port     int    `validate:"required"` // Port 端口号
	Username string `validate:"required"` // Username 用户名
	Password string // Password 密码
	Database string `validate:"required"` // Database 数据库名
}

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	//databaseName string
	//host string
	//port int
	//user,password string
	//outdir string
	mysqlInfo MysqlDbInfo

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
	//rootCmd.PersistentFlags().StringP("host", "H", "127.0.0.1", "sql host")

	rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Host, "host", "H", "127.0.0.1", "数据库地址.")
	//err := rootCmd.MarkPersistentFlagRequired("host")
	//if err != nil {
	//	fmt.Println("error:")
	//	os.Exit(1)
	//}
	//rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Username, "user", "u", "", "用户名.")
	//rootCmd.MarkFlagRequired("user")
	//
	//rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Password, "password", "p", "", "密码.")
	//rootCmd.MarkFlagRequired("password")
	//
	//rootCmd.PersistentFlags().StringVarP(&mysqlInfo.Database, "database", "d", "", "数据库名")
	//rootCmd.MarkFlagRequired("database")
	//
	//rootCmd.PersistentFlags().StringVarP(&outDir, "outdir", "o", "", "输出目录")
	//rootCmd.MarkFlagRequired("outdir")

	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "mit", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().String("keo", "keo", "keo")
	_ = viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	_ = viper.BindPFlag("keo", rootCmd.PersistentFlags().Lookup("keo"))
	_ = viper.BindEnv("keo", "TEST_KEO")
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
}

func initConfig() {
	viper.SetEnvPrefix("test")
	viper.AutomaticEnv()
}
