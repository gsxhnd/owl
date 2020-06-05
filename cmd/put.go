package cmd

import (
	"errors"
	"github.com/gsxhnd/owl/backend"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
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
			log.Panic(err)
		}
		yamlFile, err := ioutil.ReadFile(filePath)

		err = conn.Put(key, string(yamlFile))
		if err != nil {
			log.Panic(err)
		}
		return nil
	},
}

func init() {
	putComd.PersistentFlags().StringVarP(&endPoint, "endpoint", "e", "http://127.0.0.1", "etcd endpoint")
}
