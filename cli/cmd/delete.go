package cmd

import (
	"log"

	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
)

var deleteCmd = &cli.Command{
	Name:        "delete",
	Usage:       "delete value by key",
	UsageText:   "owl get [key]",
	Description: "the [key] what you want value at the etcd",
	Action: func(c *cli.Context) error {
		var key = c.Args().Get(0)
		_ = owl.SetRemoteAddr([]string{endPoint})
		err := owl.DeleteRemote(key)
		if err != nil {
			log.Panic(err)
		}
		return nil
	},
}
