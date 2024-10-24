package services

import (
	"encoding/json"
	"fmt"

	"github.com/pchaves21/star-wars/models"
	"github.com/pchaves21/star-wars/swapi"
)

func GetPlanets() ([]models.Planet, error) {
	response, err := swapi.GetResponse("planets")
	if err != nil {
		return nil, err
	}

	var planets []models.Planet
	for _, raw := range response.Results {
		var planet models.Planet
		if err := json.Unmarshal(raw, &planet); err == nil {
			planets = append(planets, planet)
		} else {
			fmt.Println("Error unmarshalling planet:", err)
		}
	}

	return planets, nil
}
