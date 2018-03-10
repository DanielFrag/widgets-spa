package infra

import (
	"testing"
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/DanielFrag/widgets-spa-rv/model"
)

func TestUserMGO(t *testing.T) {
	var userID bson.ObjectId
	userLogin, userPass, userGravatar := "sunda", "adnus", "www.sunda.com.br/adnus"
	t.Run("StartDB", func(t *testing.T) {
		startDBError := StartDB()
		if startDBError != nil {
			t.Error("Can't starts the DB")
		}
		ds.dbName = ds.dbName + "_test"
	})
	defer func() {
		mgoSession := getSession()
		dropDatabaseError := mgoSession.DB(getDbName()).DropDatabase()
		if dropDatabaseError != nil {
			panic(dropDatabaseError)
		}
		StopDB()
	}()
	t.Run("EmptyDB", func(t *testing.T) {
		userMGO := GetUserDB()
		users, usersError := userMGO.GetUsers()
		if usersError != nil {
			t.Error("Error retrieving users")
		}
		if len(users) != 0 {
			t.Error("Testing with no empty DB")
		}
	})
	t.Run("InsertFirstUser", func(t *testing.T) {
		user := model.User {
			Login: userLogin,
			Password: userPass,
			Gravatar: userGravatar,
		}
		mgoSession := getSession()
		userCollection := mgoSession.DB(getDbName()).C("User")
		insertUserError := userCollection.Insert(user)
		if insertUserError != nil {
			t.Error("Inserting user error: " + insertUserError.Error())
		}
	})
	t.Run("RecoverUsers1", func(t *testing.T) {
		userMGO := GetUserDB()
		users, usersError := userMGO.GetUsers()
		if usersError != nil {
			t.Error("Error retrieving users")
		}
		if len(users) != 1 {
			t.Error("Can't find the inserted user")
		}
		userID = users[0].ID
	})
	t.Run("RecoverSingleUser", func(t *testing.T) {
		userMGO := GetUserDB()
		user, userError := userMGO.GetUserByID(userID.Hex())
		if userError != nil {
			t.Error("Recovering user error: " + userError.Error())
		}
		if user.Gravatar != userGravatar || user.Login != userLogin || user.Password != userPass {
			t.Error("Inconsistent user data")
		}
	})
	t.Run("InsertSecondUser", func(t *testing.T) {
		user := model.User {
			Login: userLogin + "2",
			Password: userPass + "2",
			Gravatar: userGravatar + "2",
		}
		mgoSession := getSession()
		userCollection := mgoSession.DB(getDbName()).C("User")
		insertUserError := userCollection.Insert(user)
		if insertUserError != nil {
			t.Error("Inserting user error: " + insertUserError.Error())
		}
	})
	t.Run("RecoverUsers2", func(t *testing.T) {
		userMGO := GetUserDB()
		users, usersError := userMGO.GetUsers()
		if usersError != nil {
			t.Error("Error retrieving users")
		}
		if len(users) != 2 {
			t.Error("Can't find the inserted users")
		}
	})
	t.Run("SearchUnexistingUser", func(t *testing.T) {
		userMGO := GetUserDB()
		wrongID := bson.NewObjectIdWithTime(time.Now())
		_, userError := userMGO.GetUserByID(wrongID.Hex())
		if userError == nil {
			t.Error("Found an uninserted document")
		}
	})
}