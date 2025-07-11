package api

import (
	"anvil-cli/config"
	"anvil-cli/schemas"
	"anvil-cli/utils"
	"encoding/json"
	"fmt"
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
	body := utils.GETRequest(url)

	prefixes := schemas.PrefixSchema{}
	json.Unmarshal(body, &prefixes)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	if outputJson {
		fmt.Println(utils.JsonPrettyPrint(string(body)))
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
