package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// create page data structure
type FormData struct {
	Name  string
	Email string
}

// create a variable templates of type pointer to templates to store all html files
var templates *template.Template

// create a function that parses all the html files in the template folder into the just created template variable
func init() {
	// template.ParseGlob: Parses all templates in the specified directory.
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

// this function handles requests for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

// this function handles requests for the form page
func formHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "form.html", nil)
}

// this function handles requests for the form page
func submitHandler(w http.ResponseWriter, r *http.Request) {
	// checks if the http method used in the request is POST
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/form", http.StatusSeeOther)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "unable to parse form", http.StatusBadRequest)
	}

	data := FormData{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}

	templates.ExecuteTemplate(w, "result.html", data)
}

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
