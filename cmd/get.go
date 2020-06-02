package cmd

import (
	"errors"
	"fmt"
	"github.com/gsxhnd/owl/backend"
	"github.com/gsxhnd/owl/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"io/ioutil"
)

var getComd = &cobra.Command{
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
			logger.Panic("", zap.Error(err))
		}
		conn.Key = key
		v, err := conn.Get()
		if err != nil {
			logger.Panic("", zap.Error(err))
		}
		fmt.Println("value: ", v)
		return nil
	},
}

func init() {
	putComd.PersistentFlags().StringVarP(&endPoint, "endpoint", "e", "http://127.0.0.1", "etcd endpoint")
}
