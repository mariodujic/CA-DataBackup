package service

import (
	"CA-DataBackup/api/data"
	"CA-DataBackup/core"
	"encoding/json"
)

func GetPrayers(locale string) (data.PrayersResponse, error) {
	res, err := core.GetMiddlewareResponse("prayers", locale)
	if err != nil {
		println("Unable to process request")
	}
	var response data.PrayersResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		println("Unable to decode response")
	}
	return response, err
}
