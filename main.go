package main

import (
	"CA-DataBackup/api/service"
	"CA-DataBackup/storage"
	"fmt"
	"sync"
)

func main() {
	var localeArray [6]string
	localeArray[0] = "hr"
	localeArray[1] = "en"
	localeArray[2] = "sk"
	localeArray[3] = "hr-staging"
	localeArray[4] = "en-staging"
	localeArray[5] = "sk-staging"

	getMiddlewareDataAndStoreToJson(localeArray)
}

func getMiddlewareDataAndStoreToJson(localeArray [6]string) {
	wg := sync.WaitGroup{}
	wg.Add(len(localeArray) * 2)

	for _, locale := range localeArray {
		locale := locale
		go func() {
			data := service.GetMiddlewareResponseData("prayers", locale)
			storage.WriteJsonArray(fmt.Sprintf("prayers-%s", locale), data)
			defer wg.Done()
		}()
		go func() {
			data := service.GetMiddlewareResponseData("thoughts", locale)
			storage.WriteJsonArray(fmt.Sprintf("thoughts-%s", locale), data)
			defer wg.Done()
		}()
	}
	wg.Wait()
}
