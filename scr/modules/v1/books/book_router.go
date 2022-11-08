package books

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/book").Subrouter()

	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
}
