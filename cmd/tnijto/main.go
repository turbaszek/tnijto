package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	utils "github.com/turbaszek/tnijto/pkg"
	"log"
	"net/http"
	"time"
)

var config = utils.NewConfig()

const port = 1317

func main() {
	log.Printf("The app is running under: http://%s:%d/", config.Hostname, port)

	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/new", submitNewLinkHandler)

	router.Use(utils.LoggingMiddleware)
	router.NotFoundHandler = utils.Handle404()

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// Link represents the link
type Link struct {
	URL          string
	Name         string
	GeneratedURL string
}

func submitNewLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	link := r.FormValue("link")
	name := r.FormValue("name")
	generated := fmt.Sprintf("https://%s/%s", config.Hostname, name)

	linkResponse := Link{link, name, generated}

	js, err := json.Marshal(linkResponse)
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
