package workerpool

import (
	"context"
	"sync"
)

type WorkerPool struct {
	taskQueue     chan Task
	ResultChan    chan Result
	maxConcurrent int
	Wg            *sync.WaitGroup
}

func NewWorkerPool(maxConcurrent int) *WorkerPool {
	return &WorkerPool{
		taskQueue:     make(chan Task),
		ResultChan:    make(chan Result),
		maxConcurrent: maxConcurrent,
		Wg:            &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < wp.maxConcurrent; i++ {
		Worker := NewWorker(wp.taskQueue, wp.ResultChan, wp.Wg)
		Worker.Start(ctx)
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.Wg.Add(1)
	wp.taskQueue <- task
}

func (wp *WorkerPool) CloseTaskQueue() {
	close(wp.taskQueue)
}
