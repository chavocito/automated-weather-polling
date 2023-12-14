package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Data struct {
	Elevation        float64        `json:"elevation"`
	GenerationTimeMs float64        `json:"generationtime_ms"`
	Hourly           map[string]any `json:"hourly"`
}

const (
	pollInterval = time.Second * 5
	endpoint     = "https://api.open-meteo.com/v1/forecast"
)

func main() {
	ticker := time.NewTicker(pollInterval)
	fmt.Printf("Ticker Counter %s", ticker)

	for {
		<-ticker.C
		data, err := getWeatherResults(52.5, 10.23)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
	}
}

func getWeatherResults(lat, long float64) (*Data, error) {
	uri := fmt.Sprintf("%s?latitude=%f&longitude=%f&hourly=temperature_2m", endpoint, lat, long)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Timeout: time.Second * 100,
	}

	var data Data
	resp, err := client.Do(req)
	if er := json.NewDecoder(resp.Body).Decode(&data); er != nil {
		log.Fatal(er)
	}

	return &data, nil
}
