package hw06pipelineexecution

import "sync"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for {
			select {
			case value := <-in:
				if value == nil {
					wg.Done()
					return
				}

				wg.Add(1)

				for _, stage := range stages {
					newIn := make(Bi)
					go func() {
						newIn <- value
					}()

					value = <-stage(newIn)

					close(newIn)
				}

				out <- value
				wg.Done()

			case <-done:
				close(out)
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
