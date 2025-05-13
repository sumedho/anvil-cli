package config

import (
	"anvil-cli/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

const ANVIL_CONFIG_DIR = ".anvil-cli"

type Configuration struct {
	UserName string `json:"UserName"`
	BaseUrl  string `json:"BaseUrl"`
}

func CreateConfigDir() {
	// get home dir
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	// create anvil-cli dir in homedir
	newpath := filepath.Join(dirname, ANVIL_CONFIG_DIR)
	if err := utils.MakeDir(newpath); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
}

func CreateConfigOrOpen() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configpath := filepath.Join(dirname, ANVIL_CONFIG_DIR, "config.json")
	_, err = os.Stat(configpath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("no config, creating empty config.")
		config := Configuration{UserName: "", BaseUrl: ""}
		utils.SaveJSONToFile(configpath, config, true)
	}

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
	fmt.Println(config.UserName)
	fmt.Println(config.BaseUrl)
	if (config.UserName == "") || (config.BaseUrl == "") {
		SetConfig()
	}

}

func SetConfig() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configpath := filepath.Join(dirname, ANVIL_CONFIG_DIR, "config.json")

	userName, _ := pterm.DefaultInteractiveTextInput.Show("Username")
	baseUrl, _ := pterm.DefaultInteractiveTextInput.Show("BaseUrl")
	fmt.Println(userName)
	fmt.Println(baseUrl)
	config := Configuration{UserName: userName, BaseUrl: baseUrl}
	utils.SaveJSONToFile(configpath, config, true)

}
