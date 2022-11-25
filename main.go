package main

import (
	//"fmt" //this is used for directly writing if needed
	"html/template"
	"net/http"
	"os"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("templates/index.html")).Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static")) //static
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", indexHandler)  //index
	http.ListenAndServe(":"+port, mux) //load
}
