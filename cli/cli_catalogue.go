package cli

import (
	"anvil-cli/api"

	"github.com/urfave/cli/v2"
)

var catalogueCli = cli.Command{
	Name:    "cat",
	Aliases: []string{"c"},
	Usage:   "Catalogue operations",
	Subcommands: []*cli.Command{
		{
			Name:  "query",
			Usage: "query the catalogue",
			Action: func(cCtx *cli.Context) error {
				api.CatalogueQuery()
				// fmt.Println("new task template: ", cCtx.Args().First())
				return nil
			},
		},
		{
			Name:  "ls",
			Usage: "list catalogues",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "json",
					Usage: "JSON output to STDOUT",
				},
				&cli.StringFlag{
					Name:        "email",
					Usage:       "Filter by user email",
					Destination: &email,
				},
			},
			Action: func(cCtx *cli.Context) error {
				api.CatalogueSummary(cCtx.Bool("json"), email)
				return nil
			},
		},
	},
}
