package models

import "encoding/json"

type Response struct {
	Count int `json:count`
	Result *json.RawMessage `json:result`
}