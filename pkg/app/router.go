package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/turbaszek/tnijto/pkg/util"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var fs = util.NewFirestore(util.Config.GcpProject)

// NewRouter creates instance of new tnijto router
func NewRouter() *http.Server {
	log.Printf("The app is running under: http://%s:%s/", util.Config.Hostname, util.Config.Port)

	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/api/new", submitNewLinkHandler)
	router.HandleFunc("/{.*}", redirectHandler)

	router.Use(LoggingMiddleware)
	router.NotFoundHandler = Handle404()

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", util.Config.Port),
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
	value := r.FormValue("value")

	// To avoid redirect loop
	if originalURL == value {
		http.Error(w, "Link value must be different than url", http.StatusBadRequest)
		return
	}

	l := util.NewLink(originalURL, value)
	if err := fs.SaveLink(l); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	value := strings.TrimLeft(r.RequestURI, "/")
	if err := fs.ReadLink(value, &link); err != nil {
		handleErrorWithRedirect(w, r, err)
		return
	}

	u, err := url.QueryUnescape(link.URL)
	if err != nil {
		handleErrorWithRedirect(w, r, err)
		return
	}

	// To avoid redirect loop
	if link.URL == link.Value {
		http.Error(w, "Alternative url value must be different than the url", http.StatusBadRequest)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Update views count in async manner
	go fs.UpdateViewsCount(link)
	http.Redirect(w, r, u, http.StatusMovedPermanently)
}

func handleErrorWithRedirect(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("An error has occurred: %s", err)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
