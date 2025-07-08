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
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "limit",
					Usage: "Number of returned objects",
				},
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Catalogue id",
					Required: true,
				},
			},
			Action: func(cCtx *cli.Context) error {
				api.CatalogueQuery(*cCtx)
				return nil
			},
		},
		{
			Name:  "ls",
			Usage: "list catalogues",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "email",
					Usage: "Filter by user email",
				},
				&cli.StringFlag{
					Name:  "name",
					Usage: "Filter by catalogue name",
				},
			},
			Action: func(cCtx *cli.Context) error {
				api.CatalogueSummary(*cCtx)
				return nil
			},
		},
	},
}
