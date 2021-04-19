package manager

import (
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/bocommand"
	"geometris-go/core/processes/configuration"
	"geometris-go/core/processes/immobilizer"
	"geometris-go/core/processes/location"
	"geometris-go/core/processes/message"
	"geometris-go/core/processes/synchronization"
	"geometris-go/logger"
	"sync"
)

//New ...
func New(_syncParams map[string]string) interfaces.IProcesses {
	manager := &Manager{
		processes: make(map[string]interfaces.IProcess),
		pauseMap:  make(map[string][]string),
		mutex:     &sync.Mutex{},
		paused:    make(map[string]int),
	}
	manager.newLocationProcess(_syncParams)
	return manager
}

//Manager ...
type Manager struct {
	processes map[string]interfaces.IProcess
	pauseMap  map[string][]string
	mutex     *sync.Mutex
	paused    map[string]int
}

//NewSynchParameter ...
func (p *Manager) NewSynchParameter(crc, syncParam string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "location"
	if process, f := p.processes[key]; f {
		process.Current().(interfaces.ILocationTask).NewSyncParam(crc, syncParam)
	}
}

//All ...
func (p *Manager) All() []interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	processes := []interfaces.IProcess{}
	for _, proc := range p.processes {
		processes = append(processes, proc)
	}
	return processes
}

//RemoveProcess ...
func (p *Manager) RemoveProcess(_key string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.processes, _key)
}

//BOCommand ...
func (p *Manager) BOCommand(_command string) interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := fmt.Sprintf("boCommand_%v", _command)
	if _, f := p.paused[key]; f {
		logger.Logger().WriteToLog(logger.Info, "[Processes] Process "+key+" is paused")
		return nil
	}
	proc := bocommand.New(key)
	p.processes[key] = proc
	return proc
}

//Immobilizer ...
func (p *Manager) Immobilizer(index int, trigger string) interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := fmt.Sprintf("immo_%v_%v", index, trigger)
	if _, f := p.paused[key]; f {
		logger.Logger().WriteToLog(logger.Info, "[Processes] Process "+key+" is paused")
		return nil
	}
	proc, f := p.processes[key]
	if !f {
		proc = immobilizer.New(index, trigger, key)
		p.processes[key] = proc
	}
	return proc
}

//LocationRequest ...
func (p *Manager) LocationRequest() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "location_request"
	return p.getOrCreateProcess(key, location.New)
}

func (p *Manager) newLocationProcess(_syncParams map[string]string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "location"
	process := message.New(key, _syncParams)
	if _, f := p.paused[key]; f {
		process.Pause()
	}
	p.processes[key] = process
}

//Configuration ...
func (p *Manager) Configuration() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "configuration"
	return p.getOrCreateProcess(key, configuration.New)
}

//Synchronization ...
func (p *Manager) Synchronization() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "synch"
	return p.getOrCreateProcess(key, synchronization.New)
}

func (p *Manager) appendToPaused(_key string) {
	if _, f := p.paused[_key]; f {
		p.paused[_key]++
	} else {
		p.paused[_key] = 1
	}
}

func (p *Manager) getOrCreateProcess(_key string, _constructor func(string) interfaces.IProcess) interfaces.IProcess {
	if _, f := p.paused[_key]; f {
		logger.Logger().WriteToLog(logger.Info, "[Processes] Process "+_key+" is paused")
		return nil
	}
	process, f := p.processes[_key]
	if !f {
		process = _constructor(_key)
		p.processes[_key] = process
	}
	switch _key {
	case "configuration", "synch":
		{
			p.pauseProcesses(_key, "location_request")
		}
	}
	return process
}

func (p *Manager) pauseProcesses(_processKey string, _keys ...string) {
	if _, f := p.pauseMap[_processKey]; !f {
		p.pauseMap[_processKey] = []string{}
	}
	p.pauseMap[_processKey] = append(p.pauseMap[_processKey], _keys...)
	for _, key := range _keys {
		p.appendToPaused(key)
		if process, f := p.processes[key]; f {
			process.Pause()
		}
	}
}

//ProcessComplete ...
func (p *Manager) ProcessComplete(_processKey string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	processes, f := p.pauseMap[_processKey]
	if !f {
		return
	}
	for _, pr := range processes {
		if _, f := p.paused[pr]; !f {
			continue
		}
		p.paused[pr]--
		if p.paused[pr] == 0 {
			delete(p.paused, pr)
			if p, f := p.processes[pr]; f {
				p.Resume()
			}
		}
	}
}
