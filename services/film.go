package services

import (
	"encoding/json"
	"fmt"
)

func getFilms() ([]Film, error) {
	response, err := getResponse("films")
	if err != nil {
		return nil, err
	}

	var films []Film
	for _, raw := range response.Results {
		var film Film
		if err := json.Unmarshal(raw, &film); err == nil {
			films = append(films, film)
		} else {
			fmt.Println("Error unmarshalling film:", err)
		}
	}

	return films, nil
}
