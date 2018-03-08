package repository

import (
	"github.com/DanielFrag/widgets-spa/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUserByID(string) (model.User, error)
}
