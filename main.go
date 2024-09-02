package main

import (
	"fmt"
	"net/http"
)

func main() {
	// this allows you to serve static files from the "static" directory,
	// making them accessible over HTTP with a URL prefix of "/static/".
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.StripPrefix("/static/", fs) creates a new http.Handler that serves
	// files from the fs handler but removes the "/static/" prefix from the URL path before passing it to fs.

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
