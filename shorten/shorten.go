package shorten

import (
	"fmt"
	"net/http"
	"urlshortener/shortkey"
	
)

var Urls=make(map[string]string)
var Urls1=make(map[string]int)

func checkValue(originalUrl string,Urls map[string]string)(bool,string,string){

	for key,value:=range Urls{

		 if value == originalUrl{

			  return true,key,value
		 }
	}

	return false,"",""
}

func HandleShorten(w http.ResponseWriter,r *http.Request){

	if r.Method != http.MethodPost{

		http.Error(w,"Invalid Request Method",http.StatusMethodNotAllowed)
		return 
	}

	
	originalUrl:=r.FormValue("url")

	if originalUrl == ""{

		http.Error(w,"Url parameter is missing",http.StatusBadRequest)
		return 
	}

	check,key,value:=checkValue(originalUrl,Urls)

	 if check==true{
		
           Urls1[value]++
			w.Header().Set("Content-Type","text/html")
		    fmt.Fprint(w, `
				<!DOCTYPE html>
				<html>
				<head>
					<title>URL Shortener</title>
				</head>
				<body>
					<h2>URL Shortener</h2>
					<p>Original URL: `,originalUrl,`</p>
					<p>Shortened URL:<a href="http://localhost:3030/short/`, key, `">http://localhost:3030/short/`, key, `</a></p>
				</body>
				</html>
		`)

          
	 }else{

		shortkey:=shortkey.GenerateShortKey()
	
		Urls[shortkey]=originalUrl
		
		// fmt.Println("The Value of Map is : ",Urls)
	
		shortendUrl:=fmt.Sprintf("http://localhost:3030/short/%s",shortkey)
	
		w.Header().Set("Content-Type","text/html")
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>URL Shortener</title>
			</head>
			<body>
				<h2>URL Shortener</h2>
				<p>Original URL: `,originalUrl,`</p>
				<p>Shortened URL: <a href="`, shortendUrl, `">`, shortendUrl, `</a></p>
			</body>
			</html>
		`)

	 }
}