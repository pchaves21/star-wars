package models

import "encoding/json"

type Result interface {
	GetType() string
}

type BaseResponse struct {
	Results []json.RawMessage `json:"results"`
}
