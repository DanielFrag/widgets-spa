package infra

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

//UserMGO wrap the session to access user data
type UserMGO struct {
	session *mgo.Session
}

//GetUsers return all the users
func (u *UserMGO) GetUsers() ([]model.User, error) {
	u.session = getSession()
	defer u.session.Close()
	usersCollection := u.session.DB(getDbName()).C("User")
	var users []model.User
	err := usersCollection.Find(bson.M{}).All(&users)
	return users, err
}

//GetUserByID return a single user based on his ID
func (u *UserMGO) GetUserByID(userID string) (model.User, error) {
	u.session = getSession()
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

//GetUserDB return the object to access the users data
func GetUserDB() *UserMGO {
	return &UserMGO {}
}