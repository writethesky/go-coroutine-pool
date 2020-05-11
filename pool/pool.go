package pool

import (
	"time"
)

// New create coroutine pool
// len coroutine pool length
func New(len int) *Pool {
	pool := new(Pool)
	pool.create(len)
	return pool
}

// Pool coroutine pool struct
type Pool struct {
	// 容量
	len     int
	workers []*Worker
}

// create coroutine pool
func (p *Pool) create(len int) {
	p.len = len
}

// AddTask add a task
// task a task as function
func (p *Pool) AddTask(task func()) {

	// find free worker
	for _, worker := range p.workers {
		// finded a free worker and run task
		if worker.runTask(task) {
			return
		}
	}

	// not found free worker
	// pool has free
	if p.len > len(p.workers) {

		// create a worker
		worker := p.newWorker()

		// run a task
		worker.runTask(task)
		return
	}

	// pool is full
	// await worker free
	worker := p.untilFindFreeWorker()
	// run a task
	worker.runTask(task)
}

// create a worker
func (p *Pool) newWorker() *Worker {
	worker := newWorker()
	p.workers = append(p.workers, worker)
	return worker
}

// find worker，until
func (p *Pool) untilFindFreeWorker() *Worker {
	for _, worker := range p.workers {
		if worker.isFree {
			// finded a free worker
			return worker
		}
	}

	time.Sleep(time.Second)

	return p.untilFindFreeWorker()
}
