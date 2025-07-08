package cli

import (
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
			&objectCli,
			&workflowCli,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	return 0
}
