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
			Name:  "ls",
			Usage: "list the scheduled workflows",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "Catalogue id",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "schedule-id",
					Usage: "The schedule id",
				},
			},
			Action: func(cCtx *cli.Context) error {
				api.WorkflowHandler(*cCtx)
				return nil
			},
		},
	},
}
