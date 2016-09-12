package webservice

import (
	"fmt"
	"net/http"
	//"strings"
	"html/template"
	//"io/ioutil"
)

func MakeFormHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("user-form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	}
}

func MakeViewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("Name")
	fmt.Fprintf(w, "Hello " + name + "!")
	}
}