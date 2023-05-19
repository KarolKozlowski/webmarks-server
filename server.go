package main

import (
	"net/http"
	"os"

	"github.com/KarolKozlowski/webmarks-server/redirector"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func redirectToStatic(w http.ResponseWriter, r *http.Request) {
	urlBase := r.URL.Path
	log.Debug().Msg("Redirect to /static/: " + urlBase + " -> " + "/static" + r.URL.Path)
	http.Redirect(w, r, "/static"+r.URL.Path, http.StatusSeeOther)
}

func main() {
	// Options
	portNumber := "9000"

	// setup logger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// handle static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// redirect root static resources
	http.HandleFunc("/favicon.ico", redirectToStatic)

	// main handler
	http.HandleFunc("/", redirector.RootHandler)

	log.Info().Msg("Server listening on port " + portNumber)

	http.ListenAndServe(":"+portNumber, nil)

	log.Info().Msg("Server shutting down")
}
