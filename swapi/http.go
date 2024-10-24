package swapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const API_URL = "https://swapi.dev/api/"

func getResponse(extension string) (*BaseResponse, error) {
	resp, err := http.Get(API_URL + extension)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received status code", resp.StatusCode)
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	var response BaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
