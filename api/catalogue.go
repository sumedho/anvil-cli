package api

import (
	"anvil-cli/config"
	"anvil-cli/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pterm/pterm"
)

type OwnerSchema struct {
	Id        string `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	GivenName string `json:"givenName"`
	LastName  string `json:"lastName"`
}

type CatalogueSchema struct {
	Id                                 string      `json:"id"`
	Name                               string      `json:"name"`
	Description                        string      `json:"description"`
	Version                            int         `json:"version"`
	IsOfficial                         bool        `json:"isOfficial"`
	IsDraft                            bool        `json:"isDraft"`
	Owner                              OwnerSchema `json:"owner"`
	OwnerUserName                      string      `json:"ownerUsername"`
	CreatedBy                          string      `json:"createdBy"`
	DateCreated                        string      `json:"dateCreated"`
	Prefix                             string      `json:"prefix"`
	CustomSeriesNaming                 string      `json:"customSeriesNaming"`
	CustomObjectNaming                 string      `json:"customObjectNaming"`
	ApprovalProcessProcessFLowSeriesId string      `json:"approvalProcessProcessFlowSeriesId"`
	ObjectClasses                      []string    `json:"objectClasses"`
}

type CatalogueSummarySchema struct {
	Limit              int               `json:"limit"`
	Offset             int               `json:"offset"`
	CatalogueSummaries []CatalogueSchema `json:"catalogueSummaries"`
	TotalCount         int               `json:"totalCount"`
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func CatalogueSummary(outputJson bool) {
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
	catalogs := CatalogueSummarySchema{}
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
			data := []string{cata.Id, cata.Name, cata.Owner.Email}
			// fmt.Println(cata.Id, cata.Name)
			tableData = append(tableData, data)
		}
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
	}

}
