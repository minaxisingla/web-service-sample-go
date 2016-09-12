package main

import (
	"net/http"
	"go.corp.yahoo.com/minaxi/web-service-sample/webservice"
)

func main() {
	http.HandleFunc("/hello", webservice.MakeFormHandler())
	http.HandleFunc("/submit", webservice.MakeViewHandler())
	http.ListenAndServe(":8080", nil)
}