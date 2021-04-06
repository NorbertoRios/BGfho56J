package watchdog

import (
	"geometris-go/core/interfaces"
	"time"
)

//New ...
func New(_task interfaces.ITask, _state interfaces.ITaskState, _duration int) interfaces.IWatchdog {
	return &Watchdog{
		task:     _task,
		state:    _state,
		duration: _duration,
		stopChan: make(chan struct{}),
	}
}

//Watchdog ...
type Watchdog struct {
	task     interfaces.ITask
	state    interfaces.ITaskState
	stopChan chan struct{}
	duration int
}

//Start ...
func (w *Watchdog) Start() {
	go func() {
		select {
		case <-time.After(time.Duration(w.duration) * time.Second):
			{
				w.task.ChangeState(w.state)
				return
			}
		case <-w.stopChan:
			{
				return
			}
		}
	}()
}

//Stop ...
func (w *Watchdog) Stop() {
	w.stopChan <- struct{}{}
}
