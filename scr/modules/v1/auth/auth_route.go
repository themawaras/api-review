package auth

import (
	"github.com/gorilla/mux"
	"github.com/intv/scr/modules/v1/users"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()

	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.Signin).Methods("POST")
}
