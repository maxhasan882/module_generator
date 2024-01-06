package json

import "encoding/json"

func Cast(src, target interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, target)
}
