// Color and Shapes used by the Eternity tile sides.
package model

// Colour types
const (
	COLOR_EDGE       string = "00"
	COLOR_LIGHT_BLUE        = "01"
	COLOR_PINK              = "02"
	COLOR_ORANGE            = "03"
	COLOR_GREEN             = "04"
	COLOR_DARK_BLUE         = "05"
	COLOR_YELLOW            = "06"
	COLOR_PURPLE            = "07"
	COLOR_BROWN             = "08"
)

// Shape types
const (
	SHAPE_EDGE          string = "00"
	SHAPE_MALTES_CROSS         = "01"
	SHAPE_SANDWICH             = "02"
	SHAPE_NUT                  = "03"
	SHAPE_FLOWER               = "04"
	SHAPE_CLUB                 = "05"
	SHAPE_SWORD_CROSS          = "06"
	SHAPE_EHBO                 = "07"
	SHAPE_PICTURE_FRAME        = "08"
	SHAPE_CASTLE               = "09"
	SHAPE_SQUARE               = "10"
	SHAPE_STAR                 = "11"
)

// Edge side definition (conveniance for 000000)
var EDGE_SIDE_DEF string

// Map maps the six char string side definition to the byte used to represent the
// side internally
var defToSideMap map[string]byte

// Map maps the byte used internally to represent a side to the
// six char string definition
var sideToDefMap map[byte]string

// initSide initilizes the side data structures
func initSide() {
	if defToSideMap == nil || sideToDefMap == nil {
		
		// Define the conveniance 000000
		EDGE_SIDE_DEF = SHAPE_EDGE + COLOR_EDGE + COLOR_EDGE

		defToSideMap = make(map[string]byte)
		sideToDefMap = make(map[byte]string)

		// add the definitions
		addSideDefinition(SHAPE_EDGE, COLOR_EDGE, COLOR_EDGE, byte(0))
		addSideDefinition(SHAPE_MALTES_CROSS, COLOR_LIGHT_BLUE, COLOR_ORANGE, byte('A'))
		addSideDefinition(SHAPE_SANDWICH, COLOR_LIGHT_BLUE, COLOR_PINK, byte('B'))
		addSideDefinition(SHAPE_NUT, COLOR_DARK_BLUE, COLOR_GREEN, byte('C'))
		addSideDefinition(SHAPE_FLOWER, COLOR_YELLOW, COLOR_DARK_BLUE, byte('D'))
		addSideDefinition(SHAPE_CLUB, COLOR_PINK, COLOR_GREEN, byte('E'))
		addSideDefinition(SHAPE_CLUB, COLOR_ORANGE, COLOR_DARK_BLUE, byte('F'))
		addSideDefinition(SHAPE_CLUB, COLOR_DARK_BLUE, COLOR_ORANGE, byte('G'))
		addSideDefinition(SHAPE_CLUB, COLOR_YELLOW, COLOR_PINK, byte('H'))
		addSideDefinition(SHAPE_SWORD_CROSS, COLOR_LIGHT_BLUE, COLOR_PURPLE, byte('I'))
		addSideDefinition(SHAPE_SWORD_CROSS, COLOR_PINK, COLOR_LIGHT_BLUE, byte('J'))
		addSideDefinition(SHAPE_SWORD_CROSS, COLOR_ORANGE, COLOR_GREEN, byte('K'))
		addSideDefinition(SHAPE_EHBO, COLOR_PINK, COLOR_DARK_BLUE, byte('L'))
		addSideDefinition(SHAPE_EHBO, COLOR_GREEN, COLOR_BROWN, byte('M'))
		addSideDefinition(SHAPE_EHBO, COLOR_YELLOW, COLOR_PURPLE, byte('N'))
		addSideDefinition(SHAPE_PICTURE_FRAME, COLOR_ORANGE, COLOR_BROWN, byte('O'))
		addSideDefinition(SHAPE_CASTLE, COLOR_PINK, COLOR_LIGHT_BLUE, byte('P'))
		addSideDefinition(SHAPE_CASTLE, COLOR_DARK_BLUE, COLOR_YELLOW, byte('Q'))
		addSideDefinition(SHAPE_CASTLE, COLOR_YELLOW, COLOR_PINK, byte('R'))
		addSideDefinition(SHAPE_SQUARE, COLOR_LIGHT_BLUE, COLOR_DARK_BLUE, byte('S'))
		addSideDefinition(SHAPE_SQUARE, COLOR_GREEN, COLOR_YELLOW, byte('T'))
		addSideDefinition(SHAPE_STAR, COLOR_LIGHT_BLUE, COLOR_YELLOW, byte('U'))
		addSideDefinition(SHAPE_STAR, COLOR_YELLOW, COLOR_BROWN, byte('V'))
		addSideDefinition(SHAPE_STAR, COLOR_PURPLE, COLOR_ORANGE, byte('W'))
	}
}

// addSideDefinition adds the side definition and the byte used for the internal represention
// to the definition to side and side to definition maps.
func addSideDefinition(shape string, fgColor string, bgColor string, representation byte) {
	def := shape + fgColor + bgColor

	defToSideMap[def] = representation
	sideToDefMap[representation] = def
}

// DefinitionToSide returns the internal representation of a side for the given definition string.
// Definition strings are of the form AABBCC whre AA is the Shape, BB is the Foreground color
// and CC is the back ground color.
// Returns the representation of the Side and true if the definition is valid, 0 and false if the
// definition is invalid.
func DefinitionToSide(definition string) (byte, bool) {
	// ensure that the sides are initilized
	initSide()
	result, ok := defToSideMap[definition]
	return result, ok
}

// SideToDefinition returns the side definition from an internal representation.
// Definition strings are of the form AABBCC whre AA is the Shape, BB is the Foreground color
// and CC is the back ground color.
// Returns the definition of the Side and true if the side is valid, 0 and false if the
// side is invalid.
func SideToDefinition(side byte) (string, bool) {
	// ensure that the sides are initilized
	initSide()
	result, ok := sideToDefMap[side]
	return result, ok
}
