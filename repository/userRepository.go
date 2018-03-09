package repository

import (
	"github.com/DanielFrag/widgets-spa-rv/infra"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUserByID(string) (model.User, error)
}

func GetUserRepository() UserRepository {
	return infra.GetUserDB()
}
