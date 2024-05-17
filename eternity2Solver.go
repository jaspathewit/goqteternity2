// eternity2Solver
package main

import (
	"bitbucket.com/jaspathewit/eternity2/event"
	"bitbucket.com/jaspathewit/eternity2/model"
	"bitbucket.com/jaspathewit/eternity2/option"
	"bitbucket.com/jaspathewit/eternity2/simulatedAnnealing"
	"bitbucket.com/jaspathewit/eternity2/view"
	"bitbucket.com/jaspathewit/mvvm"
	"gopkg.in/qml.v1"

	// "bitbucket.com/jaspathewit/eternity2/solver"
	"fmt"
	"os"

	log "github.com/cihub/seelog"
)

// Main entry point
func main() {

	// initialize the logging framework
	logger, err := log.LoggerFromConfigAsFile("eternity2Log4go.xml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	log.ReplaceLogger(logger)
	defer log.Flush()

	// get the options
	if err = option.GetOptions(); err != nil {
		log.Errorf("error: %v\n", err)

		// return the error code
		os.Exit(1)
	}

	log.Debugf("Processing with options:\n%s", option.String())

	// get the qml to run the given method
	if err := qml.Run(run); err != nil {
		// report and errors form the qml runtime
		log.Errorf("error: %v\n", err)
		// return the error code
		os.Exit(1)
	}

	// run main internal so that defer methods will be called on exit os.Exit can
	// be used to return an error code to the OS.
	//	err = mainInternal()
	//	if err != nil {
	//		os.Exit(1)
	//	}
}

// the worker method for the qml runtime
func run() error {

	// create the viewModel
	viewModel, err := mvvm.New("ui/eternity2UI.qml")
	if err != nil {
		log.Errorf("error: %v\n", err)
		return err
	}

	// set the ViewData
	viewModel.Context.SetVar("viewData", &view.ViewData)

	// add a Image provider for png files viewModel
	viewModel.Engine.AddImageProvider("png", view.LoadPngImage)

	// register the event handlers
	viewModel.RegisterEventHandler(event.SurfaceUpdate, view.HandleSurfaceUpdate)

	// create the model and make it available in the package
	theModel, err := model.New(simulatedAnnealing.Process)
	if err != nil {
		log.Errorf("error: %v\n", err)
		return err
	}

	theModel.ViewModel = viewModel

	// set the event channel
	// viewModel.SetEventChannel(model.PModel.EventChannel)

	// create the view
	window := viewModel.NewMainWindow()

	// show the created window
	window.Show()

	// start the View model Handling events
	// go viewModel.HandleEvents()

	// initilize the model
	// model.PModel.Init()

	// start running the simulated annealing
	go theModel.Process()

	// wait for exit
	window.Wait()

	// perform the clean up
	log.Debugf("Application is closing")

	// no errors
	return nil
}

//// mainInternal internal entry point
//func mainInternal() error {
//
//	log.Info("Started")
//
//	//	if err := ui.NewControl(); err != nil {
//	//		log.Errorf(os.Stderr, "error: %v\n", err)
//	//		return err
//	//	}
//	//
//	//	// close the window on exit
//	//	defer ctrl.Window.Destroy()
//	//
//	//	// show the created window
//	//	ctrl.Window.Show()
//	//
//	//	// load the text file
//	//	//go ctrl.loadImages()
//	//
//	//	// wait for exit
//	//	ctrl.Window.Wait()
//
//	// start processing
//	err := simulatedAnnealing.Process()
//
//	return err
//}
