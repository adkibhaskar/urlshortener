package redirectapi

import (
	"net/http"
	"strings"
	"urlshortener/shorten"
)
func HandleRedirect(w http.ResponseWriter,r *http.Request){

	shortkey:=strings.TrimPrefix(r.URL.Path,"/short/")

	if shortkey == ""{

		http.Error(w,"Shortened Key is Missing",http.StatusBadRequest)
		return
	}

	originalUrl,found:=shorten.Urls[shortkey]

	if !found{

		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w,r,originalUrl,http.StatusMovedPermanently)


}
