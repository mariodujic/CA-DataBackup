package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func WriteJsonArray(fileName string, i interface{}) {
	jsonArray, _ := json.Marshal(i)
	file, _ := json.MarshalIndent(string(jsonArray), "", " ")
	_ = ioutil.WriteFile(fmt.Sprintf("database/%s.json", fileName), file, 0644)
}
