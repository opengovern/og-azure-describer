package workerpool

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

type Worker struct {
	id        uuid.UUID
	taskQueue <-chan Task
	results   chan<- Result
	wg        *sync.WaitGroup
}

func NewWorker(
	taskQueue <-chan Task,
	results chan<- Result,
	wg *sync.WaitGroup,
) Worker {
	return Worker{
		id:        uuid.New(),
		taskQueue: taskQueue,
		results:   results,
		wg:        wg,
	}
}

func (w *Worker) Start(ctx context.Context) {
	go func() {
		for task := range w.taskQueue {
			err := task.Run(ctx)
			w.results <- *NewResult(
				task.Properties().ID,
				err,
			)
			w.wg.Done()
		}
		// close(w.results)
	}()
}
