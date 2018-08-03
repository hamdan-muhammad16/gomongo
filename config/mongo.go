package config

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {
	os.Setenv("MONGO_HOST", "localhost")
	os.Setenv("MONGO_DB_NAME", "mongogo")

	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}
