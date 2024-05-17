// Package view contains the view data structs for interfacing with the qml view.
package view

import (
	"bitbucket.com/jaspathewit/eternity2/model"
	"bitbucket.com/jaspathewit/mvvm"
	log "github.com/cihub/seelog"
	"image"
	"image/png"
	"fmt"
	"os"
)

// A struct containing the ViewData
type Data struct {
	Message string
}

var ViewData Data
var PModel *model.Model

//// Event handlers
//// method Handles the initilisation of the model
//func HandleModelInit(viewModel *mvvm.ViewModel, data interface{}) {
//	log.Debugf("HandleModelInit")
//	message := data.(string)
//	ViewData.Message = message
//
//	// tell qml that the view data has changed
//	qml.Changed(&ViewData, &ViewData.Message)
//	log.Debugf("UI Updated")
//}

// create a map index by image name for the tile images
var imageCache = make(map[string]image.Image)

// HandleSurfaceUpdate the handler to handle when the surface is updated
func HandleSurfaceUpdate(viewModel *mvvm.ViewModel, data interface{}) {
	log.Debugf("HandleMessageUpdate")
	
	surface := data.(*model.Surface)
	
	// loop through the cells of the grid
	for row := 1; row < 17; row++ {
		for col := 1; col < 17; col++ {

			// create the string that identifies the
			// qmlImage in the ui
			qmlId := fmt.Sprintf("tile_%d_%d", row, col)

			// get the image component
			imageTile := viewModel.Root.ObjectByName(qmlId)

			// create the string that identifies the tile source
			tile := surface.GetTile(row, col)
			source := fmt.Sprintf("image://png/ui/tiles/Tile-%s%s.png", 
				tile.GetNumberAsString(),
				tile.GetOrientationAsString())

			// change the tiles source
			log.Debugf("Setting image source %s", source)
			imageTile.Set("source", source)
		}
	}
}
	
	
//	ViewData.Message = message
//
//	// tell qml that the view data has changed
//	qml.Changed(&ViewData, &ViewData.Message)
//	log.Debugf("UI Updated")
//}


// method called when the button is clicked
//func (data *Data) ButtonClicked(button qml.Object) {
//	log.Debugf("Button was clicked")
//	PModel.SetMessage("Hello Jason")
//	log.Debugf("SetMessage Completed")
//}

// loadPngImage loads the given png image called by the QT ui.
func LoadPngImage(source string, width, height int) image.Image {
	log.Debugf("Loading Image %s", source)
	result, ok := imageCache[source]
	
	// check if image was found in the cache
	if !ok {
		log.Debugf("Image %s Not Found in cache", source)
		result = loadPngImageFromFile(source)
		imageCache[source] = result
	}
	
	return result
}

// loadPngImage loads the given png image from the source.
func loadPngImageFromFile(source string) image.Image {
	// open the given file
	f, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	// close it later
	defer f.Close()

	// create an image from the png file
	image, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	return image
}
