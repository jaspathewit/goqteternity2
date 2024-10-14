package mvvm

import (
	log "github.com/cihub/seelog"
	"gopkg.in/qml.v1"
)

// A struct containing an event
type Event struct {
	Id   int
	Data interface{}
}

// A struct containing a reference control
type ViewModel struct {
	Root      qml.Object
	Context   *qml.Context
	Component qml.Object
	Engine    *qml.Engine

	EventChannel  chan Event
	EventHandlers map[int]func(*ViewModel, interface{})
}

// Method creates a new Mvvm from the qml file
func New(filename string) (*ViewModel, error) {
	// create the qml engine
	engine := qml.NewEngine()

	// load the ui definition from the qml file
	component, err := engine.LoadFile(filename)
	if err != nil {
		return nil, err
	}

	context := engine.Context()

	viewModel := ViewModel{Component: component,
		Context: context,
		Engine:  engine}

	//context.SetVar("viewData", viewData)

	// create the window
	//win := component.CreateWindow(nil)

	// add the window root to the viewModel
	//viewModel.Root = win.Root()

	return &viewModel, nil

}

// Method creates a new Mvvm from the qml file
func (viewModel *ViewModel) NewMainWindow() *qml.Window {

	// create the window
	win := viewModel.Component.CreateWindow(nil)

	// add the window root to the viewModel
	viewModel.Root = win.Root()

	return win
}

// Method sets an event channel on the view
func (viewModel *ViewModel) SetEventChannel(eventChannel chan Event) {
	viewModel.EventChannel = eventChannel
}

// method Handles an events arrival at the view model
func (viewModel *ViewModel) HandleEvent(event Event) {
	// get the handler from the map
	eventHandler, ok := viewModel.EventHandlers[event.Id]
	if !ok {
		log.Debugf("Handler for event not registered")
	}

	// call the event handler
	eventHandler(viewModel, event.Data)
}

// Register an event handler
func (viewModel *ViewModel) RegisterEventHandler(id int, eventHandler func(*ViewModel, interface{})) {
	// create the map if we don't have it already
	if viewModel.EventHandlers == nil {
		viewModel.EventHandlers = make(map[int]func(*ViewModel, interface{}))
	}

	// put the event handle into the map
	viewModel.EventHandlers[id] = eventHandler
}
