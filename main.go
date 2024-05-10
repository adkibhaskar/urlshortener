package main

import (
	"fmt"
	"sync"

	"net/http"
	"urlshortener/constant"
	"urlshortener/database"
	"urlshortener/logger"
	"urlshortener/shorten"
	// "urlshortener/shorten"
)

func main() {

	zapLog := logger.GetLogger()

	// err := godotenv.Load(".env")

	// if err != nil {

	// 	zapLog.Error("Unable to get Env Variables")
	// 	panic((err))

	// }

	// UrlsToShorten:=make(chan string)

	// ShortenUrls:=make(chan string)

	// database.ConnectDb()

	// zapLog.Info("Successfully Connected to DB")

	// router := mux.NewRouter()

	// router.HandleFunc(constant.UrlShortenPath, shorten.HandleShorten)

	// router.HandleFunc(constant.RedirectUrlPath, redirectapi.GoToOriginalPage)
	// router.HandleFunc(constant.Summary, summary.GetSummary)

	// http.Handle("/", router)

	// // zapLog.Info("URL Shortener is running on :8000")

	// http.ListenAndServe(constant.BaseUrl, nil)

	// go shorten.UrlShortener(UrlsToShorten,ShortenUrls)

	database.ConnectDb()

	urlsToShorten := make(chan string)
	shortenUrl := make(chan string)

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go urlShortener(urlsToShorten, shortenUrl, wg)

	http.HandleFunc(constant.UrlShortenPath, func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {

			http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)

			zapLog.Error("Invalid Request Method")

			return

		}

		url := r.FormValue("url")

		urlsToShorten <- url

		shortUrl := <-shortenUrl

		fmt.Fprintf(w, "Shortened Url : %s", shortUrl)
	})

	http.ListenAndServe(constant.BaseUrl, nil)

	wg.Wait()

}

func urlShortener(input chan string, output chan string, wg *sync.WaitGroup) {

	url := <-input

	shortURL:= shorten.ShortenUrl(url)

	output <- shortURL

	wg.Done()

}
