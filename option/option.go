// options.go
package option

import (
	"flag"
	"fmt"
	"os"
)

// define varables to hold the options
var Verbose bool
var Validate bool

// Method creates a human readable form of the Options struct
func String() string {
	result := fmt.Sprintf("Verbose: %t Validate: %t",
		Verbose, Validate)
	return result
}

// Method validates the options
func ValidateOptions() error {
	return nil
}

// print out the help text
func usage() {
	fmt.Fprintf(os.Stderr, "usage: entity2Solver [-validate] [-verbose]")
	flag.PrintDefaults()
}

func GetOptions() error {

	// define the options
	flag.BoolVar(&Verbose, "verbose", false, "Verbose output during processing.")
	flag.BoolVar(&Validate, "validate", false, "Validate Surface after each manipulation.")

	// add the default usage method
	flag.Usage = usage

	// parse the command line
	flag.Parse()

	// validate the options
	var err error
	if err = ValidateOptions(); err != nil {
		fmt.Printf("%s\n", err.Error())
		flag.Usage()
	}

	return err
}
