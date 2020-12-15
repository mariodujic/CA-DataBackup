package main

import (
	"CA-DataBackup/api/service"
	"CA-DataBackup/storage"
	"fmt"
	"sync"
)

func main() {
	localeArray := [6]string{"hr", "en", "sk", "hr-staging", "en-staging", "sk-staging"}
	featureArray := [6]string{"quizzes", "thoughts", "information-block", "saints", "prayers", "user-report"}

	getMiddlewareDataAndStoreToJson(localeArray, featureArray)
}

func getMiddlewareDataAndStoreToJson(localeArray [6]string, featureArray [6]string) {
	wg := sync.WaitGroup{}
	wg.Add(len(localeArray) * len(featureArray))

	for _, locale := range localeArray {
		for _, feature := range featureArray {
			go func(feature string, locale string) {
				data := service.GetMiddlewareResponseData(feature, locale)
				storage.WriteJsonArray(fmt.Sprintf("%s-%s", feature, locale), data)
				defer wg.Done()
			}(feature, locale)
		}
	}
	wg.Wait()
}
