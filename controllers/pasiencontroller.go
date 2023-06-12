package controllers

import (
	"html/template"
	"net/http"
)

func Index(response http.ResponseWriter, request *http.Request) {
	templ, err := template.ParseFiles("views/pasien/index.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(response, nil)
}

func Add(response http.ResponseWriter, request *http.Request) {
	templ, err := template.ParseFiles("views/pasien/add.html")
	if err != nil {
		panic(err)
	}
	templ.Execute(response, nil)

}

func Edit(response http.ResponseWriter, request *http.Request) {

}

func Delete(response http.ResponseWriter, request *http.Request) {

}
