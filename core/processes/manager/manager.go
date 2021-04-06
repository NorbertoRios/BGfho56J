package manager

import (
	"context"
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/configuration"
	"geometris-go/core/processes/immobilizer"
	"geometris-go/core/processes/location"
	"geometris-go/core/processes/message"
	"geometris-go/core/processes/synchronization"
	"geometris-go/logger"
	"sync"
)

//BuildProcesses ...
func BuildProcesses(_syncParam string) interfaces.IProcesses {
	manager := &Processes{
		mutex:     &sync.Mutex{},
		processes: make(map[string]interfaces.IProcess),
		paused:    make(map[string]int),
	}
	if _syncParam == "" {
		manager.Synchronization()
	} else {
		manager.Location(_syncParam)
	}

	return manager
}

//Processes ...
type Processes struct {
	processes map[string]interfaces.IProcess
	mutex     *sync.Mutex
	paused    map[string]int
}

//Configuration ...
func (p *Processes) Configuration() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "configuration"
	if _, f := p.paused[key]; f {
		logger.Logger().WriteToLog(logger.Info, "[Processes | Configuration] Process "+key+" is paused")
		return nil
	}
	if v, f := p.processes[key]; f {
		return v
	}
	process := configuration.New()
	p.processes[key] = process
	return process
}

//LocationRequest ...
func (p *Processes) LocationRequest() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "location_request"
	if _, f := p.paused[key]; f {
		logger.Logger().WriteToLog(logger.Info, "[Processes | LocationRequest] Process "+key+" is paused")
		return nil
	}
	if v, f := p.processes[key]; f {
		return v
	}
	process := location.New()
	p.processes[key] = process
	return process
}

//Location ...
func (p *Processes) Location(_syncParam string) interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "location"
	if _, f := p.paused[key]; f {
		logger.Logger().WriteToLog(logger.Info, "[Processes | Location] Process "+key+" is paused")
		return nil
	}
	if v, f := p.processes[key]; f {
		return v
	}
	process := message.New(_syncParam)
	p.processes[key] = process
	return process
}

//Synchronization ...
func (p *Processes) Synchronization() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "synch"
	if _, f := p.paused[key]; f {
		logger.Logger().WriteToLog(logger.Info, "[Processes | Synchronization] Process "+key+" is paused")
		return nil
	}
	if v, f := p.processes[key]; f {
		return v
	}
	process := synchronization.New(p.puaseLocation())
	p.processes[key] = process
	return process
}

//Immobilizer ...
func (p *Processes) Immobilizer(_index int, _trigger string) interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := fmt.Sprintf("immo%v%v", _index, _trigger)
	if v, f := p.processes[key]; f {
		return v
	}
	process := immobilizer.New(_index, _trigger)
	p.processes[key] = process
	return process
}

//All ...
func (p *Processes) All() []interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	processes := []interfaces.IProcess{}
	for _, proc := range p.processes {
		processes = append(processes, proc)
	}
	return processes
}

func (p *Processes) puaseLocation() context.CancelFunc {
	return p.pauseProcesses("location", "location_request")
}

func (p *Processes) appendToPaused(_key string) {
	if _, f := p.paused[_key]; f {
		p.paused[_key]++
	} else {
		p.paused[_key] = 1
	}
}

func (p *Processes) pauseProcesses(_keys ...string) context.CancelFunc {
	cnt, cancelFunc := context.WithCancel(context.Background())
	paused := make(map[string]interfaces.IProcess)
	for _, key := range _keys {
		p.appendToPaused(key)
		if process, f := p.processes[key]; f {
			process.Pause()
		}
	}
	go p.listen(cnt, paused)
	return cancelFunc
}

func (p *Processes) listen(_context context.Context, _processes map[string]interfaces.IProcess) {
	select {
	case <-_context.Done():
		{
			p.mutex.Lock()
			defer p.mutex.Unlock()
			for key, process := range _processes {
				p.paused[key]--
				if p.paused[key] == 0 {
					delete(p.paused, key)
					process.Resume()
				}
			}
			return
		}
	}
}
