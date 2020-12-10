package storage

import (
	"encoding/json"
	"io/ioutil"
)

func WriteJsonArray(v interface{}) {
	jsonArray, _ := json.Marshal(v)
	file, _ := json.MarshalIndent(string(jsonArray), "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
