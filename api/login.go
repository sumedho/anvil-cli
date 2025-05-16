package api

import (
	"anvil-cli/config"
	"anvil-cli/schemas"
	"anvil-cli/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pterm/pterm"
)

func Login() {
	config := config.ReadConfig()
	passwordInput := pterm.DefaultInteractiveTextInput.WithMask("*")
	apiKey, _ := passwordInput.Show("Enter ApiKey")
	fmt.Println(apiKey)
	auth := schemas.Auth{Username: config.UserName, Apikey: apiKey}

	authJSON, err := json.Marshal(auth)
	if err != nil {
		fmt.Println(err)
	}

	url, err := url.JoinPath(config.BaseUrl, "api/2.0/auth/apiKeyLogin")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(authJSON))
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("HTTP Response Status:", resp.StatusCode, resp.Status)
	} else {
		token := schemas.Token{}
		json.Unmarshal(body, &token)
		fmt.Println("Token")
		fmt.Println(token.Token)
		fmt.Println("Expiry")
		fmt.Println(token.Expiry)
		tokenpath := utils.GetTokenCacheFilePath()
		utils.SaveJSONToFile(tokenpath, token, true)
	}

}
