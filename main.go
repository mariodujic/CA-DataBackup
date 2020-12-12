package main

import (
	"CA-DataBackup/api/service"
	"CA-DataBackup/storage"
	"fmt"
	"sync"
)

func main() {
	localeSlice := [6]string{"hr", "en", "sk", "hr-staging", "en-staging", "sk-staging"}
	featureSlice := [7]string{"quizzes", "thoughts", "information-block", "saints", "prayers", "ad-config", "user-report"}

	getMiddlewareDataAndStoreToJson(localeSlice, featureSlice)
}

func getMiddlewareDataAndStoreToJson(localeSlice [6]string, featureSlice [7]string) {
	wg := sync.WaitGroup{}
	wg.Add(len(localeSlice) * len(featureSlice))

	for _, locale := range localeSlice {
		for _, feature := range featureSlice {
			go func(feature string, locale string) {
				data := service.GetMiddlewareResponseData(feature, locale)
				storage.WriteJsonArray(fmt.Sprintf("%s-%s", feature, locale), data)
				defer wg.Done()
			}(feature, locale)
		}
	}
	wg.Wait()
}
