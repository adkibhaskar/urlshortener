package handleform

import (
	"html/template"
	"net/http"
)



func HandleForm(w http.ResponseWriter,r *http.Request){
	
	tmpl:=template.Must(template.ParseFiles("form.html"))

	tmpl.Execute(w,nil)

	
}