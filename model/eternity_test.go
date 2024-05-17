// Unit tests for the Side entity
package model

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test_Eternity(t *testing.T) {
	fmt.Println("Eternity Tests")
	TestingT(t)
}

type EternitySuite struct{}

var _ = Suite(&EternitySuite{})

// "Test" the tile statistics
func (context *EternitySuite) Test_Statistics(c *C) {
	PrintTileStatistics()
}
