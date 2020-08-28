package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thanhpk/randstr"
	"github.com/turbaszek/tnijto/pkg/util"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var config = util.NewConfig()
var fs = util.NewFirestore(config.GcpProject)

// NewRouter creates instance of new tnijto router
func NewRouter() *http.Server {
	log.Printf("The app is running under: http://%s:%s/", config.Hostname, config.Port)

	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/api/new", submitNewLinkHandler)
	router.HandleFunc("/{.*}", redirectHandler)

	router.Use(LoggingMiddleware)
	router.NotFoundHandler = Handle404()

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", config.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
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

	l := util.Link{URL: originalURL, ID: id, GeneratedURL: generatedURL}
	fs.SaveLink(l)

	js, err := json.Marshal(l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	var link util.Link
	if r.RequestURI == "/favicon.ico" {
		return
	}
	linkID := strings.TrimLeft(r.RequestURI, "/")

	if err := fs.ReadLink(linkID, &link); err != nil {
		handleErrorWithRedirect(w, r, err)
		return
	}

	u, err := url.QueryUnescape(link.URL)
	if err != nil {
		handleErrorWithRedirect(w, r, err)
		return
	}

	http.Redirect(w, r, u, http.StatusMovedPermanently)
}

func handleErrorWithRedirect(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("An error has occurred: %s", err)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
