// Unit tests for the Side entity
package model

import (
	"fmt"

	. "gopkg.in/check.v1"
)

// Test maping side def to the internal representation
func (context *EternitySuite) Test_DefToSide(c *C) {
	side, found := DefinitionToSide("000000")
	c.Assert(side, Equals, byte(0))
	c.Assert(found, Equals, true)

	_, found = DefinitionToSide("999999")
	c.Assert(found, Equals, false)
}

// Test maping internal representation to side def
func (context *EternitySuite) Test_SideToDef(c *C) {
	def, found := SideToDefinition(byte(0))
	c.Assert(def, Equals, "000000")
	c.Assert(found, Equals, true)

	_, found = SideToDefinition(byte(23))
	c.Assert(found, Equals, false)
}

// Test maping side definition to internal representation
func (context *EternitySuite) Test_PrintingSide(c *C) {
	def, _ := DefinitionToSide("000000")
	actual := fmt.Sprintf("%c", getSideChar(def))
	c.Assert(actual, Equals, "X")
}

// Sanity check the side definitions
func (context *EternitySuite) Test_SanityCheck(c *C) {
	key := byte('A')
	def, found := SideToDefinition(key)

	for found {
		// fmt.Printf("Testing: \n" + def)
		subDef := def[0:1]
		c.Assert(subDef, Not(Equals), "00")
		subDef = def[2:3]
		c.Assert(subDef, Not(Equals), "00")
		subDef = def[4:5]
		c.Assert(subDef, Not(Equals), "00")
		key = key + 1
		def, found = SideToDefinition(key)
	}
}
