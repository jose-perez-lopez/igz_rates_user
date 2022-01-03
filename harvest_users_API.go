package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func getHarvestActiveUsers() []interface{} {
	url := "https://api.harvestapp.com/v2/users?is_active=true"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Go Harvest API Sample")
	req.Header.Set("Harvest-Account-ID", os.Getenv("HarvestAccountID"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("Authorization"))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var jsonResponse map[string]interface{}
	json.Unmarshal(body, &jsonResponse)
	listSlice, _ := jsonResponse["users"].([]interface{})
	return listSlice
}
