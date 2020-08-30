package gopool

import (
	"sync"
)

type Pool struct {
	limit int
	jobs chan func()
	wg sync.WaitGroup
}

const defaultLimit = 10

func New(limit int) *Pool {
	if limit <= 0 {
		limit = defaultLimit
	}
	pool := Pool{limit:limit, jobs:make(chan func()), wg:sync.WaitGroup{}}

	pool.wg.Add(limit)
	for i := 0; i < limit; i++ {
		go pool.do()
	}

	return &pool
}

func (p *Pool) Do(job func()) {
	p.jobs <- job
}

func (p *Pool) Done() {
	for i := 0; i < p.limit; i++ {
		p.Do(nil)
	}
	p.limit = 0
	p.wg.Wait()
}

func (p *Pool) do() {
	for job := range p.jobs {
		if job == nil {
			p.wg.Done()
			return
		}
		job()
	}
}