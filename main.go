package main

import (
	"net/http"

	"github.com/MeizalunaWulandari/golang-crud/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/pasien", controllers.Index)
	http.HandleFunc("/pasien/index", controllers.Index)
	http.HandleFunc("/pasien/add", controllers.Add)
	http.HandleFunc("/pasien/edit", controllers.Edit)
	http.HandleFunc("/pasien/delete", controllers.Delete)

	http.ListenAndServe(":8000", nil)
}
