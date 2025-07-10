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
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func ListCatalogueWorkflowSchedules(cCtx cli.Context) {
	outputJson := cCtx.Bool("json")
	catalogue_id := cCtx.String("id")
	config := config.ReadConfig()
	url, err := url.JoinPath(config.BaseUrl, "api/2.0/catalogues", catalogue_id, "workflowSchedules")
	if err != nil {
		fmt.Println(err)
	}

	client := http.Client{Timeout: 10 * time.Second}
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
	var objects schemas.WorkflowSchedules
	json.Unmarshal(body, &objects)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	fmt.Println(objects.TotalCount)
	if outputJson {
		fmt.Println(utils.JsonPrettyPrint(string(body)))
	} else {
		alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
		tableData := pterm.TableData{{"Schedule Id", "Name", "Timeout", "Object Name"}}
		for _, obj := range objects.WorkflowSchedules {
			data := []string{obj.Id, obj.Name, strconv.Itoa(obj.JobTimeoutMinutes), obj.WorkflowInfo.SeriesName}
			tableData = append(tableData, data)
		}
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
	}
}

func ListWorkflowOccurenceschedules(cCtx cli.Context) {
	// outputJson := cCtx.Bool("json")
	catalogue_id := cCtx.String("id")
	schedule_id := cCtx.String("schedule-id")
	config := config.ReadConfig()
	url, err := url.JoinPath(config.BaseUrl, "api/2.0/catalogues", catalogue_id, "workflowSchedules", schedule_id, "occurrences")
	if err != nil {
		fmt.Println(err)
	}

	now := time.Now().UTC()
	// subtract one year
	then := now.Add(-24 * 365 * time.Hour).UTC()
	time_filter := schemas.WorkflowQueryTimeFilter{StartDate: then.Format(time.RFC3339), EndDate: now.Format(time.RFC3339)}

	query := schemas.WorkflowQuerySchema{ScheduleTimeFilter: time_filter, Skip: 0, Limit: 0}

	settings, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
	}

	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(settings))
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
	var objects []schemas.WorkflowOccurrenceSchema
	json.Unmarshal(body, &objects)
	fmt.Println(utils.JsonPrettyPrint(string(body)))
}
