package manager

import (
	"context"
	"fmt"
	"geometris-go/core/interfaces"
	"geometris-go/core/processes/configuration"
	"geometris-go/core/processes/immobilizer"
	"geometris-go/core/processes/message"
	"geometris-go/core/processes/synchronization"
	"geometris-go/logger"
	"sync"
)

func New(_syncParam string) interfaces.IProcesses {
	manager := &Manager{
		processes: make(map[string]interfaces.IProcess),
		mutex:     &sync.Mutex{},
		paused:    make(map[string]int),
	}
	if _syncParam == "" {
		manager.NewSynchProcess()
	} else {
		manager.NewLocationProcess(_syncParam)
	}
	return manager
}

type Manager struct {
	processes map[string]interfaces.IProcess
	mutex     *sync.Mutex
	paused    map[string]int
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
		proc = immobilizer.New(index, trigger)
	}
	return proc
}

func (p *Manager) LocationRequest() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "location_request"
	return p.getOrCreateProcess(key, configuration.New)
}

func (p *Manager) NewSynchProcess() {
	p.Synchronization()
}

func (p *Manager) NewLocationProcess(_syncParam string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "location"
	process := message.New(_syncParam)
	if _, f := p.paused[key]; f {
		process.Pause()
	}
	p.processes[key] = process
}

func (p *Manager) Configuration() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "configuration"
	return p.getOrCreateProcess(key, configuration.New)
}

func (p *Manager) Synchronization() interfaces.IProcess {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	key := "synch"
	return p.getOrCreateProcess(key, synchronization.New)
}

func (p *Manager) puaseLocation() context.CancelFunc {
	return p.pauseProcesses("location", "location_request")
}

func (p *Manager) appendToPaused(_key string) {
	if _, f := p.paused[_key]; f {
		p.paused[_key]++
	} else {
		p.paused[_key] = 1
	}
}

func (p *Manager) getOrCreateProcess(_key string, _constructor func() interfaces.IProcess) interfaces.IProcess {
	if _, f := p.paused[_key]; f {
		logger.Logger().WriteToLog(logger.Info, "[Processes] Process "+_key+" is paused")
		return nil
	}
	process, f := p.processes[_key]
	if !f {
		process = _constructor()
		p.processes[_key] = process
	}
	switch _key {
	case "configuration", "synch":
		{
			process.NewFuncOnEnd(p.puaseLocation())
		}
	}
	return process
}

func (p *Manager) pauseProcesses(_keys ...string) context.CancelFunc {
	cnt, cancelFunc := context.WithCancel(context.Background())
	paused := []string{}
	for _, key := range _keys {
		p.appendToPaused(key)
		if process, f := p.processes[key]; f {
			process.Pause()
			paused = append(paused, key)
		}
	}
	go p.listen(cnt, paused)
	return cancelFunc
}

func (p *Manager) listen(_context context.Context, _keys []string) {
	select {
	case <-_context.Done():
		{
			p.mutex.Lock()
			defer p.mutex.Unlock()
			for _, key := range _keys {
				p.paused[key]--
				if p.paused[key] == 0 {
					p.processes[key].Resume()
				}
			}
			return
		}
	}
}
