package main

import (
	"anvil-cli/cli"
	"anvil-cli/config"
	"os"
)

func main() {
	config.CreateConfigDir()
	config.CreateConfigOrOpen()
	os.Exit(cli.CLI())
}
