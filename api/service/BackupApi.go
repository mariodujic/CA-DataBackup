package service

import (
	"CA-DataBackup/api/data"
	"CA-DataBackup/constants"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMiddlewareResponse(locale string) (*data.PrayersResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/prayers/%s", constants.BaseUrl, locale), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(constants.HeaderSecretKey, constants.HeaderSecretValue)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var response data.PrayersResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
