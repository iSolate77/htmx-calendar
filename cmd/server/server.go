package main

import (
	"html/template"
	"net/http"
)

func StartServer() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("../../web/templates/base.html")
  t.Execute(w, nil)
}

func main() {
  
}
