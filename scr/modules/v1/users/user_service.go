package users

import (
	"github.com/intv/scr/database/gorm/models"
	"github.com/intv/scr/helpers"
	"github.com/intv/scr/interfaces"
)

type users_service struct {
	rep interfaces.UserRepo
}

func NewService(rep interfaces.UserRepo) *users_service {
	return &users_service{rep}
}

func (re *users_service) FindAll() (*helpers.Response, error) {
	data, err := re.rep.FindAll()
	if err != nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}

func (re *users_service) FindByUsername(username string) (*helpers.Response, error) {
	data, err := re.rep.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}

func (re *users_service) Save(user *models.User) (*helpers.Response, error) {
	hashPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPassword
	data, err := re.rep.Add(user)
	if err != nil {
		return nil, err
	}

	response := helpers.New(data, 200, false)
	return response, nil
}
