package main

import(
	"fmt"
	"net/http"
	"html/template"
)


var templates *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "index.html", nil)
}



func main(){

	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", indexHandler)
	fmt.Println("Http serve on!")
	defer http.ListenAndServe(":8080", nil)
}