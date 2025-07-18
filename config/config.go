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

// Create the directory to store the configuration
func CreateConfigDir() {
	newpath := utils.GetAnvilDir()
	if err := utils.MakeDir(newpath); err != nil {
		fmt.Println("Unable to create configuration directory: " + err.Error())
		os.Exit(1)
	}
}

// Create configuration file
func CreateConfigOrOpen() {
	configpath := utils.GetAnvilConfigFilePath()
	_, err := os.Stat(configpath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("no config, creating empty config.")
		bookmarks := make([]schemas.Bookmark, 0)
		config := schemas.Configuration{UserName: "", BaseUrl: "", Bookmarks: bookmarks}
		utils.SaveJSONToFile(configpath, config, true)
	}
}

// Set the configuration
func SetConfig() {
	configpath := utils.GetAnvilConfigFilePath()
	userName, _ := pterm.DefaultInteractiveTextInput.Show("Username")
	baseUrl, _ := pterm.DefaultInteractiveTextInput.Show("BaseUrl")
	config := ReadConfig()
	config.UserName = userName
	config.BaseUrl = baseUrl
	utils.SaveJSONToFile(configpath, config, true)
	fmt.Println("configuration set")
}

// Read config from disk
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
