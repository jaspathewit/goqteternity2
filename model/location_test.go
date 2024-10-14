// Unit tests for the Tile1 entity
package model

import (
	. "gopkg.in/check.v1"
)

// Test the cloning of a location
func (context *EternitySuite) Test_Location_Clone(c *C) {

	location := new(Location)

	location.Row = 34
	location.Column = 300

	clone := location.Clone()

	c.Assert(clone.Row, Equals, location.Row)
	c.Assert(clone.Column, Equals, location.Column)

}
