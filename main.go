package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })
	

	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	//r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	//r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	//r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	/** gorilla mux kullanıldığında çalışmıyor
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
	**/

	//Varsayılan nil değeri alırken gorilla/mux kulllanıldığında r parametresini alır
    http.ListenAndServe(":80", r)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    title := vars["title"]
	fmt.Fprintf(w, "You've requested the book: %s\n", title)
}