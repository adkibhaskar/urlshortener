package main

import (
	"fmt"
	"net/http"
	"urlshortener/shorten"
	"urlshortener/redirectapi"
	"urlshortener/summary"
	"urlshortener/handleform"
	
	
)

func main(){


	
	// tmpl:=template.Must(template.ParseFiles("form.html"))

	// http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {

	// 	// if r.Method != http.MethodPost{

	// 	// 	tmpl.Execute(w,nil)
	// 	// 	return
	// 	// }

	// 	tmpl.Execute(w,nil)
	// })

	
	http.HandleFunc("/",handleform.HandleForm)
    http.HandleFunc("/shorten",shorten.HandleShorten)
	http.HandleFunc("/short/",redirectapi.HandleRedirect)
    http.HandleFunc("/summary",summary.GetSummary)
	fmt.Println("URL Shortener is running on :3030")
	http.ListenAndServe(":3030",nil)
	
}