package spotifyAPI

import (
	"log"
	"net/http"
)

// Root handler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Index page handler.
func albumJulHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "albumJul", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Creating user page handler.
func trackSdmHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "trackSdm", nil)
	if err != nil {
		log.Fatal(err)
	}
}
