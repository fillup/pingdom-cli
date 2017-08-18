package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// SummaryAverage - Holds API response for Check summery.average data
type SummaryAverage struct {
	Summary struct {
		Responsetime struct {
			From        int `json:"from"`
			To          int `json:"to"`
			Avgresponse int `json:"avgresponse"`
		} `json:"responsetime"`
		Status struct {
			Totalup      int `json:"totalup"`
			Totaldown    int `json:"totaldown"`
			Totalunknown int `json:"totalunknown"`
		} `json:"status"`
	} `json:"summary"`
}

// GetSummaryAverage - Make API call to get uptime summary information for given checkID
func GetSummaryAverage(checkID int, from int, to int) SummaryAverage {
	url := fmt.Sprintf("https://api.pingdom.com/api/2.0/summary.average/%d?includeuptime=true&from=%d&to=%d", checkID, from, to)
	// fmt.Println("Calling API: ", url)

	resp, err := CallAPI(url, "GET")
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	summaryAverage := SummaryAverage{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &summaryAverage)

	return summaryAverage
}
