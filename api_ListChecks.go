package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ChecksList - Holds API response with list of Uptime Checks
type ChecksList struct {
	Checks []struct {
		Hostname         string `json:"hostname"`
		ID               int    `json:"id"`
		Lasterrortime    int    `json:"lasterrortime"`
		Lastresponsetime int    `json:"lastresponsetime"`
		Lasttesttime     int    `json:"lasttesttime"`
		Name             string `json:"name"`
		Resolution       int    `json:"resolution"`
		Status           string `json:"status"`
		Type             string `json:"type"`
		Tags             []struct {
			Name  string `json:"name"`
			Type  string `json:"type"`
			Count int    `json:"count"`
		} `json:"tags"`
	} `json:"checks"`
}

// ListChecks - Make API call to get list of Uptime checks
func ListChecks(tags string) ChecksList {
	url := fmt.Sprintf("https://api.pingdom.com/api/2.0/checks?tags=%v&include_tags=true", tags)

	resp, err := CallAPI(url, "GET")
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	checksList := ChecksList{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &checksList)

	return checksList
}
