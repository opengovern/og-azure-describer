package workerpool

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
)

type TestTask struct {
	TaskProperties
}

func (d TestTask) Properties() TaskProperties {
	return d.TaskProperties
}

func (d TestTask) Run(_ context.Context) error {

	log.Println("Executing task")
	// time.Sleep(2 * time.Second)
	return nil
}

func TestPool(t *testing.T) {

	pool := NewWorkerPool(1)

	pool.Start(context.Background())

	for range 10 {
		pool.AddTask(
			TestTask{
				TaskProperties: TaskProperties{
					ID:          uuid.New(),
					Description: "Test Task",
				},
			},
		)
	}

	// pool.wg.Wait()

}
