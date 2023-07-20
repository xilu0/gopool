package gopool

import (
	"fmt"
	"sync"
)

type workerFunc func() error

type Pool interface {
	Add(fn workerFunc)
	Run()
	Stop()
}

type pool struct {
	workers chan workerFunc
	wg      sync.WaitGroup
	limit   chan struct{}
}

func NewPool(size int) Pool {
	return &pool{
		workers: make(chan workerFunc, size),
		limit:   make(chan struct{}, size),
	}
}

func (p *pool) Add(fn workerFunc) {
	p.workers <- fn
}

func (p *pool) Run() {
	for fn := range p.workers {
		p.limit <- struct{}{}
		p.wg.Add(1)
		go func(fn workerFunc) {
			defer func() {
				<-p.limit
				p.wg.Done()
			}()
			if err := fn(); err != nil {
				// handle error
				fmt.Printf("err: %v\n", err)
			}
		}(fn)
	}
	p.wg.Wait()
}

func (p *pool) Stop() {
	close(p.workers)
	p.wg.Wait()
}
