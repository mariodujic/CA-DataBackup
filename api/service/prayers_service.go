package service

import (
	"CA-DataBackup/api/data"
	"CA-DataBackup/core"
	"encoding/json"
	"fmt"
)

func GetPrayers(locale string) interface{} {
	res, err := core.GetMiddlewareResponse("prayers", locale)
	if err != nil {
		println("Unable to process request")
	}
	var response data.PrayersResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		println("Unable to decode response")
	}
	if response.Status != 200 {
		println(fmt.Sprintf("Invalid request with code %s", res.Status))
	}
 	return response.Data
}
