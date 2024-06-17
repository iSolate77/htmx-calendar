package main

import (
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type Templatereader struct {
	templates *template.Template
}

func (t *Templatereader) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	tmpls, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		logger.Error("Error parsing templates", "err", err)
	}

	renderer := &Templatereader{
		templates: tmpls,
	}

	mux := http.NewServeMux()
	http.Handle("static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "base.html", nil)
	})

	mux.HandleFunc("GET /toggle-display-endpoint", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "dropdownMenu.html", nil)
	})

	logger.Info("Server started on port 8080")
	http.ListenAndServe(":8080", mux)
}
