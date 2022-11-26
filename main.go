package main

import (
	"fmt" //this is used for directly writing if needed
	//"html/template"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	//mux := http.NewServeMux()
	//fs := http.FileServer(http.Dir("static")) //static
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))

	//mux.HandleFunc("/", indexHandler)  //index
	r := mux.NewRouter()
    r.PathPrefix("/static/").Handler(fs)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>This is the homepage. Try /hello and /hello/Sammy\n</h1>")
		//template.Must(template.ParseFiles("templates/index.html")).Execute(w, nil)
	})
	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>This is the homepage. Try /hello and /hello/Sammy\n</h1>")
		//template.Must(template.ParseFiles("templates/index.html")).Execute(w, nil)
	})
	http.ListenAndServe(":"+port, r) //load
}
