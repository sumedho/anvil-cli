package utils

import (
	"anvil-cli/schemas"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	ConfigFileDir      = ".anvil-cli"
	ConfigFileName     = "config.json"
	TokenCacheFileName = ".cache.json"
)

// Return platform independant HOME path
func getHomeDir() string {
	home, _ := os.UserHomeDir()
	if home == "" && runtime.GOOS != "windows" {
		if u, err := user.Current(); err == nil {
			return u.HomeDir
		}
	}
	return home
}

// Return the configuration file path
func GetAnvilConfigFilePath() string {
	homedir := getHomeDir()
	configpath := filepath.Join(homedir, ConfigFileDir, ConfigFileName)
	return configpath
}

// Return the configuration directory
func GetAnvilDir() string {
	homedir := getHomeDir()
	configpath := filepath.Join(homedir, ConfigFileDir)
	return configpath
}

// Get the token cache file path
func GetTokenCacheFilePath() string {
	homedir := getHomeDir()
	tokenpath := filepath.Join(homedir, ConfigFileDir, TokenCacheFileName)
	return tokenpath
}

// Create a new directory
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

// Get the valid token
func GetValidToken() schemas.Token {
	cachepath := GetTokenCacheFilePath()
	file, err := os.OpenFile(cachepath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	config := schemas.Token{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return config
}

// Save a JSON file to disk
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

// Pretty print json
func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

// Parse an elapsed time for workflow schedules
func ParseMinutesTimeString(timetaken string) float64 {
	parts := strings.Split(timetaken, ":")
	minutes, err := strconv.ParseFloat(parts[0], 3)
	if err != nil {
		log.Fatal("Parse of minutes failed:", err)
	}
	seconds, err := strconv.ParseFloat(parts[1], 3)
	if err != nil {
		log.Fatal("Parse of seconds failed:", err)
	}
	return minutes + seconds
}

// GET
func GETRequest(url_path string) []byte {
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url_path, nil)
	if err != nil {
		log.Fatal("Request failed", err)
	}
	cache := GetValidToken()
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + cache.Token},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

// POST
func POSTRequest(url_path string, json_data []byte) []byte {
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url_path, bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal("Request failed", err)
	}
	cache := GetValidToken()
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + cache.Token},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}
