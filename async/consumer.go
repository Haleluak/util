package async

import (
	"context"
	"runtime"
)

func Consumer(ctx context.Context, concurrency int, task chan Task)  Task{
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}

	return Invoke(ctx, func(taskCtx context.Context) (interface{}, error) {
		workers := make(chan int, concurrency)
		concurrentTasks := make([]Task, concurrency)

		for id := 0; id < concurrency; id++ {
			workers <- id
		}

		for{
			select {
			case <-taskCtx.Done():
				WaitAll(concurrentTasks)
				return nil, taskCtx.Err()

			case workerID := <- workers:
				select {
				case <-taskCtx.Done():
					WaitAll(concurrentTasks)
					return nil, taskCtx.Err()
				case t, ok := <-task:
					if !ok {
						WaitAll(concurrentTasks)
						return nil, nil
					}

					concurrentTasks[workerID] = t
					t.Run(taskCtx).ContinueWith(taskCtx, func( interface{},  error) (interface{}, error) {
						workers <- workerID
						return nil, nil
					})
				}
			}
		}
		return nil, nil
	})
}
