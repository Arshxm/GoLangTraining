package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type APIResponse struct {
	Status string               `json:"status"`
	Stats  map[string]PairStats `json:"stats"`
}

type PairStats struct {
	Latest string `json:"latest"`
}

func main() {
	response, err := GetExchangeRate("btc", "usdt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

func GetExchangeRate(source, destination string) (string, error) {
	source = strings.TrimSpace(strings.ToLower(source))
	destination = strings.TrimSpace(strings.ToLower(destination))
	if source == "" {
		return "", errors.New("source is required")
	}
	if destination == "" {
		destination = "rls"
	}
	url := fmt.Sprintf("http://localhost:4001/rates?srcCurrency=%s&dstCurrency=%s", source, destination)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var response APIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return statsDecoder(response, source, destination)
}

func statsDecoder(response APIResponse, source, destination string) (string, error) {
	if response.Stats == nil {
		return "", errors.New("stats is missing in response")
	}
	key := fmt.Sprintf("%s-%s", source, destination)
	pairStats, ok := response.Stats[key]
	if !ok {
		return "", fmt.Errorf("pair %s not found in stats", key)
	}
	latest := strings.TrimSpace(pairStats.Latest)
	if latest == "" {
		return "", errors.New("latest is empty in stats")
	}
	return latest, nil
}
