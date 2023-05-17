package main

import (
	"net/http"
	"os"

	"github.com/KarolKozlowski/webmarks-server/redirector"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Options
	portNumber := "9000"

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	http.HandleFunc("/", redirector.RootHandler)

	log.Info().Msg("Server listening on port " + portNumber)
	http.ListenAndServe(":"+portNumber, nil)

	log.Info().Msg("Server shutting down")
}
