package model

import "encoding/json"

type Entity struct {
	Entity string          `json:"entity"`
	ID     string          `json:"id"`
	V      json.RawMessage `json:"v"`
}

type Insertable map[string]json.RawMessage
