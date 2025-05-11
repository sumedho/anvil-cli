package main

import "anvil-cli/config"

type ConfigFile struct {
	UserName string `json:"UserName"`
	BaseUrl  string `json:"BaseUrl"`
}

func main() {
	//get home dir
	// dirname, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// configpath := filepath.Join(dirname, ".anvil-cli", "config.json")
	// // create anvil-cli dir in homedir
	// newpath := filepath.Join(dirname, ".anvil-cli")
	// // err = os.MkdirAll(newpath, os.ModePerm)
	// if err := ensureDir(newpath); err != nil {
	// 	fmt.Println("Directory creation failed with error: " + err.Error())
	// 	os.Exit(1)
	// }
	// fmt.Println(dirname)
	// os.Exit(cli.CLI())
	// Create an interactive text input with a mask for password input
	config.CreateConfigDir()
	config.CreateConfigOrOpen()
	// passwordInput := pterm.DefaultInteractiveTextInput.WithMask("*")
	// result, _ := passwordInput.Show("Enter your password")
	// logger := pterm.DefaultLogger
	// logger.Info("Password received", logger.Args("password", result))
}
