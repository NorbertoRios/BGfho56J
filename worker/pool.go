package worker

import "geometris-go/interfaces/unitofwork"

//NewPool ...
func NewPool(workersCount int, _uow unitofwork.IUnitOfWork) *Pool {
	_workers := []IWorker{}
	for i := 0; i < workersCount; i++ {
		w := NewWorker(_uow)
		_workers = append(_workers, w)
	}
	return &Pool{
		currentNum: 0,
		Workers:    _workers,
	}
}

//Pool ...
type Pool struct {
	currentNum int
	Workers    []IWorker
}

func (p *Pool) all() []IWorker {
	return p.Workers
}

func (p *Pool) next() IWorker {
	defer func() { p.currentNum++ }()
	if p.currentNum == len(p.Workers) {
		p.currentNum = 0
	}
	return p.Workers[p.currentNum]
}
