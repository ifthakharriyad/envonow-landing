// It's the landing page for envonow.com
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	mux := http.NewServeMux()

	// Serves static files from the public directory
	stfiles := http.FileServer(http.Dir(os.Getenv("STATIC")))
	mux.Handle("/static/", http.StripPrefix("/static/", stfiles))

	// GET /error
	// mux.HandleFunc("/error",err)

	// GET /
	// Index route
	mux.HandleFunc("/", index)

	// defining and starting server
	server := &http.Server{
		Addr:    os.Getenv("ADDRESS"),
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}

// handles GET / route
func index(w http.ResponseWriter, r *http.Request) {
	// Directories containing template files
	temptdirs := []string{"templates", "templates/components", "templates/sections"}
	temptfiles := []string{}
	for _, dir := range temptdirs {
		// Getting .html files from the dir directory
		files, err := dirfiles(dir, ".+.html")
		if err != nil {
			danger(err)
			continue
		}
		// Getting the full file paths(relative to root of the repo) of the files
		files = filepaths(dir, files)
		temptfiles = append(temptfiles, files...)
	}
	template, err := template.ParseFiles(temptfiles...)
	if err != nil {
		danger(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	template.ExecuteTemplate(w, "layout", os.Getenv("FORM_HERO_LINK"))
}
