package config

import (
	"anvil-cli/schemas"
	"anvil-cli/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

func CreateConfigDir() {
	newpath := utils.GetAnvilDir()
	if err := utils.MakeDir(newpath); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
}

func CreateConfigOrOpen() {

	configpath := utils.GetAnvilConfigFilePath()
	_, err := os.Stat(configpath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("no config, creating empty config.")
		config := schemas.Configuration{UserName: "", BaseUrl: ""}
		utils.SaveJSONToFile(configpath, config, true)
	}

	// file, err := os.OpenFile(configpath, os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// config := Configuration{}
	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&config)
	// if err != nil {
	// 	fmt.Println("Error decoding JSON:", err)
	// }
}

func SetConfig() {
	configpath := utils.GetAnvilConfigFilePath()
	userName, _ := pterm.DefaultInteractiveTextInput.Show("Username")
	baseUrl, _ := pterm.DefaultInteractiveTextInput.Show("BaseUrl")
	config := schemas.Configuration{UserName: userName, BaseUrl: baseUrl}
	utils.SaveJSONToFile(configpath, config, true)
	fmt.Println("configuration set")
}

func ReadConfig() schemas.Configuration {
	configpath := utils.GetAnvilConfigFilePath()
	file, err := os.OpenFile(configpath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	config := schemas.Configuration{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return config
}
