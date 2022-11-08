package main

import (
	"log"
	"net/http"

	"github.com/intv/scr/routers"
)

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := http.ListenAndServe(":8008", mainRoute); err != nil {
		log.Fatal("Application error")
	}
}
