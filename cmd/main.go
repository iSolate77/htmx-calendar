package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type Templatereader struct {
	templates *template.Template
}

func (t *Templatereader) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	tmpls, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %s", err.Error())
	}

	renderer := &Templatereader{
		templates: tmpls,
	}

	mux := http.NewServeMux()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "base.html", nil)
	})

	mux.HandleFunc("GET /toggle-display-endpoint", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "dropdownMenu.html", nil)
	})

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", mux)
}
