package cmd

import (
	"errors"
	"github.com/gsxhnd/gecko/backend"
	"github.com/gsxhnd/gecko/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"io/ioutil"
)

var putComd = &cobra.Command{
	Use:     "put",
	Short:   "put",
	Long:    "put",
	Example: "put",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 2 {
			return nil
		} else {
			return errors.New("need two args")
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var key = args[0]
		var filePath = args[1]
		conn, err := backend.NewEtcdClient([]string{endPoint})
		if err != nil {
			logger.Panic("", zap.Error(err))
		}
		yamlFile, err := ioutil.ReadFile(filePath)

		conn.Key = key
		conn.Value = string(yamlFile)
		err = conn.Put()
		if err != nil {
			logger.Panic("", zap.Error(err))
		}
		return nil
	},
}

func init() {
	putComd.PersistentFlags().StringVarP(&endPoint, "endpoint", "e", "http://127.0.0.1", "etcd endpoint")
}
