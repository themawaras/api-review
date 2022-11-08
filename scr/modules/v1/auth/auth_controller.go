package auth

import (
	"encoding/json"
	"net/http"

	"github.com/intv/scr/database/gorm/models"
	"github.com/intv/scr/interfaces"
)

type auth_ctrl struct {
	repo interfaces.AuthService
}

func NewCtrl(rep interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{rep}
}

func (re *auth_ctrl) Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User

	json.NewDecoder(r.Body).Decode(&data)
	result, err := re.repo.Login(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result.Send(w)
}
