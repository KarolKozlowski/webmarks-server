package redirector

import (
	"io/fs"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
)

const (
	Main    string = "main.tmpl"
	Landing string = "main.tmpl"
)

var templates *template.Template

type PageData struct {
	Title string
}

type Shortcut struct {
	Target string
}

func renderPage(w http.ResponseWriter, t string, data any) {
	header := templates.Lookup("header.tmpl")
	header.Execute(w, data)
	body := templates.Lookup(t)
	body.Execute(w, data)
	footer := templates.Lookup("footer.tmpl")
	footer.Execute(w, data)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var path = strings.Replace(r.URL.Path, "/", "", 1)
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

	if path != "" {
		var shortcuts = map[string]Shortcut{}
		shortcuts["dotnot"] = Shortcut{Target: "https://dotnot.pl"}
		shortcuts["wp"] = Shortcut{Target: "https://wp.pl"}

		log.Warn().Msg("Received request: " + r.Host + path)

		shortcut, ok := shortcuts[path]
		if ok {
			log.Info().Msg("Redirecting " + path + " -> " + shortcut.Target)
			http.Redirect(w, r, shortcut.Target, http.StatusSeeOther)
		} else {
			log.Warn().Msg("Redirect not found")

			data := PageData{
				Title: "Redirect: " + path + " not found!",
			}
			renderPage(w, Main, data)

		}
	} else {
		data := PageData{
			Title: "Welcome!",
		}
		renderPage(w, Landing, data)
	}

}
