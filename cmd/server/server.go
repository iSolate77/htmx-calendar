package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	// _ "github.com/lib/pq" // PostgreSQL driver
)

// var db *sql.DB

func init() {
	// Initialize the database connection
	// connStr := "user=username dbname=mydb sslmode=disable password=mypassword"
	// db, _ = sql.Open("postgres", connStr)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Initialize templates
	tmpl, err := template.ParseGlob("../../web/templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/get-subjects", getSubjectsHandler)
	// http.HandleFunc("/fetch-curriculum", fetchCurriculumHandler)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on port 8080")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the main page
	tmpl.ExecuteTemplate(w, "base.html", nil)
}

// func getSubjectsHandler(w http.ResponseWriter, r *http.Request) {
// 	// Fetch available subjects from the database and send them back
// 	// This can be used to populate the checkboxes for subject selection
// }
//
// func fetchCurriculumHandler(w http.ResponseWriter, r *http.Request) {
// 	// Based on selected subjects, fetch the respective curriculum
// 	// Render the curriculum into HTML and send it back
// }
