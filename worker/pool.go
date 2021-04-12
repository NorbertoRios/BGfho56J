package worker

import (
	"container/list"
	"geometris-go/repository"
	"sync"
)

//NewPool ...
func NewPool(workersCount int, _mysql, rabbit repository.IRepository) *Pool {
	_workers := list.New()
	for i := 0; i < workersCount; i++ {
		w := NewWorker(_mysql, rabbit)
		go w.Run()
		_workers.PushBack(w)
	}
	return &Pool{
		current: _workers.Front(),
		workers: _workers,
		mutex:   &sync.Mutex{},
	}
}

//Pool ...
type Pool struct {
	current *list.Element
	workers *list.List
	mutex   *sync.Mutex
}

func (p *Pool) all() []IWorker {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	workers := []IWorker{}
	for w := p.workers.Front(); w != nil; w = w.Next() {
		workers = append(workers, w.Value.(IWorker))
	}
	return workers
}

func (p *Pool) next() IWorker {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	defer func() {
		if v := p.current.Next(); v != nil {
			p.current = p.current.Next()
		} else {
			p.current = p.workers.Front()
		}
	}()
	return p.current.Value.(IWorker)
}
