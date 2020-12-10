package main

import (
	"CA-DataBackup/api/service"
)

func main() {
	status, err := service.GetMiddlewareResponse("hr")
	if err != nil {}
	for _, prayer := range status.Data {
		print(prayer.Title)
	}
}
