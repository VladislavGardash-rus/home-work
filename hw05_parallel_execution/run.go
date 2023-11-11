package hw05parallelexecution

import (
	"errors"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, maxProcessCount, maxErrorCount int) error {
	taskCh := make(chan Task)
	closeCh := make(chan struct{})

	defer func() {
		close(taskCh)
		close(closeCh)
	}()

	errorCount := 0

	wg := new(sync.WaitGroup)
	wg.Add(len(tasks))

	mtx := new(sync.Mutex)

	for i := 0; i < maxProcessCount; i++ {
		go func(taskCh chan Task, wg *sync.WaitGroup, mtx *sync.Mutex, errorCount *int) {
			for {
				select {
				case task := <-taskCh:
					if task != nil {
						err := task()
						mtx.Lock()
						if err != nil {
							*errorCount++
						}
						mtx.Unlock()
						wg.Done()
					}
				case <-closeCh:
					return
				}
			}
		}(taskCh, wg, mtx, &errorCount)
	}

	for i := range tasks {
		if errorCount >= maxErrorCount {
			for i := 0; i < maxProcessCount; i++ {
				closeCh <- struct{}{}
			}

			time.Sleep(500 * time.Millisecond)
			return ErrErrorsLimitExceeded
		}

		taskCh <- tasks[i]
	}

	wg.Wait()

	return nil
}
