package cli

import (
	"anvil-cli/api"

	"github.com/urfave/cli/v2"
)

var loginCli = cli.Command{
	Name:    "login",
	Aliases: []string{"l"},
	Usage:   "login to Anvil",
	Action: func(cCtx *cli.Context) error {
		api.Login()
		return nil
	},
}
