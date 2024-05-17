// Single Eternity tile.
package model

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"bitbucket.com/jaspathewit/eternity2/util"
)

// Single Eternity tile
type Tile1 struct {
	TypeOf byte
	// Number of a tile (the fist tile 001 is Number 0)
	Number      byte
	Orientation byte

	// The sides of the tile
	Top    byte
	Right  byte
	Bottom byte
	Left   byte

	Energy int
}

// The empty Tile1
var TILE1INSTANCE *Tile1

// the array that holds the 256 tiles in each of the 4 orientations, indexed from 0
var TileSet [][]*Tile1

// initilise the set of tiles
func init() {
	// create the slices for the TileSet
	TileSet = make([][]*Tile1, 4)
	for i := range TileSet {
		TileSet[i] = make([]*Tile1, 256)
	}

	// set up the tile1 instance
	TILE1INSTANCE = new(Tile1)

	addTile("C", "001", "010103", "020102", "000000", "000000")
	addTile("C", "002", "010103", "030504", "000000", "000000")
	addTile("C", "003", "040605", "020102", "000000", "000000")
	addTile("C", "004", "020102", "040605", "000000", "000000")
	addTile("E", "005", "050602", "010103", "000000", "010103")
	addTile("E", "006", "060107", "040605", "000000", "010103")
	addTile("E", "007", "070607", "010103", "000000", "010103")
	addTile("E", "008", "070607", "080308", "000000", "010103")
	addTile("E", "009", "060304", "020102", "000000", "010103")
	addTile("E", "010", "090201", "030504", "000000", "010103")
	addTile("E", "011", "100406", "040605", "000000", "010103")
	addTile("E", "012", "050305", "030504", "000000", "010103")
	addTile("E", "013", "050305", "080308", "000000", "010103")
	addTile("E", "014", "090602", "030504", "000000", "010103")
	addTile("E", "015", "060107", "010103", "000000", "040605")
	addTile("E", "016", "110106", "020102", "000000", "040605")
	addTile("E", "017", "050204", "080308", "000000", "040605")
	addTile("E", "018", "110608", "080308", "000000", "040605")
	addTile("E", "019", "090201", "040605", "000000", "040605")
	addTile("E", "020", "100406", "040605", "000000", "040605")
	addTile("E", "021", "060201", "030504", "000000", "040605")
	addTile("E", "022", "090506", "010103", "000000", "040605")
	addTile("E", "023", "090506", "080308", "000000", "040605")
	addTile("E", "024", "110703", "010103", "000000", "040605")
	addTile("E", "025", "090602", "010103", "000000", "040605")
	addTile("E", "026", "050602", "040605", "000000", "020102")
	addTile("E", "027", "050602", "020102", "000000", "020102")
	addTile("E", "028", "060107", "020102", "000000", "020102")
	addTile("E", "029", "110106", "020102", "000000", "020102")
	addTile("E", "030", "090201", "080308", "000000", "020102")
	addTile("E", "031", "100406", "040605", "000000", "020102")
	addTile("E", "032", "110703", "020102", "000000", "020102")
	addTile("E", "033", "050305", "040605", "000000", "020102")
	addTile("E", "034", "050305", "030504", "000000", "020102")
	addTile("E", "035", "100105", "080308", "000000", "020102")
	addTile("E", "036", "070205", "030504", "000000", "020102")
	addTile("E", "037", "110106", "010103", "000000", "030504")
	addTile("E", "038", "070408", "080308", "000000", "030504")
	addTile("E", "039", "060304", "080308", "000000", "030504")
	addTile("E", "040", "110608", "040605", "000000", "030504")
	addTile("E", "041", "110608", "020102", "000000", "030504")
	addTile("E", "042", "100406", "010103", "000000", "030504")
	addTile("E", "043", "100406", "040605", "000000", "030504")
	addTile("E", "044", "100406", "020102", "000000", "030504")
	addTile("E", "045", "060201", "010103", "000000", "030504")
	addTile("E", "046", "110703", "030504", "000000", "030504")
	addTile("E", "047", "050305", "030504", "000000", "030504")
	addTile("E", "048", "100105", "030504", "000000", "030504")
	addTile("E", "049", "050602", "080308", "000000", "080308")
	addTile("E", "050", "060107", "010103", "000000", "080308")
	addTile("E", "051", "060107", "040605", "000000", "080308")
	addTile("E", "052", "070607", "010103", "000000", "080308")
	addTile("E", "053", "090201", "030504", "000000", "080308")
	addTile("E", "054", "060201", "030504", "000000", "080308")
	addTile("E", "055", "060201", "080308", "000000", "080308")
	addTile("E", "056", "050305", "020102", "000000", "080308")
	addTile("E", "057", "100105", "010103", "000000", "080308")
	addTile("E", "058", "100105", "080308", "000000", "080308")
	addTile("E", "059", "090602", "040605", "000000", "080308")
	addTile("E", "060", "070205", "020102", "000000", "080308")
	addTile("M", "061", "070607", "110106", "050602", "050602")
	addTile("M", "062", "050204", "090201", "050602", "050602")
	addTile("M", "063", "060107", "070408", "050602", "060107")
	addTile("M", "064", "050602", "050305", "050602", "110106")
	addTile("M", "065", "110106", "070205", "050602", "110106")
	addTile("M", "066", "050204", "050204", "050602", "110106")
	addTile("M", "067", "060304", "060107", "050602", "110106")
	addTile("M", "068", "110703", "070607", "050602", "110106")
	addTile("M", "069", "070205", "050305", "050602", "110106")
	addTile("M", "070", "070408", "090201", "050602", "070408")
	addTile("M", "071", "090201", "090506", "050602", "070408")
	addTile("M", "072", "050204", "110106", "050602", "060304")
	addTile("M", "073", "100406", "060201", "050602", "060304")
	addTile("M", "074", "110703", "100406", "050602", "060304")
	addTile("M", "075", "050305", "070408", "050602", "060304")
	addTile("M", "076", "050204", "100406", "050602", "110608")
	addTile("M", "077", "110608", "100406", "050602", "110608")
	addTile("M", "078", "070408", "100105", "050602", "090201")
	addTile("M", "079", "110703", "070408", "050602", "090201")
	addTile("M", "080", "100105", "090602", "050602", "090201")
	addTile("M", "081", "110608", "110106", "050602", "100406")
	addTile("M", "082", "110106", "110106", "050602", "060201")
	addTile("M", "083", "060304", "060201", "050602", "060201")
	addTile("M", "084", "110106", "110608", "050602", "090506")
	addTile("M", "085", "070607", "050204", "050602", "090506")
	addTile("M", "086", "050305", "090506", "050602", "090506")
	addTile("M", "087", "100105", "110703", "050602", "090506")
	addTile("M", "088", "050602", "090602", "050602", "110703")
	addTile("M", "089", "070607", "070205", "050602", "110703")
	addTile("M", "090", "060201", "100105", "050602", "110703")
	addTile("M", "091", "060304", "090506", "050602", "050305")
	addTile("M", "092", "110608", "100406", "050602", "050305")
	addTile("M", "093", "110608", "060201", "050602", "050305")
	addTile("M", "094", "060201", "090602", "050602", "050305")
	addTile("M", "095", "090506", "050204", "050602", "050305")
	addTile("M", "096", "090602", "070408", "050602", "090602")
	addTile("M", "097", "060201", "110608", "050602", "070205")
	addTile("M", "098", "110703", "050305", "050602", "070205")
	addTile("M", "099", "090602", "070607", "050602", "070205")
	addTile("M", "100", "070205", "090602", "050602", "070205")
	addTile("M", "101", "090506", "100406", "060107", "060107")
	addTile("M", "102", "090506", "100105", "060107", "060107")
	addTile("M", "103", "100105", "110608", "060107", "060107")
	addTile("M", "104", "070205", "070607", "060107", "060107")
	addTile("M", "105", "060201", "100406", "060107", "110106")
	addTile("M", "106", "070408", "050305", "060107", "070607")
	addTile("M", "107", "110608", "050305", "060107", "070607")
	addTile("M", "108", "060201", "100406", "060107", "070607")
	addTile("M", "109", "100105", "060304", "060107", "070607")
	addTile("M", "110", "100406", "090506", "060107", "050204")
	addTile("M", "111", "090506", "100406", "060107", "050204")
	addTile("M", "112", "110703", "110608", "060107", "070408")
	addTile("M", "113", "110703", "100105", "060107", "070408")
	addTile("M", "114", "050204", "060201", "060107", "060304")
	addTile("M", "115", "090201", "090506", "060107", "060304")
	addTile("M", "116", "090506", "060304", "060107", "060304")
	addTile("M", "117", "070205", "100105", "060107", "060304")
	addTile("M", "118", "070408", "090602", "060107", "110608")
	addTile("M", "119", "100105", "090506", "060107", "090201")
	addTile("M", "120", "050305", "070205", "060107", "100406")
	addTile("M", "121", "050204", "070205", "060107", "060201")
	addTile("M", "122", "070607", "100105", "060107", "110703")
	addTile("M", "123", "050204", "110608", "060107", "110703")
	addTile("M", "124", "110703", "100406", "060107", "110703")
	addTile("M", "125", "090506", "070205", "060107", "050305")
	addTile("M", "126", "090602", "100406", "060107", "050305")
	addTile("M", "127", "050204", "070607", "060107", "100105")
	addTile("M", "128", "110608", "090602", "060107", "100105")
	addTile("M", "129", "060201", "070408", "060107", "100105")
	addTile("M", "130", "110703", "050305", "060107", "100105")
	addTile("M", "131", "070607", "110703", "060107", "090602")
	addTile("M", "132", "090506", "050204", "060107", "090602")
	addTile("M", "133", "050204", "100105", "060107", "070205")
	addTile("M", "134", "060304", "060201", "060107", "070205")
	addTile("M", "135", "060201", "070408", "060107", "070205")
	addTile("M", "136", "100105", "110703", "060107", "070205")
	addTile("M", "137", "110703", "090201", "110106", "110106")
	addTile("M", "138", "070607", "070408", "110106", "070607")
	addTile("M", "139", "070607", "060304", "110106", "070607")
	addTile("M", "140", "070607", "090506", "110106", "070607")
	addTile("M", "141", "110608", "090602", "110106", "070607")
	addTile("M", "142", "100406", "070607", "110106", "070607")
	addTile("M", "143", "100105", "090506", "110106", "070607")
	addTile("M", "144", "090602", "090602", "110106", "070607")
	addTile("M", "145", "070408", "060201", "110106", "050204")
	addTile("M", "146", "110106", "090506", "110106", "070408")
	addTile("M", "147", "110106", "070205", "110106", "070408")
	addTile("M", "148", "070408", "050204", "110106", "070408")
	addTile("M", "149", "100406", "090506", "110106", "070408")
	addTile("M", "150", "070607", "060304", "110106", "110608")
	addTile("M", "151", "060201", "070205", "110106", "110608")
	addTile("M", "152", "060304", "060304", "110106", "090201")
	addTile("M", "153", "060304", "110608", "110106", "090201")
	addTile("M", "154", "070205", "100105", "110106", "090201")
	addTile("M", "155", "090201", "090201", "110106", "060201")
	addTile("M", "156", "090201", "090506", "110106", "060201")
	addTile("M", "157", "070205", "090201", "110106", "060201")
	addTile("M", "158", "090201", "100105", "110106", "110703")
	addTile("M", "159", "050305", "070607", "110106", "110703")
	addTile("M", "160", "090602", "090602", "110106", "050305")
	addTile("M", "161", "070607", "110703", "110106", "100105")
	addTile("M", "162", "050204", "110703", "110106", "100105")
	addTile("M", "163", "100406", "060304", "110106", "070205")
	addTile("M", "164", "060201", "100105", "110106", "070205")
	addTile("M", "165", "070408", "060201", "070607", "050204")
	addTile("M", "166", "060201", "050305", "070607", "050204")
	addTile("M", "167", "070408", "070408", "070607", "060304")
	addTile("M", "168", "060304", "100406", "070607", "110608")
	addTile("M", "169", "110608", "090602", "070607", "110608")
	addTile("M", "170", "060201", "050305", "070607", "090201")
	addTile("M", "171", "110703", "100105", "070607", "090201")
	addTile("M", "172", "090602", "060304", "070607", "090201")
	addTile("M", "173", "100406", "100406", "070607", "100406")
	addTile("M", "174", "090506", "110703", "070607", "100406")
	addTile("M", "175", "090201", "090602", "070607", "060201")
	addTile("M", "176", "090201", "110608", "070607", "090506")
	addTile("M", "177", "050204", "060201", "070607", "110703")
	addTile("M", "178", "090506", "100105", "070607", "050305")
	addTile("M", "179", "050305", "100406", "070607", "050305")
	addTile("M", "180", "090201", "100105", "070607", "100105")
	addTile("M", "181", "060304", "100105", "070607", "090602")
	addTile("M", "182", "090201", "060304", "070607", "090602")
	addTile("M", "183", "110608", "050305", "050204", "050204")
	addTile("M", "184", "070205", "090201", "050204", "070408")
	addTile("M", "185", "110608", "090506", "050204", "060304")
	addTile("M", "186", "050305", "050305", "050204", "060304")
	addTile("M", "187", "090602", "090201", "050204", "110608")
	addTile("M", "188", "050204", "090602", "050204", "090201")
	addTile("M", "189", "070408", "110608", "050204", "090201")
	addTile("M", "190", "100105", "110608", "050204", "090201")
	addTile("M", "191", "070408", "070408", "050204", "100406")
	addTile("M", "192", "100406", "110703", "050204", "100406")
	addTile("M", "193", "060304", "090201", "050204", "060201")
	addTile("M", "194", "090602", "060304", "050204", "090506")
	addTile("M", "195", "090602", "070205", "050204", "090506")
	addTile("M", "196", "070205", "100406", "050204", "090506")
	addTile("M", "197", "060304", "070205", "050204", "110703")
	addTile("M", "198", "110608", "050305", "050204", "110703")
	addTile("M", "199", "110703", "110703", "050204", "110703")
	addTile("M", "200", "110608", "070408", "050204", "050305")
	addTile("M", "201", "090602", "050305", "050204", "100105")
	addTile("M", "202", "070205", "090201", "050204", "100105")
	addTile("M", "203", "090506", "110608", "050204", "090602")
	addTile("M", "204", "090506", "050305", "050204", "090602")
	addTile("M", "205", "100105", "070408", "050204", "090602")
	addTile("M", "206", "070205", "090602", "050204", "090602")
	addTile("M", "207", "070205", "090201", "070408", "070408")
	addTile("M", "208", "110608", "070205", "070408", "060304")
	addTile("M", "209", "050305", "100406", "070408", "060304")
	addTile("M", "210", "060304", "110608", "070408", "090201")
	addTile("M", "211", "100105", "100406", "070408", "090201")
	addTile("M", "212", "070408", "100105", "070408", "100406")
	addTile("M", "213", "050305", "050305", "070408", "060201")
	addTile("M", "214", "070408", "110703", "070408", "090506")
	addTile("M", "215", "060201", "070205", "070408", "090506")
	addTile("M", "216", "070205", "090602", "070408", "090506")
	addTile("M", "217", "110608", "100406", "070408", "110703")
	addTile("M", "218", "060201", "090506", "070408", "100105")
	addTile("M", "219", "060304", "060201", "070408", "090602")
	addTile("M", "220", "060304", "100105", "070408", "070205")
	addTile("M", "221", "090602", "090602", "070408", "070205")
	addTile("M", "222", "090602", "070205", "070408", "070205")
	addTile("M", "223", "090506", "070205", "060304", "060304")
	addTile("M", "224", "110703", "090201", "060304", "060304")
	addTile("M", "225", "100105", "100406", "060304", "060304")
	addTile("M", "226", "110608", "100406", "060304", "090201")
	addTile("M", "227", "090506", "090506", "060304", "090201")
	addTile("M", "228", "110608", "050305", "060304", "060201")
	addTile("M", "229", "090201", "070205", "060304", "110703")
	addTile("M", "230", "100105", "050305", "060304", "110703")
	addTile("M", "231", "090506", "110703", "060304", "050305")
	addTile("M", "232", "090506", "090602", "060304", "050305")
	addTile("M", "233", "110608", "110703", "110608", "110608")
	addTile("M", "234", "100105", "060201", "110608", "090201")
	addTile("M", "235", "090201", "060201", "110608", "060201")
	addTile("M", "236", "090201", "110703", "110608", "060201")
	addTile("M", "237", "090506", "100406", "110608", "060201")
	addTile("M", "238", "060201", "100105", "110608", "090506")
	addTile("M", "239", "100406", "070205", "110608", "110703")
	addTile("M", "240", "090602", "100406", "110608", "110703")
	addTile("M", "241", "090201", "090602", "110608", "050305")
	addTile("M", "242", "060201", "090602", "110608", "050305")
	addTile("M", "243", "100406", "090506", "090201", "100406")
	addTile("M", "244", "110703", "050305", "090201", "100406")
	addTile("M", "245", "070205", "110703", "090201", "060201")
	addTile("M", "246", "100105", "070205", "090201", "090602")
	addTile("M", "247", "090602", "070205", "100406", "100406")
	addTile("M", "248", "090506", "060201", "100406", "060201")
	addTile("M", "249", "060201", "090602", "100406", "090506")
	addTile("M", "250", "100105", "090602", "100406", "110703")
	addTile("M", "251", "070205", "050305", "060201", "060201")
	addTile("M", "252", "050305", "090506", "060201", "090506")
	addTile("M", "253", "100105", "110703", "090506", "050305")
	addTile("M", "254", "090602", "100105", "110703", "100105")
	addTile("M", "255", "100105", "070205", "110703", "070205")
	addTile("M", "256", "090602", "070205", "050305", "070205")
}

// addTile adds a tile created from the given typeOf, number and the side definitions.
// Side definitions are given as strings containing 6 digits.
func addTile(typeOf string, number string, top string, right string, bottom string, left string) {
	var err error

	tile := new(Tile1)

	// set the properties
	tile.TypeOf, err = getTypeOfFrom(typeOf)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	tile.Number, err = getNumberFrom(number)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// tiles always start out north
	tile.Orientation = ORIENTATION_NORTH

	tile.Top, err = getSideFrom(top)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	tile.Right, err = getSideFrom(right)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	tile.Bottom, err = getSideFrom(bottom)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	tile.Left, err = getSideFrom(left)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// store the four orientations tiles
	for i := 0; i < 4; i++ {
		TileSet[tile.Orientation][tile.Number] = tile
		tile = tile.Clone()
		tile.Rotate()
	}
}

// getTypeOfFrom returns the byte representation for the given string representation of
// the type of tile. C=Corner, E=Edge M=Middle
// Returns a byte containg the type of tile and an error if the string representation is
// not recognised.
func getTypeOfFrom(typeOf string) (byte, error) {
	var err error
	var result = TYPEOF_BORDER

	switch typeOf {
	case "C":
		result = TYPEOF_CORNER
	case "E":
		result = TYPEOF_EDGE
	case "M":
		result = TYPEOF_MIDDLE
	default:
		err = errors.New(fmt.Sprintf("TypeOf: %s Not recognised", typeOf))
	}

	return result, err
}

// getNumberFrom converts the string representation of the tile Id to a byte value - 1.
// Returns the byte tile number and an error if the string contains a number that cannot
// be converted.
func getNumberFrom(number string) (byte, error) {
	var result byte
	intNumber, err := strconv.Atoi(number)
	if err == nil {
		result = byte(intNumber - 1)
	}
	return result, err
}

// getSideFrom converts the string representation of a tile side to the byte representation.
// Returns the byte representation of the side and an error if the side cannot be created.
func getSideFrom(sideDef string) (byte, error) {
	var err error

	result, ok := DefinitionToSide(sideDef)
	if !ok {
		err = errors.New(fmt.Sprintf("SideDef %s is not recognised.", sideDef))
	}

	return result, err
}

// function clones a tile
func (tile *Tile1) Clone() *Tile1 {
	result := new(Tile1)

	result.Number = tile.Number
	result.TypeOf = tile.TypeOf
	result.Orientation = tile.Orientation

	result.Top = tile.Top
	result.Left = tile.Left
	result.Bottom = tile.Bottom
	result.Right = tile.Right

	result.Energy = tile.Energy

	return result
}

// function rotates the tile 90 degees clockwise
func (tile *Tile1) Rotate() {
	tile.Orientation = rotateOrientation(tile.Orientation)

	// remember the current top
	currentTop := tile.Top

	tile.Top = tile.Left
	tile.Left = tile.Bottom
	tile.Bottom = tile.Right
	tile.Right = currentTop
}

// function rotates the tile until it has the required orientation
func (tile *Tile1) RotateTill(orientation byte) {
	for tile.Orientation != orientation {
		tile.Rotate()
	}
}

// Function returns two slices containing the sides that match and those that don't between the tile and the rhs tile.
func (tile *Tile1) SidesMatching(rhs *Tile1) ([]string, []string) {
	var matching []string
	var nonMatching []string

	side, _ := SideToDefinition(tile.Top)
	if tile.Top == rhs.Top {
		matching = append(matching, side)
	} else {
		// store the side for the tile
		nonMatching = append(nonMatching, side)
		// store the side for the rhs tile
		side, _ := SideToDefinition(rhs.Top)
		nonMatching = append(nonMatching, side)
	}

	side, _ = SideToDefinition(tile.Right)
	if tile.Right == rhs.Right {

		matching = append(matching, side)
	} else {
		// store the side for the tile
		nonMatching = append(nonMatching, side)
		// store the side for the rhs tile
		side, _ := SideToDefinition(rhs.Top)
		nonMatching = append(nonMatching, side)
	}

	side, _ = SideToDefinition(tile.Bottom)
	if tile.Bottom == rhs.Bottom {

		matching = append(matching, side)
	} else {
		// store the side for the tile
		nonMatching = append(nonMatching, side)
		// store the side for the rhs tile
		side, _ := SideToDefinition(rhs.Top)
		nonMatching = append(nonMatching, side)
	}

	side, _ = SideToDefinition(tile.Left)
	if tile.Left == rhs.Left {

		matching = append(matching, side)
	} else {
		// store the side for the tile
		nonMatching = append(nonMatching, side)
		// store the side for the rhs tile
		side, _ := SideToDefinition(rhs.Top)
		nonMatching = append(nonMatching, side)
	}

	return matching, nonMatching

}

// Function returns the number of sides that match a slice of side definitions
func (tile *Tile1) NumberOfSidesMatchingDefinitions(definitions []string) int {
	result := 0

	side, _ := SideToDefinition(tile.Top)
	if util.Contains(definitions, side) {
		result++
	}

	side, _ = SideToDefinition(tile.Right)
	if util.Contains(definitions, side) {
		result++
	}

	side, _ = SideToDefinition(tile.Bottom)
	if util.Contains(definitions, side) {
		result++
	}

	side, _ = SideToDefinition(tile.Left)
	if util.Contains(definitions, side) {
		result++
	}

	return result
}

// Function returns the number of sides that match between two tiles.
func (tile *Tile1) NumberOfSidesMatchingTile(rhsTile *Tile1) int {
	definitions := rhsTile.SideDefinitions()
	result := tile.NumberOfSidesMatchingDefinitions(definitions)
	return result
}

// Function returns true if the two tiles can connect.
func (tile *Tile1) CanConnectTo(rhsTile *Tile1) bool {
	if tile.Number == rhsTile.Number {
		return false
	}
	lhsT := tile.TypeOf
	rhsT := rhsTile.TypeOf

	result := lhsT == TYPEOF_CORNER && rhsT == TYPEOF_EDGE
	result = result || (lhsT == TYPEOF_EDGE && (rhsT == TYPEOF_CORNER || rhsT == TYPEOF_EDGE || rhsT == TYPEOF_MIDDLE))
	result = result || (lhsT == TYPEOF_MIDDLE && (rhsT == TYPEOF_EDGE || rhsT == TYPEOF_MIDDLE))
	return result
}

// Function obtains the side definitions from a tile
// Returns a slice containing the definitions (Edge Sides are not included)
func (tile *Tile1) SideDefinitions() []string {
	var definitions []string

	side, _ := SideToDefinition(tile.Top)
	if !util.Contains(definitions, side) &&
		side != EDGE_SIDE_DEF {
		definitions = append(definitions, side)
	}

	side, _ = SideToDefinition(tile.Right)
	if !util.Contains(definitions, side) &&
		side != EDGE_SIDE_DEF {
		definitions = append(definitions, side)
	}

	side, _ = SideToDefinition(tile.Bottom)
	if !util.Contains(definitions, side) &&
		side != EDGE_SIDE_DEF {
		definitions = append(definitions, side)
	}

	side, _ = SideToDefinition(tile.Left)
	if !util.Contains(definitions, side) &&
		side != EDGE_SIDE_DEF {
		definitions = append(definitions, side)
	}

	return definitions
}

// function rotates the tile 90 degees clockwise
func rotateOrientation(orientation byte) byte {
	result := ORIENTATION_UNKNOWN
	switch orientation {
	case ORIENTATION_NORTH:
		result = ORIENTATION_EAST
	case ORIENTATION_EAST:
		result = ORIENTATION_SOUTH
	case ORIENTATION_SOUTH:
		result = ORIENTATION_WEST
	case ORIENTATION_WEST:
		result = ORIENTATION_NORTH
	default:
		panic(fmt.Sprintf("Attempted to Rotate Illegal Orientation %c", orientation))
	}

	return result
}

// function Marshals the tile to a byte array
func (tile *Tile1) Marshal() []byte {
	// make a slice for the bytes
	result := make([]byte, 2)
	result[0] = tile.Number
	result[1] = tile.Orientation

	return result
}

// function Unmarshals the given byte array to an ITile
func (tile *Tile1) Unmarshal(bytes []byte) *Tile1 {
	// get the tile from the tile set by number byte and orientation byte
	number := bytes[0]
	orientation := bytes[1]
	ref := TileSet[orientation][number]

	// clone it
	result := ref.Clone()
	return result
}

// function returns the unique id that Identifies this tile.
func (tile *Tile1) GetNumberAsString() string {
	number := 1 + int(tile.Number)
	result := fmt.Sprintf("%0.3d", number)
	return result
}

// function returns the orianation of this tile as a string.
func (tile *Tile1) GetOrientationAsString() string {
	result := "X"
	switch tile.Orientation {
	case ORIENTATION_NORTH:
		result = "N"
	case ORIENTATION_EAST:
		result = "E"
	case ORIENTATION_SOUTH:
		result = "S"
	case ORIENTATION_WEST:
		result = "W"
	default:
		panic(fmt.Sprintf("Unknown orientation %c", tile.Orientation))
	}

	return result
}

// function returns the unique id that Identifies this tile.
func (tile *Tile1) GetId() string {
	bytes := tile.Marshal()
	result := fmt.Sprintf("%X", bytes)
	return result
}

// function returns the Top Pattern of this tile (The string formed by reading the
// sides from the tile accross the top, left to right).
func (tile *Tile1) GetTopPattern() string {
	result := fmt.Sprintf("%c", tile.Top)
	return result
}

// function returns the Right Pattern of this tile (The string formed by reading the
// sides from the tile down the right hand side, top to bottom).
func (tile *Tile1) GetRightPattern() string {
	result := fmt.Sprintf("%c", tile.Right)
	return result
}

// function returns the Bottom Pattern of this tile (The string formed by reading the
// sides from the tile accross the top, left to right).
func (tile *Tile1) GetBottomPattern() string {
	result := fmt.Sprintf("%c", tile.Bottom)
	return result
}

// function returns the Left Pattern of this tile (The string formed by reading the
// sides from the tile down the left hand side, top to bottom).
func (tile *Tile1) GetLeftPattern() string {
	result := fmt.Sprintf("%c", tile.Left)
	return result
}

// function returns the string representation of this tile
func (tile *Tile1) String() string {
	var result string
	for i := 0; i < 5; i++ {
		result += fmt.Sprintf("%s\n", tile.Line(i))
	}
	return result
}

// function returns the string which is the n line of the text representation of the tile
func (tile *Tile1) Line(index int) string {
	var result string

	if index < 0 || index > 4 {
		panic("Index out of range")
	}
	switch index {
	case 0:
		result = tile.line0()
	case 1:
		result = tile.line1()
	case 2:
		result = tile.line2()
	case 3:
		result = tile.line3()
	case 4:
		result = tile.line4()
	}

	return result
}

// return the string represention of the line of the tile
// "+--X--+"
func (tile *Tile1) line0() string {
	result := fmt.Sprintf("+--%c--+", getSideChar(tile.Top))
	return result
}

// return the string represention of the line of the tile
// "|  *  |"
func (tile *Tile1) line1() string {
	north := ' '
	// check the orientation of the tile
	if tile.Orientation == ORIENTATION_NORTH {
		north = '*'
	}

	result := fmt.Sprintf("|  %c  |", north)
	if tile.TypeOf == TYPEOF_BORDER {
		result = "|     |"
	}

	return result
}

// return the string represention of the line of the tile
// "X*999*X"
func (tile *Tile1) line2() string {
	west := ' '
	east := ' '

	if tile.Orientation == ORIENTATION_EAST {
		east = '*'
	}

	if tile.Orientation == ORIENTATION_WEST {
		west = '*'
	}

	result := fmt.Sprintf("%c%c%s%c%c", getSideChar(tile.Left), west,
		tile.GetNumberAsString(), east, getSideChar(tile.Right))
	if tile.TypeOf == TYPEOF_BORDER {
		result = fmt.Sprintf("%c     %c", getSideChar(tile.Left), getSideChar(tile.Right))
	}

	return result
}

// return the string represention of the line of the tile
// "|  *  |"
func (tile *Tile1) line3() string {
	south := ' '
	// check the orientation of the tile
	if tile.Orientation == ORIENTATION_SOUTH {
		south = '*'
	}

	result := fmt.Sprintf("|  %c  |", south)
	if tile.TypeOf == TYPEOF_BORDER {
		result = "|     |"
	}

	//result = fmt.Sprintf("| %0.3d |", tile.Energy)

	return result
}

// return the string represention of the line of the tile
// "+--X--+"
func (tile *Tile1) line4() string {
	result := fmt.Sprintf("+--%c--+", getSideChar(tile.Bottom))
	return result
}

// get the character that should be used for the side
// Returns 'X' when the side char == byte(0)
// the side char otherwise
func getSideChar(sideChar byte) byte {
	result := sideChar
	if result == byte(0) {
		result = byte('X')
	}
	return result
}
