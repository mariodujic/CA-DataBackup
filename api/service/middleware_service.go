package service

import (
	"CA-DataBackup/api/data"
	"encoding/json"
	"fmt"
)

func GetMiddlewareResponseData(endPoint string, locale string) interface{} {
	res, err := data.GetMiddlewareRequest(endPoint, locale)
	if err != nil {
		println("Unable to process request")
	}
	var response data.MiddlewareResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		println("Unable to decode response for" + endPoint + "  " + locale)
	}
	if response.Status != 200 {
		println(fmt.Sprintf("Invalid request with code %s", res.Status))
	}
	return response.Data
}
