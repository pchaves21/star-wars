package services

import (
	"encoding/json"
	"fmt"

	"github.com/pchaves21/star-wars/models"
	"github.com/pchaves21/star-wars/swapi"
)

func GetFilms() ([]models.Film, error) {
	response, err := swapi.GetResponse("films")
	if err != nil {
		return nil, err
	}

	var films []models.Film
	for _, raw := range response.Results {
		var film models.Film
		if err := json.Unmarshal(raw, &film); err == nil {
			films = append(films, film)
		} else {
			fmt.Println("Error unmarshalling film:", err)
		}
	}

	return films, nil
}
