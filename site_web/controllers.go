package spotifyAPI

import (
	"log"
	"net/http"
)

// Root handler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	jul := getArtistData("3IW7ScrzXmPvZhB27hmfgy")
	sdm := getArtistData("0LKAV3zJ8a8AIGnyc5OvfB")
	data := IndexData{Base: BaseData{
		Title:      "Thorgify - Accueil",
		StaticPath: "static/",
		Line:       "enabled",
	},
		Artists: []Artist{
			{Name: jul.Name, ImgUrl: jul.Images[0].Url, Title: "Tous les albums de Jul", Link: "/album/jul"},
			{Name: sdm.Name, ImgUrl: sdm.Images[0].Url, Title: "Bolide Allemand de SDM", Link: "/track/sdm"},
		},
	}
	err := tmpl["index"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

// Index page handler.
func albumJulHandler(w http.ResponseWriter, r *http.Request) {
	data := AlbumsJulData{Base: BaseData{
		Title:      "Thorgify - Albums",
		StaticPath: "../static/",
		Line:       "enabled",
	},
		Data:   getAlbumsJul(),
		Artist: getArtistData("3IW7ScrzXmPvZhB27hmfgy"),
	}
	err := tmpl["albumJul"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

// Creating user page handler.
func trackSdmHandler(w http.ResponseWriter, r *http.Request) {
	data := TrackSdmData{Base: BaseData{
		Title:      "Thorgify - Morceaux",
		StaticPath: "../static/",
		Line:       "enabled",
	},
		Data:   getTrackSdm(),
		Artist: getArtistData("0LKAV3zJ8a8AIGnyc5OvfB"),
	}
	err := tmpl["trackSdm"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}
