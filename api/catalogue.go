package api

import (
	"anvil-cli/config"
	"anvil-cli/schemas"
	"anvil-cli/utils"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

type Query struct {
	OrderByAsc bool `json:"orderByAsc"`
	Limit      int  `json:"limit"`
}

// Query a single catalogue
func CatalogueQuery(cCtx cli.Context) {
	outputJson := cCtx.Bool("json")
	limit := cCtx.Int("limit")
	catalogue_id := cCtx.String("id")

	config := config.ReadConfig()
	url, err := url.JoinPath(config.BaseUrl, "api/2.0/catalogues", catalogue_id, "query")
	if err != nil {
		fmt.Println(err)
	}

	query := Query{OrderByAsc: true, Limit: limit}
	data, _ := json.Marshal(query)

	body := utils.POSTRequest(url, data)

	var objects []schemas.CatalogueQueryObject
	json.Unmarshal(body, &objects)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	if outputJson {
		fmt.Println(utils.JsonPrettyPrint(string(body)))
	} else {
		alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
		tableData := pterm.TableData{{"ID", "Name"}}
		for _, obj := range objects {

			data := []string{obj.Id, obj.Name}
			tableData = append(tableData, data)
		}
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
	}
}

// Get summary of all catalogues
func CatalogueSummary(cCtx cli.Context) {
	email := cCtx.String("email")
	name := cCtx.String("name")
	outputJson := cCtx.Bool("json")

	config := config.ReadConfig()

	url, err := url.JoinPath(config.BaseUrl, "api/2.0/catalogueSummaries")
	if err != nil {
		fmt.Println(err)
	}

	body := utils.GETRequest(url)

	catalogs := schemas.CatalogueSummarySchema{}
	json.Unmarshal(body, &catalogs)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	// fmt.Println(catalogs.TotalCount)
	if outputJson {
		fmt.Println(utils.JsonPrettyPrint(string(body)))
	} else {
		alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
		tableData := pterm.TableData{{"ID", "Name", "Owner"}}
		catalogueSummary := catalogs.CatalogueSummaries
		if len(email) > 0 {
			catalogueSummary = filterEmail(email, catalogueSummary)
		}

		if len(name) > 0 {
			catalogueSummary = filterName(name, catalogueSummary)
		}

		for _, cata := range catalogueSummary {
			data := []string{cata.Id, cata.Name, cata.Owner.Email}
			tableData = append(tableData, data)
		}
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
	}
}

// Filter slice by partial email string
func filterEmail(email string, data []schemas.CatalogueSchema) []schemas.CatalogueSchema {
	var filteredData []schemas.CatalogueSchema
	for _, cata := range data {
		if strings.Contains(cata.Owner.Email, email) {
			filteredData = append(filteredData, cata)
		}
	}
	return filteredData
}

// Filter slice by partial catalogue name
func filterName(name string, data []schemas.CatalogueSchema) []schemas.CatalogueSchema {
	var filteredData []schemas.CatalogueSchema
	for _, cata := range data {
		if strings.Contains(cata.Name, name) {
			filteredData = append(filteredData, cata)
		}
	}
	return filteredData
}
