package main

import (
	"anvil-cli/cli"
	"anvil-cli/config"
	"os"
)

func main() {
	// fmt.Println(utils.GetHomeDir())
	config.CreateConfigDir()
	config.CreateConfigOrOpen()
	os.Exit(cli.CLI())
}
