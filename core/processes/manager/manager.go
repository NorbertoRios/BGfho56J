package manager

import (
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/immobilizer"
	"geometris-go/core/processes/message"
	"geometris-go/core/processes/synchronization"
	"sync"
)

//BuildProcesses ...
func BuildProcesses(_syncParam string) interfaces.IProcesses {
	manager := &Processes{
		mutex:     &sync.Mutex{},
		processes: make(map[string]interfaces.IProcess),
	}
	if _syncParam == "" {
		manager.addProcess("synch", synchronization.New())
	} else {
		manager.addProcess("location", message.New(_syncParam))
	}

	return manager
}

//Processes ...
type Processes struct {
	processes map[string]interfaces.IProcess
	mutex     *sync.Mutex
}

func (p *Processes) addProcess(_key string, _process interfaces.IProcess) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.processes[_key] = _process
}

//Synchronization ...
func (p *Processes) Synchronization() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "synch"
	if v, f := p.processes[key]; f {
		return v
	}
	process := synchronization.New()
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
