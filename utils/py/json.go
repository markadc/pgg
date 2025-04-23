package py

import (
	"encoding/json"
	"github.com/markadc/pgg"
)

func Dumps(jsonMap pgg.A) (string, error) {
	bs, err := json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}
	jsonStr := string(bs)
	return jsonStr, err
}

func Loads(jsonStr string) (pgg.A, error) {
	var jsonMap pgg.A
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	return jsonMap, err
}
