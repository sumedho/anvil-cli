package cli

import (
	"anvil-cli/config"

	"github.com/urfave/cli/v2"
)

var configureCli = cli.Command{
	Name: "configure",
	//Aliases: []string{"l"},
	Usage: "set configuration",
	Action: func(cCtx *cli.Context) error {
		config.SetConfig()
		return nil
	},
}
