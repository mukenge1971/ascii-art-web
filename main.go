package main

import (
	"fmt"
	"html/template"
	"net/http"

	"ascii-art-web/functions"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		fmt.Fprint(w, "404 not found.", http.StatusNotFound)

		return

	}

	type Ascii struct {
		Output string
	}
	s := r.FormValue("writing")
	t := r.FormValue("asciiChoice")

	var data Ascii

	if s != "" && t != "" {
		output := functions.AsciiArt(t, s)
		data = Ascii{
			Output: output,
		}
	} else if s == "" && t != "" {
		empty := "Vous n'avez rien écrit."
		data = Ascii{
			Output: empty,
		}
	}
	tpl := template.Must(template.ParseFiles("template/index.html"))
	tpl.Execute(w, data)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	
	http.HandleFunc("/", Hello)

	fmt.Printf("Starting server for testing HTTP POST http://localhost:8000...\n")
	http.ListenAndServe(":8000", nil)
}
