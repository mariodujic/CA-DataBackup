package data

import (
	"CA-DataBackup/constants"
	"fmt"
	"net/http"
)

func GetMiddlewareRequest(endpoint string, locale string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", constants.BaseUrl, endpoint, locale), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(constants.HeaderSecretKey, constants.HeaderSecretValue)

	client := &http.Client{}
	return client.Do(req)
}
