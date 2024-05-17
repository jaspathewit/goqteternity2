package ui
//
//import (
//	"errors"
//	"fmt"
//	"gopkg.in/qml."
//	"image"
//	"image/png"
//	"os"
//)
//
//// A struct containing the UI's Control. The struct acts as the bridge between
//// the GO code and the qml code.
//type Control struct {
//	Root qml.Object
//}
//
//// The control for the UI
//var ctrl Control
//
//// The main window containing the UI
//var win *qml.Window
//
//// The chanel to which the UI listens for events
//var EventChannel chan string
//
//// Initialize the UI
//func Initalize() error {
//
//	// create the qml engine
//	engine := qml.NewEngine()
//
//	// add an image provider for jpeg to the engine
//	engine.AddImageProvider("png", loadPngImage)
//
//	// load the ui definition
//	component, err := engine.LoadFile("eternity2UI.qml")
//	if err != nil {
//		return err
//	}
//
//	// Create a new control for the ui
//	ctrl = Control{}
//
//	// Create an event channel
//	EventChannel = make(chan<- string)
//
//	// get the engine qt context
//	context := engine.Context()
//
//	// set a varable into the qml context with the control
//	context.SetVar("ctrl", &ctrl)
//
//	// create the window
//	win = component.CreateWindow(nil)
//
//	//defer win.Destroy()
//
//	// add the window root to the control
//	ctrl.Root = win.Root()
//
//	// show the created window
//	//win.Show()
//
//	// load the text file
//	//go ctrl.loadImages()
//
//	// wait for exit
//	//win.Wait()
//
//	// no errors
//	return nil
//}
//
//// method loads the given png image
//func loadPngImage(id string, width, height int) image.Image {
//	// open the given file
//	f, err := os.Open(id)
//	if err != nil {
//		panic(err)
//	}
//	// close it later
//	defer f.Close()
//
//	// create an image from the png file
//	image, err := png.Decode(f)
//	if err != nil {
//		panic(err)
//	}
//
//	return image
//}
//
//// function shows the window only returns when the window is closed
//func Start() error {
//	if win == nil {
//		err := errors.New("Window has not been created. Has ui.Initilize() been called?")
//		return err
//	}
//
//	defer win.Destroy()
//	win.Show()
//
//	// wait for exit
//	win.Wait()
//
//	return nil
//}
//
//// function handles events that arrive on the event channel
//func handleEvent() error {
//	if win == nil || ctrl == nil || eventChannel == nil {
//		err := errors.New("UI has not been created. Has ui.Initilize() been called?")
//		return err
//	}
//
//	// start handeling events that arrive on the event channel
//	go handleEvent(eventChannel)
//
//	defer win.Destroy()
//	win.Show()
//
//	// wait for exit
//	win.Wait()
//
//	return nil
//}
//
//// method called when the find button is clicked
////func (ctrl *Control) FindButtonClicked(sender qml.Object) {
////	textField := ctrl.Root.ObjectByName("textField")
////	findStr := textField.String("text")
////	fmt.Printf("Find Button was clicked: %s", findStr)
//
////	// find the position of the matching string
////	findPos := strings.Index(ctrl.TextArea_Text, findStr)
//
////	fmt.Printf("Found Text at: %d", findPos)
//
////	// if the text was found
////	if findPos != -1 {
////		textArea := ctrl.Root.ObjectByName("textArea")
////		// set the cursor position to the position
////		textArea.Set("cursorPosition", findPos)
////		textArea.Call("moveCursorSelection", findPos+len(findStr), 0)
//
////		textArea.Set("focus", true)
////	}
////}
//
//// method called when the open button is clicked
////func (ctrl *Control) OpenButtonClicked(sender qml.Object) {
////	fmt.Printf("Open Button was clicked")
//
////	fileDialog := ctrl.Root.ObjectByName("fileDialog")
////	fileDialog.Set("visible", true)
////}
//
//// method called when the file dialog is accepted
////func (ctrl *Control) FindDialogAccepted(sender qml.Object) {
////	rawurl := sender.String("fileUrl")
////	fileUrl, _ := url.Parse(rawurl)
////	filePath, _ := asFilepath(fileUrl)
////	fmt.Printf("Find dialog was accepted %s", filePath)
////	go ctrl.loadTextFile(filePath)
////}
//
//// method called when the file dialog is accepted
////func (ctrl *Control) FindDialogRejected(sender qml.Object) {
////	fmt.Printf("Find dialog was rejected")
////}
//
//func (ctrl *Control) loadImages() {
//
//	// get the grid
//	// grid := ctrl.Root.ObjectByName("grid")
//
//	tileNumber := 1
//
//	for row := 1; row < 17; row++ {
//
//		for col := 1; col < 17; col++ {
//
//			// create the string that identifies the
//			// qmlImage in the ui
//			qmlId := fmt.Sprintf("tile_%d_%d", row, col)
//
//			// get the image component
//			image := ctrl.Root.ObjectByName(qmlId)
//
//			// create the string that identifies the tile source
//			source := fmt.Sprintf("image://png/tiles/Tile-%03dN.png", tileNumber)
//
//			// change the tiles source
//			image.Set("source", source)
//
//			tileNumber++
//		}
//	}
//}
