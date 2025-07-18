package cli

import (
	"anvil-cli/api"

	"github.com/urfave/cli/v2"
)

var prefixesCli = cli.Command{
	Name:    "prefix",
	Aliases: []string{"p"},
	Usage:   "Prefix Operations",
	Subcommands: []*cli.Command{
		{
			Name:  "ls",
			Usage: "list prefixes",
			Action: func(cCtx *cli.Context) error {
				api.GetPrefixes(cCtx.Bool("json"))
				return nil
			},
		},
	},
}
