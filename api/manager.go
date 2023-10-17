package api

import (
	// "errors"
	"fmt"
	"net/http"
	"os"
	"io"
	"encoding/json"
	// "time"
)

type Task struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Done int `json:"done"`
}

var (
	SERVER_URL = "http://gamch1k.v6.navy:3000"
)

func GetTasks() []string {   // GET TASKS ENDPOINT: /api/get_tasks
	requestUrl := SERVER_URL + "/api/get_tasks"
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)

	if err != nil {
		fmt.Printf("Could not make http request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error on making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	data := []Task{}
	json.Unmarshal([]byte(resBody), &data)
	fmt.Println(data)
	
	final := []string{}
	for _, el := range data {
		final = append(final, el.Text)
	}
	return final
}