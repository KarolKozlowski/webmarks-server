package redirector

import (
	"io/fs"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
)

var templates *template.Template

type PageData struct {
	Title string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path
	var allFiles []string

	template_dir := os.DirFS("./templates")
	template_files, err := fs.Glob(template_dir, "*.tmpl")
	if err != nil {
		log.Error().Msg(err.Error())
	}

	// find all templates
	for _, filename := range template_files {
		if strings.HasSuffix(filename, ".tmpl") {
			allFiles = append(allFiles, "./templates/"+filename)
		}
	}
	// parse all templates
	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		log.Error().Msg(err.Error())
	}

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

		data := PageData{
			Title: "Redirect not found!",
		}

		s1 := templates.Lookup("main.tmpl")
		s1.Execute(w, data)
	}
}
