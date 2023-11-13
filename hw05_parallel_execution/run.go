package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type TaskResult struct {
	Err error
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, maxProcessCount, maxErrorCount int) error {
	if maxErrorCount <= 0 {
		return ErrErrorsLimitExceeded
	}

	taskCh := make(chan Task, len(tasks))

	func() {
		defer close(taskCh)
		for i := range tasks {
			taskCh <- tasks[i]
		}
	}()

	errorCount := 0
	var mainErr error

	wg := new(sync.WaitGroup)
	mtx := new(sync.Mutex)

	for i := 0; i < maxProcessCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for task := range taskCh {
				err := task()
				if err != nil {
					mtx.Lock()
					errorCount++
					mtx.Unlock()
					if errorCount >= maxErrorCount {
						mainErr = ErrErrorsLimitExceeded
						return
					}
				}
			}
		}()
	}

	wg.Wait()

	return mainErr
}
