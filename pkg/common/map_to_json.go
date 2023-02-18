package common

import (
	"encoding/json"
	"log"
)

func MapToJson(v map[string]interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(b)
		return ""
	}
	return string(b)
}
