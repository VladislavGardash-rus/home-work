package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, maxProcessCount, maxErrorCount int) error {
	taskCh := make(chan Task)
	defer close(taskCh)

	errorCount := 0
	wg := new(sync.WaitGroup)
	wg.Add(len(tasks))
	mtx := new(sync.Mutex)
	for i := 0; i < maxProcessCount; i++ {
		go func(taskCh chan Task, wg *sync.WaitGroup, mtx *sync.Mutex, errorCount *int) {
			for task := range taskCh {
				err := task()
				if err != nil {
					mtx.Lock()
					*errorCount++
					mtx.Unlock()
				}
				wg.Done()
			}
		}(taskCh, wg, mtx, &errorCount)
	}

	for i := range tasks {
		if errorCount == maxErrorCount {
			return ErrErrorsLimitExceeded
		}

		taskCh <- tasks[i]
	}

	wg.Wait()

	return nil
}
