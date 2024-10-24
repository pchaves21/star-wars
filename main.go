package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/pchaves21/star-wars/models"
	"github.com/pchaves21/star-wars/services"
)

func createResponse() (models.Response, error) {
	response := models.Response{}
	planetMap := make(map[string]string)

	planets, err := services.GetPlanets()
	if err != nil {
		return response, err
	}

	for _, planet := range planets {
		planetMap[planet.Url] = planet.Name
	}

	films, err := services.GetFilms()
	if err != nil {
		return response, err
	}

	sort.Slice(films, func(i, j int) bool {
		return films[i].ReleaseDate < films[j].ReleaseDate
	})

	for _, film := range films {
		var planetNames []string

		for _, planetURL := range film.Planets {
			if name, exists := planetMap[planetURL]; exists {
				planetNames = append(planetNames, name)
			}
		}

		sort.Strings(planetNames)

		organizedMovie := models.OrganizedMovie{
			MovieTitle: film.Title,
			Planets:    planetNames,
		}
		response.Movies = append(response.Movies, organizedMovie)
	}

	return response, nil
}

func main() {

	response, err := createResponse()
	if err != nil {
		fmt.Println("Error creating response:", err)
		return
	}

	prettyJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))

}
