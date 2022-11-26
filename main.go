package main

import (
	//"fmt" //this is used for directly writing if needed
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

func main() {

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))

	r := mux.NewRouter()

    r.PathPrefix("/static/").Handler(fs)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("templates/index.html")).Execute(w, nil)
	})

	http.ListenAndServe(":8080", r) //load
}
