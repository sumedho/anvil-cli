package cli

import (
	"anvil-cli/api"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func CLI() int {

	app := &cli.App{
		Name:     "anvil-cli - A cli app for managing Anvil",
		Version:  "1.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{Name: "John Doe", Email: "john.doe@gmail.com"},
		},
		Copyright: "(c) Serious Enterprise",
		HelpName:  "anvil-cli",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "json",
				Usage: "JSON output to STDOUT",
			},
		},
		Commands: []*cli.Command{

			&loginCli,
			&configureCli,
			&catalogueCli,
			&prefixesCli,
			{
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	return 0
}
