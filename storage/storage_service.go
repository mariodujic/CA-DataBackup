package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var currentTime = time.Now()
var currentTimeDirectory = currentTime.Format("2006-01-02 15-04-05")
var databaseDirectoryPath = "database"

func WriteJsonArray(fileName string, data interface{}) {
	directoryExistsOrCreate()
	createFile(fileName, data)
}

func directoryExistsOrCreate() {
	if _, err := os.Stat(databaseDirectoryPath); os.IsNotExist(err) {
		_ = os.Mkdir(databaseDirectoryPath, 0755)
	}

	var currentTimeDirectoryPath = fmt.Sprintf("%s/%s", databaseDirectoryPath, currentTimeDirectory)
	if _, err := os.Stat(currentTimeDirectoryPath); os.IsNotExist(err) {
		_ = os.Mkdir(currentTimeDirectoryPath, 0755)
	}
}

func createFile(fileName string, data interface{}) {
	jsonArray, _ := json.Marshal(data)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s.json", databaseDirectoryPath, currentTimeDirectory, fileName), jsonArray, 0644)
}
