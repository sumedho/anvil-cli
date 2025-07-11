package api

import (
	"anvil-cli/config"
	"anvil-cli/schemas"
	"anvil-cli/utils"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func WorkflowHandler(cCtx cli.Context) {
	schedule_id := cCtx.String("schedule-id")
	if len(schedule_id) > 0 {
		ListWorkflowOccurenceschedules(cCtx)
	} else {
		ListCatalogueWorkflowSchedules(cCtx)
	}
}

func ListCatalogueWorkflowSchedules(cCtx cli.Context) {
	outputJson := cCtx.Bool("json")
	outputCSV := cCtx.Bool("csv")
	catalogue_id := cCtx.String("id")
	config := config.ReadConfig()
	url, err := url.JoinPath(config.BaseUrl, "api/2.0/catalogues", catalogue_id, "workflowSchedules")
	if err != nil {
		fmt.Println(err)
	}

	body := utils.GETRequest(url)

	var objects schemas.WorkflowSchedules
	json.Unmarshal(body, &objects)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	fmt.Println(objects.TotalCount)
	if outputJson {
		fmt.Println(utils.JsonPrettyPrint(string(body)))
	} else if outputCSV {
		fmt.Printf("IDX,Id,Name,Timeout,SeriesName\n")
		for idx, obj := range objects.WorkflowSchedules {
			fmt.Printf("%d,%s,%s,%d,%s\n", idx+1, obj.Id, obj.Name, obj.JobTimeoutMinutes, obj.WorkflowInfo.SeriesName)
		}
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
	outputJson := cCtx.Bool("json")
	outputCSV := cCtx.Bool("csv")
	catalogue_id := cCtx.String("id")
	schedule_id := cCtx.String("schedule-id")
	config := config.ReadConfig()
	url, err := url.JoinPath(config.BaseUrl, "api/2.0/catalogues", catalogue_id, "workflowSchedules", schedule_id, "occurrences")
	if err != nil {
		fmt.Println(err)
	}

	// get the current time and back 1 year
	now := time.Now().UTC()
	then := now.Add(-24 * 365 * time.Hour).UTC()
	time_filter := schemas.WorkflowQueryTimeFilter{StartDate: then.Format(time.RFC3339), EndDate: now.Format(time.RFC3339)}

	query := schemas.WorkflowQuerySchema{ScheduledTimeFilter: time_filter, Skip: 0, Limit: 0}

	settings, err := json.Marshal(query)
	if err != nil {
		fmt.Println(err)
	}

	body := utils.POSTRequest(url, settings)
	var objects schemas.WorkflowScheduleOccurrencesSchema
	json.Unmarshal(body, &objects)
	if outputJson {
		fmt.Println(utils.JsonPrettyPrint(string(body)))
	} else if outputCSV {
		fmt.Printf("IDX,SessionId,Status,Duration,DateTime,ManualRunBy\n")
		for idx, obj := range objects.Occurrences {
			fmt.Printf("%d,%s,%s,%s,%s,%s\n", idx+1, obj.WorkflowSessionId, obj.Status, obj.Duration, obj.ScheduledTime, obj.ManualRunByUsername)
		}
	} else {
		alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
		tableData := pterm.TableData{{"IDX", "Session Id", "Status", "Duration", "DateTime", "Manual Run By"}}
		for idx, obj := range objects.Occurrences {
			data := []string{strconv.Itoa(idx + 1), obj.WorkflowSessionId, obj.Status, obj.Duration, obj.ScheduledTime, obj.ManualRunByUsername}
			tableData = append(tableData, data)
		}
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
	}
}
