package pool

// Worker Worker struct
type Worker struct {
	isFree   bool
	taskChan chan func()
}

// create a worker
func newWorker() *Worker {
	worker := new(Worker)
	worker.isFree = true
	worker.taskChan = make(chan func(), 1)
	go worker.run()
	return worker
}

// run task
// as push a task
func (w *Worker) runTask(task func()) bool {

	if !w.isFree {
		return false
	}

	// push a task
	w.taskChan <- task
	return true
}

// run task（real）
func (w *Worker) run() {
	for {
		task := <-w.taskChan
		w.isFree = false
		task()
		w.isFree = true
	}

}
