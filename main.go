package main

import (
	"CA-DataBackup/api/service"
	"CA-DataBackup/storage"
)

func main() {
	res, err := service.GetPrayers("hr")
	if err != nil {
		println("Unable to process request")
	}
	storage.WriteJsonArray(res.Data)
}
