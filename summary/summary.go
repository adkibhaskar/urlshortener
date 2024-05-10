package summary

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"urlshortener/database"
	"urlshortener/logger"
	

	"go.uber.org/zap"
)

func GetSummary(w http.ResponseWriter, r *http.Request) {

	zapLog := logger.GetLogger()

	// var urls []types.UrlDb

	urls := database.Mgr.SortDocument()

	count := 0

	for count < 3 {

		urlFromDB := urls[count].LongUrl

		url, err := url.Parse(urlFromDB)

		if err != nil {

			zapLog.Error("Error Occured", zap.Any("Error : ", err))

		}

		zapLog.Info("The value of url is", zap.Any("url : ", url))

		zapLog.Info("The HostName is", zap.Any("HostName : ", url.Hostname()))

		parts := strings.Split(url.Hostname(), ".")

		zapLog.Info("The Value of parts is", zap.Any("Parts : ", parts))

		domain := parts[len(parts)-2]

		results := fmt.Sprintf("%v : %d", domain, urls[count].Count)

		zapLog.Info("The Value of Results is", zap.Any("Results : ", results))

		fmt.Fprintln(w, results)

		count++

	}

}
