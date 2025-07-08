package cli

import (
	"anvil-cli/api"

	"github.com/urfave/cli/v2"
)

var objectCli = cli.Command{
	Name:    "obj",
	Aliases: []string{"o"},
	Usage:   "Object operations",
	Subcommands: []*cli.Command{
		{
			Name:  "info",
			Usage: "Object information",
			Action: func(cCtx *cli.Context) error {
				api.ObjectInfo()
				// fmt.Println("new task template: ", cCtx.Args().First())
				return nil
			},
		},
		{
			Name:  "attrs",
			Usage: "List object attributes",
			Action: func(cCtx *cli.Context) error {
				api.ObjectAttributes()
				//fmt.Println("removed task template: ", cCtx.Args().First())
				return nil
			},
		},
	},
}
