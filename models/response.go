package models

type Response struct {
	Movies []OrganizedMovie
}

type OrganizedMovie struct {
	MovieTitle string
	Planets    []string
}
