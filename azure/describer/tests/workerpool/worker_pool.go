package workerpool

import (
	"context"
	"sync"
)

type WorkerPool struct {
	taskQueue     chan Task
	resultChan    chan Result
	maxConcurrent int
	Wg            *sync.WaitGroup
}

func NewWorkerPool(maxConcurrent int) *WorkerPool {
	return &WorkerPool{
		taskQueue:     make(chan Task),
		resultChan:    make(chan Result),
		maxConcurrent: maxConcurrent,
		Wg:            &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < wp.maxConcurrent; i++ {
		worker := NewWorker(wp.taskQueue, wp.resultChan, wp.Wg)
		worker.Start(ctx)
	}
	resultWorker := NewResultWorker(wp.resultChan, wp.Wg)
	resultWorker.Start(ctx)
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.Wg.Add(2) // task and result
	wp.taskQueue <- task
}

// func (wp *WorkerPool) Wait() {
// 	wp.wg.Wait()
// }
