package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/intv/scr/configs/database"
	"github.com/intv/scr/modules/v1/auth"
	"github.com/intv/scr/modules/v1/books"
	"github.com/intv/scr/modules/v1/users"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	books.New(mainRoute, db)
	users.New(mainRoute, db)
	auth.New(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Home!"))
}
