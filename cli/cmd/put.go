package cmd

import (
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
)

var putCmd = &cli.Command{
	Name:        "put",
	Usage:       "read file then put value to etcd",
	UsageText:   "owl put [key] [file_path]",
	Description: "example: owl put /conf/test.yaml ../mock/test.yaml",
	Action: func(c *cli.Context) error {
		var key = c.Args().Get(0)
		var filePath = c.Args().Get(1)
		_ = owl.SetRemoteAddr([]string{endPoint})
		yamlFile, err := ioutil.ReadFile(filePath)

		err = owl.PutRemote(key, string(yamlFile))
		if err != nil {
			log.Panic(err)
		}
		return nil
	},
}
