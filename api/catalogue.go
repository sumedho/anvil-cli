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
	"strings"

	"github.com/pterm/pterm"
)

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func CatalogueSummary(outputJson bool, email string) {
	// GET
	config := config.ReadConfig()

	url, err := url.JoinPath(config.BaseUrl, "api/2.0/catalogueSummaries")
	if err != nil {
		fmt.Println(err)
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//Handle Error
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
	catalogs := schemas.CatalogueSummarySchema{}
	json.Unmarshal(body, &catalogs)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	// fmt.Println(catalogs.TotalCount)

	if outputJson {
		fmt.Println(jsonPrettyPrint(string(body)))
	} else {
		alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
		tableData := pterm.TableData{{"ID", "Name", "Owner"}}
		for _, cata := range catalogs.CatalogueSummaries {
			if len(email) > 0 {
				if strings.Contains(cata.Owner.Email, email) {
					data := []string{cata.Id, cata.Name, cata.Owner.Email}
					// fmt.Println(cata.Id, cata.Name)
					tableData = append(tableData, data)
				}

			} else {
				data := []string{cata.Id, cata.Name, cata.Owner.Email}
				// fmt.Println(cata.Id, cata.Name)
				tableData = append(tableData, data)
			}
		}
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
	}
}
