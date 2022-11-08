package auth

import (
	"github.com/intv/scr/database/gorm/models"
	"github.com/intv/scr/helpers"
	"github.com/intv/scr/interfaces"
)

type token_response struct {
	Tokens string `json:"token"`
}

type auth_service struct {
	rep interfaces.UserRepo
}

func NewService(rep interfaces.UserRepo) *auth_service {
	return &auth_service{rep}
}

func (r *auth_service) Login(body models.User) (*helpers.Response, error) {
	user, err := r.rep.FindByUsername(body.Username)
	if err != nil {
		return nil, err
	}

	if !helpers.CheckPassword(user.Password, body.Password) {
		return helpers.New("wrong password", 401, true), nil
	}

	token := helpers.NewToken(body.Username)
	theToken, err := token.Create()
	if err != nil {
		return nil, err
	}

	response := helpers.New(token_response{Tokens: theToken}, 200, false)

	return response, nil

}
