package main

import (
	"log"
	"net/http"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path

	shortcuts := make(map[string]string)
	shortcuts["/dotnot"] = "https://dotnot.pl"
	shortcuts["/wp"] = "https://wp.pl"

	WarningLogger.Printf("Received request %s%s\n", r.Host, path)

	target, ok := shortcuts[path]
	if ok {
		InfoLogger.Printf("Redirecting -> %s\n", target)
		http.Redirect(w, r, target, http.StatusSeeOther)
	} else {
		WarningLogger.Printf("Received request %s%s\n", r.Host, path)
	}
}

func main() {
	// Options
	portNumber := "9000"

	// main handler
	http.HandleFunc("/", rootHandler)
	InfoLogger.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
