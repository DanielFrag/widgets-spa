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
	usersCollection := u.session.DB(getDbName()).C("user")
	var users []model.User
	err := usersCollection.Find(bson.M{}).All(&users)
	return users, err
}

//GetUserByID return a single user based on its ID
func (u *UserMGO) GetUserByID(userID string) (model.User, error) {
	u.session = getSession()
	defer u.session.Close()
	usersCollection := u.session.DB(getDbName()).C("user")
	var user model.User
	err := usersCollection.
		Find(bson.M{
			"_id": bson.ObjectIdHex(userID),
		}).
		One(&user)
	return user, err
}

//GetUserByLogin return a single user based on its Login
func (u *UserMGO) GetUserByLogin(userLogin, userPass string) (model.User, error) {
	u.session = getSession()
	defer u.session.Close()
	usersCollection := u.session.DB(getDbName()).C("user")
	var user []model.User
	err := usersCollection.
		Find(bson.M{
			"login": userLogin,
			//"password": userPass,
		}).
		All(&user)
	return user[0], err
}

//UpdateUserSession set new values for an user session
func (u *UserMGO) UpdateUserSession(userID, session string) error {
	u.session = getSession()
	defer u.session.Close()
	usersCollection := u.session.DB(getDbName()).C("user")
	return usersCollection.Update(bson.M {
		"_id": bson.ObjectIdHex(userID),
	}, bson.M {
		"$set": bson.M{
			"session": session,
		},
	})
}

//GetUserDB return the object to access the users data
func GetUserDB() *UserMGO {
	return &UserMGO {}
}