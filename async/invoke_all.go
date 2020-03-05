package async

import "context"

func InvokeAll(ctx context.Context, concurrency int, tasks []Task) Task{
	if concurrency == 0{
		return ForkJoin(ctx, tasks)
	}

	return Invoke(ctx, func(context.Context) (i interface{}, e error) {
		sem := make(chan struct{}, concurrency)
		for _, task := range tasks{
			sem <- struct{}{}
			task.Run(ctx).ContinueWith(ctx, func( interface{}, error) (interface{}, error) {
				<-sem
				return nil, nil
			})
		}
		WaitAll(tasks)
		return nil, nil
	})
}