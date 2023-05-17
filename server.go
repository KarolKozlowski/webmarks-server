package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path

	shortcuts := make(map[string]string)
	shortcuts["/dotnot"] = "https://dotnot.pl"
	shortcuts["/wp"] = "https://wp.pl"

	log.Warn().Msg("Received request: " + r.Host + path)

	target, ok := shortcuts[path]
	if ok {
		log.Info().Msg("Redirecting -> " + target)
		http.Redirect(w, r, target, http.StatusSeeOther)
	} else {
		log.Warn().Msg("Redirect not found")
	}
}

func main() {
	// Options
	portNumber := "9000"

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// main handler
	http.HandleFunc("/", rootHandler)
	log.Info().Msg("Server listening on port " + portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
