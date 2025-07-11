package api

import (
	"anvil-cli/config"
	"anvil-cli/schemas"
	"anvil-cli/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pterm/pterm"
)

func Login() {
	config := config.ReadConfig()
	passwordInput := pterm.DefaultInteractiveTextInput.WithMask("*")
	apiKey, _ := passwordInput.Show("Enter ApiKey")
	auth := schemas.Auth{Username: config.UserName, Apikey: apiKey}

	authJSON, err := json.Marshal(auth)
	if err != nil {
		fmt.Println(err)
	}

	url, err := url.JoinPath(config.BaseUrl, "api/2.0/auth/apiKeyLogin")
	if err != nil {
		fmt.Println(err)
	}

	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(authJSON))
	if err != nil {
		fmt.Println("Request failed", err)
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	token := schemas.Token{}
	json.Unmarshal(body, &token)
	fmt.Println("Token")
	fmt.Println(token.Token)
	fmt.Println("Expiry")
	fmt.Println(token.Expiry)
	tokenpath := utils.GetTokenCacheFilePath()
	utils.SaveJSONToFile(tokenpath, token, true)
}
