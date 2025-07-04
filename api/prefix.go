package api

import (
	"anvil-cli/config"
	"anvil-cli/schemas"
	"anvil-cli/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pterm/pterm"
)

func GetPrefixes(outputJson bool) {
	// GET
	config := config.ReadConfig()

	url, err := url.JoinPath(config.BaseUrl, "api/2.0/cataloguePrefixes")
	if err != nil {
		fmt.Println(err)
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Request failed", err)
	}

	cache := utils.GetValidToken()
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + cache.Token},
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	prefixes := schemas.PrefixSchema{}
	json.Unmarshal(body, &prefixes)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	if outputJson {
		fmt.Println(jsonPrettyPrint(string(body)))
	} else {
		alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
		tableData := pterm.TableData{{"PrefixName", "Label", "Description", "CreatedBy"}}
		for _, prefix := range prefixes.CataloguePrefixes {
			data := []string{prefix.Prefix, prefix.Label, prefix.Description, prefix.CreatedBy}
			tableData = append(tableData, data)
		}
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
	}
}
