// mongodb. Provides the functions for managing mongodb instance used for the
// persistance of Tiles and jobs.
package data

import (
	// "bitbucket.com/jaspathewit/eternity2/domain"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	log "github.com/cihub/seelog"
	// "fmt"
)

var session *mgo.Session
var db string

// Open the session to the mongodb
func Open(url string, dbName string) error {
	var err error
	if session == nil {
		log.Info("Opening Mongodb")

		session, err = mgo.Dial(url)
		if err != nil {
			return log.Error(err)
		}

		db = dbName

		log.Info("Opened")
	}

	return err
}

// Close the session to the mongodb
func Close() {
	if session != nil {
		log.Info("Closing Mongodb")
		session.Close()
		session = nil
		log.Info("Closed")
	}
}

// function drops the given database
func DropDatabase(dbName string) error {
	var err error
	if err = checkOpened(); err == nil {
		if db, err := getDatabase(dbName); err == nil {
			err = db.DropDatabase()
		}
	}
	return err
}

// save a tile in the mongodb
func Save(tile *MongoDocumentTile) error {
	var err error
	if err = checkOpened(); err == nil {
		err = session.DB(db).C("Doc").Insert(tile)
	}
	return err
}

// function checks if the database has been opened
func checkOpened() error {
	var err error
	if session == nil {
		err = log.Error("Mongodb has not been opened yet.")
	}
	return err
}

// function gets the mgo.Database instance for the given name
func getDatabase(dbName string) (*mgo.Database, error) {
	var err error
	var result *mgo.Database
	if err = checkOpened(); err == nil {
		result = session.DB(dbName)
	}
	return result, err
}
