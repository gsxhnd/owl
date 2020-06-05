package cmd

import (
	"errors"
	"fmt"
	"github.com/gsxhnd/owl/backend"
	"github.com/spf13/cobra"
	"log"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "get",
	Long:    "get",
	Example: "get",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return nil
		} else {
			return errors.New("need one args")
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var key = args[0]
		conn, err := backend.NewEtcdClient([]string{endPoint})
		if err != nil {
			log.Panic(err)
		}
		v, err := conn.Get(key)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("value: ", v)
		return nil
	},
}

func init() {
	getCmd.PersistentFlags().StringVarP(&endPoint, "endpoint", "e", "http://127.0.0.1", "etcd endpoint")
}
