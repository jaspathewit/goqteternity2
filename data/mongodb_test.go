// Unit tests for the mongodb
package data

import (
	. "launchpad.net/gocheck"
)

const (
	dbName = "testEternity2"
)

func (context *DataSuite) SetUpTest(c *C) {
	// open the test database
	err := Open("localhost", dbName)
	if err != nil {
		panic(err)
	}

	// delete the database
	DropDatabase(dbName)
}

func (context *DataSuite) TearDownTest(c *C) {
	Close()
}

// Test saving a tile
func (context *DataSuite) Test_Save(c *C) {
	var doc *MongoDocumentTile

	doc = new(MongoDocumentTile)

	doc.Id = "TestTile"

	Save(doc)

}
