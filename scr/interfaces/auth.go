package interfaces

import (
	"github.com/intv/scr/database/gorm/models"
	"github.com/intv/scr/helpers"
)

type AuthService interface {
	Login(body models.User) (*helpers.Response, error)
}
