package main

import (
	"CA-DataBackup/api/service"
)

func main() {

	res, err := service.GetPrayers("hr")
	if err != nil {
		println("Unable to process request")
	}

	for _, prayer := range res.Data {
		print(prayer.Title)
	}
}
