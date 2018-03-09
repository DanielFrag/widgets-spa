package infra

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

type UserMGO struct {
	session *mgo.Session
}

func (u *UserMGO) GetUsers() ([]model.User, error) {
	defer u.session.Close()
	usersCollection := u.session.DB(getDbName()).C("User")
	var users []model.User
	err := usersCollection.Find(bson.M{}).All(&users)
	return users, err
}

func (u *UserMGO) GetUserByID(userID string) (model.User, error) {
	defer u.session.Close()
	usersCollection := u.session.DB(getDbName()).C("User")
	var user model.User
	err := usersCollection.
		Find(bson.M{
			"_id": bson.ObjectIdHex(userID),
		}).
		All(&user)
	return user, err
}

func GetUserDB() *UserMGO {
	return &UserMGO {
		session: getSession(),
	}
}