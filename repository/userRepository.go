package repository

import (
	"github.com/DanielFrag/widgets-spa-rv/infra"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUserByID(string) (model.User, error)
	GetUserByLogin(string, string) (model.User, error)
	UpdateUserSession(string, string) error
}

func GetUserRepository() UserRepository {
	return infra.GetUserDB()
}
