package models

type Planet struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (p Planet) GetType() string {
	return "Planet"
}
