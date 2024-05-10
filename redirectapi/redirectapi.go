package redirectapi

import (
	"net/http"

	"urlshortener/constant"
	"urlshortener/database"
	"urlshortener/logger"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func GoToOriginalPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	Shortkey := vars["id"]

	if Shortkey == ""{

		http.Error(w,"Key is Missing",http.StatusNotFound)
	}

	zapLog := logger.GetLogger()

	zapLog.Info("The value of shortkey is", zap.Any("shortKey : ", Shortkey))

	record, err := database.Mgr.GetUrlFromCode(Shortkey, constant.UrlCollection)

	if err != nil{

		http.Error(w,"Error Occured while fetching shortkey",http.StatusNotFound)
		zapLog.Error("Error Occured while fetching shortkey")
	}
    
	zapLog.Info("The Value of OriginalUrl is",zap.Any("OriginalUrl : ",record.LongUrl))

	http.Redirect(w,r,record.LongUrl,http.StatusFound)

}
