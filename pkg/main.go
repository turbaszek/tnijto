package main

import (
	"github.com/turbaszek/tnijto/pkg/app"
	"log"
)

func main() {
	srv := app.NewRouter()
	log.Fatal(srv.ListenAndServe())
}
