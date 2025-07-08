package cli

import (
	"anvil-cli/api"

	"github.com/urfave/cli/v2"
)

var workflowCli = cli.Command{
	Name:  "wf",
	Usage: "Workflow operations",
	Subcommands: []*cli.Command{
		{
			Name:  "schedule",
			Usage: "list the scheduled workflows",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Catalogue id",
					Required: true,
				},
			},
			Action: func(cCtx *cli.Context) error {
				api.ListWorkflowSchedules(*cCtx)
				return nil
			},
		},
	},
}
