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
	// input := "2025-05-16T15:50:00.7362658Z"
	// tokenexpiry, err := time.Parse(time.RFC3339, input)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// tnow := time.Now()
	// validminutes := tokenexpiry.Sub(tnow).Minutes()
	// fmt.Println(validminutes)
	// fmt.Println("Parsed Time:", tnow)
	// fmt.Println("Local Time:", tnow.Local())
	// token := utils.GetValidToken()
	// tokenexpiry, err := time.Parse(time.RFC3339, token.Expiry)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// tnow := time.Now()
	// validminutes := tokenexpiry.Sub(tnow).Minutes()
	// fmt.Println(validminutes)
	// api.CatalogueSummary()
}
