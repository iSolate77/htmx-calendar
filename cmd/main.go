package main

import (
	"calendar/pkg/api"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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
	e.Static("/static", "web/static")

	e.Renderer = &Templatereader{
		templates: tmpls,
	}

	e.GET("/", func(c echo.Context) error {
		mockData := api.GenerateMockData("2020-01-01", 16)
		err = c.Render(http.StatusOK, "base.html", map[string]interface{}{
			"Weeks": mockData,
		})
		if err != nil {
			log.Printf("Error rendering template: %s", err.Error())
			return c.String(http.StatusInternalServerError, "Error rendering template")
		}
		return nil
	})

	e.GET("/toggle-display-endpoint", menuDropdown)

	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test response")
	})

	log.Println("Server started on port 8080")
	e.Start(":8080")
}

func menuDropdown(c echo.Context) error {
	return c.Render(http.StatusOK, "dropdownMenu.html", nil)
}
