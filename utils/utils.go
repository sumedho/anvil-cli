package utils

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const (
	ConfigFileDir  = ".anvil-cli"
	ConfigFileName = ".config.json"
)

func getHomeDir() string {
	home, _ := os.UserHomeDir()
	if home == "" && runtime.GOOS != "windows" {
		if u, err := user.Current(); err == nil {
			return u.HomeDir
		}
	}
	return home
}

func GetAnvilConfigFilePath() string {
	homedir := getHomeDir()
	configpath := filepath.Join(homedir, ConfigFileDir, ConfigFileName)
	return configpath
}

func GetAnvilDir() string {
	homedir := getHomeDir()
	configpath := filepath.Join(homedir, ConfigFileDir)
	return configpath
}

func MakeDir(dirName string) error {
	err := os.Mkdir(dirName, 0777)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

func SaveJSONToFile(filePath string, data interface{}, indent bool) error {
	// Marshal the data into JSON format.
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if indent {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
	return nil
}
