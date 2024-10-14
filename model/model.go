package model

import (
	"bitbucket.com/jaspathewit/eternity2/event"
	"bitbucket.com/jaspathewit/mvvm"
	log "github.com/cihub/seelog"
)

// A struct containing the model
type Model struct {
	ViewModel *mvvm.ViewModel
	Surface   *Surface

	// the process for processing the model
	process func(pmodel *Model) error
}

// var PModel *Model

// Method handles the creation of the model with the given Processor callback.
func New(process func(pmodel *Model) error) (*Model, error) {
	//create the event channel
	//eventChannel := make(chan mvvm.Event, 5)

	// create the model
	model := Model{process: process}

	return &model, nil
}

// Process starts the processor working on the model
func (model *Model) Process() {
	log.Debugf("Starting processing the model")
	go model.process(model)
	log.Debugf("Processing the model")
}

// method Sets a pointer to a surface in the model
func (model *Model) SetSurface(surface *Surface) {
	log.Debugf("Changing the surface")
	// change the model
	model.Surface = surface

	// send the event to the UI
	event := mvvm.Event{Id: event.SurfaceUpdate,
		Data: model.Surface}

	log.Debugf("Send the SetSurface event")
	//model.EventChannel <- event
	model.ViewModel.HandleEvent(event)
	log.Debugf("Event sent")
}
