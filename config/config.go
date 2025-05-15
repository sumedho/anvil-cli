package config

import (
	"anvil-cli/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

type Configuration struct {
	UserName string `json:"Username"`
	BaseUrl  string `json:"Baseurl"`
}

func CreateConfigDir() {
	// // get home dir
	// dirname, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // create anvil-cli dir in homedir
	// newpath := filepath.Join(dirname, ANVIL_CONFIG_DIR)
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
		config := Configuration{UserName: "", BaseUrl: ""}
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
	config := Configuration{UserName: userName, BaseUrl: baseUrl}
	utils.SaveJSONToFile(configpath, config, true)
	fmt.Println("configuration set")
}

func ReadConfig() Configuration {
	configpath := utils.GetAnvilConfigFilePath()
	file, err := os.OpenFile(configpath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	config := Configuration{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return config
}
