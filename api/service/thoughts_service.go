package service

import (
	"CA-DataBackup/api/data"
	"CA-DataBackup/core"
	"encoding/json"
	"fmt"
)

func GetThoughts(endPoint string, locale string) interface{} {
	res, err := core.GetMiddlewareResponse(endPoint, locale)
	if err != nil {
		println("Unable to process request")
	}
	var response data.ThoughtsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		println("Unable to decode response")
	}
	if response.Status != 200 {
		println(fmt.Sprintf("Invalid request with code %s", res.Status))
	}
 	return response.Data
}
