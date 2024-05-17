// Unit tests for the data package
package data

import (
	"fmt"
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test_Data(t *testing.T) {
	fmt.Println("Data Tests")
	TestingT(t)
}

type DataSuite struct{}

var _ = Suite(&DataSuite{})
