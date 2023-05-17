package redirector

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
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
