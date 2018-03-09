package infra

import (
	"os"

	"gopkg.in/mgo.v2"
)

type dataStore struct {
	session *mgo.Session
	dbName  string
}

var ds dataStore

//StartDB initialize DB connection
func StartDB() {
	var err error
	mongoURL := os.Getenv("MONGODB_URI")
	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}
	if ds.session == nil {
		ds.session, err = mgo.Dial(mongoURL)
	}
	ds.dbName = os.Getenv("DB_NAME")
	if ds.dbName == "" {
		ds.dbName = "widgets-spa"
	}
	if err != nil {
		panic(err)
	}
}

//StopDB close DB session
func StopDB() {
	ds.session.Close()
}

//GetSession return the current DB session
func getSession() *mgo.Session {
	return ds.session.Clone()
}

func getDbName() string {
	return ds.dbName
}
