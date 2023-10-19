package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Templatereader struct {
	templates *template.Template
}

func (t *Templatereader) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	tmpls, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %s", err.Error())
	}

	e := echo.New()

	e.Renderer = &Templatereader{
		templates: tmpls,
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "base.html", nil)
	})
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
