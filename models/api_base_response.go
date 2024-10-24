package models

import "encoding/json"

type BaseResponse struct {
	Results []json.RawMessage `json:"results"`
}
