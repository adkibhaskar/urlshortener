package shorten

import (
	"fmt"
	"time"
	"urlshortener/constant"
	"urlshortener/database"
	"urlshortener/logger"
	"urlshortener/shortkey"
	"urlshortener/types"
)

// import (
// 	"net/http"
// )

func ShortenUrl(originalUrl string) string {

	zapLog := logger.GetLogger()

	// if r.Method != http.MethodPost {

	// 	http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)

	// 	zapLog.Error("Invalid Request Method")

	// 	return
	// }

	// originalUrl := r.FormValue("url")

	// zapLog.Info("Printing Original Url handle", zap.Any("OriginalUrl : ", originalUrl))

	// if originalUrl == "" {

	// 	http.Error(w, "Url parameter is missing", http.StatusBadRequest)
	// 	zapLog.Error("Url Parameter is missing")
	// 	return
	// }

	record, err := database.Mgr.GetUrlFromOriginalString(originalUrl, constant.UrlCollection)

	if err != nil {

		// http.Error(w, "Error Occured in recieving Record", http.StatusNotFound)

		zapLog.Error("Error Occured in receiving Record")
	}

	if record.LongUrl == "" {

		var url types.UrlDb

		Skey := shortkey.GenerateShortKey()

		ShortendUrl := fmt.Sprintf("http://localhost:8000/short/%s", Skey)

		url.CreatedAt = time.Now().Unix()
		url.ExpiredAt = time.Now().Unix()

		url.UrlCode = Skey

		url.LongUrl = originalUrl

		url.ShortUrl = ShortendUrl

		url.Count = 1

		_, err := database.Mgr.Insert(url, constant.UrlCollection)

		if err != nil {

			// http.Error(w, "Error Occured While Inserting Record", http.StatusInternalServerError)

			zapLog.Error("Error Occurred While Inserting Record")
		}

		record, err := database.Mgr.GetUrlFromOriginalString(originalUrl, constant.UrlCollection)

		if err != nil {

			// http.Error(w, "Error Occured in recieving Record", http.StatusNotFound)

			zapLog.Error("Error Occured in receiving Record")
		}

		// response := fmt.Sprintf("The Original Url is : %v \n The Shortened Url is %v \n ", record.LongUrl, record.ShortUrl)

		// zapLog.Info("The response is", zap.Any("Response : ", response))

		// fmt.Fprintln(w, response)

		return record.ShortUrl

	} else {

		count := record.Count

		newCount := count + 1

		err:= database.Mgr.UpdateCount(newCount, originalUrl)

		if err != nil {

			// http.Error(w, "Error Occured while Updating Count", http.StatusInternalServerError)
			zapLog.Error("Error Occured while Updating")
		}

		// response := fmt.Sprintf("The Original Url is : %v \n The Shortened Url is %v \n ", record.LongUrl, record.ShortUrl)
		// zapLog.Info("The response is", zap.Any("Response : ", response))
		// fmt.Fprintln(w, response)

		return record.ShortUrl

	}

}
