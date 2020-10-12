package cmd

import (
	"fmt"
	"github.com/gsxhnd/owl"
	"github.com/urfave/cli/v2"
	"log"
)

var getKeysCmd = &cli.Command{
	Name:      "get_keys",
	Usage:     "get keys by prefix",
	UsageText: "owl get_keys [prefix]",
	Action: func(c *cli.Context) error {
		var key = c.Args().Get(0)
		_ = owl.SetRemoteAddr([]string{endPoint})

		v, err := owl.GetRemoteKeys(key)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("value: ", v)
		return nil
	},
}
