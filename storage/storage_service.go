package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var databaseDirectoryPath = "database"

func WriteJsonArray(fileName string, data interface{}) {
	directoryExistsOrCreate()
	createFile(fileName, data)
}

func directoryExistsOrCreate() {
	if _, err := os.Stat(databaseDirectoryPath); os.IsNotExist(err) {
		_ = os.Mkdir(databaseDirectoryPath, 0755)
	}
}

func createFile(fileName string, data interface{}) {
	jsonArray, _ := json.Marshal(data)
	file, _ := json.MarshalIndent(string(jsonArray), "", " ")
	_ = ioutil.WriteFile(fmt.Sprintf("%s/%s.json", databaseDirectoryPath, fileName), file, 0644)
}
