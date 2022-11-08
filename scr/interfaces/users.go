package interfaces

import (
	"github.com/intv/scr/database/gorm/models"
	"github.com/intv/scr/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	FindByUsername(username string) (*models.User, error)
	Add(*models.User) (*models.User, error)
}

type UserService interface {
	FindAll() (*helpers.Response, error)
	FindByUsername(username string) (*helpers.Response, error)
	Save(*models.User) (*helpers.Response, error)
}
