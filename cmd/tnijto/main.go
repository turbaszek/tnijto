package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thanhpk/randstr"
	utils "github.com/turbaszek/tnijto/internal"
	"log"
	"net/http"
	"net/url"
	"time"
)

var config = utils.NewConfig()
var fs = utils.NewFirestore(config.GcpProject)

func main() {
	log.Printf("The app is running under: http://%s:%s/", config.Hostname, config.Port)

	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/api/new", submitNewLinkHandler)
	router.HandleFunc("/{.*}", redirectHandler)

	router.Use(utils.LoggingMiddleware)
	router.NotFoundHandler = utils.Handle404()

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", config.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func submitNewLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	originalURL := r.FormValue("originalURL")
	id := r.FormValue("id")

	if id == "" {
		id = randstr.String(11)
	}

	generatedURL := fmt.Sprintf("https://%s/%s", config.Hostname, id)

	l := utils.Link{URL: originalURL, ID: id, GeneratedURL: generatedURL}
	fs.SaveLink(l)

	js, err := json.Marshal(l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}
	// Skip leading / in redirect link
	id := r.RequestURI[1:]
	link, err := fs.ReadLink(id)

	if err != nil {
		log.Printf("An error has occurred: %s", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u, err := url.QueryUnescape(link.URL)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, u, http.StatusMovedPermanently)
}
