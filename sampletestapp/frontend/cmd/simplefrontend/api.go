package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Myresponse struct {
	Date     string   `json:"timestamp"`
	Origen   string   `json:"request_origin"`
	Xff      string   `json:"request_x-forwarded-for"`
	Hostname string   `json:"response_hostname"`
	Ips      []string `json:"response_ips"`
}

func getMessage() (Myresponse, error) {
	backendURL := getEnv("BACKEND_URL", "http://localhost:8080")

	fmt.Println("Calling API...")

	client := &http.Client{}

	req, err := http.NewRequest("GET", backendURL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Myresponse
	json.Unmarshal(bodyBytes, &responseObject)

	fmt.Printf("API Response as struct %+v\n", responseObject)

	return responseObject, nil
}
