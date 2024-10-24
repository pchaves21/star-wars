package services

import (
	"encoding/json"
	"fmt"
)

func getPlanets() ([]Planet, error) {
	response, err := getResponse("planets")
	if err != nil {
		return nil, err
	}

	var planets []Planet
	for _, raw := range response.Results {
		var planet Planet
		if err := json.Unmarshal(raw, &planet); err == nil {
			planets = append(planets, planet)
		} else {
			fmt.Println("Error unmarshalling planet:", err)
		}
	}

	return planets, nil
}
